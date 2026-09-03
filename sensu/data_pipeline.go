package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourcePipeline() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePipelineRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"workflow": {
				Type:     schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"filters": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"mutator": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"handler": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePipelineRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Get("name").(string)

	pipeline, err := config.client.FetchPipeline(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve pipeline %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved pipeline %s: %#v", name, pipeline)

	d.Set("name", pipeline.GetName())
	d.Set("namespace", pipeline.GetNamespace())

	workflow := flattenPipelineWorkflows(pipeline.Workflows)
	if err := d.Set("workflow", workflow); err != nil {
		return fmt.Errorf("Unable to set %s.workflow: %s", name, err)
	}

	d.SetId(name)

	return nil
}
