<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer: generation & orchestration (Bricks 6–11) — DONE

Design record for the generation phase — moving code generation off the legacy
`make`/`go:generate` machinery and into bigdiffer, which **owns the whole engine,
generates every artifact in-process and in parallel, and keeps the output**. All
bricks below (6–11) have shipped; this stays as the historical design record and
generator-surface investigation rather than being discarded, since the reasoning
(parity-first, then concurrency, then incremental) and the surface-level findings
(registration is self-contained, pluralization hazard, etc.) remain the reference
for anyone touching the engine later. Section refs point at `new-generation.md`,
whose §11 roadmap tracks what is still outstanding beyond this phase (compile
gate, absent-row probe, checkout retirement, self-heal).

## Thesis

The legacy flow regenerates all ~1580 types on every run via 3000+
`//go:generate` directives — each one `go run`-recompiling a generator and
running serially. That is the design we are replacing, not extending. bigdiffer
instead:

- **owns the engine** — a copy of the emitter, template-data assembly, file
  writer, and templates, with no import of `internal/provider/generators/**`;
- **generates in-process and in parallel**, so a full regeneration is fast;
- **generates only New + Changed** on the weekly run (the byte-compare already
  identifies them), falling back to a full run on demand;
- **keeps the output** — validation is the front half of generation (§6), so the
  same pass that gates produces the `*_gen.go`.

Legacy is not deleted in this work. It is kept as a documented, deprecated
fallback ("kept for fallback, do not use, slated for removal"); the cutover is
**conceptual** — bigdiffer relies on none of it — and proven by an
import-boundary test plus full-corpus parity.

## Runtime vs. generation (the separation)

- **Generation** = bigdiffer: the tool at `internal/tools/bigdiffer` (idiomatic
  home for a repo-internal, non-shipped dev tool) plus the copied engine in a
  subpackage (`internal/tools/bigdiffer/codegen`).
- **Runtime** = the emitted output (`internal/aws/**` and the registration file)
  and the runtime packages that output references — `internal/registry`,
  `internal/generic`, `internal/identity`, `internal/validators`,
  `internal/defaults`. Runtime is untouched by this work; the generated output
  is **not** relocated (moving `internal/aws/**` is enormous churn for no runtime
  gain).
- **Assurance:** an import-boundary test asserts bigdiffer imports nothing under
  `internal/provider/generators/**`. It goes green once the engine copy lands and
  `gate.go`'s two current `generators/**` imports are gone. The residual
  conflation — legacy generation living under `internal/provider/generators` —
  resolves when legacy is eventually deleted.

## Investigation findings (the generator surface)

Grounded by reading the generators, not assumed:

- **Registration is self-contained.** Each generated file self-registers in its
  own `func init()` via `registry.AddResourceFactory` /
  `AddListResourceFactory` / data-source equivalents. There is **no central
  factory list**. List resources are not a separate artifact — the resource
  generator emits them via `-listresource` → `AddListResourceFactory`.
- **The only build-time aggregate is a set of blank imports.** The three
  `internal/provider/{resources,singular_data_sources,plural_data_sources}.go`
  files contain nothing but `//go:generate` directives (inert comments) and a
  `package provider` + blank-import block. Go runs a package's `init()` only if
  that package is in the binary's import graph, so those `_ "…/internal/aws/<svc>"`
  imports are what make registration happen. bigdiffer therefore emits **one
  collapsed registration file** (blank imports only, no directives). The three
  legacy directive files stay as deprecated fallback; bigdiffer does not
  re-emit them. During coexistence, duplicate blank imports across files are
  legal Go and harmless.
- **Output layout:** `internal/aws/<svc>/<res>_<artifact>_gen.go` (+ `_test.go`),
  `-package <svc>`, cfschema from `internal/service/cloudformation/schemas/`.
  `plan.go` already derives `<svc>`, `<res>`, the file names, and the safe
  plural Terraform name.
