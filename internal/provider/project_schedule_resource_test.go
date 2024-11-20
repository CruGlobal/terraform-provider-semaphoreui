package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"strconv"
	"terraform-provider-semaphoreui/semaphoreui/client/schedule"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectScheduleExists(resourceName string) resource.TestCheckFunc {
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

		response, err := testClient().Schedule.GetProjectProjectIDSchedulesScheduleID(&schedule.GetProjectProjectIDSchedulesScheduleIDParams{
			ProjectID:  projectId,
			ScheduleID: id,
		}, nil)
		if err != nil {
			return fmt.Errorf("error reading project schedule: %s", err.Error())
		}

		if response.Payload.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("schedule name mismatch: %s != %s", response.Payload.Name, rs.Primary.Attributes["name"])
		}

		if response.Payload.CronFormat != rs.Primary.Attributes["cron_format"] {
			return fmt.Errorf("schedule cron_format mismatch: %s != %s", response.Payload.CronFormat, rs.Primary.Attributes["cron_format"])
		}

		return nil
	}
}

func testAccProjectScheduleDependencyConfig(nameSuffix string) string {
	return fmt.Sprintf(`
resource "semaphoreui_project" "test" {
  name = "test-%[1]s"
}

resource "semaphoreui_project_key" "test" {
  project_id = semaphoreui_project.test.id
  name       = "None-%[1]s"
  none       = {}
}

resource "semaphoreui_project_repository" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Repo-%[1]s"
  url        = "git@github.com:example/test.git"
  branch     = "main"
  ssh_key_id = semaphoreui_project_key.test.id
}

resource "semaphoreui_project_inventory" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Inventory-%[1]s"
  ssh_key_id = semaphoreui_project_key.test.id
  file = {
    path          = "path/to/inventory"
    repository_id = semaphoreui_project_repository.test.id
  }
}

resource "semaphoreui_project_environment" "test" {
  project_id = semaphoreui_project.test.id
  name       = "Env-%[1]s"
}

resource "semaphoreui_project_template" "test" {
  project_id     = semaphoreui_project.test.id
  environment_id = semaphoreui_project_environment.test.id
  inventory_id   = semaphoreui_project_inventory.test.id
  repository_id  = semaphoreui_project_repository.test.id
  name           = "Template-%[1]s"
  playbook       = "playbook.yml"
}
`, nameSuffix)
}

func testAccProjectScheduleConfig(nameSuffix string, enabled bool) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_schedule" "test" {
  project_id  = semaphoreui_project.test.id
  name        = "Test %[2]s"
  template_id = semaphoreui_project_template.test.id
  cron_format = "0 0 * * *"
  enabled = %[3]t
}
`, testAccProjectScheduleDependencyConfig(nameSuffix), nameSuffix, enabled)
}

func testAccProjectScheduleImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/schedule/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectScheduleResource_basic(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectScheduleConfig(nameSuffix, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectScheduleExists("semaphoreui_project_schedule.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_schedule.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_schedule.test", "cron_format", "0 0 * * *"),
					resource.TestCheckResourceAttr("semaphoreui_project_schedule.test", "enabled", "false"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_schedule.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_schedule.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_schedule.test", "template_id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "semaphoreui_project_schedule.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectScheduleImportID("semaphoreui_project_schedule.test"),
			},
			// Update testing
			{
				Config: testAccProjectScheduleConfig(nameSuffix, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectScheduleExists("semaphoreui_project_schedule.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_schedule.test", "name", fmt.Sprintf("Test %s", nameSuffix)),
					resource.TestCheckResourceAttr("semaphoreui_project_schedule.test", "cron_format", "0 0 * * *"),
					resource.TestCheckResourceAttr("semaphoreui_project_schedule.test", "enabled", "true"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_schedule.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_schedule.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_schedule.test", "template_id"),
				),
			},
			// Delete testing
			{
				Config: testAccProjectScheduleDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_schedule.test"),
				),
			},
		},
	})
}
