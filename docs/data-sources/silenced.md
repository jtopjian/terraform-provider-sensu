# sensu_silenced

Get information about a Sensu Silencing.

For full documentation on Sensu Silencing, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/silencing/).

## Basic Example

```hcl
data "sensu_silenced" "silence_1" {
  check = "foo"
  subscription = "entity:bar"
}
```

## Argument Reference

* `check` - *Required* - The name of the check the entry should match

* `subscription` - *Required* - The name of the subscription the entry should match.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `begin` - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/silencing/#silencing-specification).
* `expire` - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/silencing/#silencing-specification). 
* `expire_on_resolve` - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/silencing/#silencing-specification). 
* `reason` - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/silencing/#silencing-specification). 
* `labels` - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/reference/silencing/#spec-attributes). 
* `annotations` - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/reference/silencing/#spec-attributes).
