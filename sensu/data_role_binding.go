package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/sensu/sensu-go/types"
)

func dataSourceRoleBinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRoleBindingRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"binding_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"role": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Get("name").(string)

	roleBinding, err := config.client.FetchRoleBinding(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role binding %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved role binding %s: %#v", name, roleBinding)

	d.Set("namespace", roleBinding.ObjectMeta.Namespace)
	d.Set("role", roleBinding.RoleRef.Name)

	switch roleBinding.RoleRef.Type {
	case "ClusterRole":
		d.Set("binding_type", "cluster_role")
	case "Role":
		d.Set("binding_type", "role")
	}

	users := []string{}
	groups := []string{}

	for _, subject := range roleBinding.Subjects {
		switch subject.Type {
		case types.GroupType:
			groups = append(groups, subject.Name)
		case types.UserType:
			users = append(users, subject.Name)
		}
	}

	d.Set("users", users)
	d.Set("groups", groups)

	d.SetId(name)

	return nil
}
