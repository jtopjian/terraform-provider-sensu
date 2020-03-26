package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sensu/sensu-go/types"
)

func resourceCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceCheckCreate,
		Read:   resourceCheckRead,
		Update: resourceCheckUpdate,
		Delete: resourceCheckDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"name": resourceNameSchema,

			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"subscriptions": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// Optional
			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"cron": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"interval"},
			},

			"env_vars": resourceEnvVarsSchema,

			"handlers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"high_flap_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"check_hook": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hook": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"trigger": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"interval": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"cron"},
			},

			"low_flap_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"output_metric_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"output_metric_handlers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"proxy_entity_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			/*
				"proxy_requests": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"entity_attributes": &schema.Schema{
								Type:     schema.TypeList,
								Optional: true,
								Elem:     &schema.Schema{Type: schema.TypeString},
							},
							"splay": &schema.Schema{
								Type:     schema.TypeBool,
								Optional: true,
							},
							"splay_coverage": &schema.Schema{
								Type:     schema.TypeInt,
								Optional: true,
							},
						},
					},
				},
			*/

			"publish": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},

			"round_robin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},

			"runtime_assets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"stdin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},

			"subdue": resourceTimeWindowsSchema,

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"namespace": resourceNamespaceSchema,
		},
	}
}

func resourceCheckCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	name := d.Get("name").(string)

	// One of cron or interval is required
	cron := d.Get("cron").(string)
	interval := d.Get("interval").(int)
	if cron == "" && interval == 0 {
		return fmt.Errorf("Must specify one of cron or interval for check %s", name)
	}

	// standard string lists
	subscriptions := expandStringList(d.Get("subscriptions").([]interface{}))
	handlers := expandStringList(d.Get("handlers").([]interface{}))
	outputMetricHandlers := expandStringList(d.Get("output_metric_handlers").([]interface{}))
	runtimeAssets := expandStringList(d.Get("runtime_assets").([]interface{}))

	// detailed structures
	//proxyRequests := expandCheckProxyRequests(d.Get("proxy_requests").([]interface{}))
	envVars := expandEnvVars(d.Get("env_vars").(map[string]interface{}))
	subdues := expandTimeWindows(d.Get("subdue").(*schema.Set).List())
	annotations := expandAnnotations(d.Get("annotations").(map[string]interface{}))
	labels := expandLabels(d.Get("labels").(map[string]interface{}))

	// Using partial to resume hook configuration if there's a failure.
	d.Partial(true)

	check := &types.CheckConfig{
		ObjectMeta: types.ObjectMeta{
			Name:        name,
			Namespace:   config.determineNamespace(d),
			Annotations: annotations,
			Labels:      labels,
		},
		Command:              d.Get("command").(string),
		Subscriptions:        subscriptions,
		Cron:                 cron,
		EnvVars:              envVars,
		Handlers:             handlers,
		HighFlapThreshold:    uint32(d.Get("high_flap_threshold").(int)),
		Interval:             uint32(interval),
		LowFlapThreshold:     uint32(d.Get("low_flap_threshold").(int)),
		OutputMetricFormat:   d.Get("output_metric_format").(string),
		OutputMetricHandlers: outputMetricHandlers,
		ProxyEntityName:      d.Get("proxy_entity_name").(string),
		//ProxyRequests:        &proxyRequests,
		Publish:       d.Get("publish").(bool),
		RoundRobin:    d.Get("round_robin").(bool),
		RuntimeAssets: runtimeAssets,
		Stdin:         d.Get("stdin").(bool),
		Subdue:        &subdues,
		Timeout:       uint32(d.Get("timeout").(int)),
		Ttl:           int64(d.Get("ttl").(int)),
	}

	log.Printf("[DEBUG] Creating check %s: %#v", name, check)

	if err := check.Validate(); err != nil {
		return fmt.Errorf("Invalid check %s: %s", name, err)
	}

	if err := config.client.CreateCheck(check); err != nil {
		return fmt.Errorf("Error creating check %s: %s", name, err)
	}

	d.SetId(name)

	d.SetPartial("namespace")
	d.SetPartial("command")
	d.SetPartial("subscriptions")
	d.SetPartial("cron")
	d.SetPartial("env_vars")
	d.SetPartial("annotations")
	d.SetPartial("labels")
	d.SetPartial("handlers")
	d.SetPartial("high_flap_threshold")
	d.SetPartial("interval")
	d.SetPartial("low_flap_threshold")
	d.SetPartial("output_metric_format")
	d.SetPartial("output_metric_handlers")
	d.SetPartial("proxy_entity_name")
	//d.SetPartial("proxy_requests")
	d.SetPartial("publish")
	d.SetPartial("round_robin")
	d.SetPartial("runtime_assets")
	d.SetPartial("stdin")
	d.SetPartial("timeout")
	d.SetPartial("ttl")

	config.SaveNamespace(config.determineNamespace(d))

	check, err := config.client.FetchCheck(name)
	if err != nil {
		return fmt.Errorf("Unable to retrive check %s after creation: %s", name, err)
	}

	// Add any hooks
	checkHooks := expandCheckHooks(d.Get("check_hook").(*schema.Set).List())
	for _, checkHook := range checkHooks {
		if err := checkHook.Validate(); err != nil {
			return fmt.Errorf("Invalid check_hook for %s: %s", name, err)
		}

		if err := config.client.AddCheckHook(check, &checkHook); err != nil {
			return fmt.Errorf("Error adding check_hooks to check %s: %s", name, err)
		}
	}

	d.Partial(false)

	return resourceCheckRead(d, meta)
}

func resourceCheckRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Id()
	check, err := config.client.FetchCheck(name)
	if err != nil {
		if err.Error() == "not found" {
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("Unable to retrieve check %s: %s", name, err)
		}
	}

	log.Printf("[DEBUG] Retrieved check %s: %#v", name, check)

	d.Set("name", name)
	d.Set("namespace", check.ObjectMeta.Namespace)
	d.Set("command", check.Command)
	d.Set("annotations", check.ObjectMeta.Annotations)
	d.Set("labels", check.ObjectMeta.Labels)
	d.Set("cron", check.Cron)
	d.Set("high_flap_threshold", check.HighFlapThreshold)
	d.Set("interval", check.Interval)
	d.Set("low_flap_threshold", check.LowFlapThreshold)
	d.Set("output_metric_format", check.OutputMetricFormat)
	d.Set("proxy_entity_name", check.ProxyEntityName)
	d.Set("publish", check.Publish)
	d.Set("round_robin", check.RoundRobin)
	d.Set("stdin", check.Stdin)
	d.Set("timeout", check.Timeout)
	d.Set("ttl", check.Ttl)

	envVars := flattenEnvVars(check.EnvVars)
	if err := d.Set("env_vars", envVars); err != nil {
		return fmt.Errorf("Unable to set %s.env_vars: %s", name, err)
	}

	if err := d.Set("handlers", check.Handlers); err != nil {
		return fmt.Errorf("Unable to set %s.handlers: %s", name, err)
	}

	checkHooks := flattenCheckHooks(check.CheckHooks)
	if err := d.Set("check_hook", checkHooks); err != nil {
		return fmt.Errorf("Unable to set %s.check_hook: %s", name, err)
	}

	if err := d.Set("output_metric_handlers", check.OutputMetricHandlers); err != nil {
		return fmt.Errorf("Unable to set %s.output_metric_handlers: %s", name, err)
	}

	/*
		proxyRequests := flattenCheckProxyRequests(check.ProxyRequests)
		if err := d.Set("proxy_requests", proxyRequests); err != nil {
			return fmt.Errorf("Unable to set %s.proxy_requests: %s", name, err)
		}
	*/

	if err := d.Set("runtime_assets", check.RuntimeAssets); err != nil {
		return fmt.Errorf("Unable to set %s.runtime_assets: %s", name, err)
	}

	if err := d.Set("subscriptions", check.Subscriptions); err != nil {
		return fmt.Errorf("Unable to set %s.subscriptions: %s", name, err)
	}

	subdues := flattenTimeWindows(check.Subdue)
	if err := d.Set("subdue", subdues); err != nil {
		return fmt.Errorf("Unable to set %s.subdue: %s", name, err)
	}

	return nil
}

func resourceCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Id()
	check, err := config.client.FetchCheck(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve check %s: %s", name, err)
	}

	if d.HasChange("command") {
		check.Command = d.Get("command").(string)
	}

	if d.HasChange("cron") {
		check.Cron = d.Get("cron").(string)
	}

	if d.HasChange("handlers") {
		handlers := expandStringList(d.Get("handlers").([]interface{}))
		check.Handlers = handlers
	}

	if d.HasChange("annotations") {
		check.ObjectMeta.Annotations = expandAnnotations(d.Get("annotations").(map[string]interface{}))
	}

	if d.HasChange("labels") {
		check.ObjectMeta.Labels = expandLabels(d.Get("labels").(map[string]interface{}))
	}

	if d.HasChange("high_flap_threshold") {
		check.HighFlapThreshold = uint32(d.Get("high_flap_threshold").(int))
	}

	if d.HasChange("interval") {
		check.Interval = uint32(d.Get("interval").(int))
	}

	if d.HasChange("low_flap_threshold") {
		check.LowFlapThreshold = uint32(d.Get("low_flap_threshold").(int))
	}

	if d.HasChange("metric_format") {
		check.OutputMetricFormat = d.Get("output_metric_format").(string)
	}

	if d.HasChange("metric_handlers") {
		metricHandlers := expandStringList(d.Get("output_metric_handlers").([]interface{}))
		check.OutputMetricHandlers = metricHandlers
	}

	if d.HasChange("proxy_entity_name") {
		check.ProxyEntityName = d.Get("proxy_entity_name").(string)
	}

	/*
		if d.HasChange("proxy_requests") {
			proxyRequests := expandCheckProxyRequests(d.Get("proxy_requests").([]interface{}))
			check.ProxyRequests = &proxyRequests
		}
	*/

	if d.HasChange("publish") {
		check.Publish = d.Get("publish").(bool)
	}

	if d.HasChange("round_robin") {
		check.RoundRobin = d.Get("round_robin").(bool)
	}

	if d.HasChange("runtime_assets") {
		runtimeAssets := expandStringList(d.Get("runtime_assets").([]interface{}))
		check.RuntimeAssets = runtimeAssets
	}

	if d.HasChange("stdin") {
		check.Stdin = d.Get("stdin").(bool)
	}

	if d.HasChange("subscriptions") {
		subscriptions := expandStringList(d.Get("subscriptions").([]interface{}))
		check.Subscriptions = subscriptions
	}

	if d.HasChange("subdue") {
		subdues := expandTimeWindows(d.Get("subdue").(*schema.Set).List())
		check.Subdue = &subdues
	}

	if d.HasChange("timeout") {
		check.Timeout = uint32(d.Get("timeout").(int))
	}

	if d.HasChange("ttl") {
		check.Ttl = int64(d.Get("ttl").(int))
	}

	log.Printf("[DEBUG] Updating check %s: %#v", name, check)

	if err := config.client.UpdateCheck(check); err != nil {
		return fmt.Errorf("Unable to update check %s: %s", name, err)
	}

	check, err = config.client.FetchCheck(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve updated check %s: %s", name, err)
	}

	// Update hooks
	if d.HasChange("check_hook") {
		old, new := d.GetChange("check_hook")
		oldHookLists := expandCheckHooks(old.([]interface{}))
		newHookLists := expandCheckHooks(new.([]interface{}))

		// remove all old hooks
		for _, v := range oldHookLists {
			if len(v.Hooks) > 0 {
				if err := config.client.RemoveCheckHook(check, v.Type, v.Hooks[0]); err != nil {
					return fmt.Errorf("Error removing hook %s from check %s: %s", v.Hooks[0], name, err)
				}
			}
		}

		// add new hooks
		for _, v := range newHookLists {
			if err := config.client.AddCheckHook(check, &v); err != nil {
				return fmt.Errorf("Error updating check_hooks for check %s: %s", name, err)
			}
		}
	}

	return resourceCheckRead(d, meta)
}

func resourceCheckDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Id()
	_, err := config.client.FetchCheck(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve check %s: %s", name, err)
	}

	if err := config.client.DeleteCheck(config.namespace, name); err != nil {
		return fmt.Errorf("Unable to delete check %s: %s", name, err)
	}

	return nil
}

func expandCheckHooks(v []interface{}) []types.HookList {
	var hookLists []types.HookList

	for _, v := range v {
		var hookList types.HookList

		hookData := v.(map[string]interface{})
		if raw, ok := hookData["hook"]; ok {
			hookList.Hooks = []string{raw.(string)}
		}

		if raw, ok := hookData["trigger"]; ok {
			hookList.Type = raw.(string)
		}

		hookLists = append(hookLists, hookList)
	}

	return hookLists
}

func flattenCheckHooks(v []types.HookList) []map[string]interface{} {
	var hookLists []map[string]interface{}

	log.Printf("[DEBUG] hooks: %#v", v)

	for _, hl := range v {
		for _, hook := range hl.Hooks {
			hookList := make(map[string]interface{})
			hookList["hook"] = hook
			hookList["trigger"] = hl.Type
			hookLists = append(hookLists, hookList)
		}
	}

	return hookLists
}

/*
func expandCheckProxyRequests(v []interface{}) types.ProxyRequests {
	var proxyRequests types.ProxyRequests

	for _, v := range v {
		proxyData := v.(map[string]interface{})

		// entity attributes
		if raw, ok := proxyData["entity_attributes"]; ok {
			list := raw.([]interface{})
			proxyRequests.EntityAttributes = expandStringList(list)
		}

		// splay
		if raw, ok := proxyData["splay"]; ok {
			proxyRequests.Splay = raw.(bool)
		}

		// splay coverage
		if raw, ok := proxyData["splay_coverage"]; ok {
			proxyRequests.SplayCoverage = uint32(raw.(int))
		}
	}

	return proxyRequests
}

func flattenCheckProxyRequests(v *types.ProxyRequests) []map[string]interface{} {
	var proxyRequests []map[string]interface{}
	pr := make(map[string]interface{})

	if len(v.EntityAttributes) > 0 {
		pr["entity_attributes"] = v.EntityAttributes
		pr["splay"] = v.Splay
		pr["splay_coverage"] = v.SplayCoverage
		proxyRequests = append(proxyRequests, pr)
	}

	return proxyRequests
}
*/
