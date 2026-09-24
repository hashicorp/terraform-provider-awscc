// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
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
