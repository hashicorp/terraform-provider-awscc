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
  path, output file names) is non-trivial naming logic — but it is compact and
  copyable. `schema/main.go`'s `Downloader.Schemas` + `ResourceSchema` derive it
  in ~170 lines, depending only on `internal/naming` (already a bigdiffer
  dependency) and `cfschema`. Copy that core into bigdiffer and compute recipes
  **in memory** — no temp overlay, no schema-generator subprocess, no
  directive-file round-trip.
- `schema/main.go` is *not* the code generator. It derives recipes, downloads,
  and writes the `//go:generate` directive lists. The actual code generation is
  in `resource/main.go` etc. atop `shared/`. So the gate's pass/fail comes from
  generation's **front half** — `shared.GenerateTemplateData(...)`, importable —
  where almost all generation failures surface. That is §6's "validation is the
  front half of generation," in-process, with no temp output tree.

## Prerequisites / missing pieces (dependency order)

1. **Change detection (§4).** Which types actually changed this run. Requires
   selective refresh: for each non-frozen type, fetch AWS bytes, byte-compare to
   the cached JSON, replace only if different, and record the changed set.
   New types are always candidates. Frozen types are skipped.
   - Enabler: reuse the `DescribeType` bytes from discovery (§10 / roadmap #4)
     so the fetch is the same crawl, not a second one.
2. **Recipe + refresh core (copy & redesign, not reuse).** Copy the reusable
   ~170 lines of `schema/main.go` — `Downloader.Schemas` (per-type recipe:
   Terraform type via `naming.CreateTerraformTypeName`, package/paths, plural
   name, `-listresource`, and `suppress_*` gating) and `ResourceSchema`
   (download-or-cached + `frozen_since` + meta-schema validation). Redesign for
   the in-memory pipeline: return recipes as values, refresh by
   compare-not-missing (§4), and **drop** the directive-file/template emission
   (`GenerateResources`/`GenerateDataSources`) — bigdiffer never writes
   `resources.go`; it consumes recipes directly.
3. **Gate mechanism (in-process front half).** The pass/fail signal is
   `shared.GenerateTemplateData(...)` per applicable artifact (resource /
   singular DS / plural DS). It parses the CFN schema into template data and is
   where the vast majority of generation failures surface (unsupported types,
   recursive definitions, bad identifiers). No temp output tree is needed because
   recipes are in memory and the check is in-process. Full template rendering +
   compilation is covered once by the whole-repo build (P6). New dependency:
   `shared` (the generation engine core) — acceptable now; the template layer can
   be copied later if full independence is wanted.
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
- **Stage B — gate runner.** For each candidate, derive its recipe(s) in memory
  (copied core, P2) and run the generation front half
  (`shared.GenerateTemplateData`) for the applicable artifacts, returning
  per-type results. No temp tree, no subprocess. Test with a known-good and a
  known-broken type.
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

- **New recipe/refresh module (copied core).** A bigdiffer-owned file adapted
  from `schema/main.go`'s `Downloader.Schemas` + `ResourceSchema`: derive
  per-type recipes in memory and refresh the cache by compare-not-missing. No
  directive emission. This is the "copy for free, redesign" piece.
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

- **Copy & redesign, don't reuse (decided).** bigdiffer absorbs `schema/main.go`'s
  ~170-line recipe/refresh core rather than shelling out to it. This removes the
  temp-overlay + subprocess + directive-file round-trip entirely.
- **New-type gating is no longer special (resolved).** Because recipes are
  derived in memory from a `resourceRow`, a New type's recipe is computed
  directly — nothing is written to disk to gate it, so no tentative overlay is
  needed. (This dissolves the open item from the first draft.)
- **`shared` dependency (open).** The gate's front half uses
  `shared.GenerateTemplateData`. Acceptable now; copy the template layer later if
  full independence from the generation engine is wanted.
- **Change detection source (lean byte-compare).** Byte-compare in bigdiffer
  (self-contained) vs. `git diff` on the schemas dir (cheap but git-coupled).
  Prefer byte-compare.

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

