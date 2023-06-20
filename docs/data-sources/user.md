# sensu_user

Get information about a Sensu User.

For full documentation on Sensu Users, see [here](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#users).

## Basic Example

```hcl
data "sensu_user" "user_1" {
  name = "my_user"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu user.

## Attribute Reference

* `groups` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#user-specification).
