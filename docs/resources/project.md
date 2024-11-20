---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "semaphoreui_project Resource - semaphoreui"
subcategory: ""
description: |-
  Provides a SemaphoreUI Project resource.
  A project is a place to separate management activity. All SemaphoreUI activities occur within the context of a project.
---

# semaphoreui_project (Resource)

Provides a SemaphoreUI Project resource.

A project is a place to separate management activity. All SemaphoreUI activities occur within the context of a project.

## Example Usage

```terraform
resource "semaphoreui_project" "example" {
  name               = "Example Project"
  alert              = false
  max_parallel_tasks = 0 # Unlimited
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Project name.

### Optional

- `alert` (Boolean) Allow alerts for this project. Default `false`.
- `alert_chat` (String) Telegram chat ID.
- `max_parallel_tasks` (Number) Maximum number of parallel tasks, `0` for unlimited. Default `0`.

### Read-Only

- `created` (String) Creation date of the project.
- `id` (Number) Project ID.

## Import

Import is supported using the following syntax:

```shell
# Import ID is specified by the string "project/{project_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
terraform import semaphoreui_project.example project/1
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project.example
  id = "project/1"
}
```