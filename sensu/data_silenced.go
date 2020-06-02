package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSilenced() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSilencedRead,

		Schema: map[string]*schema.Schema{
			// Required
			"check": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"subscription": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"begin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"expire_on_resolve": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},

			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSilencedRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	check := d.Get("check").(string)
	subscription := d.Get("subscription").(string)
	name := fmt.Sprintf("%s:%s", subscription, check)

	silenced, err := config.client.FetchSilenced(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve silenced %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved silenced %s: %#v", name, silenced)

	d.SetId(silenced.Name)
	d.Set("namespace", silenced.ObjectMeta.Namespace)
	d.Set("annotations", silenced.ObjectMeta.Annotations)
	d.Set("labels", silenced.ObjectMeta.Labels)
	d.Set("begin", silenced.Begin)
	d.Set("check", silenced.Check)
	d.Set("subscription", silenced.Subscription)
	d.Set("resolve", silenced.ExpireOnResolve)
	d.Set("expire", silenced.Expire)
	d.Set("reason", silenced.Reason)

	return nil
}
