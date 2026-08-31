<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer

`bigdiffer` reconciles `internal/provider/all_schemas.hcl` (the hand-curated
"overlay") against the live CloudFormation registry, regenerates only the
resources whose schema changed, and applies one declarative policy — replacing
the error-prone weekly ritual of hand-jamming new `resource_schema` blocks and
running a chain of `make` targets.

- **Weekly release process:** `contributing/docs/generating-the-provider-with-bigdiffer.md`
  (and the legacy fallback in `generating-the-provider.md`).
- **Design and status:** `contributing/docs/new-generation.md` (target design,
  §11 status table) and `contributing/docs/bigdiffer-generation.md` (the
  generation / orchestration working doc).

> Note: there is also a `make bigdiffer` target that only prints the raw `diff`
> between two dated `available_schemas` files. It is unrelated and slated for
> removal; this tool replaces it.

## Usage

```sh
# Weekly release: discover live, regenerate only the types whose schema changed,
# apply the policy to all_schemas.hcl, write the cache + aggregates (needs AWS):
go run ./internal/tools/bigdiffer -update

# Regenerate the whole provider offline from the committed overlay + schema cache
# (parallel, in-process; writes *_gen.go, registrations_gen.go,
# import_examples_gen.json). No AWS:
go run ./internal/tools/bigdiffer -generate

# Regenerate documentation: import-example docs from import_examples_gen.json,
# then orchestrate terraform fmt + tfplugindocs (both must be on PATH):
go run ./internal/tools/bigdiffer -docs

# CI guard: verify all_schemas.hcl is normalized and anomaly-free (offline,
# writes nothing; exits non-zero on a problem):
go run ./internal/tools/bigdiffer -check
```

Exactly one command runs per invocation; with no command, usage is printed.
`make bigdiffer-update`, `make bigdiffer-generate`, and `make bigdiffer-docs` are
shortcuts.

Flags:

- `-update` — the weekly release: one discovery crawl reconciles the overlay and
  regenerates only the New/Changed types from their fresh bytes, promoting files
  and cache only on clean generation (never regress), freezing or suppressing what
  breaks, then re-emitting the aggregates (needs AWS, `us-east-1`)
- `-generate` — regenerate the whole provider offline from the committed overlay +
  schema cache (writes generated code + aggregates); does not touch AWS
- `-docs` — import-example docs from `import_examples_gen.json`, then orchestrate
  `terraform fmt` and `tfplugindocs generate` (does not reimplement `tfplugindocs`)
- `-check` — verify `all_schemas.hcl` is normalized (sorted, canonical, correct
  count header) and anomaly-free; offline, writes nothing, suitable for CI
- `-all-schemas` — overlay path (default `internal/provider/all_schemas.hcl`)
- `-checkout` — `suppressions_checkout.txt` path (cross-referenced, never modified)

## What it does

`-update` runs the full weekly pipeline off a single discovery crawl:

1. **Discover** every provisionable public type in `us-east-1` (`ListTypes` +
   `DescribeType`), capturing each type's sanitized schema bytes in the same sweep.
2. **Detect changes** by byte-comparing those bytes against the committed cache;
   only New and Changed types are touched (`change.go`).
3. **Regenerate** only those types — resource + data sources — from the fresh
   bytes, in parallel and in-process (`corpus.go`, `codegen/`).
4. **Generation is the gate**: a type that generates cleanly is refreshed; one
   that breaks is handed to the policy (`policy.go`), which freezes an existing
   type at its last-good bytes or adds a new one suppressed, always with a
   `suppression_reason` — so a broken schema never blocks the release.
5. **Never regress**: generated files and the schema cache are promoted only on
   clean generation (staged in a temp dir first; `update.go`).
6. **Reconcile `all_schemas.hcl`** and apply the policy edits in one
   comment-preserving pass, then **emit the aggregates** — `registrations_gen.go`
   and `import_examples_gen.json` (`emit.go`).

`-generate` runs the generation + aggregate steps over the **whole** corpus
offline (no AWS); `-docs` generates the import-example docs and orchestrates
`terraform fmt` + `tfplugindocs`; `-check` verifies `all_schemas.hcl` is
normalized and anomaly-free (offline), for CI.

During the migration off the legacy `make` generators, a full-corpus parity test
validated bigdiffer's output against them (1579 types / 8432 files, byte-for-byte);
it remains as a regression guard while the legacy path coexists. Matching the
legacy output is not an ongoing goal — bigdiffer is now the authoritative
generator; the parity test simply catches accidental engine drift.

### Overlay reconcile (byte-preserving)

`-update` reconciles `all_schemas.hcl` against the discovered live registry (this
also powers the offline `-check`):

