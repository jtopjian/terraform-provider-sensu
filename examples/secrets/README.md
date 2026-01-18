# Sensu Secrets Testing Guide

This directory contains a comprehensive example for testing the Sensu secrets resource implementation. This guide will walk you through setting up a local Sensu environment and testing the secrets functionality end-to-end.

## Overview

This example demonstrates:
- Creating a custom namespace
- Defining an entity (agent)
- Creating secrets using the `env` provider (built-in to Sensu Go)
- Using secrets in checks
- Using secrets in handlers
- Reading secrets via data sources

## Prerequisites

- Docker and Docker Compose
- Go 1.14 or later
- Terraform 0.12 or later
- Make (optional, but recommended)

## Understanding Sensu Secrets

### The Env Secrets Provider

The `env` secrets provider is **automatically created** when you start a Sensu backend. It exposes secrets from environment variables on your Sensu backend nodes. You do NOT need to create the provider itself - it's built into Sensu Go commercial distribution.

Key points:
- Provider name is always `env`
- Secret `id` corresponds to the environment variable name on the backend
- Secret `name` is how you reference it in Terraform and other Sensu resources
- The actual secret values are stored as environment variables on the Sensu backend

### How Secrets Work in Checks and Handlers

When you reference a secret in a check or handler:

```hcl
resource "sensu_check" "my_check" {
  name    = "my-check"
  command = "check-api.rb --api-key $API_KEY"

  secrets = {
    API_KEY = "my-api-key-secret"  # Name of the sensu_secret resource
  }
}
```

The `secrets` map defines:
- **Key**: Environment variable name that will be available to the check/handler command
- **Value**: Name of the Sensu secret resource (not the environment variable on the backend)

## Step-by-Step Testing Instructions

### Step 1: Start the Sensu Backend

The `docker-compose.yaml` file in the project root is configured with environment variables for testing secrets.

```bash
cd <project-root>

# Start the Sensu backend and agent
docker-compose up -d

# Verify containers are running
docker-compose ps
```

Expected output:
```
NAME                COMMAND                  SERVICE             STATUS              PORTS
backend1            "sensu-backend start…"   backend1            running             0.0.0.0:2379-2380->2379-2380/tcp, ...
agent1              "sensu-agent start -…"   agent1              running
```

### Step 2: Initialize the Sensu Admin User

Wait a few seconds for the backend to start, then initialize the admin user:

```bash
# Initialize the admin user (this only needs to be done once)
docker-compose exec backend1 sensu-backend init

# When prompted, enter:
# Username: admin
# Password: P@ssw0rd!
```

Alternative one-liner (non-interactive):
```bash
docker-compose exec backend1 sensu-backend init <<EOF
admin
P@ssw0rd!
EOF
```

### Step 3: Verify Sensu Backend is Ready

Test API connectivity:

```bash
# Check if API is responding
curl -i http://127.0.0.1:8080/health

# Expected response: HTTP/1.1 200 OK

# Verify the env secrets provider exists (it's auto-created)
curl -u admin:P@ssw0rd! http://127.0.0.1:8080/api/enterprise/secrets/v1/providers

# Expected: You should see the "env" provider in the response
```

### Step 4: Verify Environment Variables on Backend

Confirm the environment variables are set on the backend:

```bash
docker-compose exec backend1 env | grep -E "(SLACK_WEBHOOK_URL|PAGERDUTY_API_KEY|MONITORING_API_KEY)"
```

Expected output:
```
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/YOUR/WEBHOOK/URL
PAGERDUTY_API_KEY=your-pagerduty-api-key-here
MONITORING_API_KEY=your-monitoring-api-key-here
```

### Step 5: Build and Install the Terraform Provider Locally

Build the provider from source:

```bash
cd <project-root>

# Build the provider (installs to $GOPATH/bin or ~/go/bin)
make build

# Verify the binary was created
ls -lh ~/go/bin/terraform-provider-sensu
```

For Terraform 0.13+, you need to install the provider to the local plugin directory:

```bash
# Create the local provider directory
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/jtopjian/sensu/0.15.0/darwin_amd64

# Copy the built binary (adjust architecture as needed)
cp ~/go/bin/terraform-provider-sensu ~/.terraform.d/plugins/registry.terraform.io/jtopjian/sensu/0.15.0/darwin_amd64/

# Make it executable
chmod +x ~/.terraform.d/plugins/registry.terraform.io/jtopjian/sensu/0.15.0/darwin_amd64/terraform-provider-sensu
```

