# Lookup by Key ID
data "semaphoreui_project_key" "key" {
  project_id = 1
  id         = 3
}

# Lookup by Key Name
data "semaphoreui_project_key" "none" {
  project_id = 1
  name       = "None"
}
