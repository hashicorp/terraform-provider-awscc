<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer: splitting `suppression_reason` per artifact (design)

**Transient.** Deleted once implemented and folded into `bigdiffer-design.md`
and `suppressed-and-frozen.md`. Tracks `generation-punchlist.md` item 9b.

## The problem

`resourceRow.SuppressionReason` (`model.go`) is one string per row, but
suppression is already per-artifact: a resource, its singular data source, and
its plural data source each have their own `suppress_*_generation` boolean and
can fail (or be manually suppressed) independently of each other
(`generation-punchlist.md` item 4, "per-artifact independence"). A freeze
(`frozen_since`) is a fourth, orthogonal fact — it pins the whole *schema*, not
one artifact.

One shared reason field cannot describe a row where more than one of these facts
is true at once, and — more importantly going forward — it cannot support
lifting one artifact's suppression without disturbing another's reason. That is
a real, named goal: AWS periodically adds the plural list operation to a type
that was originally missing it, and the fix should be "clear
`suppress_plural_data_source_generation` and its own reason," not "figure out
which part of one combined sentence was about the plural DS and hope the rest
still applies to the resource."

This was surfaced while designing item 9a (`-heal` re-proposing `lift` for a
plural data source that structurally cannot fail, every run, once a human has
already decided to keep it suppressed for an unrelated reason) — 9a cannot be
expressed cleanly without this.

## Is the current data actually migrated, or greenfield?

Checked directly against the real overlay (`internal/provider/all_schemas.hcl`,
1579 rows) before designing anything, since a real migration changes what's
defensible:

- **Zero rows currently have a non-empty `suppression_reason`.** The taxonomy
  (items 3/5/6) added the capability to write one; nothing has populated it yet
  in the committed file. The backlog (item 10) is exactly the not-yet-done work
  of populating it.
- 37 rows have more than one artifact suppressed (36 with all three, 1 with
  resource+plural only) and 9 rows are both frozen and have a suppressed
  artifact — the shapes that would make a single shared reason ambiguous. All
  46 of them currently have **no reason at all** in the *structured*
  `suppression_reason` field, so there is nothing to reassign or reinterpret
  there.

