# sensu_hook

Manages a Sensu Hook.

For full documentation on Sensu Hooks, see [here](https://docs.sensu.io/sensu-core/2.0/reference/hooks).

## Basic Example

```hcl
resource "sensu_hook" "hook_1" {
  name = "my_hook"
  command = "/usr/local/bin/foo"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu hook.

* `command` - *Required* - See the [Sensu hook reference](https://docs.sensu.io/sensu-core/2.0/reference/hooks/#hook-attributes).

* `timeout` - *Optional* - See the [Sensu hook reference](https://docs.sensu.io/sensu-core/2.0/reference/hooks/#hook-attributes).
  Defaults to 60.

* `stdin` - *Optional* - See the [Sensu hook reference](https://docs.sensu.io/sensu-core/2.0/reference/hooks/#hook-attributes).
  Defaults to false.

## Attribute Reference

The resource has no computed fields.
