# Lookup by Project ID
data "semaphoreui_project" "project" {
  id = 1
}

# Lookup by Project Name
data "semaphoreui_project" "example" {
  name = "Example Project"
}
