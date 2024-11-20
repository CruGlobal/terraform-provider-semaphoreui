# Import ID is specified by the string "project/{project_id}/template/{template_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {template_id} is the ID of the template in SemaphoreUI.
terraform import semaphoreui_project_template.example project/1/template/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_template.example
  id = "project/1/template/2"
}
