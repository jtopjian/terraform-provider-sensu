# sensu_check

Manages a Sensu Check.

For full documentation on Sensu Checks, see [here](https://docs.sensu.io/sensu-go/latest/reference/checks).

## Basic Example

```hcl
resource "sensu_check" "check_1" {
  name     = "my_check"
  command  = "/usr/local/bin/foo"
  interval = 600

  subscriptions = [
    "foo",
    "bar",
  ]
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu check.

* `command` - *Required* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `subscriptions` - *Required* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `annotations` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#metadata-attributes).

* `check_hook` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).
  Also see the `check_hook` section below for details on this block.

* `cron` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `handlers` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `interval` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `high_flap_threshold` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `label` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#metadata-attributes).

* `low_flap_threshold` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `proxy_entity_name` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `proxy_requests` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `publish` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `round_robin` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `runtime_assets` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `stdin` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `subdue` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).
  Also see the `subdue` section below for details on this block.

* `timeout` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

* `ttl` - *Optional* - See the [Sensu check reference](https://docs.sensu.io/sensu-go/latest/reference/checks/#check-attributes).

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

* `day` - *Required* - The day to subdue the check.

* `begin` - *Required* - The start time of the subdue. Must be in a format such as "09:00AM".

* `end` - *Required* - The end time of the subdue. Must be in a format such as "09:00AM".

## Attribute Reference

This resource has no computed fields.
