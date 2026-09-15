<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer: design

**bigdiffer** (`internal/tools/bigdiffer`) regenerates the Terraform AWSCC
provider from the current live CloudFormation (CFN) types while protecting the
provider from additions and updates that would break it.

This is the durable design reference — what bigdiffer is, how it works, and *why*
it is shaped this way — for maintenance and for reviewing the tool. For a plain
overview and the available commands, see
[the tool's README](../../internal/tools/bigdiffer/README.md). It is not a
runbook (see `generating-the-provider-with-bigdiffer.md` for the weekly process)
and not a task tracker (see "Deferred and future work" below for what is left to do).

bigdiffer is shipped and drives the weekly cycle today; §11 summarizes what is
built and what remains. Scope: generate resources, data sources, and list
resources from current CFN schemas, plus docs. Release mechanics (tagging,
changelog, PRs) are out of scope.

> Naming note: the legacy `make bigdiffer` target is an unrelated `git diff` of
> dated `available_schemas` files, slated for removal. Here "bigdiffer" always
> means `internal/tools/bigdiffer`.

## 1. Core insight

The provider already records what it is — `internal/provider/all_schemas.hcl`
(the "overlay") — and AWS already records what exists — the live CFN registry.
Updating the provider is reconciling those two facts directly: read the overlay,
ask AWS what exists, reconcile, generate resiliently, write the overlay back.

The legacy process reconciled them indirectly instead, through a dated,
git-diffed snapshot a human hand-transcribed into the overlay (see §10 for the
full comparison).

## 2. Model: two inputs, one output

```text
  AWS registry ──▶ ┌────────────┐
  (live types)     │ Reconciler ├──▶ updated all_schemas.hcl (+ change report)
  all_schemas ───▶ └────────────┘   join on CFN type name, in memory
```

- Input A — live AWS types: `ListTypes` (public, live, `FULLY_MUTABLE` +
  `IMMUTABLE`) plus `DescribeType` for each schema. Per type this yields the
  Terraform type name and whether a plural data source is supported (a `list`
  handler with no required arguments).
- Input O — the overlay: every type the provider tracks, plus per-type policy
  (`suppress_*_generation`, `frozen_since`, `non_provisionable`, and a reason
  per fact: `suppression_reason_resource` / `_singular_data_source` /
  `_plural_data_source`, `frozen_reason`).
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

Absent is not automatically "withdrawn." Because A excludes both
`NON_PROVISIONABLE` and `DEPRECATED` types, an absent row is resolved by one
`DescribeType` probe, reusing the crawl's own throttled client:

- Probe succeeds, `DeprecatedStatus == DEPRECATED`: deactivated/deregistered —
  checked first, since a deprecated type is withdrawn regardless of what
  provisioning type it still reports.
- Probe succeeds, not deprecated, `ProvisioningType == NON_PROVISIONABLE`:
  still live, only filtered out. Annotate `non_provisionable = true`; do not
  freeze.
- Probe fails with `TypeNotFoundException` (matched via `errors.As` through
  the wrapped error, not a string match): genuinely withdrawn/deregistered.
  Keep cached bytes and set `frozen_since` (`manual: withdrawn from AWS,
  pending major-version removal`).
- Anything else (a transient error, or a probe that succeeds live and
  provisionable — a `ListTypes` eventual-consistency blip) resolves to no
  decision: the row is left exactly as it was for a later run, never frozen on
  a guess.

The provider ships non-provisionable-but-live resources (e.g.
`AWS::AppStream::StackFleetAssociation`), which a naive "absent ⇒ freeze" rule
would wrongly freeze. Only the probe distinguishes the two. A row already
explained — frozen, or checkout-pinned — skips the probe entirely, since a
checkout pin may explain an absence the probe cannot see (a deliberately
checked-out older version). Withdrawn and non-provisionable rows feed the same
policy table as the New/Present classes (§7) — no separate taxonomy.

## 4. Selective refresh, not teardown

bigdiffer does a direct compare — download a type only when AWS's bytes differ
from the cache, and skip frozen types outright (the legacy equivalent forced a
refresh by deleting every cached schema and re-downloading; see §10):

- Frozen types: not re-downloaded during the normal refresh — last-good bytes
  stay (no delete/restore). The self-heal probe (§7) may re-fetch them to a
  scratch location to test recovery, but never disturbs the pinned bytes unless
  recovery is accepted.
- Present, non-frozen: downloaded only if AWS's bytes differ; a byte-identical
  refresh is a no-op that skips the gate.
- New types: downloaded once.

## 5. One declarative suppression surface

The overlay is the single policy surface, per block:

- `frozen_since` — pin: do not refresh cached bytes. This subsumes
  `suppressions_checkout.txt`; the pinned set is simply "blocks with
  `frozen_since`."
- `suppress_resource_generation`, `suppress_singular_data_source_generation`,
  `suppress_plural_data_source_generation` — do not generate that artifact.
- `non_provisionable` — cannot be provisioned via Cloud Control.
- `suppression_reason_resource` / `_singular_data_source` / `_plural_data_source`,
  `frozen_reason`, and free-form comments — human rationale, one field per
  fact rather than one shared string, preserved verbatim. The reason taxonomy
  and its enforcement live in `suppressed-and-frozen.md`.

Pinning and generation-suppression are orthogonal and may co-occur, but they now
live in one file with one mental model. (`suppressions_checkout.txt` is still
read as a read-only cross-reference; folding it fully into `frozen_since` is
deferred — see "Deferred and future work.")

## 6. The gates: generation, then compile

bigdiffer has two gates. Both route failures through one policy (§7), and neither
ever blocks a release.

### Generation is the gate

Generation and validation are the same pass: generating a type for real *is*
validating it, so there is no separate emission-only pass to keep in sync. The
pass is per-artifact and resilient — a resource, its singular data source, and
its plural data source each succeed or fail independently, and one failure
never aborts the others. Independent types run concurrently.

Only New and Changed types enter the gate (a byte-identical refresh is a no-op,
§4), so a typical week gates a handful of types, not the corpus. (The legacy
flow's two whole-corpus gates and per-failure rerun loop are gone; see §10.)

### The compile gate

Generation catches schemas that won't render; it cannot catch generated code that
renders but fails to **type-check** — a bad cross-file reference, a collision with
a sibling in the same service package, a broken import, or a stale registration
file. `go build` is that check. bigdiffer runs it itself, inside `-update`,
before promoting anything, so a build failure routes through the same policy
(§7) as a generation failure, before a broken commit can happen.

It is whole-module and serial, not per-artifact like generation, because Go's
unit of build success is the package, not the file: one bad file fails its
whole service package, and building a package while a co-package candidate
concurrently mutates it is unsafe. It runs by overlaying the staged files onto
the real tree, building, and **unconditionally reverting** — the real tree is
bit-for-bit unchanged whether the build passes or fails; only promotion, after
a green build, writes for real.

**Build until green.** On failure, bigdiffer maps each file `go build` blames
back to the candidate that staged it, downgrades only those to a
`build_failed` outcome, drops their staged files, and rebuilds — until the
build is green or a small round cap trips. Each non-green round drops at least
one artifact, so it terminates. A blamed file that maps to nothing staged (a
pre-existing base-tree break, or a bug in bigdiffer's own emitted registration
file) is unattributable: promote nothing and hard-error rather than guess.

**Scope.** The compile gate covers every generated package and the registration
file. It does **not** cover `*_gen_test.go`: `go build` never compiles test
files; a generated test that fails to compile is caught by `make
smoke`/`make test`.

**List-resource coupling guard.** A resource is generated advertising a list
resource only when its plural data source is not suppressed, but the compile
gate can drop the plural data source afterward. A post-fixpoint check refuses
to promote a resource still advertising a list resource whose plural data
source the gate rejected, rather than shipping a list resource with no
backing data source.

## 7. Policy: results become overlay edits

Change class plus gate result determine the overlay edit. "Failed" covers both a
generation failure and a compile (`build_failed`) failure; suppression is
per-artifact — only the artifacts that failed are suppressed, each with its own
categorized reason in its own field
(`suppression_reason_resource`/`_singular_data_source`/`_plural_data_source`);
a freeze, when one applies, carries its own separate `frozen_reason`.

| Change class | Gate result | Overlay action |
|--------------|-------------|----------------|
| New | ok | add plain block |
| New | failed | add block, suppressing only the failed artifact(s) + reason (backlog) |
| Present | ok | keep block; accept refreshed bytes |
| Present | partial failure | suppress the failed artifact(s) + set `frozen_since` (one JSON backs all three, so the schema cannot partially advance) |
| Present | full failure | set `frozen_since`; keep last-good bytes |
| Non-provisionable (live) | n/a | keep block; set `non_provisionable = true` |
| Withdrawn | n/a | set `frozen_since`; keep bytes |
| Frozen / suppressed | ok on re-probe | propose un-freeze / un-suppress |

Two invariants keep every row safe by default:

- Never regress a shipped resource — a broken `Present` type is frozen at its
  last-good bytes, not failed.
- Never let a broken new type block the release — a broken `New` type is added
  suppressed and becomes backlog.

The reason taxonomy (`structural` / `generation_failed` / `build_failed` /
`manual` / `unknown`), the `-check` reason anomaly, and the `-heal` re-probe are
specified in `suppressed-and-frozen.md`. The last table row is self-healing:
policy is recomputed every cycle, so a type AWS has since fixed is proposed for
recovery automatically.

## 8. The one human step

Everything above is mechanical. The human reviews the change report before
commit: confirm or override auto-applied freezes/suppressions, supply the
appropriate per-artifact `suppression_reason_*` (or `frozen_reason`) text and
issue links, and adjudicate flagged anomalies. The report is plain,
consistently-shaped text (one fact per line: CFN type, label, field, reason),
readable by a human or an LLM agent working through proposals interactively.

`-update` also drafts the `FEATURES:` half of the weekly `CHANGELOG.md` entry,
writing directly into the file's top, in-progress version block (e.g.
`## 1.100.0 (Unreleased)`). The bullets come from what the run actually
promoted, filtered to artifacts not already user-visible before this run — a
brand-new type, or a specific artifact whose suppression was just lifted. It
assumes that block's `FEATURES:` section is empty or absent when it runs (the
normal case — bigdiffer runs once per release cycle) and errors rather than
guess at a merge if it already has bullets; add the new ones by hand in that
case. The PR number and any `NOTES:`/breaking-change entries stay a human step.

