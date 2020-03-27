package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHook() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHookRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"stdin": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceHookRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Get("name").(string)

	hook, err := config.client.FetchHook(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve hook %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved hook %s: %#v", name, hook)

	d.SetId(hook.Name)
	d.Set("namespace", hook.ObjectMeta.Namespace)
	d.Set("command", hook.Command)
	d.Set("timeout", hook.Timeout)
	d.Set("stdin", hook.Stdin)

	return nil
}
