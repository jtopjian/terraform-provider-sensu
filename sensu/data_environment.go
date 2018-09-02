package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnvironmentRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Computed
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	environment, err := config.client.FetchEnvironment(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve environment %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved environment %s: %#v", name, environment)

	d.Set("name", name)
	d.Set("description", environment.Description)
	d.Set("organization", environment.Organization)

	d.SetId(name)

	return nil
}
