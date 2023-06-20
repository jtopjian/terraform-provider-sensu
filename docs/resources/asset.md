# sensu_asset

Manages a Sensu Asset.

For full documentation on Sensu Assets, see [here](https://docs.sensu.io/sensu-go/latest/plugins/assets/).

_Note_: The Sensu API currently cannot delete Assets.

## Basic Example

```hcl
resource "sensu_asset" "asset_1" {
  name = "sensu-plugins-cpu-checks"

  build {
    sha512 = "abcd1234..."
    url = "http://example.com/asset/example.tar.gz"
    filters = [
      "System.OS=='linux'",
      "System.Arch=='amd64'",
    ]
  }
}
```

## Deprecated Example

```hcl
resource "sensu_asset" "asset_1" {
  name = "asset_1"
  sha512 = "abcd1234..."
  url = "http://example.com/asset/example.tar.gz"

  filters = [
    "System.OS=='linux'",
    "System.Arch=='amd64'",
  ]
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu asset.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `build` - *Required, if `url`, `sha512` and `filters` are not provided* - Defines a build for an asset. Define more than one `build` block for
  multiple-build assets

* `sha512` - *Required, unless `builds` are provided* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
  This was for single-build assets which have been deprecated. It is recommended to use the `build` block
  for multiple-build assets.

* `url` - *Required, unless `builds` are provided* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
  This was for single-build assets which have been deprecated. It is recommended to use the `build` block
  for multiple-build assets.

* `filters` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
  This was for single-build assets which have been deprecated. It is recommended to use the `build` block
  for multiple-build assets.

* `headers` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

### build

The `build` block supports:

* `sha512` - *Required* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `url` - *Required* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `filters` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `headers` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).


## Attribute Reference

The resource has no computed fields.
