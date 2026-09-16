// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"sort"
	"strings"
)

// Overlay attribute names the policy sets on a block.
const (
	attrFrozenSince               = "frozen_since"
	attrFrozenReason              = "frozen_reason"
	attrNonProvisionable          = "non_provisionable"
	attrSuppressResource          = "suppress_resource_generation"
	attrSuppressSingular          = "suppress_singular_data_source_generation"
	attrSuppressPlural            = "suppress_plural_data_source_generation"
	attrSuppressionReasonResource = "suppression_reason_resource"
	attrSuppressionReasonSingular = "suppression_reason_singular_data_source"
	attrSuppressionReasonPlural   = "suppression_reason_plural_data_source"
)

// reasonCategory is the taxonomy every suppression_reason bigdiffer writes is
// tagged with, one of (contributing/docs/suppressed-and-frozen.md):
//   - structural: the schema doesn't support this artifact; known before
//     generation is attempted, not a failure.
//   - generationFailed: the owned engine returned an error rendering it.
//   - buildFailed: it rendered, but `go build` failed on it.
//   - manual: a human suppressed it for a reason not mechanically detected.
//   - unknown: inherited from history; no reason was ever recorded.
type reasonCategory string

const (
	reasonStructural       reasonCategory = "structural"
	reasonGenerationFailed reasonCategory = "generation_failed"
	reasonBuildFailed      reasonCategory = "build_failed"
	reasonManual           reasonCategory = "manual"
	reasonUnknown          reasonCategory = "unknown"
)

// formatReason renders a suppression_reason value as "category: detail", the
// taxonomy's machine-readable form. An empty detail yields a bare "category:".
func formatReason(category reasonCategory, detail string) string {
	if detail == "" {
		return string(category) + ":"
	}
	return string(category) + ": " + detail
}

// changeClass is the reconciliation class of a type (design §3). "Absent" splits
// into non-provisionable-live vs withdrawn after the DescribeType probe.
//
// classPresentUnchanged is classPresent's source-routed sibling
// (contributing/docs/held-artifacts-design.md §3, "The routing this still
// requires"): a failure on this class means the schema itself did not change
// — only the machinery (codegen/templates/toolchain) could be at fault — so
// it must never take classPresent's freeze/suppress-and-pin branch (pinning a
// schema that never moved makes no sense, and previously masked exactly this
// bug). decide() itself does not fail the run for this class; that decision
// belongs to the caller (settleBatch/runSync), which aborts the whole batch,
// promoting nothing, the moment any classPresentUnchanged candidate fails —
// see §3's "Why all-or-nothing, precisely." classPresentUnchanged exists so
// two different callers can each get the routing they need:
//   - reconcile/check have no discovery diff at all (they never call
//     discover()), so every candidate they feed decide() is
//     classPresentUnchanged, never classPresent — by construction, every
//     failure they produce is a machinery failure, never a schema one.
//   - sync, which does have a discovery diff, uses classPresentUnchanged for
//     exactly the statusUnchanged rows once it gates the whole corpus: a
//     statusChanged row (the schema itself moved) still uses classPresent and
//     takes today's freeze/suppress path unchanged.
type changeClass string

const (
	classNew              changeClass = "new"
	classPresent          changeClass = "present"
	classPresentUnchanged changeClass = "present_unchanged"
	classNonProvisionable changeClass = "non_provisionable"
	classWithdrawn        changeClass = "withdrawn"
)

// policyDecision is the overlay edit for one type: whether to add a new block,
// which attributes to set on it, any suppression/frozen reason text, and a
// one-line report summary. Per the design invariants, neither gate ever
// blocks a release on its own — except machineryFailure (below), the one
// case that legitimately must, because bigdiffer has no safe fallback left to
// apply (§3, "Why not a held state").
type policyDecision struct {
	addBlock         bool
	setAttrs         map[string]string // attribute -> value ("true" for bools; date for frozen_since)
	reasons          map[string]string // reason attribute -> text (suppression_reason_*, frozen_reason)
	machineryFailure bool              // classPresentUnchanged's failure branch only (§3): schema unchanged, generation/build still broke — the caller (settleBatch/runSync) must abort the whole batch, promoting nothing, rather than apply setAttrs/reasons to the overlay at all
	summary          string            // human-facing report line
}

