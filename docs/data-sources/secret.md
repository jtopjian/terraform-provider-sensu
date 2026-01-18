# sensu_secret

Get information about a Sensu Secret.

For full documentation on Sensu Secrets, see [here](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets/).

## Basic Example

```hcl
data "sensu_secret" "api_token" {
  name = "sensu-ansible-token"
}
```

## Example with Custom Namespace

```hcl
data "sensu_secret" "database_password" {
  name      = "db-password"
  namespace = "production"
}
```

## Example Usage in a Handler

```hcl
data "sensu_secret" "slack_webhook" {
  name = "slack-webhook-url"
}

resource "sensu_handler" "slack" {
  name    = "slack"
  type    = "pipe"
  command = "sensu-slack-handler --channel '#alerts'"

  secrets = {
    "SLACK_WEBHOOK_URL" = data.sensu_secret.slack_webhook.name
  }
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu secret to retrieve.

* `namespace` - *Optional* - The namespace the secret belongs to. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `secret_id` - The secret ID. For the Env provider, this is the environment variable
  name (e.g., `ANSIBLE_TOKEN`). For other providers, this follows the
  provider-specific format.

* `secrets_provider` - The name of the secrets provider used by this secret (e.g., `env`,
  `vault`).

## See Also

* [sensu_secret Resource](../resources/secret.md)
* [Sensu Secrets Documentation](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets/)
