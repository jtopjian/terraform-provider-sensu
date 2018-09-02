# sensu_hook

Get information about a Sensu Hook.

For full documentation on Sensu Hooks, see [here](https://docs.sensu.io/sensu-core/2.0/reference/hooks).

## Basic Example

```hcl
resource "sensu_hook" "hook_1" {
  name = "my_hook"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu hook.

## Attribute Reference

* `command` - See the [Sensu hook reference](https://docs.sensu.io/sensu-core/2.0/reference/hooks/#hook-attributes).

* `timeout` - See the [Sensu hook reference](https://docs.sensu.io/sensu-core/2.0/reference/hooks/#hook-attributes).

* `stdin` - See the [Sensu hook reference](https://docs.sensu.io/sensu-core/2.0/reference/hooks/#hook-attributes).
