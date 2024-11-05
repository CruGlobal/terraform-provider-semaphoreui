# Import ID is specified by the string "project/{project_id}/user/{user_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {user_id} is the ID of the user in SemaphoreUI.
terraform import semaphoreui_project_user.example project/1/user/3
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_user.example
  id = "project/1/user/3"
}
