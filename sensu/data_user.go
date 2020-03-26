package sensu

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Computed
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	user, err := findUser(meta, name)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Retrieved user %s: %#v", name, user)

	d.Set("name", name)
	d.Set("groups", user.Groups)
	d.Set("disabled", user.Disabled)

	d.SetId(name)

	return nil
}
