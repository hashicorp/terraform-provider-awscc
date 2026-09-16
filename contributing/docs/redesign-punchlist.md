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
5. **DONE. `check`.** Wired the same pipeline (`reconcileCandidates` +
   `settleBatch`, unchanged from item 4), promotion switched off
   (`run_check.go`, `runCheck`). Reports every failure via the identical
   `machineryFailures` scan `reconcile`/`sync` use, then scans everything
   `settleBatch` staged (`stagingDir/out`, `stagingDir/cache`) against the
   real, committed trees (`cfg.outputRoot`, `cfg.cacheDir`) byte-for-byte —
   `diffStagedTrees`, a new helper — and exits non-zero on either a failure
   or any byte diff, even though `check` never calls `promoteStaged` at all
   (it only ever reads the real tree, never writes it; the staging dir is
   `os.RemoveAll`'d either way). A file staged with no committed counterpart
   at all (a suppressed type a machinery fix just recovered) still counts as
   a diff — there is genuinely new output to `-reconcile` and commit, not a
   false "no diff."

   Verified against the real repo, not just synthetic fixtures: a clean
   `go run ./internal/tools/bigdiffer -check` passes with exit 0 and zero
   diffs (confirming, empirically, the open determinism question below —
   the real corpus regenerates byte-identical to committed output, twice,
   across two separate real runs). Deliberately corrupted a real,
   **non-frozen** committed file (`internal/aws/acmpca/certificate_resource_gen.go`)
   and re-ran `-check`: correctly detected the diff, named the exact file,
   and exited 1; reverted the corruption and re-ran clean. The first two
   manual attempts at this picked `AWS::Logs::LogGroup` and `AWS::S3::Bucket`
   — both frozen, so both are silently excluded from `reconcileCandidates`
   before ever reaching the diff scan, which produced a false-negative
   "check passed" the first two times and would have shipped a broken test
   if not caught by manually reading the overlay to confirm each candidate's
   `frozen_since` status before concluding the corruption should have been
   detected.

   That false start became `TestRunCheckAgainstRealRepo`
   (`run_check_test.go`) — a genuine end-to-end test against the real repo,
   feasible here in a way `TestRunReconcile` never was (`run_reconcile_test.go`'s
   doc comment): `check` never promotes, so pointing it at the real repo from
   within a test carries none of `reconcile`'s risk. Two subtests: a clean
   pass, and a corruption case that sources its target file's real path from
   an actual `settleBatch` call on one dynamically chosen non-frozen
   candidate (never a hand-picked type name) specifically so a future frozen
   row can never repeat the false-negative mistake above — the corrupted
   file is restored via `t.Cleanup` unconditionally, including on failure or
   panic. Gated behind `testing.Short()` (skipped in `-race -short`, which
   exists precisely to stay fast) since it costs ~4 minutes of real,
   full-corpus generation on top of the existing full-suite run — raises the
   full `-timeout 20m` suite from ~205s to ~455s. Also added focused,
   fast, synthetic-tree unit tests for `diffStagedTrees` itself (no-diff,
   byte-diff, missing-committed-counterpart, cache-tree coverage, empty
   staging dir) that do not pay this cost.

   Updated `main.go`'s package doc comment (was still "five modes today,"
   with `-check` listed as "designed but not yet implemented") and `run`'s
   dispatch doc comment/error message to include `-check`.

   One review point confirmed as an existing, consistent property rather
   than a gap: `diffStagedTrees` walks staged → committed only, never the
   reverse, so it cannot see a committed file the current templates/codegen
   no longer emit at all — but neither can `reconcile`'s own promotion
   (`copyTree` overwrites, never deletes), so `check`'s "passes iff a
   `reconcile` + commit would be a no-op" contract holds exactly as stated.
   Pinned as its own subsection in `held-artifacts-design.md` §4 ("Neither
   `reconcile` nor `check` removes orphaned output") since it is a
   corpus-wide property of both commands, not specific to this scan.

   Verified: `gofmt`/`vet`/`build` (whole module) clean, new
   `TestDiffStagedTrees*` (five tests) and `TestRunCheckAgainstRealRepo`
   (two subtests, full-corpus, real repo) pass, full suite `-race -short`
   and `-timeout 20m` both pass, `impi` clean, and the manual real
   `-check` runs above (clean pass, corruption detection, re-confirmed
   clean after reverting). *(core)* Detail: `held-artifacts-design.md` §4,
   §5 step 5; §0 story 3.
6. **DONE. `recheck`'s reasoned-row mode.** New opt-in flag, `-recheck-all`
   (resolves the "name TBD" open question below: chosen as a modifier on
   `-recheck` — meaningless alone — mirroring `-recheck-probe-artifact`'s
   existing `-recheck-*` naming pattern for `recheck`-related flags).
   Widened `needsHealing` from a plain predicate to `needsHealing(all bool)`:
   with `all` false (the default), unchanged — active and reason-less/`unknown`.
   With `all` true, every active fact is in scope regardless of its existing
   reason, so a year-old suppress/freeze — even one with a real, specific
   reason already recorded — can be revisited on demand (§0 story 5).
   Threaded `all` through `runRecheck` into the candidate-selection loop and
   into `writeHealReport`'s tally line (`"facts re-probed (every active
   fact, -recheck-all): N"` instead of the now-inapplicable "facts needing a
   reason" label).

   **A real bug, found by actually running `-recheck -recheck-all` against
   the real repo, not by any unit test written first:** widening the scope
   to include already-reasoned facts meant `commentOrUnknown` — the
   no-cached-schema fallback every fact still falls through to — could now
   receive a fact with a real, meaningful reason already recorded (only
   possible under `-recheck-all`, since the default scope excludes these).
   Its original behavior would have silently proposed *overwriting* that
   real reason with either a same-row comment-migration guess (which may
   describe a different fact on the same row entirely) or a bare "needs a
   human look" — actively destructive, not just unhelpful. Fixed by
   threading each fact's existing reason (`f.reason`) through
   `healArtifact`/`freezeProposal` into `commentOrUnknown`, which now keeps
   a real, non-`unknown:`-tagged existing reason as-is (appending a short
   note that there was no cached schema to weigh it against) rather than
   falling through to the comment/unknown path at all. Confirmed empirically
   against the real corpus: a real `-recheck -recheck-all` run surfaced
   **34 real rows** that hit exactly this path (frozen, with a real,
   carefully-written reason, no cached schema since a frozen row's bytes
   are never refreshed) — without the fix, all 34 would have had their real
   reasons proposed for silent replacement.

   Also confirmed the intended behavior fires correctly on the same real
   run: two genuine "lift" proposals surfaced for rows that generate cleanly
   now despite being suppressed a while back (`AWS::ARCRegionSwitch::Plan`'s
   plural data source, `AWS::AmazonMQ::Broker`'s resource and singular data
   source) — exactly story 5's "is that still true?" case. `lint` needed no
   code change beyond its item-1 rename, as anticipated.

   Verified: `gofmt`/`vet`/`build` (whole module) clean, updated
   `TestCommentOrUnknown` (two new subtests: a real existing reason kept
   verbatim and never overwritten by an unrelated comment; an existing
   reason still tagged `unknown:` correctly still falls through, not kept)
   and `TestHealFactsForNeedsHealing` (two new subtests: `-recheck-all`
   widens scope to an already-reasoned active fact; an inactive fact is
   still excluded regardless), full suite `-race -short` and `-timeout 20m`
   both pass, `impi` clean, and real `go run ./internal/tools/bigdiffer
   -recheck` (0 facts — the reason-less backlog is currently empty) and
   `-recheck -recheck-all` (557 facts re-probed, including the 34-row
   existing-reason-kept case and the two genuine lift proposals above) runs
   against the real repo, confirmed to write nothing to `all_schemas.hcl`.
   *(additive)* Detail: `held-artifacts-design.md` §4, §5 step 6; §0
   story 5.
