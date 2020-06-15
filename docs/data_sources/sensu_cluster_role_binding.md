# sensu_cluster_role_binding

Get information about a Sensu Cluser Role.

For full documentation on Sensu Cluster Roles bindings, see [here](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings).

## Basic Example

```hcl
data "sensu_cluster_role_binding" "cluster_role_binding_1" {
  name         = "cluster_role_binding_1"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu role.

## Attribute Reference

* `cluster_role` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `users` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.

* `groups` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#role-bindings-and-cluster-role-bindings) for more information.
