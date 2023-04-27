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

func TestAccNotebooksRuntime_notebookRuntimeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNotebooksRuntimeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksRuntime_notebookRuntimeBasicExample(context),
			},
			{
				ResourceName:            "google_notebooks_runtime.runtime",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNotebooksRuntime_notebookRuntimeBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime" {
  name = "tf-test-notebooks-runtime%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}
`, context)
}

func TestAccNotebooksRuntime_notebookRuntimeBasicGpuExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNotebooksRuntimeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksRuntime_notebookRuntimeBasicGpuExample(context),
			},
			{
				ResourceName:            "google_notebooks_runtime.runtime_gpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNotebooksRuntime_notebookRuntimeBasicGpuExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime_gpu" {
  name = "tf-test-notebooks-runtime-gpu%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  software_config {
    install_gpu_driver = true
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
      accelerator_config {
        core_count = "1"
        type = "NVIDIA_TESLA_V100"
      }
    }
  }
}
`, context)
}

func TestAccNotebooksRuntime_notebookRuntimeBasicContainerExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNotebooksRuntimeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksRuntime_notebookRuntimeBasicContainerExample(context),
			},
			{
				ResourceName:            "google_notebooks_runtime.runtime_container",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNotebooksRuntime_notebookRuntimeBasicContainerExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime_container" {
  name = "tf-test-notebooks-runtime-container%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
      container_images {
        repository = "gcr.io/deeplearning-platform-release/base-cpu"
        tag = "latest"
      }
      container_images {
        repository = "gcr.io/deeplearning-platform-release/beam-notebooks"
        tag = "latest"
      }
    }
  }
}
`, context)
}

func TestAccNotebooksRuntime_notebookRuntimeKernelsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNotebooksRuntimeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksRuntime_notebookRuntimeKernelsExample(context),
			},
			{
				ResourceName:            "google_notebooks_runtime.runtime_container",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNotebooksRuntime_notebookRuntimeKernelsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime_container" {
  name = "tf-test-notebooks-runtime-kernel%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  software_config {
    kernels {
      repository = "gcr.io/deeplearning-platform-release/base-cpu"
      tag        = "latest"
    }
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}
`, context)
}

func TestAccNotebooksRuntime_notebookRuntimeScriptExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNotebooksRuntimeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksRuntime_notebookRuntimeScriptExample(context),
			},
			{
				ResourceName:            "google_notebooks_runtime.runtime_container",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNotebooksRuntime_notebookRuntimeScriptExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_runtime" "runtime_container" {
  name = "tf-test-notebooks-runtime-script%{random_suffix}"
  location = "us-central1"
  access_config {
    access_type = "SINGLE_USER"
    runtime_owner = "admin@hashicorptest.com"
  }
  software_config {
    post_startup_script_behavior = "RUN_EVERY_START"
  }
  virtual_machine {
    virtual_machine_config {
      machine_type = "n1-standard-4"
      data_disk {
        initialize_params {
          disk_size_gb = "100"
          disk_type = "PD_STANDARD"
        }
      }
    }
  }
}
`, context)
}

func testAccCheckNotebooksRuntimeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_notebooks_runtime" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/runtimes/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NotebooksRuntime still exists at %s", url)
			}
		}

		return nil
	}
}
