<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer command redesign: sync / reconcile / check / lint / recheck

> **Transient design doc** for the command-surface redesign and the one
> behavior change it carries: gating the whole corpus every run. When it lands,
> the durable parts fold into `bigdiffer-design.md` (§0 stories, §6 the gates,
> §7 the policy) and this file plus `redesign-punchlist.md` are deleted.
>
> **Simplified after review.** An earlier draft added a per-artifact `held`
> state (`codegen_error`/`toolchain_error` markers, self-clearing, a new `lint`
> anomaly). That is dropped — §3 explains why. The file keeps its old name; its
> subject is now the redesign, not `held`.

## 0. The five user stories (north star)

Told plainly, no tool vocabulary — these are what the command surface must
serve (restated from `bigdiffer-design.md` §0), each with the command it lands
on:

1. **"It's been a week. Go see what's new and bring the provider up to date."**
   Ask AWS what changed, regenerate what needs it, gate it, ship it, update the
   bookkeeping. → **`sync`**
2. **"I changed how the code gets written. Redo everything the new way, don't
   ask AWS — I want today's version of everything, generated fresh."** Re-run
   the machinery against what's already known, gate it, make the output real.
   → **`reconcile`**
3. **"Before I do that — will this break anything? Just tell me, change
   nothing."** The same attempt as #2, reported instead of kept — asked by a
   person occasionally or by CI every time, no difference between the askers.
   → **`check`**
4. **"Never mind the machinery — is the bookkeeping itself in good shape?"**
   Duplicates, contradictions, unexplained gaps, things marked broken with no
   reason. → **`lint`**
5. **"Someone decided a while back this was broken. Is that still true?"**
   Revisit an old judgment call, on demand. → **`recheck`**

| # | Story | Command | AWS? | Writes? |
|---|-------|---------|------|---------|
| 1 | Weekly bring-current | `sync` | yes | yes |
| 2 | Land a machinery change, no crawl | `reconcile` | no | yes |
| 3 | Would a machinery change break anything? | `check` | no | no |
| 4 | Is the bookkeeping healthy? | `lint` | no | no |
| 5 | Revisit an old decision | `recheck` | no | no |

Naming: `update`→`sync` and `heal`→`recheck` keep their spirit; today's
`check`→`lint` frees `check` for story 3; `reconcile` is new. Chosen so no two
blur in sound or meaning. `reconcile` was pressure-tested against
`bigdiffer-design.md` §12's durable use of the word (the AWS-facing
reconciliation) and kept deliberately: story 2 reconciles the committed overlay
against fresh generation output — same kind of operation (bring two things into
agreement), no AWS required for that one.

## 1. One engine, three callers

`sync`, `reconcile`, and `check` run the **same** generate → compile → decide
pipeline over the corpus (`settleBatch`, already extracted). They differ only
in:

- **Row source** — `sync` crawls AWS and diffs against the cache;
  `reconcile`/`check` read the committed overlay + cache directly, every row a
  candidate (offline, there is no "changed" — just "attempt it").
- **After the batch settles** — `sync`/`reconcile` promote; `check` reports.

`lint` and `recheck` never enter that pipeline — they are read-only overlay
operations, small and independent.

## 2. The gap, and the one-line fix

Both gates run only on New/Changed types today (`buildCandidates`,
`update.go:40`, filters to `statusNew`/`statusChanged`). So a change to the
*machinery* — a template, `codegen/`, naming, the Go toolchain — is never
exercised against the ~1580 schema-unchanged types. Their committed output
silently ages, and nothing catches a machinery regression on them until
something downstream breaks. `TestFullCorpusParity` is the only current guard,
and it retires with the legacy engine (`bigdiffer-design.md` §10).

**Fix: gate the whole corpus every run** — generate + compile every type, not
just New/Changed. That is the entire behavioral change. The rest of this doc is
about what to *do* with a failure the wider gate now surfaces.

## 3. Disposition by source

A gate failure is handled by **why** it failed, which is fully determined by
whether the schema changed:

- **schema-changed failure** (`statusChanged`): the existing freeze/suppress
  policy, unchanged. AWS handed us a schema that won't generate or compile —
  external breakage — so defend the shipped provider (keep last-good, pin the
  schema) and ship everything else. `bigdiffer-design.md` §7 is untouched.
- **schema-unchanged failure** (`statusUnchanged`): **fail the run**, loudly,
  with per-artifact blame (which types, which artifacts, the `go build` error).
  Nothing is written to the overlay, nothing self-clears, no new state. The
  gate failing *is* the surfacing.

The asymmetry is the point. A schema break is **external** — you can't fix
AWS's schema, only pin it — so you freeze and move on. A machinery break is
**self-inflicted** — your own template/codegen/toolchain change — so the honest
response is to fix it (or revert it), not to have the tool paper over it.
bigdiffer records only states it can act on; a machinery break it cannot
mitigate, so it fails and asks a human.

