# Terraform SemaphoreUI Provider

[![Registry](https://img.shields.io/badge/registry-CruGlobal%2Fsemaphoreui-623CE4)](https://registry.terraform.io/providers/CruGlobal/semaphoreui/latest/docs)

> **Status: AI-supported, not actively maintained.** This provider was
> developed for an internal use case and is now maintained on a
> best-effort basis with AI assistance. Dependabot keeps dependencies
> and security advisories up to date automatically (patch and minor
> bumps auto-merge; majors require manual review). Feature work, bug
> fixes, and other changes happen on a best-effort basis. **Pull
> requests and issues are welcome** — they may take time to be
> reviewed. See [`CONTRIBUTING.md`](CONTRIBUTING.md) for the
> contribution workflow.

The [SemaphoreUI Provider](https://registry.terraform.io/providers/CruGlobal/semaphoreui/latest/docs)
enables [Terraform](https://terraform.io) to manage
[SemaphoreUI](https://semaphoreui.com/) resources — projects,
environments, inventories, keys, repositories, schedules, templates,
users, and views.

Releases are managed automatically by [release-please](https://github.com/googleapis/release-please)
based on Conventional Commit messages. The [`CHANGELOG.md`](CHANGELOG.md)
follows the [Keep a Changelog](https://keepachangelog.com/) spec.

## Requirements

A reachable installation of [SemaphoreUI](https://semaphoreui.com/) and
an API token. See the
[provider docs](https://registry.terraform.io/providers/CruGlobal/semaphoreui/latest/docs)
for the available configuration options.

The provider is tested against the **three most recent SemaphoreUI minor
lines** — currently `v2.16.x`, `v2.17.x`, and `v2.18.x`. See the
[`Tests` workflow](https://github.com/CruGlobal/terraform-provider-semaphoreui/blob/main/.github/workflows/test.yml)
for the exact pinned versions.

## SemaphoreUI API client

The SemaphoreUI API client (`semaphoreui/client/`, `semaphoreui/models/`)
is generated from the local
[`api-docs.yml`](api-docs.yml) via
[go-swagger](https://goswagger.io/go-swagger/). The local spec is a
patched copy of the upstream
[semaphoreui/semaphore](https://github.com/semaphoreui/semaphore/blob/develop/api-docs.yml)
spec — small nullability tweaks are re-applied on top of each upstream
import so the generated client matches the actual API behavior. See
[`CONTRIBUTING.md`](CONTRIBUTING.md) for the regeneration workflow.
