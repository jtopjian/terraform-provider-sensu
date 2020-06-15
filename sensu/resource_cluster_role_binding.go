package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/types"
)

func resourceClusterRoleBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterRoleBindingCreate,
		Read:   resourceClusterRoleBindingRead,
		Delete: resourceClusterRoleBindingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"cluster_role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"users": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceClusterRoleBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	roleName := d.Get("cluster_role").(string)
	users := d.Get("users").([]interface{})
	groups := d.Get("groups").([]interface{})

	clusterRoleBinding := &types.ClusterRoleBinding{
		ObjectMeta: types.ObjectMeta{
			Name: name,
		},
		RoleRef: v2.RoleRef{
			Type: "ClusterRole",
			Name: roleName,
		},
	}

	if len(groups) == 0 && len(users) == 0 {
		return fmt.Errorf("at least one group or user must be specified")
	}

	for _, group := range groups {
		clusterRoleBinding.Subjects = append(clusterRoleBinding.Subjects,
			v2.Subject{
				Type: "Group",
				Name: group.(string),
			})
	}

	for _, user := range users {
		clusterRoleBinding.Subjects = append(clusterRoleBinding.Subjects,
			v2.Subject{
				Type: "User",
				Name: user.(string),
			},
		)
	}

	log.Printf("[DEBUG] Creating cluster role binding %s: %#v", name, clusterRoleBinding)

	if err := clusterRoleBinding.Validate(); err != nil {
		return fmt.Errorf("Invalid cluster role binding %s: %s", name, err)
	}

	if err := config.client.CreateClusterRoleBinding(clusterRoleBinding); err != nil {
		return fmt.Errorf("Error creating cluster role binding %s: %s", name, err)
	}

	d.SetId(name)

	return resourceClusterRoleBindingRead(d, meta)
}

func resourceClusterRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	clusterRoleBinding, err := config.client.FetchClusterRoleBinding(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve cluster role binding %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved cluster role binding %s: %#v", name, clusterRoleBinding)

	d.Set("name", clusterRoleBinding.Name)
	d.Set("cluster_role", clusterRoleBinding.RoleRef.Name)

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

	return nil
}

func resourceClusterRoleBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := config.client.FetchClusterRoleBinding(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve cluster role binding %s: %s", name, err)
	}

	if err := config.client.DeleteClusterRoleBinding(name); err != nil {
		return fmt.Errorf("Unable to delete cluster role binding %s: %s", name, err)
	}

	return nil
}
