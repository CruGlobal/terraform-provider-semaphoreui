package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"regexp"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/user"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// function to clean up external users since they are not deleted by the provider
func testAccExternalUserCleanup(s *terraform.State) error {
	// loop though each semaphoreui_external_user and ensure they are deleted
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "data.semaphoreui_external_user" && rs.Type != "semaphoreui_external_user" {
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

func testAccExternalUserDataSourceConfigBasic(extras string) string {
	return fmt.Sprintf(`
data "semaphoreui_external_user" "test" {
  username = "username1"
  %s
}`, extras)
}

func testAccExternalUserDataSourceConfigExists(external bool, admin bool, extras string) string {
	return fmt.Sprintf(`
resource "semaphoreui_user" "test" {
  username = "username2"
  name = "Test User2"
  email = "test2@example.com"
  external = %[1]t
  admin = %[3]t
}

data "semaphoreui_external_user" "test" {
  username = "username2"
  %[2]s
  depends_on = [semaphoreui_user.test]
}`, external, extras, admin)
}

func TestAcc_ExternalUserDataSource_basicUsername(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccExternalUserCleanup,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExternalUserDataSourceConfigBasic(""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "username", "username1"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "name", "username1"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "email", "username1"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "admin", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "external", "true"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_external_user.test", "created"),
				),
			},
		},
	})
}

func TestAcc_ExternalUserDataSource_basicUsernameNameEmail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExternalUserDataSourceConfigBasic(`name = "Test Name"
email = "test@example.com"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "username", "username1"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "name", "Test Name"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "admin", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "external", "true"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_external_user.test", "created"),
				),
			},
		},
	})
}

func TestAcc_ExternalUserDataSource_existsError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccExternalUserCleanup,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config:      testAccExternalUserDataSourceConfigExists(false, false, ""),
				ExpectError: regexp.MustCompile("user with username (.*) is not an external user"),
			},
		},
	})
}

func TestAcc_ExternalUserDataSource_existsUsername(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccExternalUserCleanup,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExternalUserDataSourceConfigExists(true, false, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "username", "username2"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "name", "Test User2"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "email", "test2@example.com"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "admin", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "external", "true"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_external_user.test", "created"),
				),
			},
		},
	})
}

func TestAcc_ExternalUserDataSource_existsUsernameNameEmail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccExternalUserCleanup,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExternalUserDataSourceConfigExists(true, false, `name = "Test Name"
email = "test@example.com"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "username", "username2"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "name", "Test User2"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "email", "test2@example.com"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "admin", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "external", "true"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_external_user.test", "created"),
				),
			},
		},
	})
}

func TestAcc_ExternalUserDataSource_existsAdmin(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccExternalUserCleanup,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExternalUserDataSourceConfigExists(true, true, `name = "Test Name"
email = "test@example.com"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "username", "username2"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "name", "Test User2"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "email", "test2@example.com"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "admin", "true"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "alert", "false"),
					resource.TestCheckResourceAttr("data.semaphoreui_external_user.test", "external", "true"),
					resource.TestCheckResourceAttrSet("data.semaphoreui_external_user.test", "created"),
				),
			},
		},
	})
}
