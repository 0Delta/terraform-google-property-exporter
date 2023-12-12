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

package compute

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"

	"google.golang.org/api/googleapi"
)

func ResourceComputeNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkCreate,
		Read:   resourceComputeNetworkRead,
		Update: resourceComputeNetworkUpdate,
		Delete: resourceComputeNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"auto_create_subnetworks": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `When set to 'true', the network is created in "auto subnet mode" and
it will create a subnet for each region automatically across the
'10.128.0.0/9' address range.

When set to 'false', the network is created in "custom subnet mode" so
the user can explicitly connect subnetwork resources.`,
				Default: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. The resource must be
recreated to modify this field.`,
			},
			"enable_ula_internal_ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Enable ULA internal ipv6 on this network. Enabling this feature will assign
a /48 from google defined ULA prefix fd20::/20.`,
			},
			"internal_ipv6_range": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `When enabling ula internal ipv6, caller optionally can specify the /48 range
they want from the google defined ULA prefix fd20::/20. The input must be a
valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will
fail if the speficied /48 is already in used by another resource.
If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field.`,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Maximum Transmission Unit in bytes. The default value is 1460 bytes.
The minimum value for this field is 1300 and the maximum value is 8896 bytes (jumbo frames).
Note that packets larger than 1500 bytes (standard Ethernet) can be subject to TCP-MSS clamping or dropped
with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or other VPCs
with varying MTUs.`,
			},
			"network_firewall_policy_enforcement_order": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL", ""}),
				Description:  `Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"]`,
				Default:      "AFTER_CLASSIC_FIREWALL",
			},
			"routing_mode": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"REGIONAL", "GLOBAL", ""}),
				Description: `The network-wide routing mode to use. If set to 'REGIONAL', this
network's cloud routers will only advertise routes with subnetworks
of this network in the same region as the router. If set to 'GLOBAL',
this network's cloud routers will advertise routes with all
subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]`,
			},

			"gateway_ipv4": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The gateway address for default routing out of the network. This value
is selected by GCP.`,
			},
			"delete_default_routes_on_create": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `If set to 'true', default routes ('0.0.0.0/0') will be deleted
