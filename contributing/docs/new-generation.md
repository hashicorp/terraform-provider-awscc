<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# New Generation Process (Design)

This is the design for a new tool called **bigdiffer** (`internal/tools/bigdiffer`).
Its job is to regenerate the Terraform AWSCC provider from the current live
CloudFormation (CFN) types while protecting the provider from additions and
updates that would break it. It is organized and precise, communicates amply,
and treats its users' time as valuable through automation.

Status: design proposal. It defines the model, the data flow, and the decision
rules — not a specific implementation. Sections 1–8 define the target design;
section 9 records decisions already proven in the current tool; sections 10–11
map today's process onto it.

Scope: generate resources, data sources, and list resources from current CFN
schemas, plus docs. Release mechanics (tagging, changelog, PRs) are out of scope.

Naming: the legacy `make bigdiffer` target is an unrelated `git diff` of dated
`available_schemas` files. It is removed by this design; until then, "bigdiffer"
means `internal/tools/bigdiffer`.

## 1. Core insight

The provider already records what it is — `internal/provider/all_schemas.hcl`
(the "overlay") — and AWS already records what exists — the live CFN registry.
Updating the provider is just reconciling those two facts.

Today they are not reconciled directly. AWS's live list is serialized to a dated
`available_schemas.<date>.hcl`, that file is `git diff`ed against last week's
copy, and a human hand-transcribes the result into the overlay. The dated files
feed nothing but the diff. Remove that detour and the process is: read the
overlay, ask AWS what exists, reconcile, generate resiliently, write the overlay
back. No git, no dated snapshots, no manual transcription.

## 2. Model: two inputs, one output

```
  AWS registry ──▶ ┌────────────┐
  (live types)     │ Reconciler ├──▶ updated all_schemas.hcl (+ change report)
  all_schemas ───▶ └────────────┘   join on CFN type name, in memory
```

- Input A — live AWS types: `ListTypes` (public, live, `FULLY_MUTABLE` +
  `IMMUTABLE`) plus `DescribeType` for each schema. Per type this yields the
  Terraform type name and whether a plural data source is supported (a `list`
  handler with no required arguments).
- Input O — the overlay: every type the provider tracks, plus per-type policy
  (`suppress_*_generation`, `frozen_since`, `non_provisionable`,
  `suppression_reason`).
- Output — the overlay, rewritten in place, plus a change report.

The join key is the CFN type name (e.g. `AWS::S3::Bucket`); no other correlation
artifact is needed.

## 3. Change classes

The join yields three classes with mechanical default handling:

| Class | Definition | Handling |
|-------|-----------|----------|
| New | in A, not in O | probe; add to overlay |
| Present | in A and O | refresh unless frozen |
| Absent | in O, not in A | probe, then classify (below) |

Absent is not automatically "withdrawn." Because A excludes `NON_PROVISIONABLE`
types, an absent row is resolved by one `DescribeType` probe:

- Probe succeeds, `NON_PROVISIONABLE`: still live, only filtered out. Annotate
  `non_provisionable = true`; do not freeze.
- Probe returns `TypeNotFoundException` (or `DEPRECATED`): genuinely withdrawn.
  Keep cached bytes and set `frozen_since`.

The provider ships non-provisionable-but-live resources (e.g.
`AWS::AppStream::StackFleetAssociation`), which a naive "absent ⇒ freeze" rule
would wrongly freeze. Only the probe distinguishes the two.

## 4. Selective refresh, not teardown

`make schemas` downloads only schemas that are *missing*, then validates — it
never compares against or replaces an existing file. Today an update is forced
by deletion: `make cleanschemas` removes every `AWS_*.json` so the next
`make schemas` re-downloads them all. Pinning falls out of the same mechanism —
`make suppressions` git-restores frozen JSONs first, so `make schemas` finds them
present and leaves them untouched. This is a clunky stand-in for compare →
replace: delete everything to trigger a full re-download.

Replace it with a direct compare — download a type only when AWS's bytes differ
from the cache, and skip frozen types outright:

- Frozen types: not re-downloaded during the normal refresh — last-good bytes
  stay (no delete/restore). The self-heal probe (§7) may re-fetch them to a
  scratch location to test recovery, but never disturbs the pinned bytes unless
  recovery is accepted.
- Present, non-frozen: downloaded only if AWS's bytes differ; a byte-identical
  refresh is a no-op that skips the gate.
- New types: downloaded once.

So `cleanschemas` and `suppressions` cease to exist as steps.

