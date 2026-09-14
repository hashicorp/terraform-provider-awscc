<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Generating the Provider with bigdiffer

This is the weekly runbook for generating the Terraform AWS Cloud Control
Provider from the live CloudFormation registry using `bigdiffer`
(`internal/tools/bigdiffer`). For why it's built this way, see
[bigdiffer-design.md](bigdiffer-design.md).

<!--mdtoc: begin-->
* [The weekly release](#the-weekly-release)
    * [1. Setup](#1-setup)
    * [2. Update](#2-update)
    * [3. Build and smoke test](#3-build-and-smoke-test)
    * [4. Documentation](#4-documentation)
    * [5. CHANGELOG and version](#5-changelog-and-version)
    * [6. Commit and open a pull request](#6-commit-and-open-a-pull-request)
* [Reading the report](#reading-the-report)
* [Full regeneration (offline)](#full-regeneration-offline)
* [Fallback: the legacy process](#fallback-the-legacy-process)
* [How it works](#how-it-works)
<!--mdtoc: end-->

## The weekly release

> [!TIP]
> `make bigdiffer-update`, `make bigdiffer-docs`, and `make bigdiffer-generate` are
> shortcuts for the `go run ./internal/tools/bigdiffer …` commands used below.

### 1. Setup

```sh
git switch main
git pull
make tools
make newbranch
export AWS_DEFAULT_REGION=us-east-1   # bigdiffer always queries us-east-1
```

Ensure valid AWS credentials are set in the environment.

### 2. Update

```sh
go run ./internal/tools/bigdiffer -update
```

This refreshes the schema cache, tracks new CloudFormation types, regenerates
every changed resource and data source, and rewrites `all_schemas.hcl` — all in
one pass, taking roughly 13 minutes (dominated by the AWS discovery crawl; a
progress bar shows it). It prints a [report](#reading-the-report) of what
changed and what, if anything, was frozen or suppressed.

Review the report together with `git diff`. For each frozen/suppressed type,
open a GitHub issue with the reason (see
[example issue #2070](https://github.com/hashicorp/terraform-provider-awscc/issues/2070));
`internal/update/suppressions_checkout.txt` pins are still honored.

Files this changes:

* `internal/service/cloudformation/schemas/*.json` — refreshed schema cache
* `internal/provider/all_schemas.hcl` — new blocks + policy annotations
* `internal/aws/**/*_gen.go` — regenerated code for changed types only
* `internal/provider/import_examples_gen.json` — import-examples aggregate
* `internal/provider/registrations_gen.go` — the blank-import registration file

> [!NOTE]
> `registrations_gen.go` supersedes the legacy
> `internal/provider/{resources,singular_data_sources,plural_data_sources}.go`
> directive files, which bigdiffer does **not** touch. Both may be present
> during the transition — duplicate blank imports are legal Go. Commit
> `registrations_gen.go`. `bigdiffer -check` (run in CI) re-emits this file and
> fails if the committed copy is stale.

### 3. Build and smoke test

```sh
make build
make smoke
```

`-update` already compiled the generated code before writing it, so `make
build` here is a confirmation, and `make smoke` adds test compilation and
acceptance smoke tests. If `make build` fails here, investigate rather than
hand-suppressing — `-update` would already have suppressed a genuinely
non-compiling type.

### 4. Documentation

Requires Terraform `v1.14`+ and `tfplugindocs` on `PATH` (both installed by
`make tools`):

```sh
go run ./internal/tools/bigdiffer -docs
```

This regenerates import-example docs, runs `terraform fmt`, and runs
`tfplugindocs generate`. Changes `examples/**` and `docs/**`.

### 5. CHANGELOG and version

`-update` already drafted the `FEATURES:` bullets into `CHANGELOG.md`'s top,
in-progress version block (e.g. `## 1.100.0 (Unreleased)`). By hand:

1. Add the PR number link
   (`([#1234](https://github.com/hashicorp/terraform-provider-awscc/pull/1234))`)
   once the PR exists.
2. Add any `NOTES:` or breaking-change entries.
3. Update `version/VERSION` to match.

If this run promoted nothing newly user-visible, no `FEATURES:` section is
added.

`-update` assumes the top block's `FEATURES:` section is empty or absent when
it runs and errors instead of guessing at a merge if it already has bullets —
add the new bullets by hand in that case.

### 6. Commit and open a pull request

Review `git status` and `git diff`, then commit — either as a single commit or
grouped logically, for example:

```sh
git add internal/service/cloudformation/schemas internal/provider/all_schemas.hcl
git commit -m "Refresh CloudFormation schemas"

git add internal/aws internal/provider/registrations_gen.go internal/provider/import_examples_gen.json
git commit -m "Regenerate resources and data sources"

git add examples docs CHANGELOG.md version/VERSION
git commit -m "Regenerate documentation and update changelog"
```

Open a pull request and verify CI passes; once merged, cut the release.

## Reading the report

`-update` narrates each phase (`==>` lines) and prints a reconciliation report to
stderr:

* **Detected changes — new / changed / unchanged / frozen / missing.** New and
  Changed are regenerated; Unchanged are skipped; Frozen (pinned via
  `frozen_since`) are never re-evaluated; Missing means discovery could not
  describe the type.
* **Per-type lines** for anything frozen or suppressed, with the policy summary
  and reason — these are your issue-filing worklist.
* **Anomalies** — retained-but-unexplained rows (in the overlay, gone from AWS,
  with no `frozen_since`/`non_provisionable`/checkout pin), duplicate blocks, and
  naming-invariant violations. Investigate these before merging.

## Full regeneration (offline)

To regenerate the **entire** corpus from the committed overlay and schema cache
without touching AWS — for example after changing the generation engine — use:

```sh
go run ./internal/tools/bigdiffer -generate
```

This writes every `*_gen.go`, `registrations_gen.go`, and
`import_examples_gen.json` in parallel, in seconds. It does not query AWS and
does not change `all_schemas.hcl`.

## Fallback: the legacy process

The legacy machinery is intact and unaffected by bigdiffer; it produces the same
result. If bigdiffer is ever unavailable or misbehaving, follow
[generating-the-provider.md](generating-the-provider.md) instead
(`make cleanschemas suppressions schemas`, `make biglister`, `make resources`,
`make singular-data-sources plural-data-sources`, `make docs-all`, …).

Interoperating between the two:

* The **schema cache** (`internal/service/cloudformation/schemas/*.json`) and the
  **overlay** (`all_schemas.hcl`) use the identical on-disk format, so a cache and
  overlay produced by bigdiffer are consumed unchanged by the legacy `make`
  targets, and vice versa. No conversion is needed.
* Before committing a legacy-only run, delete bigdiffer's registration file so
  registration comes solely from the regenerated legacy directive files:

  ```sh
  rm -f internal/provider/registrations_gen.go
  ```

  `make resources` regenerates `internal/provider/resources.go` (and the
  data-source equivalents) complete with every current type, so the legacy path is
  self-contained.

> [!NOTE]
> The legacy `make bigdiffer` target is an unrelated helper that `diff`s dated
> `available_schemas.<date>.hcl` snapshots — it is **not** the
> `internal/tools/bigdiffer` tool described here.

## How it works

Briefly: bigdiffer reconciles the live CloudFormation registry against
`internal/provider/all_schemas.hcl` in memory, regenerates only what changed,
compile-gates the result before writing anything, and lets one policy decide
what to freeze or suppress so a broken schema never blocks a release.

For the full design — why it's built this way, what it replaces, and what's
still open — see [bigdiffer-design.md](bigdiffer-design.md).
