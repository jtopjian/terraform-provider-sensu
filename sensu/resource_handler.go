package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/sensu/sensu-go/types"
)

func resourceHandler() *schema.Resource {
	return &schema.Resource{
		Create: resourceHandlerCreate,
		Read:   resourceHandlerRead,
		Update: resourceHandlerUpdate,
		Delete: resourceHandlerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"pipe", "tcp", "udp", "set",
				}, false),
			},

			// Optional
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"env_vars": resourceEnvVarsSchema,

			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"handlers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"mutator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"socket": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceHandlerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Get("name").(string)

	// TODO: secondary validation for different combinations.

	// standard string lists
	filters := expandStringList(d.Get("filters").([]interface{}))
	handlers := expandStringList(d.Get("handlers").([]interface{}))

	// detailed structures
	envVars := expandEnvVars(d.Get("env_vars").(map[string]interface{}))

	handler := &types.Handler{
		Name:         name,
		Environment:  config.environment,
		Organization: config.organization,
		Command:      d.Get("command").(string),
		EnvVars:      envVars,
		Handlers:     handlers,
		Filters:      filters,
		Mutator:      d.Get("mutator").(string),
		Timeout:      uint32(d.Get("timeout").(int)),
		Type:         d.Get("type").(string),
	}

	if v, ok := d.GetOk("socket"); ok {
		handler.Socket = expandHandlerSocket(v.([]interface{}))
	}

	log.Printf("[DEBUG] Creating handler %s: %#v", name, handler)

	if err := handler.Validate(); err != nil {
		return fmt.Errorf("Invalid handler %s: %s", name, err)
	}

	if err := config.client.CreateHandler(handler); err != nil {
		return fmt.Errorf("Error creating handler %s: %s", name, err)
	}

	d.SetId(name)

	return resourceHandlerRead(d, meta)
}

func resourceHandlerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Id()
	handler, err := config.client.FetchHandler(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve handler %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved handler %s: %#v", name, handler)

	d.Set("name", name)
	d.Set("command", handler.Command)
	d.Set("filters", handler.Filters)
	d.Set("handlers", handler.Handlers)
	d.Set("mutator", handler.Mutator)
	d.Set("timeout", handler.Timeout)
	d.Set("type", handler.Type)

	socket := flattenHandlerSocket(handler.Socket)
	if err := d.Set("socket", socket); err != nil {
		return fmt.Errorf("Unable to set %s.socket: %s", name, err)
	}

	envVars := flattenEnvVars(handler.EnvVars)
	if err := d.Set("env_vars", envVars); err != nil {
		return fmt.Errorf("Unable to set %s.env_vars: %s", name, err)
	}

	return nil
}

func resourceHandlerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Id()
	handler, err := config.client.FetchHandler(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve handler %s: %s", name, err)
	}

	if d.HasChange("command") {
		handler.Command = d.Get("command").(string)
	}

	if d.HasChange("env_vars") {
		envVars := expandEnvVars(d.Get("env_vars").(map[string]interface{}))
		handler.EnvVars = envVars
	}

	if d.HasChange("filters") {
		filters := expandStringList(d.Get("filters").([]interface{}))
		handler.Filters = filters
	}

	if d.HasChange("handlers") {
		handlers := expandStringList(d.Get("handlers").([]interface{}))
		handler.Handlers = handlers
	}

	if d.HasChange("mutator") {
		handler.Mutator = d.Get("mutator").(string)
	}

	if d.HasChange("socket") {
		socket := expandHandlerSocket(d.Get("socket").([]interface{}))
		handler.Socket = socket
	}

	if d.HasChange("timeout") {
		handler.Timeout = uint32(d.Get("timeout").(int))
	}

	if d.HasChange("type") {
		handler.Type = d.Get("type").(string)
	}

	if err := handler.Validate(); err != nil {
		return fmt.Errorf("Invalid handler %s: %s", name, err)
	}

	if err := config.client.UpdateHandler(handler); err != nil {
		return fmt.Errorf("Error updating handler %s: %s", name, err)
	}

	return resourceHandlerRead(d, meta)
}

func resourceHandlerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Id()
	handler := &types.Handler{Name: name}

	if err := config.client.DeleteHandler(handler); err != nil {
		return fmt.Errorf("Unable to delete handler %s: %s", name, err)
	}

	return nil
}

func expandHandlerSocket(v []interface{}) *types.HandlerSocket {
	socket := types.HandlerSocket{}

	for _, v := range v {
		socketData := v.(map[string]interface{})
		if raw, ok := socketData["host"]; ok {
			socket.Host = raw.(string)
		}

		if raw, ok := socketData["port"]; ok {
			socket.Port = uint32(raw.(int))
		}
	}

	return &socket
}

func flattenHandlerSocket(v *types.HandlerSocket) []map[string]interface{} {
	var sockets []map[string]interface{}
	socket := make(map[string]interface{})

	if v != nil {
		socket["host"] = v.Host
		socket["port"] = v.Port

		sockets = append(sockets, socket)
	}

	return sockets
}
