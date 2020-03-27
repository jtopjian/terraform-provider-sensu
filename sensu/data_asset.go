package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAsset() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssetRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"sha512": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceAssetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Get("name").(string)

	asset, err := config.client.FetchAsset(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve asset %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved asset %s: %#v", name, asset)

	d.SetId(asset.Name)
	d.Set("sha512", asset.Sha512)
	d.Set("url", asset.URL)
	d.Set("namespace", asset.ObjectMeta.Namespace)

	if err := d.Set("filters", asset.Filters); err != nil {
		return fmt.Errorf("Error setting %s.filter: %s", name, err)
	}

	return nil
}
