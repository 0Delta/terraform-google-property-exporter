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

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputeNodeGroup_nodeGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNodeGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeGroup_nodeGroupBasicExample(context),
			},
			{
				ResourceName:            "google_compute_node_group.nodes",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"node_template", "initial_size", "zone"},
			},
		},
	})
}

func testAccComputeNodeGroup_nodeGroupBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_node_template" "soletenant-tmpl" {
  name      = "tf-test-soletenant-tmpl%{random_suffix}"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}

resource "google_compute_node_group" "nodes" {
  name        = "tf-test-soletenant-group%{random_suffix}"
  zone        = "us-central1-f"
  description = "example google_compute_node_group for Terraform Google Provider"

  size          = 1
  node_template = google_compute_node_template.soletenant-tmpl.id
}
`, context)
}

func TestAccComputeNodeGroup_nodeGroupAutoscalingPolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNodeGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeGroup_nodeGroupAutoscalingPolicyExample(context),
			},
			{
				ResourceName:            "google_compute_node_group.nodes",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"node_template", "initial_size", "zone"},
			},
		},
	})
}

func testAccComputeNodeGroup_nodeGroupAutoscalingPolicyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_node_template" "soletenant-tmpl" {
  name      = "tf-test-soletenant-tmpl%{random_suffix}"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}

resource "google_compute_node_group" "nodes" {
  name        = "tf-test-soletenant-group%{random_suffix}"
  zone        = "us-central1-f"
  description = "example google_compute_node_group for Terraform Google Provider"
  maintenance_policy = "RESTART_IN_PLACE"
  maintenance_window {
    start_time = "08:00"
  }
  initial_size  = 1
  node_template = google_compute_node_template.soletenant-tmpl.id
  autoscaling_policy {
    mode      = "ONLY_SCALE_OUT"
    min_nodes = 1
    max_nodes = 10
  }
}
`, context)
}

func TestAccComputeNodeGroup_nodeGroupShareSettingsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNodeGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeGroup_nodeGroupShareSettingsExample(context),
			},
			{
				ResourceName:            "google_compute_node_group.nodes",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"node_template", "initial_size", "zone"},
			},
		},
	})
}

func testAccComputeNodeGroup_nodeGroupShareSettingsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "guest_project" {
  project_id      = "tf-test-project-id%{random_suffix}"
  name            = "tf-test-project-name%{random_suffix}"
  org_id          = "%{org_id}"
}

resource "google_compute_node_template" "soletenant-tmpl" {
  name      = "tf-test-soletenant-tmpl%{random_suffix}"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}

resource "google_compute_node_group" "nodes" {
  name        = "tf-test-soletenant-group%{random_suffix}"
  zone        = "us-central1-f"
  description = "example google_compute_node_group for Terraform Google Provider"

  size          = 1
  node_template = google_compute_node_template.soletenant-tmpl.id

  share_settings {
    share_type = "SPECIFIC_PROJECTS"
    project_map {
      id = google_project.guest_project.project_id
      project_id = google_project.guest_project.project_id
    }
  }
}
`, context)
}

func testAccCheckComputeNodeGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_node_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}")
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
				return fmt.Errorf("ComputeNodeGroup still exists at %s", url)
			}
		}

		return nil
	}
}
