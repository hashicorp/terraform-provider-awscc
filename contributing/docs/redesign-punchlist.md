<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Redesign punchlist

The checklist for the `sync`/`reconcile`/`check`/`lint`/`recheck` command
redesign and the `held` feature it carries. This is a **transient tracking
doc** — it is deleted once the work lands. It does not restate the design; it
links to `held-artifacts-design.md` (the detail) and `bigdiffer-design.md`
(the durable model these changes fold into once shipped).

Status shorthand: **prereq** (unblocks other items) · **core** · **additive**
(independent, no dependency on the rest) · **deferred** (real, but not now).

Ordering matches `held-artifacts-design.md` §6's implementation plan — each
item below is independently mergeable and leaves the tool working.

## The gaps

1. **Rename pass, no behavior change.** `update`→`sync`, `heal`→`recheck`,
   today's `check`→`lint` — flags, functions (`runUpdate`→`runSync`,
   `runHeal`→`runRecheck`, `runCheck`→`runLint`), doc strings, `make
   bigdiffer-*` targets, README, runbook. Delete `run_generate.go`'s
   `runGenerate` in the same pass (replaced, not carried forward — see item 3).
   *(prereq)* Detail: `held-artifacts-design.md` §1 (all five stories), §6
   step 1.
2. **Extract the shared offline pipeline.** Candidate-build →
   `refreshCandidate` → `compileFixpoint` → `decide()`, factored out of
   `runUpdate`'s AWS-specific version so an offline row source (every overlay
   row, no `discover()` call) can drive it too. *(prereq)* Unblocks items 3
   and 4. Detail: `held-artifacts-design.md` §2, §6 step 2.
3. **`reconcile`.** Wire the shared pipeline (item 2) to promotion + overlay
   write, offline, no AWS — `-generate`'s actual replacement. *(core)* Detail:
   `held-artifacts-design.md` §1 story 2, §6 step 3.
4. **`check`.** Wire the same shared pipeline, stop before promotion, report,
   set the exit code (non-zero on any hold or a tripped volume threshold).
   Genuinely new — no command like this exists today. *(core)* Detail:
   `held-artifacts-design.md` §1 story 3, §6 step 4.
5. **`held` policy + `codegen_error`/`toolchain_error` classification.**
   Byte-diff regenerated vs. committed bytes to label the cause; the
   post-revert-build result (not the byte-diff) decides hold-vs-block; the
   volume-threshold safety valve for a large `codegen_error` batch. Implemented
   once in the shared pipeline (item 2) so `sync`/`reconcile`/`check` get it
   simultaneously. *(core)* Detail: `held-artifacts-design.md` §3, §4, §6
   step 5.
6. **`sync`'s corpus-wide gate.** Extend `buildCandidates` (or its
   replacement) so `statusUnchanged` rows also become candidates, not just
   New/Changed — the one behavior change to `sync` beyond its rename; this is
   the actual gap the whole redesign exists to close. *(core)* Detail:
   `held-artifacts-design.md` §1 story 1, §6 step 6.
7. **`lint`'s new anomaly.** Any row with `held_resource` /
   `held_singular_data_source` / `held_plural_data_source` set is reported and
   fails `lint` — one more case in `anomalyProblems()`, parallel to today's
   reason-less-suppression check. *(additive)* Detail:
   `held-artifacts-design.md` §1 story 4, §6 step 7.
8. **`recheck`'s reasoned-row mode.** An opt-in flag (name TBD) that widens
   `needsHealing`'s predicate from "reason is empty/unknown" to "active,
   regardless of reason" — revisit a year-old, already-reasoned
   suppress/freeze on demand. *(additive)* Detail:
   `held-artifacts-design.md` §1 story 5, §6 step 7; open question in §7.
9. **Reporting.** The end-of-run held summary (grouped by category, blocking
   headline first if the run blocked or tripped the threshold); `held`
   markers appearing in the `sync`/`reconcile` PR diff. *(core)* Detail:
   `held-artifacts-design.md` §4.4, §6 step 8.
10. **Reconcile every other doc in one pass, last — not incrementally.**
    Decided explicitly: `held-artifacts-design.md` and this file are the
    *only* docs that change while implementation is in progress. Everything
    else — `bigdiffer-design.md` (its out-of-sync notice, §0's command table,
    §6 the gates, §7 the policy table), the README, the runbook, and
    `suppressed-and-frozen.md` (the two new reason categories) — gets updated
    together, once, after the redesign is fully implemented and working, not
    piecemeal alongside each step above. Avoids reviewing docs against a
    design that's still moving. *(core, last, before item 11)*
11. **Delete this file and `held-artifacts-design.md`.** Once item 10's
    reconciliation is committed, both transient docs have nothing left to
    track. *(core, last)*

## Open questions (not gaps — need a decision, not just work)

Tracked in `held-artifacts-design.md` §7:

- Exact flag name for item 8's reasoned-row mode.
- Exit code of `reconcile`/`sync` on a `codegen_error` hold: exit 0 + banner,
  or non-zero?
- Exact volume threshold (item 5): fraction, absolute count, or both.
- How `check` (item 4) is wired into CI, and how CI scopes it to only the PRs
  that plausibly touch the engine.
