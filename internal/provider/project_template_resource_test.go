package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"net/http"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/template"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
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

		response, err := testClient().Template.GetProjectProjectIDTemplatesTemplateID(&template.GetProjectProjectIDTemplatesTemplateIDParams{
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

// testAccProjectTemplateServerTaskParam checks one key of the template's
// task_params as SemaphoreUI stores it. The runner and the web UI read a
// template's limit, tags, skip_tags and auto_approve from the top level of
// task_params, so the check reads the raw API response instead of the
// generated client, whose model is what is under test.
func testAccProjectTemplateServerTaskParam(resourceName string, key string, want any) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		url := fmt.Sprintf("%s://%s:%s/api/project/%s/templates/%s",
			testProtocol(), testHostname(), testPort(),
			rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"])
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+testApiToken())
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("error fetching project template: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error fetching project template: HTTP %d", resp.StatusCode)
		}

		var tpl struct {
			TaskParams map[string]any `json:"task_params"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&tpl); err != nil {
			return fmt.Errorf("error decoding project template: %s", err.Error())
		}

		got, _ := json.Marshal(tpl.TaskParams[key])
		expected, _ := json.Marshal(want)
		if string(got) != string(expected) {
			return fmt.Errorf("server task_params.%s = %s, want %s", key, got, expected)
		}
		return nil
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

func TestAcc_ProjectTemplateResource_taskParams(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with Ansible task_params — tags, skip_tags, diff.
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
  task_params = {
    arguments = "[\"-v\"]"
    ansible = {
      tags      = ["deploy", "db"]
      skip_tags = ["slow"]
      limit     = ["web"]
      diff      = true
    }
  }
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "tags", []string{"deploy", "db"}),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "skip_tags", []string{"slow"}),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "limit", []string{"web"}),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.limit.#", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.limit.0", "web"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.arguments", "[\"-v\"]"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.tags.#", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.tags.0", "deploy"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.tags.1", "db"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.skip_tags.#", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.skip_tags.0", "slow"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.diff", "true"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "task_params.terraform"),
				),
			},
			// ImportState
			{
				ResourceName:      "semaphoreui_project_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectTemplateImportID("semaphoreui_project_template.test"),
			},
			// Update tags (rotate).
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
  task_params = {
    ansible = {
      tags = ["only-this"]
    }
  }
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "tags", []string{"only-this"}),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "skip_tags", nil),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "limit", nil),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.tags.#", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.tags.0", "only-this"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "task_params.ansible.skip_tags"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.ansible.diff", "false"),
				),
			},
			// Clear task_params entirely.
			{
				Config: testAccProjectTemplateConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "task_params"),
				),
			},
			// Delete
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectTemplateResource_taskParamsTerraform(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create a Terraform app template with auto_approve.
			{
				Config: testAccProjectTemplateConfig(nameSuffix, `
  app = "terraform"
  task_params = {
    terraform = {
      auto_approve = true
      upgrade      = true
    }
  }
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "app", "terraform"),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "auto_approve", true),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.terraform.auto_approve", "true"),
					resource.TestCheckResourceAttr("semaphoreui_project_template.test", "task_params.terraform.upgrade", "true"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_template.test", "task_params.ansible"),
				),
			},
			// Delete
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}

// testAccProjectTemplateWriteNestedTaskParams overwrites a template's
// task_params with the nested `params` shape that provider v1.5.0 wrote, so a
// test can check that the next plan moves the values to the top level.
func testAccProjectTemplateWriteNestedTaskParams(t *testing.T, projectID *string, templateID *string, nested map[string]any) func() {
	return func() {
		url := fmt.Sprintf("%s://%s:%s/api/project/%s/templates/%s",
			testProtocol(), testHostname(), testPort(), *projectID, *templateID)

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+testApiToken())
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			_ = resp.Body.Close()
			t.Fatalf("error fetching project template: HTTP %d", resp.StatusCode)
		}
		var tpl map[string]any
		err = json.NewDecoder(resp.Body).Decode(&tpl)
		_ = resp.Body.Close()
		if err != nil {
			t.Fatal(err)
		}

		tpl["task_params"] = map[string]any{"params": nested}
		body, err := json.Marshal(tpl)
		if err != nil {
			t.Fatal(err)
		}
		req, err = http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+testApiToken())
		req.Header.Set("Content-Type", "application/json")
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
			t.Fatalf("error writing nested task_params: HTTP %d", resp.StatusCode)
		}
	}
}

func TestAcc_ProjectTemplateResource_taskParamsNestedUpgrade(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	var projectID, templateID string
	config := testAccProjectTemplateConfig(nameSuffix, `
  task_params = {
    ansible = {
      tags  = ["users"]
      limit = ["fivetran"]
    }
  }
`)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateExists("semaphoreui_project_template.test", ""),
					func(s *terraform.State) error {
						rs := s.RootModule().Resources["semaphoreui_project_template.test"]
						projectID = rs.Primary.Attributes["project_id"]
						templateID = rs.Primary.Attributes["id"]
						return nil
					},
				),
			},
			// A template written in the old nested shape must plan an update,
			// and the apply must leave the values where SemaphoreUI reads them.
			{
				PreConfig: testAccProjectTemplateWriteNestedTaskParams(t, &projectID, &templateID, map[string]any{
					"tags":  []string{"users"},
					"limit": []string{"fivetran"},
				}),
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("semaphoreui_project_template.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "tags", []string{"users"}),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "limit", []string{"fivetran"}),
					testAccProjectTemplateServerTaskParam("semaphoreui_project_template.test", "params", nil),
				),
			},
			{
				Config: testAccProjectTemplateDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_template.test"),
				),
			},
		},
	})
}