**This is only true for the structured field — it does not hold for the
free-form `# Suppression Reason:` code comment, which is a second, populated,
row-level reason store.** 58 rows carry one of these comments today, and 41 of
those are multi-suppressed — exactly the "one combined reason string covering
multiple artifacts" case the structured-field check above shows does not
exist. `-heal`'s existing `commentByCFN` map is keyed by CFN type (row-level)
and is passed identically into `healArtifact` for *every* suppressed artifact
of a row (`healRow`'s loop) — so a naive per-artifact split would mint the
same, possibly-single-artifact-specific prose into two or three different
fields as if it applied to all of them equally, which the original comment
never claimed. See "`-heal`'s comment fallback" below for the resolution: this
is not auto-split, it is offered as a candidate per artifact for a human to
confirm.

Conclusion: **the structured-field side of this is greenfield — no
`suppression_reason` value needs to be split, reparsed, or reassigned.** The
comment-store side is not; it is populated and does have the multi-artifact
ambiguity, and needs an explicit routing answer (below), not a mechanical
guess.

## The design

### Schema (`model.go`)

Replace the one `SuppressionReason` field with four, mirroring the existing
one-attribute-per-artifact shape the boolean suppression flags already
establish, rather than inventing a new structural idiom (a nested HCL block, a
delimited value, or folding the reason into the boolean itself) for a single
field:

```go
SuppressionReasonResource            string `hcl:"suppression_reason_resource,optional"`
SuppressionReasonSingularDataSource  string `hcl:"suppression_reason_singular_data_source,optional"`
SuppressionReasonPluralDataSource    string `hcl:"suppression_reason_plural_data_source,optional"`
FrozenReason                         string `hcl:"frozen_reason,optional"`
```

Each of the first three lines up 1:1 with its existing boolean
(`suppress_resource_generation`, `suppress_singular_data_source_generation`,
`suppress_plural_data_source_generation`). `frozen_reason` lines up with
`frozen_since` the same way — a freeze is a schema-level fact independent of
which (if any) artifacts are also suppressed, so it gets its own field rather
than being folded into one of the three or left ambiguous when a row is both
frozen and has a suppressed artifact (9 real rows today).

`SuppressionReason` (the old field) is deleted, not deprecated-and-kept — there
is no data in it to preserve (see above), and keeping a dead field around
invites exactly the "which one is authoritative" confusion this change exists
to remove.

### Writers: `decide()` and `gateFailureReason` (`policy.go`)

`gateFailureReason(gr gateResult) string` today builds one string by iterating
`gr.artifacts`, labeling each failure (`"%s: %s"`, already artifact-labeled
internally) and joining with `"; "`, collapsing information it already has in
per-artifact form. Replace it with a function returning `map[string]string`
(not `map[artifactKind]string` — it must be keyed by the same attribute-name
constants `suppressAttrsForFailures` uses, `attrSuppressResource` etc., so the
two zip together directly when `decide()` builds `policyDecision.setAttrs`; the
old draft of this doc said both, inconsistently — this is the pinned answer),
one formatted `category: detail` reason per failed artifact, reusing
`formatReason` unchanged.

**The `len(attrs) == 0` fallback in `suppressAttrsForFailures` must be mirrored
exactly, not left to only the failure loop.** Today, when `gr.artifacts`
contains no attributable failure (a total plan-derivation failure with no
per-artifact detail), `suppressAttrsForFailures` still defensively suppresses
the resource (`attrs[attrSuppressResource] = "true"`) so the row is never added
un-suppressed. A reason-map builder that only loops `gr.artifacts` would miss
this case entirely, leaving `attrSuppressResource` set with no matching reason
entry — which `-check`'s new per-field anomaly loop (below) would immediately
flag as reason-less, a regression introduced by this change rather than fixed
by it. The reason-map builder must check the identical `len(...) == 0`
condition and, in that case, also emit
`attrs[attrSuppressResource] = formatReason(reasonGenerationFailed, "no
attributable per-artifact failure; suppressed resource generation
defensively")` (or equivalent), so the two functions' fallback branches never
drift apart. Consider deriving both `setAttrs` and the reason map from one
shared internal loop (e.g. a single function returning both maps together)
specifically so this fallback can't be updated in one place and forgotten in
the other.

`policyDecision.reason string` (one field) becomes `policyDecision.reasons
map[string]string`, same attribute-name keys, covering all four possible
fields uniformly (the three per-artifact ones plus `"frozen_reason"`).

**`frozen_reason` needs its own rationale text, never the reused
per-artifact-failure string.** `classPresent`'s total-failure branch and
`classWithdrawn` already compute one reason for the freeze alone (no
per-artifact failures exist in those branches) — they write it under
`"frozen_reason"` unchanged in substance, just renamed from `"suppression_reason"`.
`classPresent`'s **partial**-failure branch is the one that needs care: the
freeze there exists for a different, structural reason than any single
artifact's failure — "the schema is pinned because one JSON backs all three
artifacts and cannot partially advance," a fact that is true regardless of
*which* artifact(s) failed. Reusing `gateFailureReason`'s (renamed) output for
`frozen_reason` in this branch would misdescribe the freeze as caused by the
specific artifact error, when the freeze is actually caused by the "can't
partially advance one shared JSON" invariant itself. This branch mints a fixed
*rationale sentence* for `frozen_reason`, but **the category prefix must still
be derived from what actually triggered the freeze, not hardcoded** — a
partial failure can be triggered by a build failure just as easily as a
generation failure (the compile gate is a first-class stage; `gr.artifacts` can
mix `gateFailedGeneration` and `gateFailedBuild` outcomes freely, `gate.go`'s
`anyOK`/`ok` don't distinguish), so hardcoding `reasonGenerationFailed` here
would mislabel a build-triggered partial freeze with a `generation_failed:`
prefix — the same category-mislabel class as the `-heal` bug fixed earlier
this session. Reuse the exact category-selection rule `gateFailureReason`
already applies (favor `reasonBuildFailed` if any failing artifact's outcome is
`gateFailedBuild`, else `reasonGenerationFailed`) rather than inventing a
second convention:

```go
category := reasonGenerationFailed
for _, a := range gr.artifacts {
	if a.outcome == gateFailedBuild {
		category = reasonBuildFailed
	}
}
attrs["frozen_reason"] = formatReason(category,
	"schema pinned — one JSON backs all three artifacts, cannot partially advance")
```

This is distinct from, and written in addition to, the per-artifact reason(s)
for whichever artifact(s) actually failed.

### Writers: `main.go`'s HCL-writing path

`setBlockAttributes` (`mutate.go`) already takes an arbitrary
`map[string]string` — no change needed there, it is already attribute-name
agnostic. `boolAttrs` (which attributes render as bare `true`/`false` vs.
quoted strings) is unaffected since all four reason fields are strings, same as
today's single field. `canonicalBlock` and any other call site currently
building a `map[string]string{... "suppression_reason": x ...}` by hand switches
to whichever of the four attribute-name constants applies.

### `-check`'s reason anomaly (item 8, `main.go`)

`isSuppressedOrFrozen` is unchanged (still "is any of the four suppression/
freeze facts true"). The anomaly loop becomes per-fact instead of per-row: for
each row, check independently whether `suppress_resource_generation` is true
with an empty `suppression_reason_resource`, same for singular/plural, and
whether `frozen_since` is set with an empty `frozen_reason` — reporting each
missing reason as its own advisory anomaly line (naming which specific field is
missing) rather than one bit per row. This is a strictly more precise version
of the same check, not a new one.

### `-heal` (item 9, `heal.go`)

`runHeal`'s row-level gate today is "skip the whole row if
`row.SuppressionReason != "" && not unknown:`". This becomes per-artifact: for
each of the three suppressed-artifact checks and the freeze check
independently, skip probing that *specific* fact if its own reason field is
already set to something real. This is what actually fixes 9a's recurrence
(next section) — a row can have a real reason recorded for its resource while
still being probed and reported for its still-reason-less plural DS, and vice
versa, instead of the whole row being skipped or included as one unit.

`healArtifact`/`healRow`/`freezeProposal`'s return type keeps carrying `kind`
(already does, `""` for the freeze) — proposals stay one-per-artifact-or-freeze,
already the right shape; only the *reading and writing* of the reason changes
from one shared field to the matching one of four.

### `-heal`'s comment fallback: propose, don't auto-split

`commentOrUnknown`'s free-form-comment path is where the populated,
multi-suppressed comment store (58 rows, 41 multi-suppressed — confirmed above)
actually needs a routing answer, since the structured field has none. The same
row-level comment text is available to every one of a row's still-reason-less
suppressed artifacts; auto-assigning it identically to each would mint
identical prose into multiple fields as if it applied to all of them equally,
which is not a fact the original comment establishes — a single "Recursive
Attribute Definitions" comment written when only the resource existed does not
necessarily explain why the plural DS was *also* suppressed later.

Consistent with this doc's own framing (comments are legacy debt being
migrated away from, not a mechanism to formalize further) and with `-heal`'s
existing "everything is a proposal" posture: **`-heal` does not auto-split the
comment.** For a multi-suppressed row with a row-level comment, `-heal` offers
the same comment text as a *candidate* `manual:` reason for each still
reason-less suppressed artifact independently (one proposal per artifact, same
candidate text, clearly labeled as a shared/unconfirmed candidate rather than a
confirmed per-artifact fact) and lets the human assign, edit, or reject it per
field. No mechanical guess about which artifact(s) the original prose actually
described. This also directly feeds item 10 (the backfill), which inherits the
same 41-row ambiguity and the same resolution.

### Reconciling item 9a in light of this

9a's actual bug (a structurally-cannot-fail plural DS re-proposing `lift`
forever once probed) is fixed as a natural consequence of the per-artifact
`-heal` gate above: once a human sets `suppression_reason_plural_data_source`
to anything real (a `manual:` reason, e.g. "kept suppressed: semantically
redundant with X"), `-heal` stops probing that artifact specifically on every
future run, regardless of what the resource's or singular DS's own reasons say
or whether they are still reason-less. No separate mechanism, adjudication
marker, or new taxonomy sub-tag is needed — 9b's schema change is 9a's fix.

The one thing worth adding on top per the review discussion: a `lift`
proposal's reason text should say which exact field to set to make the proposal
stop recurring, e.g. for a plural DS lift proposal:

> `generates cleanly now — propose lifting the suppression. To keep it
> suppressed and stop this proposal recurring, set
> suppression_reason_plural_data_source instead (e.g. "manual: ...").`

This keeps `-heal` stateless and puts the decision explicitly in the human's
hands, consistent with the "everything is a proposal, the human decides"
posture (`bigdiffer-design.md` §8) — no persisted adjudication state, no new
reason category.

### Tests and real-tree impact

No `all_schemas.hcl` data migration script is needed for the structured field
(confirmed above: nothing in `suppression_reason` to migrate). The free-form
comment store does need new *behavior* (the propose-per-artifact-candidate
routing above), but not a migration script either — it stays exactly what it
is today (unstructured prose `-heal` reads and proposes from) and is written
to the overlay, if at all, only by a human confirming a proposal, same as every
other `-heal` outcome. Existing tests referencing `SuppressionReason`
(`main_test.go`, `update_test.go`, `compile_test.go`, `heal_test.go`) update to
whichever of the four fields applies to what they are asserting — this is a
mechanical rename at each call site once the four-field shape is fixed, not a
logic change to those tests' intent; a new test is needed specifically for the
comment-routing behavior (a multi-suppressed row with a comment proposes the
same candidate text for each still-reason-less artifact, not a merged or
arbitrarily-assigned one).

The real overlay itself does not need editing as part of landing 9b — every
row's suppression/freeze booleans are untouched; only the (currently always
empty) structured reason field's name(s) change. `TestFullCorpusParity` is
unaffected (reason text is not part of generated code).

## What this does not do

- Does not populate any reasons — that is item 10 (the backfill), which lands
  after this, now targeting four narrower fields instead of one ambiguous one,
  and inherits the comment-routing resolution above for its 41 ambiguous rows.
- Does not change the three suppression booleans or `frozen_since` themselves.
- Does not add a machine-readable report format — that is item 11, sequenced
  after 9b/9a so its schema can represent four reason fields instead of one
  from the start rather than needing a second revision.
- Does not formalize the free-form `# Suppression Reason:` code-comment
  convention into a structured input beyond reading and proposing from it, as
  `-heal` already does — it remains legacy debt being migrated away from.

## Sequencing

9b (this) → 9a (the recurring-lift fix, now just the per-artifact `-heal` gate
plus the actionable proposal text) → 11 (machine-readable report, designed
against the four-field shape) → 10 (the backfill, executed against the final
shape and consumable structured output).
