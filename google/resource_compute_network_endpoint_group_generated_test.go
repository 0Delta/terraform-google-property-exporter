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

func TestAccComputeNetworkEndpointGroup_networkEndpointGroupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkEndpointGroup_networkEndpointGroupExample(context),
			},
			{
				ResourceName:            "google_compute_network_endpoint_group.neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "subnetwork", "zone"},
			},
		},
	})
}

func testAccComputeNetworkEndpointGroup_networkEndpointGroupExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_compute_network_endpoint_group" "neg" {
  name         = "tf-test-my-lb-neg%{random_suffix}"
  network      = google_compute_network.default.id
  subnetwork   = google_compute_subnetwork.default.id
  default_port = "90"
  zone         = "us-central1-a"
}

resource "google_compute_network" "default" {
  name                    = "tf-test-neg-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-neg-subnetwork%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccComputeNetworkEndpointGroup_networkEndpointGroupNonGcpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkEndpointGroup_networkEndpointGroupNonGcpExample(context),
			},
			{
				ResourceName:            "google_compute_network_endpoint_group.neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "subnetwork", "zone"},
			},
		},
	})
}

func testAccComputeNetworkEndpointGroup_networkEndpointGroupNonGcpExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_compute_network_endpoint_group" "neg" {
  name                  = "tf-test-my-lb-neg%{random_suffix}"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "NON_GCP_PRIVATE_IP_PORT"
}

resource "google_compute_network_endpoint" "default-endpoint" {
  network_endpoint_group = google_compute_network_endpoint_group.neg.name
  port = google_compute_network_endpoint_group.neg.default_port
  ip_address = "127.0.0.1"
}

resource "google_compute_network" "default" {
  name = "tf-test-neg-network%{random_suffix}"
}
`, context)
}

func testAccCheckComputeNetworkEndpointGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_network_endpoint_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
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
				return fmt.Errorf("ComputeNetworkEndpointGroup still exists at %s", url)
			}
		}

		return nil
	}
}
