package provider

import (
	"fmt"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/integration"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccProjectIntegrationExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["id"] == "" {
			return fmt.Errorf("no ID is set")
		}
		if rs.Primary.Attributes["project_id"] == "" {
			return fmt.Errorf("no ProjectID is set")
		}

		id, _ := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		projectId, _ := strconv.ParseInt(rs.Primary.Attributes["project_id"], 10, 64)

		response, err := testClient().Integration.GetProjectProjectIDIntegrationsIntegrationID(&integration.GetProjectProjectIDIntegrationsIntegrationIDParams{
			ProjectID:     projectId,
			IntegrationID: id,
		}, nil)
		if err != nil {
			return fmt.Errorf("project integration does not exist: %s", err.Error())
		}

		if response.Payload.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("integration name mismatch: %s != %s", response.Payload.Name, rs.Primary.Attributes["name"])
		}
		return nil
	}
}

func testAccProjectIntegrationDependencyConfig(nameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None-%[1]s"
  none       = {}
}

resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Repo-%[1]s"
  url        = "git@github.com:example/test.git"
  branch     = "main"
  ssh_key_id = semaphoreui_project_key.test.id
}

resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Inventory-%[1]s"
  ssh_key_id = semaphoreui_project_key.test.id
  file = {
    path          = "path/to/inventory"
    repository_id = semaphoreui_project_repository.test.id
  }
}

resource "semaphoreui_project_environment" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Env-%[1]s"
}

resource "semaphoreui_project_template" "test" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Template-%[1]s"
  playbook       = "playbook.yml"
}
`, nameSuffix)
}

func testAccProjectIntegrationConfig(nameSuffix string, extras string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_integration" "test" {
  project_id  = semaphoreui_project.test.id
  template_id = semaphoreui_project_template.test.id
  name        = "Integration-%[2]s"
  %[3]s
}
`, testAccProjectIntegrationDependencyConfig(nameSuffix), nameSuffix, extras)
}

func testAccProjectIntegrationImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}
		return fmt.Sprintf("project/%[1]s/integration/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectIntegrationResource_basic(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with defaults (auth_method=none).
			{
				Config: testAccProjectIntegrationConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectIntegrationExists("semaphoreui_project_integration.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "name", fmt.Sprintf("Integration-%s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "auth_method", "none"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "auth_header", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "searchable", "false"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_integration.test", "auth_secret_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_integration.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_integration.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_integration.test", "template_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_integration.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectIntegrationImportID("semaphoreui_project_integration.test"),
			},
			// Update — switch on HMAC auth with a key reference.
			{
				Config: testAccProjectIntegrationConfig(nameSuffix, `
  auth_method    = "hmac"
  auth_secret_id = semaphoreui_project_key.test.id
  auth_header    = "X-Hub-Signature-256"
  searchable     = true
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectIntegrationExists("semaphoreui_project_integration.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "auth_method", "hmac"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "auth_header", "X-Hub-Signature-256"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "searchable", "true"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_integration.test", "auth_secret_id"),
				),
			},
			// Update — clear auth back to defaults.
			{
				Config: testAccProjectIntegrationConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectIntegrationExists("semaphoreui_project_integration.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "auth_method", "none"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "auth_header", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "searchable", "false"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_integration.test", "auth_secret_id"),
				),
			},
			// Delete
			{
				Config: testAccProjectIntegrationDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_integration.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectIntegrationResource_taskParams(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with task_params.ansible set.
			{
				Config: testAccProjectIntegrationConfig(nameSuffix, `
  task_params = {
    environment = "{}"
    ansible = {
      tags      = ["deploy"]
      skip_tags = ["slow"]
    }
  }
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectIntegrationExists("semaphoreui_project_integration.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "task_params.environment", "{}"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "task_params.ansible.tags.#", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "task_params.ansible.tags.0", "deploy"),
					resource.TestCheckResourceAttr("semaphoreui_project_integration.test", "task_params.ansible.skip_tags.0", "slow"),
				),
			},
			// Clear task_params.
			{
				Config: testAccProjectIntegrationConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectIntegrationExists("semaphoreui_project_integration.test"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_integration.test", "task_params"),
				),
			},
			// Delete
			{
				Config: testAccProjectIntegrationDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_integration.test"),
				),
			},
		},
	})
}