**Note**: Adjust the path based on your OS and architecture:
- macOS (Intel): `darwin_amd64`
- macOS (Apple Silicon): `darwin_arm64`
- Linux: `linux_amd64`
- Windows: `windows_amd64`

### Step 6: Initialize Terraform

Navigate to the examples directory and initialize Terraform:

```bash
cd <project-root>/examples/secrets

# Initialize Terraform (downloads required providers)
terraform init
```

Expected output:
```
Initializing the backend...

Initializing provider plugins...
- Finding latest version of registry.terraform.io/jtopjian/sensu...
- Installing registry.terraform.io/jtopjian/sensu v0.15.0...

Terraform has been successfully initialized!
```

### Step 7: Review the Terraform Plan

Preview what Terraform will create:

```bash
terraform plan
```

You should see plans to create:
- 1 namespace (`sensu_namespace.dev`)
- 1 entity (`sensu_entity.test_agent`)
- 3 secrets (`sensu_secret.slack_webhook`, `sensu_secret.pagerduty_key`, `sensu_secret.monitoring_api_key`)
- 2 checks (`sensu_check.http_check`, `sensu_check.api_check`)
- 1 handler (`sensu_handler.slack`)

### Step 8: Apply the Configuration

Create the resources:

```bash
terraform apply

# When prompted, type 'yes' to confirm
```

Expected output:
```
Apply complete! Resources: 8 added, 0 changed, 0 destroyed.

Outputs:

checks = [
  "http-health-check",
  "api-health-check",
]
entity_name = "test-agent"
handler = "slack-alerts"
namespace = "development"
secrets = [
  "slack-webhook-url",
  "pagerduty-api-key",
  "monitoring-api-key",
]
```

### Step 9: Verify Resources Were Created

Use the Sensu CLI or API to verify the resources:

```bash
# List namespaces
docker-compose exec backend1 sensuctl namespace list

# Configure sensuctl
docker-compose exec backend1 sensuctl configure -n \
  --url http://127.0.0.1:8080 \
  --username admin \
  --password 'P@ssw0rd!' \
  --namespace development

# List secrets
docker-compose exec backend1 sensuctl secret list --namespace development

# Get secret details
docker-compose exec backend1 sensuctl secret info slack-webhook-url --namespace development

# List checks
docker-compose exec backend1 sensuctl check list --namespace development

# List handlers
docker-compose exec backend1 sensuctl handler list --namespace development

# List entities
docker-compose exec backend1 sensuctl entity list --namespace development
```

Alternatively, use the API:

```bash
# List secrets
curl -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/secrets

# Get a specific secret
curl -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/secrets/slack-webhook-url

# List checks
curl -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/checks

# Get check that uses secrets
curl -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/checks/api-health-check
```

### Step 10: Verify Secret References in Resources

Inspect a check to confirm it references the secret correctly:

```bash
# Get the check configuration as JSON
docker-compose exec backend1 sensuctl check info api-health-check \
  --namespace development --format json | grep -A 5 secrets
```

You should see output like:
```json
"secrets": [
  {
    "name": "monitoring-api-key",
    "secret": "MONITORING_API_KEY"
  }
]
```

This shows that the check has a secret named `MONITORING_API_KEY` that references the `monitoring-api-key` secret resource.

### Step 11: Test Updates

Test updating a secret's ID:

```bash
# Edit main.tf and change one of the secret IDs, e.g.:
# id = "MONITORING_API_KEY" -> id = "MONITORING_API_KEY_V2"

# Plan the changes
terraform plan

# Apply the update
terraform apply
```

### Step 12: Test Import

Test importing an existing secret:

```bash
# First, create a secret manually via API
curl -X POST -u admin:P@ssw0rd! \
  -H 'Content-Type: application/json' \
  -d '{
    "type": "Secret",
    "api_version": "secrets/v1",
    "metadata": {
      "name": "test-import-secret",
      "namespace": "development"
    },
    "secrets_provider": "env",
    "id": "TEST_IMPORT_VAR"
  }' \
  http://127.0.0.1:8080/api/enterprise/secrets/v1/namespaces/development/secrets

# Import it into Terraform state
terraform import sensu_secret.test_import development/test-import-secret

# Note: You need to add the resource block to main.tf first:
# resource "sensu_secret" "test_import" {
#   name      = "test-import-secret"
#   namespace = "development"
#   id        = "TEST_IMPORT_VAR"
#   secrets_provider = "env"
# }
```

