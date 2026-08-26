<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer

`bigdiffer` normalizes `internal/provider/all_schemas.hcl` (the hand-curated
"overlay") against the newest generated
`internal/provider/generators/allschemas/available_schemas.<date>.hcl` (the
"base"). It replaces the error-prone weekly ritual of hand-jamming new
`resource_schema` blocks after `make bigdiffer`.

> Note: there is also a `make bigdiffer` target that only prints the raw `diff`
> between two `available_schemas` files. This tool is the thing you run *instead
> of* hand-editing after looking at that diff.

## Usage

```sh
# Write the normalized all_schemas.hcl (uses the newest available_schemas.<date>.hcl):
go run ./internal/tools/bigdiffer

# Report only; fail (non-zero) if all_schemas.hcl is not normalized or has
# blocking anomalies (duplicate blocks or naming-invariant violations):
go run ./internal/tools/bigdiffer -check
```

Flags:

- `-all-schemas` — overlay path (default `internal/provider/all_schemas.hcl`)
- `-available` — base path (default: newest `available_schemas.<date>.hcl` in `-allschemas-dir`)
- `-allschemas-dir` — directory of `available_schemas.<date>.hcl` files
- `-checkout` — `suppressions_checkout.txt` path (cross-referenced, never modified)
- `-check` — verify only; write nothing

## What it does (Phase 1)

- Joins overlay and base on the **CloudFormation type name**.
- **Adds** base rows missing from the overlay, copying the resource type name and
  `suppress_plural_data_source_generation` flag verbatim from the generated base.
  This also recovers any resources that were available previously but never added
  ("backlog").
- **Preserves every existing block byte-for-byte**, including free-form
  `# Suppression Reason` comments, manual suppression attributes, leading
  comments, and fully commented-out blocks. (A `gohcl` round-trip cannot do this:
  it discards comments.)
- **Re-sorts** all blocks by CloudFormation type name, matching the generator's
  `sort.Strings` ordering. This normalizes accumulated hand-edit disorder.
- Updates the header count to the number of available schemas.
- Is **idempotent**: running twice makes no further change. Use `-check` in CI.

### How blocks are read

Blocks are delimited by their braces (so internal blank lines and inline comments
are kept intact). The scanner's set of live blocks is then **cross-validated
against a real HCL decode** (`hclparse` + `gohcl` into the `allschemas` structs);
a mismatch is a hard error. Fully commented-out blocks are handled by the text
scanner (HCL sees them only as loose comment tokens) and are keyed by their
commented-out type name so they sort into place.

## Suppression model (why there are two mechanisms)

- `internal/update/suppressions_checkout.txt` — **schema-byte pinning**. Consumed
  by `make suppressions` (a `git checkout` of specific `AWS_*.json` files) so a
  refresh keeps the last-good bytes when AWS's new schema is broken, or keeps the
  cached bytes for a type AWS has removed. It does not exclude anything from
  `all_schemas.hcl`.
- HCL `suppress_*` flags in `all_schemas.hcl` — **generation control** (resource /
  singular data source / plural data source), independent of pinning.

These are orthogonal; `bigdiffer` treats the checkout list as a read-only
cross-reference and preserves the HCL flags verbatim.

## Anomaly reports (advisory)

The tool prints, but does not auto-fix:

- **retained but not pinned** — a block that is gone from the base and has no
  `suppressions_checkout.txt` pin (its schema bytes may be at risk on refresh).
- **duplicate live blocks** — more than one live block for the same CloudFormation
  type (a latent hand-edit error).
- **naming-invariant violations** — `resource_type_name` does not match the
  deterministic transform of the CloudFormation type name (the classic
  copy-paste-wrong-CFN mistake).

## Roadmap

- **Phase 2**: derive the `suppress_*` flags from an actual single-schema
  download + generation dry-run instead of carrying them forward.
- **Phase 3**: re-validate currently-suppressed/pinned resources to detect
  AWS-side fixes and propose un-suppression.
