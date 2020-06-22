# sensu_cluster_role

Get information about a Sensu Cluser Role.

For full documentation on Sensu Cluster Roles, see [here](https://docs.sensu.io/sensu-go/latest/reference/rbac/#roles-and-cluster-roles).

## Basic Example

```hcl
data "sensu_cluster_role" "cluster_role_1" {
  name = "my_role"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu cluster role.

## Attribute Reference

* `rule` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).

### rule

The `rule` block supports:

* `verbs` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).

* `resources` - *Required* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).

* `resource_names` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/reference/rbac/#rule-attributes).
