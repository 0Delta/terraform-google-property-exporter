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

package apigee

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceApigeeOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeOrganizationCreate,
		Read:   resourceApigeeOrganizationRead,
		Update: resourceApigeeOrganizationUpdate,
		Delete: resourceApigeeOrganizationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeOrganizationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
			Update: schema.DefaultTimeout(45 * time.Minute),
			Delete: schema.DefaultTimeout(45 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The project ID associated with the Apigee organization.`,
			},
			"analytics_region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).`,
			},
			"authorized_network": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
Valid only when 'RuntimeType' is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".`,
			},
			"billing_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description of the Apigee organization.`,
			},
			"disable_vpc_peering": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Flag that specifies whether the VPC Peering through Private Google Access should be
disabled between the consumer network and Apigee. Required if an 'authorizedNetwork'
on the consumer project is not provided, in which case the flag should be set to 'true'.
Valid only when 'RuntimeType' is set to CLOUD. The value must be set before the creation
of any Apigee runtime instance and can be updated only when there are no runtime instances.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The display name of the Apigee organization.`,
			},
			"properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Properties defined in the Apigee organization profile.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"property": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `List of all properties in the object.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Name of the property.`,
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Value of the property.`,
									},
								},
							},
						},
					},
				},
			},
			"retention": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DELETION_RETENTION_UNSPECIFIED", "MINIMUM", ""}),
				Description: `Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
is not EVALUATION). It controls how long Organization data will be retained after the initial delete
operation completes. During this period, the Organization may be restored to its last known state.
After this period, the Organization will no longer be able to be restored. Default value: "DELETION_RETENTION_UNSPECIFIED" Possible values: ["DELETION_RETENTION_UNSPECIFIED", "MINIMUM"]`,
				Default: "DELETION_RETENTION_UNSPECIFIED",
			},
			"runtime_database_encryption_key_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
Update is not allowed after the organization is created.
If not specified, a Google-Managed encryption key will be used.
Valid only when 'RuntimeType' is CLOUD. For example: 'projects/foo/locations/us/keyRings/bar/cryptoKeys/baz'.`,
			},
			"runtime_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"CLOUD", "HYBRID", ""}),
				Description:  `Runtime type of the Apigee organization based on the Apigee subscription purchased. Default value: "CLOUD" Possible values: ["CLOUD", "HYBRID"]`,
				Default:      "CLOUD",
			},
			"apigee_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Project ID of the Apigee Tenant Project.`,
			},
			"ca_certificate": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Base64-encoded public certificate for the root CA of the Apigee organization.
