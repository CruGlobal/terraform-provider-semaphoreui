package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectRepositoryDataSourceConfigByID() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project 1"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test Repository"
  url        = "/path/to/repo"
  branch     = ""
  ssh_key_id = semaphoreui_project_key.test.id
}

data "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  id         = semaphoreui_project_repository.test.id
}`
}

func testAccProjectRepositoryDataSourceConfigByName() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project 1"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Semaphore"
  url        = "https://github.com/semaphoreui/semaphore.git"
  branch     = "develop"
  ssh_key_id = semaphoreui_project_key.test.id
}

data "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Semaphore"
  depends_on = [semaphoreui_project_repository.test]
}`
}

func TestAcc_ProjectRepositoryDataSource_basicID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectRepositoryDataSourceConfigByID(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_repository.test", "name", "Test Repository"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_repository.test", "url", "/path/to/repo"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_repository.test", "branch", ""),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_repository.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_repository.test", "project_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_repository.test", "ssh_key_id"),
				),
			},
		},
	})
}

func TestAcc_ProjectRepositoryDataSource_basicName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectRepositoryDataSourceConfigByName(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_repository.test", "name", "Semaphore"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_repository.test", "url", "https://github.com/semaphoreui/semaphore.git"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_repository.test", "branch", "develop"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_repository.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_repository.test", "project_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_repository.test", "ssh_key_id"),
				),
			},
		},
	})
}
