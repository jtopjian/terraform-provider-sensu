# sensu_mutator

Get information about a Sensu Mutator.

For full documentation on Sensu Mutators, see [here](https://docs.sensu.io/sensu-go/5.0/reference/mutators).

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

* `command` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/5.0/reference/mutators/#attributes).

* `timeout` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/5.0/reference/mutators/#attributes).

* `env_vars` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-go/5.0/reference/mutators/#attributes).
