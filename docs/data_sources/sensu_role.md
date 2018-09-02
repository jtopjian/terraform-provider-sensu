# sensu_role

Get information about a Sensu Role.

For full documentation on Sensu Roles, see [here](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#role).

## Basic Example

```hcl
resource "sensu_role" "role_1" {
  name = "my_role"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu role.

## Attribute Reference

* `rule` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#role).

### rule

The `rule` block supports:

* `type` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).

* `environment` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).

* `organization` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).

* `permissions` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).