## 5. One declarative suppression surface

The overlay is the single policy surface, per block:

- `frozen_since` — pin: do not refresh cached bytes. This subsumes
  `suppressions_checkout.txt`; the pinned set is simply "blocks with
  `frozen_since`."
- `suppress_resource_generation`, `suppress_singular_data_source_generation`,
  `suppress_plural_data_source_generation` — do not generate that artifact.
- `non_provisionable` — cannot be provisioned via Cloud Control.
- `suppression_reason` and free-form comments — human rationale, preserved
  verbatim.

Pinning and generation-suppression are orthogonal and may co-occur, but they now
live in one file with one mental model. `suppressions_checkout.txt` and
`make suppressions` are retired.

## 6. One resilient gate

Today two global gates (validate, then generate) each traverse the whole corpus
(10–15 min each) and surface one failure at a time, forcing a rerun per broken
schema. Replace them with a single pass:

1. Merge validate and generate — validation is the front half of generation.
2. Make the pass per-type and resilient — one type's failure never aborts the
   others. Each type yields `ok`, `failed-validation`, or `failed-generation`
   with its error; independent types run concurrently.

Only changed types enter the gate (a byte-identical refresh is a no-op, §4), so
in a typical week the gate runs on a handful of types, not the corpus.

The gate is the crux, not the control flow. Today's generators traverse the
whole overlay and abort on first error; they are neither per-type nor resilient.
Making generation callable per type with per-type error capture — without paying
a process-per-type cost at ~1500 types — is the single largest implementation
risk here. Everything else is bookkeeping by comparison.

Compilation is a final gate: a type can generate valid-looking code that fails to
build. A build failure feeds the same policy as a generation failure (§7) —
suppress the offending new type, or freeze a broken present one — so it too can
require touching the overlay.

## 7. Policy: results become overlay edits

Change class plus gate result determine the overlay edit:

| Change class | Gate result | Overlay action |
|--------------|-------------|----------------|
| New | ok | add plain block |
| New | failed | add block, suppressing only the artifacts that failed + reason (backlog) |
| Present | ok | keep block; accept refreshed bytes |
| Present | failed | set `frozen_since`; keep last-good bytes; flag |
| Non-provisionable (live) | n/a | keep block; set `non_provisionable = true` |
| Withdrawn | n/a | set `frozen_since`; keep bytes |
| Frozen / suppressed | ok on re-probe | propose un-freeze / un-suppress |

Two invariants keep every row safe by default:

- Never regress a shipped resource — a broken `Present` type is frozen at its
  last-good bytes, not failed.
- Never let a broken new type block the release — a broken `New` type is added
  suppressed and becomes backlog.

The last row is self-healing: policy is recomputed every cycle, so a type AWS has
since fixed is detected and proposed for recovery automatically. For frozen types
this means re-fetching to a scratch location and re-running the gate there (§4);
the pinned bytes are replaced only once recovery is accepted.

## 8. The one human step

Everything above is mechanical. The human reviews the change report before
commit: confirm or override auto-applied freezes/suppressions, supply
`suppression_reason` text and issue links, and adjudicate flagged anomalies. The
report is machine-readable, so it also seeds release notes.

## 9. Design decisions (proven in the current tool)

These choices are already made in `internal/tools/bigdiffer` and are adopted as
design commitments:

- Verbatim preservation. Existing blocks are rewritten byte-for-byte by a
  text scanner that delimits blocks on their braces — comments, suppression
  reasons, and fully commented-out blocks survive. A `gohcl` round-trip cannot
  do this; it discards comments.
- Cross-validation. The scanner's set of live blocks is checked against a real
  HCL decode (`hclparse` + `gohcl`); a mismatch is a hard error. This guards the
  verbatim text path with a structural check and yields per-block attributes
  (`frozen_since`, `non_provisionable`).
- Naming invariant. `resource_type_name` must equal the deterministic transform
  of the CFN type name; violations are reported (the classic copy-paste error).
- Explained-retained rule. A row absent from the base is benign only when
  explained by `frozen_since`, `non_provisionable`, or a checkout pin; otherwise
  it is flagged.
- Deterministic output. Blocks are sorted by CFN type name (matching the
  generator) and the header count is updated, giving a stable, reviewable file.
- Idempotent, two check modes. Running twice changes nothing. A structural
  `-check` verifies normalization and anomalies with no AWS access (safe on
  every PR); a live drift report against the registry drives the weekly run and
  is advisory, never a hard PR gate (the overlay is expected to lag live AWS).
