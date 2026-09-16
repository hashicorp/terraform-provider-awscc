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

Ordering matches `held-artifacts-design.md` §6's implementation plan (revised
after review — `held`'s routing now lands with, not after, `reconcile`; see
item 3).

## The gaps

1. **Rename pass, no behavior change.** `update`→`sync`, `heal`→`recheck`,
   today's `check`→`lint` — flags, functions (`runUpdate`→`runSync`,
   `runHeal`→`runRecheck`, `runCheck`→`runLint`), doc strings, `make
   bigdiffer-*` targets, README, runbook. Delete `run_generate.go`'s
   `runGenerate` in the same pass (replaced, not carried forward — see item 3)
   **and** delete/rewrite `run_generate_test.go` (calls `runGenerate`
   directly at line 29; `generateCorpus`'s own coverage elsewhere is
   unaffected). *(prereq)* Detail: `held-artifacts-design.md` §1 (all five
   stories), §6 step 1.
2. **Extract the shared offline pipeline.** Candidate-build →
   `refreshCandidate` → `compileFixpoint` → `decide()`, factored out of
   `runUpdate`'s AWS-specific version so an offline row source (every overlay
   row, no `discover()` call) can drive it too. *(prereq)* Unblocks items 3
   and 4. Detail: `held-artifacts-design.md` §2, §6 step 2.
3. **`held` policy + `codegen_error`/`toolchain_error` classification +
   freeze-vs-held routing — must land with `reconcile` (item 4), not after
   it.** Found in review: `decide()` routes purely on `changeClass` with no
   schema-changed signal, so `reconcile` (which feeds every row as
   `classPresent`) would mis-freeze offline on its first failure if `held`
   shipped later. This item now includes the routing rule itself —
   `reconcile`/`check` have no discovery diff, so every failure they produce
   is a `held` candidate, never freeze/suppress — alongside the byte-diff
   classification, the post-revert-build hold-vs-block discriminator, and the
   volume threshold. *(prereq for item 4)* Detail:
   `held-artifacts-design.md` §3, §4, §6 step 3.
4. **`reconcile`.** Wire the shared pipeline (item 2), now `held`-aware
   (item 3), to promotion + overlay write, offline, no AWS — `-generate`'s
   actual replacement. *(core)* Detail: `held-artifacts-design.md` §1
   story 2, §6 step 4.
5. **`check`.** Wire the same shared pipeline, stop before promotion, report,
   set the exit code — non-zero on any hold, a tripped volume threshold, **or
   any output-diff at all** (found in review: holds-only misses valid,
   uncommitted drift, which is still "silent staleness"). Genuinely new — no
   command like this exists today. *(core)* Detail:
   `held-artifacts-design.md` §1 story 3, §6 step 5.
6. **`sync`'s corpus-wide gate + discovery-diff routing.** Extend
   `buildCandidates` (or its replacement) so `statusUnchanged` rows also
   become candidates, not just New/Changed — the actual gap the whole
   redesign exists to close — and thread the discovery-diff signal into
   `decide()` (item 3's routing rule) so a `sync` failure on a changed schema
   still freezes/suppresses while a failure on an unchanged schema holds.
   Also promotes any output-diff that passes the gate regardless of
   discovery-diff status (found in review: `sync` must self-heal valid drift
   on unchanged rows, not just gate it). *(core)* Detail:
   `held-artifacts-design.md` §1 story 1, §6 step 6.
7. **`lint`'s new anomaly.** Any populated per-artifact slot (`resource` /
   `singular_data_source` / `plural_data_source`) is reported and fails
   `lint`, immediately — surfacing a machinery regression the moment it
   appears is the point. One more case in `anomalyProblems()`, parallel to
   today's reason-less-suppression check. *(additive)* Detail:
   `held-artifacts-design.md` §1 story 4, §4.4, §5, §6 step 7.
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
    together, once, after the redesign is
    fully implemented and working, not piecemeal alongside each step above.
    Avoids reviewing docs against a design that's still moving. *(core, last,
    before item 11)*
11. **Delete this file and `held-artifacts-design.md`.** Once item 10's
    reconciliation is committed, both transient docs have nothing left to
    track. *(core, last)*

## Open questions (not gaps — need a decision, not just work)

Tracked in `held-artifacts-design.md` §7:

- Exact flag name for item 8's reasoned-row mode.
- Exit code of `reconcile`/`sync` on a `codegen_error` hold: exit 0 + banner,
  or non-zero?
- Exact volume threshold (item 3): fraction, absolute count, or both.
- How `check` (item 5) is wired into CI, and how CI scopes it to only the PRs
  that plausibly touch the engine.
- **New, from review** — the exact shape of `decide()`'s signature change for
  item 3's routing: a new `changeClass` value, or a separate parameter.
- **New, from review** — whether item 5's "fail on any output-diff" widening
  needs its own noise/volume consideration, or whether zero-diff-by-default
  for non-engine PRs already makes that moot; check for any incidental,
  codegen-unrelated byte churn that could produce a spurious diff.
