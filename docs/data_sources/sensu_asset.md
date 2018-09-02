# sensu_asset

Get information about a Sensu Asset.

For full documentation on Sensu Hooks, see [here](https://docs.sensu.io/sensu-core/2.0/reference/assets).

## Basic Example

```hcl
resource "sensu_asset" "asset_1" {
  name = "my_asset"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu asset.

## Attribute Reference

* `sha512` - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets/).

* `url` - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets/).

* `filter` - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets/).

* `metadata` - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets/).
