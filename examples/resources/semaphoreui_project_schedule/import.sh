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
