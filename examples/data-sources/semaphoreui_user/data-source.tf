# Lookup User by ID
data "semaphoreui_user" "user" {
  # SemaphoreUI User ID
  id = 1
}

# Lookup User by Username
data "semaphoreui_user" "batman" {
  username = "batman"
}

# Lookup User by Email
data "semaphoreui_user" "superman" {
  username = "clark.kent@dailyplanet.com"
}
