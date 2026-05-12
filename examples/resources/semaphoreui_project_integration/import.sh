# Import ID is specified by the string "project/{project_id}/integration/{integration_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {integration_id} is the ID of the integration in SemaphoreUI.
terraform import semaphoreui_project_integration.example project/1/integration/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_integration.example
  id = "project/1/integration/2"
}
