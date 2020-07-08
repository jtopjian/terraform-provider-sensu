package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceEntity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEntityRead,

		Schema: map[string]*schema.Schema{
			// Required
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"class": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"deregister": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},

			"deregistration": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"handler": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"last_seen": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"redact": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"subscriptions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"user": &schema.Schema{
				Type:     schema.TypeString,
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

func dataSourceEntityRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Get("name").(string)

	entity, err := config.client.FetchEntity(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve entity %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved entity %s: %#v", name, entity)

	d.SetId(entity.Name)
	d.Set("namespace", entity.ObjectMeta.Namespace)
	d.Set("class", entity.EntityClass)
	d.Set("last_seen", entity.LastSeen)
	d.Set("deregister", entity.Deregister)
	d.Set("user", entity.User)

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
