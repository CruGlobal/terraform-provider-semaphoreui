# Integration-scoped alias:
# Import ID is "project/{project_id}/integration/{integration_id}/alias/{alias_id}".
terraform import semaphoreui_integration_alias.example project/1/integration/2/alias/3

# Project-scoped alias (no integration):
# Import ID is "project/{project_id}/alias/{alias_id}".
terraform import semaphoreui_integration_alias.example project/1/alias/3
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_integration_alias.example
  id = "project/1/integration/2/alias/3"
}
