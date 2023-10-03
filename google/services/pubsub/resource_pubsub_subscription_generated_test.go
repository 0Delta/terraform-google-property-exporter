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

package pubsub_test

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

func TestAccPubsubSubscription_pubsubSubscriptionPushExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionPushExample(context),
			},
			{
				ResourceName:            "google_pubsub_subscription.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"topic", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionPushExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  ack_deadline_seconds = 20

  labels = {
    foo = "bar"
  }

  push_config {
    push_endpoint = "https://example.com/push"

    attributes = {
      x-goog-version = "v1"
    }
  }
}
`, context)
}

func TestAccPubsubSubscription_pubsubSubscriptionPullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionPullExample(context),
			},
			{
				ResourceName:            "google_pubsub_subscription.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"topic", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionPullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  labels = {
    foo = "bar"
  }

  # 20 minutes
  message_retention_duration = "1200s"
  retain_acked_messages      = true

  ack_deadline_seconds = 20

  expiration_policy {
    ttl = "300000.5s"
  }
  retry_policy {
    minimum_backoff = "10s"
  }

  enable_message_ordering    = false
}
`, context)
}

func TestAccPubsubSubscription_pubsubSubscriptionDeadLetterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionDeadLetterExample(context),
			},
			{
				ResourceName:            "google_pubsub_subscription.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"topic", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionDeadLetterExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_topic" "example_dead_letter" {
  name = "tf-test-example-topic%{random_suffix}-dead-letter"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  dead_letter_policy {
    dead_letter_topic = google_pubsub_topic.example_dead_letter.id
    max_delivery_attempts = 10
  }
}
`, context)
}

func TestAccPubsubSubscription_pubsubSubscriptionPushBqExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionPushBqExample(context),
			},
			{
				ResourceName:            "google_pubsub_subscription.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"topic", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionPushBqExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  bigquery_config {
    table = "${google_bigquery_table.test.project}.${google_bigquery_table.test.dataset_id}.${google_bigquery_table.test.table_id}"
  }

  depends_on = [google_project_iam_member.viewer, google_project_iam_member.editor]
}

data "google_project" "project" {
}

resource "google_project_iam_member" "viewer" {
  project = data.google_project.project.project_id
  role   = "roles/bigquery.metadataViewer"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "editor" {
  project = data.google_project.project.project_id
  role   = "roles/bigquery.dataEditor"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_example_dataset%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_example_table%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id

  schema = <<EOF
[
  {
    "name": "data",
    "type": "STRING",
    "mode": "NULLABLE",
    "description": "The data"
  }
]
EOF
}
`, context)
}

func TestAccPubsubSubscription_pubsubSubscriptionPushCloudstorageExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionPushCloudstorageExample(context),
			},
			{
				ResourceName:            "google_pubsub_subscription.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"topic", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionPushCloudstorageExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "example" {
  name  = "tf-test-example-bucket%{random_suffix}"
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  cloud_storage_config {
    bucket = google_storage_bucket.example.name

    filename_prefix = "pre-"
    filename_suffix = "-%{random_suffix}"
  
    max_bytes = 1000
    max_duration = "300s"
  }
  depends_on = [ 
    google_storage_bucket.example,
    google_storage_bucket_iam_member.admin,
  ]
}

data "google_project" "project" {
}

resource "google_storage_bucket_iam_member" "admin" {
  bucket = google_storage_bucket.example.name
  role   = "roles/storage.admin"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}
`, context)
}

func TestAccPubsubSubscription_pubsubSubscriptionPushCloudstorageAvroExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSubscription_pubsubSubscriptionPushCloudstorageAvroExample(context),
			},
			{
				ResourceName:            "google_pubsub_subscription.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"topic", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPubsubSubscription_pubsubSubscriptionPushCloudstorageAvroExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "example" {
  name  = "tf-test-example-bucket%{random_suffix}"
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
}

resource "google_pubsub_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_topic.example.name

  cloud_storage_config {
    bucket = google_storage_bucket.example.name

    filename_prefix = "pre-"
    filename_suffix = "-%{random_suffix}"
  
    max_bytes = 1000
    max_duration = "300s"
  
    avro_config {
      write_metadata = true
    }
  }
  depends_on = [ 
    google_storage_bucket.example,
    google_storage_bucket_iam_member.admin,
  ]
}

data "google_project" "project" {
}

resource "google_storage_bucket_iam_member" "admin" {
  bucket = google_storage_bucket.example.name
  role   = "roles/storage.admin"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}
`, context)
}

func testAccCheckPubsubSubscriptionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_pubsub_subscription" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PubsubBasePath}}projects/{{project}}/subscriptions/{{name}}")
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
				return fmt.Errorf("PubsubSubscription still exists at %s", url)
			}
		}

		return nil
	}
}