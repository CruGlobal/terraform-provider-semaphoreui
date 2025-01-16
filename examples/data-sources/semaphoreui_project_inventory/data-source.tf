# Lookup by Inventory ID
data "semaphoreui_project_inventory" "inventory" {
  project_id = 1
  id         = 2
}

# Lookup by Inventory Name
data "semaphoreui_project_inventory" "example" {
  project_id = 1
  name       = "Example Invewntory"
}
