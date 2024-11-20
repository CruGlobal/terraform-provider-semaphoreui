resource "semaphoreui_project" "project" {
  name = "Example Project"
}

data "semaphoreui_project_template" "template" {
  project_id = semaphoreui_project.project.id
  name       = "Example Template"
}

resource "semaphoreui_project_schedule" "schedule" {
  project_id  = semaphoreui_project.project.id
  template_id = data.semaphoreui_project_template.template.id
  name        = "Example Schedule"
  cron_format = "0 0 * * *"
  enabled     = true
}
