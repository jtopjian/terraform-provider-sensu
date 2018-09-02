# sensu_environment

Get information about a Sensu Environment.

For full documentation on Sensu Environments, see [here](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#environment).

## Basic Example

```hcl
resource "sensu_environment" "environment_1" {
  name = "my_environment"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu environment.

## Attribute Reference

* `description` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#environment).

* `organization` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#environment).
