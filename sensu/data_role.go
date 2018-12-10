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

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"rule": dataSourceRulesSchema,
		},
	}
}

func dataSourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Get("name").(string)

	role, err := config.client.FetchRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved role %s: %#v", name, role)

	rules := flattenRules(role.Rules)

	d.Set("namespace", role.ObjectMeta.Namespace)
	d.Set("rule", rules)

	d.SetId(name)

	return nil
}