- Joins overlay and registry on the **CloudFormation type name**.
- **Adds** rows for types present in the registry but missing from the overlay,
  copying the resource type name and `suppress_plural_data_source_generation` flag
  from discovery. This also recovers any resources available earlier but never
  added.
- **Preserves every existing block byte-for-byte**, including free-form
  `# Suppression Reason` comments, manual suppression attributes, leading
  comments, and fully commented-out blocks. (A `gohcl` round-trip cannot do this:
  it discards comments.)
- **Re-sorts** all blocks by CloudFormation type name. This normalizes accumulated
  hand-edit disorder.
- Updates the header count to the number of available schemas.
- Is **idempotent**: reconciling twice makes no further change. `-check` verifies
  this offline (overlay reconciled against itself), suitable for CI.

### Discovery: one live sweep

Discovery (`discover.go`) queries `ListTypes` + `DescribeType` directly, with no
checked-in snapshots. It is concurrent (bounded worker pool) and resilient: a
per-type `DescribeType` failure is captured in-band and does not abort the crawl.
The same sweep captures each type's sanitized schema bytes, which change detection
and generation consume.

### How overlay blocks are read

Blocks are delimited by their braces (so internal blank lines and inline comments
are kept intact). The scanner's set of live blocks is then **cross-validated
against a real HCL decode** (`hclparse` + `gohcl` into the `allschemas` structs);
a mismatch is a hard error. Fully commented-out blocks are handled by the text
scanner (HCL sees them only as loose comment tokens) and are keyed by their
commented-out type name so they sort into place.

## Generation internals

bigdiffer **owns** its generation engine (copied into `codegen/`, importing
nothing under `internal/provider/generators/**` — enforced by
`import_boundary_test.go`), so it can generate in-process and concurrently:

- **Recipe derivation** (`plan.go`) — from an overlay row alone, derives the
  artifacts to generate (resource, singular/plural data source), their Terraform
  type names, and file paths, honoring the row's `suppress_*` flags.
- **Concurrent, faithful generation** (`corpus.go`, `codegen/`) — the emitter has
  no mutable package state; the one non-thread-safe dependency (the inflection
  pluralizer, which mutates globals on every call) is isolated in the owned,
  mutex-guarded `naming/` package. Schemas are parsed from in-memory bytes
  (`NewResourceJsonSchemaDocument`) rather than by path, avoiding a
  working-directory side effect and keeping generation `-race`-clean.
- **Policy** (`policy.go`) — `decide` maps a change class and generation result to
  an overlay edit: add a plain block, add one suppressing only the artifacts that
  failed, freeze at last-good bytes, annotate non-provisionable, or freeze a
  withdrawn type. Pure and table-tested. Generation itself is the gate:
  `refreshCandidate` generates a type and `gateResultFromGenResults` maps the
  per-artifact success/failure into the `gateResult` that `decide` consumes
  (`gate.go` now holds just those outcome types).
- **Block mutation** (`mutate.go`) — sets attributes on an existing block via
  `hclwrite`, preserving comments and existing attributes; blocks for new
  additions are synthesized during the reconcile pass.
- **Aggregates** (`emit.go`) — the single `registrations_gen.go` blank-import file
  (which supersedes the legacy `resources.go` / `singular_data_sources.go` /
  `plural_data_sources.go` directive files) and `import_examples_gen.json`.

See `contributing/docs/bigdiffer-generation.md` (working doc) and
`contributing/docs/new-generation.md` (durable design, §11 status) for the full
picture, and `generating-the-provider-with-bigdiffer.md` for the weekly process.

## Suppression model (why there are two mechanisms, for now)

- `internal/update/suppressions_checkout.txt` — **schema-byte pinning**. Consumed
  by `make suppressions` (a `git checkout` of specific `AWS_*.json` files) so a
  refresh keeps the last-good bytes when AWS's new schema is broken, or keeps the
  cached bytes for a type AWS has removed. It does not exclude anything from
  `all_schemas.hcl`.
- HCL `suppress_*` flags in `all_schemas.hcl` — **generation control** (resource /
  singular data source / plural data source), independent of pinning.

These are orthogonal today; `bigdiffer` treats the checkout list as a read-only
cross-reference and preserves the HCL flags verbatim. The design
(`new-generation.md` §5) folds the checkout file into `frozen_since` and retires
it; that migration has not happened yet.

## Anomaly reports (advisory)

The tool prints, but does not auto-fix:

- **retained but not pinned** — a block that is gone from the base and has no
  `suppressions_checkout.txt` pin (its schema bytes may be at risk on refresh).
- **duplicate live blocks** — more than one live block for the same CloudFormation
  type (a latent hand-edit error).
- **naming-invariant violations** — `resource_type_name` does not match the
  deterministic transform of the CloudFormation type name (the classic
  copy-paste-wrong-CFN mistake).
