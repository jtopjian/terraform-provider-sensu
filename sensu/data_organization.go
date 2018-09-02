package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOrganizationRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Computed
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	organization, err := config.client.FetchOrganization(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve organization %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved organization %s: %#v", name, organization)

	d.Set("name", name)
	d.Set("description", organization.Description)
	d.SetId(name)

	return nil
}
