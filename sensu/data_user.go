package sensu

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Computed
			"roles": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
	d.Set("roles", user.Roles)

	d.SetId(name)

	return nil
}
