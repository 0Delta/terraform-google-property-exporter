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
)

func suppressGkeHubEndpointSelfLinkDiff(_, old, new string, _ *schema.ResourceData) bool {
	// The custom expander injects //container.googleapis.com/ if a selflink is supplied.
	selfLink := strings.TrimPrefix(old, "//container.googleapis.com/")
	if selfLink == new {
		return true
	}

	return false
}

func resourceGKEHubMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHubMembershipCreate,
		Read:   resourceGKEHubMembershipRead,
		Update: resourceGKEHubMembershipUpdate,
		Delete: resourceGKEHubMembershipDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHubMembershipImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"membership_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The client-provided identifier of the membership.`,
			},
			"authority": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Authority encodes how Google will recognize identities from this Membership.
See the workload identity documentation for more details:
https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"issuer": {
							Type:     schema.TypeString,
							Required: true,
							Description: `A JSON Web Token (JWT) issuer URI. 'issuer' must start with 'https://' and // be a valid 
with length <2000 characters. For example: 'https://container.googleapis.com/v1/projects/my-project/locations/us-west1/clusters/my-cluster' (must be 'locations' rather than 'zones'). If the cluster is provisioned with Terraform, this is '"https://container.googleapis.com/v1/${google_container_cluster.my-cluster.id}"'.`,
						},
					},
				},
			},
			"endpoint": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gke_cluster": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_link": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: suppressGkeHubEndpointSelfLinkDiff,
										Description: `Self-link of the GCP resource for the GKE cluster.
For example: '//container.googleapis.com/projects/my-project/zones/us-west1-a/clusters/my-cluster'.
It can be at the most 1000 characters in length. If the cluster is provisioned with Terraform,
this can be '"//container.googleapis.com/${google_container_cluster.my-cluster.id}"' or
'google_container_cluster.my-cluster.id'.`,
									},
								},
							},
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels to apply to this membership.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier of the membership.`,
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

func resourceGKEHubMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandGKEHubMembershipLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	endpointProp, err := expandGKEHubMembershipEndpoint(d.Get("endpoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("endpoint"); !isEmptyValue(reflect.ValueOf(endpointProp)) && (ok || !reflect.DeepEqual(v, endpointProp)) {
		obj["endpoint"] = endpointProp
	}
	authorityProp, err := expandGKEHubMembershipAuthority(d.Get("authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authority"); !isEmptyValue(reflect.ValueOf(authorityProp)) && (ok || !reflect.DeepEqual(v, authorityProp)) {
		obj["authority"] = authorityProp
	}

	url, err := replaceVars(d, config, "{{GKEHubBasePath}}projects/{{project}}/locations/global/memberships?membershipId={{membership_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Membership: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Membership: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = gKEHubOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Membership", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Membership: %s", err)
	}

	if err := d.Set("name", flattenGKEHubMembershipName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Membership %q: %#v", d.Id(), res)

	return resourceGKEHubMembershipRead(d, meta)
}

func resourceGKEHubMembershipRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{GKEHubBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("GKEHubMembership %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}

	if err := d.Set("name", flattenGKEHubMembershipName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("labels", flattenGKEHubMembershipLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("endpoint", flattenGKEHubMembershipEndpoint(res["endpoint"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}
	if err := d.Set("authority", flattenGKEHubMembershipAuthority(res["authority"], d, config)); err != nil {
		return fmt.Errorf("Error reading Membership: %s", err)
	}

	return nil
}

func resourceGKEHubMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandGKEHubMembershipLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	authorityProp, err := expandGKEHubMembershipAuthority(d.Get("authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authorityProp)) {
		obj["authority"] = authorityProp
	}

	url, err := replaceVars(d, config, "{{GKEHubBasePath}}projects/{{project}}/locations/global/memberships/{{membership_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Membership %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("authority") {
		updateMask = append(updateMask, "authority")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Membership %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Membership %q: %#v", d.Id(), res)
	}

	err = gKEHubOperationWaitTime(
		config, res, project, "Updating Membership", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGKEHubMembershipRead(d, meta)
}

func resourceGKEHubMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Membership: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{GKEHubBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Membership %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Membership")
	}

	err = gKEHubOperationWaitTime(
		config, res, project, "Deleting Membership", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Membership %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHubMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHubMembershipName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGKEHubMembershipLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGKEHubMembershipEndpoint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gke_cluster"] =
		flattenGKEHubMembershipEndpointGkeCluster(original["gkeCluster"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipEndpointGkeCluster(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resource_link"] =
		flattenGKEHubMembershipEndpointGkeClusterResourceLink(original["resourceLink"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipEndpointGkeClusterResourceLink(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGKEHubMembershipAuthority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["issuer"] =
		flattenGKEHubMembershipAuthorityIssuer(original["issuer"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHubMembershipAuthorityIssuer(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandGKEHubMembershipLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGKEHubMembershipEndpoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGkeCluster, err := expandGKEHubMembershipEndpointGkeCluster(original["gke_cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGkeCluster); val.IsValid() && !isEmptyValue(val) {
		transformed["gkeCluster"] = transformedGkeCluster
	}

	return transformed, nil
}

func expandGKEHubMembershipEndpointGkeCluster(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceLink, err := expandGKEHubMembershipEndpointGkeClusterResourceLink(original["resource_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceLink); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceLink"] = transformedResourceLink
	}

	return transformed, nil
}

func expandGKEHubMembershipEndpointGkeClusterResourceLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if strings.HasPrefix(v.(string), "//") {
		return v, nil
	} else {
		v = "//container.googleapis.com/" + v.(string)
		return v, nil
	}
}

func expandGKEHubMembershipAuthority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIssuer, err := expandGKEHubMembershipAuthorityIssuer(original["issuer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIssuer); val.IsValid() && !isEmptyValue(val) {
		transformed["issuer"] = transformedIssuer
	}

	return transformed, nil
}

func expandGKEHubMembershipAuthorityIssuer(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
