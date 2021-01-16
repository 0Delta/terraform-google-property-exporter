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
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeGlobalAddress_globalAddressBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeGlobalAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalAddress_globalAddressBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_global_address.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeGlobalAddress_globalAddressBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_global_address" "default" {
  name = "global-appserver-ip-%s"
}
`, val,
	)
}

func testAccCheckComputeGlobalAddressDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_global_address" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/addresses/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeGlobalAddress still exists at %s", url)
		}
	}

	return nil
}