# sensu_check

Get information about a Sensu Check.

For full documentation on Sensu Check's, see [here](https://docs.sensu.io/sensu-go/latest/reference/checks).

## Basic Example

```hcl
data "sensu_check" "check_1" {
  name = "foo"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu check.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `annotations` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#metadata-attributes).

* `command` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `check_hook` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).
  Also see the `check_hook` section below for details on this block.

* `cron` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `high_flap_threshold` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `handlers` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `interval` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `label` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#metadata-attributes).

* `low_flap_threshold` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `proxy_entity_name` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `proxy_requests` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `publish` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `round_robin` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `runtime_assets` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `stdin` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `subdue` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).
  Also see the `subdue` section below for details on this block.

* `subscriptions` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `timeout` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `ttl` - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

### check_hook

The `check_hook` block supports:

* `hook` - *Required* - The name of the `sensu_hook` to run.

* `trigger` - *Required* - Known as `type` in the Sensu documention, but, IMO,
  `trigger` makes more sense. Determines when the trigger will run. Valid values
  are: `1`-`255`, `ok`, `warning`, `critical`, `unknown`, and `non-zero`.

### proxy_requests

The `proxy_requests` block supports:

* `entity_attributes` - Attributes to match entities.

* `splay` - Enable splaying of coverage for checks.

* `splay_coverage` - The percentage of the check interval for Sensu to execute
  checks for entities matching the entity attributes.

### subdue

The `subdue` block supports:

* `day` - The day the check is subdued.

* `begin` - The start time of the subdue.

* `end` - The end time of the subdue.
