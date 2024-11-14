resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_project_environment" "environment" {
  project_id = semaphoreui_project.project.id
  name       = "Example Environment"

  # extraVars
  variables = {
    key1 = "value1"
    key2 = "value2"
  }

  # environment variables
  environment_variables = {
    KEY1 = "value1"
    KEY2 = "value2"
  }

  # secrets
  secrets = [{
    # extraVar Secret
    name  = "key3"
    type  = "var"
    value = "value3"
    }, {
    # environment Secret
    name  = "KEY4"
    type  = "env"
    value = "value4"
  }]
}
