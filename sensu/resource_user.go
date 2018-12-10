package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},

			// Optional
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	groups := expandStringList(d.Get("groups").([]interface{}))
	user := &types.User{
		Username: name,
		Password: d.Get("password").(string),
		Groups:   groups,
	}

	log.Printf("[DEBUG] Creating user %s: %#v", name, user)

	if err := user.Validate(); err != nil {
		return fmt.Errorf("Invalid user %s: %s", name, err)
	}

	if err := config.client.CreateUser(user); err != nil {
		return fmt.Errorf("Error creating user %s: %s", name, err)
	}

	d.SetId(name)

	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()

	user, err := findUser(meta, name)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Retrieved user %s: %#v", name, user)

	d.Set("name", name)
	d.Set("groups", user.Groups)

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := findUser(meta, name)
	if err != nil {
		return err
	}

	if d.HasChange("groups") {
		o, n := d.GetChange("groups")
		oldGroups := expandStringList(o.([]interface{}))
		newGroups := expandStringList(n.([]interface{}))

		// first remove all old groups
		for _, oldGroup := range oldGroups {
			err := config.client.RemoveGroupFromUser(name, oldGroup)
			if err != nil {
				return fmt.Errorf("Unable to remove group %s from user %s: %s", oldGroup, name, err)
			}
		}

		// next add all groups
		for _, newGroup := range newGroups {
			err := config.client.AddGroupToUser(name, newGroup)
			if err != nil {
				return fmt.Errorf("Unable to add group %s to user %s: %s", newGroup, name, err)
			}
		}
	}

	if d.HasChange("password") {
		password := d.Get("password").(string)
		if err := config.client.UpdatePassword(name, password); err != nil {
			return fmt.Errorf("Unable to update password for user %s: %s", name, err)
		}
	}

	/*
		if d.HasChange("disabled") {
			disabled := d.Get("disabled").(bool)
			if disabled {
				if err := config.client.DisableUser(name); err != nil {
					return fmt.Errorf("Unable to disable user %s: %s", name, err)
				}
			} else {
				if err := config.client.ReinstateUser(name); err != nil {
					return fmt.Errorf("Unable to reinstate user %s: %s", name, err)
				}
			}
		}
	*/

	return resourceUserRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := findUser(meta, name)
	if err != nil {
		return err
	}

	if err := config.client.DisableUser(name); err != nil {
		return fmt.Errorf("Unable to delete user %s: %s", name, err)
	}

	return nil
}

func findUser(meta interface{}, name string) (*types.User, error) {
	config := meta.(*Config)

	users, err := config.client.ListUsers()
	if err != nil {
		return nil, fmt.Errorf("Unable to list users: %s", err)
	}

	var user types.User
	var found bool
	for _, u := range users {
		if u.Username == name {
			found = true
			user = u
		}
	}

	if !found {
		return nil, fmt.Errorf("Unable to retrieve user %s: not found", name)
	}

	return &user, nil
}
