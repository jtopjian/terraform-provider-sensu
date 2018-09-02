package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrganizationCreate,
		Read:   resourceOrganizationRead,
		Update: resourceOrganizationUpdate,
		Delete: resourceOrganizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			// Optional
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	organization := &types.Organization{
		Name:        name,
		Description: d.Get("description").(string),
	}

	log.Printf("[DEBUG] Creating organization %s: %#v", name, organization)

	if err := organization.Validate(); err != nil {
		return fmt.Errorf("Invalid organization %s: %s", name, err)
	}

	if err := config.client.CreateOrganization(organization); err != nil {
		return fmt.Errorf("Error creating organization %s: %s", name, err)
	}

	d.SetId(name)

	return resourceOrganizationRead(d, meta)
}

func resourceOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	organization, err := config.client.FetchOrganization(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve organization %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved organization %s: %#v", name, organization)

	d.Set("name", name)
	d.Set("description", organization.Description)

	return nil
}

func resourceOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	organization, err := config.client.FetchOrganization(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve organization %s: %s", name, err)
	}

	if d.HasChange("description") {
		organization.Description = d.Get("description").(string)
	}

	return resourceOrganizationRead(d, meta)
}

func resourceOrganizationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := config.client.FetchOrganization(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve organization %s: %s", name, err)
	}

	if err := config.client.DeleteOrganization(name); err != nil {
		return fmt.Errorf("Unable to delete organization %s: %s", name, err)
	}

	return nil
}
