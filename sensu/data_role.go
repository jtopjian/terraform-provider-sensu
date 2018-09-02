package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRoleRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Computed
			"rule": dataSourceRulesSchema,
		},
	}
}

func dataSourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	role, err := config.client.FetchRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved role %s: %#v", name, role)

	rules := flattenRules(role.Rules)

	d.Set("name", name)
	d.Set("rule", rules)

	d.SetId(name)

	return nil
}
