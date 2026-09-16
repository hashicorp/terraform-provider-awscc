<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Generating the Provider with bigdiffer

This is the weekly runbook for generating the Terraform AWS Cloud Control
Provider from the live CloudFormation registry using `bigdiffer`
(`internal/tools/bigdiffer`). For what bigdiffer is and the available commands,
see [the tool's README](../../internal/tools/bigdiffer/README.md); for why it's
built this way, see [bigdiffer-design.md](bigdiffer-design.md).

<!--mdtoc: begin-->
* [The weekly release](#the-weekly-release)
    * [1. Setup](#1-setup)
    * [2. Sync](#2-sync)
    * [3. Build and smoke test](#3-build-and-smoke-test)
    * [4. Documentation](#4-documentation)
    * [5. CHANGELOG and version](#5-changelog-and-version)
    * [6. Commit and open a pull request](#6-commit-and-open-a-pull-request)
* [Reading the report](#reading-the-report)
* [Full regeneration, offline (-reconcile)](#full-regeneration-offline--reconcile)
* [Checking an engine change, offline and read-only (-check)](#checking-an-engine-change-offline-and-read-only--check)
* [Fallback: the legacy process](#fallback-the-legacy-process)
* [How it works](#how-it-works)
<!--mdtoc: end-->

## The weekly release

> [!TIP]
> `make bigdiffer-sync`, `make bigdiffer-docs`, and `make bigdiffer-reconcile` are
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

### 2. Sync

```sh
go run ./internal/tools/bigdiffer -sync
```

This refreshes the schema cache, tracks new CloudFormation types, gates and
regenerates the **whole** corpus (not just changed types — every type is
generate-and-compile-checked every run, so a machinery change surfaces
immediately instead of aging silently), and rewrites `all_schemas.hcl` — all in
one pass, taking roughly 13 minutes (dominated by the AWS discovery crawl; a
progress bar shows it). It prints a [report](#reading-the-report) of what
changed and what, if anything, was frozen or suppressed.

A schema-changed type that fails generation/compile freezes or suppresses just
that type, same as before. A **schema-unchanged** type that fails means the
generation machinery itself regressed — a template, `codegen/`, or a Go
toolchain change — so `-sync` aborts the **entire run** and promotes nothing at
all, not even the types that compiled cleanly this run, and reports exactly
which types and artifacts failed. That is not a bug to work around: an engine
regression makes every type it touched this run untrustworthy, not just the
one that happened to fail to compile. Fix or revert the machinery change, then
re-run.

Review the report together with `git diff`. For each frozen/suppressed type,
open a GitHub issue with the reason (see
[example issue #2070](https://github.com/hashicorp/terraform-provider-awscc/issues/2070));
`internal/update/suppressions_checkout.txt` pins are still honored.

Files this changes:

* `internal/service/cloudformation/schemas/*.json` — refreshed schema cache
* `internal/provider/all_schemas.hcl` — new blocks + policy annotations
* `internal/aws/**/*_gen.go` — regenerated code for every type the gate
  promotes this run (new/changed types, plus any type whose output changed
  from a machinery fix — schema-unchanged types otherwise regenerate
  byte-identical and are a no-op)
* `internal/provider/import_examples_gen.json` — import-examples aggregate
* `internal/provider/registrations_gen.go` — the blank-import registration file
* `CHANGELOG.md` — a drafted `FEATURES:` block for newly user-visible artifacts

> [!NOTE]
> `registrations_gen.go` supersedes the legacy
> `internal/provider/{resources,singular_data_sources,plural_data_sources}.go`
> directive files, which bigdiffer does **not** touch. Both may be present
> during the transition — duplicate blank imports are legal Go. Commit
> `registrations_gen.go`. `bigdiffer -lint` (run in CI) re-emits this file and
> fails if the committed copy is stale.

### 3. Build and smoke test

```sh
make build
make smoke
```

`-sync` already compiled the generated code before writing it, so `make
build` here is a confirmation, and `make smoke` adds test compilation and
acceptance smoke tests. If `make build` fails here, investigate rather than
hand-suppressing — `-sync` would already have suppressed a genuinely
non-compiling type, or aborted entirely on a machinery regression.

### 4. Documentation

Requires Terraform `v1.14`+ and `tfplugindocs` on `PATH` (both installed by
`make tools`):

```sh
go run ./internal/tools/bigdiffer -docs
```

This regenerates import-example docs, runs `terraform fmt`, and runs
`tfplugindocs generate`. Changes `examples/**` and `docs/**`.

### 5. CHANGELOG and version

`-sync` already drafted the `FEATURES:` bullets into `CHANGELOG.md`'s top,
in-progress version block (e.g. `## 1.100.0 (Unreleased)`). By hand:

1. Add the PR number link
   (`([#1234](https://github.com/hashicorp/terraform-provider-awscc/pull/1234))`)
   once the PR exists.
2. Add any `NOTES:` or breaking-change entries.
3. Update `version/VERSION` to match.

If this run promoted nothing newly user-visible, no `FEATURES:` section is
added.

`-sync` assumes the top block's `FEATURES:` section is empty or absent when
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

`-sync` narrates each phase (`==>` lines) and prints a reconciliation report to
stderr:

* **Detected changes — new / changed / unchanged / frozen / missing.** Every
  type is generate-and-compile-gated every run; New and Changed also refresh
  their schema bytes from AWS, Unchanged only re-verifies the machinery still
  produces the same output, Frozen (pinned via `frozen_since`) are never
  re-evaluated at all, and Missing means discovery could not describe the type.
* **Per-type lines** for anything frozen or suppressed, with the policy summary
  and reason — these are your issue-filing worklist.
* **Anomalies** — retained-but-unexplained rows (in the overlay, gone from AWS,
  with no `frozen_since`/`non_provisionable`/checkout pin), duplicate blocks, and
  naming-invariant violations. Investigate these before merging.

## Full regeneration, offline (`-reconcile`)

To regenerate the **entire** corpus from the committed overlay and schema cache
without touching AWS — for example after changing the generation engine — use:

```sh
go run ./internal/tools/bigdiffer -reconcile
```

This runs every committed type through the same generate → compile → decide
gate `-sync` uses, offline, and promotes the result if the whole run is clean.
Unlike the old, deleted `-generate` mode, `-reconcile` **does** compile-gate:
any failure — there is no schema change to distinguish here, so every failure
is a machinery regression — aborts the entire run and promotes nothing, with
the same per-type-and-artifact blame `-sync` reports. It does not query AWS
and writes the same files `-sync` would (schema cache, generated code,
registration file, import-examples aggregate, `all_schemas.hcl`, and a
CHANGELOG draft for any backlog artifact the fix newly recovers), just from
the committed base instead of a live crawl.

## Checking an engine change, offline and read-only (`-check`)

To find out whether a machinery change would break anything, without writing
anything at all:

```sh
go run ./internal/tools/bigdiffer -check
```

This runs `-reconcile`'s identical pipeline with promotion switched off. It
reports every failure, then exits non-zero on either a failure **or** any
regenerated output that differs from what's committed, even if it compiles —
the "you changed the engine but didn't `-reconcile` and commit yet" case. A
non-engine change touches nothing `-check` would regenerate differently, so it
passes with a clean diff; after an engine change, run `-check` to preview the
result, then `-reconcile` (or `-sync`) to actually land it. This is the
intended CI gate for engine-touching changes (templates, `codegen/`, naming,
`go.mod`/`go.sum`).

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
`internal/provider/all_schemas.hcl` in memory, gates the whole corpus
(generate and compile) every run, and lets one policy decide what to freeze
or suppress — a broken *schema* never blocks a release, but a regression in
the generation *machinery itself* fails the whole run rather than shipping
output from a generator just caught being broken.

For the full design — why it's built this way, what it replaces, and what's
still open — see [bigdiffer-design.md](bigdiffer-design.md).
