# sensu_namespace

Manage a Sensu Namespace.

For full documentation on Sensu Namespaces, see [here](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#namespaces)

## Basic Example

```hcl
resource "sensu_namespace" "namespace_1" {
  name = "my_namespace"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu namespace.
