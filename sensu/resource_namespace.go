package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceNamespaceCreate,
		Read:   resourceNamespaceRead,
		Update: resourceNamespaceUpdate,
		Delete: resourceNamespaceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,
		},
	}
}

func resourceNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	namespace := &types.Namespace{
		Name: name,
	}

	log.Printf("[DEBUG] Creating namespace %s: %#v", name, namespace)

	if err := namespace.Validate(); err != nil {
		return fmt.Errorf("Invalid namespace %s: %s", name, err)
	}

	if err := config.client.CreateNamespace(namespace); err != nil {
		return fmt.Errorf("Error creating namespace %s: %s", name, err)
	}

	d.SetId(name)

	return resourceNamespaceRead(d, meta)
}

func resourceNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	namespace, err := config.client.FetchNamespace(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve namespace %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved namespace %s: %#v", name, namespace)

	d.Set("name", name)

	return nil
}

func resourceNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	namespace, err := config.client.FetchNamespace(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve namespace %s: %s", name, err)
	}

	if err := config.client.UpdateNamespace(namespace); err != nil {
		return fmt.Errorf("Unable to delete namespace %s: %s", name, err)
	}

	if err := namespace.Validate(); err != nil {
		return fmt.Errorf("Invalid namespace %s: %s", name, err)
	}

	return resourceNamespaceRead(d, meta)
}

func resourceNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Id()

	_, err := config.client.FetchNamespace(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve namespace %s: %s", name, err)
	}

	if err := config.client.DeleteNamespace(name); err != nil {
		return fmt.Errorf("Unable to delete namespace %s: %s", name, err)
	}

	return nil
}
