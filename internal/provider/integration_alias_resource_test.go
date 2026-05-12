package provider

import (
	"fmt"
	"regexp"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/integration"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccIntegrationAliasExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["id"] == "" {
			return fmt.Errorf("no ID is set")
		}

		id, _ := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		projectID, _ := strconv.ParseInt(rs.Primary.Attributes["project_id"], 10, 64)
		integrationIDStr := rs.Primary.Attributes["integration_id"]

		if integrationIDStr != "" {
			integrationID, _ := strconv.ParseInt(integrationIDStr, 10, 64)
			response, err := testClient().Integration.GetProjectProjectIDIntegrationsIntegrationIDAliases(
				&integration.GetProjectProjectIDIntegrationsIntegrationIDAliasesParams{
					ProjectID:     projectID,
					IntegrationID: integrationID,
				}, nil)
			if err != nil {
				return fmt.Errorf("could not list integration aliases: %s", err.Error())
			}
			for _, a := range response.Payload {
				if a.ID == id {
					return nil
				}
			}
			return fmt.Errorf("integration alias %d not found under integration %d", id, integrationID)
		}

		response, err := testClient().Integration.GetProjectProjectIDIntegrationsAliases(
			&integration.GetProjectProjectIDIntegrationsAliasesParams{ProjectID: projectID}, nil)
		if err != nil {
			return fmt.Errorf("could not list project aliases: %s", err.Error())
		}
		for _, a := range response.Payload {
			if a.ID == id {
				return nil
			}
		}
		return fmt.Errorf("project alias %d not found", id)
	}
}

func testAccIntegrationAliasDependencyConfig(nameSuffix string) string {
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

resource "semaphoreui_project_template" "test" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Template-%[1]s"
  playbook       = "playbook.yml"
}

resource "semaphoreui_project_integration" "test" {
  project_id  = semaphoreui_project.test.id
  template_id = semaphoreui_project_template.test.id
  name        = "Integration-%[1]s"
}
`, nameSuffix)
}

func testAccIntegrationAliasImportIDScoped(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}
		return fmt.Sprintf("project/%[1]s/integration/%[2]s/alias/%[3]s",
			rs.Primary.Attributes["project_id"],
			rs.Primary.Attributes["integration_id"],
			rs.Primary.Attributes["id"]), nil
	}
}

func testAccIntegrationAliasImportIDProject(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}
		return fmt.Sprintf("project/%[1]s/alias/%[2]s",
			rs.Primary.Attributes["project_id"],
			rs.Primary.Attributes["id"]), nil
	}
}

// Integration-scoped alias: create, read URL, import, delete.
func TestAcc_IntegrationAliasResource_integrationScoped(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationAliasDependencyConfig(nameSuffix) + `
resource "semaphoreui_integration_alias" "test" {
  project_id     = semaphoreui_project.test.id
  integration_id = semaphoreui_project_integration.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIntegrationAliasExists("semaphoreui_integration_alias.test"),
					resource.TestCheckResourceAttrSet("semaphoreui_integration_alias.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_integration_alias.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_integration_alias.test", "integration_id"),
					resource.TestMatchResourceAttr("semaphoreui_integration_alias.test", "url",
						regexp.MustCompile(`^https?://[^/]+/api/integrations/[a-z0-9]+$`)),
				),
			},
			{
				ResourceName:      "semaphoreui_integration_alias.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccIntegrationAliasImportIDScoped("semaphoreui_integration_alias.test"),
			},
			// Delete
			{
				Config: testAccIntegrationAliasDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_integration_alias.test"),
				),
			},
		},
	})
}

// Project-scoped alias: no integration_id, separate API endpoint.
func TestAcc_IntegrationAliasResource_projectScoped(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationAliasDependencyConfig(nameSuffix) + `
resource "semaphoreui_integration_alias" "test" {
  project_id = semaphoreui_project.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIntegrationAliasExists("semaphoreui_integration_alias.test"),
					resource.TestCheckResourceAttrSet("semaphoreui_integration_alias.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_integration_alias.test", "project_id"),
					resource.TestCheckNoResourceAttr("semaphoreui_integration_alias.test", "integration_id"),
					resource.TestMatchResourceAttr("semaphoreui_integration_alias.test", "url",
						regexp.MustCompile(`^https?://[^/]+/api/integrations/[a-z0-9]+$`)),
				),
			},
			{
				ResourceName:      "semaphoreui_integration_alias.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccIntegrationAliasImportIDProject("semaphoreui_integration_alias.test"),
			},
			// Delete
			{
				Config: testAccIntegrationAliasDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_integration_alias.test"),
				),
			},
		},
	})
}

// Adding integration_id to an existing project-scoped alias forces
// replacement — confirms the ForceNew plan modifier.
func TestAcc_IntegrationAliasResource_scopeSwitchForcesReplace(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationAliasDependencyConfig(nameSuffix) + `
resource "semaphoreui_integration_alias" "test" {
  project_id = semaphoreui_project.test.id
}
`,
				Check: testAccIntegrationAliasExists("semaphoreui_integration_alias.test"),
			},
			{
				Config: testAccIntegrationAliasDependencyConfig(nameSuffix) + `
resource "semaphoreui_integration_alias" "test" {
  project_id     = semaphoreui_project.test.id
  integration_id = semaphoreui_project_integration.test.id
}
`,
				Check: testAccIntegrationAliasExists("semaphoreui_integration_alias.test"),
			},
		},
	})
}
