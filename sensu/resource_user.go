package sensu

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/cli/client"
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
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
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
	d.SetId(d.Get("name").(string))
	name := d.Id()

	_, err := findUser(meta, name)
	if err != nil {
		groups := expandStringList(d.Get("groups").([]interface{}))
		user := &types.User{
			Username: name,
			Password: d.Get("password").(string),
			Groups:   groups,
			Disabled: d.Get("disabled").(bool),
		}

		log.Printf("[DEBUG] Creating user %s: %#v", name, user)

		if err := user.Validate(); err != nil {
			return fmt.Errorf("Invalid user %s: %s", name, err)
		}

		if err := config.client.CreateUser(user); err != nil {
			return fmt.Errorf("Error creating user %s: %s", name, err)
		}
	} else {
		log.Printf("[DEBUG] Updating user %s", name)
		updateUser(d, meta)
	}

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
	d.Set("disabled", user.Disabled)

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()

	_, err := findUser(meta, name)
	if err != nil {
		return err
	}

	updateUser(d, meta)

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

	var header http.Header
	opts := client.ListOptions{}
	results := []corev2.User{}
	err := config.client.List("/api/core/v2/users", &results, &opts, &header)
	if err != nil {
		return nil, fmt.Errorf("Unable to list users: %s", err)
	}

	var user types.User
	var found bool
	for _, u := range results {
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

func updateUser(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

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

	return nil
}
