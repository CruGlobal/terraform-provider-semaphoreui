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

func testAccProjectExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		_, err = testClient().Project.GetProjectProjectID(&project.GetProjectProjectIDParams{ProjectID: id}, nil)
		return err
	}
}

func TestAcc_ProjectResource(t *testing.T) {
	projectNameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectConfig(projectNameSuffix, `alert = false
max_parallel_tasks = 0`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectExists("semaphoreui_project.test"),
					resource.TestCheckResourceAttr("semaphoreui_project.test", "name", fmt.Sprintf("test-%s", projectNameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project.test", "alert", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project.test", "max_parallel_tasks", "0"),
					resource.TestCheckResourceAttrSet("semaphoreui_project.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project.test", "created"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project.test",
				ImportState:       true,
				ImportStateVerify: true,
				// API returns different timestamp format between create and read, so just ignore it
				ImportStateVerifyIgnore: []string{"created"},
				ImportStateIdFunc:       getProjectImportID("semaphoreui_project.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectConfig(projectNameSuffix, `alert = true
max_parallel_tasks = 2
alert_chat = "testing"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_project.test", "name", fmt.Sprintf("test-%s", projectNameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project.test", "alert", "true"),
					resource.TestCheckResourceAttr("semaphoreui_project.test", "alert_chat", "testing"),
					resource.TestCheckResourceAttr("semaphoreui_project.test", "max_parallel_tasks", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func getProjectImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%s", rs.Primary.Attributes["id"]), nil
	}
}

func testAccProjectConfig(projectNameSuffix string, projectExtras string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
  %[2]s
}`, projectNameSuffix, projectExtras)
}
