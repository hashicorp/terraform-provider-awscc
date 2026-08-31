<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Generation punchlist

The single big-picture list of what is left to do on bigdiffer's generation
pipeline, with pointers to the doc that owns each item's detail. This is a
**transient tracking doc** — it is deleted once the work lands. It does not
restate designs; it links to them.

Status shorthand: **prereq** (unblocks other items) · **core** · **backlog**
(one-time catch-up on existing data) · **deferred** (real, but not now).

## Document map

Which generation docs are durable and which are scaffolding, so we know what to
keep, trim, or delete as the work settles.

| Doc | Role | Lifespan |
|---|---|---|
| `generating-the-provider-with-bigdiffer.md` | The current weekly release process (operational how-to). | **Long-term.** The canonical runbook. |
| `generating-the-provider.md` | The legacy manual `make`-target process. | **Long-term while the fallback exists;** delete with the legacy generators. |
| `new-generation.md` | The durable *design* (why reconcile-vs-live, change classes, one gate, policy) + a §11 roadmap. | **Long-term for §1–§10** (the rationale). §11 is superseded by *this* punchlist and should shrink to a pointer here. |
| `suppressed-and-frozen.md` | Both a spec *and* the source of truth for suppression/frozen nuances. | **Long-term.** The spec sections settle into "how it works" as items land; the taxonomy, the `frozen`-means-schema-pin clarification, and the issue guidance stay permanently. |
| `bigdiffer-generation.md` | Historical record of Bricks 6–11 (owning the engine) + the generator-surface investigation. | **Transient/historical.** Its durable nugget is the "Investigation findings (the generator surface)" section; fold that into `new-generation.md` or the tool README, then retire the rest once the legacy generators are deleted. |
| `generation-punchlist.md` (this) | Big-picture gap tracker. | **Transient.** Delete when the list is empty. |

## The gaps

### Prerequisites

1. **Compile gate** — bigdiffer runs `go build ./internal/...` itself and feeds
   failures into policy, instead of leaving `make build` as a separate manual
   step. *(prereq)* Unblocks the `build_failed` reason category and `-heal`
   step 3. Detail: `new-generation.md` §6 tail, §11.
2. **Absent-row `DescribeType` probe** — split a type that is gone from the live
   crawl into non-provisionable-but-live vs. genuinely withdrawn. *(prereq)*
   Unblocks the `withdrawn` freeze path. Detail: `new-generation.md` §3.

### Suppression & frozen reasons

All detail lives in `suppressed-and-frozen.md`.

3. **Reason taxonomy** — every `suppression_reason` bigdiffer writes is
   `category: detail`, one of `structural` / `generation_failed` / `build_failed`
   / `manual` / `unknown`. *(core)*
4. **Per-artifact independence** — a resource, singular DS, and plural DS
   succeed or fail independently in one pass. Requires changing **two** sites:
   `refreshCandidate` (promote per-artifact, not all-or-nothing on `gr.ok()`)
   **and** `decide()`'s `classPresent` branch (per-artifact `suppress_*` with a
   reason, not a blanket `frozen_since`). Preserve the resource↔plural
   `ListResource` coupling. *(core)*
   **Done when:** given a type where exactly one artifact (resource, singular
   DS, or plural DS) fails and the others succeed — for both a New type and a
   Present/Changed type — the successful artifacts' files are written/promoted,
   the failed artifact's file is left untouched (Present) or simply not written
   (New), and `all_schemas.hcl` ends up with `suppress_*` set only on the failed
   artifact (with a `category: detail` reason), not on the whole type. A New
   type with a resource that succeeds but whose plural DS fails is *not*
   promoted with `GenerateListResource: true`. Covered by tests for all three
   single-artifact-failure combinations, on both New and Present, plus the
   existing all-fail and all-succeed cases (never-regress unchanged).
5. **Reason on freeze** — `frozen_since` always carries a categorized
   `suppression_reason`; today it carries none. *(core)*
6. **Tag `structural` at the source** — the plural-DS structural determination
   is in `discover.go` (`!pluralSupported`), not the planner; that is where the
   `structural:` reason should be stamped. *(core)*
7. **GitHub-issue guidance** — for `generation_failed`/`build_failed` only,
   emit a ready-to-file issue stub (type, artifact, category, captured error)
   and let `suppression_reason` carry the issue URL. `structural` never warrants
   an issue (upstream, common). *(core)*

### Enforcement & backlog

8. **`-check` reason anomaly** — advisory report line for any suppressed/frozen
   row lacking a `suppression_reason`. Non-noisy only *after* item 9 backfills
   the existing rows. Detail: `suppressed-and-frozen.md`. *(core)*
9. **`-heal` subcommand** — re-probe reason-less/`unknown` rows
   (structural → regenerate → compile-gate → migrate any free-form
   `# Suppression Reason` comment into `manual:`), reporting proposals, never
   auto-applying. Detail: `suppressed-and-frozen.md`. *(core)*
10. **One-time reason backfill** — the bulk labor: tag the ~428 structural
    plural suppressions, and triage the 35 reason-less freezes + 3 bare
    suppressions. Includes **mining open GitHub issues** to propose reasons and
    URL links for rows that have none. Detail: `suppressed-and-frozen.md`
    ("Mining existing issues"). *(backlog)*

### Reporting & robustness

11. **Machine-readable report** — structured output so `-check`/`-heal`
    proposals and the issue stubs are consumable (and seed release notes).
    Detail: `new-generation.md` §8, §11. *(core)*
12. **Never-regress cross-type atomicity** — `-update` writes generated files in
    the per-candidate loop but the overlay only once after; a late error leaves
    files promoted and the overlay un-reconciled. Stage all, promote after the
    batch. *(core)*
    **Done when:** a hard I/O error (not a generation failure — those already
    stay in-band via `gr`) on candidate N of a multi-candidate `-update` run
    leaves the working tree exactly as it was before the run started — no
    candidate's files or cache are promoted, and `all_schemas.hcl` is
    unmodified — rather than candidates 1..N-1 promoted with the overlay never
    reconciled. Covered by a test that injects a failure partway through a
    multi-candidate batch and asserts nothing was promoted and the overlay file
    is byte-identical to its pre-run content.

### Deferred

13. **Checkout-file retirement** — fold `suppressions_checkout.txt` into
    `frozen_since`. Detail: `new-generation.md` §5, `bigdiffer-generation.md` §5.
    *(deferred)*
14. **Parity-validated naming simplification** — replace the `isCustomName`
    regex list with the general "plural == input ⇒ suffix" rule, proven by the
    parity harness. Detail: `bigdiffer-generation.md` ("Deferred"). *(deferred)*
15. **Delete the legacy generators, directive files, and `make` targets** — only
    after full-corpus parity has held for several real cycles. Detail:
    `new-generation.md` §10, `bigdiffer-generation.md` ("Deferred"). *(deferred)*