## 9. Design decisions (proven in the tool)

These choices are settled in `internal/tools/bigdiffer`:

- **Verbatim preservation.** Existing blocks are rewritten byte-for-byte by a
  text scanner that delimits blocks on their braces — comments, suppression
  reasons, and fully commented-out blocks survive. A `gohcl` round-trip cannot
  do this; it discards comments.
- **Cross-validation.** The scanner's set of live blocks is checked against a real
  HCL decode (`hclparse` + `gohcl`); a mismatch is a hard error. This guards the
  verbatim text path with a structural check and yields per-block attributes.
- **Naming invariant.** `resource_type_name` must equal the deterministic
  transform of the CFN type name; violations are reported (the classic
  copy-paste error).
- **Explained-retained rule.** A row absent from the base is benign only when
  explained by `frozen_since`, `non_provisionable`, or a checkout pin; otherwise
  it is flagged.
- **Deterministic output.** Blocks are sorted by CFN type name (matching the
  generator) and the header count updated, giving a stable, reviewable file.
- **Idempotent, two check modes.** Running twice changes nothing. A structural
  `-check` verifies normalization and anomalies with no AWS access (safe on every
  PR); the live drift report drives the weekly run and is advisory, never a hard
  PR gate (the overlay is expected to lag live AWS).
