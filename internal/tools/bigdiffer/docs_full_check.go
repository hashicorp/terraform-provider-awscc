// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

// The full tier of -check's docs freshness enforcement
// (bigdiffer-design.md §6, "Documentation is part of the same gated
// pipeline"): render the registry docs (tfplugindocs) themselves and diff
// against committed docs/, the same "-check passes iff running -reconcile
// and committing the result would be a no-op" contract
// diffStagedTrees/diffImportExamples already extend to code and the
// import-examples aggregate.
//
// tfplugindocs is not static analysis of source: by default it builds the
// provider and drives `terraform providers schema -json` against a real
// running binary. Item 2's load-bearing constraint is that the render must be
// a pure function of a tree root, shared by -sync/-reconcile (the promoted
// real tree) and -check (the staged tree) — so the seam is the schema JSON
// itself: extract it once per tree, then feed both callers through the
// identical `tfplugindocs generate --providers-schema <path>` render (a real,
// documented flag — `tfplugindocs generate --help`: "Setting this flag will
// skip building the provider and calling Terraform CLI"). That confines every
// tree-dependent step (building a binary, running Terraform) to one
// extraction call; the render itself never touches the tree again.
//
// -check has no standalone "staged tree" to build from at all: settleBatch
// only stages *changed* files under stagingDir/out, not a full buildable
// module (mirroring buildOnce's own design — see compile.go's top-of-file
// comment). So this reuses the identical overlay-onto-the-real-tree-then-
// revert mechanism the compile gate already uses (overlayFiles/revert), just
// building a real, runnable binary via `go build -o` instead of a throwaway
// `go build ./...` pass/fail signal, then reverting unconditionally before
// returning — exactly buildOnce's own safety contract, extended by one more
// step (extract + render) inside the same window.

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sort"
	"syscall"
)

// diffRenderedDocs renders the registry docs from the tree settleBatch staged
// (stagingDir/out overlaid onto the real cfg.outputRoot, exactly as the
// compile gate itself builds against it) and diffs the result against
// committed cfg.docsDir, returning the destination paths (relative to the
// repo root) of every file that differs or that has no committed counterpart
// at all. rows is the same projected row set diffImportExamples uses
// (projectRows), so the scratch import-example docs this renders from match
// exactly what item 2's cheap tier already confirmed (or, if it didn't,
// -check has already failed before this runs — see runCheck's call-site
// comment on ordering).
//
// The real tree is overlaid only for the duration of one build + one
// `terraform providers schema -json` call, then unconditionally reverted —
// never left overlaid, regardless of success, failure, or a SIGINT/SIGTERM
// arriving mid-window (see extractProviderSchema for the exact mechanism,
// shared verbatim with buildOnce).
func diffRenderedDocs(ctx context.Context, cfg config, stagingDir string, rows []resourceRow) ([]string, error) {
	files, err := collectStagedGoFiles(filepath.Join(stagingDir, "out"), cfg.outputRoot)
	if err != nil {
		return nil, fmt.Errorf("collecting staged files: %w", err)
	}
	// The registration file is part of the buildable module too, exactly as
	// compileFixpoint's own files map includes it — a docs render against a
	// tree missing it would fail to build for a reason -check has nothing to
	// do with (a stale registration file is -lint's job, not this check's).
	reg, err := emitRegistration(cfg, rows)
	if err != nil {
		return nil, fmt.Errorf("rendering registration file for the docs render: %w", err)
	}
	files[cfg.registrationPath] = reg

	schemaJSON, err := extractProviderSchema(ctx, cfg.repoRoot, files)
	if err != nil {
		return nil, fmt.Errorf("extracting provider schema from the staged tree: %w", err)
	}

	scratchDir, err := os.MkdirTemp("", "bigdiffer-docs-check-")
	if err != nil {
		return nil, fmt.Errorf("creating scratch render dir: %w", err)
	}
	defer func() { _ = os.RemoveAll(scratchDir) }()

	// examples/ is not purely bigdiffer's output: hand-authored usage
	// examples (e.g. a resource's *.tf demonstrating real-world config,
	// predating and unrelated to the import-* files bigdiffer owns) live in
	// the same per-resource directories, and doc templates reference them by
	// name via {{tffile ...}}. Start from a full copy of the real, committed
	// tree — exactly what a maintainer's working directory already has —
	// then overlay only the regenerated import-example files on top, the
	// same relationship runDocs itself has with the real cfg.examplesDir (it
	// writes generated files into a directory that already has hand-authored
	// ones sitting in it, never replacing the whole tree). Missing this and
	// rendering into an examples/ dir containing only generated files was a
	// real bug caught only by TestDiffRenderedDocsNoDiff actually running
	// against the real corpus, not by any synthetic test.
	scratchExamplesDir := filepath.Join(scratchDir, "examples")
	if err := copyDirTree(cfg.examplesDir, scratchExamplesDir); err != nil {
		return nil, fmt.Errorf("copying examples/ into the scratch render dir: %w", err)
	}
	examples, err := buildImportExamples(cfg, rows)
	if err != nil {
		return nil, fmt.Errorf("deriving import examples for the docs render: %w", err)
	}
	if _, err := writeImportExampleDocs(examples, scratchExamplesDir); err != nil {
		return nil, fmt.Errorf("rendering scratch import-example docs: %w", err)
	}
	// templates/ is a human-maintained source input, not generated output —
	// tfplugindocs reads it from providerDir, so the scratch provider-dir
	// needs a copy alongside the scratch examples, but templates/ itself is
	// never a candidate for drift the way examples/ or docs/ are.
	if err := copyDirTree(filepath.Join(cfg.repoRoot, "templates"), filepath.Join(scratchDir, "templates")); err != nil {
		return nil, fmt.Errorf("copying templates/ into the scratch render dir: %w", err)
	}

	scratchDocsDir := filepath.Join(scratchDir, "docs")
	if err := renderDocsFromSchema(scratchDir, scratchDocsDir, schemaJSON); err != nil {
		return nil, fmt.Errorf("rendering docs from the extracted schema: %w", err)
	}

	return diffDocsTrees(cfg, scratchDocsDir)
}

