package sensu

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/sensu-go/backend/apid/actions"
	"github.com/sensu/sensu-go/cli/client"
)

func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ForceNew: true,
			},

			"workflow": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"filters": {
							Type:     schema.TypeList,
							Optional: true,

							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"mutator": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"handler": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourcePipelineCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	pipeline := &corev2.Pipeline{
		ObjectMeta: corev2.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Workflows: expandPipelineWorkflows(d.Get("workflow")),
	}

	if err := pipeline.Validate(); err != nil {
		return err
	}

	bytes, err := json.Marshal(pipeline)
	if err != nil {
		return err
	}

	path := client.PipelinesPath(namespace)

	res, err := config.client.R().SetBody(bytes).Post(path)
	if err != nil {
		return err
	}

	if res.StatusCode() >= 400 {
		return client.UnmarshalError(res)
	}

	d.SetId(name)

	return resourcePipelineRead(d, meta)
}

func resourcePipelineRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Id()
	pipeline, err := config.client.FetchPipeline(name)

	if err != nil {
		if apiErr, ok := err.(client.APIError); ok && apiErr.Code == uint32(actions.NotFound) {
			d.SetId("")
			return nil
		}

		return err
	}

	d.Set("name", pipeline.GetName())
	d.Set("namespace", pipeline.GetNamespace())
	d.Set("workflow", flattenPipelineWorkflows(pipeline.Workflows))

	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	pipeline := &corev2.Pipeline{
		ObjectMeta: corev2.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Workflows: expandPipelineWorkflows(d.Get("workflow")),
	}

	if err := pipeline.Validate(); err != nil {
		return err
	}

	if err := config.client.UpdatePipeline(pipeline); err != nil {
		return err
	}

	return resourcePipelineRead(d, meta)
}

func resourcePipelineDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Id()
	namespace := d.Get("namespace").(string)

	if err := config.client.DeletePipeline(namespace, name); err != nil {
		if apiErr, ok := err.(client.APIError); ok && apiErr.Code == uint32(actions.NotFound) {
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId("")

	return nil
}

func expandPipelineWorkflows(v interface{}) []*corev2.PipelineWorkflow {
	var workflows []*corev2.PipelineWorkflow

	for _, raw := range v.([]interface{}) {
		data := raw.(map[string]interface{})

		workflow := &corev2.PipelineWorkflow{
			Name: data["name"].(string),
		}

		for _, rawFilter := range data["filters"].([]interface{}) {
			filter := rawFilter.(string)

			workflow.Filters = append(workflow.Filters, &corev2.ResourceReference{
				Name:       filter,
				APIVersion: "core/v2",
				Type:       "EventFilter",
			})
		}

		if mutator, ok := data["mutator"].(string); ok && mutator != "" {
			workflow.Mutator = &corev2.ResourceReference{
				Name:       mutator,
				APIVersion: "core/v2",
				Type:       "Mutator",
			}
		}

		workflow.Handler = &corev2.ResourceReference{
			Name:       data["handler"].(string),
			APIVersion: "core/v2",
			Type:       "Handler",
		}

		workflows = append(workflows, workflow)
	}

	return workflows
}

func flattenPipelineWorkflows(workflows []*corev2.PipelineWorkflow) []interface{} {
	result := make([]interface{}, 0, len(workflows))

	for _, workflow := range workflows {
		data := map[string]interface{}{
			"name": workflow.GetName(),
		}

		filters := make([]interface{}, 0, len(workflow.GetFilters()))
		for _, filter := range workflow.GetFilters() {
			filters = append(filters, filter.GetName())
		}

		data["filters"] = filters

		if mutator := workflow.GetMutator(); mutator != nil {
			data["mutator"] = mutator.GetName()
		}

		if handler := workflow.GetHandler(); handler != nil {
			data["handler"] = handler.GetName()
		}

		result = append(result, data)
	}

	return result
}
