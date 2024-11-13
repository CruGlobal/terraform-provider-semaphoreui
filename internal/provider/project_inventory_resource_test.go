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

func testAccProjectInventoryExists(resourceName string, inventoryType string) resource.TestCheckFunc {
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

		response, err := testClient().Project.GetProjectProjectIDInventoryInventoryID(&project.GetProjectProjectIDInventoryInventoryIDParams{
			ProjectID:   projectId,
			InventoryID: id,
		}, nil)
		if err != nil {
			return fmt.Errorf("error reading project inventory: %s", err.Error())
		}

		if response.Payload.Type != inventoryType {
			return fmt.Errorf("inventory type mismatch: %s != %s", response.Payload.Type, inventoryType)
		}

		return nil
	}
}

func testAccProjectInventoryEmptyConfig(nameSuffix string, extras string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}
resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "test-%[1]s"
  none = {}
}
%[2]s
`, nameSuffix, extras)
}

func testAccProjectInventoryConfig(nameSuffix string, extras string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test %[2]s"
  ssh_key_id = semaphoreui_project_key.test.id
  %[3]s
}`, testAccProjectInventoryEmptyConfig(nameSuffix, ""), nameSuffix, extras)
}

func testAccProjectInventoryStaticConfig(nameSuffix string, become bool) string {
	return testAccProjectInventoryConfig(nameSuffix, fmt.Sprintf(`
  static = {
    inventory = <<-EOT
      [webservers]
      foo.example.com
      bar.example.com
    EOT
	become_key_id = %t ? semaphoreui_project_key.test.id : null
  }
`, become))
}

func testAccProjectInventoryStaticYamlConfig(nameSuffix string, become bool) string {
	return testAccProjectInventoryConfig(nameSuffix, fmt.Sprintf(`
  static_yaml = {
    inventory = yamlencode({
      webservers: {
        "foo.example.com": {}
        "bar.example.com": {}
      }
    })
    become_key_id = %t ? semaphoreui_project_key.test.id : null
  }
`, become))
}

func testAccProjectInventoryFileConfig(nameSuffix string, path string) string {
	return testAccProjectInventoryConfig(nameSuffix, fmt.Sprintf(`
  file = {
    path = "%[1]s"
  }
`, path))
}

func testAccProjectInventoryFileWithRepoConfig(nameSuffix string, path string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test %[2]s"
  ssh_key_id = semaphoreui_project_key.test.id
  file = {
    path          = "%[3]s"
    repository_id = semaphoreui_project_repository.test.id
  }
}`, testAccProjectInventoryEmptyConfig(nameSuffix, fmt.Sprintf(`
resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "test-%[1]s"
  url        = "git@github.com:example/test.git"
  branch     = "main"
  ssh_key_id = semaphoreui_project_key.test.id
}
`, nameSuffix)), nameSuffix, path)
}

func testAccProjectInventoryTerraformWorkspaceConfig(nameSuffix string, workspace string) string {
	return testAccProjectInventoryConfig(nameSuffix, fmt.Sprintf(`
  terraform_workspace = {
    workspace = "%[1]s"
  }
`, workspace))
}

func testAccProjectInventoryImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/inventory/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectInventoryResource_basicFile(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectInventoryFileConfig(nameSuffix, "path/to/inventory"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryFile),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "file.%", "3"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "file.path", "path/to/inventory"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file.become_key_id"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file.repository_id"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_inventory.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectInventoryImportID("semaphoreui_project_inventory.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectInventoryFileWithRepoConfig(nameSuffix, "path/to/inventory"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryFile),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "file.%", "3"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "file.path", "path/to/inventory"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file.become_key_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "file.repository_id"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectInventoryEmptyConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_inventory.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectInventoryResource_basicStatic(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectInventoryStaticConfig(nameSuffix, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryStatic),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "static.%", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "static.inventory"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static.become_key_id"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_inventory.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectInventoryImportID("semaphoreui_project_inventory.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectInventoryStaticConfig(nameSuffix, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryStatic),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "static.%", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "static.inventory"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "static.become_key_id"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectInventoryEmptyConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_inventory.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectInventoryResource_basicStaticYaml(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectInventoryStaticYamlConfig(nameSuffix, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryStaticYaml),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "static_yaml.%", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "static_yaml.inventory"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml.become_key_id"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_inventory.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectInventoryImportID("semaphoreui_project_inventory.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectInventoryStaticYamlConfig(nameSuffix, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryStaticYaml),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "static_yaml.%", "2"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "static_yaml.inventory"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "static_yaml.become_key_id"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectInventoryEmptyConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_inventory.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectInventoryResource_basicTerraformWorkspace(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectInventoryTerraformWorkspaceConfig(nameSuffix, "name"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryTerraformWorkspace),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace.workspace", "name"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_inventory.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectInventoryImportID("semaphoreui_project_inventory.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectInventoryTerraformWorkspaceConfig(nameSuffix, "testing"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", InventoryTerraformWorkspace),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),

					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace.workspace", "testing"),

					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),

					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectInventoryEmptyConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_inventory.test"),
				),
			},
		},
	})
}
