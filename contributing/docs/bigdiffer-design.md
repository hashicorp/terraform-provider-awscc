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

## 0. User stories: the paths through bigdiffer

Six needs drive every command bigdiffer has (or should have). Stated first,
before the model and the gates, because the paths a maintainer actually walks
should shape the design, not fall out of it after the fact — this section
should have existed from the start.

1. **Weekly release prep.** Crawl AWS, refresh what changed, gate it, promote,
   open a PR. (Today's `-update`.)
2. **"Did my codegen/toolchain change just break something?"** — a human,
   locally, after touching a template, `codegen/`, naming, or bumping a
   dependency/Go version. No AWS involved; nothing should be promoted just to
   answer the question. Today's gap: nothing runs this check at all (§6,
   "Gap: the unchanged corpus is never gated") — a dependency bump today gets
   effectively no verification beyond whatever `go build ./...` already
   catches for code that happens to be touched.
3. **The same check, in CI, on every PR that could have broken generation.**
   Identical operation to story 2 — different caller (CI, not a human),
   gating a merge instead of informing a person before one.
4. **Actually land a machinery fix without an AWS crawl.** Same gated pass as
   2/3, but this time promote: write the regenerated files, clear any `held`
   markers that now resolve, commit.
5. **"Is the overlay itself healthy right now?"** Including whether anything
   is currently `held`. Offline, structural, no generation attempted — today's
   `-check`, extended to treat a `held` marker as an anomaly the same way a
   reason-less suppression already is.
