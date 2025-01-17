package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectDataSourceConfigByID() string {
	return `
resource "semaphoreui_project" "test" {
  name       = "Project 1"
  alert      = true
  alert_chat = "slack"
}

data "semaphoreui_project" "test" {
  id = semaphoreui_project.test.id
}`
}

func testAccProjectDataSourceConfigByName() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Test Project"
}

data "semaphoreui_project" "test" {
  name = semaphoreui_project.test.name
}`
}

func TestAcc_ProjectDataSource_basicID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccProjectDataSourceConfigByID(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "name", "Project 1"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "alert", "true"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "alert_chat", "slack"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "max_parallel_tasks", "0"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project.test", "created"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project.test", "id"),
				),
			},
		},
	})
}

func TestAcc_ProjectDataSource_basicName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccProjectDataSourceConfigByName(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "name", "Test Project"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "alert", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "max_parallel_tasks", "0"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project.test", "created"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project.test", "id"),
				),
			},
		},
	})
}
