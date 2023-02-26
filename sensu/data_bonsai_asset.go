package sensu

import (
	"bytes"
	"fmt"
	"log"

	goversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/sensu/core/v2"
	"github.com/sensu/sensu-go/bonsai"
	"github.com/sensu/sensu-go/cli/resource"
	"github.com/sensu/sensu-go/types/compat"
)

func dataSourceBonsaiAsset() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBonsaiAssetRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceAssetNameSchema,

			// Optional
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"labels": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"build": dataSourceAssetBuildsSchema,

			"builds": &schema.Schema{
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "Use build instead",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
		},
	}
}

func dataSourceBonsaiAssetRead(d *schema.ResourceData, meta interface{}) error {
	var version *goversion.Version
	config := meta.(*Config)
	name := d.Get("name").(string)
	ver := d.Get("version").(string)
	if ver != "" {
		name = fmt.Sprintf("%s:%s", name, ver)
	}

	bAsset, err := bonsai.NewBaseAsset(name)
	if err != nil {
		return fmt.Errorf("Error with bonsai asset name %s: %s", name, err)
	}

	if bAsset.Version != "" {
		version, err = goversion.NewVersion(bAsset.Version)
		if err != nil {
			return fmt.Errorf("Error with bonsai asset version %s:%s: %s", bAsset.Name, bAsset.Version, err)
		}
	}

	bonsaiClient := bonsai.New(bonsai.Config{})
	bonsaiAsset, err := bonsaiClient.FetchAsset(bAsset.Namespace, bAsset.Name)
	if err != nil {
		return fmt.Errorf("Error fetching bonsai asset %s/%s: %s", bAsset.Namespace, bAsset.Name, err)
	}

	bonsaiVersion, err := bonsaiAsset.BonsaiVersion(version)
	if err != nil {
		return fmt.Errorf("Error with bonsai asset version %s/%s: %s", bAsset.Namespace, bAsset.Name, err)
	}

	fullName := fmt.Sprintf("%s/%s:%s", bAsset.Namespace, bAsset.Name, bonsaiVersion.Original())
	if version == nil {
		log.Printf("[DEBUG] No bonsai asset version specified. Using latest: %s", fullName)
	}

	log.Printf("[DEBUG] Fetching bonsai asset: %s", fullName)

	asset, err := bonsaiClient.FetchAssetVersion(bAsset.Namespace, bAsset.Name, bonsaiVersion.Original())
	if err != nil {
		return fmt.Errorf("Error fetching bonsai asset %s: %s", fullName, err)
	}

	resources, err := resource.Parse(bytes.NewReader([]byte(asset)))
	if err != nil {
		return fmt.Errorf("Error parsing bonsai asset %s: %s", fullName, err)
	}

	if err := resource.Validate(resources, config.Namespace()); err != nil {
		return fmt.Errorf("Invalid bonsai asset %s: %s", fullName, err)
	}

	if len(resources) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(resources) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	var builds []map[string]interface{}
	if v, ok := resources[0].Value.(*v2.Asset); ok {
		for _, build := range v.Builds {
			b := map[string]interface{}{
				"url":     build.URL,
				"sha512":  build.Sha512,
				"filters": build.Filters,
				"headers": build.Headers,
			}

			builds = append(builds, b)
		}
	}

	d.SetId(fullName)
	if d.Set("builds", builds); err != nil {
		return fmt.Errorf("Error setting bonsai asset builds %s: %s", name, err)
	}

	if d.Set("build", builds); err != nil {
		return fmt.Errorf("Error setting bonsai asset build %s: %s", name, err)
	}

	resourceMeta := compat.GetObjectMeta(resources[0].Value)
	d.Set("labels", resourceMeta.Labels)
	d.Set("annotations", resourceMeta.Annotations)

	return nil
}