// decide maps a change class and gate result to the overlay edit (design §7).
// today is the date stamped into frozen_since (YYYY-MM-DD). Two invariants make
// every outcome safe by default: a broken Present type is frozen at its last-good
// bytes (never regressed), and a broken New type is added suppressed (never
// blocks the release).
func decide(class changeClass, gr gateResult, today string) policyDecision {
	switch class {
	case classNew:
		if gr.ok() {
			return policyDecision{addBlock: true, summary: "new: generated OK, added"}
		}
		return policyDecision{
			addBlock: true,
			setAttrs: suppressAttrsForFailures(gr),
			reasons:  reasonsForFailures(gr),
			summary:  "new: generation failed, added suppressed (backlog)",
		}

	case classPresent:
		if gr.ok() {
			return policyDecision{summary: "present: refreshed OK"}
		}
		if gr.anyOK() {
			// Partial failure: the artifacts that generated cleanly are
			// promoted against the new bytes (refreshCandidate); the schema
			// is still pinned (frozen_since) since it cannot partially
			// advance — one JSON file backs all three artifacts — and the
			// artifacts that failed against it are suppressed individually,
			// each with its own reason, so the working artifacts are not held
			// back by the broken one. The freeze itself gets its own
			// frozen_reason rationale (frozenReasonForFailures), distinct
			// from the per-artifact failure text: the freeze is caused by the
			// "can't partially advance one shared JSON" invariant, not by any
			// one artifact's specific error.
			attrs := suppressAttrsForFailures(gr)
			attrs[attrFrozenSince] = today
			reasons := reasonsForFailures(gr)
			reasons[attrFrozenReason] = frozenReasonForFailures(gr)
			return policyDecision{
				setAttrs: attrs,
				reasons:  reasons,
				summary:  "present: partial failure, schema frozen, broken artifact(s) suppressed",
			}
		}
		return policyDecision{
			setAttrs: map[string]string{attrFrozenSince: today},
			reasons:  map[string]string{attrFrozenReason: totalFailureReason(gr)},
			summary:  "present: generation broke, frozen at last-good bytes",
		}

	case classPresentUnchanged:
		if gr.ok() {
			return policyDecision{summary: "present (unchanged): refreshed OK"}
		}
		// The schema did not change, yet generation or the compile gate still
		// broke — self-inflicted, not external (§3): bigdiffer has no
		// last-good schema bytes to fall back to that differ from what it
		// just tried (there is nothing to freeze), and no way to tell whether
		// this failure is narrow (one schema shape) or symptomatic of a
		// broader engine regression that also produced compiling-but-wrong
		// output elsewhere in this same run. So this never sets setAttrs/
		// reasons at all — there is no safe partial action — and instead
		// signals machineryFailure: the caller aborts the whole batch,
		// promoting nothing (§3, "Why all-or-nothing, precisely"). The one
		// legitimate way to keep shipping past a persistent failure here is a
		// human suppress_*-ing the specific artifact deliberately (§3,
		// "Escape hatch"), which is a different, later decide() call once
		// that suppression is already on the row — not a distinct branch in
		// this function.
		return policyDecision{
			machineryFailure: true,
			summary:          "present (unchanged): generation or build broke — machinery regression, failing the run",
		}

	case classNonProvisionable:
		return policyDecision{
			setAttrs: map[string]string{attrNonProvisionable: "true"},
			summary:  "non-provisionable (live): annotated",
		}

	case classWithdrawn:
		return policyDecision{
			setAttrs: map[string]string{attrFrozenSince: today},
			reasons:  map[string]string{attrFrozenReason: formatReason(reasonManual, "withdrawn from AWS, pending major-version removal")},
			summary:  "withdrawn from AWS: frozen, kept pending major-version removal",
		}

	default:
		return policyDecision{summary: "no action"}
	}
}

// suppressAttrFor and reasonAttrFor map an artifact kind to its suppress_* flag
// and suppression_reason_* attribute name, respectively — the single source of
// truth both suppressAttrsForFailures and reasonsForFailures key from, so their
// fallback behavior (below) cannot drift apart by one being updated without the
// other (contributing/docs/bigdiffer-design.md).
func suppressAttrFor(kind artifactKind) string {
	switch kind {
	case artifactSingularDataSource:
		return attrSuppressSingular
	case artifactPluralDataSource:
		return attrSuppressPlural
	default:
		return attrSuppressResource
	}
}

