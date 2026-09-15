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

`reconcile`'s naming was pressure-tested against `bigdiffer-design.md`'s
existing durable use of "reconcile" (§12's thesis: "reconcile the live AWS
registry directly against `all_schemas.hcl`," describing the AWS-facing
operation). The concern: story 2's command is the *no-AWS* one, seemingly the
opposite of that meaning. Decision: keep `reconcile` as the command name.
"Reconcile the overlay" (the prose sense) survives fine alongside it — the
command reconciles the committed overlay against fresh generation output, no
AWS required for that particular reconciliation; the word describes the same
kind of operation (bring two things back into agreement) in both places, not
two conflicting ones.

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

**Pinned, from review: `sync` must promote valid drift on `statusUnchanged`
rows, not just gate them.** Once `sync` gates the whole corpus (not only
New/Changed), a schema-unchanged type can regenerate to output that is
**different from committed but compiles fine** — no hold, no failure, just a
machinery-driven change with nothing wrong with it. If `sync` only gated
that case (checked it, then discarded the new bytes because the row was
`statusUnchanged`), the drift would never self-heal via the weekly cycle
either — it would sit there, exactly the "committed output silently ages"
gap this feature exists to close, just relocated from "sync doesn't check
it" to "sync checks it but throws the good result away." So: `sync` promotes
*any* output-diff (§3) that passes the gate, `statusUnchanged` or not — the
byte-diff that matters for promotion is the output diff (regenerated vs.
committed), not the discovery diff (which only decides freeze-vs-held
routing, §3). The practical consequence: an engine change can churn a real
diff across a large slice of the corpus at the very next `sync`, which is the
correct, intended behavior, not a bug — better an expected, reviewable diff
in the `sync` PR than corpus-wide silent drift.

**To do:** rename the flag/function (`update` → `sync`, `runUpdate` →
`runSync`, doc strings, `make bigdiffer-update` → `make bigdiffer-sync`); wire
the corpus-wide gate + `held` policy into the one shared pipeline it now calls
(§3, §4); promote any output-diff that passes the gate regardless of
discovery-diff status (the pinned decision above).

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

