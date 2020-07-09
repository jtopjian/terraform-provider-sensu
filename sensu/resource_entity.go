package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/sensu/sensu-go/types"
)

func resourceEntity() *schema.Resource {
	return &schema.Resource{
		Create: resourceEntityCreate,
		Read:   resourceEntityRead,
		Update: resourceEntityUpdate,
		Delete: resourceEntityDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"class": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"agent", "proxy",
				}, false),
			},

			// Optional
			"namespace": resourceNamespaceSchema,

			"labels": &schema.Schema{
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: suppressDiffRedacted,
			},

			"annotations": &schema.Schema{
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: suppressDiffRedacted,
			},

			"deregister": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},

			"deregistration": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"handler": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"redact": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"subscriptions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// Computed
			"last_seen": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"system": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"os": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"platform": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"platform_family": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"platform_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"arch": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_interfaces": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"addresses": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceEntityCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Get("name").(string)
	annotations := expandStringMap(d.Get("annotations").(map[string]interface{}))
	labels := expandStringMap(d.Get("labels").(map[string]interface{}))
	subscriptions := expandStringList(d.Get("subscriptions").([]interface{}))
	deregistration := expandEntityDeregistration(d.Get("deregistration").([]interface{}))

	entity := &types.Entity{
		ObjectMeta: types.ObjectMeta{
			Name:        name,
			Namespace:   config.determineNamespace(d),
			Annotations: annotations,
			Labels:      labels,
		},
		Deregister:     d.Get("deregister").(bool),
		Deregistration: deregistration,
		EntityClass:    d.Get("class").(string),
		Subscriptions:  subscriptions,
	}

	redact := expandStringList(d.Get("redact").([]interface{}))
	if len(redact) > 0 {
		entity.Redact = redact
	}

	log.Printf("[DEBUG] Creating entity %s: %#v", name, entity)

	if err := entity.Validate(); err != nil {
		return fmt.Errorf("Invalid entity %s: %s", name, err)
	}

	if err := config.client.CreateEntity(entity); err != nil {
		return fmt.Errorf("Error creating entity %s: %s", name, err)
	}

	d.SetId(name)

	return resourceEntityRead(d, meta)
}

func resourceEntityRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	entity, err := config.client.FetchEntity(name)
	if err != nil {
		if err.Error() == "not found" {
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("Unable to retrieve entity %s: %s", name, err)
		}
	}

	log.Printf("[DEBUG] Retrieved entity %s: %#v", name, entity)

	d.Set("name", name)
	d.Set("namespace", entity.ObjectMeta.Namespace)

	d.Set("annotations", entity.ObjectMeta.Annotations)
	d.Set("class", entity.EntityClass)
	d.Set("deregister", entity.Deregister)
	d.Set("labels", entity.ObjectMeta.Labels)
	d.Set("last_seen", entity.LastSeen)
	if err := d.Set("subscriptions", entity.Subscriptions); err != nil {
		return fmt.Errorf("Unable to set %s.subscriptions: %s", name, err)
	}

	deregistration := flattenEntityDeregistration(entity.Deregistration)
	if err := d.Set("deregistration", deregistration); err != nil {
		return fmt.Errorf("Unable to set %s.deregistration: %s", name, err)
	}

	system := flattenEntitySystem(entity.System)
	if err := d.Set("system", system); err != nil {
		return fmt.Errorf("Unable to set %s.system: %s", name, err)
	}

	return nil
}

func resourceEntityUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Id()
	entity, err := config.client.FetchEntity(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve entity %s: %s", name, err)
	}

	if d.HasChange("annotations") {
		entity.ObjectMeta.Annotations = expandStringMap(d.Get("annotations").(map[string]interface{}))
	}

	if d.HasChange("class") {
		entity.EntityClass = d.Get("class").(string)
	}

	if d.HasChange("deregister") {
		entity.Deregister = d.Get("deregister").(bool)
	}

	if d.HasChange("deregistration") {
		entity.Deregistration = expandEntityDeregistration(d.Get("deregistration").([]interface{}))
	}

	if d.HasChange("labels") {
		entity.ObjectMeta.Labels = expandStringMap(d.Get("labels").(map[string]interface{}))
	}

	if d.HasChange("redact") {
		entity.Redact = expandStringList(d.Get("redact").([]interface{}))
	}

	if d.HasChange("subscriptions") {
		entity.Subscriptions = expandStringList(d.Get("subscriptions").([]interface{}))
	}

	log.Printf("[DEBUG] Updating entity %s: %#v", name, entity)

	if err := entity.Validate(); err != nil {
		return fmt.Errorf("Invalid entity %s: %s", name, err)
	}

	if err := config.client.UpdateEntity(entity); err != nil {
		return fmt.Errorf("Unable to update entity %s: %s", name, err)
	}

	return resourceEntityRead(d, meta)
}

func resourceEntityDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Id()
	_, err := config.client.FetchEntity(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve entity %s: %s", name, err)
	}

	if err := config.client.DeleteEntity(config.namespace, name); err != nil {
		return fmt.Errorf("Unable to delete entity %s: %s", name, err)
	}

	return nil
}
