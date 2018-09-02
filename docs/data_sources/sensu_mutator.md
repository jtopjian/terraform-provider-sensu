# sensu_mutator

Get information about a Sensu Mutator.

For full documentation on Sensu Mutators, see [here](https://docs.sensu.io/sensu-core/2.0/reference/mutators).

## Basic Example

```hcl
resource "sensu_mutator" "mutator_1" {
  name = "my_mutator"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu mutator.

## Attribute Reference

* `command` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-core/2.0/reference/mutators/#attributes).

* `timeout` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-core/2.0/reference/mutators/#attributes).

* `env_vars` - See the [Sensu mutator reference](https://docs.sensu.io/sensu-core/2.0/reference/mutators/#attributes).
