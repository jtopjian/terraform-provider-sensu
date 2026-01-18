# sensu_secret

Manages a Sensu Secret.

For full documentation on Sensu Secrets, see [here](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets/).

## Basic Example

```hcl
resource "sensu_secret" "api_token" {
  name      = "sensu-ansible-token"
  secret_id = "ANSIBLE_TOKEN"
  secrets_provider = "env"
}
```

## Example with Custom Namespace

```hcl
resource "sensu_secret" "database_password" {
  name      = "db-password"
  namespace = "production"
  secret_id = "DATABASE_PASSWORD"
  secrets_provider = "env"
}
```

## Example with Vault Provider

```hcl
resource "sensu_secret" "vault_secret" {
  name      = "vault-api-key"
  secret_id = "secret/api#key"
  secrets_provider = "vault"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu secret.

* `secret_id` - *Required* - The secret ID. For the Env provider, this is the environment
  variable name (e.g., `ANSIBLE_TOKEN`). For other providers like Vault, this
  follows the provider-specific format (e.g., `secret/path#key`).

* `secrets_provider` - *Required* - The name of the secrets provider to use. This must
  match the name of a configured secrets provider in your Sensu backend. Common
  values include `env` for the built-in environment variable provider, or `vault`
  for HashiCorp Vault integration. Changing this value will force recreation of
  the resource.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

The resource has no additional computed fields beyond the arguments.

## Import

Secrets can be imported using the namespace and name:

```
$ terraform import sensu_secret.api_token default/sensu-ansible-token
```

## Sensu Secrets Overview

Sensu secrets allow you to securely reference sensitive information (like API keys,
passwords, and tokens) in your Sensu resources without hardcoding them in your
configuration files.

### Env Provider

The Env provider is automatically created when you start a Sensu backend and allows
you to reference environment variables set on the backend. To use it:

1. Set the environment variable on your Sensu backend:
   ```bash
   export ANSIBLE_TOKEN="your-token-here"
   ```

2. Create a secret resource that references it:
   ```hcl
   resource "sensu_secret" "ansible_token" {
     name      = "sensu-ansible-token"
     secret_id = "ANSIBLE_TOKEN"
     secrets_provider = "env"
   }
   ```

3. Use the secret in other resources:
   ```hcl
   resource "sensu_handler" "ansible_handler" {
     name    = "ansible"
     type    = "pipe"
     command = "sensu-ansible-handler"

     secrets = {
       "ANSIBLE_TOKEN" = "sensu-ansible-token"
     }
   }
   ```

### Security Considerations

* The Env provider exposes secrets from backend environment variables. Ensure
  your backend environment is properly secured.
* Secret IDs and providers are stored in Terraform state. The actual secret
  values are not stored in state, but the references are.
* When using the Env provider, ensure environment variables are set before
  the Sensu backend starts, or restart the backend after setting new variables.
* For production use, consider using a secrets management provider like Vault
  instead of the Env provider for enhanced security.

## See Also

* [Sensu Secrets Documentation](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets/)
* [Sensu Secrets Providers](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets-providers/)
* [Sensu Secrets Management](https://docs.sensu.io/sensu-go/latest/operations/manage-secrets/secrets-management/)
