resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_project_key" "none" {
  project_id = semaphoreui_project.project.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_key" "webhook_secret" {
  project_id = semaphoreui_project.project.id
  name       = "Webhook HMAC"
  login_password = {
    login    = "webhook"
    password = "shared-secret-rotate-me"
  }
}

resource "semaphoreui_project_repository" "repository" {
  project_id = semaphoreui_project.project.id
  name       = "Example Repository"
  url        = "https://github.com/semaphoreui/semaphore.git"
  branch     = "develop"
  ssh_key_id = semaphoreui_project_key.none.id
}

resource "semaphoreui_project_inventory" "inventory" {
  project_id = semaphoreui_project.project.id
  name       = "Example Inventory"
  ssh_key_id = semaphoreui_project_key.none.id
  static = {
    inventory = "localhost ansible_connection=local"
  }
}

resource "semaphoreui_project_template" "deploy" {
  project_id    = semaphoreui_project.project.id
  inventory_id  = semaphoreui_project_inventory.inventory.id
  repository_id = semaphoreui_project_repository.repository.id
  name          = "Deploy"
  playbook      = "deploy.yml"
}

# Webhook with no authentication
resource "semaphoreui_project_integration" "open" {
  project_id  = semaphoreui_project.project.id
  template_id = semaphoreui_project_template.deploy.id
  name        = "open-webhook"
}

# Webhook authenticated by a shared HMAC secret in the X-Hub-Signature header
resource "semaphoreui_project_integration" "github" {
  project_id     = semaphoreui_project.project.id
  template_id    = semaphoreui_project_template.deploy.id
  name           = "github-webhook"
  auth_method    = "hmac"
  auth_secret_id = semaphoreui_project_key.webhook_secret.id
  auth_header    = "X-Hub-Signature-256"
}
