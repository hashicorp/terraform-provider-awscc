<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Redesign punchlist

The checklist for the `sync`/`reconcile`/`check`/`lint`/`recheck` command
redesign and the one behavior change it carries (gating the whole corpus every
run). **Transient tracking doc** — deleted once the work lands. It does not
restate the design; it links to `held-artifacts-design.md` (the detail) and
`bigdiffer-design.md` (the durable model these fold into).

Simplified after review: the earlier `held` state (per-artifact
`codegen_error`/`toolchain_error` markers, self-clearing, a new `lint` anomaly)
is dropped. A schema-unchanged gate failure now just **fails the run** — no new
overlay attributes. See `held-artifacts-design.md` §3.

Status shorthand: **prereq** (unblocks others) · **core** · **additive**
(independent) · **deferred**.

## The gaps

1. **DONE. Rename pass, no behavior change.** `update`→`sync`, `heal`→`recheck`,
   today's `check`→`lint` — flags, functions (`runUpdate`→`runSync`,
   `runHeal`→`runRecheck`, `runCheck`→`runLint`), doc strings, `make
   bigdiffer-*` targets. Deleted `run_generate.go`'s `runGenerate` (replaced,
   not carried forward — see item 3) **and** `run_generate_test.go` (called
   `runGenerate` at line 29; moved its one unrelated test,
   `TestCheckRegistrationUpToDate` — which exercises `checkRegistrationUpToDate`
   in `main.go` — into `main_test.go` rather than lose its coverage). Also
   renamed the hidden probe re-exec surface (`-heal-probe-artifact` →
   `-recheck-probe-artifact`, `runHealProbeArtifact` → `runRecheckProbeArtifact`).
   **`-reconcile`/`-check` are deliberately not flags yet** — no dispatch case
   until item 4/5; `GNUmakefile`'s `bigdiffer-generate` was removed (with a
   comment noting why) rather than point at a nonexistent flag. README/runbook
   were left un-updated on purpose (docs stay out of sync until the final
   reconciliation, item 8). `GNUmakefile` and `.github/workflows/bigdiffer.yml`
   *were* updated — they are executable config the rename mechanically breaks
   (CI ran `-check`; `make bigdiffer-update` invoked `-update`), not narrative
   docs. Verified: `gofmt -l`, `go build ./...`, `go vet`, `go test
   ./internal/tools/bigdiffer/... -race -short` and `-timeout 20m` (full
   parity), `impi`, and a real `go run ./internal/tools/bigdiffer -lint` all
   pass. *(prereq)* Detail: `held-artifacts-design.md` §0, §5 step 1.
2. **DONE. Extract the shared offline pipeline.** Candidate-build →
   `refreshCandidate` → `compileFixpoint` → `decide()`, factored out of
   `runSync`'s AWS-specific version into `pipeline.go`/`settleBatch` so an
   offline row source can drive it too. Pure extraction: `decide()` is
   byte-for-byte unchanged — the source routing (item 3) is not this item's
   scope, so `reconcile`/`check` still cannot safely call `settleBatch` until
   item 3 lands (they would mis-freeze on their first failure). `runSync` now
   builds `baseDecisions` + the pre-fixpoint overlay read, then calls
   `settleBatch` for everything from the candidate loop through the compile gate
   and list-resource coupling check, receiving a `settledBatch`. Verified:
   `gofmt`/`vet`/`build` clean, `go test ./internal/tools/bigdiffer/... -race
   -short` and `-timeout 20m` (full parity + the batch-atomicity e2e tests that
   exercise this path) pass, `impi` clean, `-lint` passes. *(prereq)* Unblocks
   items 3–5. Detail: `held-artifacts-design.md` §1, §5 step 2.
