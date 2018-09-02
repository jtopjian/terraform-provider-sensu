# sensu_handler

Get information about a Sensu Handler.

For full documentation on Sensu Handlers, see [here](https://docs.sensu.io/sensu-core/2.0/reference/handlers).

## Basic Example

```hcl
resource "sensu_handler" "handler_1" {
  name = "my_handler"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu handler.

## Attribute Reference

* `type` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `command` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `env_vars` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `filters` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `handlers` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `mutator` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `socket` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `timeout` - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

### socket

The `socket` block supports:

* `host` - The host to connect to.

* `port` - The port to connect to.
