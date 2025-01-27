# Import ID is specified by the string "project/{project_id}/view/{view_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {view_id} is the ID of the view in SemaphoreUI.
terraform import semaphoreui_project_view.example project/1/view/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_view.example
  id = "project/1/view/2"
}
