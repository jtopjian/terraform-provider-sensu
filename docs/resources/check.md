# sensu_check

Manages a Sensu Check.

For full documentation on Sensu Checks, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/checks/).

## Basic Example

```hcl
resource "sensu_check" "check_1" {
  name          = "my_check"
  command       = "/usr/local/bin/foo"
  interval      = 600
  subscriptions = ["foo", "bar"]
}
```

## Output Metrics Example

The `sensu_check` resource supports configuring output metrics, including metric format, tags, thresholds, and handlers.

```hcl
resource "sensu_check" "certificate" {
  name          = "certificate-check"
  command       = "cert-checks -c https://example.com/certificate | grep cert_days_left"
  interval      = 600
  timeout       = 30
  publish       = true
  subscriptions = ["system"]

  output_metric_format = "prometheus_text"

  output_metric_tags {
    name  = "entity"
    value = "{{ .name }}"
  }

  output_metric_thresholds {
    name        = "cert_days_left"
    null_status = 1

    thresholds {
      min    = "30.0"
      status = 1
    }

    thresholds {
      min    = "14.0"
      status = 2
    }
  }

  output_metric_handlers = ["prometheus"]

  runtime_assets = ["sensu/cert-checks"]
}
```

## Argument Reference

* `name` - (Required) The name of the check.
* `command` - (Required) The command to execute.
* `subscriptions` - (Required) The subscriptions that will cause the check to execute.
* `annotations` - (Optional) A map of annotations to add to the check.
* `check_hook` - (Optional) A set of check hooks to execute based on the check status.
* `cron` - (Optional) A cron expression specifying when the check should execute. Conflicts with `interval`.
* `env_vars` - (Optional) Environment variables to make available to the check command.
* `handlers` - (Optional) A list of handlers to execute when the check produces an event.
* `high_flap_threshold` - (Optional) The number of status changes within the flap detection window required to enter a flapping state.
* `interval` - (Optional) The interval, in seconds, between check executions. Conflicts with `cron`.
* `labels` - (Optional) A map of labels to add to the check.
* `low_flap_threshold` - (Optional) The number of status changes within the flap detection window required to leave a flapping state.
* `namespace` - (Optional) The Sensu namespace in which to create the check.
* `output_metric_format` - (Optional) The output metric format to use when parsing metrics from check output.
* `output_metric_handlers` - (Optional) A list of handlers to use for output metrics.
* `output_metric_tags` - (Optional) A list of tags to add to output metrics.
* `output_metric_thresholds` - (Optional) A list of metric definitions and thresholds used to determine check status.
* `pipelines` - (Optional) A list of pipelines to associate with the check.
* `proxy_entity_name` - (Optional) The name of the entity to proxy the check request to.
* `proxy_requests` - (Optional) Configuration for proxy requests.
* `publish` - (Optional) Whether the check should publish events.
* `round_robin` - (Optional) Whether to execute the check using round-robin scheduling.
* `runtime_assets` - (Optional) A list of runtime assets required by the check.
* `secrets` - (Optional) A list of secrets to expose to the check command.
* `stdin` - (Optional) Whether to provide standard input to the check command.
* `subdue` - (Optional) Time windows during which the check should be subdued.
* `timeout` - (Optional) The check execution timeout, in seconds.
* `ttl` - (Optional) The amount of time, in seconds, before a check is considered stale.

### check_hook

The `check_hook` block configures a hook to execute when a check produces a particular status.

```hcl
check_hook {
  hook    = "my-hook"
  trigger = "non-zero"
}
```

* `hook` - (Optional) The name of the hook to execute.
* `trigger` - (Optional) The event that triggers the hook.

### output_metric_tags

The `output_metric_tags` block defines tags that are added to metrics produced by the check.

```hcl
output_metric_tags {
  name  = "entity"
  value = "{{ .name }}"
}
```

* `name` - (Required) The name of the metric tag.
* `value` - (Required) The value of the metric tag. Sensu template expressions may be used.

Multiple `output_metric_tags` blocks may be specified.

### output_metric_thresholds

The `output_metric_thresholds` block defines a metric and the thresholds that determine the check status based on the metric's value.

```hcl
output_metric_thresholds {
  name        = "cert_days_left"
  null_status = 1

  thresholds {
    min    = "30.0"
    status = 1
  }

  thresholds {
    min    = "14.0"
    status = 2
  }
}
```

* `name` - (Required) The name of the metric to which the thresholds apply.
* `null_status` - (Optional) The status to return when the metric value is null.
* `tags` - (Optional) A list of tags to associate with the metric.
* `thresholds` - (Optional) A list of threshold definitions.

Multiple `output_metric_thresholds` blocks may be specified.

#### tags

The `tags` block defines tags associated with an output metric threshold definition.

```hcl
output_metric_thresholds {
  name = "cert_days_left"

  tags {
    name  = "hostname"
    value = "{{ .labels.hostname }}"
  }
}
```

* `name` - (Required) The name of the metric tag.
* `value` - (Required) The value of the metric tag. Sensu template expressions may be used.

Multiple `tags` blocks may be specified.

#### thresholds

The `thresholds` block defines the minimum and/or maximum value for a metric and the Sensu status to return when the threshold is met.

```hcl
thresholds {
  min    = "30.0"
  max    = "60.0"
  status = 1
}
```

* `min` - (Optional) The minimum value for the threshold.
* `max` - (Optional) The maximum value for the threshold.
* `status` - (Required) The Sensu status code to return when the threshold condition is met.

Multiple `thresholds` blocks may be specified.

### pipelines

The `pipelines` block associates one or more Sensu pipelines with the check.

```hcl
pipelines {
  api_version = "core/v2"
  type        = "Pipeline"
  name        = "my-pipeline"
}
```

* `api_version` - (Optional) The API version of the pipeline.
* `type` - (Optional) The resource type of the pipeline.
* `name` - (Optional) The name of the pipeline.

### proxy_requests

The `proxy_requests` block configures how a check is proxied to another entity.

```hcl
proxy_requests {
  entity_attributes = [
    "entity.system.os",
    "entity.system.arch",
  ]

  splay          = true
  splay_coverage = 90
}
```

* `entity_attributes` - (Optional) A list of entity attributes used to select a proxy entity.
* `splay` - (Optional) Whether to enable splay.
* `splay_coverage` - (Optional) The percentage of the interval over which requests are distributed.

### subdue

The `subdue` block defines time windows during which the check should not execute.

## Attribute Reference

This resource has no computed fields.
