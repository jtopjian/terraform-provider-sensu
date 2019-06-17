package sensu

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCheck() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCheckRead,
		Schema: map[string]*schema.Schema{
			// Required
			"name": dataSourceNameSchema,

			// Optional
			"namespace": resourceNamespaceSchema,

			// Computed
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"cron": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"env_vars": dataSourceEnvVarsSchema,

			"handlers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"high_flap_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"check_hook": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hooks": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"trigger": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"low_flap_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"output_metric_format": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"output_metric_handlers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"proxy_entity_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
				Computed: true,
			},

			"round_robin": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},

			"runtime_assets": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"stdin": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},

			"subdue": dataSourceTimeWindowsSchema,

			"subscriptions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceCheckRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))

	name := d.Get("name").(string)
	check, err := config.client.FetchCheck(name)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Retrieved check %s: %#v", name, check)

	d.SetId(check.Name)
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

	checkHooks := flattenCheckHooks(check.CheckHooks)
	if err := d.Set("check_hook", checkHooks); err != nil {
		return fmt.Errorf("Unable to set %s.check_hook: %s", name, err)
	}

	envVars := flattenEnvVars(check.EnvVars)
	if err := d.Set("env_vars", envVars); err != nil {
		return fmt.Errorf("Unable to set %s.env_vars: %s", name, err)
	}

	if err := d.Set("handlers", check.Handlers); err != nil {
		return fmt.Errorf("Unable to set %s.handlers: %s", name, err)
	}

	if err := d.Set("subscriptions", check.Subscriptions); err != nil {
		return fmt.Errorf("Unable to set %s.subscriptions: %s", name, err)
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
		return fmt.Errorf("Unable to set %s.proxy_requests: %s", name, err)
	}

	subdues := flattenTimeWindows(check.Subdue)
	if err := d.Set("subdue", subdues); err != nil {
		return fmt.Errorf("Unable to set %s.subdue: %s", name, err)
	}

	return nil
}
