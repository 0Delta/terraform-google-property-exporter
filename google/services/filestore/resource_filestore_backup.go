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

package filestore

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

func ResourceFilestoreBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFilestoreBackupCreate,
		Read:   resourceFilestoreBackupRead,
		Update: resourceFilestoreBackupUpdate,
		Delete: resourceFilestoreBackupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFilestoreBackupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the location of the instance. This can be a region for ENTERPRISE tier instances.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name of the backup. The name must be unique within the specified instance.

The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"source_file_share": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the file share in the source Cloud Filestore instance that the backup is created from.`,
			},
			"source_instance": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels to represent user-provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"capacity_gb": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The amount of bytes needed to allocate a full copy of the snapshot content.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the snapshot was created in RFC3339 text format.`,
			},
			"download_bytes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Amount of bytes that will be downloaded if the backup is restored.`,
			},
			"kms_key_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `KMS key name used for data encryption.`,
			},
			"source_instance_tier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The service tier of the source Cloud Filestore instance that this backup is created from.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The backup state.`,
			},
			"storage_bytes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.`,
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

func resourceFilestoreBackupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreBackupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandFilestoreBackupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	sourceInstanceProp, err := expandFilestoreBackupSourceInstance(d.Get("source_instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceInstanceProp)) && (ok || !reflect.DeepEqual(v, sourceInstanceProp)) {
		obj["sourceInstance"] = sourceInstanceProp
	}
	sourceFileShareProp, err := expandFilestoreBackupSourceFileShare(d.Get("source_file_share"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_file_share"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceFileShareProp)) && (ok || !reflect.DeepEqual(v, sourceFileShareProp)) {
		obj["sourceFileShare"] = sourceFileShareProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "filestore/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{location}}/backups?backupId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Backup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
	})
	if err != nil {
		return fmt.Errorf("Error creating Backup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FilestoreOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Backup", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Backup: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Backup %q: %#v", d.Id(), res)

	return resourceFilestoreBackupRead(d, meta)
}

func resourceFilestoreBackupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{location}}/backups/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FilestoreBackup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}

	if err := d.Set("description", flattenFilestoreBackupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("state", flattenFilestoreBackupState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("create_time", flattenFilestoreBackupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("labels", flattenFilestoreBackupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("capacity_gb", flattenFilestoreBackupCapacityGb(res["capacityGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("storage_bytes", flattenFilestoreBackupStorageBytes(res["storageBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("source_instance", flattenFilestoreBackupSourceInstance(res["sourceInstance"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("source_file_share", flattenFilestoreBackupSourceFileShare(res["sourceFileShare"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("source_instance_tier", flattenFilestoreBackupSourceInstanceTier(res["sourceInstanceTier"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("download_bytes", flattenFilestoreBackupDownloadBytes(res["downloadBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("kms_key_name", flattenFilestoreBackupKmsKeyName(res["kmsKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}

	return nil
}

func resourceFilestoreBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreBackupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandFilestoreBackupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	sourceInstanceProp, err := expandFilestoreBackupSourceInstance(d.Get("source_instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceInstanceProp)) {
		obj["sourceInstance"] = sourceInstanceProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "filestore/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{location}}/backups/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Backup %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("source_instance") {
		updateMask = append(updateMask, "sourceInstance")
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
		Config:               config,
		Method:               "PATCH",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutUpdate),
		ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
	})

	if err != nil {
		return fmt.Errorf("Error updating Backup %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Backup %q: %#v", d.Id(), res)
	}

	err = FilestoreOperationWaitTime(
		config, res, project, "Updating Backup", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceFilestoreBackupRead(d, meta)
}

func resourceFilestoreBackupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "filestore/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{location}}/backups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Backup %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Backup")
	}

	err = FilestoreOperationWaitTime(
		config, res, project, "Deleting Backup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Backup %q: %#v", d.Id(), res)
	return nil
}

func resourceFilestoreBackupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/backups/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFilestoreBackupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupCapacityGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupStorageBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupSourceInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupSourceFileShare(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupSourceInstanceTier(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupDownloadBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFilestoreBackupKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFilestoreBackupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreBackupLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFilestoreBackupSourceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreBackupSourceFileShare(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
