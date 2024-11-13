resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_project_key" "login_password" {
  project_id = semaphoreui_project.project.id
  name       = "Example Login"
  login_password = {
    login    = "username"
    password = "password"
  }
}

resource "semaphoreui_project_key" "ssh" {
  project_id = semaphoreui_project.project.id
  name       = "Example SSH"
  ssh = {
    passphrase  = "password"
    private_key = file("./id_rsa")
  }
}

resource "semaphoreui_project_key" "none" {
  project_id = semaphoreui_project.project.id
  name       = "Example None"
  none       = {}
}
