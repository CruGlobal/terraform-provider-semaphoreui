# Lookup or Create External User
data "semaphoreui_user" "user" {
  username = "batman"
}

# Lookup or Create External User with additional attributes
data "semaphoreui_user" "batman" {
  username = "batman"
  name     = "Bruce Wayne"
  email    = "batman@wayneenterprises.com"
}