- **Self-contained.** The tool owns the overlay's shape, its own discovery, its
  own naming, and its own copy of the generation engine rather than importing the
  legacy generator/update packages, so those can eventually be deleted. An
  import-boundary test forbids any bigdiffer file from importing
  `internal/provider/generators/**` or the legacy `internal/naming`.
- **Resilient, concurrent discovery.** `DescribeType` fans out over a bounded
  worker pool; a per-type failure is captured in-band, not aborting the crawl —
  the same "one failure never blocks the rest" principle the gates use.
- **The shared pluralizer is not shared.** The legacy `internal/naming.Pluralize`
  mutates the `inflection` library's global rules on every call, making it
  unsafe under concurrency. bigdiffer owns a copy (`internal/tools/bigdiffer/naming`)
  that registers the irregular once behind a mutex. `codegen.Emitter` itself
  has no mutable package-level state and is driven concurrently in production
  (`-race`-clean, ~7× over serial).
- **Generation and runtime are separate; generated output is not relocated.**
  bigdiffer (the tool + its `codegen` copy) is generation; the emitted
  `internal/aws/**`, the registration file, and the runtime packages they
  reference are runtime, untouched by the tool. Moving `internal/aws/**` would be
  enormous churn for no runtime gain.

## 10. What is eliminated

