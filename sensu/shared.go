package sensu

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/sensu/sensu-go/types"
)

// Name
var resourceNameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
	ValidateFunc: validation.StringMatch(
		regexp.MustCompile(`\A[\w\.\-]+\z`),
		"Invalid name"),
}

var dataSourceNameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
	ValidateFunc: validation.StringMatch(
		regexp.MustCompile(`\A[\w\.\-]+\z`),
		"Invalid name"),
}

// Namespace
var resourceNamespaceSchema = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	Computed: true,
	ForceNew: true,
}

// Environment Variables
var resourceEnvVarsSchema = &schema.Schema{
	Type:     schema.TypeMap,
	Optional: true,
}

var dataSourceEnvVarsSchema = &schema.Schema{
	Type:     schema.TypeMap,
	Computed: true,
}

func expandAnnotations(v map[string]interface{}) (annotations map[string]string) {

	annotations = make(map[string]string)
	for key, val := range v {
		annotations[key] = val.(string)
	}

	return
}

func expandLabels(v map[string]interface{}) (labels map[string]string) {

	labels = make(map[string]string)
	for key, val := range v {
		labels[key] = val.(string)
	}

	return
}

func expandEnvVars(v map[string]interface{}) []string {
	var envVars []string

	for key, val := range v {
		raw := val.(string)
		envVar := fmt.Sprintf("%s=%s", key, raw)
		envVars = append(envVars, envVar)
	}

	return envVars
}

func flattenEnvVars(v []string) map[string]string {
	envVars := make(map[string]string)

	for _, v := range v {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) == 2 {
			envVars[parts[0]] = parts[1]
		}
	}

	return envVars
}

// Time Window
var resourceTimeWindowsSchema = &schema.Schema{
	Type:     schema.TypeSet,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"all", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday",
				}, false),
			},
			"begin": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	},
}

var dataSourceTimeWindowsSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"begin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	},
}

func expandTimeWindows(v []interface{}) types.TimeWindowWhen {
	var timeWindows types.TimeWindowWhen

	for _, v := range v {
		timeRange := new(types.TimeWindowTimeRange)

		subdueData := v.(map[string]interface{})

		// subdue day
		var day string
		if raw, ok := subdueData["day"]; ok {
			day = strings.ToLower(raw.(string))
		}

		// begin and end
		if raw, ok := subdueData["begin"]; ok {
			timeRange.Begin = raw.(string)
		}

		if raw, ok := subdueData["end"]; ok {
			timeRange.End = raw.(string)
		}

		switch day {
		case "all":
			timeWindows.Days.All = append(timeWindows.Days.All, timeRange)
		case "monday":
			timeWindows.Days.Monday = append(timeWindows.Days.Monday, timeRange)
		case "tuesday":
			timeWindows.Days.Tuesday = append(timeWindows.Days.Tuesday, timeRange)
		case "wednesday":
			timeWindows.Days.Wednesday = append(timeWindows.Days.Wednesday, timeRange)
		case "thursday":
			timeWindows.Days.Thursday = append(timeWindows.Days.Thursday, timeRange)
		case "friday":
			timeWindows.Days.Friday = append(timeWindows.Days.Friday, timeRange)
		case "saturday":
			timeWindows.Days.Saturday = append(timeWindows.Days.Saturday, timeRange)
		case "sunday":
			timeWindows.Days.Sunday = append(timeWindows.Days.Sunday, timeRange)
		}
	}

	return timeWindows
}

func flattenTimeWindows(v *types.TimeWindowWhen) []map[string]interface{} {
	var timeWindows []map[string]interface{}

	if v == nil {
		return timeWindows
	}

	for _, v := range v.Days.All {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "all"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	for _, v := range v.Days.Monday {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "monday"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	for _, v := range v.Days.Tuesday {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "tuesday"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	for _, v := range v.Days.Wednesday {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "wednesday"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	for _, v := range v.Days.Thursday {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "thursday"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	for _, v := range v.Days.Friday {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "friday"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	for _, v := range v.Days.Saturday {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "saturday"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	for _, v := range v.Days.Sunday {
		timeWindow := make(map[string]interface{})
		timeWindow["day"] = "sunday"
		timeWindow["begin"] = v.Begin
		timeWindow["end"] = v.End
		timeWindows = append(timeWindows, timeWindow)
	}

	return timeWindows
}

// RBAC Rules
var allVerbs = []string{
	"get", "list", "create", "update", "delete", "*",
}

var resourceRulesSchema = &schema.Schema{
	Type:     schema.TypeList,
	Required: true,
	ForceNew: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"verbs": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resources": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}

var dataSourceRulesSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"verbs": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resources": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}

func expandRules(v []interface{}) []types.Rule {
	var rules []types.Rule

	for _, v := range v {
		rule := new(types.Rule)
		ruleData := v.(map[string]interface{})

		if raw, ok := ruleData["verbs"]; ok {
			for _, verb := range raw.([]interface{}) {
				rule.Verbs = append(rule.Verbs, verb.(string))
			}
		}

		if raw, ok := ruleData["resources"]; ok {
			for _, resource := range raw.([]interface{}) {
				rule.Resources = append(rule.Resources, resource.(string))
			}
		}

		if raw, ok := ruleData["resource_names"]; ok {
			for _, resourceNames := range raw.([]interface{}) {
				rule.ResourceNames = append(rule.ResourceNames, resourceNames.(string))
			}
		}

		rules = append(rules, *rule)
	}

	return rules
}

func flattenRules(v []types.Rule) []map[string]interface{} {
	var rules []map[string]interface{}

	if v == nil {
		return rules
	}

	for _, v := range v {
		rule := make(map[string]interface{})
		rule["verbs"] = v.Verbs
		rule["resources"] = v.Resources
		rule["resource_names"] = v.ResourceNames

		rules = append(rules, rule)

	}

	return rules
}

// StringList to StringSlice
func expandStringList(v []interface{}) []string {
	var vs []string

	for _, v := range v {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}

	return vs
}
