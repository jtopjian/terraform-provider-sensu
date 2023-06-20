# sensu_cluster_role

Manages a Sensu Cluster Role.

For full documentation on Sensu Cluster Roles, see [here](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#cluster-roles).

## Basic Example

```hcl
resource "sensu_cluster_role" "cluster_role_1" {
  name = "my_role"
  rule {
    verbs = ["get", "list"]
    resource = ["checks"]
  }
}
```

## Argument Reference

* `name` - *Required* - The rolename of the Sensu role.

* `rule` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).

### rule

The `rule` block supports:

* `verbs` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).

* `resources` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).

* `resource_names` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-and-cluster-role-specification).

## Attribute Reference

The resource has no computed fields.
