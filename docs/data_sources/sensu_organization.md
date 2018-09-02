# sensu_organization

Get information about a Sensu Organization.

For full documentation on Sensu Organizations, see [here](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#organization).

## Basic Example

```hcl
resource "sensu_organization" "organization_1" {
  name = "my_organization"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu organization.

## Attribute Reference

* `description` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#organization).
