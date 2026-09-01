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
	field  string // the reason attribute this proposal fills, e.g. suppression_reason_plural_data_source
	action string // "lift" (generation now succeeds) or "reason" (a reason was determined/migrated)
	reason string // the proposed reason value
}

// healFact is one suppressed artifact or the freeze, named by its flag/date
// field, its own reason field, and (for artifacts) its artifactKind — the
// four independent facts a row can carry (item 9b). runHeal only re-probes a
// fact whose own reason is still empty or tagged unknown; a row can have a
// real reason recorded for one fact while another of its facts is still in
// the backlog.
type healFact struct {
	kind   artifactKind // "" for the freeze
	field  string       // the reason attribute to fill
	active bool         // is this fact true on the row (suppressed / frozen)
	reason string       // the fact's own current reason text
}

func healFactsFor(row resourceRow) []healFact {
	return []healFact{
		{kind: artifactResource, field: attrSuppressionReasonResource, active: row.SuppressResourceGeneration, reason: row.SuppressionReasonResource},
		{kind: artifactSingularDataSource, field: attrSuppressionReasonSingular, active: row.SuppressSingularDataSourceGeneration, reason: row.SuppressionReasonSingularDataSource},
		{kind: artifactPluralDataSource, field: attrSuppressionReasonPlural, active: row.SuppressPluralDataSourceGeneration, reason: row.SuppressionReasonPluralDataSource},
		{field: attrFrozenReason, active: row.FrozenSince != "", reason: row.FrozenReason},
	}
}

// needsHealing is true for an active fact whose own reason is still empty or
// tagged unknown — the reason-less/unknown backlog this fact belongs to.
func (f healFact) needsHealing() bool {
	return f.active && (f.reason == "" || strings.HasPrefix(f.reason, string(reasonUnknown)+":"))
}

// runHeal re-probes every suppressed artifact or freeze that has no reason (or
// one tagged unknown) and proposes a reclassification. Offline except for
// reading the committed schema cache; never touches AWS and never writes
// all_schemas.hcl.
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

	stepf("Re-probing suppressed/frozen facts with no recorded reason…")
	var proposals []healProposal
	needsReason := 0
	for _, row := range rows {
		var pending []healFact
		for _, f := range healFactsFor(row) {
			if f.needsHealing() {
				pending = append(pending, f)
			}
		}
		if len(pending) == 0 {
			continue
		}
		needsReason += len(pending)
		proposals = append(proposals, healRow(cfg, row, pending, commentByCFN[row.CloudFormationTypeName])...)
	}
	sort.Slice(proposals, func(i, j int) bool {
		if proposals[i].cfn != proposals[j].cfn {
			return proposals[i].cfn < proposals[j].cfn
		}
		return proposals[i].field < proposals[j].field
	})

	writeHealReport(needsReason, proposals)
	return nil
}

// healRow probes exactly the row's still-reason-less facts (pending) and
// returns a proposal per fact it can say something about. It never mutates
// row or the overlay. multiPending is true when the row has more than one
// still-reason-less fact — used to phrase a migrated free-form comment as a
// shared candidate rather than a confirmed per-fact reason, since the same
// row-level comment text cannot be assumed to describe more than one fact
// (contributing/docs/suppressed-and-frozen.md, "-heal: re-probe and fill
// gaps").
func healRow(cfg config, row resourceRow, pending []healFact, comment string) []healProposal {
	var out []healProposal
	multiPending := len(pending) > 1

	schema, schemaErr := os.ReadFile(row.CloudFormationSchemaPath)
	if schemaErr != nil {
		schema, schemaErr = os.ReadFile(schemaCachePath(cfg.cacheDir, row.CloudFormationTypeName))
	}

	for _, f := range pending {
		if f.field == attrFrozenReason {
			out = append(out, freezeProposal(row, comment, multiPending))
			continue
		}
		out = append(out, healArtifact(cfg, row, f, schema, schemaErr, comment, multiPending))
	}
	return out
}

