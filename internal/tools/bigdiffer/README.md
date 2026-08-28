<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer

`bigdiffer` reconciles `internal/provider/all_schemas.hcl` (the hand-curated
"overlay") against the live CloudFormation registry, and is meant to replace
the error-prone weekly ritual of hand-jamming new `resource_schema` blocks.
Design and status: see `contributing/docs/new-generation.md` (target design,
§11 for the authoritative status table) and
`contributing/docs/bigdiffer-gate-policy.md` (gate + policy implementation plan).

> Note: there is also a `make bigdiffer` target that only prints the raw `diff`
> between two dated `available_schemas` files. It is unrelated and slated for
> removal; this tool replaces it.

## Usage

```sh
# Reconcile against AWS live (recommended) and write the normalized all_schemas.hcl:
go run ./internal/tools/bigdiffer -discover

# Or reconcile against a snapshot file instead of live AWS:
go run ./internal/tools/bigdiffer

# Report only; fail (non-zero) if all_schemas.hcl is not normalized or has
# blocking anomalies (duplicate blocks or naming-invariant violations):
go run ./internal/tools/bigdiffer -check
```

Flags:

- `-all-schemas` — overlay path (default `internal/provider/all_schemas.hcl`)
- `-discover` — query AWS live (`us-east-1`) for the base instead of reading a
  snapshot file
- `-available` — snapshot base path (default: newest `available_schemas.<date>.hcl`
  in `-allschemas-dir`); ignored when `-discover` is set
- `-allschemas-dir` — directory of `available_schemas.<date>.hcl` snapshot files
- `-checkout` — `suppressions_checkout.txt` path (cross-referenced, never modified)
- `-check` — verify only; write nothing

## What it does today (reconcile + report)

- Joins overlay and base on the **CloudFormation type name**.
- **Adds** base rows missing from the overlay, copying the resource type name and
  `suppress_plural_data_source_generation` flag verbatim from the generated base.
  This also recovers any resources that were available previously but never added
  ("backlog") — snapshot mode only; `-discover` has no "previous" to diff against.
- **Preserves every existing block byte-for-byte**, including free-form
  `# Suppression Reason` comments, manual suppression attributes, leading
  comments, and fully commented-out blocks. (A `gohcl` round-trip cannot do this:
  it discards comments.)
- **Re-sorts** all blocks by CloudFormation type name, matching the generator's
  `sort.Strings` ordering. This normalizes accumulated hand-edit disorder.
- Updates the header count to the number of available schemas.
- Is **idempotent**: running twice makes no further change. Use `-check` in CI.

### `-discover`: live AWS, one sweep

`-discover` queries `ListTypes` + `DescribeType` directly (`discover.go`) instead
of reading a checked-in snapshot. Discovery is concurrent (bounded worker pool)
and resilient: a per-type `DescribeType` failure is captured in-band and does not
abort the crawl. The same sweep captures each type's sanitized schema bytes,
which change detection and the gate consume — see below.

### How overlay blocks are read

Blocks are delimited by their braces (so internal blank lines and inline comments
are kept intact). The scanner's set of live blocks is then **cross-validated
against a real HCL decode** (`hclparse` + `gohcl` into the `allschemas` structs);
a mismatch is a hard error. Fully commented-out blocks are handled by the text
scanner (HCL sees them only as loose comment tokens) and are keyed by their
commented-out type name so they sort into place.

## Built but not yet wired in

These exist, are unit-tested, and are not yet called from `run()`:

- **Change detection** (`change.go`) — byte-compares each discovered schema
  against the on-disk cache, respects `frozen_since`, and classifies each type
  as new / changed / unchanged / frozen / missing.
- **Recipe derivation** (`plan.go`) — derives, purely in memory from an overlay
  row, the set of generated artifacts (resource, singular/plural data source),
  their Terraform type names, and file paths, honoring the row's `suppress_*`
  flags. Naming/pluralization logic (`naming.go`) is copied from
  `internal/naming` rather than imported, because the legacy pluralizer mutates
  global state on every call — unsafe under concurrency.
- **Gate** (`gate.go`) — runs the generation front half for the resource and
  singular-data-source artifacts (`shared.NewResource` + `codegen.Emitter`,
  deliberately not `shared.GenerateTemplateData`, to avoid its CWD-coupled side
  effects) and reports `ok` / `failed-validation` / `failed-generation`. Plural
  data sources aren't schema-emission-driven and are excluded from this gate by
  design; a plural-only failure surfaces at the compile gate instead. Runs
  candidates serially — the reused emission engine isn't safe to parallelize
  in-process, confirmed by the race detector — which is fine since the
  candidate set (new + changed types) is small.
- **Policy** (`policy.go`) — `decide` maps a change class and gate result to an
  overlay edit: add a plain block, add a block suppressing only the artifacts
  that failed, freeze at last-good bytes, annotate non-provisionable, or freeze
  a withdrawn type. Pure and table-tested.
- **Block mutation** (`mutate.go`) — sets attributes on an existing block via
  `hclwrite`, preserving comments and existing attributes. The capability
  policy needs to act on an existing block; synthesizing a new block's text for
  a suppressed New type is not yet built.

See `bigdiffer-gate-policy.md` for the remaining work: wiring all of this
together, writing refreshed bytes back to the cache, synthesizing new-block text
for suppressed additions, and the compile gate.

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
