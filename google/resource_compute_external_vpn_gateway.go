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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceComputeExternalVpnGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeExternalVpnGatewayCreate,
		Read:   resourceComputeExternalVpnGatewayRead,
		Delete: resourceComputeExternalVpnGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeExternalVpnGatewayImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"interface": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `A list of interfaces on this external VPN gateway.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Description: `The numeric ID for this interface. Allowed values are based on the redundancy type
of this external VPN gateway
* '0 - SINGLE_IP_INTERNALLY_REDUNDANT'
* '0, 1 - TWO_IPS_REDUNDANCY'
* '0, 1, 2, 3 - FOUR_IPS_REDUNDANCY'`,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `IP address of the interface in the external VPN gateway.
Only IPv4 is supported. This IP address can be either from
your on-premise gateway or another Cloud provider's VPN gateway,
it cannot be an IP address from Google Compute Engine.`,
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Labels for the external VPN gateway resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"redundancy_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"FOUR_IPS_REDUNDANCY", "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY", ""}),
				Description:  `Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY", "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"]`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeExternalVpnGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeExternalVpnGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandComputeExternalVpnGatewayLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	nameProp, err := expandComputeExternalVpnGatewayName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	redundancyTypeProp, err := expandComputeExternalVpnGatewayRedundancyType(d.Get("redundancy_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redundancy_type"); !isEmptyValue(reflect.ValueOf(redundancyTypeProp)) && (ok || !reflect.DeepEqual(v, redundancyTypeProp)) {
		obj["redundancyType"] = redundancyTypeProp
	}
	interfacesProp, err := expandComputeExternalVpnGatewayInterface(d.Get("interface"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interface"); !isEmptyValue(reflect.ValueOf(interfacesProp)) && (ok || !reflect.DeepEqual(v, interfacesProp)) {
		obj["interfaces"] = interfacesProp
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/externalVpnGateways")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ExternalVpnGateway: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ExternalVpnGateway: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ExternalVpnGateway: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/externalVpnGateways/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating ExternalVpnGateway", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ExternalVpnGateway: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ExternalVpnGateway %q: %#v", d.Id(), res)

	return resourceComputeExternalVpnGatewayRead(d, meta)
}

func resourceComputeExternalVpnGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/externalVpnGateways/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ExternalVpnGateway: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeExternalVpnGateway %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}

	if err := d.Set("description", flattenComputeExternalVpnGatewayDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("labels", flattenComputeExternalVpnGatewayLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("name", flattenComputeExternalVpnGatewayName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("redundancy_type", flattenComputeExternalVpnGatewayRedundancyType(res["redundancyType"], d, config)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("interface", flattenComputeExternalVpnGatewayInterface(res["interfaces"], d, config)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}

	return nil
}

func resourceComputeExternalVpnGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ExternalVpnGateway: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/externalVpnGateways/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ExternalVpnGateway %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ExternalVpnGateway")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting ExternalVpnGateway", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ExternalVpnGateway %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeExternalVpnGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/global/externalVpnGateways/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/externalVpnGateways/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeExternalVpnGatewayDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeExternalVpnGatewayLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeExternalVpnGatewayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeExternalVpnGatewayRedundancyType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeExternalVpnGatewayInterface(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"id":         flattenComputeExternalVpnGatewayInterfaceId(original["id"], d, config),
			"ip_address": flattenComputeExternalVpnGatewayInterfaceIpAddress(original["ipAddress"], d, config),
		})
	}
	return transformed
}
func flattenComputeExternalVpnGatewayInterfaceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeExternalVpnGatewayInterfaceIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeExternalVpnGatewayDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeExternalVpnGatewayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayRedundancyType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandComputeExternalVpnGatewayInterfaceId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["id"] = transformedId
		}

		transformedIpAddress, err := expandComputeExternalVpnGatewayInterfaceIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeExternalVpnGatewayInterfaceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayInterfaceIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
