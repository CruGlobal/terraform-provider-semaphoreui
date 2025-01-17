package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectInventoryDataSourceConfigByID() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project 1"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test Inventory"
  ssh_key_id = semaphoreui_project_key.test.id
  static = {
	inventory = <<-EOT
      [all]
      hostname
    EOT
  }
}

data "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  id         = semaphoreui_project_inventory.test.id
  depends_on = [semaphoreui_project_inventory.test]
}`
}

func testAccProjectInventoryDataSourceConfigByName() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project 1"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test Inventory"
  ssh_key_id = semaphoreui_project_key.test.id
  file = {
    path = "inventory.yml"
  }
}

data "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test Inventory"
  depends_on = [semaphoreui_project_inventory.test]
}`
}

func TestAcc_ProjectInventoryDataSource_basicID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectInventoryDataSourceConfigByID(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_inventory.test", "name", "Test Inventory"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_inventory.test", "ssh_key_id"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_inventory.test", "static.%", "2"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_inventory.test", "static.inventory", "[all]\nhostname\n"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_inventory.test", "terraform_workspace"),
				),
			},
		},
	})
}

func TestAcc_ProjectInventoryDataSource_basicName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectInventoryDataSourceConfigByName(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_inventory.test", "name", "Test Inventory"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_inventory.test", "ssh_key_id"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_inventory.test", "file.%", "3"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_inventory.test", "file.path", "inventory.yml"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_inventory.test", "terraform_workspace"),
				),
			},
		},
	})
}
