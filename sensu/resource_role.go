package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"rule": resourceRulesSchema,
		},
	}
}

func resourceRoleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	rules := expandRules(d.Get("rule").([]interface{}))

	role := &types.Role{
		Name:  name,
		Rules: rules,
	}

	log.Printf("[DEBUG] Creating role %s: %#v", name, role)

	if err := role.Validate(); err != nil {
		return fmt.Errorf("Invalid role %s: %s", name, err)
	}

	if err := config.client.CreateRole(role); err != nil {
		return fmt.Errorf("Error creating role %s: %s", name, err)
	}

	d.SetId(name)

	return resourceRoleRead(d, meta)
}

func resourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	role, err := config.client.FetchRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved role %s: %#v", name, role)

	rules := flattenRules(role.Rules)

	d.Set("name", name)
	d.Set("rule", rules)

	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := config.client.FetchRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role %s: %s", name, err)
	}

	if d.HasChange("rule") {
		o, n := d.GetChange("rule")
		oldRules := expandRules(o.([]interface{}))
		newRules := expandRules(n.([]interface{}))

		// first remove all old rules
		for _, oldRule := range oldRules {
			err := config.client.RemoveRule(name, oldRule.Type)
			if err != nil {
				return fmt.Errorf("Unable to remove rule %s from role %s: %s", oldRule.Type, name, err)
			}
		}

		// next add all roles
		for _, newRule := range newRules {
			err := config.client.AddRule(name, &newRule)
			if err != nil {
				return fmt.Errorf("Unable to add rule %s to role %s: %s", newRule.Type, name, err)
			}
		}
	}

	return resourceRoleRead(d, meta)
}

func resourceRoleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := config.client.FetchRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role %s: %s", name, err)
	}

	if err := config.client.DeleteRole(name); err != nil {
		return fmt.Errorf("Unable to delete role %s: %s", name, err)
	}

	return nil
}
