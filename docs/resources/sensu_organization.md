# sensu_organization

Manages a Sensu Organization.

For full documentation on Sensu Organizations, see [here](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#organization).

## Basic Example

```hcl
resource "sensu_organization" "organization_1" {
  name = "my_organization"
  description = "some organization"
}
```

_Note_: Creating an organization implicitly creates an environment called `default`.

## Argument Reference

* `name` - *Required* - The name of the Sensu organization.

* `description` - *Optional* - See the [Sensu rbac reference](https://docs.sensu.io/sensu-core/2.0/reference/rbac/#organization).

## Attribute Reference

The resource has no computed fields.
