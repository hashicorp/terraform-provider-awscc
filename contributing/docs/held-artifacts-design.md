<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer command redesign: sync / reconcile / check / lint / recheck

> **Transient design doc.** This is the working design for the command-surface
> redesign and the `held` feature it exists to carry, while both are being
> implemented. When they land, the durable parts fold into `bigdiffer-design.md`
> (§0's stories, §6 "The gates," §7's policy table) and this file is deleted.
> It exists so the gap, the five target commands, and what each one is missing
> today are stated cleanly in one place before code is written. See
> `redesign-punchlist.md` for the tracked, checkable-off list of what's left;
> that file links back here for every item's detail rather than restating it.

## 0. The five user stories (north star)

Told plainly first, with no tool vocabulary, because these are what the
command surface has to serve — restated from `bigdiffer-design.md` §0 in the
canonical order and with the command each one lands on:

1. **"It's been a week. Go see what's new and bring the provider up to
   date."** Ask AWS what changed, regenerate what needs it, gate it, ship it,
   update the bookkeeping. → **`sync`**
2. **"I just changed how the code gets written. Redo everything with the new
   way, and don't ask AWS anything — I want today's version of everything,
   generated fresh."** Re-run the machinery against what's already known,
   gate it, and make the new output real. → **`reconcile`**
3. **"Before I do that — will changing how the code gets written break
   anything? Just tell me, don't change anything yet."** The same attempt as
   #2, reported instead of kept — asked by a person once in a while or a
   machine (CI) every single time, no difference between the two askers. →
   **`check`**
4. **"Never mind the code-writing machinery — is everything else in good
   shape right now?"** Is the bookkeeping itself consistent: duplicates,
   contradictions, unexplained gaps, things marked broken with no reason.
   → **`lint`**
5. **"Someone decided a while back that this one thing was broken or
   excluded. Is that still true, or has it quietly become fine?"** Revisit an
   old judgment call, on demand. → **`recheck`**

Naming: `update` → `sync` and `heal` → `recheck` keep their spirit under a new
name; today's `check` is being repurposed to the new story-3 need, and its old
job (story 4) takes the name `lint` instead; `reconcile` is a new name for a
new command. Chosen as a set so no two are close enough in sound or meaning to
blur (the specific failure `bigdiffer-design.md` §0 flagged with the earlier
`check`/`verify` pairing).

| # | Story | Command | AWS? | Writes files/overlay? |
|---|-------|---------|------|------------------------|
| 1 | Weekly bring-current | `sync` | yes | yes |
| 2 | Land a machinery fix, no crawl | `reconcile` | no | yes |
| 3 | Would a machinery change break anything? | `check` | no | no |
| 4 | Is the bookkeeping healthy? | `lint` | no | no (no generation attempted) |
| 5 | Revisit an old decision | `recheck` | no | no (proposes only) |

## 1. Status quo → target, one row per story

The core insight driving this redesign: **`sync`, `reconcile`, and `check` are
one engine, not three.** All three run the identical generate → compile →
decide pipeline over the full corpus; they differ only in whether AWS is
consulted first and whether the result is promoted. Today that pipeline exists
in exactly one place (`runUpdate`, gated, New/Changed-only) and is duplicated,
badly, in a second place (`runGenerate`, ungated, whole-corpus). The redesign's
real work is deleting the second implementation and making the one true
pipeline runnable in all three modes the stories need.

### Story 1 → `sync` (rename only, plus the corpus-wide gate)

**Today:** `-update` (`runUpdate`, `update.go`). Crawls AWS, byte-diffs to find
New/Changed, gates *only those*, promotes, reconciles the overlay.

**Gap:** none in shape — `sync` is `-update` renamed. But it inherits the core
gap this whole redesign exists to close: `buildCandidates` (`update.go:40`)
filters to `statusNew`/`statusChanged` only, so a machinery regression on a
schema-unchanged type is invisible to `sync` today. Closing that (running
every type through the gate, not just New/Changed) is `held`'s core mechanic
(§3 below) and belongs to `sync` as much as to `reconcile`/`check` — it's the
same pipeline extension applied to whichever mode is running.

