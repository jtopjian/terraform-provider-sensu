# sensu_user

Manages a Sensu User.

For full documentation on Sensu Users, see [here](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#user).

_Note_: When Terraform deletes a user, the user is only disabled.
This prevents Terraform from creating another user with the same
username.

## Basic Example

```hcl
resource "sensu_user" "user_1" {
  name = "my_user"
  password = "abcd1234"
  groups = ["admin"]
  disabled = false
}
```

_Note_: if user already exist, it will be re-enabled and updated with resource data.

## Argument Reference

* `name` - *Required* - The username of the Sensu user.

* `password` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#user).

* `groups` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#user).

* `disabled` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#user).

## Attribute Reference

The resource has no computed fields.
