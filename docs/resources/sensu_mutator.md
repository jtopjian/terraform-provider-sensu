# sensu_mutator

Manages a Sensu Mutator.

For full documentation on Sensu Mutators, see [here](https://docs.sensu.io/sensu-core/2.0/reference/mutators).

## Basic Example

```hcl
resource "sensu_mutator" "mutator_1" {
  name = "my_mutator"
  command = "/usr/local/bin/foo"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu mutator.

* `command` - *Required* - See the [Sensu mutator reference](https://docs.sensu.io/sensu-core/2.0/reference/mutators/#attributes).

* `timeout` - *Optional* - See the [Sensu mutator reference](https://docs.sensu.io/sensu-core/2.0/reference/mutators/#attributes).
  Defaults to 60.

* `env_vars` - *Optional* - See the [Sensu mutator reference](https://docs.sensu.io/sensu-core/2.0/reference/mutators/#attributes).

## Attribute Reference

The resource has no computed fields.
