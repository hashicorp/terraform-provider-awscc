<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->
<!-- markdownlint-disable MD029 -->

# bigdiffer punchlist: adoption & follow-ups

The single big-picture list of what is left to do on bigdiffer, with pointers to
the doc that owns each item's detail. This is a **transient tracking doc** — it
is deleted once the work lands. It does not restate designs; it links to them.

The generation-pipeline build-out (the original punchlist, items 1–12/16) is
**complete**; it is summarized under "Completed" below, with full detail in git
history and the durable docs. What remains is **adoption** (switching the weekly
process over to bigdiffer and retiring the legacy generators) plus a few
follow-ups the build-out surfaced.

Status shorthand: **prereq** (unblocks other items) · **core** · **backlog**
(one-time catch-up on existing data) · **deferred** (real, but not now) · **done**.

## Document map

| Doc | Role | Lifespan |
|---|---|---|
| `bigdiffer-design.md` | The durable design reference (model, change classes, the two gates, policy, the generator surface, deferred work). | **Long-term.** The maintenance/review reference. |
| `generating-the-provider-with-bigdiffer.md` | The current weekly release process (operational how-to). | **Long-term.** The canonical runbook. |
| `generating-the-provider.md` | The legacy manual `make`-target process. | **Long-term while the fallback exists;** delete with the legacy generators (item D). |
| `suppressed-and-frozen.md` | Both a spec *and* the source of truth for suppression/frozen nuances. | **Long-term.** The taxonomy, the `frozen`-means-schema-pin clarification, and the issue guidance stay permanently. |
| `generation-punchlist.md` (this) | Big-picture gap tracker. | **Transient.** Delete when the list is empty (item I). |

## Completed — generation pipeline build-out

All landed; detail in git history + the durable docs above.

- **1 Compile gate**, **2 Absent-row `DescribeType` probe** — prereqs (`bigdiffer-design.md` §6, §3).
- **3 Reason taxonomy**, **4 Per-artifact independence**, **5 Reason on freeze**, **6 Tag `structural` at source**, **7 GitHub-issue guidance** — the reason machinery (`a08072eb6`; `suppressed-and-frozen.md`).
- **9 `-heal` subcommand** (`04b6a264d`, + compile-gate wiring + subprocess-isolated probes), **9a** recurring-`lift` suppression, **9b** per-artifact reason split — all done.
- **8 `-check` reason anomaly** (`ad401cd95`) — landed as advisory; **now upgraded to a hard `-check` failure in #3323** (item A) once the backlog hit zero.
- **10 One-time reason backfill** — ✅ **done in #3323**: reason-less facts 558 → 0 across 8 data-only commits (structural plural, Tier A/B, deferred-lift candidates, mined + issue-verified freezes). `suppressed-and-frozen.md` ("Mining existing issues").
- **11 Machine-readable report** — de-scoped/closed (no deterministic consumer).
- **12 Never-regress cross-type atomicity** (`a08072eb6`), **16 CHANGELOG automation** — done.

## Current work — adoption & follow-ups

### Enforcement

A. ~~**`-check` reason enforcement**~~ — ✅ **Done (#3323).** Now that the item-10
   backfill drove reason-less facts to zero, a suppressed artifact or a set
   `frozen_since` with an empty reason field is a **hard `-check` failure**
   (non-zero exit), not an advisory line — the invariant is enforced so the gap
   cannot silently reopen. `UnexplainedRetained` stays advisory (separate
   workstream). Detail: `suppressed-and-frozen.md` "`-check`: enforce a reason";
   `internal/tools/bigdiffer` README "Anomaly reports". *(core)*

### Adoption (the switchover)

C. **Adopt — phase 1: switch to bigdiffer's generated init approach, keep the
   legacy fallback.** Make the weekly/release process generate the provider via
   bigdiffer, but **leave the legacy `make` targets and generators in place**
   behind a documented "safety hatch": a short section in the runbook on how to
   revert to the legacy process if bigdiffer output regresses. No legacy code is
   deleted in this phase. Pairs with item E (release docs). Branch:
   `b-bigdiffer-adopt`. Depends on A. *(core)*
D. **Adopt — phase 2: delete the redundant legacy machinery.** After phase 1 has
   held for several real cycles, remove the legacy generators, directive files,
   and `make` targets now redundant with bigdiffer, and land the deferred design
   items that only make sense once the legacy path is gone:

- **13 Checkout-file retirement** — fold `suppressions_checkout.txt` into
  `frozen_since` (`bigdiffer-design.md` §5).
- **14 Parity-validated naming simplification** — replace the `isCustomName`
  regex list with the general "plural == input ⇒ suffix" rule, proven by the
  parity harness (`bigdiffer-design.md` "Deferred and future work"). Related
  to item B.
- **15 Delete legacy generators/directives/`make` targets**
  (`bigdiffer-design.md` §10). Delete `generating-the-provider.md` with them.

Branch: later (e.g. `b-bigdiffer-legacy-removal`). Highest risk → last.
Depends on C. *(deferred)*

E. **Release-process docs → bigdiffer (Jira).** Point the release runbook in Jira
   at the bigdiffer process. Owner: **Dirk** (external to this repo). Sequence
   alongside phase 1 (C), once bigdiffer is the default path. *(core)*

### Generator fix

B. **Pluralizer fix for already-plural type names.** `AWS::SSMGuiConnect::Preferences`
   is suppressed on its plural DS with a `build_failed` reason: the type name is
   already plural, so the singular and plural data-source generators emit the
   same Go identifier (`preferencesDataSource redeclared`). Fix the pluralizer to
   disambiguate this **class** (already-plural type names) rather than adding a
   one-at-a-time exception list. When it lands, un-suppress the plural DS and drop
   its `build_failed` reason from `all_schemas.hcl`. Overlaps with item 14's
   naming simplification — do together or sequence B before 14. Standalone
   generator PR. *(core)*

### Thaw / lift (each carries generated-code diffs — its own reviewed PR, never a data-only ride-along)

F. **Move dedup "tags" out of the schema bytes into a new `all_schemas.hcl`
   argument.** Today the deduplication marker lives inside the pinned schema JSON;
   make it a first-class overlay argument (a new `resourceRow` field + generator
   support) so it is visible and diffable in the overlay rather than buried in
   bytes. Enabler for G. Overlay-model PR. *(core)*
G. **Thaw resources unlocked by the new deduplication schema approach.**
   Re-evaluate freezes that only existed to avoid the old dedup behavior; thaw
   where the new approach makes them generate cleanly. Depends on F. *(backlog)*
H. **Thaw resources that now compile fine.** The 63 deferred lift-candidates
   recorded in #3323 (`manual: lift candidate; generates and compiles cleanly …
   pending a batched lift`) plus any freezes whose pinned bytes now generate.
   Confirm with `-heal`, lift in batches. Sequence after C so the mechanism is
   authoritative. *(backlog)*

### Cleanup

I. **Retire this punchlist.** Delete `generation-punchlist.md` once A–H (and the
   deferred items folded into D) have landed. *(deferred)*