**To do:** rename the flag/function (`update` → `sync`, `runUpdate` →
`runSync`, doc strings, `make bigdiffer-update` → `make bigdiffer-sync`); wire
the corpus-wide gate + `held` policy into the one shared pipeline it now calls
(§3, §4).

### Story 2 → `reconcile` (new implementation; `-generate`'s replacement)

**Today:** `-generate` (`runGenerate`, `run_generate.go`) regenerates every
type unconditionally, but:

- Never compile-gates. `runGenerate` calls `generateCorpus` then
  `writeCorpus` directly — no `compileFixpoint`, no staging-then-revert, no
  policy decision at all.
- Fails all-or-nothing. Any generation error anywhere aborts the whole run
  (`if genErrs > 0 { return fmt.Errorf(...) }`) — nothing is written, not even
  the types that generated fine.
- Never touches the overlay. `all_schemas.hcl` is read for input rows only;
  no policy is ever applied back to it, so it cannot record a hold, a freeze,
  or anything else.

This is precisely the gate-less second implementation §1 above calls out —
`runGenerate` was never wired into the gate/policy pipeline `runUpdate` uses.
It cannot back story 2 as it stands; it can only "regenerate or abort," which
is not what "land a machinery fix" means once holding broken artifacts
(rather than aborting on the first one) is the requirement.

**Gap:** `runGenerate` needs to be deleted, not extended. `reconcile`'s
implementation is the *same* pipeline `sync` uses — candidate-build,
generate+compile fixpoint, `decide()`, promote, reconcile the overlay — with
one input swapped: rows and schema bytes come from the committed overlay +
cache (no `discover()` call), and **every** row is a candidate, not just
New/Changed (there is no "changed" concept without a live AWS comparison —
offline, every type is "attempt it," full stop).

**To do:**

- Delete `run_generate.go`'s `runGenerate` and its all-or-nothing generation
  loop entirely.
- Add an offline row source: build the candidate list from `loadOverlay`'s
  rows directly (every row, `class` derived from the row's own state —
  `classPresent` for all of them, since there is no New/Absent without a live
  AWS set) instead of from `detectChanges`/`buildCandidates`.
- Reuse `refreshCandidate` → `compileFixpoint` → `decide()` → `promoteStaged`
  → `normalizeWithDecisions` unchanged — the exact sequence `runUpdate`
  already runs, just fed offline candidates instead of AWS-diffed ones.
- New per-artifact byte-diff (regenerated vs. committed) feeding `codegen_error`
  / `toolchain_error` classification (§3, §4) — this is net-new; nothing today
  compares regenerated bytes against the committed file for this purpose.

### Story 3 → `check` (repurposed name; wholly new functionality)

**Today:** does not exist. No command runs the gate and throws the result
away — every existing command either doesn't gate (`-generate`) or gates and
commits (`-update`).

**Gap:** total — this is genuinely new, not a rename. It is `reconcile`'s
identical pipeline with promotion switched off: run candidate-build →
generate+compile fixpoint → `decide()`, then stop — print the report (which
`held`/`codegen_error`/`toolchain_error` decisions *would* be made), never
call `promoteStaged`, never write `all_schemas.hcl`. Nothing on disk changes.

**To do:**

- A `-dry-run`-style switch through the same offline pipeline `reconcile`
  uses, short-circuiting before `promoteStaged`/the overlay write.
- Exit non-zero if anything would be held (or the volume threshold, §4, would
  trip) — this is the command a CI gate runs on every PR that touches
  `codegen/`, templates, naming, or `go.mod`/`go.sum`, so its exit code is the
  actual gate signal, not just informational output for a human.
- This is the natural replacement for `TestFullCorpusParity`'s role as the
  engine-change guard once the legacy comparison is retired (`bigdiffer-design.md`
  §10) — self-referential (current engine vs. committed output) rather than
  legacy-referential.

