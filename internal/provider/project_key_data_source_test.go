package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectKeyDataSourceConfigByID() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project 1"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None"
  none       = {}
}

data "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  id         = semaphoreui_project_key.test.id
  depends_on = [semaphoreui_project_key.test]
}`
}

func testAccProjectKeyDataSourceConfigByName() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project 1"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Password"
  login_password = {
    password = "hello123"
  }
}

data "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Password"
  depends_on = [semaphoreui_project_key.test]
}`
}

func TestAcc_ProjectKeyDataSource_basicID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectKeyDataSourceConfigByID(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_key.test", "name", "None"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_key.test", "none.%", "0"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_key.test", "login_password"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_key.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_key.test", "project_id"),
				),
			},
		},
	})
}

func TestAcc_ProjectKeyDataSource_basicName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectKeyDataSourceConfigByName(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_key.test", "name", "Password"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_key.test", "login_password.%", "2"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_key.test", "login_password.password", ""),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_key.test", "none"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_key.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_key.test", "project_id"),
				),
			},
		},
	})
}