### Why not a `held` state (dropped from the earlier draft)

Two realizations killed it:

- **The unmitigable case.** If the schema is unchanged *and* regeneration
  produces byte-identical output that still won't compile (a toolchain or
  dependency break), there is nothing to fall back to — the committed file *is*
  the failing file. bigdiffer can do nothing but fail; marking the overlay
  "held" and moving on would be dishonest, since `go build` (a required CI
  check) fails regardless.
- **The redundancy.** The only case where a `held` marker could let the run
  *proceed* is a codegen change whose old output still compiles (keep last-good,
  ship). But we require CI to block on a machinery regression — so a marker
  would block the PR anyway. `held` then reduces to "fail the run, plus a
  self-clearing marker, a byte-diff classifier, and a new `lint` anomaly" —
  elaborate bookkeeping around an outcome that is just *blocked until a human
  fixes the machinery.* Failing the run gets the same result with none of the
  machinery, and **adds no new overlay attributes.**

### Escape hatch for a genuinely hard straggler

If one type is legitimately hard to fix under a new engine and you want to ship
the rest, a human `suppress`es that artifact deliberately — the existing
`suppress_*` lever and reason taxonomy. A standing, human-owned decision with a
reason, not a state the tool invented. No new mechanism.

### Why all-or-nothing, precisely (pinned)

A `statusUnchanged` failure blocks the **entire** `sync` run — nothing
promotes, not even the `statusChanged` half that compiled cleanly this run.
Not a limitation to work around later; the correct default, for a reason
sharper than "stay consistent with existing batch atomicity":

**A machinery failure impugns the whole run's output, not just the failing
types.** A `statusUnchanged` failure means the engine itself regressed — and
every `statusChanged` type this same run was generated by that same engine.
Those types "compiled," but an engine bug that produces broken output for one
schema shape can equally produce output that compiles yet is subtly wrong for
another. A compile failure cannot tell you whether the regression is narrow
(one schema shape the bug happens to hit) or broad (every type, silently). Once
you can't tell, shipping the `statusChanged` half ships this week's new/updated
schemas out of a generator you have just caught being broken — trusting output
you have direct evidence not to trust. Don't ship anything a known-suspect
engine produced this run; that is not a weakened invariant, it is what "never
ship broken output" actually requires once the gate can see a self-inflicted
failure at all. (Previously it couldn't — New/Changed-only gating never
exercised the corpus widely enough to catch this, so the question never came
up; §2.)

### The routing this still requires

`sync` must tell a `statusChanged` failure (→ freeze/suppress) from a
`statusUnchanged` failure (→ fail the run). Today `decide()` (`policy.go`)
routes only on `changeClass`, and both are `classPresent`, so without this bit a
whole-corpus gate would send unchanged-schema failures down the
`classPresent`→`frozen_since` branch — silently freezing schema-fine types and
stopping their future refresh (the wrong lever, the worst outcome). So `decide()`
needs the discovery-diff signal (`statusChanged` vs `statusUnchanged`) threaded
in. `reconcile` and `check` are offline and have no discovery diff at all: every
failure they see is, by construction, a machinery failure → fail the run. They
never freeze or suppress.

**`reconcile`/`check` must exclude frozen rows (pinned now, before item 4).**
`sync` already excludes frozen types from the gate entirely — `classifyChange`
returns `statusFrozen` before any byte comparison, upstream of and prior to
`buildCandidates`, so a frozen row never becomes a candidate at all (a type is
frozen precisely because its schema is authoritative at its pinned bytes, not
because the tool should keep re-attempting it). `reconcile`'s offline row
source has no `classifyChange` call to inherit that exclusion from for free —
it must apply the same exclusion explicitly (skip any row with `frozen_since`
set, or generate from its pinned bytes and treat that as authoritative rather
than a fresh attempt). Missing this is not a cosmetic gap: a frozen type is
frozen *because* it fails to generate/compile, so feeding it through
`classPresentUnchanged` would hit `machineryFailure` on every single
`reconcile` run — the command would never succeed at all, not even in the
best case of "the engine change fixed everything else."

## 4. Per-command behavior

- **`sync`** (AWS, writes): crawl, diff, gate the **whole** corpus.
  `statusChanged` failure → freeze/suppress, proceed. `statusUnchanged` failure
  → fail the run, **and promote nothing at all this run — not even the
  `statusChanged` half that compiled** (pinned; see "Why all-or-nothing,
  precisely" below). Absent any `statusUnchanged` failure, promote any type
  whose regenerated output differs from committed *and* compiles — valid
  machinery-driven drift — so an engine change self-heals across the corpus at
  the next `sync` instead of aging silently.
- **`reconcile`** (offline, writes): regenerate the whole corpus from the
  committed overlay + cache through the same gate; any failure fails the run;
  promote the clean result. Replaces the deleted `-generate` (which never
  gated).
- **`check`** (offline, read-only): `reconcile`'s pipeline with promotion off.
  Reports every failure, then exits non-zero on any failure **or** any
  output-diff (regenerated ≠ committed) even when it compiles — the latter being
  the "you changed the engine but didn't `reconcile` + commit" case. A
  non-engine PR has zero diff and passes; an engine PR must `reconcile` + commit
  until `check` is clean. This is the CI gate on engine-touching PRs and the
  replacement for `TestFullCorpusParity`'s guard role (self-referential: current
  engine vs. committed output).
