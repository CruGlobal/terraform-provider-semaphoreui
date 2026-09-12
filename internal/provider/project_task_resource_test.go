package provider

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"terraform-provider-semaphoreui/semaphoreui/client/task"
)

func testAccProjectTaskExists(resourceName string) resource.TestCheckFunc {
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

		response, err := testClient().Task.GetProjectProjectIDTasksTaskID(&task.GetProjectProjectIDTasksTaskIDParams{
			ProjectID: projectId,
			TaskID:    id,
		}, nil)
		if err != nil {
			return fmt.Errorf("error reading project task: %s", err.Error())
		}

		templateId, _ := strconv.ParseInt(rs.Primary.Attributes["template_id"], 10, 64)
		if response.Payload.TemplateID != templateId {
			return fmt.Errorf("task template_id mismatch: %d != %d", response.Payload.TemplateID, templateId)
		}

		return nil
	}
}

// testAccProjectTaskIDChanged asserts that the task ID recorded in state
// differs from previous, i.e. a new job was actually queued rather than the
// existing one being reused.
func testAccProjectTaskIDChanged(resourceName string, previous *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		current := rs.Primary.Attributes["id"]
		if current == "" {
			return fmt.Errorf("no ID is set")
		}
		if *previous != "" && *previous == current {
			return fmt.Errorf("expected a new task to be queued, but task ID is still %s", current)
		}
		*previous = current
		return nil
	}
}

func testAccProjectTaskDependencyConfig(nameSuffix string) string {
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

// testAccProjectTaskConfig queues a task without waiting for it. The test
// environment's repository is not clonable, so any job that actually runs ends
// in "error"; not waiting keeps these steps deterministic.
func testAccProjectTaskConfig(nameSuffix, message string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_task" "test" {
  project_id  = semaphoreui_project.test.id
  template_id = semaphoreui_project_template.test.id

  message = "%[2]s"

  wait_for_completion = false
}
`, testAccProjectTaskDependencyConfig(nameSuffix), message)
}

func testAccProjectTaskTriggersConfig(nameSuffix, triggerValue string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_task" "test" {
  project_id  = semaphoreui_project.test.id
  template_id = semaphoreui_project_template.test.id

  wait_for_completion = false

  triggers = {
    version = "%[2]s"
  }
}
`, testAccProjectTaskDependencyConfig(nameSuffix), triggerValue)
}

// testAccProjectTaskWaitConfig waits on a job that cannot succeed in the test
// environment, exercising the polling and failure-reporting path.
func testAccProjectTaskWaitConfig(nameSuffix string) string {
	return fmt.Sprintf(`
%[1]s
resource "semaphoreui_project_task" "test" {
  project_id  = semaphoreui_project.test.id
  template_id = semaphoreui_project_template.test.id

  wait_for_completion = true
  timeout             = "3m"
}
`, testAccProjectTaskDependencyConfig(nameSuffix))
}

func testAccProjectTaskImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}

		return fmt.Sprintf("project/%[1]s/task/%[2]s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["id"]), nil
	}
}

func TestAcc_ProjectTaskResource_basic(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccProjectTaskConfig(nameSuffix, "Queued by Terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTaskExists("semaphoreui_project_task.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "message", "Queued by Terraform"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "wait_for_completion", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "capture_output", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "timeout", "30m"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "debug", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "dry_run", "false"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "diff", "false"),
					resource.TestCheckNoResourceAttr("semaphoreui_project_task.test", "output"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_task.test", "id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_task.test", "project_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_task.test", "template_id"),
					resource.TestCheckResourceAttrSet("semaphoreui_project_task.test", "status"),
				),
			},
			// ImportState testing. status/output and the client-side knobs are
			// not part of the job record, so an import cannot reproduce them.
			{
				ResourceName:      "semaphoreui_project_task.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectTaskImportID("semaphoreui_project_task.test"),
				ImportStateVerifyIgnore: []string{
					"wait_for_completion", "timeout", "capture_output", "triggers", "status", "output",
				},
			},
			// Delete testing
			{
				Config: testAccProjectTaskDependencyConfig(nameSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccResourceNotExists("semaphoreui_project_task.test"),
				),
			},
		},
	})
}

// TestAcc_ProjectTaskResource_triggers verifies that changing an input replaces
// the resource, queueing a genuinely new job rather than updating the old one.
func TestAcc_ProjectTaskResource_triggers(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	taskID := ""
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTaskTriggersConfig(nameSuffix, "1.0.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTaskExists("semaphoreui_project_task.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "triggers.version", "1.0.0"),
					testAccProjectTaskIDChanged("semaphoreui_project_task.test", &taskID),
				),
			},
			{
				Config: testAccProjectTaskTriggersConfig(nameSuffix, "2.0.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectTaskExists("semaphoreui_project_task.test"),
					resource.TestCheckResourceAttr("semaphoreui_project_task.test", "triggers.version", "2.0.0"),
					testAccProjectTaskIDChanged("semaphoreui_project_task.test", &taskID),
				),
			},
		},
	})
}

// TestAcc_ProjectTaskResource_waitFails covers the wait path: the test
// environment's repository cannot be cloned, so the job ends in "error" and the
// apply must fail rather than silently succeed.
func TestAcc_ProjectTaskResource_waitFails(t *testing.T) {
	nameSuffix := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectTaskWaitConfig(nameSuffix),
				// Deliberately does NOT accept the timeout diagnostic: the point
				// of this test is that polling observes the terminal "error"
				// status. Passing on timeout would keep it green even if
				// terminal-status detection were broken.
				ExpectError: regexp.MustCompile(`SemaphoreUI Project Task Failed`),
			},
		},
	})
}
