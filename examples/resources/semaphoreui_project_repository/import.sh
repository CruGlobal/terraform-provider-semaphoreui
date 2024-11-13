# Import ID is specified by the string "project/{project_id}/repository/{repository_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {repository_id} is the ID of the repository in SemaphoreUI.
terraform import semaphoreui_project_repository.example project/1/repository/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_repository.example
  id = "project/1/repository/2"
}
