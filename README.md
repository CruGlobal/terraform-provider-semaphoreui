# Terraform SemaphoreUI Provider

The [SemaphoreUI Provider](https://registry.terraform.io/providers/CruGlobal/semaphoreui/latest/docs) enables [Terraform](https://terraform.io) to manage [SemaphoreUI](https://semaphoreui.com/) resources.

## WARNING
This provider requires changes to SemaphoreUI API which are not yet released. It will not function until those changes have been released. https://github.com/semaphoreui/semaphore/pull/2553


### SemaphoreUI API Client
The SemaphoreUI API client is generated from the Swagger (OpenAPI-2.0) [api-docs.yml](https://github.com/semaphoreui/semaphore/blob/develop/api-docs.yml) using [go-swagger](https://goswagger.io/go-swagger/).
To re-generate the client, ensure you have [go-swagger](https://goswagger.io/go-swagger/install/install-binary/) installed and configured on your system and then run `task client`.

