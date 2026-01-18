package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Note: Secret type is defined in resource_secret.go

func dataSourceSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecretRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Computed
			"secret_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The secret ID (e.g., environment variable name for Env provider)",
			},

			"secrets_provider": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the secrets provider (e.g., env)",
			},

			"namespace": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The namespace the secret belongs to",
			},
		},
	}
}

func dataSourceSecretRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	namespace := config.determineNamespace(d)
	config.SaveNamespace(namespace)
	name := d.Get("name").(string)

	var secret Secret
	path := fmt.Sprintf("/api/enterprise/secrets/v1/namespaces/%s/secrets/%s", namespace, name)

	err := config.client.Get(path, &secret)
	if err != nil {
		return fmt.Errorf("Unable to retrieve secret %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved secret %s: %#v", name, secret)

	d.SetId(name)
	d.Set("name", secret.Metadata.Name)
	d.Set("namespace", secret.Metadata.Namespace)
	d.Set("secret_id", secret.Spec.ID)
	d.Set("secrets_provider", secret.Spec.Provider)

	return nil
}