func reasonAttrFor(kind artifactKind) string {
	switch kind {
	case artifactSingularDataSource:
		return attrSuppressionReasonSingular
	case artifactPluralDataSource:
		return attrSuppressionReasonPlural
	default:
		return attrSuppressionReasonResource
	}
}

// suppressAttrsForFailures maps each failed gated artifact to its suppress_* flag,
// so only the artifacts that actually broke are suppressed. If the type failed but
// no specific artifact is attributable, the resource is suppressed as a safe
// default.
func suppressAttrsForFailures(gr gateResult) map[string]string {
	attrs := make(map[string]string)
	for _, a := range gr.artifacts {
		if a.outcome == gateOK {
			continue
		}
		attrs[suppressAttrFor(a.kind)] = "true"
	}
	if len(attrs) == 0 {
		attrs[attrSuppressResource] = "true"
	}
	return attrs
}

// reasonsForFailures renders one suppression_reason_* value per failed
// artifact, keyed by the same attribute name suppressAttrsForFailures uses for
// that artifact's suppress_* flag — so decide() can zip the two maps directly
// when building policyDecision.setAttrs. Each artifact is tagged by which stage
// rejected it: generation_failed (the owned engine returned an error) or
// build_failed (it generated fine but the compile gate's `go build` rejected
// it — item 1).
//
// Mirrors suppressAttrsForFailures' len(attrs) == 0 fallback exactly: if no
// specific artifact is attributable, that function still defensively
// suppresses the resource, so this function must still supply a matching
// reason for it — otherwise -lint's per-field reason anomaly would flag a
// suppression this same decision just made as reason-less.
func reasonsForFailures(gr gateResult) map[string]string {
	reasons := make(map[string]string)
	for _, a := range gr.artifacts {
		if a.outcome == gateOK || a.err == nil {
			continue
		}
		category := reasonGenerationFailed
		if a.outcome == gateFailedBuild {
			category = reasonBuildFailed
		}
		reasons[reasonAttrFor(a.kind)] = formatReason(category, firstLine(a.err.Error()))
	}
	if len(reasons) == 0 {
		reasons[attrSuppressionReasonResource] = formatReason(reasonGenerationFailed,
			"no attributable per-artifact failure; suppressed resource generation defensively")
	}
	return reasons
}

// frozenReasonForFailures renders the frozen_reason value for a classPresent
// partial failure: the freeze there is caused by a structural invariant ("one
// JSON backs all three artifacts, cannot partially advance"), not by any one
// artifact's specific error, so it gets its own fixed rationale sentence
// distinct from reasonsForFailures' per-artifact text — but its category
// prefix must still be derived from what actually triggered the freeze (any
// failing artifact rejected by the compile gate favors build_failed, mirroring
// reasonsForFailures' own per-artifact rule) rather than hardcoded, since a
// partial failure can be triggered by a build failure just as easily as a
// generation failure.
func frozenReasonForFailures(gr gateResult) string {
	category := reasonGenerationFailed
	for _, a := range gr.artifacts {
		if a.outcome == gateFailedBuild {
			category = reasonBuildFailed
		}
	}
	return formatReason(category, "schema pinned — one JSON backs all three artifacts, cannot partially advance")
}

// totalFailureReason renders the frozen_reason for a classPresent type where
// every artifact failed (gr.anyOK() is false): nothing is suppressed in this
// branch (no artifact is promoted, so there is no suppress_* flag to attach a
// per-artifact reason to), the type simply doesn't advance and is frozen at its
// last-good bytes — so, unlike reasonsForFailures, this renders one combined,
// artifact-labeled string (mirroring the pre-split gateFailureReason) directly
// into frozen_reason.
func totalFailureReason(gr gateResult) string {
	var parts []string
	category := reasonGenerationFailed
	for _, a := range gr.artifacts {
		if a.outcome == gateOK || a.err == nil {
			continue
		}
		parts = append(parts, string(a.kind)+": "+firstLine(a.err.Error()))
		if a.outcome == gateFailedBuild {
			category = reasonBuildFailed
		}
	}
	sort.Strings(parts)
	return formatReason(category, strings.Join(parts, "; "))
}

func firstLine(s string) string {
	if before, _, ok := strings.Cut(s, "\n"); ok {
		return before
	}
	return s
}
