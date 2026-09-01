// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Overlay attribute names the policy sets on a block.
const (
	attrFrozenSince       = "frozen_since"
	attrNonProvisionable  = "non_provisionable"
	attrSuppressResource  = "suppress_resource_generation"
	attrSuppressSingular  = "suppress_singular_data_source_generation"
	attrSuppressPlural    = "suppress_plural_data_source_generation"
	attrSuppressionReason = "suppression_reason"
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
type changeClass string

const (
	classNew              changeClass = "new"
	classPresent          changeClass = "present"
	classNonProvisionable changeClass = "non_provisionable"
	classWithdrawn        changeClass = "withdrawn"
)

// policyDecision is the overlay edit for one type: whether to add a new block,
// which attributes to set on it, an optional suppression_reason, and a one-line
// report summary. Per the design invariants, no decision blocks a release.
type policyDecision struct {
	addBlock bool
	setAttrs map[string]string // attribute -> value ("true" for bools; date for frozen_since)
	reason   string            // suppression_reason text, when suppressing/freezing
	summary  string            // human-facing report line
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
			reason:   gateFailureReason(gr),
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
			// back by the broken one.
			attrs := suppressAttrsForFailures(gr)
			attrs[attrFrozenSince] = today
			return policyDecision{
				setAttrs: attrs,
				reason:   gateFailureReason(gr),
				summary:  "present: partial failure, schema frozen, broken artifact(s) suppressed",
			}
		}
		return policyDecision{
			setAttrs: map[string]string{attrFrozenSince: today},
			reason:   gateFailureReason(gr),
			summary:  "present: generation broke, frozen at last-good bytes",
		}

	case classNonProvisionable:
		return policyDecision{
			setAttrs: map[string]string{attrNonProvisionable: "true"},
			summary:  "non-provisionable (live): annotated",
		}

	case classWithdrawn:
		return policyDecision{
			setAttrs: map[string]string{attrFrozenSince: today},
			reason:   formatReason(reasonManual, "withdrawn from AWS, pending major-version removal"),
			summary:  "withdrawn from AWS: frozen, kept pending major-version removal",
		}

	default:
		return policyDecision{summary: "no action"}
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
		switch a.kind {
		case artifactResource:
			attrs[attrSuppressResource] = "true"
		case artifactSingularDataSource:
			attrs[attrSuppressSingular] = "true"
		case artifactPluralDataSource:
			attrs[attrSuppressPlural] = "true"
		}
	}
	if len(attrs) == 0 {
		attrs[attrSuppressResource] = "true"
	}
	return attrs
}

// gateFailureReason renders a compact, deterministic suppression_reason from the
// failed artifacts' errors. Each artifact is tagged by which stage rejected it:
// generation_failed (the owned engine returned an error) or build_failed (it
// generated fine but the compile gate's `go build` rejected it — item 1). A
// gateResult with a mix of both across its artifacts renders a mixed reason;
// each artifact's own tag is still correct individually via suppressAttrsForFailures.
func gateFailureReason(gr gateResult) string {
	var parts []string
	category := reasonGenerationFailed
	for _, a := range gr.artifacts {
		if a.outcome == gateOK || a.err == nil {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %s", a.kind, firstLine(a.err.Error())))
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