### Story 4 → `lint` (rename of today's `-check`, plus a new anomaly)

**Today:** `-check` (`runCheck`, `main.go`). Purely structural: normalizes the
overlay against itself, checks sort order/formatting, flags duplicate blocks,
naming-invariant violations, and reason-less suppressed/frozen facts
(`anomalyProblems()`). Never attempts generation.

**Gap:** small and additive. `lint` is `-check` renamed, plus one new anomaly
class: a `held_*` marker present on any row. A hold that survives past the PR
that should have fixed it is exactly the kind of thing `anomalyProblems()`
already exists to catch (parallel to today's reason-less-suppression check),
so this is one more case in that same function, not a new mechanism.

**To do:**

- Rename `-check` → `lint`, `runCheck` → `runLint`.
- Add a `held` case to `Report`/`anomalyProblems()`: any row with
  `held_resource`/`held_singular_data_source`/`held_plural_data_source` set is
  reported and fails `lint`.

### Story 5 → `recheck` (rename of `-heal`, plus a scope expansion)

**Today:** `-heal` (`runHeal`, `heal.go`). `healFact.needsHealing()`
(`heal.go:60`) re-probes an active suppress/freeze fact only when its own
reason is empty or tagged `unknown` — it never revisits a fact that already
carries a real, categorized reason.

**Gap:** `recheck` is `-heal` renamed, plus an opt-in mode that drops the
"reason is empty/unknown" half of `needsHealing`'s condition — re-probe every
active fact regardless of whether it already has a reason, so a year-old
`manual` or `structural` call can be asked "is this still true?" on demand,
not just backlog with no reason yet.

**To do:**

- Rename `-heal` → `recheck`, `runHeal` → `runRecheck`.
- Add a flag (name TBD — `-recheck-all`? `-force`?) that widens
  `needsHealing`'s predicate to `active` alone when set, leaving the default
  (reason-less/`unknown` only) as today's behavior for a plain `recheck` run.
- Whether `recheck` should also re-probe `held` rows is a real question, but
  the answer is no by construction: `held` is self-clearing under `sync`/
  `reconcile`/`check` already (§4's self-clearing requirement) — it doesn't
  need `-heal`'s propose-and-report model, because it isn't a backlog item
  waiting on a human's classification, it's a mechanical state re-derived
  every run. `recheck` stays scoped to `suppress`/`frozen`, as today.

## 2. What's shared vs. what's genuinely new

Restating §1's per-story gaps as one engineering fact: there is exactly **one**
generate→compile→decide pipeline to build correctly, and three thin callers
around it (`sync`, `reconcile`, `check`), differing only in:

- **Row source**: `sync` discovers from AWS and diffs against the cache;
  `reconcile`/`check` read the committed overlay + cache directly, every row a
  candidate.
- **What happens after the fixpoint settles**: `sync`/`reconcile` promote and
  write the overlay; `check` reports and stops.

`lint` and `recheck` never enter this pipeline at all — they are the two
read-only paths from `bigdiffer-design.md` §0, and their gaps (one new anomaly
case; one widened predicate) are independent, small, additive changes to code
that already exists and already works.

## 3. Failure modes (`codegen_error` / `toolchain_error`)

A schema-unchanged failure is not automatically a codegen defect. bigdiffer
already holds both the freshly regenerated bytes (in memory, before promotion)
and the committed file (on disk), so a cheap byte-comparison — no second
compile — separates two distinct causes:

| Mode | Regenerated bytes vs. committed | Diagnosis |
|------|--------------------------------|-----------|
| `codegen_error` | **differ**, and the new output is what fails | This run's codegen produced something different from what is shipping, and the different output broke. The committed file — unchanged, still in production — is known-good. |
| `toolchain_error` | **identical**, and that same file now fails | Codegen produced the exact bytes it always has; the environment around them (Go version, a dependency bump, a shared runtime package) changed. The committed file is not "known-good" in the same sense — it was good under the *previous* environment. |

**The nuance: cause is not disposition.** The byte-diff labels the cause; the
actual hold-vs-block decision keys on the **post-revert build result** (does
the committed file, alone, still build under the current environment right
now), which `compileFixpoint`'s existing revert-and-rebuild already computes:

- `codegen_error` → revert goes green → hold and ship.
- `toolchain_error` → revert stays red → block the run; nothing promoted,
  nothing persisted to the overlay.

This also correctly handles the case a byte-diff alone would mislabel:
regenerated bytes differ (looks like `codegen_error`), yet the committed file
*also* fails post-revert — that's still a block, not a hold, because the
committed tree does not build.

Two more points of nuance:

- **Attribution reuses the existing blame-mapping** (`compileFixpoint`'s
  `file:line:col` → staged-artifact mapping, `bigdiffer-design.md` §6). A
  dependency bump breaking *some* resources and not others holds each
  affected artifact individually; a blamed file mapping to nothing staged is
  unattributable exactly as §6 defines it today and hard-errors the run.
- **Per-artifact, matching `suppress_*`**: resource / singular data source /
  plural data source each hold independently; list resource is not a fourth
  bucket (it rides on the resource artifact, per §6's list-resource coupling
  guard).

## 4. How `sync`/`reconcile`/`check` should behave on a hold

### 4.1 `codegen_error`: hold, mark, proceed — loudly

- `sync`/`reconcile`: keep the committed file (never regress); set
  `held_<artifact>` + `held_reason_<artifact> = "codegen_error: <detail>"` on
  the overlay row; let the run complete — the committed tree builds — but
  surface the hold prominently (§4.4). `check` computes and reports the
  identical decision, promotes nothing.

### 4.2 `toolchain_error`: block the run

- All three: the committed tree will not build under the current environment,
  so the run cannot produce a shippable/reportable-as-fine result. Hard-error,
  with a per-artifact report of every `toolchain_error` and its blame. Nothing
  is promoted; nothing is persisted — a `toolchain_error` never becomes a
  committed overlay marker, it is a run-blocking condition, not a durable state.

### 4.3 Volume threshold (a `codegen_error` safety valve)

`toolchain_error` already blocks regardless of count. For `codegen_error`:
each is individually holdable, but holding hundreds in one run is a systemic
regression that "completing successfully" would bury. If the held count
crosses a threshold (illustrative, not committed: ~10–20% of attempted types,
or an absolute count in the low hundreds), hard-error instead of completing —
same instinct as an unattributable blame, triggered by volume instead.

### 4.4 Noisy reporting

An engine regression must be impossible to miss, across all three commands:

1. **End-of-run summary**, distinct from per-type progress: *N artifacts held
   this run*, grouped by category, each with type/artifact/blame. A
   `toolchain_error` spike reads as one event, not scattered rows. If the run
   blocked or tripped the threshold, that headline comes first.
2. **The committed diff** (`sync`/`reconcile` only): `codegen_error` markers
   land in `all_schemas.hcl`, so they appear in the PR diff.
3. **`lint` fails on them**: any `held_*` marker is an anomaly (§1, story 4);
   a hold surviving past the PR that should have fixed it fails CI.

### 4.5 Self-clearing

Because `sync`/`reconcile` re-attempt every row every run, a held artifact is
re-evaluated every time one of them runs against it. The run after the
regression is fixed, that artifact regenerates/compiles cleanly, is promoted,
and its marker clears automatically — no `recheck` step needed for this
category (§1, story 5's "why not held too" note).

## 5. Overlay markers (near-term shape)

`held` follows today's two-attribute-per-fact convention (bool + separate
`_reason` string), matching `suppress_*_generation`/`suppression_reason_*` and
`frozen_since`/`frozen_reason`, so it ships consistent with the levers already
in the overlay rather than inventing a third shape for maintainers to learn:

- `held_resource` / `held_singular_data_source` / `held_plural_data_source`
  (bool).
- `held_reason_resource` / `_singular_data_source` / `_plural_data_source`
  (string, `category: detail`; category is `codegen_error` only —
  `toolchain_error` blocks and is never persisted, §4.2).

Orthogonal to both existing levers:

- **Not `frozen_since`** — freezing pins the *schema*; a held type's schema is
  fine and keeps refreshing normally under `sync`. A held artifact never sets
  `frozen_since`.
- **Not `suppress_*`** — suppression is a standing, human-owned decision to
  not generate; `held` is the tool temporarily unable to keep a promise it
  already kept, intent unchanged.

> Collapsing `suppress`/`held`/`frozen` into one self-describing attribute per
> fact is **out of scope here** — tracked separately in `bigdiffer-design.md`
> "Deferred and future work." `held`'s shape is chosen to be additive toward
> that direction, not to prejudge it.

## 6. Implementation plan

Ordered so each step is independently mergeable and leaves the tool working:

1. **Rename pass first, no behavior change.** `update`→`sync`, `heal`→`recheck`,
   today's `check`→`lint` (flags, functions, doc strings, Makefile targets,
   README, runbook). Zero risk, unblocks everything else being named correctly
   from the start. Delete `run_generate.go`'s `runGenerate` in the same pass —
   its behavior is being replaced, not carried forward, so there is nothing to
   preserve by keeping it alive under a new name.
2. **Extract the shared offline pipeline.** A function that takes "every
   overlay row as a candidate" and runs candidate-build → `refreshCandidate` →
   `compileFixpoint` → `decide()`, returning the settled decisions —
   factored out of `runUpdate`'s AWS-specific version so `reconcile` and
   `check` can both call it with an offline row source.
3. **`reconcile`**: wire the shared pipeline to promotion + overlay write,
   offline row source, no AWS.
4. **`check`**: wire the same shared pipeline, stop before promotion, report,
   set the exit code.
5. **`held` policy + byte-diff classification**: the `codegen_error`/
   `toolchain_error` split (§3), the post-revert-build discriminator (§4), the
   volume threshold (§4.3) — implemented once, in the shared pipeline, so
   `sync`, `reconcile`, and `check` all get it simultaneously rather than
   risking the two-implementation drift that created this gap originally.
6. **`sync`'s corpus-wide gate**: extend `buildCandidates` (or its replacement)
   so `statusUnchanged` rows also become candidates for the gate (not just
   New/Changed) — the one behavior change to `sync` beyond its rename.
7. **`lint`'s new anomaly** + **`recheck`'s reasoned-row flag**: small,
   independent, no dependency on 2–6.
8. **Reporting** (§4.4) — the end-of-run held summary and its presence in
   the PR diff. Code/behavior only; no doc changes here (see step 9).
9. **Reconcile every other doc in one pass, last.** Decided explicitly:
   this file and `redesign-punchlist.md` are the only docs touched while
   steps 1–8 are in progress. Once they're done, update
   `bigdiffer-design.md` (drop its out-of-sync notice; fold in §0's command
   table, §6, §7), the README, the runbook, and `suppressed-and-frozen.md`
   (the two new reason categories) together in one pass — not incrementally
   alongside steps 1–8 — then delete this file and `redesign-punchlist.md`.

## 7. Open questions

- Exact flag name for `recheck`'s reasoned-row mode.
- Exit code of `reconcile`/`sync` when they complete with `codegen_error`
  holds: exit 0 with the loud banner (the tree is shippable, `lint`/`check`
  are the hard gates), or non-zero to force attention? Leaning exit 0 + banner,
  matching the prior draft's reasoning.
- Exact threshold value for §4.3, and whether it's a fraction, an absolute
  count, or both.
- Whether `check` is wired into CI as a new Go test, a `make` target, or a
  dedicated CI job/workflow step — and, either way, how CI scopes it to only
  the PRs that plausibly touch the engine (`codegen/`, templates, naming,
  `go.mod`/`go.sum`) rather than running the full offline sweep on every PR
  regardless of what it touches.
