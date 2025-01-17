# Lookup by Template ID
data "semaphoreui_project_template" "template" {
  project_id = 1
  id         = 3
}

# Lookup by Template Name
data "semaphoreui_project_template" "build" {
  project_id = 1
  name       = "Build Application"
}