// extractProviderSchema builds a real, runnable provider binary from the
// staged tree and returns the bytes of `terraform providers schema -json`
// against it, via a scratch Terraform CLI config pointing a dev_overrides
// entry at the binary (the identical mechanism GNUmakefile's
// check-startup-errors target already uses for local testing) — no
// `terraform init`, no registry, no lock file, so this is safe to run offline
// and repeatedly.
//
// The overlay window (staging files onto the real tree, building, extracting
// the schema, then reverting) is buildOnce's own safety contract (compile.go)
// verbatim, extended by the schema-extraction step: overlayFiles records
// prior state before writing, a SIGINT/SIGTERM handler reverts and exits
// rather than leaving Go's default terminate-immediately behavior to skip the
// deferred revert, and the deferred revert itself runs regardless of the
// build's or extraction's outcome. The real tree is bit-for-bit unchanged by
// the time this returns, success or failure.
func extractProviderSchema(ctx context.Context, repoRoot string, stagedFiles map[string][]byte) (schemaJSON []byte, err error) {
	overlay, oerr := overlayFiles(stagedFiles)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	done := make(chan struct{})
	go func() {
		select {
		case sig := <-sigCh:
			_ = overlay.revert()
			fmt.Fprintf(os.Stderr, "\ndocs check: received %s, reverted overlay before exiting\n", sig)
			os.Exit(1)
		case <-done:
		}
	}()
	defer func() {
		close(done)
		signal.Stop(sigCh)
		if rerr := overlay.revert(); rerr != nil && err == nil {
			err = fmt.Errorf("reverting docs-check overlay: %w", rerr)
		}
	}()
	if oerr != nil {
		return nil, fmt.Errorf("staging docs-check overlay: %w", oerr)
	}

	binDir, berr := os.MkdirTemp("", "bigdiffer-docs-check-bin-")
	if berr != nil {
		return nil, fmt.Errorf("creating scratch binary dir: %w", berr)
	}
	defer func() { _ = os.RemoveAll(binDir) }()
	binPath := filepath.Join(binDir, "terraform-provider-awscc")

	buildCmd := exec.CommandContext(ctx, "go", "build", "-o", binPath, ".")
	buildCmd.Dir = repoRoot
	if out, berr := buildCmd.CombinedOutput(); berr != nil {
		if ctx.Err() != nil {
			return nil, fmt.Errorf("go build -o (docs check) cancelled: %w", ctx.Err())
		}
		return nil, fmt.Errorf("go build -o %s .: %w\n%s", binPath, berr, out)
	}

	tfConfigDir, terr := os.MkdirTemp("", "bigdiffer-docs-check-tfrc-")
	if terr != nil {
		return nil, fmt.Errorf("creating scratch terraform config dir: %w", terr)
	}
	defer func() { _ = os.RemoveAll(tfConfigDir) }()
	tfConfigPath := filepath.Join(tfConfigDir, "terraform.rc")
	tfConfig := fmt.Sprintf("provider_installation {\n  dev_overrides {\n    \"hashicorp/awscc\" = %q\n  }\n  direct {}\n}\n", binDir)
	if err := os.WriteFile(tfConfigPath, []byte(tfConfig), filePerm); err != nil {
		return nil, fmt.Errorf("writing scratch terraform config: %w", err)
	}
	// terraform providers schema -json only reports schemas for providers a
	// configuration actually requires — a dev_overrides entry alone tells
	// Terraform where to find the awscc binary, not that anything in this
	// directory wants it. Without this file, the command succeeds but
	// returns an empty provider_schemas map, and tfplugindocs fails later
	// with "unable to find schema in JSON for provider \"awscc\"" — a real
	// bug caught only by TestDiffRenderedDocsNoDiff actually running the
	// full extraction against the real repo, not by any synthetic test.
	mainTF := "terraform {\n  required_providers {\n    awscc = {\n      source = \"hashicorp/awscc\"\n    }\n  }\n}\n"
	if err := os.WriteFile(filepath.Join(tfConfigDir, "main.tf"), []byte(mainTF), filePerm); err != nil {
		return nil, fmt.Errorf("writing scratch main.tf: %w", err)
	}

	schemaCmd := exec.CommandContext(ctx, "terraform", "providers", "schema", "-json")
	schemaCmd.Dir = tfConfigDir
	schemaCmd.Env = append(os.Environ(), "TF_CLI_CONFIG_FILE="+tfConfigPath)
	var stdout, stderr bytes.Buffer
	schemaCmd.Stdout = &stdout
	schemaCmd.Stderr = &stderr
	if err := schemaCmd.Run(); err != nil {
		if ctx.Err() != nil {
			return nil, fmt.Errorf("terraform providers schema -json cancelled: %w", ctx.Err())
		}
		return nil, fmt.Errorf("terraform providers schema -json: %w\n%s", err, stderr.String())
	}
	return stdout.Bytes(), nil
}

