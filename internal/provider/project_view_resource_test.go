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

func testAccProjectViewExists(resourceName string) resource.TestCheckFunc {
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

		response, err := testClient().Project.GetProjectProjectIDViewsViewID(&project.GetProjectProjectIDViewsViewIDParams{
			ProjectID: projectId,
			ViewID:    id,
		}, nil)
		if err != nil {
			return fmt.Errorf("error reading project view: %s", err.Error())
		}

		if response.Payload.Title != rs.Primary.Attributes["title"] {
			return fmt.Errorf("view title mismatch: %s != %s", response.Payload.Title, rs.Primary.Attributes["title"])
		}

		return nil
	}
}

func testAccProjectViewConfig(title string, position int64) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-project"
}

resource "semaphoreui_project_view" "test" {
  project_id = semaphoreui_project.test.id
  title      = "%s"
  position   = %d
}
`, title, position)
}

func testAccProjectViewImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/view/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectViewResource_basic(t *testing.T) {
	title := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectViewConfig(title, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectViewExists("semaphoreui_project_view.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_view.test", "title", title),
					resource.TestCheckResourceAttr("semaphoreui_project_view.test", "position", "1"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_view.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_view.test", "project_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_view.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectViewImportID("semaphoreui_project_view.test"),
			},
			// Update testing
			{
				Config: testAccProjectViewConfig(title, 5),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectViewExists("semaphoreui_project_view.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_view.test", "title", title),
					resource.TestCheckResourceAttr("semaphoreui_project_view.test", "position", "5"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_view.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_view.test", "project_id"),
				),
			},
			// Delete testing
			{
				Config: `
resource "semaphoreui_project" "test" {
  name = "test-project"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_view.test"),
				),
			},
		},
	})
}
