# Terraform SemaphoreUI Provider

The [SemaphoreUI Provider](https://registry.terraform.io/providers/CruGlobal/semaphoreui/latest/docs) enables [Terraform](https://terraform.io) to manage [SemaphoreUI](https://semaphoreui.com/) resources.

This project uses Conventional Commits in order to automatically manage releases as well as keeping the CHANGELOG.md updated. CHANGELOG follows the Keep a Changelog spec.

### Requirements
This provider requires an installation of [SemaphoreUI](https://semaphoreui.com/).

The provider is tested against the latest 3 versions of SemaphoreUI. See [Terraform Provider Acceptance Tests](https://github.com/CruGlobal/terraform-provider-semaphoreui/blob/main/.github/workflows/test.yml#L64) for a list of versions.

### SemaphoreUI API Client
The SemaphoreUI API client is generated from the Swagger (OpenAPI-2.0) [api-docs.yml](https://github.com/semaphoreui/semaphore/blob/develop/api-docs.yml) using [go-swagger](https://goswagger.io/go-swagger/).
To re-generate the client, ensure you have [go-swagger](https://goswagger.io/go-swagger/install/install-binary/) installed and configured on your system and then run `task client`.

### Support
This provider was developed for an internal use case and released as open source for anyone to use. It is not actively maintained, but we welcome contributions and issues.
