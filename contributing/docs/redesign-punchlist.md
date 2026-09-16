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
3. **Whole-corpus gate + source routing — the one behavior change.** Extend
   `buildCandidates` (or its replacement) so `statusUnchanged` rows also become
   candidates, not just New/Changed — the actual gap the redesign closes. Thread
   the discovery-diff signal (`statusChanged` vs `statusUnchanged`) into
   `decide()` so: a `statusChanged` failure still takes today's freeze/suppress
   path (unchanged), while a `statusUnchanged` failure **fails the run** with
   per-artifact blame instead of falling into `classPresent`→`frozen_since`
   (which would silently freeze schema-fine types — the worst outcome). Offline,
   `reconcile`/`check` have no discovery diff, so every failure they see is a
   machinery failure → fail the run; they never freeze/suppress. `sync` also
   promotes any type whose regenerated output differs from committed and
   compiles (valid drift self-heals). *(prereq for items 4/5, core)* Detail:
   `held-artifacts-design.md` §2, §3, §5 step 3; §0 story 1.
4. **`reconcile`.** Wire the shared pipeline (item 2), now source-routed
   (item 3), to promotion + overlay write, offline, no AWS — `-generate`'s
   replacement. *(core)* Detail: `held-artifacts-design.md` §4, §5 step 4;
   §0 story 2.
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
   that unchanged-schema failures hard-error), the README, and the runbook —
   is updated together, once, after the redesign works. *(core, last)*
9. **Delete this file and `held-artifacts-design.md`.** Once item 8 is
   committed, both transient docs have nothing left to track. *(core, last)*

## Open questions (need a decision, not just work)

Tracked in `held-artifacts-design.md` §6:

- Exact flag name for item 6's reasoned-row mode.
- The shape of `decide()`'s discovery-diff signal (item 3): a new `changeClass`
  value, or a `byteStatus` parameter alongside it.
- Item 5's "fail on any output-diff" relies on deterministic generation
  (`TestFullCorpusParity` already does); confirm no incidental non-determinism
  (timestamps, map order) could make an unrelated PR fail `check` on spurious
  diff.
- How `check` (item 5) is wired into CI, and how it's scoped to engine-touching
  PRs rather than every PR.
