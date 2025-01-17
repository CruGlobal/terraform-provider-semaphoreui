package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectTemplateDataSourceConfigByID() string {
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
  name           = "Test Template"
  playbook       = "playbook.yml"
  description    = "Description"
  arguments = [
    "--help",
    "--vvv",
  ]
  allow_override_args_in_task = true

  survey_vars = [{
    name     = "age"
    title    = "What is your age?"
    required = true
    type     = "integer"
    }, {
    name  = "question"
    title = "Pick one."
    type  = "enum"
    enum_values = {
      "First Value"  = "1"
      "Second Value" = "2"
    }
  }]

  vaults = [{
    name = "" # default vault
    password = {
      vault_key_id = semaphoreui_project_key.test.id
    }
    }, {
    name = "database"
    script = {
      script = "path/to/script-client.py"
    }
  }]
}

data "semaphoreui_project_template" "test" {
  project_id = semaphoreui_project.test.id
  id         = semaphoreui_project_template.test.id
}`
}

func testAccProjectTemplateDataSourceConfigByName() string {
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
  name           = "Build Template"
  playbook       = "playbook.yml"
  description    = "Description"
  build = {
    start_version = "1.0.0"
  }
}

data "semaphoreui_project_template" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Build Template"
  depends_on = [semaphoreui_project_template.test]
}`
}

func TestAcc_ProjectTemplateDataSource_basicID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTemplateDataSourceConfigByID(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "name", "Test Template"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "description", "Description"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "arguments.#", "2"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "survey_vars.#", "2"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "vaults.#", "2"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_template.test", "build"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_template.test", "deploy"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "repository_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "inventory_id"),
				),
			},
		},
	})
}

func TestAcc_ProjectTemplateDataSource_basicName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTemplateDataSourceConfigByName(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "name", "Build Template"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "description", "Description"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_template.test", "arguments"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_template.test", "vaults"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "build.%", "1"),
					resource.TestCheckResourceAttr("data.semaphoreui_project_template.test", "build.start_version", "1.0.0"),
					resource.TestCheckNoResourceAttr("data.semaphoreui_project_template.test", "deploy"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "repository_id"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_project_template.test", "inventory_id"),
				),
			},
		},
	})
}