- **`lint`** (offline, read-only): today's `-check` renamed. Existing checks
  (normalization, duplicates, naming, reason-less suppress/freeze) unchanged.
  **No new anomaly** — with `held` gone there is nothing new to flag.
- **`recheck`** (offline, read-only): today's `-heal` renamed, plus an opt-in
  flag that widens `needsHealing` (`heal.go:60`) from "reason empty/`unknown`"
  to "active, regardless of reason," to revisit a year-old suppress/freeze on
  demand.

## 5. Implementation plan

Ordered so each step is independently mergeable and leaves the tool working.

1. **Rename pass** *(done)*. `update`→`sync`, `heal`→`recheck`, `check`→`lint`
   (flags, functions, doc strings, Makefile). Deleted `runGenerate` +
   `run_generate_test.go`.
2. **Extract the shared pipeline** *(done)*. `settleBatch`: candidate-build →
   `refreshCandidate` → `compileFixpoint` → `decide()`, factored out of
   `runSync`. Pure extraction; `decide()` unchanged, so `reconcile`/`check`
   cannot call it safely until step 3's routing lands.
3. **Whole-corpus gate + source routing** *(done)*. Gated `statusUnchanged`
   rows too (`buildCandidates`), and gave `decide()` a `classPresentUnchanged`
   branch so a `statusChanged` failure still freezes/suppresses while a
   `statusUnchanged` failure sets `machineryFailure` instead (no overlay
   attributes at all — nothing safe to write). `machineryFailures`
   (`pipeline.go`) scans a settled batch for the flag; `runSync` calls it
   right after `settleBatch` returns, before any promotion, and aborts with
   full per-artifact blame on any hit, promoting nothing that run. Two
   forward-looking notes from review, both minor, neither blocking:
   - **Reporting asymmetry.** The formatted, per-artifact `machineryFailures`
     report only appears when `compileFixpoint` itself reaches green (the
     committed fallback compiles). In the broad-toolchain shape, where the
     committed files also won't build, the fixpoint can't reach green and
     `runSync` returns the raw `compile gate: ...` error instead of the
     formatted report — both abort correctly, and `go build` in CI backstops
     that case regardless, so this is a message-quality nuance, not a
     correctness hole. Worth a future polish if the raw error proves hard to
     read in practice; not urgent.
   - **Frozen-row exclusion, pinned ahead of item 4** — see "The routing this
     still requires," above: `reconcile`'s offline row source must exclude
     frozen rows explicitly (or treat their pinned bytes as authoritative),
     the same outcome `sync` gets for free via `classifyChange`'s
     `statusFrozen` short-circuit upstream of `buildCandidates`. Missing this
     would make `reconcile` hit `machineryFailure` on every frozen type,
     every run, with no path to ever completing.
4. **`reconcile`**: wire the shared pipeline to promotion, offline row source.
5. **`check`**: wire the same pipeline, report-only, exit non-zero on any
   failure or output-diff.
6. **`recheck`'s reasoned-row flag** (small, independent). `lint` needs no code
   change beyond its step-1 rename.
7. **Reporting**: per-artifact blame on a hard-errored run (types, artifacts,
   `go build` errors) — the failing run's own output; no overlay writes.
8. **Doc reconciliation, one pass, last.** This file and `redesign-punchlist.md`
   are the only docs touched during steps 1–7. Then update `bigdiffer-design.md`
   (§0 command table, §6, and a §7 note that a `statusUnchanged` failure fails
   the whole `sync` run and promotes nothing — including the `statusChanged`
   half that compiled — because a self-inflicted engine regression makes every
   type that engine touched this run untrustworthy, not just the type that
   happened to fail; §3's "Why all-or-nothing, precisely" carries the actual
   rationale, not just the outcome), the README, and the runbook — together —
   and delete both transient docs.

## 6. Open questions

- Exact flag name for `recheck`'s reasoned-row mode.
- The shape of the discovery-diff signal into `decide()`: a new `changeClass`
  (e.g. `classUnchanged`), or a `byteStatus` parameter alongside `changeClass`.
- `check`'s "fail on any output-diff" relies on generation being deterministic
  (same inputs → identical bytes, which `TestFullCorpusParity` already depends
  on); confirm no incidental non-determinism (timestamps, map order) would make
  an unrelated PR fail `check` on spurious diff.
- How `check` is wired into CI, and how it's scoped to engine-touching PRs
  (`codegen/`, templates, naming, `go.mod`/`go.sum`) rather than every PR.
