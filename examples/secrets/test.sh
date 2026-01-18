#!/bin/bash
set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# Project root is two levels up from examples/secrets
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

echo -e "${GREEN}================================${NC}"
echo -e "${GREEN}Sensu Secrets Testing Script${NC}"
echo -e "${GREEN}================================${NC}"
echo ""

# Function to print section headers
section() {
    echo ""
    echo -e "${YELLOW}>>> $1${NC}"
    echo ""
}

# Function to wait for service
wait_for_service() {
    local url=$1
    local max_attempts=30
    local attempt=0

    section "Waiting for Sensu backend to be ready..."

    while [ $attempt -lt $max_attempts ]; do
        if curl -s -f "${url}/health" > /dev/null 2>&1; then
            echo -e "${GREEN}Backend is ready!${NC}"
            return 0
        fi
        attempt=$((attempt + 1))
        echo "Attempt $attempt/$max_attempts - waiting..."
        sleep 2
    done

    echo -e "${RED}Backend failed to start after $max_attempts attempts${NC}"
    return 1
}

# Parse command line arguments
ACTION=${1:-"all"}

case $ACTION in
    "start")
        section "Starting Sensu environment..."
        cd "$SCRIPT_DIR"

        # Stop any existing containers first to avoid port conflicts
        echo "Stopping any existing containers..."
        docker-compose down 2>/dev/null || true

        # Also stop containers from root docker-compose if they exist
        if [ -f "$PROJECT_ROOT/docker-compose.yaml" ]; then
            (cd "$PROJECT_ROOT" && docker-compose down 2>/dev/null) || true
        fi

        # Check if ports are still in use and warn
        for port in 2379 2380 8080 8081 3000; do
            if lsof -i ":$port" >/dev/null 2>&1; then
                echo -e "${YELLOW}Warning: Port $port is in use. Attempting to identify...${NC}"
                lsof -i ":$port" | head -3
            fi
        done

        docker-compose up -d
        wait_for_service "http://127.0.0.1:8080"

        section "Initializing admin user..."
        echo "Using default credentials: admin / P@ssw0rd!"
        # sensu-backend init returns exit code 3 if already initialized, which is not an error
        docker-compose exec -T backend1 sensu-backend init <<EOF || [ $? -eq 3 ]
