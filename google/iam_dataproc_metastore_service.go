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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var DataprocMetastoreServiceIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"service_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type DataprocMetastoreServiceIamUpdater struct {
	project   string
	location  string
	serviceId string
	d         TerraformResourceData
	Config    *transport_tpg.Config
}

func DataprocMetastoreServiceIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("service_id"); ok {
		values["service_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/services/(?P<service_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<service_id>[^/]+)", "(?P<location>[^/]+)/(?P<service_id>[^/]+)", "(?P<service_id>[^/]+)"}, d, config, d.Get("service_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataprocMetastoreServiceIamUpdater{
		project:   values["project"],
		location:  values["location"],
		serviceId: values["service_id"],
		d:         d,
		Config:    config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("service_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting service_id: %s", err)
	}

	return u, nil
}

func DataprocMetastoreServiceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/services/(?P<service_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<service_id>[^/]+)", "(?P<location>[^/]+)/(?P<service_id>[^/]+)", "(?P<service_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataprocMetastoreServiceIamUpdater{
		project:   values["project"],
		location:  values["location"],
		serviceId: values["service_id"],
		d:         d,
		Config:    config,
	}
	if err := d.Set("service_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting service_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataprocMetastoreServiceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyServiceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "GET", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataprocMetastoreServiceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyServiceUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataprocMetastoreServiceIamUpdater) qualifyServiceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataprocMetastoreBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/services/%s", u.project, u.location, u.serviceId), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataprocMetastoreServiceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s", u.project, u.location, u.serviceId)
}

func (u *DataprocMetastoreServiceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-dataprocmetastore-service-%s", u.GetResourceId())
}

func (u *DataprocMetastoreServiceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("dataprocmetastore service %q", u.GetResourceId())
}
