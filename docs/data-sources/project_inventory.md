---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "semaphoreui_project_inventory Data Source - semaphoreui"
subcategory: ""
description: |-
  The project inventory data source allows you to read the Ansible inventory or a Terraform/OpenTofu workspace for a project.
---

# semaphoreui_project_inventory (Data Source)

The project inventory data source allows you to read the Ansible inventory or a Terraform/OpenTofu workspace for a project.

## Example Usage

```terraform
# Lookup by Inventory ID
data "semaphoreui_project_inventory" "inventory" {
  project_id = 1
  id         = 2
}

# Lookup by Inventory Name
data "semaphoreui_project_inventory" "example" {
  project_id = 1
  name       = "Example Invewntory"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_id` (Number) The project ID that the inventory belongs to.

### Optional

- `id` (Number) The inventory ID. Ensure that one and only one attribute from this collection is set : `id`, `name`.
- `name` (String) The display name of the inventory or workspace. Ensure that one and only one attribute from this collection is set : `id`, `name`.

### Read-Only

- `file` (Attributes) Inventory File. (see [below for nested schema](#nestedatt--file))
- `ssh_key_id` (Number) The Project Key ID to use for accessing hosts in the inventory. This attribute is required for all inventory types in SemaphoreUI. You should set it to the ID of a Key of type `none` if the inventory doesn't require credentials, or for Workspace type inventories.
- `static` (Attributes) Static Inventory. (see [below for nested schema](#nestedatt--static))
- `static_yaml` (Attributes) Static YAML Inventory. (see [below for nested schema](#nestedatt--static_yaml))
- `terraform_workspace` (Attributes) Terraform Workspace. (see [below for nested schema](#nestedatt--terraform_workspace))

<a id="nestedatt--file"></a>
### Nested Schema for `file`

Read-Only:

- `become_key_id` (Number) The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.
- `path` (String) The path to the inventory file, relative to the Template or custom Repository. Example: `folder/hosts.yml`.
- `repository_id` (Number) The ID of the Repository that contains the inventory file.


<a id="nestedatt--static"></a>
### Nested Schema for `static`

Read-Only:

- `become_key_id` (Number) The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.
- `inventory` (String) Static inventory content in INI format.


<a id="nestedatt--static_yaml"></a>
### Nested Schema for `static_yaml`

Read-Only:

- `become_key_id` (Number) The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.
- `inventory` (String) Static inventory content in YAML format.


<a id="nestedatt--terraform_workspace"></a>
### Nested Schema for `terraform_workspace`

Read-Only:

- `workspace` (String) The Terraform workspace name.
