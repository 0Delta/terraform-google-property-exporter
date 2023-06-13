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

package accessapproval

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

func ResourceAccessApprovalOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessApprovalOrganizationSettingsCreate,
		Read:   resourceAccessApprovalOrganizationSettingsRead,
		Update: resourceAccessApprovalOrganizationSettingsUpdate,
		Delete: resourceAccessApprovalOrganizationSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessApprovalOrganizationSettingsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enrolled_services": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `A list of Google Cloud Services for which the given resource has Access Approval enrolled.
Access requests for the resource given by name against any of these services contained here will be required
to have explicit approval. Enrollment can be done for individual services.

A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.`,
				Elem: accessapprovalOrganizationSettingsEnrolledServicesSchema(),
				Set:  accessApprovalEnrolledServicesHash,
			},
			"organization_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the organization of the access approval settings.`,
			},
			"active_key_version": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The asymmetric crypto key version to use for signing approval requests.
Empty active_key_version indicates that a Google-managed key should be used for signing.`,
			},
			"notification_emails": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Description: `A list of email addresses to which notifications relating to approval requests should be sent.
Notifications relating to a resource will be sent to all emails in the settings of ancestor
resources of that resource. A maximum of 50 email addresses are allowed.`,
				MaxItems: 50,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"ancestor_has_active_key_version": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `This field will always be unset for the organization since organizations do not have ancestors.`,
			},
			"enrolled_ancestor": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `This field will always be unset for the organization since organizations do not have ancestors.`,
			},
			"invalid_key_version": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `If the field is true, that indicates that there is some configuration issue with the active_key_version
configured on this Organization (e.g. it doesn't exist or the Access Approval service account doesn't have the
correct permissions on it, etc.).`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the settings. Format is "organizations/{organization_id}/accessApprovalSettings"`,
			},
		},
		UseJSONNumber: true,
	}
}

func accessapprovalOrganizationSettingsEnrolledServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_product": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
  all
  appengine.googleapis.com
  bigquery.googleapis.com
  bigtable.googleapis.com
  cloudkms.googleapis.com
  compute.googleapis.com
  dataflow.googleapis.com
  iam.googleapis.com
  pubsub.googleapis.com
  storage.googleapis.com`,
			},
			"enrollment_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"BLOCK_ALL", ""}),
				Description:  `The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"]`,
				Default:      "BLOCK_ALL",
			},
		},
	}
}

func resourceAccessApprovalOrganizationSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalOrganizationSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationEmailsProp)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalOrganizationSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !tpgresource.IsEmptyValue(reflect.ValueOf(enrolledServicesProp)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalOrganizationSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("active_key_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(activeKeyVersionProp)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessApprovalBasePath}}organizations/{{organization_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationSettings: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
	}

	if d.HasChange("active_key_version") {
		updateMask = append(updateMask, "activeKeyVersion")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating OrganizationSettings: %s", err)
	}
	if err := d.Set("name", flattenAccessApprovalOrganizationSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization_id}}/accessApprovalSettings")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating OrganizationSettings %q: %#v", d.Id(), res)

	return resourceAccessApprovalOrganizationSettingsRead(d, meta)
}

func resourceAccessApprovalOrganizationSettingsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessApprovalBasePath}}organizations/{{organization_id}}/accessApprovalSettings")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AccessApprovalOrganizationSettings %q", d.Id()))
	}

	if err := d.Set("name", flattenAccessApprovalOrganizationSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("notification_emails", flattenAccessApprovalOrganizationSettingsNotificationEmails(res["notificationEmails"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("enrolled_services", flattenAccessApprovalOrganizationSettingsEnrolledServices(res["enrolledServices"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("enrolled_ancestor", flattenAccessApprovalOrganizationSettingsEnrolledAncestor(res["enrolledAncestor"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("active_key_version", flattenAccessApprovalOrganizationSettingsActiveKeyVersion(res["activeKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("ancestor_has_active_key_version", flattenAccessApprovalOrganizationSettingsAncestorHasActiveKeyVersion(res["ancestorHasActiveKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("invalid_key_version", flattenAccessApprovalOrganizationSettingsInvalidKeyVersion(res["invalidKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}

	return nil
}

func resourceAccessApprovalOrganizationSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalOrganizationSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalOrganizationSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalOrganizationSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("active_key_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessApprovalBasePath}}organizations/{{organization_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
	}

	if d.HasChange("active_key_version") {
		updateMask = append(updateMask, "activeKeyVersion")
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
		return fmt.Errorf("Error updating OrganizationSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating OrganizationSettings %q: %#v", d.Id(), res)
	}

	return resourceAccessApprovalOrganizationSettingsRead(d, meta)
}

func resourceAccessApprovalOrganizationSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["notificationEmails"] = []string{}
	obj["enrolledServices"] = []string{}
	obj["activeKeyVersion"] = ""

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessApprovalBasePath}}organizations/{{organization_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Emptying OrganizationSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	updateMask = append(updateMask, "notificationEmails")
	updateMask = append(updateMask, "enrolledServices")
	updateMask = append(updateMask, "activeKeyVersion")

	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error emptying OrganizationSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished emptying OrganizationSettings %q: %#v", d.Id(), res)
	}

	return nil
}

func resourceAccessApprovalOrganizationSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"organizations/(?P<organization_id>[^/]+)/accessApprovalSettings",
		"(?P<organization_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization_id}}/accessApprovalSettings")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAccessApprovalOrganizationSettingsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalOrganizationSettingsNotificationEmails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenAccessApprovalOrganizationSettingsEnrolledServices(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(accessApprovalEnrolledServicesHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"cloud_product":    flattenAccessApprovalOrganizationSettingsEnrolledServicesCloudProduct(original["cloudProduct"], d, config),
			"enrollment_level": flattenAccessApprovalOrganizationSettingsEnrolledServicesEnrollmentLevel(original["enrollmentLevel"], d, config),
		})
	}
	return transformed
}
func flattenAccessApprovalOrganizationSettingsEnrolledServicesCloudProduct(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalOrganizationSettingsEnrolledServicesEnrollmentLevel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalOrganizationSettingsEnrolledAncestor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalOrganizationSettingsActiveKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalOrganizationSettingsAncestorHasActiveKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessApprovalOrganizationSettingsInvalidKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAccessApprovalOrganizationSettingsNotificationEmails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessApprovalOrganizationSettingsEnrolledServices(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCloudProduct, err := expandAccessApprovalOrganizationSettingsEnrolledServicesCloudProduct(original["cloud_product"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCloudProduct); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["cloudProduct"] = transformedCloudProduct
		}

		transformedEnrollmentLevel, err := expandAccessApprovalOrganizationSettingsEnrolledServicesEnrollmentLevel(original["enrollment_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnrollmentLevel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["enrollmentLevel"] = transformedEnrollmentLevel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessApprovalOrganizationSettingsEnrolledServicesCloudProduct(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalOrganizationSettingsEnrolledServicesEnrollmentLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalOrganizationSettingsActiveKeyVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
