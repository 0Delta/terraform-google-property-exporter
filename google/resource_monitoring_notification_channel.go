// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var sensitiveLabels = []string{"auth_token", "service_key", "password"}

func sensitiveLabelCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	for _, sl := range sensitiveLabels {
		mapLabel := diff.Get("labels." + sl).(string)
		authLabel := diff.Get("sensitive_labels.0." + sl).(string)
		if mapLabel != "" && authLabel != "" {
			return fmt.Errorf("Sensitive label [%s] cannot be set in both `labels` and the `sensitive_labels` block.", sl)
		}
	}
	return nil
}

func ResourceMonitoringNotificationChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringNotificationChannelCreate,
		Read:   resourceMonitoringNotificationChannelRead,
		Update: resourceMonitoringNotificationChannelUpdate,
		Delete: resourceMonitoringNotificationChannelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringNotificationChannelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: sensitiveLabelCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.`,
				Default:     true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Configuration fields that define the channel and its behavior. The
permissible and required labels are specified in the
NotificationChannelDescriptor corresponding to the type field.

Labels with sensitive data are obfuscated by the API and therefore Terraform cannot
determine if there are upstream changes to these fields. They can also be configured via
the sensitive_labels block, but cannot be configured in both places.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"sensitive_labels": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Different notification type behaviors are configured primarily using the the 'labels' field on this
resource. This block contains the labels which contain secrets or passwords so that they can be marked
sensitive and hidden from plan output. The name of the field, eg: password, will be the key
in the 'labels' map in the api request.

Credentials may not be specified in both locations and will cause an error. Changing from one location
to a different credential configuration in the config will require an apply to update state.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_token": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  `An authorization token for a notification channel. Channel types that support this field include: slack`,
							Sensitive:    true,
							ExactlyOneOf: []string{"sensitive_labels.0.auth_token", "sensitive_labels.0.password", "sensitive_labels.0.service_key"},
						},
						"password": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  `An password for a notification channel. Channel types that support this field include: webhook_basicauth`,
							Sensitive:    true,
							ExactlyOneOf: []string{"sensitive_labels.0.auth_token", "sensitive_labels.0.password", "sensitive_labels.0.service_key"},
						},
						"service_key": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  `An servicekey token for a notification channel. Channel types that support this field include: pagerduty`,
							Sensitive:    true,
							ExactlyOneOf: []string{"sensitive_labels.0.auth_token", "sensitive_labels.0.password", "sensitive_labels.0.service_key"},
						},
					},
				},
			},
			"user_labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The full REST resource name for this channel. The syntax is:
projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
The [CHANNEL_ID] is automatically assigned by the server on creation.`,
			},
			"verification_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.`,
			},
			"force_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `If true, the notification channel will be deleted regardless
of its use in alert policies (the policies will be updated
to remove the channel). If false, channels that are still
referenced by an existing alerting policy will fail to be
deleted in a delete operation.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceMonitoringNotificationChannelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandMonitoringNotificationChannelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	typeProp, err := expandMonitoringNotificationChannelType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	userLabelsProp, err := expandMonitoringNotificationChannelUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_labels"); !isEmptyValue(reflect.ValueOf(userLabelsProp)) && (ok || !reflect.DeepEqual(v, userLabelsProp)) {
		obj["userLabels"] = userLabelsProp
	}
	descriptionProp, err := expandMonitoringNotificationChannelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringNotificationChannelDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandMonitoringNotificationChannelEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); ok || !reflect.DeepEqual(v, enabledProp) {
		obj["enabled"] = enabledProp
	}

	obj, err = resourceMonitoringNotificationChannelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "stackdriver/notifications/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{MonitoringBasePath}}v3/projects/{{project}}/notificationChannels")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NotificationChannel: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NotificationChannel: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), IsMonitoringConcurrentEditError)
	if err != nil {
		return fmt.Errorf("Error creating NotificationChannel: %s", err)
	}
	if err := d.Set("name", flattenMonitoringNotificationChannelName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	log.Printf("[DEBUG] Finished creating NotificationChannel %q: %#v", d.Id(), res)

	return resourceMonitoringNotificationChannelRead(d, meta)
}

func resourceMonitoringNotificationChannelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NotificationChannel: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil, IsMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringNotificationChannel %q", d.Id()))
	}

	res, err = resourceMonitoringNotificationChannelDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing MonitoringNotificationChannel because it no longer exists.")
		d.SetId("")
		return nil
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("force_delete"); !ok {
		if err := d.Set("force_delete", false); err != nil {
			return fmt.Errorf("Error setting force_delete: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}

	if err := d.Set("labels", flattenMonitoringNotificationChannelLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("name", flattenMonitoringNotificationChannelName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("verification_status", flattenMonitoringNotificationChannelVerificationStatus(res["verificationStatus"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("type", flattenMonitoringNotificationChannelType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("user_labels", flattenMonitoringNotificationChannelUserLabels(res["userLabels"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("description", flattenMonitoringNotificationChannelDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringNotificationChannelDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}
	if err := d.Set("enabled", flattenMonitoringNotificationChannelEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading NotificationChannel: %s", err)
	}

	return nil
}

func resourceMonitoringNotificationChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NotificationChannel: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandMonitoringNotificationChannelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	typeProp, err := expandMonitoringNotificationChannelType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	userLabelsProp, err := expandMonitoringNotificationChannelUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userLabelsProp)) {
		obj["userLabels"] = userLabelsProp
	}
	descriptionProp, err := expandMonitoringNotificationChannelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringNotificationChannelDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandMonitoringNotificationChannelEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); ok || !reflect.DeepEqual(v, enabledProp) {
		obj["enabled"] = enabledProp
	}

	obj, err = resourceMonitoringNotificationChannelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "stackdriver/notifications/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating NotificationChannel %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate), IsMonitoringConcurrentEditError)

	if err != nil {
		return fmt.Errorf("Error updating NotificationChannel %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating NotificationChannel %q: %#v", d.Id(), res)
	}

	return resourceMonitoringNotificationChannelRead(d, meta)
}

func resourceMonitoringNotificationChannelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NotificationChannel: %s", err)
	}
	billingProject = project

	lockName, err := ReplaceVars(d, config, "stackdriver/notifications/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}?force={{force_delete}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NotificationChannel %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), IsMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, "NotificationChannel")
	}

	log.Printf("[DEBUG] Finished deleting NotificationChannel %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringNotificationChannelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringNotificationChannelLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringNotificationChannelName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringNotificationChannelVerificationStatus(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringNotificationChannelType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringNotificationChannelUserLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringNotificationChannelDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringNotificationChannelDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringNotificationChannelEnabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandMonitoringNotificationChannelLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringNotificationChannelType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelUserLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringNotificationChannelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceMonitoringNotificationChannelEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	labelmap, ok := obj["labels"]
	if !ok {
		labelmap = make(map[string]string)
	}

	var labels map[string]string
	labels = labelmap.(map[string]string)

	for _, sl := range sensitiveLabels {
		if auth, _ := d.GetOkExists("sensitive_labels.0." + sl); auth != "" {
			labels[sl] = auth.(string)
		}
	}

	obj["labels"] = labels

	return obj, nil
}

func resourceMonitoringNotificationChannelDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if labelmap, ok := res["labels"]; ok {
		labels := labelmap.(map[string]interface{})
		for _, sl := range sensitiveLabels {
			if _, apiOk := labels[sl]; apiOk {
				if _, exists := d.GetOkExists("sensitive_labels.0." + sl); exists {
					delete(labels, sl)
				} else {
					labels[sl] = d.Get("labels." + sl)
				}
			}
		}
	}

	return res, nil
}
