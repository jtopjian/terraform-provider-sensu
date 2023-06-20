# sensu_asset

Get information about a Sensu Asset.

For full documentation on Sensu Asset, see [here](https://docs.sensu.io/sensu-go/latest/plugins/assets/).

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

* `build` - Defines a build for an asset. Define more than one `build` block for

* `sha512` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
  This was for single-build assets which have been deprecated. It is recommended to use the `build` block
  for multiple-build assets.

* `url` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
  This was for single-build assets which have been deprecated. It is recommended to use the `build` block
  for multiple-build assets.

* `filter` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
  This was for single-build assets which have been deprecated. It is recommended to use the `build` block
  for multiple-build assets.

* `headers` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
  This was for single-build assets which have been deprecated. It is recommended to use the `build` block
  for multiple-build assets.

### build

The `build` block supports:

* `sha512` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `url` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `filters` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `headers` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

