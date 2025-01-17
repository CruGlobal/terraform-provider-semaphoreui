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

func testAccProjectKeyExists(resourceName string, keyType string) resource.TestCheckFunc {
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

		response, err := testClient().Project.GetProjectProjectIDKeys(&project.GetProjectProjectIDKeysParams{
			ProjectID: projectId,
		}, nil)
		if err != nil {
			return fmt.Errorf("error fetching project keys: %s", err.Error())
		}

		for _, key := range response.Payload {
			if key.ID == id {
				if key.Type == keyType {
					return nil
				}
				return fmt.Errorf("key type mismatch: %s != %s", key.Type, keyType)
			}
		}
		return fmt.Errorf("project key not found: %d", id)
	}
}

func testAccProjectKeyEmptyConfig(nameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}
`, nameSuffix)
}

func testAccProjectKeyConfig(nameSuffix string, keyExtras string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test %[2]s"
  %[3]s
}`, testAccProjectKeyEmptyConfig(nameSuffix), nameSuffix, keyExtras)
}

func testAccProjectKeyNoneConfig(nameSuffix string) string {
	return testAccProjectKeyConfig(nameSuffix, `none = {}`)
}

func testAccProjectKeyLoginPasswordConfig(nameSuffix string, login string, password string) string {
	return testAccProjectKeyConfig(nameSuffix, fmt.Sprintf(`login_password = {
  login    = "%[1]s"
  password = "%[2]s"
}`, login, password))
}

func testAccProjectKeySSHConfig(nameSuffix string, login string, privateKey string, passphrase string) string {
	return testAccProjectKeyConfig(nameSuffix, fmt.Sprintf(`ssh = {
  login       = "%[1]s"
  private_key = <<-EOT
%[2]sEOT
  passphrase = "%[3]s"
}`, login, privateKey, passphrase))
}

func testAccProjectKeyImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/key/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectKeyResource_basicNone(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectKeyNoneConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeNone),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "none.%", "0"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "login_password"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_key.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectKeyImportID("semaphoreui_project_key.test"),
			},
			// Delete testing
			{
				Config: testAccProjectKeyEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_key.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectKeyResource_basicLoginPassword(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectKeyLoginPasswordConfig(nameSuffix, "username", "password"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeLoginPassword),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "none"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.%", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.login", "username"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.password", "password"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_key.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectKeyImportID("semaphoreui_project_key.test"),
				// API doesn't return login_password details, required attributes are imported as empty strings
				ImportStateVerifyIgnore: []string{"login_password"},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.password", ""),
				),
			},
			// Update and Read testing
			{
				Config: testAccProjectKeyLoginPasswordConfig(nameSuffix, "foo", "bar"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeLoginPassword),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "none"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.%", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.login", "foo"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.password", "bar"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectKeyEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_key.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectKeyResource_basicSSH(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	_, privateKey, _ := acctest.RandSSHKeyPair("")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectKeySSHConfig(nameSuffix, "username", privateKey, "passphrase"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeSSH),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "none"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "login_password"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.%", "3"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.login", "username"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.private_key", privateKey),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.passphrase", "passphrase"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_key.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectKeyImportID("semaphoreui_project_key.test"),
				// API doesn't return ssh details, required attributes are imported as empty strings
				ImportStateVerifyIgnore: []string{"ssh"},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.%", "1"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.private_key", ""),
				),
			},
			// Update and Read testing
			{
				Config: testAccProjectKeySSHConfig(nameSuffix, "testing", privateKey, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeSSH),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "none"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "login_password"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.%", "3"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.login", "testing"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.private_key", privateKey),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.passphrase", ""),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectKeyEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_key.test"),
				),
			},
		},
	})
}

func TestAcc_ProjectKeyResource_changeName(t *testing.T) {
	nameSuffix1 := acctest.RandString(8)
	nameSuffix2 := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectKeyNoneConfig(nameSuffix1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeNone),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix1)),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "none.%", "0"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "login_password"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testAccProjectKeyNoneConfig(nameSuffix2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeNone),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix2)),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "none.%", "0"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "login_password"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
		},
	})
}

func TestAcc_ProjectKeyResource_changeType(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	_, privateKey, _ := acctest.RandSSHKeyPair("")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectKeyLoginPasswordConfig(nameSuffix, "username", "password"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeLoginPassword),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "none"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.%", "2"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.login", "username"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "login_password.password", "password"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "ssh"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testAccProjectKeySSHConfig(nameSuffix, "username", privateKey, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectKeyExists("semaphoreui_project_key.test", ProjectKeyTypeSSH),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "none"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_key.test", "login_password"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.%", "3"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.login", "username"),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.private_key", privateKey),
					resource.TestCheckResourceAttr("semaphoreui_project_key.test", "ssh.passphrase", ""),
					resource.TestCheckResourceAttrSet("semaphoreui_project_key.test", "id"),
				),
			},
		},
	})
}
