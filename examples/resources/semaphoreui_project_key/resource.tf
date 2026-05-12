resource "semaphoreui_project" "project" {
  name = "Example Project"
}

resource "semaphoreui_project_key" "login_password" {
  project_id = semaphoreui_project.project.id
  name       = "Example Login"
  login_password = {
    login    = "username"
    password = "password"
  }
}

resource "semaphoreui_project_key" "ssh" {
  project_id = semaphoreui_project.project.id
  name       = "Example SSH"
  ssh = {
    passphrase  = "password"
    private_key = file("./id_rsa")
  }
}

resource "semaphoreui_project_key" "none" {
  project_id = semaphoreui_project.project.id
  name       = "Example None"
  none       = {}
}

# Write-only / ephemeral secrets — for SSH keys or passwords fetched from a
# secret store like Vault. The `*_wo` values are sent to SemaphoreUI on apply
# but never persisted to Terraform state. Bump the matching `_wo_version` to
# rotate the secret.
ephemeral "vault_kv_secret_v2" "deploy_key" {
  mount = "secret"
  name  = "deploy-ssh-key"
}

resource "semaphoreui_project_key" "ephemeral_ssh" {
  project_id = semaphoreui_project.project.id
  name       = "Ephemeral SSH"
  ssh = {
    login                  = "deploy"
    private_key_wo         = ephemeral.vault_kv_secret_v2.deploy_key.data["private_key"]
    private_key_wo_version = 1 # bump to push a rotated key
  }
}
