// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"testing"
)

// absSchema resolves a committed schema to an absolute path (immune to any
// working-directory state) and skips the test if it is not present.
func absSchema(t *testing.T, rel string) string {
	t.Helper()
	p, err := filepath.Abs(filepath.Join("../../service/cloudformation/schemas", rel))
	if err != nil {
		t.Fatalf("resolving schema path: %v", err)
	}
	if _, err := os.Stat(p); err != nil {
		t.Skipf("committed schema not present: %v", err)
	}
	return p
}

func absSchemaDir(t *testing.T) string {
	t.Helper()
	p, err := filepath.Abs("../../service/cloudformation/schemas")
	if err != nil {
		t.Fatalf("resolving schema dir: %v", err)
	}
	return p
}

func TestGateArtifactOK(t *testing.T) {
	// No t.Parallel: the reused legacy emitter is not concurrency-safe.
	schema := absSchema(t, "AWS_Logs_LogGroup.json")

	if outcome, err := gateArtifact(schema, "awscc_logs_log_group", false); outcome != gateOK || err != nil {
		t.Errorf("resource gate = %q, %v; want ok, nil", outcome, err)
	}
	if outcome, err := gateArtifact(schema, "awscc_logs_log_group", true); outcome != gateOK || err != nil {
		t.Errorf("singular DS gate = %q, %v; want ok, nil", outcome, err)
	}
}

func TestGateArtifactFailedValidation(t *testing.T) {
	// No t.Parallel: the reused legacy emitter is not concurrency-safe.
	// Missing schema file: NewResource can't read it -> failed-validation.
	outcome, err := gateArtifact("does/not/exist.json", "awscc_svc_thing", false)
	if outcome != gateFailedValidation || err == nil {
		t.Errorf("gate = %q, %v; want failed-validation with an error", outcome, err)
	}
}

func TestGateTypeAndRunGate(t *testing.T) {
	// No t.Parallel: the reused legacy emitter is not concurrency-safe.
	_ = absSchema(t, "AWS_Logs_LogGroup.json") // presence check / skip

	good, err := generationPlan(resourceRow{
		ResourceTypeName:       "aws_logs_log_group",
		CloudFormationTypeName: "AWS::Logs::LogGroup",
	}, "awscc", absSchemaDir(t))
	if err != nil {
		t.Fatalf("generationPlan(good): %v", err)
	}

	bad := plan{
		cfType:     "AWS::Svc::Broken",
		tfType:     "awscc_svc_broken",
		schemaFile: "does/not/exist.json",
		artifacts: []genArtifact{
			{kind: artifactResource, tfType: "awscc_svc_broken"},
		},
	}

	results := runGate([]plan{good, bad})
	if len(results) != 2 {
		t.Fatalf("results = %d, want 2", len(results))
	}

	byType := map[string]gateResult{}
	for _, r := range results {
		byType[r.cfType] = r
	}

	if g := byType["AWS::Logs::LogGroup"]; !g.ok() {
		t.Errorf("LogGroup should gate ok, got %+v", g.artifacts)
	} else if len(g.artifacts) != 2 {
		t.Errorf("LogGroup gated artifacts = %d, want 2 (resource + singular)", len(g.artifacts))
	}

	if g := byType["AWS::Svc::Broken"]; g.ok() {
		t.Errorf("broken plan should not gate ok")
	} else if g.artifacts[0].outcome != gateFailedValidation {
		t.Errorf("broken outcome = %q, want failed-validation", g.artifacts[0].outcome)
	}
}
