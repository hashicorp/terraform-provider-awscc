// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

// Generation outcome types consumed by the policy engine. Generation itself is
// the gate: refreshCandidate generates a type and gateResultFromGenResults
// (update.go) maps the per-artifact success/failure into a gateResult, which
// decide() (policy.go) turns into an overlay edit. (An earlier standalone
// front-half gate that re-emitted schemas to pass/fail types was superseded by
// this generate-for-real approach and removed.)

type gateOutcome string

const (
	gateOK               gateOutcome = "ok"
	gateFailedGeneration gateOutcome = "failed-generation" // schema won't generate
)

// artifactResult is the generation outcome for one artifact of a type.
type artifactResult struct {
	kind    artifactKind
	outcome gateOutcome
	err     error
}

// gateResult aggregates a type's per-artifact generation outcomes.
type gateResult struct {
	cfType    string
	artifacts []artifactResult
}

// ok reports whether every artifact generated cleanly.
func (g gateResult) ok() bool {
	for _, a := range g.artifacts {
		if a.outcome != gateOK {
			return false
		}
	}
	return true
}

// anyOK reports whether at least one artifact generated cleanly. Used to tell a
// partial failure (some artifacts promoted, others suppressed) apart from a
// total one (nothing to promote; the type is frozen/suppressed as a whole).
func (g gateResult) anyOK() bool {
	for _, a := range g.artifacts {
		if a.outcome == gateOK {
			return true
		}
	}
	return false
}
