# sensu_filter

Get information about a Sensu Filter

For full documentation on Sensu Filters, see [here](https://docs.sensu.io/sensu-go/latest/reference/filters).

## Basic Example

```hcl
data "sensu_filter" "filter_1" {
  name = "filter_1"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu filter.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `action` - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/reference/filters/#filter-attributes).
  Valid values are `allow` and `deny`.

* `expressions` - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/reference/filters/#filter-attributes).

* `runtime_assets` - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/reference/filters/#filter-attributes).

* `when` - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/reference/filters/#filter-attributes).

### when

The `when` block supports:

* `day` - The day for when the filter is valid.

* `begin` - The start time for when the filter is valid.

* `end` - The end time for when the filter is valid.
