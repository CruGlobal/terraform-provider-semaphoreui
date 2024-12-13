package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectsDataSourceConfig() string {
	return `
resource "semaphoreui_project" "project1" {
  name = "Project 1"
}
resource "semaphoreui_project" "project2" {
  name  = "Project 2"
  alert = true
}
data "semaphoreui_projects" "test" {
  depends_on = [semaphoreui_project.project1]
}`
}

func TestAcc_ProjectsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccProjectsDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_projects.test", "projects.#", "2"),
					resource.TestCheckResourceAttr("data.semaphoreui_projects.test", "projects.0.name", "Project 1"),
					resource.TestCheckResourceAttr("data.semaphoreui_projects.test", "projects.1.name", "Project 2"),
					resource.TestCheckResourceAttr("data.semaphoreui_projects.test", "projects.1.alert", "true"),
				),
			},
		},
	})
}
