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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceDialogflowCXIntent() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowCXIntentCreate,
		Read:   resourceDialogflowCXIntentRead,
		Update: resourceDialogflowCXIntentUpdate,
		Delete: resourceDialogflowCXIntentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowCXIntentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Description:  `The human-readable name of the intent, unique within the agent.`,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 140),
				Description:  `Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.`,
			},
			"is_fallback": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"language_code": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The language of the following fields in intent:
Intent.training_phrases.parts.text
If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.`,
			},
			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The collection of parameters associated with the intent.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entity_type": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The entity type of the parameter.
Format: projects/-/locations/-/agents/-/entityTypes/<System Entity Type ID> for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID> for developer entity types.`,
						},
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The unique identifier of the parameter. This field is used by training phrases to annotate their parts.`,
						},
						"is_list": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Indicates whether the parameter represents a list of values.`,
						},
						"redact": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: `Indicates whether the parameter content should be redacted in log. If redaction is enabled, the parameter content will be replaced by parameter name during logging.
Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.`,
						},
					},
				},
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The agent to create an intent for.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.`,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The priority of this intent. Higher numbers represent higher priorities.
If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
If the supplied value is negative, the intent is ignored in runtime detect intent requests.`,
			},
			"training_phrases": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The collection of training phrases the agent is trained on to identify the intent.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parts": {
							Type:     schema.TypeList,
							Required: true,
							Description: `The ordered list of training phrase parts. The parts are concatenated in order to form the training phrase.
Note: The API does not automatically annotate training phrases like the Dialogflow Console does.
Note: Do not forget to include whitespace at part boundaries, so the training phrase is well formatted when the parts are concatenated.
If the training phrase does not need to be annotated with parameters, you just need a single part with only the Part.text field set.
If you want to annotate the training phrase, you must create multiple parts, where the fields of each part are populated in one of two ways:
Part.text is set to a part of the phrase that has no parameters.
Part.text is set to a part of the phrase that you want to annotate, and the parameterId field is set.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"text": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The text for this part.`,
									},
									"parameter_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The parameter used to annotate this part of the training phrase. This field is required for annotated parts of the training phrase.`,
									},
								},
							},
						},
						"repeat_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Indicates how many times this example was added to the intent.`,
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The unique identifier of the training phrase.`,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique identifier of the intent.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDialogflowCXIntentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXIntentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	trainingPhrasesProp, err := expandDialogflowCXIntentTrainingPhrases(d.Get("training_phrases"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("training_phrases"); !isEmptyValue(reflect.ValueOf(trainingPhrasesProp)) && (ok || !reflect.DeepEqual(v, trainingPhrasesProp)) {
		obj["trainingPhrases"] = trainingPhrasesProp
	}
	parametersProp, err := expandDialogflowCXIntentParameters(d.Get("parameters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parameters"); !isEmptyValue(reflect.ValueOf(parametersProp)) && (ok || !reflect.DeepEqual(v, parametersProp)) {
		obj["parameters"] = parametersProp
	}
	priorityProp, err := expandDialogflowCXIntentPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	isFallbackProp, err := expandDialogflowCXIntentIsFallback(d.Get("is_fallback"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_fallback"); !isEmptyValue(reflect.ValueOf(isFallbackProp)) && (ok || !reflect.DeepEqual(v, isFallbackProp)) {
		obj["isFallback"] = isFallbackProp
	}
	labelsProp, err := expandDialogflowCXIntentLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandDialogflowCXIntentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	languageCodeProp, err := expandDialogflowCXIntentLanguageCode(d.Get("language_code"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("language_code"); !isEmptyValue(reflect.ValueOf(languageCodeProp)) && (ok || !reflect.DeepEqual(v, languageCodeProp)) {
		obj["languageCode"] = languageCodeProp
	}

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/intents")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Intent: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Intent: %s", err)
	}
	if err := d.Set("name", flattenDialogflowCXIntentName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{parent}}/intents/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Intent %q: %#v", d.Id(), res)

	return resourceDialogflowCXIntentRead(d, meta)
}

func resourceDialogflowCXIntentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/intents/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DialogflowCXIntent %q", d.Id()))
	}

	if err := d.Set("name", flattenDialogflowCXIntentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowCXIntentDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("training_phrases", flattenDialogflowCXIntentTrainingPhrases(res["trainingPhrases"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("parameters", flattenDialogflowCXIntentParameters(res["parameters"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("priority", flattenDialogflowCXIntentPriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("is_fallback", flattenDialogflowCXIntentIsFallback(res["isFallback"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("labels", flattenDialogflowCXIntentLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("description", flattenDialogflowCXIntentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}
	if err := d.Set("language_code", flattenDialogflowCXIntentLanguageCode(res["languageCode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Intent: %s", err)
	}

	return nil
}

func resourceDialogflowCXIntentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXIntentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	trainingPhrasesProp, err := expandDialogflowCXIntentTrainingPhrases(d.Get("training_phrases"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("training_phrases"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trainingPhrasesProp)) {
		obj["trainingPhrases"] = trainingPhrasesProp
	}
	parametersProp, err := expandDialogflowCXIntentParameters(d.Get("parameters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parameters"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, parametersProp)) {
		obj["parameters"] = parametersProp
	}
	priorityProp, err := expandDialogflowCXIntentPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	isFallbackProp, err := expandDialogflowCXIntentIsFallback(d.Get("is_fallback"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_fallback"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, isFallbackProp)) {
		obj["isFallback"] = isFallbackProp
	}
	labelsProp, err := expandDialogflowCXIntentLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandDialogflowCXIntentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/intents/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Intent %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("training_phrases") {
		updateMask = append(updateMask, "trainingPhrases")
	}

	if d.HasChange("parameters") {
		updateMask = append(updateMask, "parameters")
	}

	if d.HasChange("priority") {
		updateMask = append(updateMask, "priority")
	}

	if d.HasChange("is_fallback") {
		updateMask = append(updateMask, "isFallback")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
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

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Intent %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Intent %q: %#v", d.Id(), res)
	}

	return resourceDialogflowCXIntentRead(d, meta)
}

func resourceDialogflowCXIntentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/intents/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	log.Printf("[DEBUG] Deleting Intent %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Intent")
	}

	log.Printf("[DEBUG] Finished deleting Intent %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowCXIntentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value and parent contains slashes
	if err := parseImportId([]string{
		"(?P<parent>.+)/intents/(?P<name>[^/]+)",
		"(?P<parent>.+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{parent}}/intents/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowCXIntentName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenDialogflowCXIntentDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentTrainingPhrases(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"id":           flattenDialogflowCXIntentTrainingPhrasesId(original["id"], d, config),
			"parts":        flattenDialogflowCXIntentTrainingPhrasesParts(original["parts"], d, config),
			"repeat_count": flattenDialogflowCXIntentTrainingPhrasesRepeatCount(original["repeatCount"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowCXIntentTrainingPhrasesId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentTrainingPhrasesParts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"text":         flattenDialogflowCXIntentTrainingPhrasesPartsText(original["text"], d, config),
			"parameter_id": flattenDialogflowCXIntentTrainingPhrasesPartsParameterId(original["parameterId"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowCXIntentTrainingPhrasesPartsText(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentTrainingPhrasesPartsParameterId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentTrainingPhrasesRepeatCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenDialogflowCXIntentParameters(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"id":          flattenDialogflowCXIntentParametersId(original["id"], d, config),
			"entity_type": flattenDialogflowCXIntentParametersEntityType(original["entityType"], d, config),
			"is_list":     flattenDialogflowCXIntentParametersIsList(original["isList"], d, config),
			"redact":      flattenDialogflowCXIntentParametersRedact(original["redact"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowCXIntentParametersId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentParametersEntityType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentParametersIsList(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentParametersRedact(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentPriority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenDialogflowCXIntentIsFallback(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXIntentLanguageCode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandDialogflowCXIntentDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrases(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandDialogflowCXIntentTrainingPhrasesId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedParts, err := expandDialogflowCXIntentTrainingPhrasesParts(original["parts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedParts); val.IsValid() && !isEmptyValue(val) {
			transformed["parts"] = transformedParts
		}

		transformedRepeatCount, err := expandDialogflowCXIntentTrainingPhrasesRepeatCount(original["repeat_count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRepeatCount); val.IsValid() && !isEmptyValue(val) {
			transformed["repeatCount"] = transformedRepeatCount
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXIntentTrainingPhrasesId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrasesParts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXIntentTrainingPhrasesPartsText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
			transformed["text"] = transformedText
		}

		transformedParameterId, err := expandDialogflowCXIntentTrainingPhrasesPartsParameterId(original["parameter_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedParameterId); val.IsValid() && !isEmptyValue(val) {
			transformed["parameterId"] = transformedParameterId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXIntentTrainingPhrasesPartsText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrasesPartsParameterId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrasesRepeatCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParameters(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandDialogflowCXIntentParametersId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedEntityType, err := expandDialogflowCXIntentParametersEntityType(original["entity_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEntityType); val.IsValid() && !isEmptyValue(val) {
			transformed["entityType"] = transformedEntityType
		}

		transformedIsList, err := expandDialogflowCXIntentParametersIsList(original["is_list"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIsList); val.IsValid() && !isEmptyValue(val) {
			transformed["isList"] = transformedIsList
		}

		transformedRedact, err := expandDialogflowCXIntentParametersRedact(original["redact"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRedact); val.IsValid() && !isEmptyValue(val) {
			transformed["redact"] = transformedRedact
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXIntentParametersId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParametersEntityType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParametersIsList(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParametersRedact(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentPriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentIsFallback(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDialogflowCXIntentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentLanguageCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
