# sensu_asset

Manages a Sensu Asset.

For full documentation on Sensu Assets, see [here](https://docs.sensu.io/sensu-core/2.0/reference/assets).

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

  metadata {
    Content-Type = "application/zip"
    X-Intended-Distribution = "trusty-14"
  }
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu asset.

* `sha512` - *Required* - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets).

* `url` - *Required* - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets).

* `filters` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets).

* `metadata` - *Optional* - See the [Sensu asset reference](https://docs.sensu.io/sensu-core/2.0/reference/assets).

## Attribute Reference

The resource has no computed fields.