Valid only when 'RuntimeType' is CLOUD. A base64-encoded string.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Name of the Apigee organization.`,
			},
			"subscription_type": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Subscription type of the Apigee organization.
Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased).`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandApigeeOrganizationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApigeeOrganizationDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	analyticsRegionProp, err := expandApigeeOrganizationAnalyticsRegion(d.Get("analytics_region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("analytics_region"); !tpgresource.IsEmptyValue(reflect.ValueOf(analyticsRegionProp)) && (ok || !reflect.DeepEqual(v, analyticsRegionProp)) {
		obj["analyticsRegion"] = analyticsRegionProp
	}
	authorizedNetworkProp, err := expandApigeeOrganizationAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	disableVpcPeeringProp, err := expandApigeeOrganizationDisableVpcPeering(d.Get("disable_vpc_peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_vpc_peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableVpcPeeringProp)) && (ok || !reflect.DeepEqual(v, disableVpcPeeringProp)) {
		obj["disableVpcPeering"] = disableVpcPeeringProp
	}
	runtimeTypeProp, err := expandApigeeOrganizationRuntimeType(d.Get("runtime_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeTypeProp)) && (ok || !reflect.DeepEqual(v, runtimeTypeProp)) {
		obj["runtimeType"] = runtimeTypeProp
	}
	billingTypeProp, err := expandApigeeOrganizationBillingType(d.Get("billing_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("billing_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(billingTypeProp)) && (ok || !reflect.DeepEqual(v, billingTypeProp)) {
		obj["billingType"] = billingTypeProp
	}
	runtimeDatabaseEncryptionKeyNameProp, err := expandApigeeOrganizationRuntimeDatabaseEncryptionKeyName(d.Get("runtime_database_encryption_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_database_encryption_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeDatabaseEncryptionKeyNameProp)) && (ok || !reflect.DeepEqual(v, runtimeDatabaseEncryptionKeyNameProp)) {
		obj["runtimeDatabaseEncryptionKeyName"] = runtimeDatabaseEncryptionKeyNameProp
	}
	propertiesProp, err := expandApigeeOrganizationProperties(d.Get("properties"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}

	obj, err = resourceApigeeOrganizationEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations?parent=projects/{{project_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Organization: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating Organization: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApigeeOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating Organization", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Organization: %s", err)
	}

	if err := d.Set("name", flattenApigeeOrganizationName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "organizations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Organization %q: %#v", d.Id(), res)

	return resourceApigeeOrganizationRead(d, meta)
}

func resourceApigeeOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeOrganization %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeOrganizationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("display_name", flattenApigeeOrganizationDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("description", flattenApigeeOrganizationDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("analytics_region", flattenApigeeOrganizationAnalyticsRegion(res["analyticsRegion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("authorized_network", flattenApigeeOrganizationAuthorizedNetwork(res["authorizedNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("disable_vpc_peering", flattenApigeeOrganizationDisableVpcPeering(res["disableVpcPeering"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("runtime_type", flattenApigeeOrganizationRuntimeType(res["runtimeType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("subscription_type", flattenApigeeOrganizationSubscriptionType(res["subscriptionType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("billing_type", flattenApigeeOrganizationBillingType(res["billingType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("ca_certificate", flattenApigeeOrganizationCaCertificate(res["caCertificate"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("runtime_database_encryption_key_name", flattenApigeeOrganizationRuntimeDatabaseEncryptionKeyName(res["runtimeDatabaseEncryptionKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("properties", flattenApigeeOrganizationProperties(res["properties"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}
	if err := d.Set("apigee_project_id", flattenApigeeOrganizationApigeeProjectId(res["apigeeProjectId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Organization: %s", err)
	}

	return nil
}

func resourceApigeeOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandApigeeOrganizationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApigeeOrganizationDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	analyticsRegionProp, err := expandApigeeOrganizationAnalyticsRegion(d.Get("analytics_region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("analytics_region"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, analyticsRegionProp)) {
		obj["analyticsRegion"] = analyticsRegionProp
	}
	authorizedNetworkProp, err := expandApigeeOrganizationAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	disableVpcPeeringProp, err := expandApigeeOrganizationDisableVpcPeering(d.Get("disable_vpc_peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_vpc_peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disableVpcPeeringProp)) {
		obj["disableVpcPeering"] = disableVpcPeeringProp
	}
	runtimeTypeProp, err := expandApigeeOrganizationRuntimeType(d.Get("runtime_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, runtimeTypeProp)) {
		obj["runtimeType"] = runtimeTypeProp
	}
	billingTypeProp, err := expandApigeeOrganizationBillingType(d.Get("billing_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("billing_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, billingTypeProp)) {
		obj["billingType"] = billingTypeProp
	}
	runtimeDatabaseEncryptionKeyNameProp, err := expandApigeeOrganizationRuntimeDatabaseEncryptionKeyName(d.Get("runtime_database_encryption_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_database_encryption_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, runtimeDatabaseEncryptionKeyNameProp)) {
		obj["runtimeDatabaseEncryptionKeyName"] = runtimeDatabaseEncryptionKeyNameProp
	}
	propertiesProp, err := expandApigeeOrganizationProperties(d.Get("properties"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}

	obj, err = resourceApigeeOrganizationEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Organization %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Organization %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Organization %q: %#v", d.Id(), res)
	}

	err = ApigeeOperationWaitTime(
		config, res, "Updating Organization", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceApigeeOrganizationRead(d, meta)
}

func resourceApigeeOrganizationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{name}}?retention={{retention}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Organization %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Organization")
	}

	log.Printf("[DEBUG] Finished deleting Organization %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeOrganizationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	parts := strings.Split(d.Get("name").(string), "/")

	var projectId string
	switch len(parts) {
	case 1:
		projectId = parts[0]
	case 2:
		projectId = parts[1]
	default:
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s or %s",
			d.Get("name"),
			"{{name}}",
			"organizations/{{name}}",
		)
	}

	if err := d.Set("name", projectId); err != nil {
		return nil, fmt.Errorf("Error setting organization: %s", err)
	}

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("Error setting organization: %s", err)
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeOrganizationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationAnalyticsRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationAuthorizedNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationDisableVpcPeering(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationRuntimeType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationSubscriptionType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationBillingType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationCaCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationRuntimeDatabaseEncryptionKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["property"] =
		flattenApigeeOrganizationPropertiesProperty(original["property"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeOrganizationPropertiesProperty(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"name":  flattenApigeeOrganizationPropertiesPropertyName(original["name"], d, config),
			"value": flattenApigeeOrganizationPropertiesPropertyValue(original["value"], d, config),
		})
	}
	return transformed
}
func flattenApigeeOrganizationPropertiesPropertyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationPropertiesPropertyValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeOrganizationApigeeProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeOrganizationDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationAnalyticsRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationAuthorizedNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationDisableVpcPeering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationRuntimeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationBillingType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationRuntimeDatabaseEncryptionKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProperty, err := expandApigeeOrganizationPropertiesProperty(original["property"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperty); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["property"] = transformedProperty
	}

	return transformed, nil
}

func expandApigeeOrganizationPropertiesProperty(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandApigeeOrganizationPropertiesPropertyName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandApigeeOrganizationPropertiesPropertyValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApigeeOrganizationPropertiesPropertyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationPropertiesPropertyValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceApigeeOrganizationEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["name"] = d.Get("project_id").(string)
	return obj, nil
}
