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
