resource "semaphoreui_project" "example" {
  name               = "Example Project"
  alert              = false
  max_parallel_tasks = 0 # Unlimited
}
