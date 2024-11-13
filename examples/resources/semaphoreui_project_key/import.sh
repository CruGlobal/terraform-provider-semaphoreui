# When semaphoreui_project_key is imported, the `login_password` and `ssh` attributes
# will be blank as SemaphoreUI does not return these values on the API.
#
# Import ID is specified by the string "project/{project_id}/key/{key_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {key_id} is the ID of the key in SemaphoreUI.
terraform import semaphoreui_project_key.example project/1/key/2
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_key.example
  id = "project/1/key/2"
}
