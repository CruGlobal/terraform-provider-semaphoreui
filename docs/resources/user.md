---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "semaphoreui_user Resource - semaphoreui"
subcategory: ""
description: |-
  Provides a SemaphoreUI User resource.
---

# semaphoreui_user (Resource)

Provides a SemaphoreUI User resource.

## Example Usage

```terraform
resource "semaphoreui_user" "example" {
  username = "login_name"
  name     = "Full Name"
  email    = "name@example.com"
  password = "abc123"

  admin    = false
  alert    = false
  external = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) Email address.
- `name` (String) Display name.
- `username` (String) Username.

### Optional

- `admin` (Boolean) Is the user an admin? Default `false`.
- `alert` (Boolean) Send alerts to the user's email? Default `false`.
- `external` (Boolean) Is the user linked to an external identity provider? Default `false`.
- `password` (String, Sensitive) Login Password.

### Read-Only

- `created` (String) Creation date of the user.
- `id` (Number) User ID.

## Import

Import is supported using the following syntax:

```shell
# Import ID is specified by the string "user/{user_id}".
# - {user_id} is the ID of the user in SemaphoreUI.
terraform import semaphoreui_user.example user/1
```
Or using `import {}` block in the configuration file:
```hcl
import {
  to = semaphoreui_user.example
  id = "user/1"
}
```
