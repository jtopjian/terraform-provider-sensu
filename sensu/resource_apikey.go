package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAPIKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceAPIKeyCreate,
		Read:   resourceAPIKeyRead,
		Update: resourceAPIKeyUpdate,
		Delete: resourceAPIKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAPIKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	username := d.Get("username").(string)

	log.Printf("[DEBUG] Creating apikey for user '%s'", username)

	result, err := config.client.CreateAPIKey(username)
	if err != nil {
		return fmt.Errorf("Unable to create apikey for user '%s': %s", username, err)
	}

	d.SetId(result)

	return resourceAPIKeyRead(d, meta)
}

func resourceAPIKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	result, err := config.client.FetchAPIKey(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve apikey %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved apikey %s: %#v", name, result)

	d.Set("username", result.Username)

	return nil
}

func resourceAPIKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()
	username := d.Get("username").(string)

	log.Printf("[DEBUG] Updating apikey '%s' with user '%s'", name, username)

	err := config.client.UpdateAPIKey(name, username)
	if err != nil {
		return fmt.Errorf("Unable to update apikey '%s' with user '%s': %s", name, username, err)
	}

	return resourceAPIKeyRead(d, meta)
}

func resourceAPIKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := config.client.FetchAPIKey(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve apikey %s: %s", name, err)
	}

	if err := config.client.DeleteAPIKey(name); err != nil {
		return fmt.Errorf("Unable to delete apikey %s: %s", name, err)
	}

	return nil
}