6. **Revisit an old suppression, freeze, or hold and see if it still holds
   up.** Today's `-heal` only proposes reasons for reason-less/`unknown` rows;
   this story is different — a human going back to a year-old, *already
   reasoned* decision and asking bigdiffer to re-attempt it, on demand, to see
   if the underlying problem is still real (queued: "Deferred and future
   work").

Stories 2 and 3 are the identical operation; so are 1 and 4, up to whether AWS
is consulted. That collapses six *needs* into three *operations* —
crawl-and-promote, gate-and-report, gate-and-promote — each runnable with or
without an AWS crawl, plus the two read-only paths (5, 6) that never generate
anything. The exact command surface for these three operations (today's
`-update` and `-generate`, plus a currently-missing gate-and-report path) is
not yet settled — see "Deferred and future work." `held` is the one state distinct from `suppress`/`frozen` in exactly
this respect: `suppress_*`/`frozen_since` are inputs the gate honors (skip this,
don't refresh that); `held` is only ever an *output* of an attempt, so it must
be re-attempted, not just skipped, every time one of these paths runs against
that artifact — and cleared the moment the attempt succeeds. A `held` marker
that outlives its cause is a bug in the removal step, not an acceptable steady
state, the same way a stale `frozen_since` on a type AWS un-deprecated would be
(§7's self-healing row already handles that case for `frozen`; `held` needs
the equivalent guarantee, actively enforced, not just implied by "self-clearing").

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

### Gap: the unchanged corpus is never gated (skeleton design, not yet implemented)

Both gates above only run on New/Changed types (§4) — a type whose schema
bytes are byte-identical to the cache is skipped entirely (`statusUnchanged`,
`change.go`). That is correct for schema drift, but it means a change to the
*machinery itself* — a template, `codegen/`, the naming package, or the Go
toolchain — is never exercised against the ~1580 types whose schema didn't
move this week. Today the only thing that would ever catch such a regression
is `TestFullCorpusParity`, and only for as long as that test's legacy-engine
comparison exists (it is itself slated for removal once the legacy engine is
deleted, per §10). Once that guard is gone, a machinery regression on an
unchanged-schema type would go undetected until something downstream actually
breaks — silent, not surfaced.

**Design.** Run every type through generate-then-compile every `-update`, not
just New/Changed — the gate already exists and is deterministic and fast
(§6's compile gate alone measures in seconds; running generation for the full
corpus is the same order of magnitude). A schema-unchanged type that now
fails either gate needs a result distinct from both existing outcomes:

- **Not `frozen_since`.** Freezing means "the *schema* is pinned" (§7,
  `suppressed-and-frozen.md`, "Frozen, precisely") — it also stops future
  schema refresh for that type. A machinery regression has nothing wrong with
  the schema; freezing it would be both the wrong claim and the wrong
  behavior (the type should keep refreshing normally the moment the schema
  actually changes, independent of whether the machinery currently compiles
  it).
- **Not `suppress_*_generation`.** Suppression means "don't generate this
  artifact" as a standing, reason-carrying decision about the artifact. A
  machinery regression is the tool failing to keep a promise it already made
  last week (the artifact generated and compiled fine then) — the intent is
  still "generate this," the tool just currently can't.

**Two causes, one outcome.** A schema-unchanged failure is not automatically
a codegen defect. Compare the freshly regenerated bytes against the committed
file — bigdiffer already has both, no second compile pass needed — and two
distinct causes emerge, but both resolve to the *same* thing happening to the
artifact: it keeps its last-good committed file, unpromoted, with a reason
attached.

| Cause | Generated bytes vs. committed | What it means |
|---|---|---|
| `codegen_error` | **Changed**, and the new output fails | This run's codegen/plan output is different from what's shipping, and the different output is what broke. The committed file — unmodified, still what's running in production — is known-good. |
| `toolchain_error` | **Unchanged** (byte-identical to committed), and that same file now fails | Codegen produced the exact same file it always has; the environment around it (Go version, a dependency bump, a shared runtime package) is what changed. The committed file is not "known-good" in the same sense as `codegen_error` — it was good under the *previous* environment, and might fail too if it were ever rebuilt from scratch — but it is what's *already shipping*, so leaving it in place changes nothing about the *already-released* provider — though the tree this run would commit still won't build (see disposition below). |

Both are **held** for detection and reporting — the committed file is not
replaced, and the overlay records why — under one mechanism with two reason
categories rather than two different mechanisms. But they part ways on *release
disposition*, and the byte-diff is the wrong thing to decide that on. The
invariant that must hold is "the committed tree builds," and §6's fixpoint
already computes exactly that when it reverts a failed artifact to its
committed file and rebuilds:

- `codegen_error`: reverting to the committed file goes **green** — that file
  is genuinely known-good — so the artifact is held, its marker persisted, and
  the release proceeds.
- `toolchain_error`: the committed file *is* the file that failed (the
  regenerated bytes were byte-identical), so reverting to it stays **red** —
  the tree this run would commit does not build under the current environment.
  That cannot ship, so a `toolchain_error` **blocks the run** (hard-error,
  reported per-artifact) rather than completing with a held marker. "Nothing
  changes about the *already-released* provider" is true; "the release proceeds"
  is not.

So the byte-diff still labels the cause for the report, but the hold-vs-block
decision keys on the fixpoint's post-revert build result — which also correctly
handles the case the byte-diff would mislabel: regenerated bytes differ, yet the
committed file *also* fails under a bumped toolchain (labeled `codegen_error`,
but post-revert red, so blocked). Only `codegen_error` ever reaches the overlay
as a persisted `held` marker; `toolchain_error` is a transient, run-blocking
report artifact. This also sidesteps needing to treat either as a suppression:
nothing is removed from the provider, so there is no "backlog" framing to force
onto it the way an actual suppression implies.

**Proposed: a new `held` per-artifact status, orthogonal to `frozen_since` and
`suppress_*`.**

- **Three-way per-artifact split, matching `suppress_*_generation` exactly.**
  `held_resource` / `held_singular_data_source` / `held_plural_data_source`
  (bool) plus `held_reason_resource` / `_singular_data_source` /
  `_plural_data_source` (string, `category: detail` shape, category one of
  `codegen_error` / `toolchain_error` — a proposed addition to
  `suppressed-and-frozen.md`'s existing five-category taxonomy, §7). List
  resource is not a fourth bucket, for the same reason it isn't one for
  `suppress_*_generation` today: it isn't an independent artifact, it's an
  attribute of the resource artifact (`-listresource`, gated by whether the
  plural data source is suppressed — §6's "list-resource coupling guard").
  A list-resource-specific compile failure is the *resource* artifact's own
  `held`, with the reason text naming the list resource specifically, not a
  separate `held_list_resource`.
- **Self-clearing.** Because the full corpus is re-attempted every `-update`
  (not just on a schema change), a held artifact is re-evaluated every run:
  the run after the regression is fixed (in codegen, or by the environment
  moving again — a further dependency bump, a template fix, etc.), that
  artifact generates/compiles cleanly, is promoted, and the marker clears
  automatically — no `-heal` step needed for this category specifically,
  unlike `unknown`/`manual` rows.
- **Distinguishable from a real suppression in the report and the overlay.**
  A held row is not backlog (nobody needs to open a "fix this resource's
  schema" issue — the resource's schema is fine) and not something a human
  approves; it is bigdiffer telling its own maintainers "the last cycle's
  output changed or the environment changed, either way this needs a code
  fix, not an overlay edit." The change report gets a distinct section for
  it (parallel to the existing frozen/suppressed section), grouped by reason
  category so a `toolchain_error` spike after a dependency bump is
  immediately visible as one event, not scattered among unrelated
  `codegen_error` rows — and `-check` flags any held row as an anomaly
  (parallel to today's `anomalyProblems()`) rather than passing silently, so
  a held marker surviving past the PR that should have fixed it is itself a
  signal the fix didn't land.
- **Only schema-unchanged failures are eligible.** A schema-changed type that
  fails is already covered by the existing freeze/suppress branches in §7's
  policy table — the schema itself is the newest variable there, so
  `generation_failed`/`build_failed` under a freeze or suppression remains
  the right call, not `held`.

**Attribution, and when to give up on attribution.** A build failure is
attributable when `go build`'s blame maps cleanly to a staged artifact/package
(exactly the mapping §6's fixpoint loop already does) — that covers both
causes above, including a dependency bump that breaks some resources and not
others: each affected resource is held individually, with its own
`toolchain_error` reason, and the rest of the corpus proceeds unaffected. A
blamed file that maps to nothing staged is unattributable exactly as §6
already defines it (a pre-existing base-tree break, or a bug in bigdiffer's
own emitted registration file) — that case is not "held," it hard-errors the
run today and should continue to.

**A wide `codegen_error` batch should still stop the run.** The systemic
dependency-bump case is already caught — its committed files stay red on
revert, so `toolchain_error` blocks regardless of count. What a volume
threshold still guards is a large `codegen_error` batch: each is individually
holdable and shippable (post-revert green), but a run that holds hundreds of
them is a systemic codegen regression, and completing "successfully" while
silently skipping a large fraction of the corpus would bury the signal in a
huge, technically-accurate report. Proposed: track the held count as a fraction
of the gated corpus during the full-corpus sweep, and above a threshold
(illustrative, not committed: something like 10–20% of attempted types, or a
held count over N in the low hundreds), hard-error instead of completing — the
same "abort on aggregate" instinct §6 already has, triggered by volume instead
of by an unmappable blame. Below the threshold, held rows are exactly the
clear, per-artifact signal the report already makes them, and the release
proceeds.

**Noted, not designed here: unifying the taxonomy.** `held_*` above still
follows today's two-attribute-per-fact convention (a bool plus a separate
`_reason` string), matching `suppress_*_generation`/`suppression_reason_*`
and `frozen_since`/`frozen_reason` as they exist now, so `held` ships
consistent with the levers already in the overlay. Collapsing all three into
one self-describing attribute per fact (dropping the boolean *and* the
`_reason` suffix — the value already reads as an explanation on its own,
e.g. `resource_suppressed = "structural: <detail>"` conveys everything
`suppress_resource_generation = true` + `suppression_reason_resource = "..."`
does today) is a separate, larger design problem, called out as its own item
in "Deferred and future work" rather than folded into `held` here — `held`'s
shape is chosen to be additive toward a possible future unification, not another
format to migrate away from later, but the unification itself needs its own
design pass before scheduling.

This keeps all three of bigdiffer's guarantees simultaneously once
implemented: never regress a shipped artifact, always reflect the latest
schema, and always reflect the latest machinery wherever it can — with every
place it can't recorded per-artifact, never silent. It also removes
`TestFullCorpusParity`'s role as the only engine-change guard, unblocking that
test's own retirement (§10, "Deferred and future work").

Design only — not yet implemented. Tracked as the next `held`/machinery-gap
work item in "Deferred and future work."

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
| Unchanged (schema) | failed, committed file still builds (`codegen_error`) | hold the failed artifact(s): keep last-good bytes, persist `held_*` marker, ship; never `frozen_since` *(design, §6, not yet implemented)* |
| Unchanged (schema) | failed, committed file also fails (`toolchain_error`) | block the run — the committed tree won't build; report per-artifact, persist nothing *(design, §6, not yet implemented)* |

Two invariants keep every row safe by default:

- Never regress a shipped resource — a broken `Present` type is frozen at its
  last-good bytes, not failed.
- Never let a broken new type block the release — a broken `New` type is added
  suppressed and becomes backlog.

The reason taxonomy (`structural` / `generation_failed` / `build_failed` /
`manual` / `unknown`), the `-check` reason anomaly, and the `-heal` re-probe are
specified in `suppressed-and-frozen.md`. The last table row is self-healing:
policy is recomputed every cycle, so a type AWS has since fixed is proposed for
recovery automatically. Two more categories, `codegen_error` and
`toolchain_error`, are proposed in §6 ("Gap: the unchanged corpus is never
gated") but not yet implemented.

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

- **`-heal`: add an opt-in mode to re-probe rows that already have a reason,**
  not just reason-less/`unknown` ones. Real story: revisit a year-old
  suppression/freeze to see if it's still deserved, not just fill gaps.
  Quick note only — not designed.
- **`held`: gate the unchanged corpus too (engine-change safety).** Full
  skeleton design in §6 ("Gap: the unchanged corpus is never gated") and §7's
  policy table. Implementation: run generate-then-compile over the whole
  corpus every `-update`, not just New/Changed; add the `held_*` per-artifact
  attributes and the `codegen_error` / `toolchain_error` reason categories
  (`policy.go`, `suppressed-and-frozen.md`); surface held rows as a distinct,
  category-grouped report section and a `-check` anomaly; add the
  volume-based abort threshold (§6) so a wide, attributable regression still
  stops the run instead of completing as a large held batch. Replaces
  `TestFullCorpusParity` as the engine-change guard, unblocking that test's
  own retirement. *The one medium item outstanding.*
- **Settle the command surface around one unified gate/policy pipeline.**
  §0 collapses six user stories into three operations (crawl-and-promote,
  gate-and-report, gate-and-promote), each with/without an AWS crawl. Today's
  `-generate` is a second, gate-less code path (`run_generate.go`) — it
  regenerates the whole corpus but never compile-gates and never applies
  policy, so it cannot produce a `held` marker and cannot back stories 2/3/4.
  That gate-less shortcut needs to be eliminated, not preserved alongside the
  gated pipeline `-update` uses — same underlying engine, parameterized by
  whether rows/schema bytes come from an AWS crawl or the committed base, and
  by whether the result is promoted or only reported. Exact flag/command
  names not decided.
- **Unify the overlay's exclusion levers — candidate shape, needs its own
  design pass.** Not scheduled, not decided; documented so the next design step
  has a starting point, not a conclusion. Today every exclusion is a
  boolean-plus-separate-reason-string pair — `suppress_resource_generation` +
  `suppression_reason_resource`, `frozen_since` + `frozen_reason`, and (once
  implemented) `held_resource` + `held_reason_resource` — three levers,
  six-plus attributes per artifact, with a flag/reason-mismatch class `-check`
  has to guard against.

  The realization that shapes this: these facts live on **two different
  granularities**, so "unify" cannot mean "one attribute per row."

    - **Output disposition — per-artifact (the trio).** `suppress` and `held`
      are both "what happened to this artifact's output," and each of the three
      artifacts (resource / singular data source / plural data source) carries it
      independently. *Candidate:* one self-describing attribute per artifact —
      `resource = "held: codegen_error: <detail>"`,
      `singular_data_source = "suppressed: manual: <detail> (issue: URL)"`; absent
      means "generate normally." The `state:` prefix (`suppressed` vs `held`) is
      how consumers route — `held:` is always a `-check` anomaly; a `suppressed:`
      with no `category: detail` is the reason-less anomaly — collapsing two
      levers into one and removing the flag/reason mismatch.
    - **Schema input state — per-type.** `frozen_since` pins the *shared* schema
      JSON (one file backs all three artifacts — the "cannot partially advance"
      invariant), so it is inherently whole-type and does not belong in the trio;
      forcing it in would triplicate it and break that invariant. It also carries
      a *date*, which resists the "one self-describing string" collapse.

  This is one candidate, not a plan. **Open questions, and reasons it may land
  differently:** whether `frozen_since`'s date folds into a value or stays a
  distinct field; whether the trio is three flat attributes or a nested
  `artifacts {}` block; how the intra-trio coupling (a suppressed/held
  `plural_data_source` forces the `resource` to drop its list-resource
  capability — §6's guard) is represented, if at all, versus left as a
  generation-time rule; and a migration path for ~1580+ existing rows and every
  consumer (`resourceRow`, `-check`, `-heal`, the HCL schema). It is a breaking
  change to the overlay format, well beyond a point release, and worth weighing
  against simpler alternatives — e.g. keeping today's pairs but adding a single
  `-check` invariant for the mismatch class — before committing to a rewrite.
  `held` (§6) is being designed to sit *additively* toward this direction, not
  to foreclose it; because it has no existing rows it could serve as a low-risk
  pilot for the disposition shape, but that too is a call for the dedicated
  design pass, not this work.
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
