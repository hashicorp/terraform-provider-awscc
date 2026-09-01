<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Suppressed and frozen: what the knobs mean, and what they should mean

`all_schemas.hcl` has five levers for excluding a type (or one artifact of a
type) from generation, plus a sixth external one (the checkout file). None of
them currently record *why* they were pulled. This doc has two parts:

- **Part 1** describes exactly where things stand today, grounded in the real
  overlay and the real code — including what we don't know, because
  reconstructing intent after the fact is the problem this doc exists to stop.
- **Part 2** is the spec: a taxonomy for *why*, a clarified meaning for
  `frozen_since`, a `-heal` subcommand that re-probes ungrounded rows, and
  `-check` enforcement so the gap cannot silently reopen.

<!--mdtoc: begin-->
- [Part 1 — Where we are](#part-1--where-we-are)
    - [The six levers](#the-six-levers)
    - [How each lever gets set today](#how-each-lever-gets-set-today)
    - [What the real overlay looks like](#what-the-real-overlay-looks-like)
    - [What we don't know](#what-we-dont-know)
- [Part 2 — The spec](#part-2--the-spec)
    - [Suppression reason taxonomy](#suppression-reason-taxonomy)
    - [Frozen, precisely](#frozen-precisely)
    - [Per-artifact independence](#per-artifact-independence)
    - [GitHub issues: when to file, and what to say](#github-issues-when-to-file-and-what-to-say)
    - [Mining existing issues to backfill reasons](#mining-existing-issues-to-backfill-reasons)
    - [`-check`: enforce a reason](#-check-enforce-a-reason)
    - [`-heal`: re-probe and fill gaps](#-heal-re-probe-and-fill-gaps)
    - [What does not change](#what-does-not-change)
<!--mdtoc: end-->

## Part 1 — Where we are

### The six levers

| Lever | Type | Scope | Meaning today |
|---|---|---|---|
| `suppress_resource_generation` | bool | one artifact | don't generate/emit the resource |
| `suppress_singular_data_source_generation` | bool | one artifact | don't generate/emit the singular data source |
| `suppress_plural_data_source_generation` | bool | one artifact (+ gates `ListResource`, see below) | don't generate/emit the plural data source |
| `frozen_since` | date string | whole type | stop refreshing this type's schema from AWS; keep last-good bytes |
| `non_provisionable` | bool | whole type | annotation only: AWS lists the type but it cannot actually be provisioned (no generation impact) |
| `suppression_reason` | string | whole type | **HCL field exists, but is never populated** — see below |
| `internal/update/suppressions_checkout.txt` | file, external | whole type | a separate, older mechanism: pins specific `internal/service/cloudformation/schemas/AWS_*.json` files via `git checkout` so a refresh keeps last-good bytes. Orthogonal to everything above; read-only cross-reference in bigdiffer today (see `bigdiffer-design.md` §5 and "Deferred and future work" for the planned fold into `frozen_since`) |

One structural fact shapes everything else: a resource, its singular data
source, and its plural data source are generated from **the same schema file**.
There is no way for one artifact of a type to be generated from schema version 5
while another uses version 3 — freezing or advancing the schema is necessarily a
whole-type decision, even though *suppression* of an individual artifact against
whichever schema version is active is not.

The resource and the plural data source have one more coupling: generating a
`ListResource` alongside the resource is driven by whether the plural data
source is suppressed (`plan.go`: `listResource: !row.SuppressPluralDataSourceGeneration`).
A `ListResource` without a working plural data source makes no sense, so the two
travel together. The singular data source has no such coupling — it is
independent of both.

### How each lever gets set today

**The legacy AWS → available-schemas transform (`make biglister`,
`generators/allschemas/allschemas.go`).** Before generation is ever attempted,
this step sets `suppress_plural_data_source_generation` on **every** newly
discovered type, purely structurally: it calls `DescribeType`, parses the `list`
handler, and only clears the suppression if that handler exists and has zero
required arguments. This is not a failure being recorded — it is a fact about
the schema, known in advance. This is the origin of the great majority of the
428 `suppress_plural_data_source_generation = true` rows in the current overlay.

**The legacy weekly process, by hand (`generating-the-provider.md`).** Steps 2–5
of the legacy doc are the entire mechanism for everything else: run a `make`
target, and *"if there are any errors, add `suppress_resource_generation = true`
... and open an issue with details about the suppression reason."* There is:

- no structured place to put that reason (the doc says to put it in "an issue,"
  not in the HCL file),
- no distinction between *why* the target failed (a schema that can't be parsed
  vs. one that parses but won't compile vs. one that compiles but breaks the
  build),
- and no mechanism that ever revisits a suppression later to check whether the
  underlying problem was fixed.

**The automated weekly-update-PR tooling (`internal/update/makes.go`).** This
newer, still-live, still-separate package (imported by nothing in bigdiffer) is
more structured than the fully manual process, but still legacy: it greps
`make`'s output for known error shapes, classifies by which `make` target was
running (`BuildTypeSchemas` / `Resources` / `SingularDataSources` /
`PluralDataSources`), sets exactly one corresponding `suppress_*` flag — except
a `BuildTypeSingularDataSources` failure also suppresses the plural data source,
a coupling specific to the legacy generator's shared template-data plumbing that
**does not exist in bigdiffer's owned engine** (`codegen.GeneratePluralDataSource`
takes only the CloudFormation type name — no dependency on the singular data
source's generation or success). It does write a reason, automatically:
`suppression_reason = "<raw error text>:<GitHub issue URL>"`. This is the only
place in the whole system that has ever populated the `suppression_reason`
field programmatically — and even it is not run against the *current* overlay
format bigdiffer manages, since it operates through its own parallel copy of
the make-target flow.

**bigdiffer today.** `decide()` (`policy.go`) sets suppression for a **New**
type per the failed artifact (correctly per-artifact), and `gateFailureReason()`
builds a `suppression_reason` string from the generation error automatically.
This is real progress — a machine-written reason, always present for
suppressions bigdiffer itself makes going forward. The gaps this section
originally called out have since been closed: it now runs for **Present** types
on a partial failure (per-artifact suppression + freeze), the reason carries a
**taxonomy** (`category: detail`, Part 2), and the **compile gate** runs inside
`-update` so `build_failed` is a real category, not a manual `make build`
afterthought (`bigdiffer-design.md` §6). What remains is applying that machinery
to the *existing* reason-less backlog — Part 2's `-check`/`-heal` and the
one-time backfill.

### What the real overlay looks like

Counted directly against the current `internal/provider/all_schemas.hcl`
(1579 `resource_schema` blocks):

| Flag | Count |
|---|---|
| `suppress_resource_generation = true` | 47 |
| `suppress_singular_data_source_generation = true` | 46 |
| `suppress_plural_data_source_generation = true` | 428 |
| `frozen_since` set | 35 |
| `non_provisionable = true` | 3 |
| `suppression_reason` attribute set | **0** |

428 plural suppressions is consistent with the structural pre-check being the
dominant cause (most CFN types simply don't have a list handler with no
required arguments), not 428 recorded failures.

What actually carries a "reason" today is a free-form `# Suppression Reason:`
comment above the block — a convention, not a schema field, so it is invisible
to any tooling:

```hcl
  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
```

```hcl
  # suppression reason: recursive object definitions
  suppress_resource_generation             = true
```

Note the inconsistent casing and phrasing even between adjacent examples — this
is hand-written prose, not structured data.

### What we don't know

This is the part that matters most, because it is the actual problem: a
meaningful share of the existing suppressions and every existing freeze carry
**no recorded reason at all**, structured or free-form:

- **All 35 `frozen_since` rows have zero comment explaining why they were
  frozen.** We cannot currently distinguish, for any of them, "AWS broke this
  schema and nobody has revisited it" from "we are deliberately holding this
  type back for an unrelated reason" from "this type was withdrawn from AWS and
  frozen pending removal." The dates range from 2021-08-11 to 2026-08-12, so
  some are years old.
- **3 of the 47 `suppress_resource_generation = true` rows have no comment at
  all** — not even free-form prose.
- We do not know how many of the remaining suppressions with a comment are
  still accurate. A comment like "Recursive Attribute Definitions" recorded in
  2022 may or may not still be true against bigdiffer's owned engine and
  today's schema — nobody has re-checked.
- We do not know how many `suppress_plural_data_source_generation = true` rows
  are structural (no list handler — expected, not a problem) versus a genuine
  unresolved failure, because the structural pre-check's output was never
  distinguished from a failure's output in the data model. They are the same
  bit.
- We do not have a compile-gate history at all — no record of which
  suppressions, if any, exist because generated code built but failed to
  `go build`. The compile gate runs inside `-update` now
  (`bigdiffer-design.md` §6), but the existing backlog predates it, so no
  `build_failed` history exists for these old rows.

Part 2 exists to make sure this list is never this long again.

## Part 2 — The spec

### Suppression reason taxonomy

Every `suppression_reason` bigdiffer writes going forward names exactly one of
five categories, machine-readable as a prefix (`category: detail`):

| Category | Meaning | Set by |
|---|---|---|
| `structural` | The schema itself doesn't support this artifact — known before generation is attempted, not a failure. Today's only case: a plural data source with no `list` handler, or one requiring arguments. | The planner, before generation runs |
| `generation_failed` | The owned engine (`codegen.Generate*`) returned an error rendering this artifact. | `-update` / `-generate`, from the real generation error |
| `build_failed` | The artifact rendered, but `go build` failed on it (or on the package it lives in, attributable to it). | The compile gate, inside `-update` (`bigdiffer-design.md` §6) |
| `manual` | A human suppressed this for a reason not mechanically detected (e.g. a schema shape the engine accepts but that produces something wrong or undesirable). | A human, editing `all_schemas.hcl` directly |
| `unknown` | Inherited from history; no reason was recorded, and `-heal` has not yet been able to reclassify it. | `-heal`, as a placeholder, or left over from before this taxonomy existed |

`unknown` is not a permanent category — it exists so `-heal` and `-check` (next
two sections) have something to grab onto, and so that pre-existing rows are
never silently reported as reasoned when they aren't.

Every suppression/freeze bigdiffer itself sets is tagged with a category from
day one. The `manual` category is the escape hatch for the cases the mechanical
categories can't currently explain — anyone can still hand-edit
`suppression_reason` directly, exactly as today.

### Frozen, precisely

`frozen_since` pins the **schema version**, not "this type builds nothing."
Concretely:

- A frozen type's schema bytes are not refreshed from AWS. `detectChanges`
  continues to short-circuit it (`statusFrozen`), exactly as today.
- Generation still runs for a frozen type's non-suppressed artifacts, against
  the **pinned** schema bytes — not against a new AWS schema, since there is
  none in play. A type frozen at version N can have its resource and singular
  data source generating cleanly at version N while its plural data source is
  independently suppressed (say, version N's plural template broke on the
  engine and nobody has revisited it — that combination is exactly `frozen_since`
  \+ `suppress_plural_data_source_generation` + a `generation_failed` or
  `manual` reason on the plural artifact specifically).
- "Build nothing" is simply the case where every artifact also happens to be
  suppressed. It is not a distinct mode — it is what `frozen_since` plus
  `suppress_resource_generation` plus `suppress_singular_data_source_generation`
  plus `suppress_plural_data_source_generation` already expresses, with each
  suppression carrying its own reason.
- **Freezing itself gets a reason too.** `frozen_since` is accompanied by a
  `suppression_reason` explaining why the *schema* is pinned (typically
  `generation_failed` or `build_failed` against the newly discovered bytes, or
  `manual` for a deliberate hold). This is new: today `frozen_since` carries no
  reason at all (see Part 1).
- A **withdrawn** type (gone from AWS's live listing) is frozen with reason
  category `manual` and a fixed detail string (`withdrawn from AWS, pending
  major-version removal`) — this is the one case where the reason is about the
  type's existence, not a build problem.

This resolves the ambiguity in Part 1: "frozen" was never meant to mean "stop
everything," only "stop advancing the schema." Per-artifact suppression on top
of a frozen (or live) schema version is what determines what actually builds.

### Per-artifact independence

This is the behavior change tracked in `generation-punchlist.md` (implementation,
not this doc):
a resource, its singular data source, and its plural data source succeed or
fail **independently** within one generation pass against one schema version,
and the outcome for each is reflected both in the provider (only files for
artifacts that generated cleanly are written/promoted) and in `all_schemas.hcl`
(only the artifacts that failed are suppressed, each with its own
`category: detail` reason). Today, one artifact failing discards the whole
type's generation for that cycle (see the worked example in the session that
led to this doc); that is the bug this spec's taxonomy and Part 2 exist to let
us fix correctly instead of papering over.

The resource↔plural `ListResource` coupling is preserved: a resource is
generated with `GenerateListResource` only when the plural data source for the
same pass is not suppressed (whether previously or newly).

### GitHub issues: when to file, and what to say

Not every suppression deserves a GitHub issue, and bigdiffer should say which do.
The taxonomy maps directly onto the decision:

| Category | File an issue? | Why |
|---|---|---|
| `structural` | **No** | Upstream and common — the CFN schema simply has no qualifying `list` handler. Not our defect, nothing to track; filing here would be constant noise. |
| `generation_failed` | **Yes** | The owned engine couldn't render the artifact from a schema AWS considers valid — a real defect (ours or the schema's) worth tracking. Validation failures fold in here (validation is the front half of generation, `bigdiffer-design.md` §6). |
| `build_failed` | **Yes** | The artifact rendered but didn't compile — a real defect. |
| `manual` | Human's call | Whoever suppressed it knows the reason and whether it warrants an issue. |
| `unknown` | Triage | Reason lost to history; `-heal` reclassifies it first, and whether it then warrants an issue follows the category it lands on. |

For the issue-worthy categories (`generation_failed`, `build_failed`), bigdiffer
should emit, in its report, a ready-to-file **issue stub** per affected type: the
CloudFormation type name, the failed artifact(s), the category, and the captured
error text (the `detail` half of `category: detail`) — so filing is copy-paste,
not archaeology. This is the structured successor to the legacy process's "open
an issue with details" instruction (and to `internal/update/makes.go`'s
`suppression_reason = "<error>:<issue URL>"` convention), except the error is
captured automatically and the human only files it and pastes the URL back.

Once the issue exists, its URL belongs in the row. `suppression_reason` may carry
the issue URL alongside the category (`category: detail (issue: URL)`) — the
structured form of the legacy `<error>:<issue URL>` intent. bigdiffer never
invents or auto-files issues; it tells the human exactly when one is warranted
and hands them the text.

### Mining existing issues to backfill reasons

The reason-less backlog (Part 1) predates all of this, but the reasons often
already exist — in open GitHub issues filed by the legacy process, which mandated
one per suppression. Those issues carry the error detail and the type name. A
one-time (or `-heal`-assisted) pass can mine the repo's open issues — matching on
CloudFormation type name and error signatures — to *propose* `suppression_reason`
values and issue-URL links for rows that have none today.

This is best-effort and advisory, like every `-heal` proposal: a mined reason is
reported for human confirmation, never written silently. It won't cover every
row, but it can convert a meaningful share of the 35 reason-less freezes and the
bare suppressions from `unknown` into a categorized reason with an issue link,
cheaply — turning the "what we don't know" list in Part 1 back into recorded
intent.

### `-check`: enforce a reason

`-check` (offline, safe on every PR) gains one more anomaly class: **any row
with a suppressed artifact or a set `frozen_since` that has no
`suppression_reason`** is reported the same way `UnexplainedRetained` is
today — advisory, printed, not a hard failure (matching the existing "advisory,
never a hard PR gate" posture for anomalies), but now visible instead of
silent. This is what makes the 35-frozen-with-no-reason and 3-bare-suppression
situations in Part 1 impossible to reintroduce without at least a report line
calling it out.

### `-heal`: re-probe and fill gaps

A new subcommand, `-heal`, for exactly the backlog Part 1 describes. For every
row that is suppressed or frozen **and** has no `suppression_reason` (or has
one tagged `unknown`), `-heal`:

1. **Checks structural first.** For a suppressed plural data source, re-run the
   same list-handler check the legacy transform used. If it still doesn't
   qualify, tag `structural` and stop — nothing else to probe.
2. **Re-runs generation.** Using the current schema bytes (frozen types use
   their pinned bytes; live types use the current cache), call the owned
   engine for the specific suppressed artifact. If it now succeeds, propose
   lifting the suppression (report it; do not silently un-suppress without
   review — a human confirms, matching the existing "everything is a proposal,
   the human reviews the report" posture in `bigdiffer-design.md` §8). If it
   still fails, tag `generation_failed` with the current error text.
3. **Re-runs the compile gate** — the gate now exists (`bigdiffer-design.md`
   §6); wiring it into this `-heal` probe (so a suppressed artifact that
   generates but fails `go build` retags `build_failed` rather than proposing a
   `lift`) is the remaining follow-up. If generation succeeds but the build
   fails, retag `build_failed`.
4. **Falls back to `manual`/`unknown`.** If none of the above apply — most
   commonly, an existing free-form `# Suppression Reason:` comment already
   explains it in prose — `-heal` does not overwrite a human's comment; it
   leaves the row as `manual` (best-effort: if a `# Suppression Reason:`
   comment exists, migrate its text into the `manual:` detail rather than
   discarding it) and reports it as still needing a human look if the comment
   doesn't parse into something confident.

`-heal` never suppresses or un-suppresses anything on its own for a row that
already has a reason — it only targets the `unknown`/reason-less backlog, and
every proposed change is reported for human review rather than applied silently,
consistent with every other bigdiffer policy decision.

### What does not change

- The checkout file (`suppressions_checkout.txt`) remains a separate, read-only
  cross-reference, exactly as documented in `bigdiffer-design.md` §5. Folding
  it into `frozen_since` is unrelated to this spec and stays a future migration.
  `-heal` does not touch it.
- `non_provisionable` remains a bare annotation with no generation effect and no
  reason requirement (it says AWS lists the type as un-provisionable, not that
  bigdiffer failed to generate it).
- Everything here is advisory and additive: no existing suppression, freeze, or
  comment is removed or overwritten without a human accepting a proposed change
  (`-heal`'s report) or editing the file directly.
