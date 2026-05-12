package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectIntegrationDataSourceFixture() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Integration DS Project"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Repo"
  url        = "/path/to/repo"
  branch     = ""
  ssh_key_id = semaphoreui_project_key.test.id
}

resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Inventory"
  ssh_key_id = semaphoreui_project_key.test.id
  file = {
    path          = "path/to/inventory"
    repository_id = semaphoreui_project_repository.test.id
  }
}

resource "semaphoreui_project_environment" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Env"
}

resource "semaphoreui_project_template" "test" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Template"
  playbook       = "playbook.yml"
}

resource "semaphoreui_project_integration" "test" {
  project_id  = semaphoreui_project.test.id
  template_id = semaphoreui_project_template.test.id
  name        = "github-webhook"
  auth_method = "github"
  auth_header = "X-Hub-Signature-256"
}
`
}

func TestAcc_ProjectIntegrationDataSource_basicID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectIntegrationDataSourceFixture() + `
data "semaphoreui_project_integration" "test" {
  project_id = semaphoreui_project.test.id
  id         = semaphoreui_project_integration.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_integration.test", "name", "github-webhook"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_integration.test", "auth_method", "github"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_integration.test", "auth_header", "X-Hub-Signature-256"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_integration.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_integration.test", "project_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_integration.test", "template_id"),
				),
			},
		},
	})
}

func TestAcc_ProjectIntegrationDataSource_basicName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectIntegrationDataSourceFixture() + `
data "semaphoreui_project_integration" "test" {
  project_id = semaphoreui_project.test.id
  name       = "github-webhook"
  depends_on = [semaphoreui_project_integration.test]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_integration.test", "name", "github-webhook"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_integration.test", "auth_method", "github"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_integration.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_integration.test", "project_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_integration.test", "template_id"),
				),
			},
		},
	})
}
