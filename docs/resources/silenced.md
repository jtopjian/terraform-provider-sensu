# sensu_silenced

Get information about a Sensu Silencing.

For full documentation on Sensu Silencing, see [here](https://docs.sensu.io/sensu-go/latest/reference/silencing).

## Basic Example

```hcl
resource "sensu_silenced" "silence_1" {
  check = "foo"
  subscription = "entity:bar"
  begin = "Jun 02 2020 3:04PM MST"
}
```

## Argument Reference

* `check` - *Required* - The name of the check the entry should match

* `subscription` - *Required* - The name of the subscription the entry should match.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `begin` - *Optional* - Time at which silence entry goes into effect
  in human readable time (Format: Jan 02 2006 3:04PM MST)". If not set,
  this defaults to `now`.

* `expire` - *Optional* - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/reference/silencing/#spec-attributes). 

* `expire_on_resolve` - *Optional* - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/reference/silencing/#spec-attributes).

* `reason` - *Optional* - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/reference/silencing/#spec-attributes). 

* `labels` - *Optional* - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/reference/silencing/#metadata-attributes). 

* `annotations` - *Optional* - See the [Sensu silence reference](https://docs.sensu.io/sensu-go/latest/reference/silencing/#metadata-attributes).

## Attribute Reference

The resource has no computed fields.