resource "semaphoreui_project" "project" {
  name = "Project"
}

resource "semaphoreui_project_key" "none" {
  project_id = semaphoreui_project.project.id
  name       = "None"
  none       = {}
}

resource "semaphoreui_project_repository" "repo" {
  project_id = semaphoreui_project.project.id
  name       = "Repo"
  url        = "git@github.com:example/test.git"
  branch     = "main"
  ssh_key_id = semaphoreui_project_key.none.id
}

resource "semaphoreui_project_inventory" "inventory" {
  project_id = semaphoreui_project.project.id
  name       = "Inventory"
  ssh_key_id = semaphoreui_project_key.none.id
  file = {
    path          = "path/to/inventory"
    repository_id = semaphoreui_project_repository.repo.id
  }
}

resource "semaphoreui_project_environment" "environment" {
  project_id = semaphoreui_project.project.id
  name       = "Environment"
  secrets = [{
    name  = "SECRET_ONE"
    type  = "var"
    value = "VALUE_ONE"
  }]
}

resource "semaphoreui_project_repository" "repository" {
  project_id = semaphoreui_project.project.id
  name       = "Example Repository"
  url        = "https://github.com/semaphoreui/semaphore.git"
  branch     = "develop"
  ssh_key_id = semaphoreui_project_key.none.id
}

# Task Template
resource "semaphoreui_project_template" "task" {
  project_id     = semaphoreui_project.project.id
  environment_id = semaphoreui_project_environment.environment.id
  inventory_id   = semaphoreui_project_inventory.inventory.id
  repository_id  = semaphoreui_project_repository.repo.id
  name           = "Template"
  playbook       = "playbook.yml"
  description    = "Description"
  arguments = [
    "--help",
    "--vvv",
  ]
  allow_override_args_in_task = true

  survey_vars = [{
    name     = "age"
    title    = "What is your age?"
    required = true
    type     = "integer"
    }, {
    name  = "question"
    title = "Pick one."
    type  = "enum"
    enum_values = {
      "First Value"  = "1"
      "Second Value" = "2"
    }
  }]

  vaults = [{
    name = "" # default vault
    password = {
      vault_key_id = semaphoreui_project_key.none.id
    }
    }, {
    name = "database"
    script = {
      script = "path/to/script-client.py"
    }
  }]
}

# Build Template
resource "semaphoreui_project_template" "build" {
  project_id     = semaphoreui_project.project.id
  environment_id = semaphoreui_project_environment.environment.id
  inventory_id   = semaphoreui_project_inventory.inventory.id
  repository_id  = semaphoreui_project_repository.repo.id
  name           = "Build"
  playbook       = "build.yml"
  build = {
    start_version = "1.0.0"
  }
}

# Deploy Template
resource "semaphoreui_project_template" "deploy" {
  project_id     = semaphoreui_project.project.id
  environment_id = semaphoreui_project_environment.environment.id
  inventory_id   = semaphoreui_project_inventory.inventory.id
  repository_id  = semaphoreui_project_repository.repo.id
  name           = "Deploy"
  playbook       = "deploy.yml"
  deploy = {
    build_template_id = semaphoreui_project_template.build.id
    autorun           = false
  }
}
