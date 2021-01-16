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
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeRegionAutoscaler() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionAutoscalerCreate,
		Read:   resourceComputeRegionAutoscalerRead,
		Update: resourceComputeRegionAutoscalerUpdate,
		Delete: resourceComputeRegionAutoscalerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionAutoscalerImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"autoscaling_policy": {
				Type:     schema.TypeList,
				Required: true,
				Description: `The configuration parameters for the autoscaling algorithm. You can
define one or more of the policies for an autoscaler: cpuUtilization,
customMetricUtilizations, and loadBalancingUtilization.

If none of these are specified, the default will be to autoscale based
on cpuUtilization to 0.6 or 60%.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_replicas": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `The maximum number of instances that the autoscaler can scale up
to. This is required when creating or updating an autoscaler. The
maximum number of replicas should not be lower than minimal number
of replicas.`,
						},
						"min_replicas": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `The minimum number of replicas that the autoscaler can scale down
to. This cannot be less than 0. If not provided, autoscaler will
choose a default value depending on maximum number of instances
allowed.`,
						},
						"cooldown_period": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The number of seconds that the autoscaler should wait before it
starts collecting information from a new instance. This prevents
the autoscaler from collecting information when the instance is
initializing, during which the collected usage would not be
reliable. The default time autoscaler waits is 60 seconds.

Virtual machine initialization times might vary because of
numerous factors. We recommend that you test how long an
instance may take to initialize. To do this, create an instance
and time the startup process.`,
							Default: 60,
						},
						"cpu_utilization": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Description: `Defines the CPU utilization policy that allows the autoscaler to
scale based on the average CPU utilization of a managed instance
group.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `The target CPU utilization that the autoscaler should maintain.
Must be a float value in the range (0, 1]. If not specified, the
default is 0.6.

If the CPU level is below the target utilization, the autoscaler
scales down the number of instances until it reaches the minimum
number of instances you specified or until the average CPU of
your instances reaches the target utilization.

If the average CPU is above the target utilization, the autoscaler
scales up until it reaches the maximum number of instances you
specified or until the average utilization reaches the target
utilization.`,
									},
								},
							},
						},
						"load_balancing_utilization": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration parameters of autoscaling based on a load balancer.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `Fraction of backend capacity utilization (set in HTTP(s) load
balancing configuration) that autoscaler should maintain. Must
be a positive float value. If not defined, the default is 0.8.`,
									},
								},
							},
						},
						"metric": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration parameters of autoscaling based on a custom metric.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										Description: `The identifier (type) of the Stackdriver Monitoring metric.
The metric cannot have negative values.

The metric must have a value type of INT64 or DOUBLE.`,
									},
									"target": {
										Type:     schema.TypeFloat,
										Optional: true,
										Description: `The target value of the metric that autoscaler should
maintain. This must be a positive value. A utilization
metric scales number of virtual machines handling requests
to increase or decrease proportionally to the metric.

For example, a good metric to use as a utilizationTarget is
www.googleapis.com/compute/instance/network/received_bytes_count.
The autoscaler will work to keep this value constant for each
of the instances.`,
									},
									"type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE", ""}, false),
										Description: `Defines how target utilization value is expressed for a
