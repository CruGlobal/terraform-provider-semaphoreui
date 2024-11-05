resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_user" "user" {
  name     = "Example User"
  username = "example"
  email    = "user@example.com"
}

resource "semaphoreui_project_user" "project_user" {
  project_id = semaphoreui_project.project.id
  user_id    = semaphoreui_user.user.id
  role       = "owner"
}
