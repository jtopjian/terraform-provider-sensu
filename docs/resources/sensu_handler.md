# sensu_handler

Manages a Sensu Handler.

For full documentation on Sensu Handlers, see [here](https://docs.sensu.io/sensu-core/2.0/reference/handlers).

## Basic Example

```hcl
resource "sensu_handler" "handler_1" {
  name = "my_handler"
  type = "pipe"
  command = "/usr/local/bin/foo"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu handler.

* `type` - *Required* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `command` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `env_vars` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `filters` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `handlers` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `mutator` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `socket` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

* `timeout` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-core/2.0/reference/handlers/#handler-attributes).

### socket

The `socket` block supports:

* `host` - *Required* The host to connect to.

* `port` - *Required* The port to connect to.
## Attribute Reference

The resource has no computed fields.
