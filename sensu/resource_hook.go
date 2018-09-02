package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceHookCreate,
		Read:   resourceHookRead,
		Update: resourceHookUpdate,
		Delete: resourceHookDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},

			"stdin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceHookCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	hook := &types.HookConfig{
		Name:         name,
		Environment:  config.environment,
		Organization: config.organization,
		Command:      d.Get("command").(string),
		Timeout:      uint32(d.Get("timeout").(int)),
		Stdin:        d.Get("stdin").(bool),
	}

	log.Printf("[DEBUG] Creating hook %s: %#v", name, hook)

	if err := hook.Validate(); err != nil {
		return fmt.Errorf("Invalid hook %s: %s", name, err)
	}

	if err := config.client.CreateHook(hook); err != nil {
		return fmt.Errorf("Error creating hook %s: %s", name, err)
	}

	d.SetId(name)

	return resourceHookRead(d, meta)
}

func resourceHookRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	hook, err := config.client.FetchHook(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve hook %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved hook %s: %#v", name, hook)

	d.Set("name", name)
	d.Set("command", hook.Command)
	d.Set("timeout", hook.Timeout)
	d.Set("stdin", hook.Stdin)

	return nil
}

func resourceHookUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	hook, err := config.client.FetchHook(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve hook %s: %s", name, err)
	}

	if d.HasChange("command") {
		hook.Command = d.Get("command").(string)
	}

	if d.HasChange("timeout") {
		hook.Timeout = uint32(d.Get("timeout").(int))
	}

	if d.HasChange("stdin") {
		hook.Stdin = d.Get("stdin").(bool)
	}

	if err := hook.Validate(); err != nil {
		return fmt.Errorf("Invalid hook %s: %s", name, err)
	}

	if err := config.client.UpdateHook(hook); err != nil {
		return fmt.Errorf("Error updating hook %s: %s", name, err)
	}

	return resourceHookRead(d, meta)
}

func resourceHookDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	hook, err := config.client.FetchHook(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve hook %s: %s", name, err)
	}

	if err := config.client.DeleteHook(hook); err != nil {
		return fmt.Errorf("Unable to delete hook %s: %s", name, err)
	}

	return nil
}
