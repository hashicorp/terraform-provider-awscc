<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# New Generation Process (Design)

This is the design for a new tool called **bigdiffer** (`internal/tools/bigdiffer`).
Its job is to regenerate the Terraform AWSCC provider from the current live
CloudFormation (CFN) types while protecting the provider from additions and
updates that would break it. It is organized and precise, communicates amply,
and treats its users' time as valuable through automation.

Status: design proposal. It defines the model, the data flow, and the decision
rules — not a specific implementation. Sections 1–8 define the target design;
section 9 records decisions already proven in the current tool; sections 10–11
map today's process onto it.

Scope: generate resources, data sources, and list resources from current CFN
schemas, plus docs. Release mechanics (tagging, changelog, PRs) are out of scope.

Naming: the legacy `make bigdiffer` target is an unrelated `git diff` of dated
`available_schemas` files. It is removed by this design; until then, "bigdiffer"
means `internal/tools/bigdiffer`.

## 1. Core insight

The provider already records what it is — `internal/provider/all_schemas.hcl`
(the "overlay") — and AWS already records what exists — the live CFN registry.
Updating the provider is just reconciling those two facts.

Today they are not reconciled directly. AWS's live list is serialized to a dated
`available_schemas.<date>.hcl`, that file is `git diff`ed against last week's
copy, and a human hand-transcribes the result into the overlay. The dated files
feed nothing but the diff. Remove that detour and the process is: read the
overlay, ask AWS what exists, reconcile, generate resiliently, write the overlay
back. No git, no dated snapshots, no manual transcription.

## 2. Model: two inputs, one output

```
  AWS registry ──▶ ┌────────────┐
  (live types)     │ Reconciler ├──▶ updated all_schemas.hcl (+ change report)
  all_schemas ───▶ └────────────┘   join on CFN type name, in memory
```

- Input A — live AWS types: `ListTypes` (public, live, `FULLY_MUTABLE` +
  `IMMUTABLE`) plus `DescribeType` for each schema. Per type this yields the
  Terraform type name and whether a plural data source is supported (a `list`
  handler with no required arguments).
- Input O — the overlay: every type the provider tracks, plus per-type policy
  (`suppress_*_generation`, `frozen_since`, `non_provisionable`,
  `suppression_reason`).
- Output — the overlay, rewritten in place, plus a change report.

The join key is the CFN type name (e.g. `AWS::S3::Bucket`); no other correlation
artifact is needed.

## 3. Change classes

The join yields three classes with mechanical default handling:

| Class | Definition | Handling |
|-------|-----------|----------|
| New | in A, not in O | probe; add to overlay |
| Present | in A and O | refresh unless frozen |
| Absent | in O, not in A | probe, then classify (below) |

Absent is not automatically "withdrawn." Because A excludes `NON_PROVISIONABLE`
types, an absent row is resolved by one `DescribeType` probe:

- Probe succeeds, `NON_PROVISIONABLE`: still live, only filtered out. Annotate
  `non_provisionable = true`; do not freeze.
- Probe returns `TypeNotFoundException` (or `DEPRECATED`): genuinely withdrawn.
  Keep cached bytes and set `frozen_since`.

The provider ships non-provisionable-but-live resources (e.g.
`AWS::AppStream::StackFleetAssociation`), which a naive "absent ⇒ freeze" rule
would wrongly freeze. Only the probe distinguishes the two.

## 4. Selective refresh, not teardown

`make schemas` downloads only schemas that are *missing*, then validates — it
never compares against or replaces an existing file. Today an update is forced
by deletion: `make cleanschemas` removes every `AWS_*.json` so the next
`make schemas` re-downloads them all. Pinning falls out of the same mechanism —
`make suppressions` git-restores frozen JSONs first, so `make schemas` finds them
present and leaves them untouched. This is a clunky stand-in for compare →
replace: delete everything to trigger a full re-download.

Replace it with a direct compare — download a type only when AWS's bytes differ
from the cache, and skip frozen types outright:

