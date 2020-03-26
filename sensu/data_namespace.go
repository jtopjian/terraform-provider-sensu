package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNamespace() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNamespaceRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,
		},
	}
}

func dataSourceNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	namespace, err := config.client.FetchNamespace(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve namespace %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved namespace %s: %#v", name, namespace)

	d.Set("name", name)

	d.SetId(name)

	return nil
}