// healArtifact probes one suppressed artifact: structural check first (plural
// only), then a real regeneration attempt, then a comment-migration fallback.
func healArtifact(cfg config, row resourceRow, f healFact, schema []byte, schemaErr error, comment string, multiPending bool) healProposal {
	base := healProposal{cfn: row.CloudFormationTypeName, label: row.ResourceTypeName, kind: string(f.kind), field: f.field}

	if f.kind == artifactPluralDataSource && schemaErr == nil {
		if !pluralSupported(string(schema)) {
			base.action = "reason"
			base.reason = formatReason(reasonStructural, "no list handler with zero required arguments")
			return base
		}
	}

	if schemaErr == nil {
		if err := probeArtifact(cfg, row, f.kind, schema); err == nil {
			base.action = "lift"
			base.reason = fmt.Sprintf(
				"generates cleanly now — propose lifting the suppression. To keep it suppressed and stop this proposal recurring, set %s instead (e.g. %q).",
				f.field, formatReason(reasonManual, "<why this stays suppressed>"))
			return base
		} else {
			base.action = "reason"
			var gateFailure *buildGateFailure
			if errors.As(err, &gateFailure) {
				base.reason = formatReason(reasonBuildFailed, firstLine(gateFailure.detail))
			} else {
				base.reason = formatReason(reasonGenerationFailed, firstLine(err.Error()))
			}
			return base
		}
	}

	return commentOrUnknown(base, comment, multiPending)
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
//
// A package-level var, not a plain func: os.Executable() resolves to the
// test binary under `go test` (a binary whose CLI is testing.Main, not
// bigdiffer's real main()), so healArtifact-level tests that need a real
// probe against a real built binary swap this var for the duration of the
// test and restore it afterward.
var probeArtifact = func(cfg config, row resourceRow, kind artifactKind, schema []byte) error {
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
		if errors.As(runErr, &exitErr) && exitErr.ExitCode() == exitCodeBuildGateFailed {
			// The probe generated cleanly but the compile gate rejected it —
			// runHealProbeArtifact's message is exactly buildGateFailure's
			// Error() text, printed verbatim to stdout/stderr. Reconstruct
			// the typed error here so healArtifact can distinguish this from
			// a plain generation failure and tag the proposal build_failed
			// instead of generation_failed.
			return &buildGateFailure{detail: strings.TrimPrefix(trimmed, "generates cleanly but fails the compile gate: ")}
		}
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

// exitCodeBuildGateFailed is runHealProbeArtifact's exit code for a
// buildGateFailure specifically — distinct from the generic exit 1 a plain
// generation error uses. probeArtifactWithBinary keys on this to reconstruct
// a *buildGateFailure across the process boundary, since the parent can't
// see the child's Go error values directly, only its exit code and output.
const exitCodeBuildGateFailed = 3

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
			return &buildGateFailure{detail: formatBuildErrors(relativizeBuildErrors(buildErrs, cfg.repoRoot))}
		}
		return nil
	}
	return fmt.Errorf("artifact %s not derivable from this row", kind)
}

// buildGateFailure is runHealProbeArtifact's distinct signal that an artifact
// generated cleanly but was rejected by the compile gate — as opposed to a
// plain error, which means generation itself failed. probeArtifactWithBinary
// maps this to exitCodeBuildGateFailed so the parent process (running in a
// separate binary, so it cannot see this type directly) can still tell the
// two stages apart and healArtifact can tag the proposal's reason
// accordingly (build_failed vs generation_failed) — the whole point of
// wiring the compile gate into -heal.
type buildGateFailure struct{ detail string }

func (e *buildGateFailure) Error() string {
	return "generates cleanly but fails the compile gate: " + e.detail
}

// relativizeBuildErrors rewrites each buildError's absolute file path
// (buildError.file is always absolute, joined against repoRoot by
// parseBuildErrors) to be repo-root-relative for a tidier proposal detail —
// cosmetic only, does not affect attribution.
func relativizeBuildErrors(errs []buildError, repoRoot string) []buildError {
	out := make([]buildError, len(errs))
	for i, e := range errs {
		if rel, err := filepath.Rel(repoRoot, e.file); err == nil {
			e.file = rel
		}
		out[i] = e
	}
	return out
}

// freezeProposal handles a frozen_since with no reason of its own — the
// freeze is an orthogonal, schema-level fact (item 9b), never derived from an
// artifact's suppression reason, so it always falls through to the same
// comment-migration/unknown fallback every other reason-less fact uses.
func freezeProposal(row resourceRow, comment string, multiPending bool) healProposal {
	base := healProposal{cfn: row.CloudFormationTypeName, label: row.ResourceTypeName, kind: "", field: attrFrozenReason}
	return commentOrUnknown(base, comment, multiPending)
}

// commentOrUnknown is step 4 of the heal probe (suppressed-and-frozen.md): if
// a free-form "# Suppression Reason:" comment exists, migrate its text into a
// manual: reason; otherwise the row still needs a human look.
//
// When multiPending is true (more than one of the row's facts is still
// reason-less), the same row-level comment text is offered identically to
// each — it is not auto-assigned as a confirmed per-fact reason, since a
// single comment written for one artifact does not necessarily explain a
// different artifact's suppression or the freeze
// (contributing/docs/suppressed-and-frozen.md, "-heal: re-probe and fill
// gaps"). The proposal text is worded as a shared candidate for a human
// to assign, edit, or reject per field, rather than a confirmed fact.
func commentOrUnknown(base healProposal, comment string, multiPending bool) healProposal {
	base.action = "reason"
	if comment != "" {
		if multiPending {
			base.reason = formatReason(reasonManual, comment) +
				" (candidate: this comment is shared with other still-reason-less facts on this row; confirm or edit per field before applying)"
		} else {
			base.reason = formatReason(reasonManual, comment)
		}
		return base
	}
	base.reason = formatReason(reasonUnknown, "no schema cached and no existing comment to migrate; needs a human look")
	return base
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
	fmt.Fprintf(os.Stderr, "facts needing a reason: %d\n", needsReason)
	fmt.Fprintf(os.Stderr, "proposals: %d\n", len(proposals))
	for _, p := range proposals {
		switch p.action {
		case "lift":
			fmt.Fprintf(os.Stderr, "  ~ %s (%s) [%s]: %s\n", p.cfn, p.label, p.field, p.reason)
		default:
			fmt.Fprintf(os.Stderr, "  + %s (%s) [%s]: %s = %q\n", p.cfn, p.label, p.field, p.field, p.reason)
		}
	}
	fmt.Fprintln(os.Stderr, "Nothing above was written; review and apply by hand.")
}
