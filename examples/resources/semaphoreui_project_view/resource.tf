resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_project_view" "view" {
  project_id = semaphoreui_project.project.id
  title      = "Section A"
  position   = 0
}
