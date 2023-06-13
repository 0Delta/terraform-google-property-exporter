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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDialogflowCXPage_dialogflowcxPageFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDialogflowCXPageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXPage_dialogflowcxPageFullExample(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_page.basic_page",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDialogflowCXPage_dialogflowcxPageFullExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_dialogflow_cx_agent" "agent" {
  display_name = "tf-test-dialogflowcx-agent%{random_suffix}"
  location = "global"
  default_language_code = "en"
  supported_language_codes = ["fr","de","es"]
  time_zone = "America/New_York"
  description = "Example description."
  avatar_uri = "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"
  enable_stackdriver_logging = true
  enable_spell_correction    = true
	speech_to_text_settings {
		enable_speech_adaptation = true
	}
}


resource "google_dialogflow_cx_page" "basic_page" {
  parent       = google_dialogflow_cx_agent.agent.start_flow
  display_name = "MyPage"

  entry_fulfillment {
		messages {
			text {
				text = ["Welcome to page"]
			}
		}
   }

   form {
		parameters {
			display_name = "param1"
			entity_type  = "projects/-/locations/-/agents/-/entityTypes/sys.date"
			fill_behavior {
				initial_prompt_fulfillment {
					messages {
						text {
							text = ["Please provide param1"]
						}
					}
				}
			}
			required = "true"
			redact   = "true"
		}
	}

    transition_routes {
		condition = "$page.params.status = 'FINAL'"
		trigger_fulfillment {
			messages {
				text {
					text = ["information completed, navigating to page 2"]
				}
			}
		}
		target_page = google_dialogflow_cx_page.my_page2.id
	}
} 

resource "google_dialogflow_cx_page" "my_page2" {
    parent       = google_dialogflow_cx_agent.agent.start_flow
    display_name  = "MyPage2"
}
`, context)
}

func testAccCheckDialogflowCXPageDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dialogflow_cx_page" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DialogflowCXBasePath}}{{parent}}/pages/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DialogflowCXPage still exists at %s", url)
			}
		}

		return nil
	}
}
