package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// SecretMetadata represents the metadata section of a Sensu secret
type SecretMetadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

// SecretSpec represents the spec section of a Sensu secret
type SecretSpec struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
}

// Secret represents a Sensu Enterprise secret resource
type Secret struct {
	Type       string         `json:"type"`
	APIVersion string         `json:"api_version"`
	Metadata   SecretMetadata `json:"metadata"`
	Spec       SecretSpec     `json:"spec"`
}

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretCreate,
		Read:   resourceSecretRead,
		Update: resourceSecretUpdate,
		Delete: resourceSecretDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"secret_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The secret ID (e.g., environment variable name for Env provider)",
			},

			"secrets_provider": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the secrets provider (e.g., env)",
			},

			// Optional
			"namespace": resourceNamespaceSchema,
		},
	}
}

func resourceSecretCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)
	namespace := config.determineNamespace(d)

	secret := &Secret{
		Type:       "Secret",
		APIVersion: "secrets/v1",
		Metadata: SecretMetadata{
			Name:      name,
			Namespace: namespace,
		},
		Spec: SecretSpec{
			ID:       d.Get("secret_id").(string),
			Provider: d.Get("secrets_provider").(string),
		},
	}

	log.Printf("[DEBUG] Creating secret %s in namespace %s: %#v", name, namespace, secret)

	// Validate required fields
	if secret.Spec.ID == "" {
		return fmt.Errorf("Secret ID is required for secret %s", name)
	}
	if secret.Spec.Provider == "" {
		return fmt.Errorf("Secret provider is required for secret %s", name)
	}

	// Use the enterprise secrets API endpoint
	path := fmt.Sprintf("/api/enterprise/secrets/v1/namespaces/%s/secrets/%s", namespace, name)
	if err := config.client.Put(path, secret); err != nil {
		return fmt.Errorf("Error creating secret %s: %s", name, err)
	}

	d.SetId(name)

	return resourceSecretRead(d, meta)
}

func resourceSecretRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	namespace := config.determineNamespace(d)
	config.SaveNamespace(namespace)
	name := d.Id()

	var secret Secret
	path := fmt.Sprintf("/api/enterprise/secrets/v1/namespaces/%s/secrets/%s", namespace, name)

	err := config.client.Get(path, &secret)
	if err != nil {
		// Check if it's a not found error
		if err.Error() == "not found" {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Unable to retrieve secret %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved secret %s: %#v", name, secret)

	d.Set("name", secret.Metadata.Name)
	d.Set("namespace", secret.Metadata.Namespace)
	d.Set("secret_id", secret.Spec.ID)
	d.Set("secrets_provider", secret.Spec.Provider)

	return nil
}

func resourceSecretUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	namespace := config.determineNamespace(d)
	config.SaveNamespace(namespace)
	name := d.Id()

	// Fetch current secret
	var secret Secret
	path := fmt.Sprintf("/api/enterprise/secrets/v1/namespaces/%s/secrets/%s", namespace, name)

	err := config.client.Get(path, &secret)
	if err != nil {
		return fmt.Errorf("Unable to retrieve secret %s: %s", name, err)
	}

	// Update the ID if it changed
	if d.HasChange("secret_id") {
		secret.Spec.ID = d.Get("secret_id").(string)
	}

	// Note: secrets_provider is ForceNew, so it cannot be changed without recreating the resource

	log.Printf("[DEBUG] Updating secret %s: %#v", name, secret)

	// Validate required fields
	if secret.Spec.ID == "" {
		return fmt.Errorf("Secret ID is required for secret %s", name)
	}
	if secret.Spec.Provider == "" {
		return fmt.Errorf("Secret provider is required for secret %s", name)
	}

	// Update via PUT
	if err := config.client.Put(path, &secret); err != nil {
		return fmt.Errorf("Error updating secret %s: %s", name, err)
	}

	return resourceSecretRead(d, meta)
}

func resourceSecretDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	namespace := config.determineNamespace(d)
	config.SaveNamespace(namespace)
	name := d.Id()

	log.Printf("[DEBUG] Deleting secret %s from namespace %s", name, namespace)

	path := fmt.Sprintf("/api/enterprise/secrets/v1/namespaces/%s/secrets/%s", namespace, name)
	if err := config.client.Delete(path); err != nil {
		return fmt.Errorf("Unable to delete secret %s: %s", name, err)
	}

	return nil
}
