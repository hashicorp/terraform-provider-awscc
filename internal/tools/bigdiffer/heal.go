// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/cli"
)

// healProposal is one row -heal has something to say about. -heal never
// mutates all_schemas.hcl; every proposal is reported for a human to apply,
// matching every other bigdiffer policy decision
// (contributing/docs/suppressed-and-frozen.md, "-heal: re-probe and fill gaps").
type healProposal struct {
	cfn    string
	label  string
	kind   string // which artifact this proposal is about, or "" for the freeze itself
	action string // "lift" (generation now succeeds) or "reason" (a reason was determined/migrated)
	reason string // the proposed suppression_reason value
}

// runHeal re-probes every suppressed or frozen row that has no
// suppression_reason (or one tagged unknown) and proposes a reclassification.
// Offline except for reading the committed schema cache; never touches AWS and
// never writes all_schemas.hcl.
func runHeal(allSchemasPath string) error {
	cfg, rows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}

	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}
	loc := countLineRE.FindStringIndex(string(overlayContent))
	if loc == nil {
		return fmt.Errorf("could not find the count header comment in all_schemas.hcl")
	}
	items := extractItems(string(overlayContent)[loc[1]:])
	commentByCFN := make(map[string]string, len(items))
	for _, it := range items {
		if !it.live || it.key == "" {
			continue
		}
		if c := suppressionComment(it.text); c != "" {
			commentByCFN[it.key] = c
		}
	}

	stepf("Re-probing suppressed/frozen rows with no recorded reason…")
	var proposals []healProposal
	needsReason := 0
	for _, row := range rows {
		if !isSuppressedOrFrozen(row) {
			continue
		}
		if row.SuppressionReason != "" && !strings.HasPrefix(row.SuppressionReason, string(reasonUnknown)+":") {
			continue // already has a real reason; -heal only targets the reason-less/unknown backlog
		}
		needsReason++
		proposals = append(proposals, healRow(cfg, row, commentByCFN[row.CloudFormationTypeName])...)
	}
	sort.Slice(proposals, func(i, j int) bool {
		if proposals[i].cfn != proposals[j].cfn {
			return proposals[i].cfn < proposals[j].cfn
		}
		return proposals[i].kind < proposals[j].kind
	})

	writeHealReport(needsReason, proposals)
	return nil
}

// healRow probes one row's suppressed artifacts (and its freeze, if any) and
// returns a proposal per artifact/freeze it can say something about. It never
// mutates row or the overlay.
func healRow(cfg config, row resourceRow, comment string) []healProposal {
	var out []healProposal
	label := row.ResourceTypeName
	cfn := row.CloudFormationTypeName

	schema, schemaErr := os.ReadFile(row.CloudFormationSchemaPath)
	if schemaErr != nil {
		schema, schemaErr = os.ReadFile(schemaCachePath(cfg.cacheDir, cfn))
	}

	type suppressedArtifact struct {
		kind artifactKind
		flag bool
	}
	for _, sa := range []suppressedArtifact{
		{artifactResource, row.SuppressResourceGeneration},
		{artifactSingularDataSource, row.SuppressSingularDataSourceGeneration},
		{artifactPluralDataSource, row.SuppressPluralDataSourceGeneration},
	} {
		if !sa.flag {
			continue
		}
		out = append(out, healArtifact(cfg, row, sa.kind, schema, schemaErr, comment))
	}

	if row.FrozenSince != "" && (len(out) == 0 || allReasonProposals(out) == 0) {
		// The freeze itself has no reason distinct from its artifacts' causes;
		// propose migrating the free-form comment (if any) or unknown.
		out = append(out, freezeProposal(cfn, label, comment))
	}
	return out
}

// healArtifact probes one suppressed artifact: structural check first (plural
// only), then a real regeneration attempt, then a comment-migration fallback.
func healArtifact(cfg config, row resourceRow, kind artifactKind, schema []byte, schemaErr error, comment string) healProposal {
	base := healProposal{cfn: row.CloudFormationTypeName, label: row.ResourceTypeName, kind: string(kind)}

	if kind == artifactPluralDataSource && schemaErr == nil {
		if !pluralSupported(string(schema)) {
			base.action = "reason"
			base.reason = formatReason(reasonStructural, "no list handler with zero required arguments")
			return base
		}
	}

	if schemaErr == nil {
		if err := probeArtifact(cfg, row, kind, schema); err == nil {
			base.action = "lift"
			base.reason = "generates cleanly now — propose lifting the suppression"
			return base
		} else {
			base.action = "reason"
			base.reason = formatReason(reasonGenerationFailed, firstLine(err.Error()))
			return base
		}
	}

	return commentOrUnknown(base, comment)
}

