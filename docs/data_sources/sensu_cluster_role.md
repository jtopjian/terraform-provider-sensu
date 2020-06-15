# sensu_cluster_role

Get information about a Sensu Cluser Role.

For full documentation on Sensu Cluster Roles, see [here](https://docs.sensu.io/sensu-go/5.20/reference/rbac/#roles-and-cluster-roles).

## Basic Example

```hcl
data "sensu_cluster_role" "cluster_role_1" {
  name = "my_role"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu cluster role.

## Attribute Reference

* `rule` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).

### rule

The `rule` block supports:

* `verbs` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).

* `resources` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).

* `resource_names` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).
