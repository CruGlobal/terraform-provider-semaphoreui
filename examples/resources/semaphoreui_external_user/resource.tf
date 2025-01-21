resource "semaphoreui_external_user" "example" {
  username = "login_name"
  name     = "Full Name"
  email    = "name@example.com"

  admin = false
  alert = false
}
