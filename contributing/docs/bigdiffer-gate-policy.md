<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer: gate & policy — implementation plan

Working doc (may be discarded once the gap is closed). Tracks how to implement
the gate + policy gap — §6–§7 of `new-generation.md`, roadmap item #3 in §11.
Section refs below point at `new-generation.md`.

## Status

Built and unit-tested, not yet wired into `run()`:

- `discover.go` — concurrent discovery, captures sanitized schema bytes per type.
- `change.go` — `detectChanges`/`classifyChange`: byte-compares discovered vs.
  cached schemas, respects `frozen_since`, classifies
  new/changed/unchanged/frozen/missing. Closes prerequisite #1 below.
- `plan.go` — `generationPlan`: pure, in-memory recipe derivation (artifacts,
  Terraform type, package/paths) per overlay row, honoring `suppress_*`. This
  is the reimplementation the "recipe" half of prerequisite #2 called for — see
  correction below on its actual size.
- `naming.go` — bigdiffer-owned naming/pluralization logic, copied from
  `internal/naming` rather than imported, because the legacy `Pluralize` mutates
  global `inflection` state on every call — a data race under concurrency,
  confirmed by the race detector. Fixes it by registering the irregular once and
  serializing access.
- `gate.go` — `gateType`/`runGate`: in-process front-half gate. Gates only the
  resource and singular-data-source artifacts (plural data sources are not
  schema-emission-driven and cannot fail this gate by construction — see
  correction below). Runs candidates **serially**, not concurrently: the reused
  legacy emission engine (`codegen.Emitter`, via `naming.go`'s dependencies) is
  not safe for concurrent in-process use, also confirmed by the race detector.
  Acceptable because the candidate set (New + Changed) is small by design.
- `policy.go` — `decide(class, gateResult, today)`: pure implementation of the
  §7 matrix. New+ok adds a plain block; New+failed adds a block suppressing only
  the artifacts that actually failed (a finer-grained action than §7's table
  originally stated — see correction below) with a reason; Present+ok keeps the
  block; Present+failed freezes at last-good bytes; non-provisionable annotates;
  withdrawn freezes. Encodes both §7 invariants. Table-tested.
- `mutate.go` — `setBlockAttributes`: edits a single block via `hclwrite`,
  adding/updating attributes while preserving comments and existing attributes,
  and re-aligning the result. The one capability beyond verbatim preservation
  that policy needs. Tested for add/update/preserve/no-op.

Not yet built: cache write-back (bytes compared but not persisted), synthesizing
a *new* block's text from a `policyDecision` (mutation only edits existing
blocks so far), the compile gate, and wiring all of the above into one pipeline.
See `new-generation.md` §11 for the authoritative status table.

## Corrections to this plan's original claims

- **Recipe core size.** The "Key finding" below estimated the reusable core at
  "~170 lines, copy & redesign." The actual reimplementation, `plan.go`, is
  ~95 lines including doc comments. The estimate was in the right range; treat
  it as resolved rather than a remaining unknown.
- **Plural data sources have no front-half gate.** The "gate mechanism" section
  below states the front half runs "per applicable artifact (resource /
  singular DS / plural DS)." That's wrong for plural DS: `plural-data-source/main.go`
  never calls the schema-parsing front half — it builds its template data from
  the CFN type name alone and never touches `shared.NewResource`. `gate.go` now
  reflects this correctly: `gateType` explicitly skips `artifactPluralDataSource`
  with a comment explaining why, rather than gating it and getting a
  meaningless `ok`. Treat `failed-validation` / `failed-generation` as only
  possible for `artifactResource` and `artifactSingularDataSource`.
- **The `GenerateTemplateData` CWD hazard is resolved, not merely noted.**
  `gate.go` avoids `shared.GenerateTemplateData` entirely — it calls
  `shared.NewResource` + `codegen.Emitter.EmitRootPropertiesSchema` directly, so
  the `last_resource.txt` write and the CWD-relative `services.hcl` read never
  happen during gating. No `shared` refactor is needed for concurrency safety.
- **In-process concurrency for gating is not safe, and the plan's "resilience is
  free" framing undersold this.** The original "Key finding" argued per-type
  subprocess invocation buys resilience for free; the actual implementation
  went in-process instead (a stronger, harder property) and discovered via the
  race detector that the reused legacy emission engine mutates global state
  (`inflection.AddIrregular` in what became `naming.go`) — unsafe under
  goroutines even though the legacy pipeline never hit this bug (it parallelizes
  across OS processes). `gate.go`'s `runGate` is therefore serial by design, not
  as a placeholder. This is the right call for now: the candidate set is small,
  so serial gating is cheap, and it avoids auditing the rest of `codegen.Emitter`
  for other shared state before it's needed.
