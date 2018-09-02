# sensu_check

Get information about a Sensu Check.

For full documentation on Sensu Check's, see [here](https://docs.sensu.io/sensu-core/2.0/reference/checks).

## Basic Example

```hcl
data "sensu_check" "check_1" {
  name = "foo"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu check.

## Attribute Reference

* `command` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `cron` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `high_flap_threshold` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `handlers` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `interval` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `low_flap_threshold` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `proxy_entity_id` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `publish` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `round_robin` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `runtime_assets` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `stdin` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `subdue` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).
  Also see the `subdue` section below for details on this block.

* `subscriptions` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `timeout` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

* `ttl` - See the [Sensu check reference](https://docs.sensu.io/sensu-core/2.0/reference/checks/#check-attributes).

### subdue

The `subdue` block supports:

* `day` - The day the check is subdued.

* `begin` - The start time of the subdue.

* `end` - The end time of the subdue.
