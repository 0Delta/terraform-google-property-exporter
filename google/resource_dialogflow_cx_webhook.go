// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceDialogflowCXWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowCXWebhookCreate,
		Read:   resourceDialogflowCXWebhookRead,
		Update: resourceDialogflowCXWebhookUpdate,
		Delete: resourceDialogflowCXWebhookDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowCXWebhookImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The human-readable name of the webhook, unique within the agent.`,
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Indicates whether the webhook is disabled.`,
			},
			"enable_spell_correction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Indicates if automatic spell correction is enabled in detect intent requests.`,
			},
			"enable_stackdriver_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Determines whether this agent should log conversation queries.`,
			},
			"generic_web_service": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration for a generic web service.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Whether to use speech adaptation for speech recognition.`,
						},
						"allowed_ca_certs": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"request_headers": {
							Type:        schema.TypeMap,
							Optional:    true,
							ForceNew:    true,
							Description: `The HTTP request headers to send together with webhook requests.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The agent to create a webhook for.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.`,
			},
			"security_settings": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.`,
			},
			"service_directory": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration for a Service Directory service.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"generic_web_service": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The name of Service Directory service.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uri": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Whether to use speech adaptation for speech recognition.`,
									},
									"allowed_ca_certs": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"request_headers": {
										Type:        schema.TypeMap,
										Optional:    true,
										ForceNew:    true,
										Description: `The HTTP request headers to send together with webhook requests.`,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"service": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The name of Service Directory service.`,
						},
					},
				},
			},
			"timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Webhook execution timeout.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique identifier of the webhook.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.`,
			},
			"start_flow": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDialogflowCXWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXWebhookDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	timeoutProp, err := expandDialogflowCXWebhookTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeoutProp)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}
	disabledProp, err := expandDialogflowCXWebhookDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	genericWebServiceProp, err := expandDialogflowCXWebhookGenericWebService(d.Get("generic_web_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("generic_web_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(genericWebServiceProp)) && (ok || !reflect.DeepEqual(v, genericWebServiceProp)) {
		obj["genericWebService"] = genericWebServiceProp
	}
	serviceDirectoryProp, err := expandDialogflowCXWebhookServiceDirectory(d.Get("service_directory"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_directory"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceDirectoryProp)) && (ok || !reflect.DeepEqual(v, serviceDirectoryProp)) {
		obj["serviceDirectory"] = serviceDirectoryProp
	}
	securitySettingsProp, err := expandDialogflowCXWebhookSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("security_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(securitySettingsProp)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXWebhookEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXWebhookEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableSpellCorrectionProp)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/webhooks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Webhook: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Webhook: %s", err)
	}
	if err := d.Set("name", flattenDialogflowCXWebhookName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/webhooks/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Webhook %q: %#v", d.Id(), res)

	return resourceDialogflowCXWebhookRead(d, meta)
}

func resourceDialogflowCXWebhookRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/webhooks/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DialogflowCXWebhook %q", d.Id()))
	}

	if err := d.Set("name", flattenDialogflowCXWebhookName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowCXWebhookDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("timeout", flattenDialogflowCXWebhookTimeout(res["timeout"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("disabled", flattenDialogflowCXWebhookDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("generic_web_service", flattenDialogflowCXWebhookGenericWebService(res["genericWebService"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("service_directory", flattenDialogflowCXWebhookServiceDirectory(res["serviceDirectory"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("start_flow", flattenDialogflowCXWebhookStartFlow(res["startFlow"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("security_settings", flattenDialogflowCXWebhookSecuritySettings(res["securitySettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("enable_stackdriver_logging", flattenDialogflowCXWebhookEnableStackdriverLogging(res["enableStackdriverLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}
	if err := d.Set("enable_spell_correction", flattenDialogflowCXWebhookEnableSpellCorrection(res["enableSpellCorrection"], d, config)); err != nil {
		return fmt.Errorf("Error reading Webhook: %s", err)
	}

	return nil
}

func resourceDialogflowCXWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXWebhookDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	timeoutProp, err := expandDialogflowCXWebhookTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}
	disabledProp, err := expandDialogflowCXWebhookDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	genericWebServiceProp, err := expandDialogflowCXWebhookGenericWebService(d.Get("generic_web_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("generic_web_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, genericWebServiceProp)) {
		obj["genericWebService"] = genericWebServiceProp
	}
	serviceDirectoryProp, err := expandDialogflowCXWebhookServiceDirectory(d.Get("service_directory"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_directory"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serviceDirectoryProp)) {
		obj["serviceDirectory"] = serviceDirectoryProp
	}
	securitySettingsProp, err := expandDialogflowCXWebhookSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("security_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXWebhookEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXWebhookEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/webhooks/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Webhook %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("timeout") {
		updateMask = append(updateMask, "timeout")
	}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
	}

	if d.HasChange("generic_web_service") {
		updateMask = append(updateMask, "genericWebService")
	}

	if d.HasChange("service_directory") {
		updateMask = append(updateMask, "serviceDirectory")
	}

	if d.HasChange("security_settings") {
		updateMask = append(updateMask, "securitySettings")
	}

	if d.HasChange("enable_stackdriver_logging") {
		updateMask = append(updateMask, "enableStackdriverLogging")
	}

	if d.HasChange("enable_spell_correction") {
		updateMask = append(updateMask, "enableSpellCorrection")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Webhook %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Webhook %q: %#v", d.Id(), res)
	}

	return resourceDialogflowCXWebhookRead(d, meta)
}

func resourceDialogflowCXWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/webhooks/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	log.Printf("[DEBUG] Deleting Webhook %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Webhook")
	}

	log.Printf("[DEBUG] Finished deleting Webhook %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowCXWebhookImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value and parent contains slashes
	if err := tpgresource.ParseImportId([]string{
		"(?P<parent>.+)/webhooks/(?P<name>[^/]+)",
		"(?P<parent>.+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/webhooks/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowCXWebhookName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenDialogflowCXWebhookDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookTimeout(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookGenericWebService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["uri"] =
		flattenDialogflowCXWebhookGenericWebServiceUri(original["uri"], d, config)
	transformed["request_headers"] =
		flattenDialogflowCXWebhookGenericWebServiceRequestHeaders(original["requestHeaders"], d, config)
	transformed["allowed_ca_certs"] =
		flattenDialogflowCXWebhookGenericWebServiceAllowedCaCerts(original["allowedCaCerts"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXWebhookGenericWebServiceUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookGenericWebServiceRequestHeaders(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookGenericWebServiceAllowedCaCerts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookServiceDirectory(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service"] =
		flattenDialogflowCXWebhookServiceDirectoryService(original["service"], d, config)
	transformed["generic_web_service"] =
		flattenDialogflowCXWebhookServiceDirectoryGenericWebService(original["genericWebService"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXWebhookServiceDirectoryService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookServiceDirectoryGenericWebService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["uri"] =
		flattenDialogflowCXWebhookServiceDirectoryGenericWebServiceUri(original["uri"], d, config)
	transformed["request_headers"] =
		flattenDialogflowCXWebhookServiceDirectoryGenericWebServiceRequestHeaders(original["requestHeaders"], d, config)
	transformed["allowed_ca_certs"] =
		flattenDialogflowCXWebhookServiceDirectoryGenericWebServiceAllowedCaCerts(original["allowedCaCerts"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXWebhookServiceDirectoryGenericWebServiceUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookServiceDirectoryGenericWebServiceRequestHeaders(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookServiceDirectoryGenericWebServiceAllowedCaCerts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookStartFlow(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookSecuritySettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookEnableStackdriverLogging(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXWebhookEnableSpellCorrection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDialogflowCXWebhookDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookGenericWebService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandDialogflowCXWebhookGenericWebServiceUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedRequestHeaders, err := expandDialogflowCXWebhookGenericWebServiceRequestHeaders(original["request_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestHeaders"] = transformedRequestHeaders
	}

	transformedAllowedCaCerts, err := expandDialogflowCXWebhookGenericWebServiceAllowedCaCerts(original["allowed_ca_certs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedCaCerts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedCaCerts"] = transformedAllowedCaCerts
	}

	return transformed, nil
}

func expandDialogflowCXWebhookGenericWebServiceUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookGenericWebServiceRequestHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDialogflowCXWebhookGenericWebServiceAllowedCaCerts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookServiceDirectory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandDialogflowCXWebhookServiceDirectoryService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedGenericWebService, err := expandDialogflowCXWebhookServiceDirectoryGenericWebService(original["generic_web_service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGenericWebService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["genericWebService"] = transformedGenericWebService
	}

	return transformed, nil
}

func expandDialogflowCXWebhookServiceDirectoryService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandDialogflowCXWebhookServiceDirectoryGenericWebServiceUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedRequestHeaders, err := expandDialogflowCXWebhookServiceDirectoryGenericWebServiceRequestHeaders(original["request_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestHeaders"] = transformedRequestHeaders
	}

	transformedAllowedCaCerts, err := expandDialogflowCXWebhookServiceDirectoryGenericWebServiceAllowedCaCerts(original["allowed_ca_certs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedCaCerts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedCaCerts"] = transformedAllowedCaCerts
	}

	return transformed, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebServiceUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebServiceRequestHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebServiceAllowedCaCerts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookSecuritySettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookEnableStackdriverLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookEnableSpellCorrection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