- **Policy is finer-grained than §7's table text.** `decide()` suppresses only
  the artifacts that actually failed (e.g. a resource that fails but whose
  singular DS would succeed only sets `suppress_resource_generation`), not a
  blanket "suppress everything" for the type. `new-generation.md` §7 has been
  updated to reflect this.
- **The S3::Bucket integration fixture (Verification strategy, below) is
  unconfirmed and likely wrong.** `all_schemas.hcl`'s `aws_s3_bucket` block
  carries no `frozen_since` and no `suppress_*` — its schema is only pinned in
  the legacy `suppressions_checkout.txt`, and nothing here confirms it fails
  `gateType`. Before using it as a known-breaker fixture, run `gateType` against
  its live schema and confirm the outcome; if it passes, pick a different fixture
  or drop the claim that this exercises `failed-generation`.

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

1. **Change detection (§4). Done — `change.go`.** `detectChanges` fetches AWS
   bytes (via the shared discovery sweep), byte-compares against the cache,
   respects `frozen_since`, and classifies each type. `gateCandidates` narrows
   to New + Changed. Not yet done: writing the compared bytes back to the cache
   (still read-only) and calling this from `run()`.
2. **Recipe core (§7 prereq). Done — `plan.go`.** `generationPlan` derives
   artifacts, Terraform type, package, and paths per row, in memory, honoring
   `suppress_*`. ~95 lines; no directive-file emission, no subprocess.
3. **Gate mechanism (§6). Done — `gate.go`.** `gateType` runs the front half
   per artifact via `shared.NewResource` + `codegen.Emitter` directly (bypassing
   `shared.GenerateTemplateData`'s CWD hazards — see Corrections, above).
   `runGate` fans candidates out with `errgroup`. Per the plural-DS correction
   above, only `artifactResource` and `artifactSingularDataSource` produce a
   meaningful `failed-validation`/`failed-generation`; treat plural-DS gate
   outcomes as advisory until the compile gate exists.
4. **Wiring.** Not built. `run()` still ends at reconcile-and-report; discovery,
   change detection, recipe derivation, the gate, policy, and block mutation are
   all correct in isolation but not connected. This is integration work, not
   new design: feed `discovered.schema` into `detectChanges`, write refreshed
   bytes to the cache for New/Changed types (a New type's `plan.schemaFile`
   only resolves to a real file once this exists), restrict `generationPlan` +
   `runGate` to `gateCandidates`, run `decide` on each result, and apply the
   decision — `setBlockAttributes` for existing blocks, and a new "synthesize a
   block's text from a fresh `resourceRow`" step for New types (not yet built;
   `mutate.go` only edits existing blocks).
5. **Block mutation in the overlay. Done — `mutate.go`.** `setBlockAttributes`
   sets `frozen_since` / `suppress_*` / `non_provisionable` on an existing block
   via `hclwrite`, preserving comments and existing attributes. Still open:
   rendering a brand-new block's text for a New+failed decision (today's
   `canonicalBlock` in `main.go` only handles the plain New+ok case with no
   suppression attributes).
6. **Result taxonomy + policy mapping. Done — `policy.go`.** `decide(class,
   gateResult, today)` maps (class × result) → a `policyDecision` per the §7
   matrix, table-tested, suppressing only the artifacts that actually failed
   (see Corrections, above).
7. **Compile gate (§6 tail, §10).** Not built. Per-type compile is impractical
   (needs package context). Keep it as a single whole-repo `go build` /
   `make build` after edits are applied; a build failure feeds the same policy
   (freeze the present offender / suppress the new one) and re-runs. This is
   the only gate that can catch a plural-DS-only failure.

## Build plan (incremental, each independently testable)

