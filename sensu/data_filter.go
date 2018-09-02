package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceFilter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFilterRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"statements": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"when": dataSourceTimeWindowsSchema,
		},
	}
}

func dataSourceFilterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	filter, err := config.client.FetchFilter(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve filter %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved filter %s: %#v", name, filter)

	d.Set("name", name)
	d.Set("action", filter.Action)
	d.Set("statements", filter.Statements)

	when := flattenTimeWindows(filter.When)
	if err := d.Set("when", when); err != nil {
		return fmt.Errorf("Unable to set %s.when: %s", name, err)
	}

	d.SetId(name)

	return nil
}
