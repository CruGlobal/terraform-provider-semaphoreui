package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectTemplateExists(resourceName string, templateType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.Attributes["id"] == "" {
			return fmt.Errorf("no ID is set")
		}
		if rs.Primary.Attributes["project_id"] == "" {
			return fmt.Errorf("no ProjectID is set")
		}

		id, _ := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		projectId, _ := strconv.ParseInt(rs.Primary.Attributes["project_id"], 10, 64)

		response, err := testClient().Project.GetProjectProjectIDTemplatesTemplateID(&project.GetProjectProjectIDTemplatesTemplateIDParams{
			ProjectID:  projectId,
			TemplateID: id,
		}, nil)
		if err != nil {
			return fmt.Errorf("error fetching project template: %s", err.Error())
		}

		if rs.Primary.Attributes["name"] != response.Payload.Name {
			return fmt.Errorf("template name mismatch: %s != %s", rs.Primary.Attributes["name"], response.Payload.Name)
		}
		if rs.Primary.Attributes["playbook"] != response.Payload.Playbook {
			return fmt.Errorf("template playbook mismatch: %s != %s", rs.Primary.Attributes["playbook"], response.Payload.Playbook)
		}
		if rs.Primary.Attributes["app"] != response.Payload.App {
			return fmt.Errorf("template app mismatch: %s != %s", rs.Primary.Attributes["app"], response.Payload.App)
		}
		if response.Payload.Type != templateType {
			return fmt.Errorf("template app mismatch: %s != %s", templateType, response.Payload.Type)
		}

		return nil
	}
}

func testAccProjectTemplateDependencyConfig(nameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None-%[1]s"
  none       = {}
}

resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Repo-%[1]s"
  url        = "git@github.com:example/test.git"
  branch     = "main"
  ssh_key_id = semaphoreui_project_key.test.id
}

resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Inventory-%[1]s"
  ssh_key_id = semaphoreui_project_key.test.id
  file = {
    path          = "path/to/inventory"
    repository_id = semaphoreui_project_repository.test.id
  }
}

resource "semaphoreui_project_environment" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Env-%[1]s"
}

resource "semaphoreui_project_view" "test" {
  project_id = semaphoreui_project.test.id
  title      = "Test View"
  position   = 0
}`, nameSuffix)
}

func testAccProjectTemplateConfig(nameSuffix string, extras string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_template" "test" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Test %[2]s"
  playbook	     = "playbook.yml"
  %[3]s
}`, testAccProjectTemplateDependencyConfig(nameSuffix), nameSuffix, extras)
}

func testAccProjectTemplateBuildConfig(nameSuffix string, startVersion bool, extras string) string {
	startVersionConfig := ""
	if startVersion {
		startVersionConfig = `start_version = "1.0.0"`
	}
	buildConfig := fmt.Sprintf(`
  build = {
    %[1]s
}
%[2]s
`, startVersionConfig, extras)
	return testAccProjectTemplateConfig(nameSuffix, buildConfig)
}

func testAccProjectTemplateDeployConfig(nameSuffix string, extras string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_template" "build" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Build %[2]s"
  playbook	     = "playbook.yml"
  build = {
    start_version = "2.0.0"
  }
}

resource "semaphoreui_project_template" "test" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Test %[2]s"
  playbook	     = "playbook.yml"
  deploy = {
    build_template_id = semaphoreui_project_template.build.id
  }
  %[3]s
}
`, testAccProjectTemplateDependencyConfig(nameSuffix), nameSuffix, extras)
}

func testAccProjectTemplateImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/template/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectTemplateResource_basic(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectTemplateConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "arguments"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "view_id"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectTemplateImportID("semaphoreui_project_template.test"),
			},
			// Update testing
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
allow_override_args_in_task = true
git_branch = "staging"
arguments = [
  "--help",
  "--verbose",
]
view_id = semaphoreui_project_view.test.id
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "true"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "git_branch", "staging"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.0", "--help"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.1", "--verbose"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "view_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectTemplateResource_basicBuild(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectTemplateBuildConfig(nameSuffix, false, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", "build"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "build.%", "1"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build.start_version"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "arguments"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectTemplateImportID("semaphoreui_project_template.test"),
			},
			// Update testing
			{
				Config: testAccProjectTemplateBuildConfig(nameSuffix, true, `
