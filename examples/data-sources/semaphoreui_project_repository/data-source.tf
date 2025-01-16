# Lookup by Repository ID
data "semaphoreui_project_repository" "repo" {
  project_id = 1
  id         = 3
}

# Lookup by Repository Name
data "semaphoreui_project_repository" "semaphore" {
  project_id = 1
  name       = "Semaphore"
}