- Self-contained. The tool owns the overlay's shape and its own discovery rather
  than importing the legacy generator/update packages, so those can be deleted.
- Resilient, concurrent discovery (`discover.go`). `DescribeType` fans out over a
  bounded worker pool (`errgroup`, limit 10); a per-type failure is captured
  in-band (`discovered.err`) rather than aborting the crawl. The same "one type's
  failure never blocks the rest" principle §6 requires of the gate is already
  proven here, one phase earlier.
- Gate without the CWD-coupled wrapper (`gate.go`). The gate deliberately calls
  `shared.NewResource` + `codegen.Emitter.EmitRootPropertiesSchema` directly
  instead of `shared.GenerateTemplateData`, because that wrapper writes
  `last_resource.txt` to the CWD (a data race under concurrency) and reads
  `../identity/names/services.hcl` via a CWD-relative path (a false failure
  outside `internal/provider`). This is a cleaner in-process front half than
  reusing the existing wrapper, and it resolves in one step what would otherwise
  be a `shared` refactor.
- Plural data sources are not schema-emission gated (`gate.go`). A plural data
  source is generated from the CFN type name alone, not via
  `codegen.EmitRootPropertiesSchema`, so it cannot fail the front-half gate by
  construction. `gateType` only runs the gate for the resource and singular data
  source artifacts; a plural-DS-only failure is left to the compile gate (§6
  tail).
- Gating runs serially today, but the engine is concurrency-clean (`naming.go`,
  `gate.go`). The data race the detector caught was localized to the shared
  pluralizer — the legacy `naming.Pluralize` calls `inflection.AddIrregular` on
  every invocation, mutating that library's global rules — not to the emitter.
  bigdiffer copies the naming logic and serializes that one unsafe dependency
  (`naming.go`). On inspection `codegen.Emitter` has **no mutable package-level
  state** (its only global is a read-only slice), so it is safe to drive from
  concurrent goroutines. The current serial gate was therefore a defensive
  interim from when the gate was the sole consumer, not a hard limit: the only
  remaining in-process concurrency hazards are two CWD-coupled lines in
  `shared.GenerateTemplateData` (the `last_resource.txt` write and the
  CWD-relative `services.hcl` read), both fixable by owning a corrected copy.
  Concurrent in-process generation is thus viable, and is the plan for the
  generation step (Brick 6, see `bigdiffer-generation.md`).

## 10. What is eliminated

| Current step(s) | Fate |
|-----------------|------|
| `make cleanschemas` | Removed — refresh is selective (§4). |
| `make suppressions` + `suppressions_checkout.txt` | Removed — folded into `frozen_since` (§5). |
| `make schemas` (download-if-missing) | Replaced by a compare-and-replace AWS sweep (§2, §4). |
| `make biglister` + dated `available_schemas.<date>.hcl` | Removed — reconcile live AWS against the overlay in memory (§1). |
| `make bigdiffer` (git diff of dated files) | Removed — nothing to diff (§1). |
| Manual overlay edit | Automated — the reconciler writes blocks (§3, §7). |
| Validate gate (`make schemas`, its validation half) + rerun loop | Collapsed into the resilient pass — validation is its front half (§6). |
| Generate — `make resources`, `make singular-data-sources`, `make plural-data-sources` + rerun loops | bigdiffer **owns** the engine and generates in-process and in parallel, only for New + Changed types; validation is the front half (§6). |
| `make docs-all` | bigdiffer orchestrates it — owns `docs-import`, invokes the standard `tfplugindocs` + `docs-fmt`; it does not reimplement `tfplugindocs`. |
| 3000+ `//go:generate` directives + the three `resources.go` / `*_data_sources.go` directive files | Replaced by in-process parallel generation and a single blank-import registration file (§6, `bigdiffer-generation.md`). |
| `make build` | Kept as the final compile gate; its failures feed the same overlay policy as §7. |

