# Import ID is specified by the string "project/{project_id}/task/{task_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {task_id} is the ID of the task (job run) in SemaphoreUI.
terraform import semaphoreui_project_task.example project/1/task/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_task.example
  id = "project/1/task/2"
}
