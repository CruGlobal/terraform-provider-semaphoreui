# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
