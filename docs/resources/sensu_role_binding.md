# sensu_role_binding

Manages a Sensu Role Binding.

For full documentation on Sensu Roles bindings, see [here](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings).

## Basic Example

```hcl
resource "sensu_role_binding" "role_binding_1" {
  name = "my_role_binding"
  binding_type = "role"
  role = "role_1"
  users = ["bill", "ted"]
  groups = ["devops"]
}
```

## Argument Reference

* `name` - *Required* - The rolename of the Sensu role.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `binding_type` - *Required* - The type of binding. Valid values are either
  "role" or "cluster_role". See the
  [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `role` - *Required* - The name of the role. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `users` - *Optional* - The users to bind. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `groups` - *Optional* - The groups to bind. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.
