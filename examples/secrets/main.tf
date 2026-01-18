terraform {
  required_providers {
    sensu = {
      source = "registry.terraform.io/jtopjian/sensu"
    }
  }
}

provider "sensu" {
  api_url   = "http://127.0.0.1:8080"
  username  = "admin"
  password  = "P@ssw0rd!"
  namespace = "default"
}

# =============================================================================
# Namespace
# =============================================================================
resource "sensu_namespace" "dev" {
  name = "development"
}

# =============================================================================
# Entity (agent)
# =============================================================================
resource "sensu_entity" "test_agent" {
  name      = "test-agent"
  namespace = sensu_namespace.dev.name
  class     = "agent"

  subscriptions = [
    "linux",
    "web",
  ]

  labels = {
    environment = "development"
    team        = "platform"
  }
}

# =============================================================================
# Secrets (using env provider)
# =============================================================================

# Secret for Slack webhook URL
resource "sensu_secret" "slack_webhook" {
  name      = "slack-webhook-url"
  secret_id = "SLACK_WEBHOOK_URL"
  secrets_provider = "env"
  namespace = sensu_namespace.dev.name
}

# Secret for PagerDuty API key
resource "sensu_secret" "pagerduty_key" {
  name      = "pagerduty-api-key"
  secret_id = "PAGERDUTY_API_KEY"
  secrets_provider = "env"
  namespace = sensu_namespace.dev.name
}

# Secret for external monitoring API
resource "sensu_secret" "monitoring_api_key" {
  name      = "monitoring-api-key"
  secret_id = "MONITORING_API_KEY"
  secrets_provider = "env"
  namespace = sensu_namespace.dev.name
}

# =============================================================================
# Check
# =============================================================================
resource "sensu_check" "http_check" {
  name      = "http-health-check"
  namespace = sensu_namespace.dev.name
  command   = "check-http.rb -u https://example.com/health"
  interval  = 60
  timeout   = 30
  publish   = true

  subscriptions = [
    "web",
  ]

  labels = {
    severity = "critical"
  }

  annotations = {
    documentation = "https://wiki.example.com/http-check"
  }
}

# Check that uses secrets (e.g., for authenticated API check)
resource "sensu_check" "api_check" {
  name      = "api-health-check"
  namespace = sensu_namespace.dev.name
  command   = "check-api.rb --api-key $MONITORING_API_KEY"
  interval  = 120
  timeout   = 30
  publish   = true

  subscriptions = [
    "linux",
  ]

  # Reference the secret - key is env var name, value is secret name
  # The key (MONITORING_API_KEY) becomes an environment variable in the check execution
  # The value (sensu_secret.monitoring_api_key.name) is the name of the secret resource
  secrets = {
    MONITORING_API_KEY = sensu_secret.monitoring_api_key.name
  }

  handlers = [
    sensu_handler.slack.name,
  ]
}

# =============================================================================
# Handler (using secrets)
# =============================================================================
resource "sensu_handler" "slack" {
  name      = "slack-alerts"
  namespace = sensu_namespace.dev.name
  type      = "pipe"
  command   = "sensu-slack-handler --channel '#alerts'"
  timeout   = 30

  # Reference the Slack webhook secret
  secrets = {
    SLACK_WEBHOOK_URL = sensu_secret.slack_webhook.name
  }
}

# =============================================================================
# Data Sources (read back what we created)
# =============================================================================
data "sensu_secret" "slack_webhook_data" {
  name      = sensu_secret.slack_webhook.name
  namespace = sensu_namespace.dev.name
}

# =============================================================================
# Outputs
# =============================================================================
output "namespace" {
  value = sensu_namespace.dev.name
}

output "entity_name" {
  value = sensu_entity.test_agent.name
}

output "secrets" {
  value = [
    sensu_secret.slack_webhook.name,
    sensu_secret.pagerduty_key.name,
    sensu_secret.monitoring_api_key.name,
  ]
}

output "checks" {
  value = [
    sensu_check.http_check.name,
    sensu_check.api_check.name,
  ]
}

output "handler" {
  value = sensu_handler.slack.name
}
