// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestRunToolStreamsToLogOnly confirms runTool routes an external tool's full
// output to the log sink (logOnly) rather than the console — the core of the
// console-quieting change. Uses `sh` so it needs no project tooling; the
// console spinner is silent under `go test` (stderr is not a terminal).
func TestRunToolStreamsToLogOnly(t *testing.T) {
	var buf bytes.Buffer
	orig := logOnly
	logOnly = &buf
	t.Cleanup(func() { logOnly = orig })

	if err := runTool("", "sh", toolBar{label: "sh"}, "-c", "printf 'alpha\\nbeta\\ngamma\\n'"); err != nil {
		t.Fatalf("runTool: %v", err)
	}
	got := buf.String()
	for _, want := range []string{"alpha", "beta", "gamma"} {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in the log sink, got %q", want, got)
		}
	}
}

// TestRunToolReportsFailure confirms a non-zero exit still surfaces as an error
// (the output routing must not swallow it).
func TestRunToolReportsFailure(t *testing.T) {
	var buf bytes.Buffer
	orig := logOnly
	logOnly = &buf
	t.Cleanup(func() { logOnly = orig })

	if err := runTool("", "sh", toolBar{label: "sh"}, "-c", "echo boom; exit 3"); err == nil {
		t.Fatal("expected an error from a non-zero exit, got nil")
	}
	if !strings.Contains(buf.String(), "boom") {
		t.Errorf("expected the failing tool's output in the log, got %q", buf.String())
	}
}

// TestPrintActionItemsPrefersArtifactReasonOverFrozenReason covers the
// classPresent partial-failure branch (policy.go), where a type's reasons map
// holds both a specific per-artifact failure (suppression_reason_resource)
// and the generic shared-schema frozen_reason. The actionable, specific
// reason must always win — never nondeterministically hidden behind the
// generic one depending on Go's random map iteration order.
func TestPrintActionItemsPrefersArtifactReasonOverFrozenReason(t *testing.T) {
	var buf bytes.Buffer
	origStructOut, origLogOnly := structOut, logOnly
	structOut, logOnly = &buf, io.Discard
	t.Cleanup(func() { structOut, logOnly = origStructOut, origLogOnly })

	decisions := map[string]policyDecision{
		"AWS::Test::Widget": {
			reasons: map[string]string{
				attrSuppressionReasonResource: "generation_failed: recursive schema",
				attrFrozenReason:              "generation_failed: schema frozen (shared bytes, one artifact broke)",
			},
		},
	}
	printActionItems(decisions)

	got := buf.String()
	if !strings.Contains(got, "recursive schema") {
		t.Errorf("expected the specific artifact reason in the recap, got %q", got)
	}
	if strings.Contains(got, "shared bytes") {
		t.Errorf("frozen_reason should not appear when a specific artifact reason exists, got %q", got)
	}
}

// TestPrintActionItemsUsesFrozenReasonForTotalFailure covers the classPresent
// total-failure branch, where frozen_reason is the only reason recorded (no
// artifact generated cleanly enough to have its own suppression_reason_*) —
// here it's the only explanation there is, so it must be shown.
func TestPrintActionItemsUsesFrozenReasonForTotalFailure(t *testing.T) {
	var buf bytes.Buffer
	origStructOut, origLogOnly := structOut, logOnly
	structOut, logOnly = &buf, io.Discard
	t.Cleanup(func() { structOut, logOnly = origStructOut, origLogOnly })

	decisions := map[string]policyDecision{
		"AWS::Test::Widget": {
			reasons: map[string]string{
				attrFrozenReason: "build_failed: does not compile",
			},
		},
	}
	printActionItems(decisions)

	if got := buf.String(); !strings.Contains(got, "does not compile") {
		t.Errorf("expected the frozen_reason in the recap for a total failure, got %q", got)
	}
}

// TestPrintActionItemsShowsEveryQualifyingArtifactReason confirms a type with
// more than one broken artifact (e.g. both the resource and the singular data
// source) surfaces every qualifying reason, not just one — and in a
// deterministic order, run to run.
func TestPrintActionItemsShowsEveryQualifyingArtifactReason(t *testing.T) {
	var buf bytes.Buffer
	origStructOut, origLogOnly := structOut, logOnly
	structOut, logOnly = &buf, io.Discard
	t.Cleanup(func() { structOut, logOnly = origStructOut, origLogOnly })

	decisions := map[string]policyDecision{
		"AWS::Test::Widget": {
			reasons: map[string]string{
				attrSuppressionReasonResource: "generation_failed: resource broke",
				attrSuppressionReasonSingular: "build_failed: singular data source broke",
			},
		},
	}
	printActionItems(decisions)

	got := buf.String()
	for _, want := range []string{"resource broke", "singular data source broke"} {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in the recap, got %q", want, got)
		}
	}
}

// TestStartRunLogFallsBackLogOnlyToStderr covers a log-open failure: logOnly
// must route to stderr for the run's lifetime rather than silently discarding
// everything an external tool (terraform/tfplugindocs) prints, so the
// "console only" fallback the warning promises is genuinely complete —
// nothing vanishes, it just all lands on the console instead of the log.
func TestStartRunLogFallsBackLogOnlyToStderr(t *testing.T) {
	origStructOut, origLogOnly := structOut, logOnly
	t.Cleanup(func() { structOut, logOnly = origStructOut, origLogOnly })

	dir := t.TempDir()
	t.Chdir(dir)
	const command = "testcmd"
	// Pre-create a directory at the exact path startRunLog will try to
	// os.Create, forcing a deterministic open failure ("is a directory").
	if err := os.Mkdir(".bigdiffer-"+command+".log", 0o755); err != nil {
		t.Fatalf("setting up the log-open failure: %v", err)
	}

	closer := startRunLog(command)
	defer closer()

	if logOnly != os.Stderr {
		t.Errorf("expected logOnly to fall back to os.Stderr on a log-open failure, got %v", logOnly)
	}

	closer()
	if logOnly != io.Discard {
		t.Errorf("expected logOnly to be restored to io.Discard after the closer runs, got %v", logOnly)
	}
}
