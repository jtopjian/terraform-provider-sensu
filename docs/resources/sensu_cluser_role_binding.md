# sensu_cluster_role_binding

Manages a Sensu Cluster Role Binding.

For full documentation on Sensu Cluster Roles bindings, see [here](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings).

## Basic Example

```hcl
resource "sensu_cluster_role_binding" "cluster_role_binding_1" {
  name = "my_role_binding"
  cluster_role = "cluster_role_1"
  users = ["bill", "ted"]
  groups = ["devops"]
}
```

## Argument Reference

* `name` - *Required* - The rolename of the Sensu role.

  [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `role` - *Required* - The name of the role. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `users` - *Optional* - The users to bind. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `groups` - *Optional* - The groups to bind. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.