| Legacy step(s) | Fate |
|-----------------|------|
| `make cleanschemas` | Removed — refresh is selective (§4). |
| `make suppressions` + `suppressions_checkout.txt` | Folded into `frozen_since` (§5; checkout retirement deferred). |
| `make schemas` (download-if-missing) | Replaced by a compare-and-replace AWS sweep (§2, §4). |
| `make biglister` + dated `available_schemas.<date>.hcl` | Removed — reconcile live AWS against the overlay in memory (§1). |
| `make bigdiffer` (git diff of dated files) | Removed — nothing to diff (§1). |
| Manual overlay edit | Automated — the reconciler writes blocks (§3, §7). |
| Validate + generate gates + rerun loops | Collapsed into the resilient per-artifact generation-is-the-gate pass (§6). |
| `make resources`, `make singular-data-sources`, `make plural-data-sources` | bigdiffer owns the engine and generates in-process and in parallel, only for New + Changed types (§6, "The generator surface"). |
| `make docs-all` | bigdiffer orchestrates it — owns `docs-import`, invokes the standard `tfplugindocs` + `docs-fmt`; it does not reimplement `tfplugindocs`. |
| 3000+ `//go:generate` directives + the three `resources.go` / `*_data_sources.go` directive files | Replaced by in-process parallel generation and a single blank-import registration file ("The generator surface"). |
| `make build` as the manual compile check | bigdiffer runs the compile gate itself, in `-update`, before promotion (§6); `make build` afterward is a confirmation, not the primary catch. |

Nothing above is deleted in the cutover. The legacy generators, directive files,
and `make` targets stay as a documented, deprecated fallback. The cutover is
**conceptual** — bigdiffer relies on none of it, proven by the import-boundary
test plus full-corpus parity — and the legacy code is removed only after parity
has held for several cycles ("Deferred and future work").

## 11. Implementation status

bigdiffer is shipped and drives the weekly cycle. Built, wired, and tested:
discovery, byte-compare change detection, per-artifact generation-as-gate, the
compile gate, the class × gate-result policy (including the absent-row
`DescribeType` probe, §3), comment-preserving block mutation, the single
blank-import registration file, the import-examples aggregate, and docs
orchestration. Correctness is anchored by full-corpus parity — the owned
engine is byte-identical to the legacy generators (0 drift, ~1580 types) — kept
as a regression guard. The `-check`, `-update`, `-generate`, `-docs`, and `-heal`
modes are all live.

Remaining work — GitHub-issue guidance, the one-time reason backfill, and the
deferred cleanups below — is listed under "Deferred and future work" below.
None of it blocks the weekly cycle today; see
`generating-the-provider-with-bigdiffer.md` for the operational process and the
legacy fallback.

## 12. Thesis

Reconcile the live AWS registry directly against `all_schemas.hcl`, refresh only
what is not frozen, generate every changed type resiliently in one pass, gate it
against a real compile, and let one declarative policy — keyed on change-class ×
gate-result — write the overlay back, so git diffs, dated snapshots, the checkout
file, the teardown, the double gate, and manual transcription all disappear.

---

## The generator surface (reference)

How generation and registration work — reference for anyone touching the
engine.

- **Registration is self-contained.** Each generated file self-registers in its
  own `func init()` via `registry.AddResourceFactory` / `AddListResourceFactory`
  / data-source equivalents. There is **no central factory list**. List resources
  are not a separate artifact — the resource generator emits them via
  `-listresource` → `AddListResourceFactory`.
- **The only build-time aggregate is a set of blank imports.** Go runs a
  package's `init()` only if that package is in the binary's import graph, so a
  `_ "…/internal/aws/<svc>"` blank import per service package is what makes
  registration happen. bigdiffer emits **one collapsed registration file**
  (`internal/provider/registrations_gen.go`, blank imports only), replacing the
  three legacy `internal/provider/{resources,singular_data_sources,plural_data_sources}.go`
  directive files. During coexistence, duplicate blank imports across files are
  legal Go and harmless.
- **Output layout:** `internal/aws/<svc>/<res>_<artifact>_gen.go` (+ `_test.go`),
  `-package <svc>`, cfschema from `internal/service/cloudformation/schemas/`.
  `plan.go` derives `<svc>`, `<res>`, the file names, and the safe plural
  Terraform name, purely from an overlay row.
- **The owned engine** (`internal/tools/bigdiffer/codegen`, copied not imported):
  the emitter (`EmitRootPropertiesSchema`; concurrency-clean — no mutable
  package state); `GenerateTemplateData` (with the CWD hazards fixed — no
  `last_resource.txt` write, `services.hcl` resolved absolutely, schema read from
  in-memory bytes); `common.Generator` (per-call template parse + `go/format`);
  the six templates (`//go:embed`); and `naming.SnakeCase`.
