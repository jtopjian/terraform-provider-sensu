# sensu_role

Manages a Sensu Role.

For full documentation on Sensu Roles, see [here](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#role).

## Basic Example

```hcl
resource "sensu_role" "role_1" {
  name = "my_role"
  password = "abcd1234"
  roles = ["admin"]
}
```

## Argument Reference

* `name` - *Required* - The rolename of the Sensu role.

* `rule` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#role).

### rule

The `rule` block supports:

* `type` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).

* `environment` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).

* `organization` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).

* `permissions` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#rule).

## Attribute Reference

The resource has no computed fields.
