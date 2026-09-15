<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer

`bigdiffer` is the tool that generates the Terraform AWSCC provider from AWS
CloudFormation resource type schemas.

Once a week (or on demand), it checks AWS for new or changed CloudFormation
types, regenerates the Terraform code for anything that changed, and updates
`internal/provider/all_schemas.hcl` — the file that tracks every type the
provider knows about — to match. If a type's schema is broken and won't
generate or compile, bigdiffer freezes or suppresses just that type instead of
failing the whole run, so one bad schema never blocks a release.

- **To run the weekly release:** see
  [generating-the-provider-with-bigdiffer.md](../../../contributing/docs/generating-the-provider-with-bigdiffer.md).
- **For the full design — why it's built this way, the gates, the policy, and
  what's left to do:** see
  [bigdiffer-design.md](../../../contributing/docs/bigdiffer-design.md).

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

## What it touches

`-update` reads AWS and the overlay, and writes:

- `internal/service/cloudformation/schemas/*.json` — the schema cache
- `internal/provider/all_schemas.hcl` — the overlay, reconciled and re-sorted
- `internal/aws/**/*_gen.go` (+ `_test.go`) — generated code for changed types
- `internal/provider/registrations_gen.go` — the blank-import registration file
- `internal/provider/import_examples_gen.json` — the import-examples aggregate

`-generate` writes the last three from the committed overlay and cache, without
touching AWS. `-docs` writes `examples/**` and `docs/**`. `-check` writes
nothing.

## Suppression and freezing, in one sentence

Every tracked type lives in one `all_schemas.hcl` block with its own policy:
`frozen_since` pins its schema bytes (don't refresh), `suppress_*_generation`
skips generating one artifact (resource / singular data source / plural data
source), and a `*_reason` field records why. See
[suppressed-and-frozen.md](../../../contributing/docs/suppressed-and-frozen.md)
for the full reason taxonomy and `-heal`.

## Where things live

- **Design reference:**
  [bigdiffer-design.md](../../../contributing/docs/bigdiffer-design.md) — the
  model, the two gates, the policy table, settled design decisions, and the
  deferred-work list.
- **Weekly runbook:**
  [generating-the-provider-with-bigdiffer.md](../../../contributing/docs/generating-the-provider-with-bigdiffer.md) —
  step-by-step release process, plus the legacy fallback.
- **Suppression/freeze spec:**
  [suppressed-and-frozen.md](../../../contributing/docs/suppressed-and-frozen.md).
- **Legacy removal plan:**
  [removing-the-legacy-generation-process.md](../../../contributing/docs/removing-the-legacy-generation-process.md).
