# sensu_cluster_role

Get information about a Sensu Cluser Role.

For full documentation on Sensu Cluster Roles, see [here](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#cluster-roles).

## Basic Example

```hcl
data "sensu_cluster_role" "cluster_role_1" {
  name = "my_role"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu cluster role.

## Attribute Reference

* `rule` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).

### rule

The `rule` block supports:

* `verbs` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).

* `resources` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).

* `resource_names` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).