// probeArtifact regenerates one artifact from schema bytes staged to a temp
// file, without writing any output — a read-only re-probe. It never runs the
// owned engine in this process: some suppressed rows are suppressed *because*
// their schema is recursive (e.g. the free-form "Recursive Attribute
// Definitions" comments, issue #95), and codegen.Emitter has no recursion-depth
// guard, so re-probing one in-process is a confirmed stack overflow — a fatal,
// unrecoverable crash that would take the whole -heal run down with it. Instead
// probeArtifact re-execs bigdiffer itself into the hidden -heal-probe-artifact
// mode (runHealProbeArtifact) under a timeout and a soft memory cap, so a crash
// or runaway probe kills only that subprocess (the same "one failure never
// blocks the rest" principle discover.go and generateCorpus already apply).
func probeArtifact(cfg config, row resourceRow, kind artifactKind, schema []byte) error {
	self, err := os.Executable()
	if err != nil {
		return fmt.Errorf("locating bigdiffer binary to probe in isolation: %w", err)
	}
	return probeArtifactWithBinary(self, cfg, row, kind, schema)
}

// probeArtifactWithBinary is probeArtifact's implementation, parameterized on
// the binary to re-exec so tests can point it at a binary built from the
// current source rather than relying on os.Executable() (which resolves to
// the test binary under `go test` — a binary whose CLI is testing.Main, not
// bigdiffer's real main(), and so can't be probed directly).
func probeArtifactWithBinary(bin string, cfg config, row resourceRow, kind artifactKind, schema []byte) error {
	tmp, err := os.CreateTemp("", "bigdiffer-heal-*.json")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tmp.Name()) }()
	if _, err := tmp.Write(schema); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), healProbeTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin,
		"-heal-probe-artifact",
		"-probe-tf-type", row.ResourceTypeName,
		"-probe-cfn-type", row.CloudFormationTypeName,
		"-probe-kind", string(kind),
		"-probe-schema", tmp.Name(),
		"-probe-prefix", cfg.prefix,
		"-probe-cache-dir", cfg.cacheDir,
		"-probe-services-path", cfg.servicesPath,
		"-probe-repo-root", cfg.repoRoot,
		"-probe-output-root", cfg.outputRoot,
	)
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOMEMLIMIT=%s", healProbeMemLimit))
	out, runErr := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("probe timed out after %s (possible runaway/recursive schema)", healProbeTimeout)
	}
	if runErr != nil {
		trimmed := strings.TrimSpace(string(out))
		var exitErr *exec.ExitError
		// A Go fatal runtime error (stack overflow, OOM inside the Go
		// runtime) exits 2 and self-reports a goroutine dump to stdout/stderr
		// — CombinedOutput captures it, so trimmed is non-empty and useful. A
		// signal-killed process (SIGKILL, an OS-level OOM kill) reports
		// ExitCode() == -1 and typically produces no output at all before
		// dying. Either way the probe is contained and reported, not crashed
		// into the parent; this only decides what text the proposal carries.
		if errors.As(runErr, &exitErr) && exitErr.ExitCode() > 0 && trimmed != "" {
			// A clean, reported generation failure: the probe printed the
			// error to stdout/stderr and exited non-zero on purpose (or a Go
			// fatal error that self-reported before dying).
			return errors.New(trimmed)
		}
		if trimmed == "" {
			trimmed = "no output captured"
		}
		return fmt.Errorf("probe crashed (signal/OOM/no self-reported error), output: %s: %w", trimmed, runErr)
	}
	return nil
}

// healProbeTimeout and healProbeMemLimit bound one isolated -heal probe.
const (
	healProbeTimeout  = 30 * time.Second
	healProbeMemLimit = "512MiB"
)

