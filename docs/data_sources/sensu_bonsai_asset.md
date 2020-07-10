# sensu_bonsai_asset

Get information about a Sensu Asset hosted at Bonsai.

For full documentation on Sensu Assets, see [here](https://docs.sensu.io/sensu-go/latest/reference/assets/).

## Basic Example

```hcl
data "sensu_bonsai_asset" "bonsai_asset_1" {
  name = "sensu-plugins/sensu-plugins-cpu-checks"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu Bonsai asset.

* `version` - *Optional* - The version of the Sensu Bonsai asset.

## Attribute Reference

* `labels` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/reference/assets/)

* `annotations` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/reference/assets/).

* `builds` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/reference/assets/).

### Builds

The `build` block contains the following, for each asset build:

* `sha512` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/reference/assets/).

* `url` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/reference/assets/).

* `filters` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/reference/assets/).

* `headers` - See the [Sensu asset reference](https://docs.sensu.io/sensu-go/latest/reference/assets/).
