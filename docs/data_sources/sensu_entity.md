# sensu_entity

Get information about a Sensu Entity.

For full documentation on Sensu Entities, see [here](https://docs.sensu.io/sensu-go/latest/reference/entities).

## Basic Example

```hcl
data "sensu_entity" "entity_1" {
  name = "my_entity"
}
```

## Argument Reference

* `name` - *Required* - The name / ID of the Sensu entity.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `annotations` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `class` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `deregister` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `deregistration` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `labels` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `last_seen` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `subscriptions` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `system` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `user` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

### deregistration

The `deregistration` block supports:

* `handler` - The handler used for deregistration

### system

The `system` block supports:

* `hostname` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#system-attributes)

* `os` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#system-attributes)

* `platform` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#system-attributes)

* `platform_family` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#system-attributes)

* `platform_version` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#system-attributes)

* `arch` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#system-attributes)

* `network_interfaces` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#system-attributes)

### network_interfaces

The `network_interfaces` block supports:

* `name` - See the [Sensu entity network reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#networkinterface-attributes).

* `mac` - See the [Sensu entity network reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#networkinterface-attributes).

* `addresses` - See the [Sensu entity network reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#networkinterface-attributes).
