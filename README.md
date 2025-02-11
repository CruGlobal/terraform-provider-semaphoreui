# Terraform SemaphoreUI Provider

The [SemaphoreUI Provider](https://registry.terraform.io/providers/CruGlobal/semaphoreui/latest/docs) enables [Terraform](https://terraform.io) to manage [SemaphoreUI](https://semaphoreui.com/) resources.

This project uses Conventional Commits in order to automatically manage releases as well as keeping the CHANGELOG.md updated. CHANGELOG follows the Keep a Changelog spec.

### Requirements
This provider requires an installation of SemaphoreUI running at least version v2.11.2.

### SemaphoreUI API Client
The SemaphoreUI API client is generated from the Swagger (OpenAPI-2.0) [api-docs.yml](https://github.com/semaphoreui/semaphore/blob/develop/api-docs.yml) using [go-swagger](https://goswagger.io/go-swagger/).
To re-generate the client, ensure you have [go-swagger](https://goswagger.io/go-swagger/install/install-binary/) installed and configured on your system and then run `task client`.

