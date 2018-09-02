package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/sensu/sensu-go/types"
)

func resourceFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceFilterCreate,
		Read:   resourceFilterRead,
		Update: resourceFilterUpdate,
		Delete: resourceFilterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{"allow", "deny"}, false),
			},

			"statements": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"when": resourceTimeWindowsSchema,
		},
	}
}

func resourceFilterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	statements := expandStringList(d.Get("statements").([]interface{}))
	when := expandTimeWindows(d.Get("when").(*schema.Set).List())

	filter := &types.EventFilter{
		Environment:  config.environment,
		Organization: config.organization,
		Name:         name,
		Action:       d.Get("action").(string),
		Statements:   statements,
		When:         &when,
	}

	log.Printf("[DEBUG] Creating filter %s: %#v", name, filter)

	if err := filter.Validate(); err != nil {
		return fmt.Errorf("Invalid filter %s: %s", name, err)
	}

	if err := config.client.CreateFilter(filter); err != nil {
		return fmt.Errorf("Unable to create filter %s: %s", name, err)
	}

	d.SetId(name)

	return resourceFilterRead(d, meta)
}

func resourceFilterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	filter, err := config.client.FetchFilter(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve filter %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved filter %s: %#v", name, filter)

	d.Set("name", name)
	d.Set("action", filter.Action)
	d.Set("statements", filter.Statements)

	when := flattenTimeWindows(filter.When)
	if err := d.Set("when", when); err != nil {
		return fmt.Errorf("Unable to set %s.when: %s", name, err)
	}

	return nil
}

func resourceFilterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	filter, err := config.client.FetchFilter(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve filter %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved filter %s: %#v", name, filter)

	if d.HasChange("action") {
		filter.Action = d.Get("action").(string)
	}

	if d.HasChange("statements") {
		statements := expandStringList(d.Get("statements").([]interface{}))
		filter.Statements = statements
	}

	if d.HasChange("when") {
		when := expandTimeWindows(d.Get("when").(*schema.Set).List())
		filter.When = &when
	}

	log.Printf("[DEBUG] Updating filter %s: %#v", name, filter)

	if err := filter.Validate(); err != nil {
		return fmt.Errorf("Invalid filter %s: %s", name, err)
	}

	if err := config.client.UpdateFilter(filter); err != nil {
		return fmt.Errorf("Unable to update filter %s: %s", name, err)
	}

	return resourceFilterRead(d, meta)
}

func resourceFilterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	filter, err := config.client.FetchFilter(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve filter %s: %s", name, err)
	}

	if err := config.client.DeleteFilter(filter); err != nil {
		return fmt.Errorf("Unable to delete filter %s: %s", name, err)
	}

	return nil
}