3. **DONE. Whole-corpus gate + source routing — the one behavior change.**
   Extended `buildCandidates` so `statusUnchanged` rows also become
   candidates, not just New/Changed — the actual gap the redesign closes —
   classed `classPresentUnchanged` (a new `changeClass`) when the row is
   already in the overlay, distinct from `classPresent`. `decide()` gained a
   `classPresentUnchanged` branch: on failure it sets no overlay attributes
   at all (no freeze, no suppress — there is no safe partial action once the
   schema is proven unchanged and generation/build still broke) and instead
   returns `machineryFailure: true` on the `policyDecision`. First attempt
   used a per-artifact `held` state (`codegen_error`/`toolchain_error`
   markers, a byte-diff classifier, self-clearing, a new `lint` anomaly) —
   dropped after review found it added real mechanism for no behavioral
   payoff: the only case `held` could let a run proceed is exactly the case
   `lint`/`check` must still hard-fail on eventually, so `held` reduced to
   "fail later, with more bookkeeping," not "don't fail." `held-artifacts-design.md`
   §3 has the full write-up, including why the failure is all-or-nothing (a
   `statusUnchanged` failure means the engine itself regressed, which makes
   every `statusChanged` type this same run compiled *by that same engine*
   untrustworthy too — not just the type that happened to fail).
   `machineryFailures` (`pipeline.go`) is the load-bearing piece: it scans
   `settleBatch`'s returned decisions for `machineryFailure` and must be
   called by every promoting caller, because `compileFixpoint` can reach a
   green build while still leaving a `machineryFailure` decision behind (a
   broken new artifact reverted to its still-compiling committed file settles
   the fixpoint cleanly) — checking only `compileFixpoint`'s own error return
   would silently promote everything else while that one type sits at
   last-good, exactly the bug this item exists to remove. `runSync` calls it
   immediately after `settleBatch` returns, before `changelogEntries`/
   promotion, and aborts with per-type-and-artifact blame on any hit — never
   promoting even the `statusChanged` half that compiled. `settleBatch`
   itself stays caller-agnostic (surfaces the flag, never aborts on its own),
   so `check` (item 5, no promote step) can reuse the identical scan for its
   own report/exit-code. Verified: `gofmt`/`vet`/`build` clean, new
   `TestDecide` cases for all three `classPresentUnchanged` shapes (ok,
   partial failure, total failure) and a new `TestMachineryFailures` (five
   cases, including the load-bearing "flagged but no compileFixpoint error"
   one) pass, `TestBuildCandidates` updated for the widened gate plus a new
   edge-case test for the rare statusChanged-but-absent-from-overlay ->
   classNew path, full suite `-race -short` and `-timeout 20m` (full parity)
   both pass, `impi` clean. *(prereq for items 4/5, core)* Detail:
   `held-artifacts-design.md` §2, §3, §5 step 3; §0 story 1.
