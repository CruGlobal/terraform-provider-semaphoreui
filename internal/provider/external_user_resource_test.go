package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/user"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccExternalUserExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.Attributes["id"] == "" {
			return fmt.Errorf("no ID is set")
		}

		id, err := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		if err != nil {
			return err
		}

		_, err = testClient().User.GetUsersUserID(&user.GetUsersUserIDParams{UserID: id}, nil)
		return err
	}
}

// function to clean up external users since they are not deleted by the provider
func testAccExternalUserCleanup(s *terraform.State) error {
	// loop though each semaphoreui_external_user and ensure they are deleted
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "semaphoreui_external_user" {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		if err != nil {
			return err
		}

		_, _ = testClient().User.DeleteUsersUserID(&user.DeleteUsersUserIDParams{UserID: id}, nil)
	}
	return nil
}

func testAccExternalUserConfig(userNameSuffix string, userExtras string) string {
	return fmt.Sprintf(`
resource "semaphoreui_external_user" "test" {
  username = "test-%[1]s"
  name     = "Test User"
  email    = "test@example.com"
  %[2]s
}`, userNameSuffix, userExtras)
}

func testAccExternalUserConfig_Exists(userNameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_user" "existing" {
  username = "test2-%[1]s"
  name	   = "Test User"
  email	   = "test2@example.com"
  external = "true"
}

resource "semaphoreui_external_user" "test" {
  username       = "test2-%[1]s"
  name           = "Example User"
  email          = "test2@example.com"
  depends_on = [semaphoreui_user.existing]
}`, userNameSuffix)
}

func testAccExternalUserImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("user/%s", rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ExternalUserResource_basic(t *testing.T) {
	userNameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccExternalUserCleanup,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccExternalUserConfig(userNameSuffix, `  admin = true`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccExternalUserExists("semaphoreui_external_user.test"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "username", fmt.Sprintf("test-%s", userNameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "external", "true"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "name", "Test User"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "admin", "true"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttrSet("semaphoreui_external_user.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_external_user.test", "created"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_external_user.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccExternalUserImportID("semaphoreui_external_user.test"),
			},
			// Update and Read testing
			{
				Config: testAccExternalUserConfig(userNameSuffix, `  admin = false
alert = true`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "username", fmt.Sprintf("test-%s", userNameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "name", "Test User"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "admin", "false"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "alert", "true"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "external", "true"),
				),
			},
		},
	})
}

func TestAcc_ExternalUserResource_exists(t *testing.T) {
	userNameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccExternalUserCleanup,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccExternalUserConfig_Exists(userNameSuffix),
				// The `semaphoreui_user` resource will show a diff as the `name` attribute changes.
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccExternalUserExists("semaphoreui_external_user.test"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "username", fmt.Sprintf("test2-%s", userNameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "external", "true"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "name", "Example User"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "admin", "false"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("semaphoreui_external_user.test", "email", "test2@example.com"),
					resource.TestCheckResourceAttrSet("semaphoreui_external_user.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_external_user.test", "created"),
				),
			},
		},
	})
}
