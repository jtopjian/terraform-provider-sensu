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
			"name": dataSourceAssetNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"build": dataSourceAssetBuildsSchema,

			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

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

			"headers": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
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
	d.Set("headers", asset.Headers)
	d.Set("namespace", asset.ObjectMeta.Namespace)
	d.Set("annotations", asset.ObjectMeta.Annotations)
	d.Set("labels", asset.ObjectMeta.Labels)

	if err := d.Set("filters", asset.Filters); err != nil {
		return fmt.Errorf("Error setting %s.filter: %s", name, err)
	}

	builds := flattenAssetBuilds(asset.Builds)
	if err := d.Set("build", builds); err != nil {
		return fmt.Errorf("Error setting %s.build: %s", name, err)
	}

	return nil
}
