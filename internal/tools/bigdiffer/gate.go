// SPDX-License-Identifier: MPL-2.0

package main

import (
	"strings"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared/codegen"
)

// The gate answers "does this schema still generate?" per type. It deliberately
// reuses the proven code-emission engine (shared.NewResource + codegen.Emitter)
// but NOT shared.GenerateTemplateData: that wrapper writes last_resource.txt to
// the CWD (a data race under concurrency) and reads ../identity/names/services.hcl
// via a CWD-relative path (a false failure when run from anywhere else). The
// front half — parse the schema, then emit the root properties schema — is where
// unsupported types, recursion, and meta-argument conflicts actually surface, so
// it is a faithful pass/fail signal. Full template rendering + compilation remain
// covered by the whole-repo build.

type gateOutcome string

const (
	gateOK               gateOutcome = "ok"
	gateFailedValidation gateOutcome = "failed-validation" // schema won't parse/expand
	gateFailedGeneration gateOutcome = "failed-generation" // schema won't emit code
)

// artifactResult is the gate outcome for one generated artifact of a type.
type artifactResult struct {
	kind    artifactKind
	outcome gateOutcome
	err     error
}

// gateResult aggregates a type's per-artifact gate outcomes.
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

// silentUI satisfies cli.Ui for the emitter without emitting anything; gate
// warnings are not failures and are discarded.
type silentUI struct{}

func (silentUI) Ask(string) (string, error)       { return "", nil }
func (silentUI) AskSecret(string) (string, error) { return "", nil }
func (silentUI) Output(string)                    {}
func (silentUI) Info(string)                      {}
func (silentUI) Error(string)                     {}
func (silentUI) Warn(string)                      {}

// gateArtifact runs the generation front half for a single artifact and reports
// whether it validated and emitted. It has no side effects and is safe to call
// concurrently (each call builds its own resource, emitter, and writer).
func gateArtifact(schemaFile, tfType string, isDataSource bool) (gateOutcome, error) {
	res, err := shared.NewResource(tfType, schemaFile)
	if err != nil {
		return gateFailedValidation, err
	}

	var sb strings.Builder
	emitter := codegen.Emitter{
		CfResource:   res.CfResource,
		IsDataSource: isDataSource,
		Ui:           silentUI{},
		Writer:       &sb,
	}
	if _, _, err := emitter.EmitRootPropertiesSchema(res.TfType, make(map[string]string)); err != nil {
		return gateFailedGeneration, err
	}
	return gateOK, nil
}

// gateType runs the gate for the schema-emission artifacts a plan would
// generate: the resource and the singular data source, which both flow through
// codegen.EmitRootPropertiesSchema. Plural data sources are generated from the
// CFN type name via a separate list-data-source path (not schema emission), so
// they are not gated here — a plural-only failure is caught by the whole build.
// A type with no gated artifacts trivially passes.
func gateType(p plan) gateResult {
	gr := gateResult{cfType: p.cfType}
	for _, a := range p.artifacts {
		switch a.kind {
		case artifactResource:
			outcome, err := gateArtifact(p.schemaFile, a.tfType, false)
			gr.artifacts = append(gr.artifacts, artifactResult{kind: a.kind, outcome: outcome, err: err})
		case artifactSingularDataSource:
			outcome, err := gateArtifact(p.schemaFile, a.tfType, true)
			gr.artifacts = append(gr.artifacts, artifactResult{kind: a.kind, outcome: outcome, err: err})
		case artifactPluralDataSource:
			// Not schema-emission-driven; deferred to the build gate.
		}
	}
	return gr
}

// runGate gates the candidate plans. Emission runs serially on purpose: the
// reused legacy code-emission engine is not safe for concurrent in-process use
// (the legacy pipeline parallelizes via separate generator processes, not
// goroutines; shared globals such as inflection rules are mutated during
// emission), which the race detector confirms. The candidate set is the small
// New+Changed set (§4), so serial gating is cheap. Making emission
// concurrency-safe is a future emitter revamp; concurrency stays where it pays
// off — the I/O-bound discovery crawl.
func runGate(plans []plan) []gateResult {
	results := make([]gateResult, len(plans))
	for i, p := range plans {
		results[i] = gateType(p)
	}
	return results
}
