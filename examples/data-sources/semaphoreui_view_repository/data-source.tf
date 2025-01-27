# Lookup by View ID
data "semaphoreui_project_view" "view" {
  project_id = 1
  id         = 3
}

# Lookup by View Name
data "semaphoreui_project_view" "prod" {
  project_id = 1
  title      = "Prod"
}
