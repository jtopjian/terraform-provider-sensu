# sensu_hook

Get information about a Sensu Hook.

For full documentation on Sensu Hooks, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/hooks/).

## Basic Example

```hcl
data "sensu_hook" "hook_1" {
  name = "my_hook"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu hook.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `command` - See the [Sensu hook reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/hooks/#hook-specification).

* `timeout` - See the [Sensu hook reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/hooks/#hook-specification).

* `stdin` - See the [Sensu hook reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/hooks/#hook-specification).
