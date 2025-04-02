# Configure the SemaphoreUI provider using required_providers.
terraform {
  required_providers {
    semaphoreui = {
      source  = "CruGlobal/semaphoreui"
      version = "~> 1.3.1"
    }
  }
}

provider "semaphoreui" {
  hostname  = "example.com"
  port      = 3000   # Default: 3000
  path      = "/api" # Default: "/api"
  protocol  = "http" # Default: "https"
  api_token = "your token"
}
