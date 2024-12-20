---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "semaphoreui_project_schedule Resource - semaphoreui"
subcategory: ""
description: |-
  Provides a SemaphoreUI Project Schedule resource.
  Allows scheduling the execution of templates in a project.
---

# semaphoreui_project_schedule (Resource)

Provides a SemaphoreUI Project Schedule resource.

Allows scheduling the execution of templates in a project.

## Example Usage

```terraform
resource "semaphoreui_project" "project" {
  name = "Example Project"
}

data "semaphoreui_project_template" "template" {
  project_id = semaphoreui_project.project.id
  name       = "Example Template"
}

resource "semaphoreui_project_schedule" "schedule" {
  project_id  = semaphoreui_project.project.id
  template_id = data.semaphoreui_project_template.template.id
  name        = "Example Schedule"
  cron_format = "0 0 * * *"
  enabled     = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cron_format` (String) The cron format of the schedule.
- `enabled` (Boolean) Whether the schedule is enabled.
- `name` (String) The display name of the repository.
- `project_id` (Number) The project ID that the repository belongs to.
- `template_id` (Number) The template ID that the schedule executes.

### Read-Only

- `id` (Number) The repository ID.

## Import

Import is supported using the following syntax:

```shell
# Import ID is specified by the string "project/{project_id}/schedule/{schedule_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {schedule_id} is the ID of the schedule in SemaphoreUI.
terraform import semaphoreui_project_schedule.example project/1/schedule/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_schedule.example
  id = "project/1/schedule/2"
}
```