Nothing above is deleted in the initial cutover. The legacy generators, directive
files, and `make` targets stay as a documented, deprecated fallback ("kept for
fallback, do not use, slated for removal"). The cutover is **conceptual** —
bigdiffer relies on none of it, proven by an import-boundary test (no import of
`internal/provider/generators/**`) plus full-corpus parity — and the legacy code
is removed only after parity has held for several cycles.

## 11. Roadmap (gaps in the current tool)

Status by design section. "Built" means the code exists and is unit-tested;
"wired" means `run()` actually calls it. As of this writing, discovery, change
detection, recipe derivation, and the gate are each built and independently
tested, but none are yet wired into `run()` — the pipeline still ends at
reconcile-and-report (§1–§3's additive half).

| Piece | Design ref | Status |
|---|---|---|
| Reconcile overlay ↔ live AWS, additive + report | §2–§3 | Wired |
| Concurrent discovery, one sweep captures schema bytes | §2, §9 | Built + wired (discovery only; bytes unused downstream yet) |
| Absent-row `DescribeType` probe (non-provisionable vs. withdrawn) | §3 | Not built |
| Selective refresh / change detection (byte-compare vs. cache) | §4 | Built (`change.go`), not wired |
| Recipe derivation (artifacts, naming, paths per row) | §7 (prereq) | Built (`plan.go`), not wired |
| Gate (in-process front half, per-artifact) | §6 | Built (`gate.go`), not wired |
| Policy mapping (class × gate result → overlay edit) | §7 | Built (`policy.go`), not wired |
| Block mutation (`frozen_since` / `suppress_*` written to an existing block) | §7 (prereq) | Built (`mutate.go`), not wired |
| Compile gate | §6 tail | Not built — planned whole-repo `go build ./internal/...` (Brick 11) |
| Own the generation engine (copy emitter / template-data / writer / templates) | §6, §10 | Not built (Brick 6) |
| Generate in-process, keep output, parallel + incremental | §6, §10 | Not built (Bricks 7–9) |
| Single blank-import registration file (replaces the directive files) | §10 | Not built (Brick 9) |
| Docs — own `docs-import`, orchestrate `tfplugindocs` | §10 | Not built (Brick 10) |
| Retire checkout file into `frozen_since` | §5 | Not built — still read as an external cross-reference |
| Drop the dated-snapshot path | §1 | **Done** — snapshot modes and `-discover` removed; discovery is only via `-update` |
| Self-healing re-probe, machine-readable report | §7 last row, §8 | Not built |

Remaining work is the generation phase, broken into digestible bricks. bigdiffer
**owns** the generation engine and replaces the legacy `make`/`go:generate`
machinery entirely — generating in-process and in parallel, keeping the output —
rather than importing it or re-emitting its 3000+ directives. The detailed
design, the generator-surface investigation, and the anticipated challenges live
in `bigdiffer-generation.md`. In dependency order:

1. **Brick 6 — config + own the engine (single type).** Copy the emitter,
   template-data assembly, file writer, and templates into
   `internal/tools/bigdiffer/codegen` (CWD fixes, in-memory bytes, pluralization
   from the `plan`); rewire the gate onto it; drop the `generators/**` imports;
   add the import-boundary test.
2. **Brick 7 — full-corpus parity (the lynchpin).** Generate all types into a
   temp tree and byte-compare against the committed `*_gen.go` +
   `import_examples_gen.json`. Proves the owned engine is faithful.
3. **Brick 8 — parallelize. Done.** `errgroup`, `-race`-clean, still
   byte-identical; 7.4× over serial.
4. **Brick 9 — incremental weekly pipeline. Done.** `-generate` (full offline
   parallel regeneration) and `-update` (live: discover → detect → regenerate only
   New/Changed from fresh bytes → policy → promote-on-success / freeze-or-suppress
   on failure, never regress → cache write-back → emit the single registration
   file + `import_examples_gen.json` → report). Engine now loads schemas in-memory
   (no CWD side effect).
5. **Brick 10 — docs. Done.** `-docs` owns `docs-import` (the import-example docs,
   copied verbatim from the legacy generator; parity-checked) and orchestrates
   `terraform fmt` + `tfplugindocs generate` — it does not reimplement
   `tfplugindocs`.
6. **Brick 11 — cutover polish. Done.** Step narration + progress bars; the weekly
   release (`generating-the-provider-with-bigdiffer.md`) and the legacy fallback
   documented; legacy `make` targets marked deprecated + `make bigdiffer-*`
   shortcuts added. Then the deferred items below.

Deferred beyond the bricks: absent-row probe (§3), checkout-file retirement (§5),
self-heal + machine-readable report (§7–§8), and the
eventual deletion of the legacy generators/directive files/`make` targets — only
after parity has held for several cycles.

## 12. Thesis

Reconcile the live AWS registry directly against `all_schemas.hcl`, refresh only
what is not frozen, generate every changed type resiliently in one pass, and let
one declarative policy — keyed on change-class × gate-result — write the overlay
back, so git diffs, dated snapshots, the checkout file, the teardown, the double
gate, and manual transcription all disappear.
