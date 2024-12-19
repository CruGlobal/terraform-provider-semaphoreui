package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectDataSourceConfig() string {
	return `
resource "semaphoreui_project" "project" {
  name       = "Project 1"
  alert      = true
  alert_chat = "slack"
}
data "semaphoreui_project" "test" {
  id = semaphoreui_project.project.id
}`
}

func TestAcc_ProjectDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccProjectDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "name", "Project 1"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "alert", "true"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "alert_chat", "slack"),
					resource.TestCheckResourceAttr("data.semaphoreui_project.test", "max_parallel_tasks", "0"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project.test", "created"),
				),
			},
		},
	})
}
