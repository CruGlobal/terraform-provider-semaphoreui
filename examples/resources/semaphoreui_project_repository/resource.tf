resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_project_key" "none" {
  project_id = semaphoreui_project.project.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_repository" "repository" {
  project_id = semaphoreui_project.project.id
  name       = "Example Repository"
  url        = "https://github.com/semaphoreui/semaphore.git"
  branch     = "develop"
  ssh_key_id = semaphoreui_project_key.none.id
}
