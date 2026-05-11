package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectProjectInventoryTofuWorkspaceConfig(nameSuffix string, workspace string) string {
	return testAccProjectInventoryConfig(nameSuffix, fmt.Sprintf(`
  tofu_workspace = {
    workspace = "%[1]s"
  }
`, workspace))
}

func TestAcc_ProjectInventoryResource_basicTofuWorkspace(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectProjectInventoryTofuWorkspaceConfig(nameSuffix, "name"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", ProjectInventoryTofuWorkspace),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "tofu_workspace.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "tofu_workspace.workspace", "name"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			{
				ResourceName:      "semaphoreui_project_inventory.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectInventoryImportID("semaphoreui_project_inventory.test"),
			},
			{
				Config: testAccProjectProjectInventoryTofuWorkspaceConfig(nameSuffix, "testing"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", ProjectInventoryTofuWorkspace),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "tofu_workspace.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "tofu_workspace.workspace", "testing"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			{
				Config: testAccProjectInventoryEmptyConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_inventory.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectInventoryResource_updateTypeToTofuWorkspace(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectProjectInventoryFileConfig(nameSuffix, "path/to/inventory"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", ProjectInventoryFile),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "file.%", "3"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "file.path", "path/to/inventory"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file.become_key_id"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file.repository_id"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "tofu_workspace"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			{
				Config: testAccProjectProjectInventoryTofuWorkspaceConfig(nameSuffix, "testing"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectInventoryExists("semaphoreui_project_inventory.test", ProjectInventoryTofuWorkspace),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "tofu_workspace.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_inventory.test", "tofu_workspace.workspace", "testing"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "static_yaml"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "file"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_inventory.test", "terraform_workspace"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_inventory.test", "ssh_key_id"),
				),
			},
			{
				Config: testAccProjectInventoryEmptyConfig(nameSuffix, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_inventory.test"),
				),
			},
		},
	})
}
