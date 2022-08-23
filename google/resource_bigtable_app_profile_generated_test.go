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
)

func TestAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigtableAppProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_profile_id", "instance", "ignore_warnings", "ignore_warnings"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-bt-instance%{random_suffix}"
  cluster {
    cluster_id   = "tf-test-bt-instance%{random_suffix}"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }

  deletion_protection  = "%{deletion_protection}"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-bt-profile%{random_suffix}"

  multi_cluster_routing_use_any = true
  ignore_warnings               = true
}
`, context)
}

func TestAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigtableAppProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_profile_id", "instance", "ignore_warnings", "ignore_warnings"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-bt-instance%{random_suffix}"
  cluster {
    cluster_id   = "tf-test-bt-instance%{random_suffix}"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }

  deletion_protection  = "%{deletion_protection}"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-bt-profile%{random_suffix}"

  single_cluster_routing {
    cluster_id                 = "tf-test-bt-instance%{random_suffix}"
    allow_transactional_writes = true
  }

  ignore_warnings = true
}
`, context)
}

func testAccCheckBigtableAppProfileDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigtable_app_profile" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{BigtableBasePath}}projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("BigtableAppProfile still exists at %s", url)
			}
		}

		return nil
	}
}
