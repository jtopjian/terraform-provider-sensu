# sensu_filter

Manages a Sensu Filter.

For full documentation on Sensu Filters, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-filter/filters/).

## Basic Example

```hcl
resource "sensu_filter" "filter_1" {
  name = "my_filter"
  action = "allow"

  statements = [
    "event.Check.Team == 'ops'",
  ]

  when {
    day = "monday"
    begin = "09:00AM"
    end = "05:00PM"
  }
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu filter.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `action` - *Required* - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-filter/filters/#event-filter-specification).
  Valid values are `allow` and `deny`.

* `expressions` - *Optional* - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-filter/filters/#event-filter-specification).

* `runtime_assets` - *Optional* - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-filter/filters/#event-filter-specification).

* `when` - *Optional* - See the [Sensu filter reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-filter/filters/#event-filter-specification).

### when

The `when` block supports:

* `day` - *Required* - The day for when the filter is valid.

* `begin` - *Required* - The start time for when the filter is valid.

* `end` - *Required* - The end time for when the filter is valid.

## Attribute Reference

The resource has no computed fields.
