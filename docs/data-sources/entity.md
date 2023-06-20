# sensu_entity

Get information about a Sensu Entity.

For full documentation on Sensu Entities, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/).

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

* `annotations` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `class` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `deregister` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `deregistration` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `labels` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `last_seen` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `subscriptions` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `system` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

* `user` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#entities-specification).

### deregistration

The `deregistration` block supports:

* `handler` - The handler used for deregistration

### system

The `system` block supports:

* `hostname` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes)

* `os` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes)

* `platform` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes)

* `platform_family` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes)

* `platform_version` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes)

* `arch` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes)

* `network_interfaces` - See the [Sensu entity system reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes)

### network_interfaces

The `network_interfaces` block supports:

* `name` - See the [Sensu entity network reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes).

* `mac` - See the [Sensu entity network reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes).

* `addresses` - See the [Sensu entity network reference](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#system-attributes).