allow_override_args_in_task = true
git_branch = "staging"
arguments = [
  "--help",
  "--verbose",
]
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", "build"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "true"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "git_branch", "staging"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.0", "--help"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.1", "--verbose"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "build.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "build.start_version", "1.0.0"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectTemplateResource_basicDeploy(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectTemplateDeployConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", "deploy"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "deploy.%", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "deploy.build_template_id"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "deploy.autorun", "false"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "arguments"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectTemplateImportID("semaphoreui_project_template.test"),
			},
			// Update testing
			{
				Config: testAccProjectTemplateDeployConfig(nameSuffix, `
allow_override_args_in_task = true
git_branch = "staging"
arguments = [
  "--help",
  "--verbose",
]
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", "deploy"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "true"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "git_branch", "staging"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.0", "--help"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.1", "--verbose"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "deploy.%", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "deploy.build_template_id"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "deploy.autorun", "false"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectTemplateResource_surveyVars(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
survey_vars = [{
  name = "var1"
  title = "Variable 1"
  description = "Description 1"
  type = "string"
  required = true
}, {
  name = "var2"
  title = "Variable 2"
  type = "enum"
  enum_values = {
    "Option 1" = "opt1"
    "Option 2" = "opt2" 
  }
}]
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.name", "var1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.title", "Variable 1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.description", "Description 1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.type", "string"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.required", "true"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars.0.enum_values"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.name", "var2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.title", "Variable 2"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars.1.description"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.type", "enum"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.required", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.enum_values.%", "2"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "arguments"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectTemplateImportID("semaphoreui_project_template.test"),
			},
			// Update testing
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
allow_override_args_in_task = true
git_branch = "staging"
arguments = [
  "--help",
  "--verbose",
]
survey_vars = [{
  name = "var1"
  title = "Variable 1"
  description = "Description 1"
  type = "integer"
}, {
  name = "var2"
  title = "Variable 2"
  type = "secret"
  required = true
}]
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "true"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "git_branch", "staging"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.0", "--help"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.1", "--verbose"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.name", "var1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.title", "Variable 1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.description", "Description 1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.type", "integer"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.0.required", "false"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars.0.enum_values"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.name", "var2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.title", "Variable 2"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars.1.description"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.type", "secret"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "survey_vars.1.required", "true"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars.1.enum_values"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectTemplateResource_vaults(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
vaults = [{
  name = ""
  password = {
    vault_key_id = semaphoreui_project_key.test.id
  }
}, {
  name = "database"
  client_script = {
    script = "path/to/script-client.py"
  }
}]
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.0.name", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.0.password.%", "1"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "vaults.0.password.vault_key_id"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults.0.client_script"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.1.name", "database"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.1.client_script.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.1.client_script.script", "path/to/script-client.py"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults.1.password"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "arguments"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectTemplateImportID("semaphoreui_project_template.test"),
			},
			// Update testing
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
allow_override_args_in_task = true
git_branch = "staging"
arguments = [
  "--help",
  "--verbose",
]
vaults = [{
  name = "testing"
  password = {
    vault_key_id = semaphoreui_project_key.test.id
  }
}]
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "playbook", "playbook.yml"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "ansible"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "allow_override_args_in_task", "true"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "suppress_success_alerts", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "git_branch", "staging"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.0", "--help"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "arguments.1", "--verbose"),

					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.#", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.0.name", "testing"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "vaults.0.password.%", "1"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "vaults.0.password.vault_key_id"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "vaults.0.client_script"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "survey_vars"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "build"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "deploy"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "inventory_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "environment_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_template.test", "repository_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}