admin
P@ssw0rd!
EOF

        section "Verifying environment variables..."
        docker-compose exec backend1 env | grep -E "(SLACK_WEBHOOK_URL|PAGERDUTY_API_KEY|MONITORING_API_KEY)" || true

        echo ""
        echo -e "${GREEN}Sensu environment is ready!${NC}"
        echo "API URL: http://127.0.0.1:8080"
        echo "Username: admin"
        echo "Password: P@ssw0rd!"
        ;;

    "build")
        section "Building Terraform provider..."
        cd "$PROJECT_ROOT"
        go install

        section "Installing provider to local plugin directory..."
        # Detect OS and architecture
        OS=$(uname -s | tr '[:upper:]' '[:lower:]')
        ARCH=$(uname -m)

        case $OS in
            darwin)
                OS_NAME="darwin"
                ;;
            linux)
                OS_NAME="linux"
                ;;
            *)
                echo -e "${RED}Unsupported OS: $OS${NC}"
                exit 1
                ;;
        esac

        case $ARCH in
            x86_64)
                ARCH_NAME="amd64"
                ;;
            arm64|aarch64)
                ARCH_NAME="arm64"
                ;;
            *)
                echo -e "${RED}Unsupported architecture: $ARCH${NC}"
                exit 1
                ;;
        esac

        PLUGIN_DIR="$HOME/.terraform.d/plugins/registry.terraform.io/jtopjian/sensu/0.15.0/${OS_NAME}_${ARCH_NAME}"

        mkdir -p "$PLUGIN_DIR"
        # Find the binary - check GOBIN first, then GOPATH/bin, then common locations
        GOBIN_PATH=$(go env GOBIN)
        GOPATH_PATH=$(go env GOPATH)
        if [ -n "$GOBIN_PATH" ] && [ -f "$GOBIN_PATH/terraform-provider-sensu" ]; then
            PROVIDER_BIN="$GOBIN_PATH/terraform-provider-sensu"
        elif [ -n "$GOPATH_PATH" ] && [ -f "$GOPATH_PATH/bin/terraform-provider-sensu" ]; then
            PROVIDER_BIN="$GOPATH_PATH/bin/terraform-provider-sensu"
        elif [ -f "$HOME/go/bin/terraform-provider-sensu" ]; then
            PROVIDER_BIN="$HOME/go/bin/terraform-provider-sensu"
        else
            echo -e "${RED}Could not find terraform-provider-sensu binary${NC}"
            echo "Checked: $GOBIN_PATH, $GOPATH_PATH/bin, $HOME/go/bin"
            exit 1
        fi
        cp "$PROVIDER_BIN" "$PLUGIN_DIR/"
        chmod +x "$PLUGIN_DIR/terraform-provider-sensu"

        echo -e "${GREEN}Provider installed to: $PLUGIN_DIR${NC}"
        ;;

    "apply")
        section "Applying Terraform configuration..."
        cd "$SCRIPT_DIR"

        # Remove stale lock file to avoid checksum mismatch after rebuild
        rm -f .terraform.lock.hcl
        terraform init
        terraform plan
        terraform apply -auto-approve

        echo ""
        echo -e "${GREEN}Resources created successfully!${NC}"
        ;;

    "verify")
        section "Verifying created resources..."
        cd "$SCRIPT_DIR"

        echo "Configuring sensuctl..."
        docker-compose exec -T backend1 sensuctl configure -n \
          --url http://127.0.0.1:8080 \
          --username admin \
          --password 'P@ssw0rd!' \
          --namespace development

        echo ""
        echo "=== Namespaces ==="
        docker-compose exec backend1 sensuctl namespace list

        echo ""
        echo "=== Secrets ==="
        docker-compose exec backend1 sensuctl secret list --namespace development

        echo ""
        echo "=== Checks ==="
        docker-compose exec backend1 sensuctl check list --namespace development

        echo ""
        echo "=== Handlers ==="
        docker-compose exec backend1 sensuctl handler list --namespace development

        echo ""
        echo "=== Entities ==="
        docker-compose exec backend1 sensuctl entity list --namespace development

        echo ""
        echo "=== API Health Check (with secrets) ==="
        curl -s -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/checks/api-health-check | jq -r '.secrets'
        ;;

    "destroy")
        section "Destroying Terraform resources..."
        cd "$SCRIPT_DIR"
        terraform destroy -auto-approve

        echo -e "${GREEN}Resources destroyed!${NC}"
        ;;

    "stop")
        section "Stopping Sensu environment..."
        cd "$SCRIPT_DIR"
        docker-compose down

        echo -e "${GREEN}Environment stopped!${NC}"
        ;;

    "clean")
        section "Cleaning up everything..."
        cd "$SCRIPT_DIR"
        terraform destroy -auto-approve || true
        docker-compose down -v
        rm -f .terraform.lock.hcl
        rm -rf .terraform

        echo -e "${GREEN}Everything cleaned up!${NC}"
        ;;

    "all"|"test")
        section "Running full test workflow..."

        # Start
        "$0" start

        # Build
        "$0" build

        # Apply
        "$0" apply

        # Verify
        "$0" verify

        echo ""
        echo -e "${GREEN}================================${NC}"
        echo -e "${GREEN}Full test completed successfully!${NC}"
        echo -e "${GREEN}================================${NC}"
        echo ""
        echo "To clean up, run: $0 clean"
        ;;

    "help"|*)
        echo "Usage: $0 [command]"
        echo ""
        echo "Commands:"
        echo "  all       - Run full test workflow (default)"
        echo "  start     - Start Sensu backend and initialize admin user"
        echo "  build     - Build and install Terraform provider locally"
        echo "  apply     - Run terraform init/plan/apply"
        echo "  verify    - Verify resources were created correctly"
        echo "  destroy   - Destroy Terraform resources"
        echo "  stop      - Stop Sensu environment (keeps data)"
        echo "  clean     - Destroy resources and stop environment (removes all data)"
        echo "  help      - Show this help message"
        echo ""
        echo "Examples:"
        echo "  $0              # Run full test"
        echo "  $0 start        # Just start Sensu"
        echo "  $0 clean        # Clean up everything"
        ;;
esac
