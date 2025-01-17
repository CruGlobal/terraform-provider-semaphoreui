package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectUserDataSourceConfig() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project 1"
}

resource "semaphoreui_user" "test" {
  username = "test"
  name = "test name"
  email = "test@example.com"
}

resource "semaphoreui_project_user" "test" {
  project_id = semaphoreui_project.test.id
  user_id = semaphoreui_user.test.id
  role = "task_runner"
}

data "semaphoreui_project_user" "test" {
  project_id = semaphoreui_project.test.id
  user_id = semaphoreui_user.test.id
  depends_on = [semaphoreui_project_user.test]
}`
}

func TestAcc_ProjectUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccProjectUserDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_user.test", "username", "test"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_user.test", "role", "task_runner"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_user.test", "name", "test name"),
				),
			},
		},
	})
}
