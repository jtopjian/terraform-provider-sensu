# sensu_handler

Manages a Sensu Handler.

For full documentation on Sensu Handlers, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/).

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

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `type` - *Required* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `command` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `env_vars` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `filters` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `handlers` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `runtime_assets` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `mutator` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `secrets` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `socket` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `timeout` - *Optional* - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

### socket

The `socket` block supports:

* `host` - *Required* The host to connect to.

* `port` - *Required* The port to connect to.

## Attribute Reference

The resource has no computed fields.