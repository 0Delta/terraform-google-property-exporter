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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// waitForAgentPoolReady waits for an agent pool to leave the
// "CREATING" state and become "CREATED", to indicate that it's ready.
func waitForAgentPoolReady(d *schema.ResourceData, config *Config, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		if err := resourceStorageTransferAgentPoolRead(d, config); err != nil {
			return resource.NonRetryableError(err)
		}

		name := d.Get("name").(string)
		state := d.Get("state").(string)
		if state == "CREATING" {
			return resource.RetryableError(fmt.Errorf("AgentPool %q has state %q.", name, state))
		} else if state == "CREATED" {
			log.Printf("[DEBUG] AgentPool %q has state %q.", name, state)
			return nil
		} else {
			return resource.NonRetryableError(fmt.Errorf("AgentPool %q has state %q.", name, state))
		}
	})
}

func ResourceStorageTransferAgentPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageTransferAgentPoolCreate,
		Read:   resourceStorageTransferAgentPoolRead,
		Update: resourceStorageTransferAgentPoolUpdate,
		Delete: resourceStorageTransferAgentPoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageTransferAgentPoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID of the agent pool to create.

The agentPoolId must meet the following requirements:
* Length of 128 characters or less.
* Not start with the string goog.
* Start with a lowercase ASCII character, followed by:
  * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
  * One or more numerals or lowercase ASCII characters.

As expressed by the regular expression: ^(?!goog)[a-z]([a-z0-9-._~]*[a-z0-9])?$.`,
			},
			"bandwidth_limit": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_mbps": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Bandwidth rate in megabytes per second, distributed across all the agents in the pool.`,
						},
					},
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the client-specified AgentPool description.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Specifies the state of the AgentPool.`,
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

func resourceStorageTransferAgentPoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandStorageTransferAgentPoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bandwidthLimitProp, err := expandStorageTransferAgentPoolBandwidthLimit(d.Get("bandwidth_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bandwidth_limit"); !isEmptyValue(reflect.ValueOf(bandwidthLimitProp)) && (ok || !reflect.DeepEqual(v, bandwidthLimitProp)) {
		obj["bandwidthLimit"] = bandwidthLimitProp
	}

	url, err := replaceVars(d, config, "{{StorageTransferBasePath}}projects/{{project}}/agentPools?agentPoolId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AgentPool: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AgentPool: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AgentPool: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/agentPools/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	if err := waitForAgentPoolReady(d, config, d.Timeout(schema.TimeoutCreate)-time.Minute); err != nil {
		return fmt.Errorf("Error waiting for AgentPool %q to be CREATED during creation: %q", d.Get("name").(string), err)
	}

	log.Printf("[DEBUG] Finished creating AgentPool %q: %#v", d.Id(), res)

	return resourceStorageTransferAgentPoolRead(d, meta)
}

func resourceStorageTransferAgentPoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{StorageTransferBasePath}}projects/{{project}}/agentPools/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AgentPool: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("StorageTransferAgentPool %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AgentPool: %s", err)
	}

	if err := d.Set("display_name", flattenStorageTransferAgentPoolDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AgentPool: %s", err)
	}
	if err := d.Set("state", flattenStorageTransferAgentPoolState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading AgentPool: %s", err)
	}
	if err := d.Set("bandwidth_limit", flattenStorageTransferAgentPoolBandwidthLimit(res["bandwidthLimit"], d, config)); err != nil {
		return fmt.Errorf("Error reading AgentPool: %s", err)
	}

	return nil
}

func resourceStorageTransferAgentPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AgentPool: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandStorageTransferAgentPoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bandwidthLimitProp, err := expandStorageTransferAgentPoolBandwidthLimit(d.Get("bandwidth_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bandwidth_limit"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bandwidthLimitProp)) {
		obj["bandwidthLimit"] = bandwidthLimitProp
	}

	url, err := replaceVars(d, config, "{{StorageTransferBasePath}}projects/{{project}}/agentPools/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AgentPool %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("bandwidth_limit") {
		updateMask = append(updateMask, "bandwidthLimit")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	if err := waitForAgentPoolReady(d, config, d.Timeout(schema.TimeoutCreate)-time.Minute); err != nil {
		return fmt.Errorf("Error waiting for AgentPool %q to be CREATED before updating: %q", d.Get("name").(string), err)
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AgentPool %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AgentPool %q: %#v", d.Id(), res)
	}

	return resourceStorageTransferAgentPoolRead(d, meta)
}

func resourceStorageTransferAgentPoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AgentPool: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{StorageTransferBasePath}}projects/{{project}}/agentPools/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AgentPool %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AgentPool")
	}

	log.Printf("[DEBUG] Finished deleting AgentPool %q: %#v", d.Id(), res)
	return nil
}

func resourceStorageTransferAgentPoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/agentPools/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/agentPools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	if err := waitForAgentPoolReady(d, config, d.Timeout(schema.TimeoutCreate)-time.Minute); err != nil {
		return nil, fmt.Errorf("Error waiting for AgentPool %q to be CREATED during importing: %q", d.Get("name").(string), err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenStorageTransferAgentPoolDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenStorageTransferAgentPoolState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenStorageTransferAgentPoolBandwidthLimit(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["limit_mbps"] =
		flattenStorageTransferAgentPoolBandwidthLimitLimitMbps(original["limitMbps"], d, config)
	return []interface{}{transformed}
}
func flattenStorageTransferAgentPoolBandwidthLimitLimitMbps(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandStorageTransferAgentPoolDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageTransferAgentPoolBandwidthLimit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLimitMbps, err := expandStorageTransferAgentPoolBandwidthLimitLimitMbps(original["limit_mbps"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLimitMbps); val.IsValid() && !isEmptyValue(val) {
		transformed["limitMbps"] = transformedLimitMbps
	}

	return transformed, nil
}

func expandStorageTransferAgentPoolBandwidthLimitLimitMbps(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
