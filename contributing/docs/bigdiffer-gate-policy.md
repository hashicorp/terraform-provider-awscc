<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer: gate & policy — implementation plan

Working doc (may be discarded once the gap is closed). Tracks how to implement
the gate + policy gap — §6–§7 of `new-generation.md`, roadmap item #1 in §11.
Section refs below point at `new-generation.md`.

## Objective

Turn bigdiffer from "reconcile + report" into "reconcile → refresh → **gate** →
**apply policy** → write." Concretely: run generation per candidate type,
capture per-type results, and let the change-class × gate-result matrix (§7)
write `frozen_since` / `suppress_*` back into the overlay automatically.

## Key finding (de-risks the crux)

Generation is **already per-type**. `make resources` etc. are just
`go generate` fanning out over ~1500 `//go:generate` lines, each a standalone
invocation, e.g.:

```
go run generators/resource/main.go -resource awscc_acmpca_certificate \
  -cfschema ../service/cloudformation/schemas/AWS_ACMPCA_Certificate.json \
  -package acmpca -- ../aws/acmpca/certificate_resource_gen.go \
                     ../aws/acmpca/certificate_resource_gen_test.go
```

Those directive lines are themselves emitted by `generators/schema/main.go`
(from `schemas.go`), which reads `all_schemas.hcl`, downloads the JSONs, and
writes `resources.go` / `singular_data_sources.go` / `plural_data_sources.go`.

Implications:

- "Resilient per-type generation" needs **no generator refactor**. bigdiffer
  invokes the existing per-type CLIs as subprocesses, one per candidate, into a
  temp output dir, and collects exit code + stderr. One failure never affects
  another — resilience is free.
- The ~1500 process-per-type cost only bites on a *full* rebuild. bigdiffer's
  gate runs only on the **changed/new** set (byte-identical skip, §4), which is
  a handful per week — so per-type subprocess is entirely acceptable.
- The per-type **recipe** (awscc type name, `-listresource`, package, cfschema
  path, output paths) is non-trivial naming logic we do *not* want to
  reimplement. It is already materialized in the directive files. bigdiffer can
  reuse those lines rather than deriving them.

## Prerequisites / missing pieces (dependency order)

1. **Change detection (§4).** Which types actually changed this run. Requires
   selective refresh: for each non-frozen type, fetch AWS bytes, byte-compare to
   the cached JSON, replace only if different, and record the changed set.
   New types are always candidates. Frozen types are skipped.
   - Enabler: reuse the `DescribeType` bytes from discovery (§10 / roadmap #4)
     so the fetch is the same crawl, not a second one.
2. **Per-type recipe access.** The candidate's generate command(s). Cleanest:
   run `generators/schema/main.go` to (re)emit the directive files against the
   working overlay + cache, then look up the candidate's line(s) in
   `resources.go` / `singular_data_sources.go` / `plural_data_sources.go`.
   Honors `suppress_*` automatically (suppressed artifacts get no directive).
3. **Temp generation sandbox.** Execute the candidate's directives with the `--`
   output paths redirected to a temp dir, so a gate run never touches
   `internal/aws/**` until the result is accepted.
4. **Block mutation in the overlay.** New capability. Today bigdiffer only *adds*
   blocks and preserves existing ones verbatim. Policy must *set* attributes
   (`frozen_since = "<date>"`, `suppress_*_generation = true`) on an existing
   block while preserving its comments/reason. Implement as a text edit on the
   block's item, then `hcl`/`terraform fmt` to realign (the trick already used
   for migrations).
5. **Result taxonomy + policy mapping.** Per candidate:
   `ok | failed-validation | failed-generation | failed-build`, plus its change
   class. A pure function maps (class × result) → overlay edit per the §7 matrix.
6. **Compile gate (§6 tail, §10).** Per-type compile is impractical (needs
   package context). Keep it as a single whole-repo `go build` / `make build`
   after edits are applied; a build failure feeds the same policy (freeze the
   present offender / suppress the new one) and re-runs.

## Build plan (incremental, each independently testable)

- **Stage A — selective refresh + changed set.** Download non-frozen to temp,
  byte-compare, replace changed, emit the changed set. Frozen skipped. New mode;
  does not yet gate. Verify against a week with known churn.
- **Stage B — gate runner.** Given a candidate set, (re)emit directives, run the
  applicable per-type generators into temp, return per-type results. Test with a
  stub generator (fast) and with one real known-good type.
- **Stage C — policy engine.** Pure `apply(class, result) → edit` matching the
  §7 table. Table-driven unit tests; no AWS, no generation.
- **Stage D — block mutation.** `setAttr(block, name, value)` preserving
  comments; fmt. Unit-tested on the complex-block fixtures already in
  `main_test.go`.
- **Stage E — orchestration.** Wire run() into the pipeline: reconcile → refresh
  (A) → gate (B) → policy (C) → mutate (D) → report → optional build (P6).
  Keep the structural `-check` untouched and offline.
- **Stage F — later.** Self-heal re-probe of frozen types to scratch (§7 last
  row); machine-readable report (§8).

## Required bigdiffer refactors

- **run() becomes a pipeline.** Currently: read → `normalize` → write.
  New: read → reconcile → refresh → gate → policy → emit → report. `normalize`
  stays the reconcile/emit core; policy edits are applied to the item set before
  the final emit.
- **Item model gains mutation.** The verbatim text-blob item needs a
  "set attribute" operation (insert or replace an attribute line inside a live
  block, preserve everything else, reformat). This is the one real departure
  from pure verbatim preservation, and is the highest-risk refactor.
- **Report gains gate results.** New/changed types with their result and the
  applied policy action, for the human review step (§8).

## Open decisions

- **New-type gating needs a tentative overlay.** To get a directive for a New
  type, its row must exist when `generators/schema/main.go` runs. Plan: add New
  rows to a *working copy* of the overlay, emit directives + gate in temp, and
  only write the tree copy once policy is decided. Confirm this is acceptable vs.
  deriving New recipes directly.
- **Reuse vs. reimplement the schema generator.** For now, reuse it as a
  subprocess (recipe emission + download), consistent with "gate = black-box
  subprocess." Reimplementing recipe/naming derivation inside bigdiffer is
  explicitly out of scope until the tool stabilizes.
- **Change detection source.** Byte-compare in bigdiffer (self-contained) vs.
  `git diff` on the schemas dir (cheap but git-coupled). Prefer byte-compare.

## Verification strategy

- Unit: policy matrix (pure), block mutation (fixtures), result parsing (stub
  generator).
- Integration (real known-breaker): gate `AWS::S3::Bucket` against its *live*
  schema. It is frozen precisely because its current schema breaks generation,
  so the gate must return `failed-generation` and policy must propose
  `frozen_since` — a real end-to-end check of the whole gate→policy path.
- Integration (known-good): gate a simple stable type (e.g. `AWS::Logs::LogGroup`)
  and expect `ok` with no overlay edit.
- Idempotence: a second run with no AWS change produces no edits.

