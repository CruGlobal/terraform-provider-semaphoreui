# When semaphoreui_project_environment is imported, the `secrets` values
# will be blank as SemaphoreUI does not return these values on the API.
#
# Import ID is specified by the string "project/{project_id}/environment/{environment_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {environment_id} is the ID of the environment in SemaphoreUI.
terraform import semaphoreui_project_environment.example project/1/environment/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_environment.example
  id = "project/1/environment/2"
}
