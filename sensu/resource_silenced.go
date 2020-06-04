package sensu

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sensu/sensu-go/cli/commands/timeutil"
	"github.com/sensu/sensu-go/types"
)

func resourceSilenced() *schema.Resource {
	return &schema.Resource{
		Create: resourceSilencedCreate,
		Read:   resourceSilencedRead,
		Update: resourceSilencedUpdate,
		Delete: resourceSilencedDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"check": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"subscription": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"begin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "now",
			},

			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"expire_on_resolve": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"namespace": resourceNamespaceSchema,
		},
	}
}

func resourceSilencedCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	check := d.Get("check").(string)
	subscription := d.Get("subscription").(string)
	name := fmt.Sprintf("%s:%s", subscription, check)

	annotations := expandAnnotations(d.Get("annotations").(map[string]interface{}))
	labels := expandLabels(d.Get("labels").(map[string]interface{}))

	begin, err := beginConvertToTimestamp(d.Get("begin").(string))
	if err != nil {
		return err
	}

	silenced := &types.Silenced{
		ObjectMeta: types.ObjectMeta{
			Name:        name,
			Namespace:   config.determineNamespace(d),
			Annotations: annotations,
			Labels:      labels,
		},
		Begin:           begin,
		Check:           check,
		Expire:          int64(d.Get("expire").(int)),
		ExpireOnResolve: d.Get("expire_on_resolve").(bool),
		Reason:          d.Get("reason").(string),
		Subscription:    subscription,
	}

	log.Printf("[DEBUG] Creating silenced %s: %#v", name, silenced)

	if err := silenced.Validate(); err != nil {
		return fmt.Errorf("Invalid silenced %s: %s", name, err)
	}

	if err := config.client.CreateSilenced(silenced); err != nil {
		return fmt.Errorf("Unable to create silenced %s: %s", name, err)
	}

	d.SetId(name)

	return resourceSilencedRead(d, meta)
}

func resourceSilencedRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	silenced, err := config.client.FetchSilenced(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve silenced %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved silenced %s: %#v", name, silenced)

	d.Set("name", name)
	d.Set("namespace", silenced.ObjectMeta.Namespace)
	d.Set("annotations", silenced.ObjectMeta.Annotations)
	d.Set("begin", silenced.Begin)
	d.Set("check", silenced.Check)
	d.Set("expire", silenced.Expire)
	d.Set("labels", silenced.ObjectMeta.Labels)
	d.Set("reason", silenced.Reason)
	d.Set("resolve", silenced.ExpireOnResolve)
	d.Set("subscription", silenced.Subscription)

	return nil
}

func resourceSilencedUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	silenced, err := config.client.FetchSilenced(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve silenced %s: %s", name, err)
	}

	log.Printf("[DEBUG] Retrieved silenced %s: %#v", name, silenced)

	if d.HasChange("annotations") {
		silenced.ObjectMeta.Annotations = expandAnnotations(d.Get("annotations").(map[string]interface{}))
	}

	if d.HasChange("labels") {
		silenced.ObjectMeta.Labels = expandLabels(d.Get("labels").(map[string]interface{}))
	}

	if d.HasChange("begin") {
		begin, err := beginConvertToTimestamp(d.Get("begin").(string))
		if err != nil {
			return err
		}
		silenced.Begin = begin
	}
	if d.HasChange("check") {
		silenced.Check = d.Get("check").(string)
	}
	if d.HasChange("creator") {
		silenced.Creator = d.Get("creator").(string)
	}
	if d.HasChange("expire") {
		silenced.Expire = int64(d.Get("expire").(int))
	}
	if d.HasChange("reason") {
		silenced.Reason = d.Get("reason").(string)
	}
	if d.HasChange("resolve") {
		silenced.ExpireOnResolve = d.Get("expire_on_resolve").(bool)
	}
	if d.HasChange("subscription") {
		silenced.Subscription = d.Get("subscription").(string)
	}

	log.Printf("[DEBUG] Updating silenced %s: %#v", name, silenced)

	if err := silenced.Validate(); err != nil {
		return fmt.Errorf("Invalid silenced %s: %s", name, err)
	}

	if err := config.client.UpdateSilenced(silenced); err != nil {
		return fmt.Errorf("Unable to update silenced %s: %s", name, err)
	}

	return resourceSilencedRead(d, meta)
}

func resourceSilencedDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	config.SaveNamespace(config.determineNamespace(d))
	name := d.Id()

	_, err := config.client.FetchSilenced(name)
	if err != nil {
		return fmt.Errorf("Unable to retrieve silenced %s: %s", name, err)
	}

	if err := config.client.DeleteSilenced(config.namespace, name); err != nil {
		return fmt.Errorf("Unable to delete silenced %s: %s", name, err)
	}

	return nil
}

func beginConvertToTimestamp(begin string) (int64, error) {
	if begin == "now" {
		return time.Now().Unix(), nil
	} else {
		timestamp, err := timeutil.ConvertToUnix(begin)
		if err != nil {
			return 0, fmt.Errorf("Unable to convert begin date to timestamp %s: %s", begin, err)
		}
		return timestamp, nil
	}
}
