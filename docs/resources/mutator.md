# sensu_mutator

Manages a Sensu Mutator.

For full documentation on Sensu Mutators, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/).

## Basic Example

```hcl
resource "sensu_mutator" "mutator_1" {
  name = "my_mutator"
  command = "/usr/local/bin/foo"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu mutator.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `command` - *Required* - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/#mutator-specification).

* `timeout` - *Optional* - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/#mutator-specification).
  Defaults to 60.

* `env_vars` - *Optional* - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/#mutator-specification).

## Attribute Reference

The resource has no computed fields.