# sensu_namespace

Get information about a Sensu Namespace.

For full documentation on Sensu Namespaces, see [here](https://docs.sensu.io/sensu-go/5.0/reference/rbac/#namespaces)

## Basic Example

```hcl
data "sensu_namespace" "namespace_1" {
  name = "my_namespace"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu namespace.
