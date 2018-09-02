# sensu_user

Get information about a Sensu User.

For full documentation on Sensu Users, see [here](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#user).

## Basic Example

```hcl
resource "sensu_user" "user_1" {
  name = "my_user"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu user.

## Attribute Reference

* `roles` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#user).
