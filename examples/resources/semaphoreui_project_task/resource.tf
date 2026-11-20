variable "app_version" {
  type    = string
  default = "1.2.3"
}

resource "semaphoreui_project" "project" {
  name = "Example Project"
}

data "semaphoreui_project_template" "template" {
  project_id = semaphoreui_project.project.id
  name       = "Deploy"
}

# Run the template once, and fail the apply if the job does not succeed.
resource "semaphoreui_project_task" "deploy" {
  project_id  = semaphoreui_project.project.id
  template_id = data.semaphoreui_project_template.template.id

  message     = "Deployed by Terraform"
  environment = jsonencode({ app_version = var.app_version })

  # Terraform waits for the job to finish (the default), for up to 15 minutes.
  wait_for_completion = true
  timeout             = "15m"

  # Re-run the job whenever the version changes.
  triggers = {
    app_version = var.app_version
  }
}

output "deploy_status" {
  value = semaphoreui_project_task.deploy.status
}

# Queue a job without blocking the apply.
resource "semaphoreui_project_task" "smoke_test" {
  project_id  = semaphoreui_project.project.id
  template_id = data.semaphoreui_project_template.template.id

  limit    = "web-servers"
  dry_run  = true
  playbook = "smoke.yml"

  wait_for_completion = false
}

# Wait for the job and keep its console output in state so it can be inspected
# or fed to another resource. Job logs can be large; leave this off unless you
# need it.
resource "semaphoreui_project_task" "migration" {
  project_id  = semaphoreui_project.project.id
  template_id = data.semaphoreui_project_template.template.id

  playbook       = "migrate.yml"
  capture_output = true
}

output "migration_log" {
  value = semaphoreui_project_task.migration.output
}