7. **DONE. Reporting.** Per-artifact blame on a hard-errored run was already
   fully in place for the common case (types, artifacts, `go build` errors)
   via `machineryFailures` (item 3) — every promoting/reporting caller
   (`sync`/`reconcile`/`check`) already surfaces exactly this when
   `compileFixpoint` reaches green with a `machineryFailure` decision
   flagged. The one remaining gap was the "reporting asymmetry" deferred in
   item 3's review and re-raised in `held-artifacts-design.md`'s open
   questions: the `downgraded == 0` hard-stop in `compileFixpoint`
   (`compile.go`) — reached when a build failure can't be attributed to any
   staged artifact at all (base tree already broken, or bigdiffer's own
   `registrations_gen.go` has a bug) — surfaced the raw `go build` output
   with absolute file paths, unlike every other report in the tool, which
   relativizes against `repoRoot`. Fixed by routing that message through
   `relativizeBuildErrors` (already existing, `heal.go`) before formatting.
   This is a single fix in `compileFixpoint` itself, so it benefits `sync`,
   `reconcile`, and `check` uniformly with no per-caller change needed —
   exactly the payoff the shared `settleBatch` pipeline (item 2) was
   designed for. There genuinely is no per-type/per-artifact structure to
   report in this specific case (that is what "unattributable" means), so
   the fix is scoped to making the one report that exists readable, not
   inventing attribution that doesn't exist. Left the separate
   `maxFixpointRounds`-exceeded circuit breaker unchanged — its message
   already states the failure mode plainly and accumulating every round's
   downgrades would mostly repeat information, not add new signal.

   Verified: `gofmt`/`vet`/`build` (whole module) clean, updated
   `TestCompileFixpointUnattributableFailureHardStops` with a new assertion
   confirming the error is relativized (no absolute `repoRoot` path present,
   the expected relative path is), full suite `-race -short` and
   `-timeout 20m` both pass, `impi` clean. *(core)* Detail:
   `held-artifacts-design.md` §5 step 7.
