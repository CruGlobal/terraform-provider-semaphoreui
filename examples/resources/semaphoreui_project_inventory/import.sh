# Import ID is specified by the string "project/{project_id}/inventory/{inventory_id}".
# - {project_id} is the ID of the project in SemaphoreUI.
# - {inventory_id} is the ID of the inventory in SemaphoreUI.
terraform import semaphoreui_project_inventory.example project/1/inventory/1
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_project_inventory.example
  id = "project/1/inventory/1"
}