immediately after network creation. Defaults to 'false'.`,
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

func resourceComputeNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeNetworkName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	autoCreateSubnetworksProp, err := expandComputeNetworkAutoCreateSubnetworks(d.Get("auto_create_subnetworks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auto_create_subnetworks"); ok || !reflect.DeepEqual(v, autoCreateSubnetworksProp) {
		obj["autoCreateSubnetworks"] = autoCreateSubnetworksProp
	}
	routingConfigProp, err := expandComputeNetworkRoutingConfig(nil, d, config)
	if err != nil {
		return err
	} else if !tpgresource.IsEmptyValue(reflect.ValueOf(routingConfigProp)) {
		obj["routingConfig"] = routingConfigProp
	}
	mtuProp, err := expandComputeNetworkMtu(d.Get("mtu"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mtu"); !tpgresource.IsEmptyValue(reflect.ValueOf(mtuProp)) && (ok || !reflect.DeepEqual(v, mtuProp)) {
		obj["mtu"] = mtuProp
	}
	enableUlaInternalIpv6Prop, err := expandComputeNetworkEnableUlaInternalIpv6(d.Get("enable_ula_internal_ipv6"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_ula_internal_ipv6"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableUlaInternalIpv6Prop)) && (ok || !reflect.DeepEqual(v, enableUlaInternalIpv6Prop)) {
		obj["enableUlaInternalIpv6"] = enableUlaInternalIpv6Prop
	}
	internalIpv6RangeProp, err := expandComputeNetworkInternalIpv6Range(d.Get("internal_ipv6_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("internal_ipv6_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(internalIpv6RangeProp)) && (ok || !reflect.DeepEqual(v, internalIpv6RangeProp)) {
		obj["internalIpv6Range"] = internalIpv6RangeProp
	}
	networkFirewallPolicyEnforcementOrderProp, err := expandComputeNetworkNetworkFirewallPolicyEnforcementOrder(d.Get("network_firewall_policy_enforcement_order"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_firewall_policy_enforcement_order"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkFirewallPolicyEnforcementOrderProp)) && (ok || !reflect.DeepEqual(v, networkFirewallPolicyEnforcementOrderProp)) {
		obj["networkFirewallPolicyEnforcementOrder"] = networkFirewallPolicyEnforcementOrderProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Network: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Network: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

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
		return fmt.Errorf("Error creating Network: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating Network", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Network: %s", err)
	}

	if d.Get("delete_default_routes_on_create").(bool) {
		token := ""
		for paginate := true; paginate; {
			network, err := config.NewComputeClient(userAgent).Networks.Get(project, d.Get("name").(string)).Do()
			if err != nil {
				return fmt.Errorf("Error finding network in proj: %s", err)
			}
			filter := fmt.Sprintf("(network=\"%s\") AND (destRange=\"0.0.0.0/0\")", network.SelfLink)
			log.Printf("[DEBUG] Getting routes for network %q with filter '%q'", d.Get("name").(string), filter)
			resp, err := config.NewComputeClient(userAgent).Routes.List(project).Filter(filter).Do()
			if err != nil {
				return fmt.Errorf("Error listing routes in proj: %s", err)
			}

			log.Printf("[DEBUG] Found %d routes rules in %q network", len(resp.Items), d.Get("name").(string))

			for _, route := range resp.Items {
				op, err := config.NewComputeClient(userAgent).Routes.Delete(project, route.Name).Do()
				if err != nil {
					return fmt.Errorf("Error deleting route: %s", err)
				}
				err = ComputeOperationWaitTime(config, op, project, "Deleting Route", userAgent, d.Timeout(schema.TimeoutCreate))
				if err != nil {
					return err
				}
			}

			token = resp.NextPageToken
			paginate = token != ""
		}
	}

	log.Printf("[DEBUG] Finished creating Network %q: %#v", d.Id(), res)

	return resourceComputeNetworkRead(d, meta)
}

func resourceComputeNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Network: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNetwork %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("delete_default_routes_on_create"); !ok {
		if err := d.Set("delete_default_routes_on_create", false); err != nil {
			return fmt.Errorf("Error setting delete_default_routes_on_create: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}

	if err := d.Set("description", flattenComputeNetworkDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("gateway_ipv4", flattenComputeNetworkGatewayIpv4(res["gatewayIPv4"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("name", flattenComputeNetworkName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("auto_create_subnetworks", flattenComputeNetworkAutoCreateSubnetworks(res["autoCreateSubnetworks"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenComputeNetworkRoutingConfig(res["routingConfig"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading Network: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				if err := d.Set(k, v); err != nil {
					return fmt.Errorf("Error setting %s: %s", k, err)
				}
			}
		}
	}
	if err := d.Set("mtu", flattenComputeNetworkMtu(res["mtu"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("enable_ula_internal_ipv6", flattenComputeNetworkEnableUlaInternalIpv6(res["enableUlaInternalIpv6"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("internal_ipv6_range", flattenComputeNetworkInternalIpv6Range(res["internalIpv6Range"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("network_firewall_policy_enforcement_order", flattenComputeNetworkNetworkFirewallPolicyEnforcementOrder(res["networkFirewallPolicyEnforcementOrder"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}

	return nil
}

func resourceComputeNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Network: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("routing_mode") || d.HasChange("network_firewall_policy_enforcement_order") {
		obj := make(map[string]interface{})

		routingConfigProp, err := expandComputeNetworkRoutingConfig(nil, d, config)
		if err != nil {
			return err
		} else if !tpgresource.IsEmptyValue(reflect.ValueOf(routingConfigProp)) {
			obj["routingConfig"] = routingConfigProp
		}
		networkFirewallPolicyEnforcementOrderProp, err := expandComputeNetworkNetworkFirewallPolicyEnforcementOrder(d.Get("network_firewall_policy_enforcement_order"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("network_firewall_policy_enforcement_order"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkFirewallPolicyEnforcementOrderProp)) {
			obj["networkFirewallPolicyEnforcementOrder"] = networkFirewallPolicyEnforcementOrderProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{name}}")
		if err != nil {
			return err
		}

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
			return fmt.Errorf("Error updating Network %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Network %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating Network", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeNetworkRead(d, meta)
}

func resourceComputeNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Network: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Network %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Network")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting Network", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Network %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/global/networks/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("delete_default_routes_on_create", false); err != nil {
		return nil, fmt.Errorf("Error setting delete_default_routes_on_create: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNetworkDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkGatewayIpv4(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkAutoCreateSubnetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkRoutingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["routing_mode"] =
		flattenComputeNetworkRoutingConfigRoutingMode(original["routingMode"], d, config)
	return []interface{}{transformed}
}
func flattenComputeNetworkRoutingConfigRoutingMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkMtu(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenComputeNetworkEnableUlaInternalIpv6(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkInternalIpv6Range(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNetworkNetworkFirewallPolicyEnforcementOrder(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeNetworkDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAutoCreateSubnetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkRoutingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedRoutingMode, err := expandComputeNetworkRoutingConfigRoutingMode(d.Get("routing_mode"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRoutingMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["routingMode"] = transformedRoutingMode
	}

	return transformed, nil
}

func expandComputeNetworkRoutingConfigRoutingMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkMtu(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEnableUlaInternalIpv6(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkInternalIpv6Range(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkNetworkFirewallPolicyEnforcementOrder(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