- **Stage A — selective refresh + changed set. Done, minus cache write.**
  `change.go` implements the compare; the write-back and the `run()` wiring
  remain (folded into prerequisite #4 above).
- **Stage B — gate runner. Done.** `gate.go` implements this exactly as
  specified, with the CWD hazard resolved rather than deferred, and runs
  serially rather than concurrently for a verified reason (see Corrections).
- **Stage C — policy engine. Done.** `policy.go`'s `decide` matches the §7
  table, table-tested, pure — no AWS, no generation.
- **Stage D — block mutation. Done for existing blocks.** `mutate.go`'s
  `setBlockAttributes` preserves comments and fmt's the result via `hclwrite`.
  Not yet done: synthesizing a new block's text for a New+failed decision (see
  prerequisite #5).
- **Stage E — orchestration.** Wire everything into the pipeline: reconcile →
  refresh (A) → gate (B) → policy (C) → mutate (D) → report → optional build
  (prerequisite #7). Keep the structural `-check` untouched and offline. Not
  started; this is prerequisite #4's wiring.
- **Stage F — later.** Self-heal re-probe of frozen types to scratch (§7 last
  row); machine-readable report (§8).

## Required bigdiffer refactors

- **Recipe/refresh module — done.** `plan.go` (recipe) and `change.go` (refresh
  compare) are the bigdiffer-owned replacements for `schema/main.go`'s
  `Downloader.Schemas` + `ResourceSchema`, with no directive emission. Remaining:
  the cache write-back inside `change.go`'s path (prerequisite #4).
- **`run()` becomes a pipeline — not started.** Currently: read → `normalize` →
  write. Target: read → reconcile → refresh → gate → policy → emit → report.
  `normalize` stays the reconcile/emit core; policy edits are applied to the
  item set before the final emit. All the stages this wiring connects
  (discovery, `detectChanges`, `generationPlan`, `runGate`, `decide`,
  `setBlockAttributes`) already exist.
- **Item model gains mutation — done, differently than planned.** The original
  plan called for a text-blob "set attribute" operation on an item. `mutate.go`
  instead re-parses a single block with `hclwrite` and re-serializes it — a
  token-preserving edit rather than string surgery, which is a better fit for
  the design's "verbatim preservation" commitment (§9) than manual line editing
  would have been. Still open: synthesizing a brand-new block's text for a
  New+failed decision (prerequisite #5) — `mutate.go` only edits blocks that
  already exist.
- **Report gains gate results — not started.** New/changed types with their
  result and the applied policy action, for the human review step (§8).

## Open decisions

- **Copy & redesign, don't reuse (decided, done).** `plan.go` is bigdiffer's own
  ~95-line recipe core rather than a shell-out to `schema/main.go`. No temp
  overlay, subprocess, or directive-file round-trip.
- **New-type gating is no longer special (resolved).** Because recipes are
  derived in memory from a `resourceRow`, a New type's recipe is computed
  directly — nothing is written to disk to gate it, so no tentative overlay is
  needed.
- **`shared` dependency (resolved, narrower than expected).** The gate uses
  `shared.NewResource` + `shared/codegen.Emitter` — not `GenerateTemplateData` —
  specifically because that avoided the CWD hazards outright (see Corrections).
  This is a smaller, safer surface of `shared` than originally planned to depend
  on, and needs no further refactor for concurrency.
- **Change detection source (resolved).** Byte-compare in bigdiffer, using the
  bytes already captured by `discover.go`'s single sweep — no `git diff`, no
  second AWS fetch.
- **Gate concurrency (resolved, differently than planned).** The plan assumed
  in-process gating would be safe to parallelize across candidates. The race
  detector disproved this: the reused legacy emission engine mutates global
  state during emission. `runGate` gates serially instead. Making it safe to
  parallelize (auditing and rewriting the emitter) is deferred — it is a
  performance optimization on a step that is already cheap because the
  candidate set is small (§6), not a correctness requirement.

## Verification strategy

- Unit: all built pieces now have test files — policy matrix (`policy_test.go`,
  pure), block mutation (`mutate_test.go`, fixtures), gate outcomes
  (`gate_test.go`, real committed schema + error paths), change detection
  (`change_test.go`), recipe derivation (`plan_test.go`), and the naming filter
  (`discover_test.go`). The suite is `-race`-clean at `-count=5`. Result parsing
  from a stub generator is obsolete: the gate is in-process, not a subprocess.
- Integration (known-breaker, TBD). Do not assume `AWS::S3::Bucket` fails the
  gate — its overlay block carries no `frozen_since`/`suppress_*`, only a
  `suppressions_checkout.txt` pin, and that pin's reason is unconfirmed against
  `gateType`. Before writing this test, run `gateType` against a handful of
  currently-frozen types' *live* schemas and use whichever one actually returns
  `failed-validation` or `failed-generation`. AppFlow `ConnectorProfile` (frozen
  for an `isSandboxEnvironment`/`IsSandboxEnvironment` naming collision, issue
  #1526) is a plausible candidate but is also unconfirmed — it may generate
  cleanly and only produce semantically wrong output, which the gate cannot
  detect (see the correctness-oracle caveat, next).
- Correctness-oracle caveat. The gate proves "does it parse and emit," not "is
  the output right." A freeze whose root cause is valid-but-wrong output (as
  AppFlow's may be) will gate as `ok`, and naive self-heal logic (§7 last row)
  would then propose an incorrect un-freeze. When policy mapping (Stage C) is
  built, either exclude semantic freezes from the self-heal candidate set or
  require a human-reviewable diff of generated output, not just a gate pass.
- Integration (known-good): gate a simple stable type (e.g. `AWS::Logs::LogGroup`,
  already used in `plan_test.go`) and expect `ok` with no overlay edit.
- Idempotence: a second run with no AWS change produces no edits.

