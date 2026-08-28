// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/cli"
)

// repoPath resolves a repo-relative path from this test's directory
// (internal/tools/bigdiffer/codegen → four levels up is the repo root).
func repoPath(t *testing.T, rel ...string) string {
	t.Helper()
	abs, err := filepath.Abs(filepath.Join(append([]string{"..", "..", "..", ".."}, rel...)...))
	if err != nil {
		t.Fatalf("resolving %v: %v", rel, err)
	}
	return abs
}

// TestGenerateResourceSmoke proves the owned engine emits a real type's code
// in-process (Brick 6's "done"): parse schema → build template data → render →
// gofmt. It asserts structure, not byte-parity — full-corpus parity is Brick 7.
func TestGenerateResourceSmoke(t *testing.T) {
	t.Parallel()

	schema := repoPath(t, "internal", "service", "cloudformation", "schemas", "AWS_Logs_LogGroup.json")
	if _, err := os.Stat(schema); err != nil {
		t.Skipf("committed schema not found (%v)", err)
	}
	services := repoPath(t, "internal", "identity", "names", "services.hcl")

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}
	code, test, err := GenerateResource(ui, schema, "awscc_logs_log_group", "logs", services, true)
	if err != nil {
		t.Fatalf("GenerateResource: %v", err)
	}

	for _, want := range []string{
		"package logs",
		"func init()",
		`registry.AddResourceFactory("awscc_logs_log_group"`,
		"func logGroupResource(",
	} {
		if !strings.Contains(string(code), want) {
			t.Errorf("generated resource code missing %q", want)
		}
	}
	if !strings.Contains(string(test), "package logs") {
		t.Errorf("generated acceptance test missing package declaration")
	}
}