- Frozen types: not re-downloaded during the normal refresh — last-good bytes
  stay (no delete/restore). The self-heal probe (§7) may re-fetch them to a
  scratch location to test recovery, but never disturbs the pinned bytes unless
  recovery is accepted.
- Present, non-frozen: downloaded only if AWS's bytes differ; a byte-identical
  refresh is a no-op that skips the gate.
- New types: downloaded once.

So `cleanschemas` and `suppressions` cease to exist as steps.

## 5. One declarative suppression surface

The overlay is the single policy surface, per block:

- `frozen_since` — pin: do not refresh cached bytes. This subsumes
  `suppressions_checkout.txt`; the pinned set is simply "blocks with
  `frozen_since`."
- `suppress_resource_generation`, `suppress_singular_data_source_generation`,
  `suppress_plural_data_source_generation` — do not generate that artifact.
- `non_provisionable` — cannot be provisioned via Cloud Control.
- `suppression_reason` and free-form comments — human rationale, preserved
  verbatim.

Pinning and generation-suppression are orthogonal and may co-occur, but they now
live in one file with one mental model. `suppressions_checkout.txt` and
`make suppressions` are retired.

## 6. One resilient gate

Today two global gates (validate, then generate) each traverse the whole corpus
(10–15 min each) and surface one failure at a time, forcing a rerun per broken
schema. Replace them with a single pass:

1. Merge validate and generate — validation is the front half of generation.
2. Make the pass per-type and resilient — one type's failure never aborts the
   others. Each type yields `ok`, `failed-validation`, or `failed-generation`
   with its error; independent types run concurrently.

Only changed types enter the gate (a byte-identical refresh is a no-op, §4), so
in a typical week the gate runs on a handful of types, not the corpus.

The gate is the crux, not the control flow. Today's generators traverse the
whole overlay and abort on first error; they are neither per-type nor resilient.
Making generation callable per type with per-type error capture — without paying
a process-per-type cost at ~1500 types — is the single largest implementation
risk here. Everything else is bookkeeping by comparison.

Compilation is a final gate: a type can generate valid-looking code that fails to
build. A build failure feeds the same policy as a generation failure (§7) —
suppress the offending new type, or freeze a broken present one — so it too can
require touching the overlay.

## 7. Policy: results become overlay edits

Change class plus gate result determine the overlay edit:

| Change class | Gate result | Overlay action |
|--------------|-------------|----------------|
| New | ok | add plain block |
| New | failed | add block with `suppress_*_generation` + reason (backlog) |
| Present | ok | keep block; accept refreshed bytes |
| Present | failed | set `frozen_since`; keep last-good bytes; flag |
| Non-provisionable (live) | n/a | keep block; set `non_provisionable = true` |
| Withdrawn | n/a | set `frozen_since`; keep bytes |
| Frozen / suppressed | ok on re-probe | propose un-freeze / un-suppress |

Two invariants keep every row safe by default:

- Never regress a shipped resource — a broken `Present` type is frozen at its
  last-good bytes, not failed.
- Never let a broken new type block the release — a broken `New` type is added
  suppressed and becomes backlog.

The last row is self-healing: policy is recomputed every cycle, so a type AWS has
since fixed is detected and proposed for recovery automatically. For frozen types
this means re-fetching to a scratch location and re-running the gate there (§4);
the pinned bytes are replaced only once recovery is accepted.

## 8. The one human step

Everything above is mechanical. The human reviews the change report before
commit: confirm or override auto-applied freezes/suppressions, supply
`suppression_reason` text and issue links, and adjudicate flagged anomalies. The
report is machine-readable, so it also seeds release notes.

## 9. Design decisions (proven in the current tool)

These choices are already made in `internal/tools/bigdiffer` and are adopted as
design commitments:

- Verbatim preservation. Existing blocks are rewritten byte-for-byte by a
  text scanner that delimits blocks on their braces — comments, suppression
  reasons, and fully commented-out blocks survive. A `gohcl` round-trip cannot
  do this; it discards comments.
