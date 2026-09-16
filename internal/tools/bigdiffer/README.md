<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer

`bigdiffer` is the tool that generates the Terraform AWSCC provider from AWS
CloudFormation resource type schemas.

Once a week (or on demand), it checks AWS for new or changed CloudFormation
types, regenerates and compile-gates the whole corpus, and updates
`internal/provider/all_schemas.hcl` — the file that tracks every type the
provider knows about — to match. If one type's *schema* is broken and won't
generate or compile, bigdiffer freezes or suppresses just that type instead of
failing the whole run. If the *generation machinery itself* regresses — a
template, codegen, or toolchain change — bigdiffer fails the whole run instead:
a schema break is external and gets frozen; a machinery break is self-inflicted
and gets fixed, not papered over.

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
# Weekly release: discover live, gate the whole corpus (generate + compile),
# apply the policy to all_schemas.hcl, write the cache + aggregates + changelog
# (needs AWS, us-east-1):
go run ./internal/tools/bigdiffer -sync

# Regenerate the whole provider offline from the committed overlay + schema
# cache through the same gate -sync uses, and promote the clean result
# (no AWS; fails the run on any machinery regression):
go run ./internal/tools/bigdiffer -reconcile

# CI gate for engine changes: -reconcile's pipeline with promotion off. Reports
# every failure, exits non-zero on any failure or on any regenerated output
# that differs from committed even though it compiles (no AWS; writes nothing):
go run ./internal/tools/bigdiffer -check

# Regenerate documentation: import-example docs from import_examples_gen.json,
# then orchestrate terraform fmt + tfplugindocs (both must be on PATH):
go run ./internal/tools/bigdiffer -docs

# CI guard: verify all_schemas.hcl is normalized and anomaly-free, and that
# registrations_gen.go is current (offline, writes nothing; exits non-zero on
# a problem):
go run ./internal/tools/bigdiffer -lint

# Re-probe suppressed/frozen rows with no recorded reason and propose a
# reclassification (offline except reading the schema cache; never writes
# all_schemas.hcl). Add -recheck-all to widen the scope to every active fact,
# not just the reason-less/unknown backlog:
go run ./internal/tools/bigdiffer -recheck
go run ./internal/tools/bigdiffer -recheck -recheck-all
```

Exactly one command runs per invocation; with no command, usage is printed.
`make bigdiffer-sync`, `make bigdiffer-reconcile`, and `make bigdiffer-docs` are
shortcuts.

Flags:

- `-sync` — the weekly release: one discovery crawl reconciles the overlay,
  gates the whole corpus (generate + compile, not just New/Changed) through the
  same policy every run, promotes on a clean gate (never regress), freezes or
  suppresses what a schema change breaks, and re-emits the aggregates + a
  CHANGELOG draft (needs AWS, `us-east-1`)
- `-reconcile` — regenerate the whole provider offline from the committed
  overlay + schema cache through the same gate `-sync` uses, and promote the
  clean result; no AWS. Any failure fails the whole run — a schema-unchanged
  failure means the machinery regressed, so nothing is promoted, not even
  types that compiled cleanly this run
- `-check` — `-reconcile`'s identical pipeline with promotion switched off:
  read-only, reports every failure, and exits non-zero on any failure or any
  regenerated output that differs from committed even though it compiles (an
  engine change that hasn't been `-reconcile`d and committed yet); no AWS,
  writes nothing
- `-docs` — import-example docs from `import_examples_gen.json`, then
  orchestrate `terraform fmt` and `tfplugindocs generate` (does not reimplement
  `tfplugindocs`)
- `-lint` — verify `all_schemas.hcl` is normalized (sorted, canonical, correct
  count header) and anomaly-free, and that `registrations_gen.go` is up to
  date; offline, writes nothing, suitable for CI
- `-recheck` — re-probe every suppressed/frozen row with no recorded reason (or
  one tagged `unknown`) and propose a reclassification; offline except reading
  the committed schema cache, never writes `all_schemas.hcl`
- `-recheck-all` — with `-recheck`: widen the scope to every active
  suppression/freeze, not just the reason-less/unknown backlog — revisit an
  old, already-explained decision on demand
- `-all-schemas` — overlay path (default `internal/provider/all_schemas.hcl`)
- `-checkout` — `suppressions_checkout.txt` path (cross-referenced, never modified)

## What it touches

`-sync` reads AWS and the overlay, and writes:

- `internal/service/cloudformation/schemas/*.json` — the schema cache
- `internal/provider/all_schemas.hcl` — the overlay, reconciled and re-sorted
- `internal/aws/**/*_gen.go` (+ `_test.go`) — generated code for every type the
  gate promotes this run (a type whose regenerated output differs from
  committed and compiles is promoted too, so a machinery change self-heals
  across the corpus)
- `internal/provider/registrations_gen.go` — the blank-import registration file
- `internal/provider/import_examples_gen.json` — the import-examples aggregate
- `CHANGELOG.md` — a `FEATURES:` draft for newly user-visible artifacts

`-reconcile` writes the same code/registration/import-examples/overlay/changelog
outputs as `-sync`, offline, from the committed overlay and cache — without
touching AWS. `-check` writes nothing. `-docs` writes `examples/**` and
`docs/**`. `-lint` and `-recheck` write nothing.

## Suppression and freezing, in one sentence

Every tracked type lives in one `all_schemas.hcl` block with its own policy:
`frozen_since` pins its schema bytes (don't refresh), `suppress_*_generation`
skips generating one artifact (resource / singular data source / plural data
source), and a `*_reason` field records why. See
[suppressed-and-frozen.md](../../../contributing/docs/suppressed-and-frozen.md)
for the full reason taxonomy and `-recheck`.

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