4. **DONE. `reconcile`.** Wired the shared pipeline (item 2), now
   source-routed (item 3), to promotion + overlay write, offline, no AWS —
   `-generate`'s replacement (`run_reconcile.go`, `runReconcile` +
   `reconcileCandidates`). Both pinned conventions honored:
   `reconcileCandidates` classes every candidate `classPresentUnchanged`
   (never `classPresent`) and excludes frozen rows entirely before they ever
   become candidates. `runReconcile` mirrors `runSync`'s structure exactly:
   builds candidates → `settleBatch` → `machineryFailures` check (abort,
   promote nothing, on any hit) → promote → reconcile the overlay → emit
   aggregates → write the CHANGELOG fragment (reused unchanged;
   `classPresentUnchanged`'s candidates never trigger the `classNew` "every
   promoted artifact is new" branch, so only a genuine backlog lift — a
   previously-suppressed artifact reconcile just fixed — gets a changelog
   entry, never a routine re-promotion).

   **A real bug found only by actually running `-reconcile` against the repo
   (not caught by any unit test written first):** the count-header line
   ("# N CloudFormation resource types schemas are available...") silently
   changed from 1581 to 1593 on a real run — `normalizeWithDecisions`
   computes that line from `len(base)`, correct for `sync` (`base` is the
   live AWS-discovered set) but wrong for `reconcile`, which has no live set
   and passes its own `overlayRows` as `base` — a structurally different,
   larger count once any row is frozen/retained/non-provisionable. Fixed by
   restoring the committed count-header line verbatim after
   `normalizeWithDecisions` runs, mirroring `runLint`'s own established
   principle that this line is unknowable offline (its count-header
   comparison already excludes it). Re-ran `-reconcile` against the real
   repo after the fix: clean, byte-for-byte no-op diff (confirmed via
   `git status`/`git diff` — zero changes to `internal/provider/`,
   `internal/service/`, `CHANGELOG.md`).

   Three additional review points, all addressed: (1) no fully-isolated
   end-to-end `runReconcile` test exists — infeasible in isolation, since the
   compile gate always builds `cfg.repoRoot` for real (documented precedent:
   `compile_fixpoint_test.go`'s `TestCompileGateFailureBlocksPromotion`) and
   `runReconcile` derives `outputRoot`/`repoRoot` entirely from the overlay's
   own on-disk location with no injection point — the underlying pipeline is
   already covered end to end elsewhere, and the manual real-repo run above
   is the practical substitute; (2) the final "Done" tally no longer reports
   an always-zero "N frozen/suppressed" (reconcile never freezes/suppresses;
   frozen rows are excluded before candidate-building) — now reports "all
   generated OK," since every decision still standing after the
   `machineryFailures` check is, by construction, clean; (3) `reconcileCandidates`
   now returns and reports `cacheMissSkipped` alongside `frozenSkipped`, so a
   cache-miss row is visible in the progress line instead of silently
   vanishing.

   Verified: `gofmt`/`vet`/`build` (whole module) clean, new
   `TestReconcileCandidates*` (five tests: class assignment, frozen
   exclusion, cache-path convention, all-frozen edge case, cache-miss
   counted separately from frozen) and `TestReconcileCountHeaderIsNeverTrustedFromRecompute`
   (the count-header regression, exercising the actual fix logic directly)
   pass, full suite `-race -short` and `-timeout 20m` both pass, `impi`
   clean, and the real `-reconcile` run above. *(core)* Detail:
   `held-artifacts-design.md` §4, §5 step 4; §0 story 2.
5. **`check`.** Wire the same pipeline, stop before promotion, report every
   failure, exit non-zero on any failure **or** any output-diff (regenerated ≠
   committed, even when it compiles — the "changed the engine, didn't reconcile
   and commit" case). Genuinely new. *(core)* Detail: `held-artifacts-design.md`
   §4, §5 step 5; §0 story 3.
6. **`recheck`'s reasoned-row mode.** An opt-in flag (name TBD) widening
   `needsHealing`'s predicate from "reason empty/`unknown`" to "active,
   regardless of reason" — revisit a year-old suppress/freeze on demand. `lint`
   needs no code change beyond its item-1 rename. *(additive)* Detail:
   `held-artifacts-design.md` §4, §5 step 6; §0 story 5.
7. **Reporting.** Per-artifact blame on a hard-errored run (types, artifacts,
   `go build` errors) — the failing run's own output, no overlay writes.
   *(core)* Detail: `held-artifacts-design.md` §5 step 7.
8. **Reconcile every other doc in one pass, last — not incrementally.**
   `held-artifacts-design.md` and this file are the *only* docs that change
   during implementation. Everything else — `bigdiffer-design.md` (its
   out-of-sync notice, §0's command table, §6 the gates, §7 the policy note
   that a `statusUnchanged` failure fails the whole `sync` run and promotes
   nothing — including the `statusChanged` half that compiled — because an
   engine regression makes every type it touched this run untrustworthy, not
   just the one that failed), the README, and the runbook — is updated
   together, once, after the redesign works. *(core, last)*
9. **Delete this file and `held-artifacts-design.md`.** Once item 8 is
   committed, both transient docs have nothing left to track. *(core, last)*

## Open questions (need a decision, not just work)

Tracked in `held-artifacts-design.md` §6:

- Exact flag name for item 6's reasoned-row mode.
- ~~The shape of `decide()`'s discovery-diff signal (item 3)~~ — **resolved**:
  a new `changeClass` value, `classPresentUnchanged`, not a separate parameter.
- Item 5's "fail on any output-diff" relies on deterministic generation
  (`TestFullCorpusParity` already does); confirm no incidental non-determinism
  (timestamps, map order) could make an unrelated PR fail `check` on spurious
  diff.
- How `check` (item 5) is wired into CI, and how it's scoped to engine-touching
  PRs rather than every PR.
- **Deferred polish, from review (item 3):** `runSync`'s abort message is the
  formatted, per-artifact `machineryFailures` report only when
  `compileFixpoint` reaches green; in the broad-toolchain shape (committed
  files also won't build), the fixpoint hard-errors first and `runSync`
  surfaces that raw `compile gate: ...` error instead. Both abort correctly;
  `go build` in CI backstops the broad case regardless. Not urgent — revisit
  if the raw error proves hard to read in practice.
