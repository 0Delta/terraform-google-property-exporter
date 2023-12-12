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

package bigqueryreservation

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

func ResourceBigqueryReservationBiReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationBiReservationCreate,
		Read:   resourceBigqueryReservationBiReservationRead,
		Update: resourceBigqueryReservationBiReservationUpdate,
		Delete: resourceBigqueryReservationBiReservationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationBiReservationImport,
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
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `LOCATION_DESCRIPTION`,
			},
			"preferred_tables": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Preferred tables to use BI capacity for.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The ID of the dataset in the above project.`,
						},
						"project_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The assigned project ID of the project.`,
						},
						"table_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The ID of the table in the above dataset.`,
						},
					},
				},
			},
			"size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `Size of a reservation, in bytes.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the singleton BI reservation. Reservation names have the form 'projects/{projectId}/locations/{locationId}/biReservation'.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The last update timestamp of a reservation.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
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

func resourceBigqueryReservationBiReservationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	sizeProp, err := expandBigqueryReservationBiReservationSize(d.Get("size"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("size"); !tpgresource.IsEmptyValue(reflect.ValueOf(sizeProp)) && (ok || !reflect.DeepEqual(v, sizeProp)) {
		obj["size"] = sizeProp
	}
	preferredTablesProp, err := expandBigqueryReservationBiReservationPreferredTables(d.Get("preferred_tables"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preferred_tables"); !tpgresource.IsEmptyValue(reflect.ValueOf(preferredTablesProp)) && (ok || !reflect.DeepEqual(v, preferredTablesProp)) {
		obj["preferredTables"] = preferredTablesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/biReservation")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BiReservation: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BiReservation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	updateMask := []string{}

	if d.HasChange("size") {
		updateMask = append(updateMask, "size")
	}

	if d.HasChange("preferred_tables") {
		updateMask = append(updateMask, "preferredTables")
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
		return fmt.Errorf("Error creating BiReservation: %s", err)
	}
	if err := d.Set("name", flattenBigqueryReservationBiReservationName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/biReservation")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BiReservation %q: %#v", d.Id(), res)

	return resourceBigqueryReservationBiReservationRead(d, meta)
}

func resourceBigqueryReservationBiReservationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/biReservation")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BiReservation: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BigqueryReservationBiReservation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BiReservation: %s", err)
	}

	if err := d.Set("name", flattenBigqueryReservationBiReservationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BiReservation: %s", err)
	}
	if err := d.Set("update_time", flattenBigqueryReservationBiReservationUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BiReservation: %s", err)
	}
	if err := d.Set("size", flattenBigqueryReservationBiReservationSize(res["size"], d, config)); err != nil {
		return fmt.Errorf("Error reading BiReservation: %s", err)
	}
	if err := d.Set("preferred_tables", flattenBigqueryReservationBiReservationPreferredTables(res["preferredTables"], d, config)); err != nil {
		return fmt.Errorf("Error reading BiReservation: %s", err)
	}

	return nil
}

func resourceBigqueryReservationBiReservationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BiReservation: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	sizeProp, err := expandBigqueryReservationBiReservationSize(d.Get("size"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("size"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sizeProp)) {
		obj["size"] = sizeProp
	}
	preferredTablesProp, err := expandBigqueryReservationBiReservationPreferredTables(d.Get("preferred_tables"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preferred_tables"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, preferredTablesProp)) {
		obj["preferredTables"] = preferredTablesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/biReservation")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BiReservation %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("size") {
		updateMask = append(updateMask, "size")
	}

	if d.HasChange("preferred_tables") {
		updateMask = append(updateMask, "preferredTables")
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
			return fmt.Errorf("Error updating BiReservation %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating BiReservation %q: %#v", d.Id(), res)
		}

	}

	return resourceBigqueryReservationBiReservationRead(d, meta)
}

func resourceBigqueryReservationBiReservationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["preferredTables"] = []string{}
	obj["size"] = 0

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/biReservation")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Clearing BIReservation %q: %#v", d.Id(), obj)
	updateMask := []string{}

	updateMask = append(updateMask, "size")
	updateMask = append(updateMask, "preferredTables")

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
		return fmt.Errorf("Error clearing BIReservation %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished clearing BIReservation %q: %#v", d.Id(), res)
	}

	return nil
}

func resourceBigqueryReservationBiReservationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/biReservation$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)$",
		"^(?P<location>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/biReservation")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryReservationBiReservationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationBiReservationUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationBiReservationSize(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBigqueryReservationBiReservationPreferredTables(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"project_id": flattenBigqueryReservationBiReservationPreferredTablesProjectId(original["projectId"], d, config),
			"dataset_id": flattenBigqueryReservationBiReservationPreferredTablesDatasetId(original["datasetId"], d, config),
			"table_id":   flattenBigqueryReservationBiReservationPreferredTablesTableId(original["tableId"], d, config),
		})
	}
	return transformed
}
func flattenBigqueryReservationBiReservationPreferredTablesProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationBiReservationPreferredTablesDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationBiReservationPreferredTablesTableId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBigqueryReservationBiReservationSize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationBiReservationPreferredTables(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProjectId, err := expandBigqueryReservationBiReservationPreferredTablesProjectId(original["project_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["projectId"] = transformedProjectId
		}

		transformedDatasetId, err := expandBigqueryReservationBiReservationPreferredTablesDatasetId(original["dataset_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["datasetId"] = transformedDatasetId
		}

		transformedTableId, err := expandBigqueryReservationBiReservationPreferredTablesTableId(original["table_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["tableId"] = transformedTableId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBigqueryReservationBiReservationPreferredTablesProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationBiReservationPreferredTablesDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationBiReservationPreferredTablesTableId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
