# sensu_handler

Get information about a Sensu Handler.

For full documentation on Sensu Handlers, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/).

## Basic Example

```hcl
data "sensu_handler" "handler_1" {
  name = "my_handler"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu handler.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `type` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `command` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `env_vars` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `filters` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `handlers` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `runtime_assets` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `mutator` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `socket` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

* `timeout` - See the [Sensu handler reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/handlers/#handler-specification).

### socket

The `socket` block supports:

* `host` - The host to connect to.

* `port` - The port to connect to.