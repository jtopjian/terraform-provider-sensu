# sensu_entity

Manage a Sensu Entity.

For full documentation on Sensu Entities, see [here](https://docs.sensu.io/sensu-go/latest/reference/entities).

NOTE: For now, the Sensu API only supports managing proxy entities.
Agent entities have their settings overwritten when the agent checks in to the backend.

## Basic Example

```hcl
resource "sensu_entity" "entity_1" {
  name = "entity_1"
  class = "proxy"

  labels = {
    foo = "bar"
    password = "secret"
  }
}
```

## Example with Proxies and Checks

```hcl
resource "sensu_entity" "entities" {
  count = 3
  name  = format("entity-%02d", count.index + 1)
  class = "proxy"
  labels = {
    "url" = format("http://example-%02d.com", count.index + 1)
    "proxy_type" = "website"
  }
}

resource "sensu_check" "check_1" {
  name     = "check-http"
  command  = "/bin/echo {{ .labels.url }}"
  interval = 60
  proxy_requests {
    entity_attributes = [
      "entity.entity_class == 'proxy'",
      "entity.labels.proxy_type == 'website'",
    ]
  }
  publish       = true
  subscriptions = ["entity:agent1"]
}
```

## Argument Reference

* `name` - *Required* - The name / ID of the Sensu entity.

* `class` - *Required* - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `annotations` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes). NOTE: if any of the annotations match a Sensu redacted field, then
the value of the annotation will always be seen as REDACTED. Redacting does
not secure sensitive data. It prevents sensitive data from ever being seen.

* `deregister` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `deregistration` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `labels` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes). NOTE: if any of the labels match a Sensu redacted field, then
the value of the label will always be seen as REDACTED. Redacting does not
secure sensitive data. It prevents sensitive data from ever being seen.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `redact` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `subscriptions` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

### deregistration

The `deregistration` block supports:

* `handler` - The handler used for deregistration

## Attribute Reference

* `last_seen` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `system` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

* `user` - See the [Sensu entity reference](https://docs.sensu.io/sensu-go/latest/reference/entities/#entity-attributes).

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
