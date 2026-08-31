// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/codegen"
)

// TestImportExampleDocsParity confirms bigdiffer's owned import-example docs match
// the committed files. It compares the files that terraform fmt does not touch
// (import.sh, import-by-string-id.tf, list-resource.tfquery.hcl); the
// import-by-identity.tf block is reformatted by docs-fmt after generation, so its
// committed form is post-fmt and is validated separately.
func TestImportExampleDocsParity(t *testing.T) {
	if testing.Short() {
		t.Skip("iterates the whole import-examples corpus; run without -short")
	}
	cfg, _ := loadCorpus(t)

	examples, err := loadImportExamples(cfg.importExamplesPath)
	if err != nil {
		t.Fatalf("loadImportExamples: %v", err)
	}
	if len(examples) == 0 {
		t.Fatal("no import examples loaded")
	}

	compare := map[string]bool{
		"import.sh":                 true,
		"import-by-string-id.tf":    true,
		"list-resource.tfquery.hcl": true,
	}

	var checked, mismatch, missing int
	for _, ex := range examples {
		files, err := codegen.GenerateImportExampleDocs(ex)
		if err != nil {
			t.Fatalf("%s: %v", ex.ResourceName, err)
		}
		for _, f := range files {
			if !compare[filepath.Base(f.RelPath)] {
				continue
			}
			want, rerr := os.ReadFile(filepath.Join(cfg.examplesDir, f.RelPath))
			if os.IsNotExist(rerr) {
				missing++
				continue
			}
			if rerr != nil {
				t.Fatalf("reading committed %s: %v", f.RelPath, rerr)
			}
			checked++
			if !bytes.Equal(f.Content, want) {
				mismatch++
				if mismatch <= 10 {
					t.Errorf("DRIFT %s:\n%s", f.RelPath, firstDiff(want, f.Content))
				}
			}
		}
	}
	t.Logf("import-example docs: %d files checked, %d drift, %d missing-committed", checked, mismatch, missing)
	if mismatch > 0 {
		t.Fatalf("import-example docs parity failed: %d drift", mismatch)
	}
}

// TestImportByIdentityFmt closes the loop on the one fmt-confounded file: it runs
// bigdiffer's raw import-by-identity.tf through `terraform fmt` and confirms the
// result matches the committed (post-fmt) file, for a single- and a
// multi-identifier resource. Skipped when terraform is not on PATH.
func TestImportByIdentityFmt(t *testing.T) {
	if _, err := exec.LookPath("terraform"); err != nil {
		t.Skip("terraform not on PATH")
	}
	cfg, _ := loadCorpus(t)
	examples, err := loadImportExamples(cfg.importExamplesPath)
	if err != nil {
		t.Fatalf("loadImportExamples: %v", err)
	}
	byName := map[string]codegen.ImportExample{}
	for _, ex := range examples {
		byName[ex.ResourceName] = ex
	}

	for _, name := range []string{"awscc_logs_log_group", "awscc_acmpca_certificate"} {
		ex, ok := byName[name]
		if !ok {
			t.Skipf("%s not in import examples", name)
		}
		files, err := codegen.GenerateImportExampleDocs(ex)
		if err != nil {
			t.Fatal(err)
		}
		var raw []byte
		for _, f := range files {
			if filepath.Base(f.RelPath) == "import-by-identity.tf" {
				raw = f.Content
			}
		}
		if raw == nil {
			t.Fatalf("%s: no import-by-identity.tf produced", name)
		}

		formatted, err := terraformFmtStdin(t, raw)
		if err != nil {
			t.Fatalf("%s: terraform fmt: %v", name, err)
		}
		want, err := os.ReadFile(filepath.Join(cfg.examplesDir, "resources", name, "import-by-identity.tf"))
		if err != nil {
			t.Fatalf("reading committed: %v", err)
		}
		if !bytes.Equal(formatted, want) {
			t.Errorf("%s import-by-identity.tf after fmt != committed:\n%s", name, firstDiff(want, formatted))
		}
	}
}

// terraformFmtStdin formats Terraform config from bytes via `terraform fmt -`.
func terraformFmtStdin(t *testing.T, in []byte) ([]byte, error) {
	t.Helper()
	cmd := exec.Command("terraform", "fmt", "-")
	cmd.Stdin = bytes.NewReader(in)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// TestGenerateImportExampleDocsWrites exercises the package-main docs-import
// writer into a temp examples dir (never the real tree), checking files land at
// the expected paths.
func TestGenerateImportExampleDocsWrites(t *testing.T) {
	t.Parallel()
	cfg, _ := loadCorpus(t)
	tmp := t.TempDir()

	n, err := generateImportExampleDocs(cfg.importExamplesPath, tmp)
	if err != nil {
		t.Fatalf("generateImportExampleDocs: %v", err)
	}
	if n == 0 {
		t.Fatal("wrote 0 files")
	}
	if _, err := os.Stat(filepath.Join(tmp, "resources", "awscc_logs_log_group", "import.sh")); err != nil {
		t.Errorf("expected import.sh for awscc_logs_log_group: %v", err)
	}
}
