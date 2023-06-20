# sensu_mutator

Get information about a Sensu Mutator.

For full documentation on Sensu Mutators, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/).

## Basic Example

```hcl
data "sensu_mutator" "mutator_1" {
  name = "my_mutator"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu mutator.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `command` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/#mutator-specification).

* `timeout` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/#mutator-specification).

* `env_vars` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/#mutator-specification).

* `secrets` - *Optional* - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-transform/mutators/#mutator-specification).