package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceMutator() *schema.Resource {
	return &schema.Resource{
		Create: resourceMutatorCreate,
		Read:   resourceMutatorRead,
		Update: resourceMutatorUpdate,
		Delete: resourceMutatorDelete,
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
			"env_vars": resourceEnvVarsSchema,

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},

			"namespace": resourceNamespaceSchema,
		},
	}
}

func resourceMutatorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Get("name").(string)

	envVars := expandEnvVars(d.Get("env_vars").(map[string]interface{}))

	mutator := &types.Mutator{
		ObjectMeta: types.ObjectMeta{
			Name:      name,
			Namespace: config.determineNamespace(d),
		},
		Command: d.Get("command").(string),
		EnvVars: envVars,
		Timeout: uint32(d.Get("timeout").(int)),
	}

	log.Printf("[DEBUG] Creating mutator %s: %#v", name, mutator)

	if err := mutator.Validate(); err != nil {
		return fmt.Errorf("Invalid mutator %s: %s", name, err)
	}

	if err := config.client.CreateMutator(mutator); err != nil {
		return fmt.Errorf("Error creating mutator %s: %s", name, err)
	}

	d.SetId(name)

	return resourceMutatorRead(d, meta)
}

func resourceMutatorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	mutator, err := config.client.FetchMutator(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve mutator %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved mutator %s: %#v", name, mutator)

	d.Set("name", name)
	d.Set("namespace", mutator.ObjectMeta.Namespace)
	d.Set("command", mutator.Command)
	d.Set("timeout", mutator.Timeout)

	envVars := flattenEnvVars(mutator.EnvVars)
	if err := d.Set("env_vars", envVars); err != nil {
		return fmt.Errorf("Unable to set %s.env_vars: %s", name, err)
	}

	return nil
}

func resourceMutatorUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	mutator, err := config.client.FetchMutator(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve mutator %s: %s", name, err)
	}

	if d.HasChange("command") {
		mutator.Command = d.Get("command").(string)
	}

	if d.HasChange("env_vars") {
		envVars := expandEnvVars(d.Get("env_vars").(map[string]interface{}))
		mutator.EnvVars = envVars
	}

	if d.HasChange("timeout") {
		mutator.Timeout = uint32(d.Get("timeout").(int))
	}

	if err := mutator.Validate(); err != nil {
		return fmt.Errorf("Invalid mutator %s: %s", name, err)
	}

	if err := config.client.UpdateMutator(mutator); err != nil {
		return fmt.Errorf("Error updating mutator %s: %s", name, err)
	}

	return resourceMutatorRead(d, meta)
}

func resourceMutatorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	_, err := config.client.FetchMutator(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve mutator %s: %s", name, err)
	}

	if err := config.client.DeleteMutator(config.namespace, name); err != nil {
		return fmt.Errorf("Unable to delete mutator %s: %s", name, err)
	}

	return nil
}
