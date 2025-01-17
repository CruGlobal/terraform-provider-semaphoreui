package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectScheduleDataSourceConfigByID() string {
	return `
resource "semaphoreui_project" "test" {
  name = "Project"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Repo"
  url        = "git@github.com:example/test.git"
  branch     = "main"
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
  name       = "Environment"
  secrets = [{
    name  = "SECRET_ONE"
    type  = "var"
    value = "VALUE_ONE"
  }]
}

# Task Template
resource "semaphoreui_project_template" "test" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Template"
  playbook       = "playbook.yml"
  description    = "Description"
}

resource "semaphoreui_project_schedule" "test" {
  project_id  = semaphoreui_project.test.id
  template_id = semaphoreui_project_template.test.id
  name        = "Test Schedule"
  cron_format = "0 0 * * *"
  enabled     = true
}

data "semaphoreui_project_schedule" "test" {
  project_id = semaphoreui_project.test.id
  id         = semaphoreui_project_schedule.test.id
}`
}

func TestAcc_ProjectScheduleDataSource_basicID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectScheduleDataSourceConfigByID(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_schedule.test", "name", "Test Schedule"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_schedule.test", "cron_format", "0 0 * * *"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_schedule.test", "enabled", "true"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_schedule.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_schedule.test", "project_id"),
				),
			},
		},
	})
}
