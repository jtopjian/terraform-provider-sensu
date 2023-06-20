# sensu_cluster_role_binding

Get information about a Sensu Cluser Role.

For full documentation on Sensu Cluster Roles bindings, see [here](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#cluster-role-bindings).

## Basic Example

```hcl
data "sensu_cluster_role_binding" "cluster_role_binding_1" {
  name         = "cluster_role_binding_1"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu role.

## Attribute Reference

* `cluster_role` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification) for more information.

* `users` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification) for more information.

* `groups` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification) for more information.