- **The engine to own (copy, do not import):**
    - `codegen.Emitter` (~1500 lines). Its `EmitRootPropertiesSchema` surface — the
      one `gate.go` already drives — is **concurrency-clean**: `emitter.go`'s only
      `naming` calls are the pure `CloudFormationPropertyToTerraformAttribute`, and
      its only package-level `var` is a read-only slice. So the copy needs no
      structural change for concurrency; it is copied for **ownership** (so legacy
      can be deleted), not to fix a bug.
    - `shared.GenerateTemplateData` — assembles `TemplateData`. Two CWD hazards to
      fix in the copy: drop the `last_resource.txt` debug write, and resolve the
      CWD-relative `ParseServicesFile("../identity/names/services.hcl")` via an
      absolute config path. Adapt it to the in-memory sanitized bytes the discovery
      sweep already holds (build the doc via `cfschema.NewResourceJsonSchemaDocument`).
      Note it does **not** pluralize (its only `naming` call is
      `ParseCloudFormationTypeName`).
    - `common.Generator` — parses each template per-call (concurrency-safe) with a
      small funcMap (`Title`, `Split`), runs `go/format.Source`, writes the file.
    - Six templates (`schema.tmpl` + `tests.tmpl` for resource / singular-DS /
      plural-DS), `//go:embed` in the copy.
    - `naming.SnakeCase` (used for `PrimaryIdentifier`) — safe (local regexes);
      ported to `bigdiffer/naming` (the `internal/tools/bigdiffer/naming`
      subpackage) alongside `Pluralize`/`PluralizeWithCustomNameSuffix`.
- **Pluralization is the one `inflection` hazard, and it is already solved.**
  The only calls to `naming.Pluralize` / `PluralizeWithCustomNameSuffix` in the
  generation path are the index/directive emission (`schema/main.go:286`) and
  the plural-DS generator (`plural-data-source/main.go:84`) — the exact global
  mutation Brick 3/4 found and fixed with `bigdiffer/naming`'s mutex-guarded
  `pluralize`. bigdiffer's `plan.go` already computes the plural name via that
  safe function, so the rule is simply: **generation consumes the plural name
  from the `plan`; it never calls `internal/naming.Pluralize`.** The hazard is
  avoided by construction.
- **The plural-DS path** builds `TemplateData` from the CFN type name alone (no
  schema emission), so it cannot fail the schema-emission gate by construction
  (§9). Confirm its exact inputs when copying it.
- **A fourth aggregate:** `import_examples_gen.json`, emitted from schemas
  (`import-examples/main.go`). bigdiffer emits it too, or it drifts.
- **`x-derecursed`** is understood and dormant: legacy already has the key check,
  it has been run, and it changes nothing on the current branch (zero cached
  schemas carry the key). The copy must preserve `Emitter.Deduplicate` (inherited
  via `schemaIsDeRecursed`), but it needs no special handling now.

## Pipeline (the weekly `-update` path)

```text
discover(ctx)                         # rows + sanitized bytes (have it)
  → detectChanges(disc, frozen, cache)  # new/changed/unchanged/frozen/missing
  → for each candidate (New|Changed), concurrently (errgroup):
        plan = generationPlan(row)        # carries the safe plural name
        emit → temp: *_gen.go (+ _test.go) via the owned engine (= validate)
        dec  = decide(class, outcome, today)
        if not a freeze: stage {promote temp→real files, promote bytes→cache}
  → apply decisions to overlay (mutate existing / synthesize new blocks)
  → emit the single registration file + import_examples_gen.json
  → re-sort overlay, write it
  → compile gate: go build ./internal/...   (failure feeds policy, re-runs)
  → report (per-candidate outcome + action)
```

Never-regress: emit to a temp dir, promote only on success; a broken Changed type
keeps its existing files + last-good cache and is frozen; a broken New type is
added suppressed and nothing is promoted for the failed artifact.

## Full generation, and how parity works

`-full` is a first-class, bigdiffer-native parallel regeneration of the whole
corpus — fast because it is in-process and concurrent, not a legacy directive
crawl. It is the correct thing to run whenever the **engine or templates change**
(then even a byte-identical schema can produce different output).

