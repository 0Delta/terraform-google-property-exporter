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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecretManagerSecretIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/secretmanager.secretAccessor",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_secret_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/secrets/%s roles/secretmanager.secretAccessor", getTestProjectFromEnv(), fmt.Sprintf("secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSecretManagerSecretIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_secret_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/secrets/%s roles/secretmanager.secretAccessor", getTestProjectFromEnv(), fmt.Sprintf("secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerSecretIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/secretmanager.secretAccessor",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecretManagerSecretIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_secret_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/secrets/%s roles/secretmanager.secretAccessor user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerSecretIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/secretmanager.secretAccessor",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_secret_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/secrets/%s", getTestProjectFromEnv(), fmt.Sprintf("secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecretManagerSecretIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_secret_manager_secret_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/secrets/%s", getTestProjectFromEnv(), fmt.Sprintf("secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecretManagerSecretIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

resource "google_secret_manager_secret_iam_member" "foo" {
  project = google_secret_manager_secret.secret-basic.project
  secret_id = google_secret_manager_secret.secret-basic.secret_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSecretManagerSecretIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_secret_manager_secret_iam_policy" "foo" {
  project = google_secret_manager_secret.secret-basic.project
  secret_id = google_secret_manager_secret.secret-basic.secret_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecretManagerSecretIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_secret_manager_secret_iam_policy" "foo" {
  project = google_secret_manager_secret.secret-basic.project
  secret_id = google_secret_manager_secret.secret-basic.secret_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecretManagerSecretIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

resource "google_secret_manager_secret_iam_binding" "foo" {
  project = google_secret_manager_secret.secret-basic.project
  secret_id = google_secret_manager_secret.secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSecretManagerSecretIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }
}

resource "google_secret_manager_secret_iam_binding" "foo" {
  project = google_secret_manager_secret.secret-basic.project
  secret_id = google_secret_manager_secret.secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
