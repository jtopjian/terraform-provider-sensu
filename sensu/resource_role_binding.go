package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/types"
)

func resourceRoleBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleBindingCreate,
		Read:   resourceRoleBindingRead,
		Delete: resourceRoleBindingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"binding_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"role", "cluster_role",
				}, false),
			},

			"role": &schema.Schema{
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

			// Optional
			"namespace": resourceNamespaceSchema,
		},
	}
}

func resourceRoleBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	roleType := d.Get("binding_type").(string)
	roleName := d.Get("role").(string)
	users := d.Get("users").([]interface{})
	groups := d.Get("groups").([]interface{})

	var rt string
	switch roleType {
	case "cluster_role":
		rt = "ClusterRole"
	case "role":
		rt = "Role"
	}

	roleBinding := &types.RoleBinding{
		ObjectMeta: types.ObjectMeta{
			Name:      name,
			Namespace: config.determineNamespace(d),
		},
		RoleRef: v2.RoleRef{
			Type: rt,
			Name: roleName,
		},
	}

	if len(groups) == 0 && len(users) == 0 {
		return fmt.Errorf("at least one group or user must be specified")
	}

	for _, group := range groups {
		roleBinding.Subjects = append(roleBinding.Subjects,
			v2.Subject{
				Type: "Group",
				Name: group.(string),
			})
	}

	for _, user := range users {
		roleBinding.Subjects = append(roleBinding.Subjects,
			v2.Subject{
				Type: "User",
				Name: user.(string),
			},
		)
	}

	log.Printf("[DEBUG] Creating role binding %s: %#v", name, roleBinding)

	if err := roleBinding.Validate(); err != nil {
		return fmt.Errorf("Invalid role binding %s: %s", name, err)
	}

	if err := config.client.CreateRoleBinding(roleBinding); err != nil {
		return fmt.Errorf("Error creating role binding %s: %s", name, err)
	}

	d.SetId(name)

	return resourceRoleBindingRead(d, meta)
}

func resourceRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	roleBinding, err := config.client.FetchRoleBinding(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role binding %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved role binding %s: %#v", name, roleBinding)

	d.Set("name", roleBinding.Name)
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

	return nil
}

func resourceRoleBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	_, err := config.client.FetchRoleBinding(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve role binding %s: %s", name, err)
	}

	if err := config.client.DeleteRoleBinding(config.namespace, name); err != nil {
		return fmt.Errorf("Unable to delete role binding %s: %s", name, err)
	}

	return nil
}
