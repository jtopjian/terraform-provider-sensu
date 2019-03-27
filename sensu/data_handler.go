package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceHandler() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHandlerRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"env_vars": dataSourceEnvVarsSchema,

			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"handlers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"runtime_assets": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"mutator": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"socket": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceHandlerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Get("name").(string)
	handler, err := config.client.FetchHandler(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve handler %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved handler %s: %#v", name, handler)

	d.SetId(name)
	d.Set("namespace", handler.ObjectMeta.Namespace)
	d.Set("command", handler.Command)
	d.Set("filters", handler.Filters)
	d.Set("handlers", handler.Handlers)
	d.Set("runtime_assets", handler.RuntimeAssets)
	d.Set("mutator", handler.Mutator)
	d.Set("timeout", handler.Timeout)
	d.Set("type", handler.Type)

	socket := flattenHandlerSocket(handler.Socket)
	if err := d.Set("socket", socket); err != nil {
		return fmt.Errorf("Unable to set %s.socket: %s", name, err)
	}

	envVars := flattenEnvVars(handler.EnvVars)
	if err := d.Set("env_vars", envVars); err != nil {
		return fmt.Errorf("Unable to set %s.env_vars: %s", name, err)
	}

	return nil
}
