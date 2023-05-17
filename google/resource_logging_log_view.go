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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceLoggingLogView() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoggingLogViewCreate,
		Read:   resourceLoggingLogViewRead,
		Update: resourceLoggingLogViewUpdate,
		Delete: resourceLoggingLogViewDelete,

		Importer: &schema.ResourceImporter{
			State: resourceLoggingLogViewImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The bucket of the resource`,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The resource name of the view. For example: \'projects/my-project/locations/global/buckets/my-bucket/views/my-view\'`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Describes this view.`,
			},
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")`,
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.`,
			},
			"parent": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The parent of the resource.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The creation timestamp of the view.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The last update timestamp of the view.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceLoggingLogViewCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandLoggingLogViewName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandLoggingLogViewDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandLoggingLogViewFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	obj, err = resourceLoggingLogViewEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/buckets/{{bucket}}/views?viewId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new LogView: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating LogView: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating LogView %q: %#v", d.Id(), res)

	return resourceLoggingLogViewRead(d, meta)
}

func resourceLoggingLogViewRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("LoggingLogView %q", d.Id()))
	}

	if err := d.Set("description", flattenLoggingLogViewDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogView: %s", err)
	}
	if err := d.Set("create_time", flattenLoggingLogViewCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogView: %s", err)
	}
	if err := d.Set("update_time", flattenLoggingLogViewUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogView: %s", err)
	}
	if err := d.Set("filter", flattenLoggingLogViewFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogView: %s", err)
	}

	return nil
}

func resourceLoggingLogViewUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandLoggingLogViewDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandLoggingLogViewFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	obj, err = resourceLoggingLogViewEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating LogView %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("filter") {
		updateMask = append(updateMask, "filter")
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

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating LogView %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating LogView %q: %#v", d.Id(), res)
	}

	return resourceLoggingLogViewRead(d, meta)
}

func resourceLoggingLogViewDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting LogView %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "LogView")
	}

	log.Printf("[DEBUG] Finished deleting LogView %q: %#v", d.Id(), res)
	return nil
}

func resourceLoggingLogViewImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"(?P<parent>.+)/locations/(?P<location>[^/]+)/buckets/(?P<bucket>[^/]+)/views/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenLoggingLogViewDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingLogViewCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingLogViewUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingLogViewFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandLoggingLogViewName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingLogViewDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingLogViewFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceLoggingLogViewEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Extract any empty fields from the bucket field.
	parent := d.Get("parent").(string)
	bucket := d.Get("bucket").(string)
	parent, err := ExtractFieldByPattern("parent", parent, bucket, "((projects|folders|organizations|billingAccounts)/[a-z0-9A-Z-]*)/locations/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting parent field: %s", err)
	}
	location := d.Get("location").(string)
	location, err = ExtractFieldByPattern("location", location, bucket, "[a-zA-Z]*/[a-z0-9A-Z-]*/locations/([a-z0-9-]*)/buckets/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting location field: %s", err)
	}
	// Set parent to the extracted value.
	d.Set("parent", parent)
	// Set all the other fields to their short forms before forming url and setting ID.
	bucket = tpgresource.GetResourceNameFromSelfLink(bucket)
	name := d.Get("name").(string)
	name = tpgresource.GetResourceNameFromSelfLink(name)
	d.Set("location", location)
	d.Set("bucket", bucket)
	d.Set("name", name)
	return obj, nil
}
