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
| `bigdiffer-design.md` | The durable design reference — what bigdiffer is, how it works, and why (model, change classes, the two gates, policy, the generator surface, deferred work). | **Long-term.** The maintenance/review reference. |
| `generating-the-provider-with-bigdiffer.md` | The current weekly release process (operational how-to). | **Long-term.** The canonical runbook. |
| `generating-the-provider.md` | The legacy manual `make`-target process. | **Long-term while the fallback exists;** delete with the legacy generators. |
| `suppressed-and-frozen.md` | Both a spec *and* the source of truth for suppression/frozen nuances. | **Long-term.** The spec sections settle into "how it works" as items land; the taxonomy, the `frozen`-means-schema-pin clarification, and the issue guidance stay permanently. |
| `generation-punchlist.md` (this) | Big-picture gap tracker. | **Transient.** Delete when the list is empty. |

## The gaps

### Prerequisites

1. ~~**Compile gate**~~ — ✅ **Done** (`d2d1a1a85` + follow-ups `e201fd7d8`,
   `5055628fa`; design: `bigdiffer-design.md` §6, "The compile gate"). `-update` now builds the exact
   code it is about to promote — every staged artifact plus the registration
   file it implies — against the real module (`go build ./...` from the repo
   root, once per fixpoint round) before writing anything real, instead of
   leaving `make build` as a separate manual step run after the fact. A build
   failure is attributed back to the specific candidate/artifact that produced
   the rejected file (parsing `go build`'s own `file:line:col` output),
   downgrades only that artifact to a new `gateFailedBuild` outcome (tagged
   `build_failed`, distinct from `generation_failed`), and the fixpoint
   re-renders the registration file and rebuilds until clean — never blocking
   the release, matching every other failure mode's safe-by-default policy. A
   build error that cannot be traced to anything staged, or a small round cap,
   promotes nothing and hard-stops rather than guessing (the design doc's
   conservative "Attribution fallback"). Unblocks the `build_failed` reason
   category (previously defined but dead) and `-heal` step 3. *(prereq)*
2. **Absent-row `DescribeType` probe** — split a type that is gone from the live
   crawl into non-provisionable-but-live vs. genuinely withdrawn. *(prereq)*
   Unblocks the `withdrawn` freeze path. Detail: `bigdiffer-design.md` §3 and
   "Deferred and future work".

### Suppression & frozen reasons

All detail lives in `suppressed-and-frozen.md`.

3. ~~**Reason taxonomy**~~ — ✅ **Done** (commit `a08072eb6`, alongside item 4).
   `policy.go`'s `reasonCategory` (`structural` / `generation_failed` /
   `build_failed` / `manual` / `unknown`) + `formatReason` (`category: detail`).
   `build_failed` became reachable once item 1 (the compile gate) landed.
   *(core)*
4. ~~**Per-artifact independence**~~ — ✅ **Done** (commit `a08072eb6`). A
   resource, singular DS, and plural DS now succeed or fail independently in one
   pass: `refreshCandidate` stages each successful artifact and skips failed
   ones, `decide()`'s `classPresent` branch suppresses only the failed
   artifact(s) (keeping `frozen_since`, since one JSON backs all three), and
   `reconcileListResource` preserves the resource↔plural `ListResource`
   coupling. Tested across all three single-artifact-failure combinations on
   both New and Present. *(core)*
5. ~~**Reason on freeze**~~ — ✅ **Done** (commit `a08072eb6`, alongside item 4).
   Every `frozen_since`-setting branch in `decide()` (`classPresent` total
   failure, `classPresent` partial failure, `classWithdrawn`) sets `reason` —
   `generation_failed`-tagged for the two `classPresent` cases,
   `manual: withdrawn from AWS, pending major-version removal` for
   `classWithdrawn`. *(core)*
6. ~~**Tag `structural` at the source**~~ — ✅ **Done** (commit `a08072eb6`,
   alongside item 4). `canonicalBlock` (`main.go`) stamps
   `structural: no list handler with zero required arguments` whenever it
   writes the plural-DS structural suppression the discovery crawl determined,
   instead of a bare, unexplained flag. *(core)*
7. **GitHub-issue guidance** — for `generation_failed`/`build_failed` only,
   emit a ready-to-file issue stub (type, artifact, category, captured error)
   and let `suppression_reason` carry the issue URL. `structural` never warrants
   an issue (upstream, common). *(core)*

