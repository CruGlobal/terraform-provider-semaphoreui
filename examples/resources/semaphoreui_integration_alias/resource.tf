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

resource "semaphoreui_project_integration" "deploy" {
  project_id  = semaphoreui_project.project.id
  template_id = semaphoreui_project_template.deploy.id
  name        = "deploy-webhook"
}

# Integration-scoped alias — incoming requests trigger this integration's
# template directly. Most common form; emit `url` to share with the upstream
# webhook caller (GitHub, etc.).
resource "semaphoreui_integration_alias" "deploy" {
  project_id     = semaphoreui_project.project.id
  integration_id = semaphoreui_project_integration.deploy.id
}

output "deploy_webhook_url" {
  value = semaphoreui_integration_alias.deploy.url
}

# Project-scoped alias — omit integration_id. Incoming requests are routed to
# an integration via matchers defined on each integration. Useful when one
# entry-point URL fans out to multiple integrations based on payload.
resource "semaphoreui_integration_alias" "router" {
  project_id = semaphoreui_project.project.id
}
