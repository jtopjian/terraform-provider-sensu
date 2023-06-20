# sensu_bonsai_asset

Get information about a Sensu Asset hosted at Bonsai.

For full documentation on Sensu Assets, see [here](https://docs.sensu.io/sensu-go/latest/plugins/assets/).

## Basic Example

```hcl
data "sensu_bonsai_asset" "bonsai_asset_1" {
  name = "sensu-plugins/sensu-plugins-cpu-checks"
}
```

## Create Asset from Bonsai

```hcl
data "sensu_bonsai_asset" "bonsai_asset_1" {
  name = "sensu-plugins/sensu-plugins-cpu-checks"
  version = "4.1.0"
}

resource "sensu_asset" "asset_1" {
  name = data.sensu_bonsai_asset.bonsai_asset_1.annotations["io.sensu.bonsai.name"]

  dynamic "build" {
    for_each = data.sensu_bonsai_asset.bonsai_asset_1.build
    content {
      sha512 = build.value["sha512"]
      url = build.value["url"]
      filters = build.value["filters"]
      headers = build.value["headers"]
    }
  }
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu Bonsai asset.

* `version` - *Optional* - The version of the Sensu Bonsai asset.

## Attribute Reference

* `labels` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification)

* `annotations` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `builds` - Deprecated. Use `build`.

* `build` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

### Build

The `build` block contains the following, for each asset build:

* `sha512` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `url` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `filters` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).

* `headers` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/plugins/assets/#asset-specification).
