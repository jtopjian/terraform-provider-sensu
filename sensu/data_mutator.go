package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMutator() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMutatorRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"env_vars": dataSourceEnvVarsSchema,

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceMutatorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Get("name").(string)

	mutator, err := config.client.FetchMutator(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve mutator %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved mutator %s: %#v", name, mutator)

	d.SetId(name)
	d.Set("namespace", mutator.ObjectMeta.Namespace)
	d.Set("command", mutator.Command)
	d.Set("timeout", mutator.Timeout)

	envVars := flattenEnvVars(mutator.EnvVars)
	if err := d.Set("env_vars", envVars); err != nil {
		return fmt.Errorf("Unable to set %s.env_vars: %s", name, err)
	}

	return nil
}