### Step 13: Cleanup

Remove all created resources:

```bash
# Destroy Terraform-managed resources
terraform destroy

# When prompted, type 'yes' to confirm

# Stop and remove Docker containers
cd <project-root>
docker-compose down

# Optional: Remove volumes to completely reset Sensu state
docker-compose down -v
```

## Troubleshooting

### Issue: Provider Not Found

If Terraform cannot find the provider:

```
Error: Failed to query available provider packages
```

**Solution**: Ensure you copied the provider binary to the correct plugin directory. For local development, you can also use a `.terraformrc` file:

```bash
cat > ~/.terraformrc <<EOF
provider_installation {
  dev_overrides {
    "registry.terraform.io/jtopjian/sensu" = "<go-bin-path>"
  }
  direct {}
}
EOF
```

Then run `terraform init` again.

### Issue: Backend Not Ready

If API calls fail with connection refused:

```
Error: Unable to connect to Sensu API
```

**Solution**: Wait a few more seconds for the backend to fully start. Check logs:

```bash
docker-compose logs backend1
```

### Issue: Secrets Not Working

If secrets aren't resolving in checks/handlers:

**Check environment variables on backend**:
```bash
docker-compose exec backend1 env | grep SLACK_WEBHOOK_URL
```

**Verify secret exists**:
```bash
docker-compose exec backend1 sensuctl secret list --namespace development
```

**Check secret-to-check mapping**:
```bash
curl -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/checks/api-health-check | jq .secrets
```

### Issue: Authentication Errors

If you get 401 Unauthorized:

**Re-initialize the admin user**:
```bash
docker-compose down -v
docker-compose up -d
docker-compose exec backend1 sensu-backend init
```

## Testing Acceptance Tests

To run the provider's acceptance tests (requires a running Sensu backend):

```bash
cd <project-root>

# Set environment variables for the test
export SENSU_API_URL="http://127.0.0.1:8080"
export SENSU_API_USER="admin"
export SENSU_API_PASS="P@ssw0rd!"

# Run secret-specific tests
make testacc TESTARGS='-run TestAccResourceSecret'

# Run all tests
make testacc
```

## Additional Resources

- [Sensu Secrets Documentation](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets/)
- [Sensu Secrets Providers](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets-providers/)
- [Sensu Secrets Management Guide](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets-management/)
- [Terraform Provider Development](https://www.terraform.io/docs/extend/writing-custom-providers.html)

## Security Notes

1. **Environment Variables**: The `env` provider exposes backend environment variables. In production, ensure your backend environment is properly secured.

2. **Terraform State**: Secret IDs and provider names are stored in Terraform state. The actual secret values are NOT stored in state, but the references are.

3. **Production Use**: For production environments, consider using a dedicated secrets management provider like HashiCorp Vault instead of the `env` provider.

4. **Docker Secrets**: For production Docker deployments, use Docker secrets or Kubernetes secrets instead of environment variables in docker-compose.

## Example Workflow

A typical development/testing workflow:

```bash
# 1. Start environment
docker-compose up -d
docker-compose exec backend1 sensu-backend init

# 2. Build provider
make build

# 3. Test with Terraform
cd examples/secrets
terraform init
terraform plan
terraform apply

# 4. Verify resources
docker-compose exec backend1 sensuctl secret list --namespace development

# 5. Make changes to provider code
# ... edit code ...

# 6. Rebuild and test
make build
terraform plan
terraform apply

# 7. Clean up
terraform destroy
docker-compose down
```

## Notes on the Example Configuration

The `main.tf` file demonstrates:

1. **Namespace isolation**: Creating a custom `development` namespace
2. **Entity definition**: Defining an agent entity with labels and subscriptions
3. **Multiple secrets**: Creating three different secrets using the `env` provider
4. **Secret usage in checks**: The `api_check` uses the `MONITORING_API_KEY` secret
5. **Secret usage in handlers**: The `slack` handler uses the `SLACK_WEBHOOK_URL` secret
6. **Data sources**: Reading back secret information via data sources
7. **Outputs**: Displaying created resource names for verification

The secrets map syntax is:
```hcl
secrets = {
  ENV_VAR_NAME = "secret-resource-name"
}
```

Where:
- `ENV_VAR_NAME`: The environment variable that will be available to the check/handler command
- `"secret-resource-name"`: The `name` attribute of the `sensu_secret` resource