**Parity oracle — free, no need to run legacy.** The committed `*_gen.go` files
*are* the legacy output. Cutover correctness = run bigdiffer over the full corpus
into a temp tree and byte-compare against the committed files (and against
`import_examples_gen.json`); an empty diff proves the owned engine introduced no
drift. This is the single most valuable thing to build early.

**Scope: the full corpus (~1580 types), not a sample.** A sample risks missing
exactly the heterogeneity that matters — recursive schemas, write-only
properties, primary identifiers with slashes, the dormant `x-derecursed` path,
types whose flags come from `services.hcl`. Generation is cheap under this
design, so the full corpus is the only sample that actually answers "did we drop
a behavior."

**Near-term expectation, stated plainly.** While the owned engine and templates
are still being built and stabilized, the engine changes across most cycles, so
most runs in that period are `-full` runs, not incremental ones. The New+Changed
speedup is real but only materializes once the copy stops changing week to week.
Early runs looking like whole-corpus regenerations is the transitional cost of
owning the engine, not a regression.

**Engine changes vs. incremental staleness.** `detectChanges` sees only schema
bytes, so an engine/template change that alters output for an unchanged schema
would be silently skipped in incremental mode. Because parallel `-full` is cheap,
this mostly dissolves: on any engine/template change, run `-full`. An optional
engine/template **version stamp** can force `-full` automatically on a mismatch;
until then, treat any engine edit as requiring a manual `-full`. Deciding whether
a given output diff is intended remains a human call on the PR that changes the
engine (the `x-derecursed` toggle is the precedent: an engine-level flag whose
effect on output is not uniform across resources).

## Anticipated challenges (grounded)

1. **Faithful engine copy.** `GenerateTemplateData` assembles more than the
   front-half emit (identifiers, timeouts, write-only paths, framework-import
   flags, `IsGlobal`/`HasMutableIdentity` from `services.hcl`). The CWD fixes are
   trivial; the risk is silently dropping behavior. **Mitigation: full-corpus
   parity, built first.**
2. **`services.hcl` resolution.** Global-resource and mutable-identity flags come
   from it; a wrong path silently changes output. Resolve absolutely via config.
3. **Template drift during coexistence.** Until legacy is deleted, our templates
   can diverge from the legacy ones. Parity guards this.
4. **Concurrency proof, not assumption.** Prove it — `-race` on concurrent
   generation of many types **plus** the parity byte-compare. Include the
   `GenerateTemplateData` and plural-DS paths, not just the emitter surface.
5. **Registration correctness.** The single collapsed registration file must
   blank-import exactly the set of service packages that have registered
   artifacts (adding a new package for a new service). The compile gate
   (`go build ./internal/...`) is the backstop, and the most likely place for a
   subtle "which package is new" bug.
6. **Determinism.** Output must be byte-stable run-to-run (map iteration, sorting)
   for clean diffs and a meaningful parity test.

**Compile-gate scope — decided: whole-repo `go build ./internal/...`.** It is the
only option that covers the registration file; affected-packages-only is blind to
registration breakage in exactly the place that matters. Revisit only if build
time is measured to be a real bottleneck.

## Brick breakdown (dependency-ordered)

- **Brick 6 — Config + own the engine (single type). DONE.** Introduced the
  `config` struct (resolves all paths, region, concurrency, prefix, mode).
  Copied the emitter, `GenerateTemplateData`, `common.Generator`, `SnakeCase`,
  and the six templates into `internal/tools/bigdiffer/codegen`, with the two
  CWD fixes, in-memory bytes, and pluralization sourced from the `plan`.
  Removed the standalone `gate.go` front-half gate in favor of
  generation-is-the-gate; dropped the `generators/**` imports. Landed the
  import-boundary test (later widened to also forbid `internal/naming`
  directly, closing the gap that would let the fixed `inflection` race back in
  silently). *Done = bigdiffer emits one type's code in-process,
  self-contained.*
