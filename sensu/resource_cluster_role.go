package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceClusterRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterRoleCreate,
		Read:   resourceClusterRoleRead,
		Delete: resourceClusterRoleDelete,
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

func resourceClusterRoleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	rules := expandRules(d.Get("rule").([]interface{}))

	cluster_role := &types.ClusterRole{
		ObjectMeta: types.ObjectMeta{
			Name:      name,
		},
		Rules: rules,
	}

	log.Printf("[DEBUG] Creating cluster role %s: %#v", name, cluster_role)

	if err := cluster_role.Validate(); err != nil {
		return fmt.Errorf("Invalid cluster role %s: %s", name, err)
	}

	if err := config.client.CreateClusterRole(cluster_role); err != nil {
		return fmt.Errorf("Error creating cluster role %s: %s", name, err)
	}

	d.SetId(name)

	return resourceClusterRoleRead(d, meta)
}

func resourceClusterRoleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	cluster_role, err := config.client.FetchClusterRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve cluster role %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved cluster role %s: %#v", name, cluster_role)

	rules := flattenRules(cluster_role.Rules)

	d.Set("name", name)
	d.Set("rule", rules)

	return nil
}

func resourceClusterRoleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := config.client.FetchClusterRole(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve cluster role %s: %s", name, err)
	}

	if err := config.client.DeleteClusterRole(name); err != nil {
		return fmt.Errorf("Unable to delete cluster role %s: %s", name, err)
	}

	return nil
}
