// SPDX-License-Identifier: MPL-2.0

package main

// Brick 10 — documentation. bigdiffer owns docs-import (the import-example docs,
// generated from import_examples_gen.json) and orchestrates the two external
// steps of `make docs-all` — terraform fmt and tfplugindocs — rather than
// reimplementing them. Order matches the legacy target: import, fmt, then
// tfplugindocs.

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/codegen"
)

// importExampleEntry mirrors one record in import_examples_gen.json (whose keys
// are lower-cased), decoupled from codegen.ImportExample's field names.
type importExampleEntry struct {
	Resource             string   `json:"resource"`
	Identifier           []string `json:"identifier"`
	GenerateListResource bool     `json:"generateListResource"`
}

// loadImportExamples decodes the import_examples_gen.json aggregate into the
// codegen entries used to render the docs.
func loadImportExamples(path string) ([]codegen.ImportExample, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var entries []importExampleEntry
	if err := json.Unmarshal(b, &entries); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	out := make([]codegen.ImportExample, len(entries))
	for i, e := range entries {
		out[i] = codegen.ImportExample{
			ResourceName:         e.Resource,
			Identifier:           e.Identifier,
			GenerateListResource: e.GenerateListResource,
		}
	}
	return out, nil
}

// generateImportExampleDocs renders every resource's import-example files from
// the aggregate and writes them under examplesDir. Returns the number of files
// written. This is bigdiffer's owned docs-import.
func generateImportExampleDocs(importExamplesPath, examplesDir string) (int, error) {
	examples, err := loadImportExamples(importExamplesPath)
	if err != nil {
		return 0, err
	}

	written := 0
	for _, ex := range examples {
		files, err := codegen.GenerateImportExampleDocs(ex)
		if err != nil {
			return written, fmt.Errorf("%s: %w", ex.ResourceName, err)
		}
		for _, f := range files {
			path := filepath.Join(examplesDir, f.RelPath)
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				return written, fmt.Errorf("creating dir for %s: %w", path, err)
			}
			if err := os.WriteFile(path, f.Content, 0o644); err != nil {
				return written, fmt.Errorf("writing %s: %w", path, err)
			}
			written++
		}
	}
	return written, nil
}

// runDocs runs the full documentation pipeline: own docs-import, then orchestrate
// terraform fmt (docs-fmt) and tfplugindocs (docs), matching `make docs-all`.
func runDocs(cfg config) error {
	n, err := generateImportExampleDocs(cfg.importExamplesPath, cfg.examplesDir)
	if err != nil {
		return fmt.Errorf("docs-import: %w", err)
	}
	fmt.Fprintf(os.Stderr, "bigdiffer: docs-import wrote %d import-example files\n", n)

	// docs-fmt: format the generated Terraform example files.
	fmt.Fprintln(os.Stderr, "bigdiffer: docs-fmt (terraform fmt -recursive)")
	if err := runTool(cfg.examplesDir, "terraform", "fmt", "-recursive"); err != nil {
		return fmt.Errorf("docs-fmt: %w", err)
	}

	// docs: regenerate the registry docs with tfplugindocs (external tool).
	for _, sub := range []string{"data-sources", "resources", "list-resources"} {
		if err := removeMarkdown(filepath.Join(cfg.docsDir, sub)); err != nil {
			return fmt.Errorf("clearing docs/%s: %w", sub, err)
		}
	}
	fmt.Fprintln(os.Stderr, "bigdiffer: docs (tfplugindocs generate)")
	if err := runTool(cfg.repoRoot, "tfplugindocs", "generate", "--provider-name", "terraform-provider-awscc"); err != nil {
		return fmt.Errorf("tfplugindocs: %w", err)
	}
	return nil
}

// runTool runs an external command in dir, streaming its output.
func runTool(dir, name string, args ...string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("%s not found on PATH (install it via `make prereq`): %w", name, err)
	}
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// removeMarkdown deletes the *.md files in dir (if it exists), matching the
// legacy docs target's `rm -f docs/<sub>/*.md`.
func removeMarkdown(dir string) error {
	matches, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return err
	}
	for _, m := range matches {
		if err := os.Remove(m); err != nil {
			return err
		}
	}
	return nil
}
