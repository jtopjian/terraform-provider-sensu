# sensu_asset

Manages a Sensu Asset.

For full documentation on Sensu Assets, see [here](https://docs.sensu.io/sensu-go/5.0/reference/assets).

_Note_: The Sensu API currently cannot delete Assets.

## Basic Example

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

* `sha512` - *Required* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets).

* `url` - *Required* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets).

* `filters` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets).

* `headers` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/5.0/reference/assets).

## Attribute Reference

The resource has no computed fields.
