# Repo Refresh Milestones (2026)

Working doc — gitignored. Drives the `chore/repo-refresh-2026` branch.

Snapshot taken 2026-05-11. Current release: `v1.4.1` (2025-06-04). Latest upstream SemaphoreUI: `v2.18.2` (2026-05-07). CI matrix currently tests against `v2.12.17 / v2.13.15 / v2.14.12` — almost a year behind.

**Positioning shift.** After this refresh the repo will be marked **AI-supported, not actively maintained** — same model as [`CruGlobal/js-hcl2`](https://github.com/CruGlobal/js-hcl2). The README gets a status callout, contribution norms move to `CONTRIBUTING.md`, and a Dependabot auto-merge workflow keeps minor/patch bumps and security advisories flowing without human review. Major bumps and feature work remain best-effort. This positioning has implications across milestones:

- **M1** ships the dependabot auto-merge workflow as part of the dep-hygiene baseline (so once `main` is current, it stays current on its own).
- **M5** adds the README status callout + `CONTRIBUTING.md` and updates the GitHub repo description / topics.

The plan is 5 sequential milestones. Land each as one PR (or a tight series) so the changelog stays legible. Conventional Commit prefixes matter — release-please reads them.

---

## Milestone 1 — Toolchain & dependency refresh (no behavior change)

**Goal:** get the repo onto modern Go + clean baseline before touching API code. Should be a pure no-op for users.

### Tasks

- [ ] Bump Go to **1.26.3** in `go.mod`, `.tool-versions`, and any GitHub Actions inputs. Verify `actions/setup-go` reads `go-version-file: 'go.mod'` (it does — `.github/workflows/test.yml:24`, `release-build.yml:23`).
- [ ] Bump `golangci-lint` to v2 latest in `.pre-commit-config.yaml` and the action (currently v8 in test.yml; dependabot PR #70 proposes v9).
- [ ] Resolve the **pending dependabot PRs** as a batch — easier to do here than carry forward:
  - #67 `gomod-updates` group (9 updates incl. `terraform-plugin-framework 1.15.0 → 1.16.1`, `terraform-plugin-framework-validators 0.18.0 → 0.19.0`, `terraform-plugin-testing 1.13.1 → 1.13.3`, `go-openapi/runtime 0.28.0 → 0.29.0`, `go-openapi/swag 0.23.1 → 0.25.1`). Just rebase + accept.
  - #72 `golang.org/x/crypto 0.38.0 → 0.45.0`
  - #71 same crypto bump for `tools/go.mod`
  - #52 `cloudflare/circl 1.6.0 → 1.6.1`; #51 same for `tools/`
  - #73 `actions/checkout 4 → 6`; #64 `actions/setup-go 5 → 6`; #62 `amannn/action-semantic-pull-request 5 → 6`; #70 `golangci-lint-action 8 → 9`
- [ ] Run `task generate` and confirm `docs/` diff is empty (CI's generate job enforces this).
- [ ] Verify `task lint` and `task testacc` pass under the bumped `SEMAPHORE_VERSION` matrix (see Milestone 2 for the matrix bump).
- [ ] Update `terraform` version in `.tool-versions` (`1.11.3 → 1.13.x`) — CI already uses 1.12.x.

### Dependabot auto-merge (set up while we're in the dep-hygiene mindset)

Pattern lifted from `CruGlobal/js-hcl2`. With the repo about to be marked not-actively-maintained, auto-merge is what keeps `main` patched without a human in the loop.

- [ ] Add `.github/workflows/dependabot-auto-merge.yml` (cribbed from js-hcl2):
  - Trigger on `pull_request` (`opened, reopened, synchronize`).
  - Permissions: `contents: write`, `pull-requests: write`.
  - Guard: `if: github.actor == 'dependabot[bot]'`.
  - Use `dependabot/fetch-metadata@v3` to read `update-type` and `alert-state`.
  - **Eligible**: any security advisory (`alert-state` set), or `version-update:semver-patch`, or `version-update:semver-minor`. For grouped PRs, `update-type` reflects the **highest** bump in the group — so any group containing a major won't auto-merge. That's the desired behavior.
  - **Not eligible**: `version-update:semver-major` — left for human review.
  - On eligible: `gh pr review --approve` (runs as `github-actions[bot]`, not dependabot, so it satisfies "no self-approval" rules), then `gh pr merge --squash --auto`.
- [ ] Confirm or add a **branch ruleset / branch protection** on `main` requiring CI checks (`build`, `generate`, the matrix `test` jobs). `--auto` is a no-op without it — GitHub falls back to immediate merge once auto-merge is "enabled" if no checks are required. The matrix tests are the real gate; without them, a broken patch bump could merge.
- [ ] Verify the **repo squash-merge setting**: auto-merge with `--squash` requires squash merging to be enabled in repo settings. Likely already on (release-please workflow expects clean commit history).
- [ ] Audit `.github/dependabot.yml`:
  - Current grouping (`gomod-updates` and `action-updates`, both `minor + patch` only) is good — majors stay ungrouped so they get individual PRs with proper migration context.
  - The `tools/` directory has its own `go.mod` but **isn't** in the dependabot config — that's why we have orphan PRs like #71 and #51 for `golang.org/x/crypto in /tools`. Add a third entry for `directory: "/tools"`.
  - Consider raising `open-pull-requests-limit` (default 5) if grouping is working well — fewer drive-by deps drift.

**Commit style:** `chore(deps): bump go to 1.26.3 and refresh dependencies`. Use one or two PRs max; don't fragment. Auto-merge workflow can be a separate `ci: add dependabot auto-merge workflow` PR if cleaner.

**Risks:** terraform-plugin-framework minor bumps occasionally tighten validation — run the full acceptance suite, not just unit tests. For auto-merge: a buggy patch release can land unattended; the CI matrix is the only safety net, so don't skip the branch-protection step.

---

## Milestone 2 — Bump tested SemaphoreUI matrix

**Goal:** confirm the provider still works against current SemaphoreUI before any client regen.

### Tasks

- [ ] Update `.github/workflows/test.yml` matrix from `v2.12.17 / v2.13.15 / v2.14.12` → `v2.16.x / v2.17.x / v2.18.2` (latest 3 minor lines). Confirm pin choices against the upstream release list.
- [ ] Update `Taskfile.yml` `SEMAPHORE_VERSION` default to the latest stable (`v2.18.2`).
- [ ] Run `task testacc` locally against each version. **Expect failures.** Catalogue them — they motivate Milestone 4.
- [ ] Known failure-likely areas based on inbound issues:
  - **#55** — 403 on `semaphoreui_projects` / `semaphoreui_project` against v2.15.0+. Reproducible? May be auth-related (token scopes changed) or an API contract change.
  - **#30 (CLOSED)** — v2.13.1 compatibility was the previous version-skew incident. Useful as a recipe for how the fix went.
- [ ] If matrix is fully red, narrow to a single version (latest) for this PR and treat the older breaks as known regressions to be fixed in Milestone 4. Don't ship a broken `main`.

**Commit style:** `ci: test against semaphoreui v2.16/v2.17/v2.18`.

---

## Milestone 3 — API client regeneration

**Goal:** rebase our patched `api-docs.yml` onto the upstream `v2.16.14`-era spec.

### Background — what diverges

The local `api-docs.yml` is `2.14.0-beta1` with several hand-edits to coerce nullability. Upstream (currently `2.16.14` on develop) moved in the **opposite direction**: it removed most `type: [string, 'null']` unions and `x-nullable: true` annotations.

Concretely, the divergence we know about:

1. **Nullable string fields** that the local fork annotates with `type: [string, 'null']`, but upstream declares as plain `type: string`:
   - `ProjectRequest`, `Project`: `alert_chat`, `type`
   - `Project`: `alert`, `max_parallel_tasks` had `x-nullable: true`
   - `ProjectBackup` nested: `alert_chat`, `type`, `arguments`, `cron`, `build_template`, `survey_vars`, `start_version`, `vault_key`, `ssh_key`, `become_key`, `password`, `env`
2. The local `convertProjectResponseToProjectModel` (and similar map functions) substitute zero values when pointers are nil — this only matters because the nullable annotation made fields pointers. If we accept upstream's non-nullable types, these become non-pointers and the workaround is moot. **But** the Semaphore API genuinely returns nulls — see the `max_parallel_tasks` fix in commit `b345643`. Need to verify which way is correct against a live API.
3. **New definitions** in upstream we'd inherit: `AcceptInviteRequest`, `AnsibleTaskParams`, `IntegrationAlias`, `ProjectInvite`, `ProjectInviteRequest`, `TaskPrams` (sic), `TerraformTaskParams`.
4. **New paths** in upstream:
   - `/project/{project_id}/integrations/{integration_id}/aliases` + `/{alias_id}`
   - `/project/{project_id}/integrations/aliases` (top-level alias index) + `/{alias_id}`
   - `/project/{project_id}/notifications/test`
   - `/project/{project_id}/templates/{template_id}/stop_all_tasks`
5. **Removed:** `/debug/gc` — fine to drop.
6. **New field:** `APIToken.name`.
7. **Integration request/response** gains `task_params` / `params` referencing `TaskPrams` (a union of `AnsibleTaskParams` / `TerraformTaskParams`).

### Tasks

- [ ] Take upstream `develop` `api-docs.yml` as the new baseline. Drop into the repo verbatim first; commit that as a discrete starting point so the patch diff is reviewable.
- [ ] **Verify against live API**: spin up SemaphoreUI v2.18.2, hit each endpoint we use, and confirm whether nullable fields actually come back as `null`. If they do, re-apply our nullability patches *on top of* the new upstream file. If they don't, drop the patches.
  - Test specifically: `GET /projects` (alert, alert_chat, type, max_parallel_tasks); `GET /project/{id}` same; `GET /project/{id}/templates/{id}` (arguments, cron, vault_key, start_version, survey_vars, build_template); `GET /project/{id}/inventory/{id}` (ssh_key, become_key); `GET /project/{id}/environment/{id}` (password, env).
- [ ] Run `task client` to regenerate `semaphoreui/client/` and `semaphoreui/models/`. Requires `go-swagger` binary — confirm install instructions in CLAUDE.md if missing.
- [ ] Hand-fix compile errors in `internal/provider/` from changed model shapes. Most likely: `models.Project.AlertChat *string → string`, `models.Project.MaxParallelTasks *int64 → int64`, etc. Update `convertXxxResponseToXxxModel` helpers accordingly — the existing nil-check pattern is the canonical example to copy from.
- [ ] Run unit + acceptance tests; iterate until green.

**Commit style:** two commits — first `chore: import upstream api-docs v2.18.x as new baseline`, then `chore(client): regenerate API client and reconcile resources`. Don't squash them — preserves the patch boundary.

**Risk hotspot:** the nullability question. Getting it wrong means either (a) compile errors, (b) silent loss of fields, or (c) "drift on every plan" UX bugs. Test against a real server, don't reason from the spec alone.

---

## Milestone 4 — Resource/data-source updates + fixes

**Goal:** repair anything broken by the client regen, ship the open-issue fixes, and add high-value new resources for the new API surface. Maintain Terraform state compatibility — schema changes that break state must include state-migration logic.

### Backwards-compatibility doctrine for this milestone

- **State migration**: any time a `tfsdk:` field changes shape or name, implement `resource.ResourceWithUpgradeState`. Don't ship a rename without a migration.
- **Required → optional**: free move, no migration needed.
- **Optional → required**: breaking, avoid. Prefer making the field optional + validating at plan-time via a custom validator.
- **Adding new attributes**: free if they have sensible computed defaults.
- **Removing attributes**: breaking. Deprecate first via `DeprecationMessage` on the schema, then remove in the next major.
- If a breaking change is unavoidable, bump major version (`v2.0.0`) and document the migration in CHANGELOG.

### Repairs (driven by open issues)

- [ ] **#68 / PR #74 / PR #69 — project_environment secret values not updating.** PR #74 is the minimal fix (6 lines, set the `Secret` field when marking `operation: update`). PR #69 is a more thorough rewrite of `convertProjectEnvironmentModelToEnvironmentRequest`. Both were authored externally and not merged. Cherry-pick PR #74 first as the safe fix; evaluate PR #69's broader refactor on top — it also handles `create` / `delete` ops more explicitly.
  - Add acceptance test: rotate a secret value, apply, confirm the new value reaches the API (mitmproxy-style; or assert via direct API check after apply).
- [ ] **#56 / PR #57 — `:443` in host header breaks HTTPS through some proxies.** PR #57 is 6 lines: skip the `:443` / `:80` suffix when it's the default for the scheme. Land it.
- [ ] **#55 — 403 on `semaphoreui_projects` data source against v2.15.0+.** Reproduce first. Hypothesis: a permission scope tightening upstream, or an Accept-header / pagination contract change. May be partially or fully resolved by Milestone 3.
- [ ] **#26 — `semaphoreui_project_template` shouldn't require `playbook` / `inventory_id` when `app = "terraform"`.** Validate the current upstream API: if `inventory_id` is now genuinely optional, drop Required and add a validator that requires it for non-terraform apps. Keep `playbook` permissive (a workaround with `noop.yml` already exists; if API allows empty, allow empty).
- [ ] **#44 (already closed without action)** — `tfplugindocs` formats nested attributes as `Attributes`/`Attributes List` instead of `Block`/`Block Set`. Confirm whether newer tfplugindocs has options here; otherwise close as wontfix and document.

### New features

- [ ] **PR #75 — OpenTofu workspace inventory.** External PR was closed; the diff is good (193 +/21 −, tests included). Resurrect it. Adds `tofu_workspace` to `semaphoreui_project_inventory` mapping to API type `tofu-workspace`. Pure addition — no migration needed.
- [ ] **#53 — Ansible `tag` / `skip_tag` on `semaphoreui_project_template`.** Maps to the new `AnsibleTaskParams` definition (`tags`, `skip_tags`, etc.) in upstream api-docs. Nested attribute under `task_params` / similar; design to mirror upstream's `TaskPrams` union (different shape for Ansible vs Terraform).
- [ ] **#58 — Ephemeral / write-only `ssh.private_key`.** Requires `terraform-plugin-framework 1.11+` (write-only attribute support). After Milestone 1's bump we're on 1.16, so feasible. Add `WriteOnly: true` to sensitive key inputs on `semaphoreui_project_key` (ssh.private_key, login_password.password). Backwards-compat: keep the old `Sensitive: true` path; gate write-only on the new attribute name (e.g. `ssh.private_key_wo`).
- [ ] **New resource: `semaphoreui_integration_alias`** — maps to `/project/{project_id}/integrations/{integration_id}/aliases` and `/{alias_id}`. New in upstream spec.
- [ ] **New resource: `semaphoreui_project_integration`** — we already have client code for integrations but no TF resource exposing it. Verify whether this is a deliberate omission or just unfinished; if no objection, add it.
- [ ] **New resource: `semaphoreui_project_invite`** — maps to the new `ProjectInvite` definition. Lower priority; only valuable if invites are how users onboard in your environments.

### Non-actionable / wontfix candidates

- **#36 — drift detection when resources deleted out-of-band.** Maintainer position (Brian's comment): this is correct behavior — Terraform should error, not silently re-create. Reaffirm and close.

**Commit style:** one PR per logical change. `fix: project_environment secret values now updated correctly (#68)`, `feat: add OpenTofu workspace inventory (#75)`, etc. Reference the issue number in the commit body, not the title (release-please dedupes on title).

---

## Milestone 5 — Documentation, positioning & release

**Goal:** make the result discoverable, shipped, and clearly marked as best-effort going forward.

### Tasks

- [ ] Regenerate docs (`task generate`) — CI gates on this anyway.
- [ ] Update `README.md`:
  - **Add a top-of-file status callout** (blockquote style, mirroring `CruGlobal/js-hcl2`'s README). Wording template:
    > **Status: AI-supported, not actively maintained.** This provider was authored for an internal use case and is now maintained on a best-effort basis with AI assistance. Dependabot keeps dependencies and security advisories up to date automatically (patch + minor bumps auto-merge; majors require manual review). Feature work, bug fixes, and other changes happen on a best-effort basis. **Pull requests and issues are welcome** — they may take time to be reviewed. See [`CONTRIBUTING.md`](CONTRIBUTING.md) for the contribution workflow.
  - Bump the "tested against latest 3 versions" line to whatever the new matrix is.
  - If any new resources landed, mention them or link to the registry.
  - Replace the existing "Support" section (currently: "developed for internal use, not actively maintained, contributions welcome") — the new status callout supersedes it.
- [ ] **Add `CONTRIBUTING.md`** modeled on `CruGlobal/js-hcl2/CONTRIBUTING.md`. Adapt for Go/Terraform-provider specifics:
  - Toolchain pinned via `.tool-versions` + `asdf`; the same file drives `actions/setup-go` (`go-version-file: 'go.mod'`).
  - Local check before PR: `task fmt && task lint && task test && task testacc`.
  - Conventional Commits required (release-please consumes them).
  - For bug fixes / features: failing-test-first norm. Acceptance tests (`TestAcc_*`) cover the API contract; unit tests cover pure helpers. Exemptions for docs-only, CI-only, and dependency PRs.
  - Link back to upstream SemaphoreUI API docs as the authoritative spec — and clarify that the local `api-docs.yml` will diverge when needed, with rationale captured in commit messages.
- [ ] Update **GitHub repo metadata**:
  - Repo description: add "AI-supported · not actively maintained" prefix or suffix so it's visible from the org page.
  - Topics: ensure `terraform-provider`, `semaphoreui`, plus a marker topic if the org uses one (e.g. `ai-supported`).
- [ ] If any user-facing breaking changes accumulated (most likely: state shape changes from Milestone 3), bump to **v2.0.0** via release-please. Add a "Migrating from v1.x" section to CHANGELOG. Each breaking commit must use `feat!:` / `fix!:` or include a `BREAKING CHANGE:` footer — release-please uses these to compute the major bump.
- [ ] If everything stayed backwards-compatible, the chain of `feat:` / `fix:` commits will release-please into `v1.5.0` (or wherever). Fine — no special action needed beyond merging release-please's PR.
- [ ] Update the **acceptance test setup script** (`scripts/setup_test_env.sh`) if the `user__token` table schema changed in newer Semaphore versions. Worth checking — the table layout has been stable but it's a single line we'd find out about on the first CI run of Milestone 2.

**Commit style:** `docs: mark provider as AI-supported and add CONTRIBUTING guide` for the positioning piece; separate from any feature/fix commits so release-please doesn't confuse it with a feature bump.

---

## Quick-reference: external PRs to evaluate

| PR | Status | Author | Disposition |
|----|--------|--------|-------------|
| #74 | OPEN | kumy | Land as-is in M4 (fixes #68) |
| #57 | OPEN | maledb404 | Land as-is in M4 (fixes #56) |
| #75 | CLOSED | SiM22 | Resurrect in M4 (tofu workspace) |
| #69 | CLOSED | lepichon | Evaluate as follow-up to #74 — broader env-secret refactor |

## Quick-reference: open issues

| # | Title | Disposition |
|---|-------|-------------|
| 68 | Env secret values not updated | M4 (PR #74) |
| 58 | Ephemeral SSH private_key | M4 (write-only attr) |
| 56 | Connection reset with :443 | M4 (PR #57) |
| 55 | 403 on projects data source v2.15.0+ | M2 reproduce, M4 fix |
| 53 | Ansible tag / skip-tag | M4 (task_params) |
| 36 | Out-of-band delete = drift | Close as wontfix |
| 26 | template playbook/inventory_id should be optional | M4 |
