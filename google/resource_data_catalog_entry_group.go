// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDataCatalogEntryGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCatalogEntryGroupCreate,
		Read:   resourceDataCatalogEntryGroupRead,
		Update: resourceDataCatalogEntryGroupUpdate,
		Delete: resourceDataCatalogEntryGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataCatalogEntryGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"entry_group_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^[A-z_][A-z0-9_]{0,63}$`),
				Description: `The id of the entry group to create. The id must begin with a letter or underscore,
contain only English letters, numbers and underscores, and be at most 64 characters.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A short name to identify the entry group, for example, "analytics data - jan 2011".`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `EntryGroup location region.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDataCatalogEntryGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogEntryGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogEntryGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}projects/{{project}}/locations/{{region}}/entryGroups?entryGroupId={{entry_group_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EntryGroup: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EntryGroup: %s", err)
	}
	if err := d.Set("name", flattenDataCatalogEntryGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EntryGroup %q: %#v", d.Id(), res)

	return resourceDataCatalogEntryGroupRead(d, meta)
}

func resourceDataCatalogEntryGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DataCatalogEntryGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading EntryGroup: %s", err)
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading EntryGroup: %s", err)
	}

	if err := d.Set("name", flattenDataCatalogEntryGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryGroup: %s", err)
	}
	if err := d.Set("display_name", flattenDataCatalogEntryGroupDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryGroup: %s", err)
	}
	if err := d.Set("description", flattenDataCatalogEntryGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntryGroup: %s", err)
	}

	return nil
}

func resourceDataCatalogEntryGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogEntryGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogEntryGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EntryGroup %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating EntryGroup %q: %s", d.Id(), err)
	}

	return resourceDataCatalogEntryGroupRead(d, meta)
}

func resourceDataCatalogEntryGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EntryGroup %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EntryGroup")
	}

	log.Printf("[DEBUG] Finished deleting EntryGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceDataCatalogEntryGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	egRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/entryGroups/(.+)")

	parts := egRegex.FindStringSubmatch(name)
	if len(parts) != 4 {
		return nil, fmt.Errorf("entry group name does not fit the format %s", egRegex)
	}
	d.Set("project", parts[1])
	d.Set("region", parts[2])
	d.Set("entry_group_id", parts[3])
	return []*schema.ResourceData{d}, nil
}

func flattenDataCatalogEntryGroupName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogEntryGroupDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogEntryGroupDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandDataCatalogEntryGroupDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
