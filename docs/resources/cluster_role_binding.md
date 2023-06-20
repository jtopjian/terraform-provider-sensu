# sensu_cluster_role_binding

Manages a Sensu Cluster Role Binding.

For full documentation on Sensu Cluster Roles bindings, see [here](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#cluster-role-bindings).

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

  [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification) for more information.

* `cluster_role` - *Required* - The name of the cluster role. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification) for more information.

* `users` - *Optional* - The users to bind. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification) for more information.

* `groups` - *Optional* - The groups to bind. See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification) for more information.