// runHealProbeArtifact is the hidden subprocess entrypoint probeArtifact
// re-execs into. It rebuilds the plan for exactly one artifact from flags,
// generates it (writing nothing to the real output tree — only the compile
// gate check below touches it, transiently, via buildOnce's own
// overlay-then-revert), and, if generation succeeds, runs that one artifact
// through the compile gate (generation-punchlist.md item 1) before reporting
// success — so a "lift" proposal is trustworthy against both stages a real
// -update run would have to pass, not just generation. repoRoot/outputRoot
// are required for the compile gate step; if either is empty (a caller that
// only wants the generation-only check, or an older binary's flag surface),
// the compile gate step is skipped rather than erroring, so probing a type
// with a kind that has no real destination concept yet degrades safely.
// Reports success/failure via exit code — never printed as a bigdiffer
// report, since this process only exists to be probed by its parent.
func runHealProbeArtifact(tfType, cfnType, kindFlag, schemaPath, prefix, cacheDir, servicesPath, repoRoot, outputRoot string) error {
	row := resourceRow{
		ResourceTypeName:         tfType,
		CloudFormationTypeName:   cfnType,
		CloudFormationSchemaPath: schemaPath,
	}
	cfg := config{prefix: prefix, cacheDir: cacheDir, servicesPath: servicesPath, repoRoot: repoRoot, outputRoot: outputRoot}

	p, err := generationPlan(row, cfg.prefix, cfg.cacheDir)
	if err != nil {
		return fmt.Errorf("plan: %w", err)
	}
	kind := artifactKind(kindFlag)
	for _, a := range p.artifacts {
		if a.kind != kind {
			continue
		}
		ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
		code, _, genErr := generateArtifact(ui, cfg, p, a)
		if genErr != nil {
			return genErr
		}
		if cfg.repoRoot == "" || cfg.outputRoot == "" {
			return nil // caller didn't ask for the compile gate step
		}
		dest := filepath.Join(cfg.outputRoot, a.pathSuffix, a.codeFile)
		ok, buildErrs, buildErr := buildOnce(cfg.repoRoot, map[string][]byte{dest: code})
		if buildErr != nil {
			return fmt.Errorf("compile gate: %w", buildErr)
		}
		if !ok {
			return fmt.Errorf("generates cleanly but fails the compile gate: %s", formatBuildErrors(buildErrs))
		}
		return nil
	}
	return fmt.Errorf("artifact %s not derivable from this row", kind)
}

// freezeProposal handles a frozen_since with no attributable per-artifact
// cause (e.g. every artifact is suppressed already with its own reason, or
// none are suppressed at all — the type itself was withdrawn or held back).
func freezeProposal(cfn, label, comment string) healProposal {
	base := healProposal{cfn: cfn, label: label, kind: "frozen_since"}
	return commentOrUnknown(base, comment)
}

// commentOrUnknown is step 4 of the heal probe (suppressed-and-frozen.md): if
// a free-form "# Suppression Reason:" comment exists, migrate its text into a
// manual: reason; otherwise the row still needs a human look.
func commentOrUnknown(base healProposal, comment string) healProposal {
	base.action = "reason"
	if comment != "" {
		base.reason = formatReason(reasonManual, comment)
		return base
	}
	base.reason = formatReason(reasonUnknown, "no schema cached and no existing comment to migrate; needs a human look")
	return base
}

func allReasonProposals(proposals []healProposal) int {
	n := 0
	for _, p := range proposals {
		if p.action == "reason" {
			n++
		}
	}
	return n
}

// suppressionCommentRE matches a "# Suppression Reason[:]" lead line, case-
// insensitive, optionally with inline text after the colon.
var suppressionCommentRE = regexp.MustCompile(`(?i)^\s*#\s*suppression reason:?\s*(.*)$`)

// suppressionComment extracts the free-form "# Suppression Reason:" comment
// text from a block's raw text (which may span multiple # lines), if present.
func suppressionComment(text string) string {
	lines := strings.Split(text, "\n")
	for i, ln := range lines {
		m := suppressionCommentRE.FindStringSubmatch(ln)
		if m == nil {
			continue
		}
		var parts []string
		if strings.TrimSpace(m[1]) != "" {
			parts = append(parts, strings.TrimSpace(m[1]))
		}
		for j := i + 1; j < len(lines); j++ {
			trimmed := strings.TrimSpace(lines[j])
			if !strings.HasPrefix(trimmed, "#") {
				break
			}
			cont := strings.TrimSpace(strings.TrimPrefix(trimmed, "#"))
			if cont == "" {
				break
			}
			parts = append(parts, cont)
		}
		return strings.Join(parts, " ")
	}
	return ""
}

func writeHealReport(needsReason int, proposals []healProposal) {
	fmt.Fprintf(os.Stderr, "== bigdiffer -heal report ==\n")
	fmt.Fprintf(os.Stderr, "rows needing a reason: %d\n", needsReason)
	fmt.Fprintf(os.Stderr, "proposals: %d\n", len(proposals))
	for _, p := range proposals {
		kind := p.kind
		if kind == "" {
			kind = "type"
		}
		switch p.action {
		case "lift":
			fmt.Fprintf(os.Stderr, "  ~ %s (%s) [%s]: %s\n", p.cfn, p.label, kind, p.reason)
		default:
			fmt.Fprintf(os.Stderr, "  + %s (%s) [%s]: suppression_reason = %q\n", p.cfn, p.label, kind, p.reason)
		}
	}
	fmt.Fprintln(os.Stderr, "Nothing above was written; review and apply by hand.")
}
