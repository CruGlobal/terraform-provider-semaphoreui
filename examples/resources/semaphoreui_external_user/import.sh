# Import ID is specified by the string "user/{user_id}".
# - {user_id} is the ID of the user in SemaphoreUI.
terraform import semaphoreui_external_user.example user/1
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_external_user.example
  id = "user/1"
}
