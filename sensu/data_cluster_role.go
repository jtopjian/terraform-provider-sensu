package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceClusterRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClusterRoleRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Computed
			"rule": dataSourceRulesSchema,
		},
	}
}

func dataSourceClusterRoleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	cluster_role, err := config.client.FetchClusterRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve cluster role %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved cluster role %s: %#v", name, cluster_role)

	rules := flattenRules(cluster_role.Rules)

	d.Set("rule", rules)

	d.SetId(name)

	return nil
}
