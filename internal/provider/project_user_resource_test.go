package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectUserConfig(nameSuffix string, role string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}

resource "semaphoreui_user" "test" {
  username = "test"
  name = "test-%[1]s"
  email = "test@example.com"
}

resource "semaphoreui_project_user" "test_test" {
  project_id = semaphoreui_project.test.id
  user_id = semaphoreui_user.test.id
  role = "%[2]s"
}`, nameSuffix, role)
}

func testAccProjectUserImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%s/user/%s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["user_id"]), nil
	}
}

func TestAcc_ProjectUserResource(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectUserConfig(nameSuffix, "guest"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_project_user.test_test", "name", fmt.Sprintf("test-%s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_user.test_test", "username", "test"),
					resource.TestCheckResourceAttr("semaphoreui_project_user.test_test", "role", "guest"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_user.test_test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_user.test_test", "user_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_user.test_test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectUserImportID("semaphoreui_project_user.test_test"),
				// Previous terraform provider SDKs required an ID attribute field, the provider framework does not.
				// We use a combination of a project_id and user_id to uniquely identify a project user,
				// but testing framework requires a single field. We picked "user_id", but testing will verify both ids
				ImportStateVerifyIdentifierAttribute: "user_id",
			},
			// Update and Read testing
			{
				Config: testAccProjectUserConfig(nameSuffix, "manager"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_project_user.test_test", "name", fmt.Sprintf("test-%s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_user.test_test", "username", "test"),
					resource.TestCheckResourceAttr("semaphoreui_project_user.test_test", "role", "manager"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_user.test_test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_user.test_test", "user_id"),
				),
			},
		},
	})
}