**Completeness gap, found in review: failing on holds only doesn't fully
close "silent staleness."** The gap this whole redesign exists to close
(§1's framing) is *committed output silently ages* — and that includes a
case a holds-only `check` cannot see. An engine PR has two distinct failure
modes: (a) regeneration breaks — a hold, which `check` catches; (b)
regeneration produces **different-but-valid** output that compiles fine but
that the author never reconciled and committed — `check` would report clean
(nothing failed to generate or build), yet the committed tree is now
observably stale relative to what the current machinery actually produces.
As specified with a holds-only exit code, (b) slips through `check` silently
— exactly the failure mode this feature exists to prevent, just for "didn't
break, but isn't what's committed" instead of "broke." **Fix: `check` exits
non-zero on any output-diff mismatch (§3's second diff), not only on a
hold-worthy one** — regenerated bytes differing from committed bytes at all,
whether or not the regenerated version compiles. A non-engine PR touches
nothing the engine reads, so its regenerated output is byte-identical to
committed and `check` passes with zero diff; an engine PR must run
`reconcile` and commit the result until `check` reports clean. This closes
the whole gap in one gate instead of half of it, and costs nothing extra —
the byte-diff already has to be computed for `codegen_error`/`toolchain_error`
classification (§3).

**To do:**

- A `-dry-run`-style switch through the same offline pipeline `reconcile`
  uses, short-circuiting before `promoteStaged`/the overlay write.
- Exit non-zero if anything would be held, if the volume threshold (§4)
  would trip, **or if any artifact's output diff is non-empty even without
  a hold** (the completeness fix above) — this is the command a CI gate runs
  on every PR that touches `codegen/`, templates, naming, or `go.mod`/`go.sum`,
  so its exit code is the actual gate signal, not just informational output
  for a human.
- This is the natural replacement for `TestFullCorpusParity`'s role as the
  engine-change guard once the legacy comparison is retired (`bigdiffer-design.md`
  §10) — self-referential (current engine vs. committed output) rather than
  legacy-referential.

### Story 4 → `lint` (rename of today's `-check`, plus a new anomaly)

**Today:** `-check` (`runCheck`, `main.go`). Purely structural: normalizes the
overlay against itself, checks sort order/formatting, flags duplicate blocks,
naming-invariant violations, and reason-less suppressed/frozen facts
(`anomalyProblems()`). Never attempts generation.

**Gap:** small and additive, but not quite as simple as originally stated —
see the resolved §4.1/§4.4 contradiction. `lint` is `-check` renamed, plus a
new anomaly class: a **stale** `held_*` marker (older than a grace window,
via a new `held_since_*` date, mirroring `frozen_since`) is reported and
fails `lint`; a `held_*` marker that just landed this cycle does not — a
`codegen_error` is designed to ship without blocking (§4.1), so `lint` failing
on it the moment it's created would contradict that guarantee. A stale hold
is exactly the kind of thing `anomalyProblems()` already exists to catch
(parallel to today's reason-less-suppression check), so this is one more case
in that same function, not a new mechanism.

**To do:**

- Rename `-check` → `lint`, `runCheck` → `runLint`.
- Add a `held` case to `Report`/`anomalyProblems()`: any row with
  `held_resource`/`held_singular_data_source`/`held_plural_data_source` set
  **and** its paired `held_since_*` older than the grace window (§4.1, §4.3's
  sibling — exact window TBD, §7) is reported and fails `lint`. A fresh hold
  is not an anomaly by itself.

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

**Two different byte-diffs answer two different questions, and only one of
them currently reaches `decide()`.** This section (and the table below) is
about the *second* one; the *first* one is missing from the design entirely,
and its absence is a real bug, not a documentation gap:

1. **Discovery diff** — freshly discovered AWS bytes vs. the committed schema
   cache (`classifyChange`, `change.go:19-21`: `statusChanged` vs.
   `statusUnchanged`). This is the signal that must decide **whether `held`
   applies at all**, as opposed to the existing freeze/suppress path (§7's
   policy table in `bigdiffer-design.md`). Only `sync` can compute it — it
   requires a live AWS crawl. `check`/`reconcile` have no schema-change signal
   at all, because they never call `discover()`.
2. **Output diff** — freshly regenerated bytes vs. the committed generated
   file. This is what the table below actually classifies (`codegen_error` vs.
   `toolchain_error`). It exists, and is computable, in all three commands.

**The bug this creates today:** `decide(class, gr, today)` (`policy.go:78`)
routes purely on `changeClass` — `classNew` / `classPresent` /
`classNonProvisionable` / `classWithdrawn` — with no schema-changed signal
threaded in at all. `classPresent`'s failure branches (both partial and total
failure) unconditionally set `frozen_since` (`policy.go`'s `classPresent`
case). Two consequences follow directly, and both are wrong:

- **`reconcile` feeds every row as `classPresent`** (§1, story 2's "to do":
  "`classPresent` for all of them, since there is no New/Absent without a live
  AWS set"). Until `held` is routed correctly, a `reconcile` codegen failure
  hits `classPresent` → `frozen_since` — freezing a type's *schema* offline,
  which is precisely what §4 says a hold must never do (a held type's schema
  is fine; only the machinery is at fault; `frozen_since` also wrongly stops
  that type's future schema refresh under `sync`).
- **Under `sync`, a Changed-present type and an Unchanged-present type are
  both `classPresent`** — `changeClass` alone cannot distinguish "the schema
  moved, and generation broke against the new bytes" (a real freeze/suppress
  case) from "the schema is identical, and only the machinery broke" (a
  `held` case). The discovery diff (point 1 above) is the only signal that
  tells them apart, and it isn't in `decide()`'s signature today.

**The fix `decide()` needs:** thread the discovery-diff signal (byteStatus, or
a new change class derived from it — e.g. `classUnchanged`/`classHeld`) into
`decide()` alongside `changeClass` and `gateResult`, and add an explicit rule:
**`reconcile` and `check` can never take the freeze/suppress branch.** Because
neither has a discovery diff at all (no AWS, no `statusChanged`/`statusUnchanged`
signal), every failure they produce is, by construction, a `held` candidate —
there is no "the schema moved" case for them to distinguish it from. Under
`sync`, the discovery diff decides which branch a `classPresent` failure takes:
`statusChanged` + failure → today's freeze/suppress path (schema-driven,
unchanged from §7's existing policy table); `statusUnchanged` + failure →
`held` (§4 below).

With that established, the output diff below is exactly what it was already
described as: a second-stage classification, computed only once a row is
already routed to `held`, that decides *which* `held` reason category applies.

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

**This creates a direct contradiction with §4.4/item 3 below as originally
stated, and it needs a pinned resolution, not just a note.** §4.1 says a
`codegen_error` hold ships without blocking the run. §4.4 item 3 says `lint`
fails on any `held_*` marker. But the `sync`/`reconcile` run that *creates*
the hold is what writes `held_*` into `all_schemas.hcl` in the first place —
so the very next `lint` (plausibly run in CI on that same `sync`/`reconcile`
PR) fails immediately on the marker the run just correctly, deliberately,
non-blockingly produced. "A hold surviving past the PR that should have fixed
it" implicitly assumed a hold is created and cleared within one flow; a
codegen fix is a *separate* engine PR from the weekly `sync` PR that first
records the hold, so the hold necessarily rides in the `sync` PR and would
trip `lint` on arrival, not after it lingers.

**Resolution: `held` gets a since-date, mirroring `frozen_since`.** A new
`held_since_<artifact>` (date, set alongside `held_<artifact>` the run a hold
is first created) lets `lint` distinguish a hold that just landed from one
that has lingered: `lint` fails only on a `held_*` marker older than some
grace window (illustrative: one release cycle), not on one created this run.
This was chosen over making `lint` merely advisory on `held_*`, because
advisory would blunt the "impossible to miss" property §4.4 exists to
guarantee — a since-date keeps `lint` a hard, unconditional gate while giving
a just-landed hold the same kind of grace `frozen_since` implicitly gets
today (nothing today fails `lint`/`check` merely for being frozen; only a
reason-less freeze does). The real preventive gate against ever landing a
`codegen_error` in the first place is `check`, run on the *engine* PR, before
it ever reaches `sync` — `lint`'s job on `held_since` is to catch a hold that
*should* have been fixed by now and wasn't, not to block the `sync` PR that
correctly, safely shipped one.

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
3. **`lint` fails on a stale `held_*` marker**: past the §4.1 grace window
   (`held_since_*` older than the grace period), a hold is an anomaly — it
   survived past the PR that should have fixed it.
4. **`check` is the real preventive gate**, run on the engine PR itself,
   before a `codegen_error` ever reaches a `sync`/`reconcile` PR at all (§1,
   story 3) — catching the regression at its source is strictly better than
   catching it after it ships, which `lint`'s grace-windowed check (item 3)
   exists only to backstop.

### 4.5 Self-clearing

Because `sync`/`reconcile` re-attempt every row every run, a held artifact is
re-evaluated every time one of them runs against it. The run after the
regression is fixed, that artifact regenerates/compiles cleanly, is promoted,
and its marker clears automatically — no `recheck` step needed for this
category (§1, story 5's "why not held too" note).

## 5. Overlay markers (near-term shape)

`held` follows today's two-attribute-per-fact convention (bool + separate
`_reason` string), matching `suppress_*_generation`/`suppression_reason_*` and
`frozen_since`/`frozen_reason` — plus one addition, `held_since_*`, needed for
`lint`'s grace-window check (§4.1, §4.4) that those two levers don't otherwise
require:

- `held_resource` / `held_singular_data_source` / `held_plural_data_source`
  (bool).
- `held_reason_resource` / `_singular_data_source` / `_plural_data_source`
  (string, `category: detail`; category is `codegen_error` only —
  `toolchain_error` blocks and is never persisted, §4.2).
- `held_since_resource` / `_singular_data_source` / `_plural_data_source`
  (date, set the run a hold first appears, mirroring `frozen_since`'s own
  shape). This is what lets `lint` (§1, story 4) distinguish a hold that just
  landed from one that has lingered past the grace window — without it,
  `lint` would have to treat every hold as equally urgent, reintroducing the
  §4.1/§4.4 contradiction a since-date exists to resolve.

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

Ordered so each step is independently mergeable and leaves the tool working.
**Revised from the original draft**: `held`'s routing (item 3 below) cannot
land after `reconcile`/`check` (items 4/5) — §3's routing bug means
`reconcile`, in particular, would mis-freeze offline on its very first
generation failure if it shipped without `held` already wired in. Held
routing now lands with (not after) the pipeline that first needs it.

1. **Rename pass first, no behavior change.** `update`→`sync`, `heal`→`recheck`,
   today's `check`→`lint` (flags, functions, doc strings, Makefile targets,
   README, runbook). Zero risk, unblocks everything else being named correctly
   from the start. Delete `run_generate.go`'s `runGenerate` in the same pass —
   its behavior is being replaced, not carried forward, so there is nothing to
   preserve by keeping it alive under a new name. **Also delete/rewrite
   `run_generate_test.go`** (`run_generate_test.go:29` calls `runGenerate`
   directly, so it breaks the moment `runGenerate` is deleted) — its
   `generateCorpus`-level coverage is not lost, since `generateCorpus` itself
   (the function `runGenerate` merely wrapped) has independent test coverage
   already (`parity_test.go`, `emit_test.go`, `corpus_test.go`,
   `run_update_test.go`), so nothing needs to be re-added, only the
   `runGenerate`-specific test removed or rewritten against whatever replaces
   it (item 3).
2. **Extract the shared offline pipeline.** A function that takes "every
   overlay row as a candidate" and runs candidate-build → `refreshCandidate` →
   `compileFixpoint` → `decide()`, returning the settled decisions —
   factored out of `runUpdate`'s AWS-specific version so `reconcile` and
   `check` can both call it with an offline row source.
3. **`held` policy + byte-diff classification, landing together with
   `reconcile` (not after it).** The output-diff classification (§3's second
   diff → `codegen_error`/`toolchain_error`), the post-revert-build
   discriminator (§4), the volume threshold (§4.3), **and** the routing rule
   §3 establishes — `reconcile`/`check` have no discovery diff, so every
   failure they produce routes to `held`, never to freeze/suppress — are all
   implemented in the shared pipeline (item 2) before or exactly alongside
   `reconcile`'s own wiring (item 4). This is what makes items 4/5
   independently safe to merge: `reconcile` is never live without `held`
   routing already in place to catch its failures correctly.
4. **`reconcile`**: wire the shared pipeline (now `held`-aware, item 3) to
   promotion + overlay write, offline row source, no AWS.
5. **`check`**: wire the same shared pipeline, stop before promotion, report,
   set the exit code.
6. **`sync`'s corpus-wide gate + discovery-diff routing.** Extend
   `buildCandidates` (or its replacement) so `statusUnchanged` rows also
   become candidates for the gate (not just New/Changed) — the core gap this
   whole redesign exists to close — **and** thread the discovery-diff signal
   (§3) into `decide()` so a `sync` failure on a `statusChanged` row still
   takes today's freeze/suppress path while a failure on a `statusUnchanged`
   row takes the `held` path. This is the one behavior change to `sync`
   beyond its rename, and it depends on item 3's routing rule already
   existing.
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
- **New, from review — the exact shape of the `decide()` signature change
  (§3).** A new `changeClass` value derived from `byteStatus` (e.g.
  `classUnchanged`), or a separate parameter threaded alongside `changeClass`
  (e.g. `decide(class, byteStatus, gr, today)`)? Either closes the routing
  bug; not yet decided which reads more clearly at the call sites in
  `runUpdate`/the new shared pipeline (§6 item 2/3).
- **New, from review — `lint`'s exact grace-window length (§4.1, §4.4).**
  Illustrative only so far ("one release cycle"); needs a real value before
  `held_since_*`'s anomaly check (§1 story 4, §5) can be implemented. Should
  probably match or relate to whatever cadence `sync` runs on, so a hold has
  a bounded number of `sync` cycles to either self-clear (§4.5) or get a
  human fix before `lint` flags it.
- **New, from review — is `check`'s "fail on any output-diff, not just
  holds" widening (§1 story 3) actually pinned, or does it need its own
  volume/noise consideration the way `held` got one (§4.3)?** A non-engine PR
  has zero diff by construction, so this shouldn't be noisy in the common
  case — but worth confirming there's no scenario (e.g. a schema-cache
  timestamp or other incidental byte churn unrelated to codegen) that would
  make ordinary, unrelated PRs fail `check` on spurious diff. If such a case
  exists, the output-diff comparison may need to exclude it explicitly.
