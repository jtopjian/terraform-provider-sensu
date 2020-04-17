# sensu_asset

Get information about a Sensu Asset.

For full documentation on Sensu Hooks, see [here](https://docs.sensu.io/sensu-go/5.0/reference/assets).

## Basic Example

```hcl
data "sensu_asset" "asset_1" {
  name = "my_asset"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu asset.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.


## Attribute Reference

* `sha512` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets/).

* `url` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets/).

* `filter` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets/).

* `headers` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets/).