8. **Reconcile every other doc in one pass, last — not incrementally.**
   `held-artifacts-design.md` and this file are the *only* docs that change
   during implementation. Everything else — `bigdiffer-design.md` (its
   out-of-sync notice, §0's command table, §6 the gates, §7 the policy note
   that a `statusUnchanged` failure fails the whole `sync` run and promotes
   nothing — including the `statusChanged` half that compiled — because an
   engine regression makes every type it touched this run untrustworthy, not
   just the one that failed), the README, and the runbook — is updated
   together, once, after the redesign works. **Also fold in the file-name
   rename, pinned now (from review, after item 4) rather than left as silent
   drift:** entry points no longer track the locked command names —
   `runSync` lives in `update.go`, `runRecheck` in `heal.go`, `runLint` has
   no file of its own (orphaned in `main.go`), while the newer `runReconcile`
   follows a `run_reconcile.go` convention. Three naming conventions
   currently coexist. Rename `update.go`→`sync.go`, `heal.go`→`recheck.go`,
   extract `runLint`→`lint.go`, and reconcile the `run_*_test.go`/bare
   `*_test.go` test-file split to match, once `check` (item 5) has landed and
   its own file-naming choice is known. *(core, last)*
9. **Delete this file and `held-artifacts-design.md`.** Once item 8 is
   committed, both transient docs have nothing left to track. *(core, last)*

## Open questions (need a decision, not just work)

Tracked in `held-artifacts-design.md` §6:

- ~~Exact flag name for item 6's reasoned-row mode.~~ — **resolved**:
  `-recheck-all`.
- ~~The shape of `decide()`'s discovery-diff signal (item 3)~~ — **resolved**:
  a new `changeClass` value, `classPresentUnchanged`, not a separate parameter.
- Item 5's "fail on any output-diff" relies on deterministic generation
  (`TestFullCorpusParity` already does); **empirically observed clean,
  item 5**: two separate real `-check` runs against the untouched corpus
  both produced zero diffs. Not a formal proof for all time — no
  incidental non-determinism (timestamps, map order) has been *seen*, but
  hasn't been exhaustively ruled out either.
- How `check` (item 5) is wired into CI, and how it's scoped to engine-touching
  PRs rather than every PR.
- ~~**Deferred polish, from review (item 3):** `runSync`'s abort message is the
  formatted, per-artifact `machineryFailures` report only when
  `compileFixpoint` reaches green; in the broad-toolchain shape (committed
  files also won't build), the fixpoint hard-errors first and `runSync`
  surfaces that raw `compile gate: ...` error instead.~~ — **resolved, item
  7**: relativized the raw error's file paths against `repoRoot`
  (`relativizeBuildErrors`), fixed once in `compileFixpoint` so `sync`,
  `reconcile`, and `check` all benefit uniformly.
