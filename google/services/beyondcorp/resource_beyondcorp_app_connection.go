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

package beyondcorp

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceBeyondcorpAppConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceBeyondcorpAppConnectionCreate,
		Read:   resourceBeyondcorpAppConnectionRead,
		Update: resourceBeyondcorpAppConnectionUpdate,
		Delete: resourceBeyondcorpAppConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBeyondcorpAppConnectionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"application_endpoint": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Address of the remote application endpoint for the BeyondCorp AppConnection.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Hostname or IP address of the remote application endpoint.`,
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `Port of the remote application endpoint.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the AppConnection.`,
			},
			"connectors": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `List of AppConnectors that are authorised to be associated with this AppConnection`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An arbitrary user-provided name for the AppConnection.`,
			},
			"gateway": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Gateway used by the AppConnection.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_gateway": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.`,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The type of hosting used by the gateway. Refer to
https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
for a list of possible values.`,
						},
						"ingress_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Ingress port reserved on the gateways for this AppConnection, if not specified or zero, the default port is 19443.`,
						},
						"uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Server-defined URI for this resource.`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Resource labels to represent user provided metadata.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the AppConnection.`,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The type of network connectivity used by the AppConnection. Refer to
https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
for a list of possible values.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceBeyondcorpAppConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandBeyondcorpAppConnectionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeProp, err := expandBeyondcorpAppConnectionType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	applicationEndpointProp, err := expandBeyondcorpAppConnectionApplicationEndpoint(d.Get("application_endpoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("application_endpoint"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationEndpointProp)) && (ok || !reflect.DeepEqual(v, applicationEndpointProp)) {
		obj["applicationEndpoint"] = applicationEndpointProp
	}
	connectorsProp, err := expandBeyondcorpAppConnectionConnectors(d.Get("connectors"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connectors"); !tpgresource.IsEmptyValue(reflect.ValueOf(connectorsProp)) && (ok || !reflect.DeepEqual(v, connectorsProp)) {
		obj["connectors"] = connectorsProp
	}
	gatewayProp, err := expandBeyondcorpAppConnectionGateway(d.Get("gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway"); !tpgresource.IsEmptyValue(reflect.ValueOf(gatewayProp)) && (ok || !reflect.DeepEqual(v, gatewayProp)) {
		obj["gateway"] = gatewayProp
	}
	labelsProp, err := expandBeyondcorpAppConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnections?app_connection_id={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AppConnection: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnection: %s", err)
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
		return fmt.Errorf("Error creating AppConnection: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = BeyondcorpOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating AppConnection", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AppConnection: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AppConnection %q: %#v", d.Id(), res)

	return resourceBeyondcorpAppConnectionRead(d, meta)
}

func resourceBeyondcorpAppConnectionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnection: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BeyondcorpAppConnection %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}

	if err := d.Set("display_name", flattenBeyondcorpAppConnectionDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}
	if err := d.Set("labels", flattenBeyondcorpAppConnectionLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}
	if err := d.Set("type", flattenBeyondcorpAppConnectionType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}
	if err := d.Set("application_endpoint", flattenBeyondcorpAppConnectionApplicationEndpoint(res["applicationEndpoint"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}
	if err := d.Set("connectors", flattenBeyondcorpAppConnectionConnectors(res["connectors"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}
	if err := d.Set("gateway", flattenBeyondcorpAppConnectionGateway(res["gateway"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}
	if err := d.Set("terraform_labels", flattenBeyondcorpAppConnectionTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}
	if err := d.Set("effective_labels", flattenBeyondcorpAppConnectionEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppConnection: %s", err)
	}

	return nil
}

func resourceBeyondcorpAppConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnection: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandBeyondcorpAppConnectionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	applicationEndpointProp, err := expandBeyondcorpAppConnectionApplicationEndpoint(d.Get("application_endpoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("application_endpoint"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, applicationEndpointProp)) {
		obj["applicationEndpoint"] = applicationEndpointProp
	}
	connectorsProp, err := expandBeyondcorpAppConnectionConnectors(d.Get("connectors"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connectors"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, connectorsProp)) {
		obj["connectors"] = connectorsProp
	}
	gatewayProp, err := expandBeyondcorpAppConnectionGateway(d.Get("gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gatewayProp)) {
		obj["gateway"] = gatewayProp
	}
	labelsProp, err := expandBeyondcorpAppConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AppConnection %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("application_endpoint") {
		updateMask = append(updateMask, "applicationEndpoint")
	}

	if d.HasChange("connectors") {
		updateMask = append(updateMask, "connectors")
	}

	if d.HasChange("gateway") {
		updateMask = append(updateMask, "gateway")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
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
			return fmt.Errorf("Error updating AppConnection %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating AppConnection %q: %#v", d.Id(), res)
		}

		err = BeyondcorpOperationWaitTime(
			config, res, project, "Updating AppConnection", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceBeyondcorpAppConnectionRead(d, meta)
}

func resourceBeyondcorpAppConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppConnection: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AppConnection %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "AppConnection")
	}

	err = BeyondcorpOperationWaitTime(
		config, res, project, "Deleting AppConnection", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AppConnection %q: %#v", d.Id(), res)
	return nil
}

func resourceBeyondcorpAppConnectionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/appConnections/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBeyondcorpAppConnectionDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectionLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenBeyondcorpAppConnectionType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectionApplicationEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] =
		flattenBeyondcorpAppConnectionApplicationEndpointHost(original["host"], d, config)
	transformed["port"] =
		flattenBeyondcorpAppConnectionApplicationEndpointPort(original["port"], d, config)
	return []interface{}{transformed}
}
func flattenBeyondcorpAppConnectionApplicationEndpointHost(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectionApplicationEndpointPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBeyondcorpAppConnectionConnectors(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectionGateway(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["app_gateway"] =
		flattenBeyondcorpAppConnectionGatewayAppGateway(original["appGateway"], d, config)
	transformed["type"] =
		flattenBeyondcorpAppConnectionGatewayType(original["type"], d, config)
	transformed["uri"] =
		flattenBeyondcorpAppConnectionGatewayUri(original["uri"], d, config)
	transformed["ingress_port"] =
		flattenBeyondcorpAppConnectionGatewayIngressPort(original["ingressPort"], d, config)
	return []interface{}{transformed}
}
func flattenBeyondcorpAppConnectionGatewayAppGateway(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectionGatewayType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectionGatewayUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBeyondcorpAppConnectionGatewayIngressPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBeyondcorpAppConnectionTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenBeyondcorpAppConnectionEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBeyondcorpAppConnectionDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionApplicationEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandBeyondcorpAppConnectionApplicationEndpointHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedPort, err := expandBeyondcorpAppConnectionApplicationEndpointPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	return transformed, nil
}

func expandBeyondcorpAppConnectionApplicationEndpointHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionApplicationEndpointPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionConnectors(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionGateway(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAppGateway, err := expandBeyondcorpAppConnectionGatewayAppGateway(original["app_gateway"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAppGateway); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["appGateway"] = transformedAppGateway
	}

	transformedType, err := expandBeyondcorpAppConnectionGatewayType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedUri, err := expandBeyondcorpAppConnectionGatewayUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedIngressPort, err := expandBeyondcorpAppConnectionGatewayIngressPort(original["ingress_port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngressPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingressPort"] = transformedIngressPort
	}

	return transformed, nil
}

func expandBeyondcorpAppConnectionGatewayAppGateway(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionGatewayType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionGatewayUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionGatewayIngressPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
