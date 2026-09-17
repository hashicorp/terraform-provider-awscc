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
  [bigdiffer-runbooks.md](../../../contributing/docs/bigdiffer-runbooks.md).
- **For the full design — why it's built this way, the gates, the policy, and
  what's left to do:** see
  [bigdiffer-design.md](../../../contributing/docs/bigdiffer-design.md).

> Note: there is also a `make bigdiffer` target that only prints the raw `diff`
> between two dated `available_schemas` files. It is unrelated and slated for
> removal; this tool replaces it.

## Usage

```sh
# Weekly release: discover live, gate the whole corpus (generate + compile),
# apply the policy to all_schemas.hcl, write the cache + aggregates +
# changelog, and regenerate documentation as the last step (needs AWS,
# us-east-1; add -no-docs to skip the documentation step for a faster loop):
go run ./internal/tools/bigdiffer -sync

# Regenerate the whole provider offline from the committed overlay + schema
# cache through the same gate -sync uses, and promote the clean result,
# including documentation (no AWS; fails the run on any machinery
# regression; -no-docs skips documentation):
go run ./internal/tools/bigdiffer -reconcile

# CI gate for engine changes: -reconcile's pipeline with promotion off. Reports
# every failure, exits non-zero on any failure or on any regenerated output
# (code, or the import-examples aggregate) that differs from committed even
# though it compiles (no AWS; writes nothing). Add -check-docs-full to also
# render the registry docs themselves via a real built binary and diff
# against committed docs/ — markedly heavier, opt-in until proven
# fast/stable enough for every PR:
go run ./internal/tools/bigdiffer -check
go run ./internal/tools/bigdiffer -check -check-docs-full

# Regenerate documentation on its own: import-example docs from
# import_examples_gen.json, then orchestrate terraform fmt + tfplugindocs
# (both must be on PATH). -sync/-reconcile already run this as their last
# step by default; use this directly for an examples-only or doc-template
# change, or to finish a run that failed only at the documentation step:
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
  suppresses what a schema change breaks, re-emits the aggregates + a
  CHANGELOG draft, and regenerates documentation as the last step (needs AWS,
  `us-east-1`)
- `-reconcile` — regenerate the whole provider offline from the committed
  overlay + schema cache through the same gate `-sync` uses, and promote the
  clean result; no AWS. Any failure fails the whole run — a schema-unchanged
  failure means the machinery regressed, so nothing is promoted, not even
  types that compiled cleanly this run. Regenerates documentation as the last
  step, same as `-sync`
- `-no-docs` — with `-sync`/`-reconcile`: skip the trailing documentation step
  for a fast codegen inner loop, since `tfplugindocs` on the whole provider is
  slow; run `-docs` separately before committing
- `-check` — `-reconcile`'s identical pipeline with promotion switched off:
  read-only, reports every failure, and exits non-zero on any failure or any
  regenerated output that differs from committed even though it compiles (an
  engine change that hasn't been `-reconcile`d and committed yet), including
  whether `import_examples_gen.json` still matches the current row set; no
  AWS, writes nothing
- `-check-docs-full` — with `-check`: also render the registry docs themselves
  (via a real built binary extracted from the staged tree) and diff against
  committed `docs/`. Markedly heavier than everything else `-check` does — a
  real whole-provider build plus a `terraform providers schema -json` call —
  so it is opt-in, not part of the CI gate, until proven fast/stable enough
  for every PR
- `-docs` — import-example docs from `import_examples_gen.json`, then
  orchestrate `terraform fmt` and `tfplugindocs generate` (does not reimplement
  `tfplugindocs`). `-sync`/`-reconcile` already run this as their last step by
  default; use this directly for an examples-only or doc-template change, or
  to finish a run that failed only at the documentation step
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
- `examples/**`, `docs/**` — regenerated documentation, unless `-no-docs`

`-reconcile` writes the same code/registration/import-examples/overlay/changelog
outputs as `-sync`, offline, from the committed overlay and cache — without
touching AWS — plus documentation, same as `-sync` (unless `-no-docs`).
`-check` writes nothing, even with `-check-docs-full` (it renders into a
scratch directory to diff, never into the real `examples/`/`docs/`). `-docs`
writes `examples/**` and `docs/**`. `-lint` and `-recheck` write nothing.

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
- **Runbooks:**
  [bigdiffer-runbooks.md](../../../contributing/docs/bigdiffer-runbooks.md) —
  task-oriented runbooks, plus the legacy fallback.
- **Suppression/freeze spec:**
  [suppressed-and-frozen.md](../../../contributing/docs/suppressed-and-frozen.md).
- **Legacy removal plan:**
  [removing-the-legacy-generation-process.md](../../../contributing/docs/removing-the-legacy-generation-process.md).
