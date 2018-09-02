# sensu_environment

Manages a Sensu Environment.

For full documentation on Sensu Environments, see [here](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#environment).

## Basic Example

```hcl
resource "sensu_environment" "environment_1" {
  name = "my_environment"
  description = "some environment"
  organization = "default"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu environment.

* `description` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#environment).

* `organization` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#environment).
  Defaults to `default`.

## Attribute Reference

The resource has no computed fields.
