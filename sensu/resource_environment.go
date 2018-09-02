package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnvironmentCreate,
		Read:   resourceEnvironmentRead,
		Update: resourceEnvironmentUpdate,
		Delete: resourceEnvironmentDelete,
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

			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
		},
	}
}

func resourceEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)
	organization := d.Get("organization").(string)

	environment := &types.Environment{
		Name:         name,
		Description:  d.Get("description").(string),
		Organization: organization,
	}

	log.Printf("[DEBUG] Creating environment %s: %#v", name, environment)

	if err := environment.Validate(); err != nil {
		return fmt.Errorf("Invalid environment %s: %s", name, err)
	}

	if err := config.client.CreateEnvironment(organization, environment); err != nil {
		return fmt.Errorf("Error creating environment %s: %s", name, err)
	}

	d.SetId(name)

	return resourceEnvironmentRead(d, meta)
}

func resourceEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	environment, err := config.client.FetchEnvironment(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve environment %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved environment %s: %#v", name, environment)

	d.Set("name", name)
	d.Set("description", environment.Description)
	d.Set("organization", environment.Organization)

	return nil
}

func resourceEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	environment, err := config.client.FetchEnvironment(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve environment %s: %s", name, err)
	}

	if d.HasChange("description") {
		environment.Description = d.Get("description").(string)
	}

	if d.HasChange("organization") {
		environment.Organization = d.Get("organization").(string)
	}

	if err := config.client.UpdateEnvironment(environment); err != nil {
		return fmt.Errorf("Unable to delete environment %s: %s", name, err)
	}

	if err := environment.Validate(); err != nil {
		return fmt.Errorf("Invalid environment %s: %s", name, err)
	}

	return resourceEnvironmentRead(d, meta)
}

func resourceEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()
	organization := d.Get("organization").(string)

	_, err := config.client.FetchEnvironment(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve environment %s: %s", name, err)
	}

	if err := config.client.DeleteEnvironment(organization, name); err != nil {
		return fmt.Errorf("Unable to delete environment %s: %s", name, err)
	}

	return nil
}
