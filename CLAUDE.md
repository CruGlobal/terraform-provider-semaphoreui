# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Terraform provider for [SemaphoreUI](https://semaphoreui.com/), built on the terraform-plugin-framework (not the legacy SDKv2). Published to the Terraform Registry as `CruGlobal/semaphoreui`.

Conventional Commits drive release-please (CHANGELOG.md + version bumps), so commit messages matter.

## Common commands

The project uses [Task](https://taskfile.dev) (`Taskfile.yml`). Common targets:

- `task build` — `go build -v ./...`
- `task fmt` — `gofmt -s -w -e .`
- `task lint` — `golangci-lint run --tests=false` (tests are intentionally excluded)
- `task test` — unit tests: `go test -v -cover -timeout=120s -parallel=10 ./internal/...`
- `task testacc` — acceptance tests against dockerized SemaphoreUI (see below)
- `task generate` — regenerates `docs/` via tfplugindocs (CI fails if the diff is non-empty)
- `task client` — regenerates `semaphoreui/` API client from `api-docs.yml` via `go-swagger` (requires `swagger` binary)
- `task docker:start` / `task docker:stop` — bring the test environment up/down manually

Run a single test:
```
go test -v -run TestAcc_ProjectResource_basic ./internal/provider/
```

For acceptance tests, `TF_ACC=1` plus the `SEMAPHOREUI_*` env vars must be set — `task testacc` does this for you. Tests named `TestAcc_*` require the live API.

## Acceptance test environment

`task testacc` orchestrates a real test environment:

1. `task docker:start` brings up SemaphoreUI + MySQL via `docker-compose.yml` on ports `13000` (Semaphore) and `13306` (MySQL).
2. `scripts/wait_for_test_env_ready.sh` polls the container health check.
3. `scripts/setup_test_env.sh` injects a freshly generated API token directly into the MySQL `user__token` table — there is no API endpoint for token creation, so this DB write is the only way to seed auth.
4. `go test` runs against the live server.
5. `task docker:stop` tears it all down (with `-v`, so all data is lost).

The SemaphoreUI version under test comes from the `SEMAPHORE_VERSION` env var (default `v2.18.2`). CI runs the matrix across the latest 3 minor lines (currently `v2.16.51 / v2.17.39 / v2.18.2`); see `.github/workflows/test.yml`.

## Architecture

```
main.go                  — providerserver entrypoint; version injected by goreleaser
internal/provider/       — all resources, data sources, schemas, and tests (single package)
internal/stringvalidator — custom validators (e.g. cron format)
semaphoreui/client/      — GENERATED go-swagger HTTP client (do not hand-edit)
semaphoreui/models/      — GENERATED request/response models
api-docs.yml             — upstream OpenAPI 2.0 spec, source of truth for the client
tools/                   — separate Go module hosting tfplugindocs for `task generate`
examples/                — TF examples consumed by tfplugindocs to render `docs/`
```

### Resource/data-source structure

Each Terraform type follows a three-file convention in `internal/provider/`:

- `<name>_schema.go` — defines the `<Name>Model` struct + a `<Name>Schema()` that returns a `superschema.Schema` (from `orange-cloudavenue/terraform-plugin-framework-superschema`). The superschema lets one definition serve both resource and data-source variants by tagging attributes with `Resource:` / `DataSource:` / `Common:` overrides.
- `<name>_resource.go` — implements `resource.Resource` with `Configure`, `Schema` (delegates to the schema file), `Create`, `Read`, `Update`, `Delete`, and `ImportState`.
- `<name>_data_source.go` — implements `datasource.DataSource`.

Resources and data sources are wired up by adding constructors to `Resources()` / `DataSources()` in `provider.go`. A new resource means: schema file, resource file, registration in `provider.go`, an `examples/resources/<name>/` directory, and a test file. Then `task generate` produces the matching `docs/` page.

### Client wiring

`provider.go` `Configure()` builds a `go-openapi/runtime/client` httptransport with bearer-token auth from `SEMAPHOREUI_API_TOKEN` and stashes the generated `*apiclient.SemaphoreUI` on both `resp.DataSourceData` and `resp.ResourceData`. Each resource's `Configure` casts `req.ProviderData` back to `*apiclient.SemaphoreUI`.

The OpenAPI-generated client splits endpoints into per-resource sub-clients (driven by `tags` in `api-docs.yml`). The `SemaphoreUI` struct exposes:

```
Authentication, Integration, Inventory, KeyStore, Operations,
Project, Repository, Schedule, Task, Template, User, VariableGroup.
```

`VariableGroup` owns the environment endpoints — non-obvious because the resource is `semaphoreui_project_environment`. Operations for templates, inventories, keys, repositories, and schedules live on their dedicated sub-clients, not on `Project`. When grepping for a new operation, search by HTTP path (e.g. `PostProjectProjectIDInventory`) rather than guessing the sub-client.

The provider supports `tls_skip_verify` for self-signed TLS; if set, `Configure` constructs an `http.Client` with `InsecureSkipVerify` and hands it to `httptransport.NewWithClient`. The host string omits the port when it matches the scheme's default (`:443` for https, `:80` for http) — some upstream proxies reject SNI/Host headers that include the default port (issue #56).

### Import IDs

Nested resources use slash-delimited compound IDs like `project/1/template/2`. `internal/provider/import.go` `parseImportFields` parses these via a `(\w+)/(\d+)` regex into a `map[string]int64`, and resources call it from `ImportState`. Each `examples/resources/<name>/import.sh` documents the format.

### Nil-handling pattern

Several upstream API responses return `nil` for fields that should be zero/false (a known SemaphoreUI quirk — see commit `b345643` on `max_parallel_tasks`). When mapping API responses to Terraform models, explicitly check for `nil` pointers and substitute the zero value rather than using `types.Int64PointerValue` directly. `convertProjectResponseToProjectModel` in `project_resource.go` is the canonical example.

### Nullability patches in `api-docs.yml`

The local `api-docs.yml` is a *patched* copy of the upstream spec from a tagged release (currently `v2.18.2`). Upstream tends to drop nullable annotations from fields it considers always-set, but the Semaphore API genuinely returns `null` for several optional fields. The patches re-add `x-nullable: true` for those:

- `Project.alert_chat` / `ProjectRequest.alert_chat`
- `ViewRequest.id` — upstream omits it, but the PUT views endpoint returns 400 without it

When bumping `api-docs.yml`, re-import upstream verbatim first (one commit), then re-apply the nullability patches based on which tests fail (a follow-up commit). The two-commit split keeps the diff legible — reviewers can see what came from upstream versus what we patched locally.

### Template `environment_id` quirk

SemaphoreUI v2.16+ introduced `environment_ids` (array) and stopped returning the singular `environment_id` on GET (always 0). The `convertProjectTemplateModelToTemplateRequest` converter writes to both fields on requests; `convertTemplateResponseToProjectTemplateModel` reads from `environment_ids[0]` with a fallback to `environment_id`. The Terraform schema's `environment_id` attribute is unchanged — the translation is invisible to users.

### Environment secret update gotcha

The Semaphore API does not honor type changes on secret update operations — only `name` and `secret` (value) are persisted. `convertProjectEnvironmentModelToEnvironmentRequest` sets the `Secret` field when marking an update; without it, the API treats the update as a no-op for the value (issue #68 / PR #74). Type changes are not supported in-place; the resource's acceptance tests reflect this.

## Regenerating code

- **API client** (after `api-docs.yml` changes): `task client`. Requires `swagger` binary from `go-swagger` (`go install github.com/go-swagger/go-swagger/cmd/swagger@latest`).
- **Docs** (after schema changes): `task generate`. Requires `terraform` in PATH (for `terraform fmt`). CI's `generate` job fails if the resulting diff isn't committed.

## Tooling notes

- Go version: 1.26.3 (see `.tool-versions` and `go.mod`).
- Linter: `golangci-lint` v2, config at `.golangci.yml` — `forcetypeassert`, `errcheck`, `staticcheck`, etc. enabled; tests excluded from lint.
- Pre-commit hooks (`.pre-commit-config.yaml`) run golangci-lint, end-of-file-fixer, and `terraform fmt`. **Pre-commit lint runs on every commit** — if the code doesn't compile or lint, the commit is blocked. This means broken-intermediate-state commits aren't possible; bundle dependent changes (e.g. client regen + provider reconciliation) into one commit.
- Documentation generation is in a separate module (`tools/go.mod`) so tfplugindocs dependencies don't bloat the main module.
- Dependabot auto-merge (`.github/workflows/dependabot-auto-merge.yml`) auto-approves and squash-merges patch/minor/security PRs but leaves majors for human review. For grouped PRs, `fetch-metadata` reports the highest semver bump — any group containing a major won't qualify.