// renderDocsFromSchema runs `tfplugindocs generate` against providerDir
// (which must already contain templates/ and examples/, real or scratch)
// using the pre-extracted schemaJSON via --providers-schema — the one flag
// that makes the render a pure function of (schema, templates, examples),
// with no build or Terraform CLI call of its own (see docs_full_check.go's
// top-of-file comment). Writes rendered docs under docsDir.
func renderDocsFromSchema(providerDir, docsDir string, schemaJSON []byte) error {
	schemaPath := filepath.Join(providerDir, ".schema.json")
	if err := os.WriteFile(schemaPath, schemaJSON, filePerm); err != nil {
		return fmt.Errorf("writing scratch schema file: %w", err)
	}
	if err := os.MkdirAll(docsDir, dirPerm); err != nil {
		return fmt.Errorf("creating %s: %w", docsDir, err)
	}
	if err := runTool(providerDir, "tfplugindocs", "generate",
		"--provider-dir", providerDir,
		"--provider-name", "terraform-provider-awscc",
		"--providers-schema", schemaPath,
		"--rendered-website-dir", docsDir,
	); err != nil {
		return fmt.Errorf("tfplugindocs: %w", err)
	}
	return nil
}

// diffDocsTrees compares every rendered file under renderedDocsDir against
// its committed counterpart in cfg.docsDir, returning the destination paths
// (relative to the repo root) of every file that differs or that has no
// committed counterpart at all. Deliberately one-directional
// (rendered -> committed, never the reverse), the identical convention
// diffStagedTrees already established for code and cfg.cacheDir, for the
// identical reason: -check's contract is "passes iff -reconcile + commit
// would be a no-op," and -reconcile has no docs-orphan-cleanup capability
// either (bigdiffer-design.md §6, "Neither -reconcile nor -check removes
// orphaned output") — scanning the other direction would make -check fail on
// a state nothing in the pipeline can fix.
func diffDocsTrees(cfg config, renderedDocsDir string) ([]string, error) {
	var diffs []string
	err := filepath.Walk(renderedDocsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(renderedDocsDir, path)
		if err != nil {
			return err
		}
		renderedBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading rendered %s: %w", path, err)
		}
		committedPath := filepath.Join(cfg.docsDir, rel)
		committedBytes, err := os.ReadFile(committedPath)
		if err != nil {
			if os.IsNotExist(err) {
				diffs = append(diffs, relToRepoRoot(cfg, committedPath))
				return nil
			}
			return fmt.Errorf("reading committed %s: %w", committedPath, err)
		}
		if !bytes.Equal(renderedBytes, committedBytes) {
			diffs = append(diffs, relToRepoRoot(cfg, committedPath))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(diffs)
	return diffs, nil
}

// copyDirTree copies every regular file under src into the identical
// relative path under dst, creating directories as needed — the same shape
// as copyTree (sync.go), duplicated here rather than shared because this one
// deliberately errors on a missing src (templates/ must exist; there is
// nothing sensible to render without it), where copyTree's silent skip on a
// missing source is specifically about an optional staged root.
func copyDirTree(src, dst string) error {
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("reading %s: %w", src, err)
	}
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, dirPerm)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		if err := os.MkdirAll(filepath.Dir(target), dirPerm); err != nil {
			return err
		}
		return os.WriteFile(target, data, filePerm)
	})
}
