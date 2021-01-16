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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeTargetInstance_targetInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeTargetInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetInstance_targetInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_compute_target_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "zone"},
			},
		},
	})
}

func testAccComputeTargetInstance_targetInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_target_instance" "default" {
  name     = "target%{random_suffix}"
  instance = google_compute_instance.target-vm.id
}

data "google_compute_image" "vmimage" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_instance" "target-vm" {
  name         = "tf-test-target-vm%{random_suffix}"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.vmimage.self_link
    }
  }

  network_interface {
    network = "default"
  }
}
`, context)
}

func testAccCheckComputeTargetInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_target_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/targetInstances/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeTargetInstance still exists at %s", url)
			}
		}

		return nil
	}
}