- **Pluralization is the one `inflection` hazard, and it is solved by
  construction.** `naming.Pluralize` / `PluralizeWithCustomNameSuffix` mutate
  `inflection`'s global rules. `plan.go` computes the plural name once via
  bigdiffer's mutex-guarded copy, and generation consumes it from the plan — it
  never calls `internal/naming.Pluralize`. The plural-DS path builds its template
  data from the CFN type name alone (no schema emission), so it cannot fail on
  schema emission by construction.
- **A fourth aggregate:** `import_examples_gen.json`, emitted from schema bytes.
  Because it reads schemas (for primary-identifier names), it is emitted from the
  promoted overlay + cache — after promotion — not folded into the compiled/gated
  set.
- **`x-derecursed`** is understood and dormant: the copy preserves
  `Emitter.Deduplicate` (via `schemaIsDeRecursed`), but zero cached schemas carry
  the key today, so it needs no special handling.
- **`path_aware_attribute_names`** (an overlay row flag, threaded through
  `plan.pathAwareNames` to `Emitter.PathAwareNames`) resolves a CloudFormation
  schema that has two properties normalizing to the same Terraform attribute
  name at different nesting levels (e.g. `AWS::SageMaker::Cluster`'s
  `FsxLustreConfig` vs. `FSxLustreConfig`) by keying the attribute name map on
  the full CloudFormation parent path instead of the bare name, bypassing
  `emitSchema`'s collision check entirely for that resource. Opt-in per type;
  every other resource keeps the flat map. Ported from the legacy generator
  (issue #3019) with no runtime change — the translation-layer half of that fix
  (`internal/generic/translate.go`/`resource.go`/`data_source.go`) is
  provider-runtime code bigdiffer does not own or copy.

## Deferred and future work

The durable list of remaining work — this section is the single tracker now that
the standalone generation punchlist is retired. Each item is either tracked by a
linked GitHub issue or described in enough detail here to act on.

- **Full-corpus regeneration + `machinery_held` (engine-change safety).** Today
  `-update` regenerates only New/Changed *schema* types, so a change to the
  generation machinery (templates, codegen, naming, or the Go toolchain) never
  reaches schema-unchanged types — their committed output silently goes stale
  (drift `TestFullCorpusParity` can only flag, not prevent). Fix: run the **whole
  corpus** through the existing generate-then-compile gate every `-update`
  (deterministic, ~30s), so machinery changes propagate under the same
  never-regress protection schema changes already get. This needs a new
  per-resource policy category — provisionally **`machinery_held`** — distinct
  from `frozen` (schema pinned) and `suppress` (deliberate omission): it means
  *the schema is current but the current machinery cannot compile this resource,
  so its last-good bytes are kept*. Because the corpus is re-attempted every run,
  the marker is **self-clearing** — it disappears the run after the codegen gap is
  fixed. It keeps all three guarantees: (a) never regress, (b) latest schema,
  (c) latest machinery — where (c) becomes "reflected wherever it compiles, every
  exception recorded per-resource, never silent." Attribution is clean only for
  schema-unchanged failures (the sole new variable is the machinery); changed-
  schema failures stay in the freeze/suppress bucket. This replaces the parity
  test as the engine-change guard (so it also unblocks retiring the parity test's
  dependency on the legacy engine). *The one medium item outstanding.*
- **Move dedup "tags" out of the schema bytes into an `all_schemas.hcl`
  argument.** The deduplication marker currently lives inside the pinned schema
  JSON; make it a first-class overlay argument (a `resourceRow` field + generator
  support) so it is visible and diffable in the overlay. Enabler for the dedup
  thaw below.
- **Thaw resources that now generate cleanly.** The ~63 deferred lift-candidates
  recorded in #3323 (`manual: lift candidate; generates and compiles cleanly …
  pending a batched lift`), plus resources unlocked by the new dedup approach and
  any freezes whose pinned bytes now generate. Confirm with `-heal`; lift in
  batches, each its own reviewed PR (generated-code diffs, never a data-only
  ride-along).
- **Checkout-file retirement (§5).** Fold `suppressions_checkout.txt` fully into
  `frozen_since` and stop reading the external file. Part of legacy removal below.
- **Delete the legacy generators, directive files, and `make` targets** — after
  the engine has driven several clean weekly cycles with no fallback. Tracked by
  #3330 and `removing-the-legacy-generation-process.md`; resolves the cutover from
  conceptual to physical (§10).
