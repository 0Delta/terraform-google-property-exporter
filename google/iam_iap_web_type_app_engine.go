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
	"strings"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var IapWebTypeAppEngineIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"app_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: IapWebTypeAppEngineDiffSuppress,
	},
}

func IapWebTypeAppEngineDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	newParts := strings.Split(new, "appengine-")

	if len(newParts) == 1 {
		// `new` is only the app engine id
		// `old` is always a long name
		if strings.HasSuffix(old, fmt.Sprintf("appengine-%s", new)) {
			return true
		}
	}
	return old == new
}

type IapWebTypeAppEngineIamUpdater struct {
	project string
	appId   string
	d       *schema.ResourceData
	Config  *Config
}

func IapWebTypeAppEngineIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("app_id"); ok {
		values["appId"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/appengine-(?P<appId>[^/]+)", "(?P<project>[^/]+)/(?P<appId>[^/]+)", "(?P<appId>[^/]+)"}, d, config, d.Get("app_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebTypeAppEngineIamUpdater{
		project: values["project"],
		appId:   values["appId"],
		d:       d,
		Config:  config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("app_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting app_id: %s", err)
	}

	return u, nil
}

func IapWebTypeAppEngineIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/appengine-(?P<appId>[^/]+)", "(?P<project>[^/]+)/(?P<appId>[^/]+)", "(?P<appId>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebTypeAppEngineIamUpdater{
		project: values["project"],
		appId:   values["appId"],
		d:       d,
		Config:  config,
	}
	if err := d.Set("app_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting app_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *IapWebTypeAppEngineIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyWebTypeAppEngineUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	obj = map[string]interface{}{
		"options": map[string]interface{}{
			"requestedPolicyVersion": iamPolicyVersion,
		},
	}

	policy, err := sendRequest(u.Config, "POST", project, url, obj)
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

func (u *IapWebTypeAppEngineIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyWebTypeAppEngineUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *IapWebTypeAppEngineIamUpdater) qualifyWebTypeAppEngineUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{IapBasePath}}%s:%s", fmt.Sprintf("projects/%s/iap_web/appengine-%s", u.project, u.appId), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *IapWebTypeAppEngineIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/iap_web/appengine-%s", u.project, u.appId)
}

func (u *IapWebTypeAppEngineIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-iap-webtypeappengine-%s", u.GetResourceId())
}

func (u *IapWebTypeAppEngineIamUpdater) DescribeResource() string {
	return fmt.Sprintf("iap webtypeappengine %q", u.GetResourceId())
}
