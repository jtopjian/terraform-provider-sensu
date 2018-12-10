# sensu_role

Manages a Sensu Role.

For full documentation on Sensu Roles, see [here](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role).

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

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `rule` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role).

### rule

The `rule` block supports:

* `verbs` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#rule).

* `resources` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#rule).

* `resource_names` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#rule).

## Attribute Reference

The resource has no computed fields.
