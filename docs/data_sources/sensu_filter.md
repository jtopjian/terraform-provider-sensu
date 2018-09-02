# sensu_filter

Get information about a Sensu Filter

For full documentation on Sensu Filters, see [here](https://docs.sensu.io/sensu-core/2.0/reference/filters).

## Basic Example

```hcl
data "sensu_filter" "filter_1" {
  name = "filter_1"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu filter.

## Attribute Reference

* `action` - See the [Sensu filter reference](https://docs.sensu.io/sensu-core/2.0/reference/filters/#filter-attributes).
  Valid values are `allow` and `deny`.

* `statements` - See the [Sensu filter reference](https://docs.sensu.io/sensu-core/2.0/reference/filters/#filter-attributes).

* `when` - See the [Sensu filter reference](https://docs.sensu.io/sensu-core/2.0/reference/filters/#filter-attributes).

### when

The `when` block supports:

* `day` - The day for when the filter is valid.

* `begin` - The start time for when the filter is valid.

* `end` - The end time for when the filter is valid.
