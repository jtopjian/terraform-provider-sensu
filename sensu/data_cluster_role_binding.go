package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/sensu/sensu-go/types"
)

func dataSourceClusterRoleBinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClusterRoleBindingRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			"cluster_role": &schema.Schema{
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

func dataSourceClusterRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	clusterRoleBinding, err := config.client.FetchClusterRoleBinding(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve cluster role binding %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved cluster role binding %s: %#v", name, clusterRoleBinding)

	d.Set("cluster_role", clusterRoleBinding.RoleRef.Name)
	d.Set("binding_type", "cluster_role")

	users := []string{}
	groups := []string{}

	for _, subject := range clusterRoleBinding.Subjects {
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
