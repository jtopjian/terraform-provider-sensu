# sensu_pipeline

Manages a Sensu Pipeline.

For full documentation on Sensu Pipelines, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/pipelines/).

## Basic Example

```
resource "sensu_pipeline" "pipeline_1" {
  name = "my_pipeline"

  workflow {
    name    = "workflow_1"
    handler = "my_handler"
  }
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu pipeline.

* `namespace` - *Optional* - The namespace to manage the pipeline in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `workflow` - *Required* - One or more workflows that define the filters,
  mutator, and handler used by the pipeline.

### workflow

The `workflow` block supports:

* `name` - *Required* - The name of the workflow.

* `filters` - *Optional* - A list of Sensu event filter names to apply to the
  workflow.

* `mutator` - *Optional* - The name of the Sensu mutator to apply to the
  workflow.

* `handler` - *Required* - The name of the Sensu handler to use for the
  workflow.

## Attribute Reference

The resource has no computed fields.
