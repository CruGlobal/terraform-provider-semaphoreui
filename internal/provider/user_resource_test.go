package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"regexp"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/user"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccUserExists(resourceName string) resource.TestCheckFunc {
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

func testAccUserConfig(userNameSuffix string, userExtras string) string {
	return fmt.Sprintf(`
resource "semaphoreui_user" "test" {
  username = "test-%[1]s"
  name     = "Test User"
  email    = "test@example.com"
  %[2]s
}`, userNameSuffix, userExtras)
}

func testAccUserConfig_Exists(userNameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_user" "existing" {
  username = "test-%[1]s"
  name	   = "Test User"
  email	   = "test@example.com"
}

resource "semaphoreui_user" "test" {
  username       = "test-%[1]s"
  name           = "Test User"
  email          = "test@example.com"
  depends_on = [semaphoreui_user.existing]
}`, userNameSuffix)
}

func testAccUserImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("user/%s", rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_UserResource_basic(t *testing.T) {
	userNameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccUserConfig(userNameSuffix, `  admin = true
password = "password!"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccUserExists("semaphoreui_user.test"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "username", fmt.Sprintf("test-%s", userNameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "external", "false"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "name", "Test User"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "password", "password!"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "admin", "true"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttrSet("semaphoreui_user.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_user.test", "created"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_user.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccUserImportID("semaphoreui_user.test"),
				// Password is encrypted and not returned by the API on import
				ImportStateVerifyIgnore: []string{"password"},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_user.test", "password", ""),
				),
			},
			// Update and Read testing
			{
				Config: testAccUserConfig(userNameSuffix, `  admin = false
password = "something"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_user.test", "username", fmt.Sprintf("test-%s", userNameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "name", "Test User"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "password", "something"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "admin", "false"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("semaphoreui_user.test", "external", "false"),
				),
			},
		},
	})
}

func TestAcc_UserResource_errorOnExists(t *testing.T) {
	userNameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,

		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config:      testAccUserConfig_Exists(userNameSuffix),
				ExpectError: regexp.MustCompile("Could not create user, unexpected error"),
			},
		},
	})
}
