# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.5.0](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v1.4.1...v1.5.0) (2026-05-12)


### Features

* add `semaphoreui_integration_alias` resource for managing webhook URLs that trigger integrations ([#83](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/83)) ([2eaa534](https://github.com/CruGlobal/terraform-provider-semaphoreui/commit/2eaa5347b8bd4d1479618b908df32e04790ed4b0))
* add `semaphoreui_project_integration` resource and data source ([#80](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/80)) ([ae8b646](https://github.com/CruGlobal/terraform-provider-semaphoreui/commit/ae8b646e092e04715bff964a76a8ca8354da4b1c))
* add `task_params` (Ansible `tags`/`skip_tags`/`limit`/`diff`/etc. and Terraform `auto_approve`/`destroy`/`plan`/`upgrade`) on `semaphoreui_project_template` and `semaphoreui_project_integration`, closes [#53](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/53) ([#82](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/82)) ([b5620ab](https://github.com/CruGlobal/terraform-provider-semaphoreui/commit/b5620ab2c75818a0e3ec30082cd417cb0b5b25fe))
* write-only (`*_wo`) attributes for sensitive `semaphoreui_project_key` fields (`password`, `passphrase`, `private_key`), supporting ephemeral values from Vault and other secret sources, closes [#58](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/58) ([#84](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/84)) ([bb80fe3](https://github.com/CruGlobal/terraform-provider-semaphoreui/commit/bb80fe3885d0b0b401ec2b5878056f577d545dae))
* add `tofu_workspace` inventory type to `semaphoreui_project_inventory` for OpenTofu workspaces ([#78](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/78))


### Bug Fixes

* `semaphoreui_project_environment` secret values now persist on update — the converter was omitting the `secret` field when marking an existing secret for update, so changes were silently dropped, closes [#68](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/68) ([#78](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/78))
* omit the default port (`:443`/`:80`) from the HTTPS Host header to support strict reverse proxies (notably F5) that reject SNI/Host containing the default port, closes [#56](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/56) ([#78](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/78))
* `playbook` is now optional on `semaphoreui_project_template` when `app` is `terraform` or `tofu`; a config validator still requires it for other apps (Ansible, Bash, etc.), closes [#26](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/26) ([#78](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/78))
* `semaphoreui_project_template` works correctly against SemaphoreUI v2.16+ — the API changed `environment_id` (singular) to `environment_ids` (array) and the previous client always read back `0`. The provider now writes both and reads from the array with a fallback ([#78](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/78))


### Documentation

* mark provider as AI-supported, not actively maintained, with a top-of-README status callout ([#79](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/79))
* add [`CONTRIBUTING.md`](CONTRIBUTING.md) covering toolchain, pre-PR checks, Conventional Commits, the failing-test-first norm, and the API client regeneration workflow ([#79](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/79))
* add `CLAUDE.md` for AI-assisted development sessions ([#79](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/79))


### Miscellaneous

* bump Go to 1.26.3 and refresh all dependencies (terraform-plugin-framework 1.15→1.19, terraform-plugin-go 0.28→0.31, terraform-plugin-testing 1.13→1.16, golangci-lint v8→v9, etc.) ([#76](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/76))
* bump test matrix to the latest three SemaphoreUI minor lines (v2.16.51 / v2.17.39 / v2.18.2); Taskfile default is now v2.18.2 ([#78](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/78))
* regenerate API client from upstream `api-docs.yml` at SemaphoreUI v2.18.2; internally restructures the client into per-resource sub-packages (Inventory, KeyStore, Repository, Schedule, Task, Template, VariableGroup) — no user-facing schema changes ([#78](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/78))
* add Dependabot auto-merge workflow that auto-approves and squash-merges patch/minor/security PRs; majors still require human review ([#76](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/76))
* bump `hashicorp/setup-terraform` from 3 to 4 ([#77](https://github.com/CruGlobal/terraform-provider-semaphoreui/pull/77))

## [1.4.1](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v1.4.0...v1.4.1) (2025-06-04)


### Bug Fixes

* project max_parallel_tasks is omitted by the API when empty. ([#49](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/49)) ([b345643](https://github.com/CruGlobal/terraform-provider-semaphoreui/commit/b345643ae4a7a89e9ac273e912cd296de35d2675))

## [1.4.0](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v1.3.0...v1.4.0) (2025-06-03)


### Features

* Add `tls_skip_verify` provider option ([#48](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/48)) ([293d862](https://github.com/CruGlobal/terraform-provider-semaphoreui/commit/293d86265695dd815678283d2bcc7770c2c0559d)), closes [#41](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/41)


### Bug Fixes

* provider docs used incorrect provider name ([#33](https://github.com/CruGlobal/terraform-provider-semaphoreui/issues/33)) ([18e14c3](https://github.com/CruGlobal/terraform-provider-semaphoreui/commit/18e14c347950d88953e22a7eecb571a137bdb8a9))

## [Unreleased](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v1.3.0...HEAD)

## [v1.0.0](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v0.1.1...v1.0.0) - 2024-11-20

### Added

- Initial release of the provider

## [v1.3.0](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v1.2.0...v1.3.0) - 2025-01-27

### Added

- feat(project_view): Add project view resource and data source @Omicron7 (#20)

### Dependency Updates

- chore(gomod): bump github.com/hashicorp/terraform-plugin-go from 0.25.0 to 0.26.0 @[dependabot[bot]](https://github.com/apps/dependabot) (#19)

## [v1.2.0](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v1.1.0...v1.2.0) - 2025-01-23

### Fixed

- fix(external_user): Refactor external_user from resource to data source. @Omicron7 (#18)

## [v1.1.0](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/v1.0.1...v1.1.0) - 2025-01-21

### Added

- feat(external_user): Add external_user resource @Omicron7 (#17)
- feat(data): Adding Data Sources @Omicron7 (#12)
- chore(dependabot): Add commit message and labels @Omicron7 (#7)

### Fixed

- chore(dependabot): Fix PR title and remove version labels @Omicron7 (#10)

### Dependency Updates

<details>
<summary>6 changes</summary>
- chore(gomod): bump golang.org/x/net from 0.28.0 to 0.33.0 @[dependabot[bot]](https://github.com/apps/dependabot) (#16)
- chore(gomod): bump github.com/hashicorp/terraform-plugin-framework-validators from 0.15.0 to 0.16.0 @[dependabot[bot]](https://github.com/apps/dependabot) (#15)
- Bump golang.org/x/crypto from 0.21.0 to 0.31.0 in /tools @[dependabot[bot]](https://github.com/apps/dependabot) (#14)
- chore(gomod): bump golang.org/x/crypto from 0.29.0 to 0.31.0 @[dependabot[bot]](https://github.com/apps/dependabot) (#13)
- chore(github-actions): bump amannn/action-semantic-pull-request from 5.4.0 to 5.5.3 @[dependabot[bot]](https://github.com/apps/dependabot) (#8)
- chore(github-actions): bump release-drafter/release-drafter from 5 to 6 @[dependabot[bot]](https://github.com/apps/dependabot) (#9)
</details>
## [v1.0.1](https://github.com/CruGlobal/terraform-provider-semaphoreui/compare/main...v1.0.1) - 2024-11-26
### Fixed
- fix: Update API client and fix GitHub workflows @Omicron7 (#6)

### Dependency Updates

- Bump github.com/hashicorp/terraform-plugin-framework-validators from 0.14.0 to 0.15.0 @dependabot (#3)
