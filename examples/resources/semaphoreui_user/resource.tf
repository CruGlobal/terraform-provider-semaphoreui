resource "semaphoreui_user" "example" {
  username = "login_name"
  name     = "Full Name"
  email    = "name@example.com"
  password = "abc123"

  admin    = false
  alert    = false
  external = false
}
