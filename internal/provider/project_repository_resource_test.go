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

func testAccProjectRepositoryExists(resourceName string) resource.TestCheckFunc {
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

		response, err := testClient().Project.GetProjectProjectIDRepositoriesRepositoryID(&project.GetProjectProjectIDRepositoriesRepositoryIDParams{
			ProjectID:    projectId,
			RepositoryID: id,
		}, nil)
		if err != nil {
			return fmt.Errorf("project repository does not exist: %s", err.Error())
		}

		if response.Payload.GitURL != rs.Primary.Attributes["url"] {
			return fmt.Errorf("repository git_url mismatch: %s != %s", response.Payload.GitURL, rs.Primary.Attributes["url"])
		}
		if response.Payload.GitBranch != rs.Primary.Attributes["branch"] {
			return fmt.Errorf("repository git_branch mismatch: %s != %s", response.Payload.GitBranch, rs.Primary.Attributes["branch"])
		}
		return nil
	}
}

func testAccProjectRepositoryEmptyConfig(nameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name	     = "test-%[1]s"
  none = {}
}
`, nameSuffix)
}

func testAccProjectRepositoryConfig(nameSuffix string, url string, branch string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Test %[2]s"
  url        = "%[3]s"
  branch     = "%[4]s"
  ssh_key_id = semaphoreui_project_key.test.id
}`, testAccProjectRepositoryEmptyConfig(nameSuffix), nameSuffix, url, branch)
}

func testAccProjectRepositoryImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/repository/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectRepositoryResource_basic(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectRepositoryConfig(nameSuffix, "https://github.com/semaphoreui/semaphore.git", "develop"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectRepositoryExists("semaphoreui_project_repository.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_repository.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_repository.test", "url", "https://github.com/semaphoreui/semaphore.git"),
					resource.TestCheckResourceAttr("semaphoreui_project_repository.test", "branch", "develop"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_repository.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_repository.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_repository.test", "ssh_key_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectRepositoryImportID("semaphoreui_project_repository.test"),
			},
			// Update and Read testing
			{
				Config: testAccProjectRepositoryConfig(nameSuffix, "/absolute/path/to/repo", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectRepositoryExists("semaphoreui_project_repository.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_repository.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_repository.test", "url", "/absolute/path/to/repo"),
					resource.TestCheckResourceAttr("semaphoreui_project_repository.test", "branch", ""),
					resource.TestCheckResourceAttrSet("semaphoreui_project_repository.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_repository.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_repository.test", "ssh_key_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectRepositoryEmptyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_repository.test"),
				),
			},
		},
	})
}