- **Brick 7 — Full-corpus parity (the lynchpin). DONE.** Generate every type's
  artifacts through the owned engine and compare to the committed `*_gen.go`. The
  true criterion is bigdiffer == legacy, so committed is only the fast first pass:
  on a mismatch the harness runs the legacy generator for that artifact and
  compares against it — equal means the committed file is merely stale (logged,
  non-fatal), only a real bigdiffer-vs-legacy difference fails. Result: 1579
  types, 8432 files, **0 drift**, ~7s in-process. Two stale committed files found
  (`appconfig::extension` resource + singular DS: bigdiffer matches current legacy
  but the release shipped an out-of-date file; they regenerate correctly on the
  next full run). One real drift found and fixed: the copied naming lacked
  internal/naming's `(D|d)ata$ -> _plural` rule. `import_examples_gen.json` is
  **not** in this brick — it is an aggregate from `schema/main.go`'s
  `GenerateResourceImportExamples`, so its parity is grouped with Brick 9.
- **Brick 8 — Parallelize.** `errgroup` over the corpus; `-race` clean; still
  byte-identical; measure the speedup. *Done: serial 5.26s → parallel(16) 710ms,
  7.4×; `-race` clean over 1067 concurrent artifacts.*
- **Brick 9 — Incremental weekly pipeline. Done.** Shipped as two modes plus an
  engine fix:
    - `9.1` `writeCorpus` (writes genResults to `internal/aws/<svc>/`) and
      `emitRegistration` — the single blank-import file that replaces the three
      legacy directive files; its import set is exactly the legacy union (285
      packages).
    - `9.2` `codegen.GenerateImportExamples` (copied `imports.tmpl`) + the
      `emitImportExamples` aggregate, byte-identical to committed
      `import_examples_gen.json` (`TestImportExamplesParity`).
    - `9.3a` **`-generate`** — a full, offline, parallel regeneration of the whole
      provider (`generateCorpus` → `writeCorpus` → both aggregates). The in-process
      replacement for the legacy `make resources/…/schemas` directive crawl; does
      not touch AWS.
    - `9.3b` **`-update`** — the live weekly incremental. One discovery crawl feeds
      both overlay reconciliation and change detection; only New/Changed types are
      regenerated from their fresh bytes; generation success/failure drives the
      policy (`decide`) applied to the overlay in one `normalizeWithDecisions` pass;
      `refreshCandidate` stages fresh bytes in a temp file and **promotes files +
      cache only on clean generation (never regress)**; the aggregates are
      re-emitted. Requires AWS (us-east-1); the live crawl is exercised manually.
    - **Engine fix:** `codegen.NewResource` now loads schemas in-memory via
      `NewResourceJsonSchemaDocument` instead of `NewResourceJsonSchemaPath`, which
      chdir'd into the schema directory without restoring — a CWD-corruption + a
      concurrency hazard. Parity unchanged (0 drift, 8432 files).
- **Brick 10 — Docs. Done.** `-docs` owns `docs-import` and orchestrates the two
  external steps of `make docs-all`. `codegen.GenerateImportExampleDocs` copies
  the legacy import-examples generator verbatim (four templates rendered through
  the same `parseTemplate`/`Split`, same snake-cased pipe-joined identifier), and
  `runDocs` writes `examples/**` then invokes `terraform fmt` (docs-fmt) and
  `tfplugindocs generate` (docs) — it does not reimplement `tfplugindocs`. Parity:
  4206 non-fmt files (`import.sh`, `import-by-string-id.tf`,
  `list-resource.tfquery.hcl`) byte-identical to committed (0 drift); the one
  fmt-touched file (`import-by-identity.tf`), run through `terraform fmt`, matches
  committed. *Done = full cutover capability (resources + data sources + docs).*
