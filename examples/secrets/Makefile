# Makefile for Sensu Secrets Testing
# Self-contained - all commands run from this directory

.PHONY: help start build apply verify destroy stop clean test all

# Default target
.DEFAULT_GOAL := help

# Get the project root (two levels up)
PROJECT_ROOT := $(shell cd ../.. && pwd)

help: ## Show this help message
	@echo "Sensu Secrets Testing Makefile"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

start: ## Start Sensu backend and initialize admin user
	@./test.sh start

build: ## Build and install Terraform provider locally
	@./test.sh build

apply: ## Run terraform init/plan/apply
	@./test.sh apply

verify: ## Verify resources were created correctly
	@./test.sh verify

destroy: ## Destroy Terraform resources
	@./test.sh destroy

stop: ## Stop Sensu environment (keeps data)
	@./test.sh stop

clean: ## Destroy resources and stop environment (removes all data)
	@./test.sh clean

test: ## Run full test workflow
	@./test.sh all

all: test ## Alias for test

# Docker shortcuts (using local docker-compose.yaml)
docker-up: ## Start docker-compose
	@docker-compose up -d

docker-down: ## Stop docker-compose
	@docker-compose down

docker-logs: ## Show docker-compose logs
	@docker-compose logs -f

docker-ps: ## Show running containers
	@docker-compose ps

# Terraform shortcuts
tf-init: ## Initialize Terraform
	@rm -f .terraform.lock.hcl
	@terraform init

tf-plan: ## Run terraform plan
	@terraform plan

tf-apply: ## Apply Terraform configuration
	@terraform apply

tf-destroy: ## Destroy Terraform resources
	@terraform destroy

tf-output: ## Show Terraform outputs
	@terraform output

# Sensu shortcuts
sensuctl-config: ## Configure sensuctl
	@docker-compose exec -T backend1 sensuctl configure -n \
		--url http://127.0.0.1:8080 \
		--username admin \
		--password 'P@ssw0rd!' \
		--namespace development

sensu-secrets: sensuctl-config ## List secrets in development namespace
	@docker-compose exec backend1 sensuctl secret list --namespace development

sensu-checks: sensuctl-config ## List checks in development namespace
	@docker-compose exec backend1 sensuctl check list --namespace development

sensu-handlers: sensuctl-config ## List handlers in development namespace
	@docker-compose exec backend1 sensuctl handler list --namespace development

sensu-entities: sensuctl-config ## List entities in development namespace
	@docker-compose exec backend1 sensuctl entity list --namespace development

# API shortcuts
api-health: ## Check Sensu API health
	@curl -s http://127.0.0.1:8080/health | jq .

api-secrets: ## List secrets via API
	@curl -s -u admin:P@ssw0rd! http://127.0.0.1:8080/api/enterprise/secrets/v1/namespaces/development/secrets | jq .

api-checks: ## List checks via API
	@curl -s -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/checks | jq .

api-check-secrets: ## Show secrets in api-health-check
	@curl -s -u admin:P@ssw0rd! http://127.0.0.1:8080/api/core/v2/namespaces/development/checks/api-health-check | jq -r '.secrets'