- Cross-validation. The scanner's set of live blocks is checked against a real
  HCL decode (`hclparse` + `gohcl`); a mismatch is a hard error. This guards the
  verbatim text path with a structural check and yields per-block attributes
  (`frozen_since`, `non_provisionable`).
- Naming invariant. `resource_type_name` must equal the deterministic transform
  of the CFN type name; violations are reported (the classic copy-paste error).
- Explained-retained rule. A row absent from the base is benign only when
  explained by `frozen_since`, `non_provisionable`, or a checkout pin; otherwise
  it is flagged.
- Deterministic output. Blocks are sorted by CFN type name (matching the
  generator) and the header count is updated, giving a stable, reviewable file.
- Idempotent, two check modes. Running twice changes nothing. A structural
  `-check` verifies normalization and anomalies with no AWS access (safe on
  every PR); a live drift report against the registry drives the weekly run and
  is advisory, never a hard PR gate (the overlay is expected to lag live AWS).
- Self-contained. The tool owns the overlay's shape and its own discovery rather
  than importing the legacy generator/update packages, so those can be deleted.

## 10. What is eliminated

| Current step(s) | Fate |
|-----------------|------|
| `make cleanschemas` | Removed — refresh is selective (§4). |
| `make suppressions` + `suppressions_checkout.txt` | Removed — folded into `frozen_since` (§5). |
| `make schemas` (download-if-missing) | Replaced by a compare-and-replace AWS sweep (§2, §4). |
| `make biglister` + dated `available_schemas.<date>.hcl` | Removed — reconcile live AWS against the overlay in memory (§1). |
| `make bigdiffer` (git diff of dated files) | Removed — nothing to diff (§1). |
| Manual overlay edit | Automated — the reconciler writes blocks (§3, §7). |
| Validate gate (`make schemas`, its validation half) + rerun loop | Collapsed into the resilient pass — validation is its front half (§6). |
| Generate gate — `make resources`, `make singular-data-sources`, `make plural-data-sources`, `make docs-all` — + rerun loops | Removed as the second half of the double gate; collapsed into the one resilient per-type pass (§6). |
| `make build` | Kept as the final compile gate; its failures feed the same overlay policy as §7. |

## 11. Roadmap (gaps in the current tool)

The current tool implements §2, §9, and the additive/reporting half of §3. The
remaining gaps, roughly in dependency order:

1. Gate + policy (§6–§7). The largest gap. Today the tool only adds rows and
   reports; it never runs generation nor writes `suppress_*` / `frozen_since`
   from results. This requires per-type, resilient, in-process generation (the
   §6 crux).
2. Absent-row probe (§3). The tool reports retained rows but does not
   `DescribeType`-probe them to split non-provisionable-live from withdrawn. It
   relies on pre-existing annotations to explain them.
3. Selective cache refresh (§4). The tool edits only the HCL; it does not
   download schemas, skip frozen types, or short-circuit byte-identical refreshes.
4. Reuse the discovery bytes (§2). `discover` calls `DescribeType` per type only
   for the plural-DS decision and discards the sanitized schema — the exact bytes
   the cache needs. One sweep should feed both the reconciler and the cache.
5. Retire the checkout file (§5). It is still read as an external cross-reference
   rather than migrated into `frozen_since` and deleted.
6. Drop the snapshot path (§1). The tool still supports reading dated
   `available_schemas.<date>.hcl`, and its New-vs-backlog split depends on a
   "previous" snapshot that live discovery lacks. Under `-discover`, backlog
   needs a different signal (or the distinction is dropped).
7. Self-healing re-probe (§7, last row) and a machine-readable report (§8) are
   not yet implemented; docs/build orchestration is out of the tool's current
   scope.

## 12. Thesis

Reconcile the live AWS registry directly against `all_schemas.hcl`, refresh only
what is not frozen, generate every changed type resiliently in one pass, and let
one declarative policy — keyed on change-class × gate-result — write the overlay
back, so git diffs, dated snapshots, the checkout file, the teardown, the double
gate, and manual transcription all disappear.