### Enforcement & backlog

8. ~~**`-check` reason anomaly**~~ — ✅ **Done** (`ad401cd95`). Added
   `Report.ReasonlessSuppressed`: any live row with a suppressed artifact or a
   set `frozen_since` and no `suppression_reason` is printed as an advisory
   anomaly (never a `-check` failure, matching `UnexplainedRetained`). Fixed
   `runCheck`'s success message in the same pass — it previously claimed
   "anomaly-free" unconditionally even when advisory anomalies (this one or
   `UnexplainedRetained`) were present; it now reports the advisory count.
   Verified against the real overlay: 463 reason-less rows correctly flagged,
   `-check` still exits 0, overlay untouched. *(core)*
9. ~~**`-heal` subcommand**~~ — ✅ **Done** (`04b6a264d`; compile-gate wiring
   added later, see below). `internal/tools/bigdiffer/heal.go`:
   for every reason-less/`unknown` row, probes each suppressed artifact
   (structural check via the existing `pluralSupported`, then a real
   regeneration attempt, then a real compile-gate check) and the freeze
   itself, proposing `lift` (generates cleanly *and* passes the compile gate)
   or a categorized `reason`, falling back to migrating a
   free-form `# Suppression Reason:` comment into `manual:` or `unknown` if
   nothing else applies. Never writes `all_schemas.hcl`.
   **Load-bearing fix found via a live run against the real overlay:** some
   rows are suppressed *because* their schema is recursive
   (`# Suppression Reason: Recursive Attribute Definitions`, issue #95), and
   `codegen.Emitter` has no recursion-depth guard — re-probing one in-process
   is a confirmed, unrecoverable Go stack overflow that would crash the whole
   `-heal` run. Fixed by isolating every regeneration probe in a subprocess
   (a hidden `-heal-probe-artifact` mode `-heal` re-execs itself into) under a
   30s timeout and a 512MiB `GOMEMLIMIT`, so a crash or runaway probe kills
   only that subprocess — the same "one failure never blocks the rest"
   principle `discover.go` and `generateCorpus` already apply. Verified end to
   end against the real overlay: all 463 reason-less rows processed to
   completion (including the two confirmed-recursive `AWS::WAFv2::*` types,
   which now report a normal `generation_failed` proposal instead of crashing),
   549 proposals (62 `lift`, 487 `reason`), overlay untouched.
   `TestProbeArtifactIsolatesRecursiveSchema` reproduces the crash against a
   real built binary and asserts it is contained (intentionally the one slow
   bigdiffer test — skipped under `-short`). *(core)*
   **Compile-gate wiring (follow-up once item 1 landed):** the isolated probe
   subprocess (`runHealProbeArtifact`) now also runs a successfully-generated
   artifact through the compile gate (`buildOnce`) before reporting success,
   so a `lift` proposal means "passes both stages a real `-update` run would
   enforce," not just generation. Reused `buildOnce` directly rather than the
   full fixpoint — a single-artifact probe has no registration file or batch
   to reconcile. Tested against a real, working type (still passes) and a
   real generates-but-rejected-by-the-gate case (a deliberately broken sibling
   file in the probed artifact's real destination package — no schema needs
   to be crafted to make the *engine* render invalid code, only the
   compiler needs a reason to reject the package).
9b. ~~**Split `suppression_reason` into one field per artifact, plus a
    separate freeze reason**~~ — ✅ **Done.**
    `resourceRow.SuppressionReason` was one string per row, but suppression is
    already per-artifact (item 4) and a freeze is an orthogonal, schema-level
    fact; one shared field could not describe a row with more than one of
    these true at once (37 real rows at the time) and could not support
    lifting one artifact's suppression without disturbing another's reason —
    a named goal (AWS adds the plural list operation to a type later; lifting
    `suppress_plural_data_source_generation` should not require touching the
    resource's unrelated reason). Confirmed against the real overlay before
    designing: zero rows had a non-empty `suppression_reason` at the time, so
    this was a greenfield schema change, not a migration. Replaced with
    `SuppressionReasonResource` / `SuppressionReasonSingularDataSource` /
    `SuppressionReasonPluralDataSource` (`model.go`) plus a separate
    `FrozenReason`, mirroring the existing one-flag-per-artifact shape of the
    `suppress_*_generation` booleans. `policy.go`'s `decide()`/
    `reasonsForFailures`/`frozenReasonForFailures`/`totalFailureReason` build a
    `map[string]string` keyed by attribute name (zipping directly against
    `suppressAttrsForFailures`, including mirroring its `len==0` defensive
    fallback so a fallback-suppressed resource still gets a matching reason);
    the partial-failure branch mints its own `frozen_reason` rationale with a
    category *derived* from whichever artifact(s) actually triggered the
    freeze (build vs. generation), not hardcoded — a review finding caught
    before this landed. `-check`'s anomaly loop and `-heal`'s gate both became
    per-fact instead of per-row as a direct consequence (folding in item 9a's
    fix, see below). Detail folded into `bigdiffer-design.md` and
    `suppressed-and-frozen.md`; the transient design doc
    (`suppression-reason-split-design.md`) is deleted per its stated
    lifecycle. Surfaced while scoping item 9a below, which could not be
    expressed cleanly without this. *(prereq — unblocked 9a, 11, 10)*
9a. ~~**Suppress recurring `lift` proposals for known-manual
    suppressions**~~ — ✅ **Done, as a consequence of 9b.**
    a plural data source manually suppressed for a semantic reason (not a
    generation failure) always re-proposes `lift` on every `-heal` run,
    because `codegen.GeneratePluralDataSource` builds from the CFN type name
    alone and structurally cannot fail (noted during item 9's review). Fixed
    with no separate mechanism: `-heal`'s gate is now per-artifact
    (`healFactsFor`/`healFact.needsHealing`) — it skips probing the plural DS
    specifically once `suppression_reason_plural_data_source` is set to
    anything real, independent of whether the resource or singular DS are
    still reason-less — no new adjudication marker or taxonomy sub-tag needed.
    On top of that, a `lift` proposal's text names the exact field to set to
    stop it recurring, keeping `-heal` stateless and putting the decision
    explicitly in the human's hands. Also resolved a related ambiguity found
    in review: the free-form `# Suppression Reason:` comment is row-level and
    populated (58 rows, 41 multi-suppressed) where the structured field was
    not — `-heal` offers it as a per-artifact *candidate*, visibly marked as
    shared/unconfirmed, rather than auto-splitting it identically into every
    still-reason-less fact's field (`commentOrUnknown`'s `multiPending`
    parameter). Detail folded into `suppressed-and-frozen.md`'s `-heal`
    section.
10. **One-time reason backfill** — the bulk labor: tag the ~428 structural
    plural suppressions, and triage the 35 reason-less freezes + 3 bare
    suppressions. Includes **mining open GitHub issues** to propose reasons and
    URL links for rows that have none. Detail: `suppressed-and-frozen.md`
    ("Mining existing issues"). Sequenced after 11 so the backfill consumes
    structured output instead of text; 9b/9a (the four-field shape it
    targets) are done. *(backlog, depends on 11)*

### Reporting & robustness

11. **Machine-readable report** — structured output so `-check`/`-heal`
    proposals and the issue stubs are consumable (and seed release notes).
    9b is done, so this can now be designed against the final four-field
    reason shape from the start rather than needing a second revision.
    Detail: `bigdiffer-design.md` §8 and "Deferred and future work". *(core)*
12. ~~**Never-regress cross-type atomicity**~~ — ✅ **Done** (commit
    `a08072eb6`). `refreshCandidate` now stages every artifact and cached schema
    under a temp `stagingDir`, writing nothing to the real tree; `promoteStaged`
    copies staging → real tree in one pass *after* the whole candidate batch has
    staged successfully, so a hard error anywhere in the per-candidate loop
    returns before any promotion and leaves the tree and `all_schemas.hcl`
    exactly as they started. `TestUpdateBatchAtomicity` injects a mid-batch
    failure and asserts nothing was promoted. *(core)*

### Deferred

13. **Checkout-file retirement** — fold `suppressions_checkout.txt` into
    `frozen_since`. Detail: `bigdiffer-design.md` §5 and "Deferred and future
    work". *(deferred)*
14. **Parity-validated naming simplification** — replace the `isCustomName`
    regex list with the general "plural == input ⇒ suffix" rule, proven by the
    parity harness. Detail: `bigdiffer-design.md` "Deferred and future work".
    *(deferred)*
15. **Delete the legacy generators, directive files, and `make` targets** — only
    after full-corpus parity has held for several real cycles. Detail:
    `bigdiffer-design.md` §10 and "Deferred and future work". *(deferred)*
