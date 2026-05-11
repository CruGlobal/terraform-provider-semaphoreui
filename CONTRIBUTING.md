# Contributing to `terraform-provider-semaphoreui`

Thanks for your interest in contributing. This provider is maintained on
a best-effort basis with AI assistance (see the status callout in the
[README](README.md)). Pull requests and issues are welcome, but reviews
may take time.

## Toolchain

This project pins exact tool versions in `.tool-versions` so local
development and CI resolve the same binaries from a single source.

```
golang        1.26.3
terraform     1.15.2
golangci-lint 2.12.2
pre-commit    4.2.0
```

[`asdf`](https://asdf-vm.com/) reads `.tool-versions` directly; CI's
`actions/setup-go@v6` reads `go-version-file: 'go.mod'`, which agrees
with `.tool-versions`. Install `asdf` and the relevant plugins, then:

```sh
asdf install
go mod download
pre-commit install
```

Do **not** bypass `asdf` (e.g. Homebrew, GVM) — divergent toolchain
versions produce hard-to-debug failures.

## Running tests

The repo uses [Task](https://taskfile.dev) (`Taskfile.yml`):

- `task lint` — `golangci-lint` v2 (tests excluded; see `.golangci.yml`)
- `task test` — unit tests, fast
- `task testacc` — acceptance tests against a Dockerized SemaphoreUI
  (orchestrates `docker compose up`, seeds an API token directly into
  the MySQL `user__token` table, runs the suite, tears down)
- `task generate` — regenerates `docs/` via `tfplugindocs`; CI fails if
  the diff is non-empty

Run a single acceptance test:

```sh
SEMAPHORE_VERSION=v2.18.2 task testacc -- -run TestAcc_ProjectResource_basic
```

Tests prefixed `TestAcc_` require the live API. The matrix in
`.github/workflows/test.yml` exercises the three most recent SemaphoreUI
minor lines.

## Failing-test-first for bug fixes and features

**Bug fixes and feature PRs should lead with a failing test.** Structure
your commits so the PR history reads:

1. **First commit** — adds an acceptance test (or unit test) that
   reproduces the bug or asserts the new feature's behavior. CI should
   fail on this commit alone.
2. **Subsequent commits** — the fix or implementation, with the test
   flipping from red to green.

For **issue reports**: include a minimal Terraform config that
reproduces the problem and the exact `terraform apply` output. Issues
with reproducible repro steps are triaged first.

### Exemptions

The failing-test-first requirement does **not** apply to:

- **Documentation-only changes** (`README.md`, `docs/**`, `CLAUDE.md`,
  `CONTRIBUTING.md`).
- **CI / workflow changes** (`.github/**`).
- **Dependency updates** (Dependabot PRs, manual `go get` bumps).
- **Generated-code changes** (`semaphoreui/client/**`, `semaphoreui/models/**`)
  — these come from `task client` and should always be accompanied by
  the matching `api-docs.yml` change.

## Commit messages

Releases are managed by
[release-please](https://github.com/googleapis/release-please), which
parses [Conventional Commit](https://www.conventionalcommits.org/)
messages to compute version bumps and assemble the changelog. The
GitHub Actions title check (`amannn/action-semantic-pull-request`)
enforces the format on every PR.

Common prefixes:

- `feat:` — new resource, new attribute, or other user-visible feature
- `fix:` — bug fix
- `chore:` / `ci:` / `docs:` — no user-visible change
- `feat!:` / `fix!:` or a `BREAKING CHANGE:` footer — major-version bump

## Regenerating the API client

The SemaphoreUI API client is generated from `api-docs.yml`. When the
upstream API changes, bump `api-docs.yml` and regenerate:

```sh
# Install go-swagger if missing
go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# Regenerate semaphoreui/client/ and semaphoreui/models/
task client

# Run lint + acceptance tests against each matrix version
task lint
SEMAPHORE_VERSION=v2.16.51 task testacc
SEMAPHORE_VERSION=v2.17.39 task testacc
SEMAPHORE_VERSION=v2.18.2  task testacc
```

The local `api-docs.yml` is a *patched* copy of upstream. When
re-importing from a new upstream tag, **do the import as a discrete
commit first**, then apply nullability patches as a follow-up commit
based on which acceptance tests fail. This keeps the diff legible —
reviewers can see what came from upstream versus what we patched
locally. The current re-applied patches are documented in
[`CLAUDE.md`](CLAUDE.md).

## Adding a new resource or data source

Each Terraform type follows a three-file convention in
`internal/provider/`:

- `<name>_schema.go` — `<Name>Model` struct + `<Name>Schema()` returning
  a `superschema.Schema` (one definition, both resource and data-source
  variants via `Resource:` / `DataSource:` / `Common:` overrides).
- `<name>_resource.go` — `resource.Resource` implementation
  (`Configure`, `Schema`, `Create`, `Read`, `Update`, `Delete`,
  `ImportState`).
- `<name>_data_source.go` — `datasource.DataSource` implementation.

A new resource requires:

1. Schema + resource + (optional) data source files.
2. Registration in `provider.go` `Resources()` / `DataSources()`.
3. An `examples/resources/<name>/` directory with `resource.tf` and
   `import.sh`.
4. Acceptance tests (`<name>_resource_test.go`) following the existing
   patterns (see `project_inventory_resource_test.go` for a thorough
   example).
5. `task generate` to produce the matching `docs/` page.

## License

By contributing, you agree that your contributions will be licensed
under this repo's [LICENSE](LICENSE).
