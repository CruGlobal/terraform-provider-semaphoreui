resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_project_key" "none" {
  project_id = semaphoreui_project.project.id
  name       = "None"
  none       = {}
}

# Static Inventory Example
resource "semaphoreui_project_inventory" "static" {
  project_id = semaphoreui_project.project.id
  name       = "Static Inventory"
  ssh_key_id = semaphoreui_project_key.none.id
  static = {
    inventory = <<-EOT
      [website]
      172.18.8.40
      172.18.8.41
    EOT
    # Optional
    become_key_id = 3
  }
}

# Static YAML Inventory Example
resource "semaphoreui_project_inventory" "static_yaml" {
  project_id = semaphoreui_project.project.id
  name       = "Static YAML Inventory"
  ssh_key_id = semaphoreui_project_key.none.id
  static_yaml = {
    inventory = yamlencode({
      all = {
        children = {
          website = {
            hosts = {
              "172.18.8.40" = {}
              "172.18.8.41" = {}
            }
          }
        }
      }
    })
  }
}

# File Inventory Example
resource "semaphoreui_project_inventory" "file" {
  project_id = semaphoreui_project.project.id
  name       = "Static Inventory"
  ssh_key_id = semaphoreui_project_key.none.id
  file = {
    path = "inventory/dev/hosts"
    # Optional
    repository_id = 1
  }
}

# Terraform/OpenTofu Workspace Inventory Example
resource "semaphoreui_project_inventory" "terraform" {
  project_id = semaphoreui_project.project.id
  name       = "Terraform Workspace Inventory"
  ssh_key_id = semaphoreui_project_key.none.id
  terraform_workspace = {
    workspace = "workspace-name"
  }
}
