# sensu_pipeline

Get information about a Sensu Pipeline.

For full documentation on Sensu Pipelines, see [here](https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-process/pipelines/).

## Basic Example

```
data "sensu_pipeline" "pipeline_1" {
  name = "my_pipeline"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu pipeline.

* `namespace` - *Optional* - The namespace to retrieve the pipeline from. This
  can also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `workflow` - The workflows configured for the Sensu pipeline.

### workflow

The `workflow` attribute contains:

* `name` - The name of the workflow.

* `filters` - The Sensu event filters applied to the workflow.

* `mutator` - The Sensu mutator applied to the workflow, if configured.

* `handler` - The Sensu handler used by the workflow.
