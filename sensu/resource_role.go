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
		Delete: resourceRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"rule": resourceRulesSchema,

			// Optional
			"namespace": resourceNamespaceSchema,
		},
	}
}

func resourceRoleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	rules := expandRules(d.Get("rule").([]interface{}))

	role := &types.Role{
		ObjectMeta: types.ObjectMeta{
			Name:      name,
			Namespace: config.determineNamespace(d),
		},
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
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	role, err := config.client.FetchRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved role %s: %#v", name, role)

	rules := flattenRules(role.Rules)

	d.Set("name", name)
	d.Set("namespace", role.ObjectMeta.Namespace)
	d.Set("rule", rules)

	return nil
}

func resourceRoleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
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
