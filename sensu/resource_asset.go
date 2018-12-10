package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssetCreate,
		Read:   resourceAssetRead,
		Update: resourceAssetUpdate,
		Delete: schema.RemoveFromState,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"sha512": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"namespace": resourceNamespaceSchema,
		},
	}
}

func resourceAssetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	filters := expandStringList(d.Get("filters").([]interface{}))

	asset := &types.Asset{
		ObjectMeta: types.ObjectMeta{
			Name:      name,
			Namespace: config.determineNamespace(d),
		},
		Sha512:  d.Get("sha512").(string),
		URL:     d.Get("url").(string),
		Filters: filters,
	}

	log.Printf("[DEBUG] Creating asset %s: %#v", name, asset)

	if err := asset.Validate(); err != nil {
		return fmt.Errorf("Invalid asset %s: %s", name, err)
	}

	// Not possible to delete an asset at this time,
	// so just update call update which will either create or update.
	//
	// https://github.com/sensu/sensu-go/issues/988
	if err := config.client.UpdateAsset(asset); err != nil {
		return fmt.Errorf("Error creating asset %s: %s", name, err)
	}

	d.SetId(name)

	return resourceAssetRead(d, meta)
}

func resourceAssetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Id()

	asset, err := config.client.FetchAsset(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve asset %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved asset %s: %#v", name, asset)

	d.Set("name", name)
	d.Set("namespace", asset.ObjectMeta.Namespace)
	d.Set("sha512", asset.Sha512)
	d.Set("url", asset.URL)

	if err := d.Set("filters", asset.Filters); err != nil {
		return fmt.Errorf("Error setting %s.filter: %s", name, err)
	}

	return nil
}

func resourceAssetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Id()

	asset, err := config.client.FetchAsset(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve asset %s: %s", name, err)
	}

	if d.HasChange("sha512") {
		asset.Sha512 = d.Get("sha512").(string)
	}

	if d.HasChange("url") {
		asset.URL = d.Get("url").(string)
	}

	// Filters can't really be updated right now...
	// This is buggy.
	if d.HasChange("filter") {
		filters := expandStringList(d.Get("filter").([]interface{}))
		asset.Filters = filters
	}

	if err := asset.Validate(); err != nil {
		return fmt.Errorf("Invalid asset %s: %s", name, err)
	}

	if err := config.client.UpdateAsset(asset); err != nil {
		return fmt.Errorf("Error updating asset %s: %s", name, err)
	}

	return resourceAssetRead(d, meta)
}

/*
func resourceAssetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	asset, err := config.client.FetchAsset(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve asset %s: %s", name, err)
	}

	if err := config.client.DeleteAsset(asset); err != nil {
		return fmt.Errorf("Unable to delete asset %s: %s", name, err)
	}

	return nil
}
*/