- **Brick 11 — Cutover polish. Done.** Step narration (`==>` headers) + progress
  bars (`schollz/progressbar/v3`, copyplop-style) across discover / generate /
  update / docs, silent off a TTY. The two things maintainers care about are
  documented: the weekly release with bigdiffer
  (`contributing/docs/generating-the-provider-with-bigdiffer.md` — `-update`,
  `make build`/`make smoke`, `-docs`, CHANGELOG, commit/PR) and the fallback to
  the legacy `make` path (the legacy doc reframed with a banner; formats
  interoperate; `rm registrations_gen.go`). Legacy `make` generation targets carry
  a deprecation note; `make bigdiffer-update/-generate/-docs` shortcuts added. The
  compile gate stays the documented `make build`; a machine-readable report and
  the deferred design items (absent-row probe, checkout-file retirement,
  self-heal re-probe, parity-validated naming simplification, eventual legacy
  deletion) remain for later.

Parity (7) before concurrency (8) before incremental (9) is deliberate: prove
faithful, then fast, then smart.

## Configuration

An **external config file** is low value: awscc is static, and the overlay's
`defaults` block already centralizes the schema settings
(`schema_cache_directory`, `terraform_type_name_prefix`, `meta_schema`). A second
file would compete with it.

An **in-code `config` struct** is worth doing, for documentation / clarity /
lever-centralization: paths and tunables are scattered across flag defaults and
constants today (`discoverRegion`, `discoverConcurrency`, cache/overlay/checkout
paths, `dateLayout`), and generation adds more (output root, registration-file
path, `import_examples` path, `services.hcl` path, meta-schema path, generation
concurrency, temp dir, full-vs-incremental). Gathering these into one struct
derived once from the overlay `defaults` + flags — and resolving all
relative-to-overlay paths to absolute there — centralizes the levers, documents
them, and improves testability (no brittle CWD-relative assumptions, which
`gate_test.go` already had to work around). This is Brick 6's first step.

## Corrections found in review (post-Brick 11)

Two policy-engine bugs (predating this doc, in Brick 5's `policy.go`/`main.go`)
surfaced during a code review of the shipped tool and were fixed in place, with
regression tests:

- `suppressAttrsForFailures` (`policy.go`) had no case for
  `artifactPluralDataSource`, so a New type whose *only* failure was its plural
  data source was added to the overlay with nothing suppressed — the generated
  plural-DS file for a failing artifact never gets written, but the type's
  registration import would still be emitted, breaking `go build`. Added the
  missing case and the `suppress_plural_data_source_generation` constant.
- `normalizeWithDecisions`'s decision-application loop (`main.go`) mutated any
  item whose `key` matched a decision, without checking `it.live`. A fully
  commented-out block can carry a matching `cloudformation_type_name` via
  `classifyItem`'s comment fallback; mutating its text hits `setBlockAttributes`
  expecting exactly one live block and aborts the whole `-update` run on a
  false collision. Added the `it.live` guard.

Both are covered by new tests (`TestDecide/new_+_plural-only_failure_...`,
`TestNormalizeWithDecisionsSkipsCommentedOutCollision`). The import-boundary
test was also widened in the same pass to forbid `internal/naming` directly
(previously it only forbade `internal/provider/generators/**`), closing a gap
that would let the fixed `inflection` race back in via a stray future import.

## Deferred (beyond this phase)

Absent-row `DescribeType` probe (§3), checkout-file retirement (§5),
self-heal re-probe + machine-readable report (§7–§8),
and the eventual deletion of the legacy generators, the three directive files,
and the legacy `make` targets — only after parity has held for several cycles.

Naming simplification (parity-validated). The `isCustomName` regex list
(`efs`/`tions`/`issions`/`windows`/`settings`/`data`) appears to approximate a
single condition: `inflection.Plural(name) == name` (pluralization left the name
unchanged, so it must be disambiguated). Replacing the list with the general rule
"if the plural equals the input, append `_plural`" would drop the hardcoded
special-cases. It is a behavior change (it would suffix *every* inflection-
unchanged name, not just the listed ones), so it must be proven equivalent by
running the full-corpus parity harness: flip the rule, and if the diff stays 0,
the list was redundant. Deferred to its own brick precisely because parity now
makes it safe to attempt.
