<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Generating the Provider with bigdiffer

This is the **current** weekly process for generating the Terraform AWS Cloud
Control Provider from [CloudFormation resource type schemas](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html).

It replaces the long, manual, multi-`make`-target process in
[generating-the-provider.md](generating-the-provider.md), which is retained as a
[fallback](#fallback-the-legacy-process) should it ever be needed.

`bigdiffer` (`internal/tools/bigdiffer`) reconciles the live CloudFormation
registry directly against `internal/provider/all_schemas.hcl`, regenerates only
the resources whose schema changed, and applies one declarative policy so a
broken schema never blocks a release. The end result — the set of files changed —
is the same as the legacy process; you just reach it with two commands instead of
~15 `make` targets and manual editing.

<!--mdtoc: begin-->
* [What bigdiffer replaces](#what-bigdiffer-replaces)
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
<!--mdtoc: end-->

## What bigdiffer replaces

| Legacy step (generating-the-provider.md) | Legacy commands | bigdiffer |
| --- | --- | --- |
| 2. Schema Refresh | `make cleanschemas suppressions schemas` (repeat), `make commitrefresh` | `-update` |
| 3. Track New Schemas | `make biglister bigdiffer`, hand-edit `all_schemas.hcl`, `make schemas commitschemas` | `-update` |
| 4. Generate Resources | `make resources commitresources` | `-update` |
| 5. Generate Data Sources | `make singular-data-sources plural-data-sources commitdatas` | `-update` |
| 8. Documentation | `make docs-all` | `-docs` |

Steps 1 (Setup), 6 (Build), 7 (Smoke), 9 (CHANGELOG), and 10 (Pull Request) are
unchanged.

The manual "run, hit an error, hand-add a suppression, re-run" loops are gone:
when a type fails to generate, bigdiffer's policy freezes it (if it already
existed, keeping the last-good output) or adds it suppressed (if it is new), with
a `suppression_reason`, so the run always completes. You review what was
frozen/suppressed afterward and open issues as before.

## The weekly release

> [!TIP]
> `make bigdiffer-update`, `make bigdiffer-docs`, and `make bigdiffer-generate` are
> shortcuts for the `go run ./internal/tools/bigdiffer …` commands used below.

### 1. Setup

Unchanged from the legacy process:

```sh
git switch main
git pull
make tools
make newbranch
export AWS_DEFAULT_REGION=us-east-1   # bigdiffer always queries us-east-1
```

Ensure valid AWS credentials are set in the environment.

### 2. Update

One command does the schema refresh, new-schema tracking, resource generation, and
data-source generation:

```sh
go run ./internal/tools/bigdiffer -update
```

What it does, in one discovery crawl:

1. **Discovers** every provisionable public CloudFormation type in `us-east-1`
   (`ListTypes` + `DescribeType`). This is rate-limited and takes roughly 13
   minutes; a progress bar shows describe progress. There is no need to repeat it
   until clean — per-type failures are handled in-band.
2. **Detects changes** by comparing each type's freshly sanitized schema bytes
   against the committed cache, so only genuinely New or Changed types are touched.
3. **Regenerates** only those types (resource + data sources) from the fresh
   bytes, in parallel. Generation *is* the first gate: a type that still generates
   is refreshed; one that breaks is frozen or suppressed by policy.
4. **Compile-gates** the staged code before writing anything real: it builds the
   exact set it is about to promote — every staged artifact plus the regenerated
   `registrations_gen.go` — with `go build ./...`, and only promotes once that is
   green. A type whose code generates but fails to type-check is suppressed
   (`build_failed`) the same way a generation failure is, so a broken build can
   no longer reach a commit.
5. **Never regresses**: generated files and the schema cache are promoted only on
   clean generation *and* a green build. A broken type keeps its last-good
   committed output and cache.
6. **Reconciles `all_schemas.hcl`**: new resource blocks are added, blocks are
   re-sorted by CloudFormation type name, and existing hand-annotations
   (`# Suppression Reason` comments, suppress flags) are preserved byte-for-byte.
   The policy edits (freeze/suppress + `suppression_reason`) are applied in the
   same pass.
7. **Emits the aggregates**: `internal/provider/registrations_gen.go` and
   `internal/provider/import_examples_gen.json`.

It prints a [report](#reading-the-report) of what changed and what (if anything)
was frozen or suppressed. Review it together with `git diff`. For each
frozen/suppressed type, open a GitHub issue with the reason (see
[example issue #2070](https://github.com/hashicorp/terraform-provider-awscc/issues/2070));
the `internal/update/suppressions_checkout.txt` pins are still honored.

Files `-update` changes (the same set the legacy process produces, plus
`registrations_gen.go`):

* `internal/service/cloudformation/schemas/*.json` — refreshed schema cache
* `internal/provider/all_schemas.hcl` — new blocks + policy annotations
* `internal/aws/**/*_gen.go` — regenerated code for changed types only
* `internal/provider/import_examples_gen.json` — import-examples aggregate
* `internal/provider/registrations_gen.go` — see the note below

> [!NOTE]
> **`registrations_gen.go`** is bigdiffer's single blank-import file that
> registers every generated resource and data source. It supersedes the legacy
> `internal/provider/{resources,singular_data_sources,plural_data_sources}.go`
> directive files, which bigdiffer does **not** touch. During the transition both
> may be present: duplicate blank imports are legal Go and each package's `init()`
> runs once, so registration is correct. Commit `registrations_gen.go`; the legacy
> three files are now redundant and will be removed after a few clean cycles.

### 3. Build and smoke test

`-update` already compile-gated the generated code internally (step 4 above), so
`make build` here is a belt-and-suspenders confirmation rather than the primary
catch it was under the legacy process, and `make smoke` adds the test compilation
and acceptance smoke tests the compile gate deliberately does not cover:

```sh
make build
make smoke
```

If `make build` somehow still fails, that points at something the internal gate
does not cover (e.g. a hand-edit after `-update`, or a bug in the gate itself) —
investigate rather than routinely hand-suppressing, since `-update` would have
already suppressed a genuinely non-compiling generated type as `build_failed`.

### 4. Documentation

Replaces `make docs-all`. Requires Terraform `v1.14`+ (for list-resource support)
and `tfplugindocs` on `PATH` (both installed by `make tools`):

```sh
go run ./internal/tools/bigdiffer -docs
```

bigdiffer owns the import-example docs (generated from `import_examples_gen.json`)
and orchestrates the two external steps — `terraform fmt` and `tfplugindocs
generate` — it does not reimplement `tfplugindocs`. This changes `examples/**` and
`docs/**`.

### 5. CHANGELOG and version

Unchanged (legacy step 9). List the new documentation files and fold them into the
changelog:

```sh
git ls-files --others --exclude-standard
```

Copy these into `CHANGELOG.md`, format appropriately, and update `version/VERSION`
to match.

### 6. Commit and open a pull request

This is the one place the workflow differs mechanically. The legacy process
committed in stages via `make commitrefresh/commitschemas/commitresources/commitdatas/commitdocs`.
bigdiffer instead leaves the entire result in your working tree in one pass. Review
`git status` and `git diff`, then commit — either as a single commit or grouped
logically, for example:

```sh
git add internal/service/cloudformation/schemas internal/provider/all_schemas.hcl
git commit -m "Refresh CloudFormation schemas"

git add internal/aws internal/provider/registrations_gen.go internal/provider/import_examples_gen.json
git commit -m "Regenerate resources and data sources"

git add examples docs CHANGELOG.md version/VERSION
git commit -m "Regenerate documentation and update changelog"
```

The resulting file set matches the legacy process (plus `registrations_gen.go`).
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
`import_examples_gen.json` in parallel, in seconds. It does not query AWS and does
not change `all_schemas.hcl`.

## Fallback: the legacy process

The legacy machinery is intact and unaffected by bigdiffer; it produces the same
result. If bigdiffer is ever unavailable or misbehaving, follow
[generating-the-provider.md](generating-the-provider.md) as before
(`make cleanschemas suppressions schemas`, `make biglister`, `make resources`,
`make singular-data-sources plural-data-sources`, `make docs-all`, …).

Interoperating between the two is seamless:

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