Stackdriver Monitoring metric. Possible values: ["GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"]`,
									},
								},
							},
						},
						"mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"OFF", "ONLY_UP", "ON", ""}, false),
							Description:  `Defines operating mode for this policy. Default value: "ON" Possible values: ["OFF", "ONLY_UP", "ON"]`,
							Default:      "ON",
						},
						"scale_in_control": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Defines scale in controls to reduce the risk of response latency
and outages due to abrupt scale-in events`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_scaled_in_replicas": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `A nested object resource`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fixed": {
													Type:     schema.TypeInt,
													Optional: true,
													Description: `Specifies a fixed number of VM instances. This must be a positive
integer.`,
												},
												"percent": {
													Type:     schema.TypeInt,
													Optional: true,
													Description: `Specifies a percentage of instances between 0 to 100%, inclusive.
For example, specify 80 for 80%.`,
												},
											},
										},
										AtLeastOneOf: []string{},
									},
									"time_window_sec": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `How long back autoscaling should look when computing recommendations
to include directives regarding slower scale down, as described above.`,
									},
								},
							},
							AtLeastOneOf: []string{},
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGCPName,
				Description: `Name of the resource. The name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"target": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL of the managed instance group that this autoscaler will scale.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL of the region where the instance group resides.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeRegionAutoscalerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeRegionAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !isEmptyValue(reflect.ValueOf(autoscalingPolicyProp)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeRegionAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	regionProp, err := expandComputeRegionAutoscalerRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionAutoscaler: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionAutoscaler: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionAutoscaler: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating RegionAutoscaler", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionAutoscaler: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionAutoscaler %q: %#v", d.Id(), res)

	return resourceComputeRegionAutoscalerRead(d, meta)
}

func resourceComputeRegionAutoscalerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionAutoscaler: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionAutoscaler %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeRegionAutoscalerCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionAutoscalerName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionAutoscalerDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("autoscaling_policy", flattenComputeRegionAutoscalerAutoscalingPolicy(res["autoscalingPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("target", flattenComputeRegionAutoscalerTarget(res["target"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionAutoscalerRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}

	return nil
}

func resourceComputeRegionAutoscalerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionAutoscaler: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeRegionAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeRegionAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	regionProp, err := expandComputeRegionAutoscalerRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers?autoscaler={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionAutoscaler %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionAutoscaler %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating RegionAutoscaler %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating RegionAutoscaler", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeRegionAutoscalerRead(d, meta)
}

func resourceComputeRegionAutoscalerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionAutoscaler: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionAutoscaler %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionAutoscaler")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting RegionAutoscaler", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionAutoscaler %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionAutoscalerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/autoscalers/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionAutoscalerCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_replicas"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMinReplicas(original["minNumReplicas"], d, config)
	transformed["max_replicas"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(original["maxNumReplicas"], d, config)
	transformed["cooldown_period"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(original["coolDownPeriodSec"], d, config)
	transformed["mode"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMode(original["mode"], d, config)
	transformed["scale_in_control"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControl(original["scaleInControl"], d, config)
	transformed["cpu_utilization"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(original["cpuUtilization"], d, config)
	transformed["metric"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMetric(original["customMetricUtilizations"], d, config)
	transformed["load_balancing_utilization"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["loadBalancingUtilization"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeRegionAutoscalerAutoscalingPolicyMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["max_scaled_in_replicas"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(original["maxScaledInReplicas"], d, config)
	transformed["time_window_sec"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(original["timeWindowSec"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["fixed"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(original["fixed"], d, config)
	transformed["percent"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(original["percent"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeRegionAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["utilizationTarget"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetric(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"name":   flattenComputeRegionAutoscalerAutoscalingPolicyMetricName(original["metric"], d, config),
			"target": flattenComputeRegionAutoscalerAutoscalingPolicyMetricTarget(original["utilizationTarget"], d, config),
			"type":   flattenComputeRegionAutoscalerAutoscalingPolicyMetricType(original["utilizationTargetType"], d, config),
		})
	}
	return transformed
}
func flattenComputeRegionAutoscalerAutoscalingPolicyMetricName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetricType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["utilizationTarget"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerTarget(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionAutoscalerRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeRegionAutoscalerName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(original["min_replicas"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["minNumReplicas"] = transformedMinReplicas
	}

	transformedMaxReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(original["max_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxReplicas); val.IsValid() && !isEmptyValue(val) {
		transformed["maxNumReplicas"] = transformedMaxReplicas
	}

	transformedCooldownPeriod, err := expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(original["cooldown_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCooldownPeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["coolDownPeriodSec"] = transformedCooldownPeriod
	}

	transformedMode, err := expandComputeRegionAutoscalerAutoscalingPolicyMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !isEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedScaleInControl, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControl(original["scale_in_control"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleInControl); val.IsValid() && !isEmptyValue(val) {
		transformed["scaleInControl"] = transformedScaleInControl
	}

	transformedCpuUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(original["cpu_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuUtilization); val.IsValid() && !isEmptyValue(val) {
		transformed["cpuUtilization"] = transformedCpuUtilization
	}

	transformedMetric, err := expandComputeRegionAutoscalerAutoscalingPolicyMetric(original["metric"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetric); val.IsValid() && !isEmptyValue(val) {
		transformed["customMetricUtilizations"] = transformedMetric
	}

	transformedLoadBalancingUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["load_balancing_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoadBalancingUtilization); val.IsValid() && !isEmptyValue(val) {
		transformed["loadBalancingUtilization"] = transformedLoadBalancingUtilization
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxScaledInReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(original["max_scaled_in_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxScaledInReplicas); val.IsValid() && !isEmptyValue(val) {
		transformed["maxScaledInReplicas"] = transformedMaxScaledInReplicas
	}

	transformedTimeWindowSec, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(original["time_window_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeWindowSec); val.IsValid() && !isEmptyValue(val) {
		transformed["timeWindowSec"] = transformedTimeWindowSec
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixed, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(original["fixed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixed); val.IsValid() && !isEmptyValue(val) {
		transformed["fixed"] = transformedFixed
	}

	transformedPercent, err := expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(original["percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !isEmptyValue(val) {
		transformed["percent"] = transformedPercent
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !isEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetric(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["metric"] = transformedName
		}

		transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !isEmptyValue(val) {
			transformed["utilizationTarget"] = transformedTarget
		}

		transformedType, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
			transformed["utilizationTargetType"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !isEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
