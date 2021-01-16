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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudRunService_cloudRunServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceBasicExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceSqlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceSqlExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceSqlExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale"      = "1000"
        "run.googleapis.com/cloudsql-instances" = "%{project}:us-central1:${google_sql_database_instance.instance.name}"
        "run.googleapis.com/client-name"        = "terraform"
      }
    }
  }
  autogenerate_revision_name = true
}

resource "google_sql_database_instance" "instance" {
  name   = "tf-test-cloudrun-sql%{random_suffix}"
  region = "us-east1"
  settings {
    tier = "db-f1-micro"
  }
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceNoauthExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceNoauthExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceNoauthExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceMultipleEnvironmentVariablesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceMultipleEnvironmentVariablesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceMultipleEnvironmentVariablesExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        env {
          name = "SOURCE"
          value = "remote"
        }
        env {
          name = "TARGET"
          value = "home"
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = true
}
`, context)
}

func testAccCheckCloudRunServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_run_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudRunBasePath}}apis/serving.knative.dev/v1/namespaces/{{project}}/services/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("CloudRunService still exists at %s", url)
			}
		}

		return nil
	}
}
