// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"encoding/json"
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

// healProposal is one row -recheck has something to say about. -recheck never
// mutates all_schemas.hcl; every proposal is reported for a human to apply,
// matching every other bigdiffer policy decision
// (contributing/docs/suppressed-and-frozen.md, "-recheck: re-probe and fill gaps").
type healProposal struct {
	cfn      string
	label    string
	kind     string         // which artifact this proposal is about, or "" for the freeze itself
	field    string         // the reason attribute this proposal fills, e.g. suppression_reason_plural_data_source
	action   string         // "lift" (generation now succeeds) or "reason" (a reason was determined/migrated)
	reason   string         // the proposed reason value
	category reasonCategory // the proposed reason's category, "" for a "lift" action
}

// healFact is one suppressed artifact or the freeze, named by its flag/date
// field, its own reason field, and (for artifacts) its artifactKind — the
// four independent facts a row can carry (item 9b). runRecheck only re-probes a
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

// needsHealing reports whether a fact belongs in -recheck's current scope.
// By default (all false) that scope is the reason-less/unknown backlog: an
// active fact whose own reason is still empty or tagged unknown. With all
// set (the -recheck-all flag), the scope widens to every active fact
// regardless of its existing reason — including one with a real, specific
// reason already recorded — so a year-old judgment call can be revisited on
// demand (§0 story 5) rather than only ever a still-open one.
func (f healFact) needsHealing(all bool) bool {
	if !f.active {
		return false
	}
	if all {
		return true
	}
	return f.reason == "" || strings.HasPrefix(f.reason, string(reasonUnknown)+":")
}

// runRecheck re-probes every active fact in scope and proposes a
// reclassification. By default that scope is the reason-less/unknown
// backlog; with all set (the -recheck-all flag), every active fact is
// in scope regardless of its existing reason, so a year-old suppress/freeze
// decision — even one with a specific, already-recorded reason — can be
// revisited on demand (§0 story 5). Offline except for reading the
// committed schema cache; never touches AWS and never writes
// all_schemas.hcl.
func runRecheck(allSchemasPath string, all bool) error {
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

	if all {
		stepf("Re-probing every suppressed/frozen fact, regardless of its existing reason…")
	} else {
		stepf("Re-probing suppressed/frozen facts with no recorded reason…")
	}
	var proposals []healProposal
	needsReason := 0
	for _, row := range rows {
		var pending []healFact
		for _, f := range healFactsFor(row) {
			if f.needsHealing(all) {
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

	writeHealReport(needsReason, proposals, all)
	return nil
}

// healRow probes exactly the row's in-scope facts (pending — reason-less/
// unknown by default, or every active fact under -recheck-all) and returns a
// proposal per fact it can say something about. It never mutates row or the
// overlay. multiPending is true when the row has more than one in-scope
// fact this run — used to phrase a migrated free-form comment as a shared
// candidate rather than a confirmed per-fact reason, since the same
// row-level comment text cannot be assumed to describe more than one fact
// (contributing/docs/suppressed-and-frozen.md, "-recheck: re-probe and fill
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
			out = append(out, freezeProposal(row, f, comment, multiPending))
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
			base.category = reasonStructural
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
				base.category = reasonBuildFailed
				base.reason = formatReason(reasonBuildFailed, firstLine(gateFailure.detail))
			} else {
				base.category = reasonGenerationFailed
				base.reason = formatReason(reasonGenerationFailed, firstLine(err.Error()))
			}
			return base
		}
	}

	return commentOrUnknown(base, f.reason, comment, multiPending)
}

// probeArtifact regenerates one artifact from schema bytes staged to a temp
// file, without writing any output — a read-only re-probe. It never runs the
// owned engine in this process: some suppressed rows are suppressed *because*
// their schema is recursive (e.g. the free-form "Recursive Attribute
// Definitions" comments, issue #95), and codegen.Emitter has no recursion-depth
// guard, so re-probing one in-process is a confirmed stack overflow — a fatal,
// unrecoverable crash that would take the whole -recheck run down with it. Instead
// probeArtifact re-execs bigdiffer itself into the hidden -recheck-probe-artifact
// mode (runRecheckProbeArtifact) under a timeout and a soft memory cap, so a crash
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
	_, _, err := runIsolatedProbe(bin, cfg, row, kind, schema, false)
	return err
}

// generateArtifactIsolated regenerates one artifact the same way probeArtifact
// does — re-exec'd into the hidden -recheck-probe-artifact subprocess mode,
// under the same timeout and memory cap, so a schema nobody has ever safely
// generated before (a brand-new CloudFormation type discovered by -sync,
// with no suppression entry yet to protect it — unlike -recheck's backlog,
// which by construction only probes rows already known to the overlay) can
// never take the parent -sync/-reconcile/-check process down with it.
//
// Unlike probeArtifact/probeArtifactWithBinary (a read-only pass/fail check
// for -recheck's own purposes), this also asks the subprocess to write the
// generated code/test bytes to two temp files and reads them back on
// success, so a real caller like refreshCandidate that needs to stage the
// actual artifact — not just learn whether generation still fails — can use
// the same isolation. row's suppress_*/path_aware_attribute_names fields are
// passed through as flags so the child's generationPlan reconstructs the
// exact same artifact shape (file names, ListResource-or-not) the parent's
// own plan already computed, not a fresh guess from bare tfType/cfnType.
var generateArtifactIsolated = func(cfg config, row resourceRow, kind artifactKind, schema []byte) (code, test []byte, err error) {
	self, err := os.Executable()
	if err != nil {
		return nil, nil, fmt.Errorf("locating bigdiffer binary to generate in isolation: %w", err)
	}
	return generateArtifactIsolatedWithBinary(self, cfg, row, kind, schema)
}

// generateArtifactIsolatedWithBinary is generateArtifactIsolated's
// implementation, parameterized on the binary to re-exec — see
// probeArtifactWithBinary's identical rationale (os.Executable() resolves to
// the `go test` binary, not bigdiffer's real main(), under tests).
func generateArtifactIsolatedWithBinary(bin string, cfg config, row resourceRow, kind artifactKind, schema []byte) (code, test []byte, err error) {
	return runIsolatedProbe(bin, cfg, row, kind, schema, true)
}

// generateCandidateArtifacts regenerates every artifact in row's plan in one
// subprocess call (runRecheckProbeArtifact's manifest mode) instead of one
// per artifact. A crash still only takes down that one candidate, never the
// parent, but now also marks any sibling artifact sharing that subprocess
// as failed, even one that would have generated cleanly alone.
// reconcileListResource's rare single-artifact correction uses
// generateArtifactIsolated directly instead, since it only regenerates one
// artifact.
//
// artifacts (the parent's own plan.artifacts) is used to know which kinds
// to expect back and to attach each manifest entry to its full genArtifact.
// The subprocess still recomputes its own plan via generationPlan; what
// keeps the two in agreement is that every input generationPlan reads is
// forwarded identically on both sides.
var generateCandidateArtifacts = func(cfg config, row resourceRow, artifacts []genArtifact, schema []byte) ([]genResult, error) {
	self, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("locating bigdiffer binary to generate in isolation: %w", err)
	}
	return generateCandidateArtifactsWithBinary(self, cfg, row, artifacts, schema)
}

// generateCandidateArtifactsWithBinary is generateCandidateArtifacts's
// implementation, parameterized on the binary to re-exec — see
// probeArtifactWithBinary's identical rationale (os.Executable() resolves to
// the `go test` binary, not bigdiffer's real main(), under tests).
func generateCandidateArtifactsWithBinary(bin string, cfg config, row resourceRow, artifacts []genArtifact, schema []byte) ([]genResult, error) {
	tmp, err := os.CreateTemp("", "bigdiffer-recheck-*.json")
	if err != nil {
		return nil, err
	}
	defer func() { _ = os.Remove(tmp.Name()) }()
	if _, err := tmp.Write(schema); err != nil {
		return nil, err
	}
	if err := tmp.Close(); err != nil {
		return nil, err
	}

	manifestTmp, err := os.CreateTemp("", "bigdiffer-probe-manifest-*")
	if err != nil {
		return nil, err
	}
	manifestPath := manifestTmp.Name()
	_ = manifestTmp.Close()
	_ = os.Remove(manifestPath) // runProbeCandidate creates it fresh; a stray empty file would confuse a crash-vs-no-output check
	defer func() {
		_ = os.Remove(manifestPath)
		for _, a := range artifacts {
			_ = os.Remove(manifestPath + "." + string(a.kind) + ".code")
			_ = os.Remove(manifestPath + "." + string(a.kind) + ".test")
		}
	}()

	args := []string{
		"-recheck-probe-artifact",
		"-probe-tf-type", row.ResourceTypeName,
		"-probe-cfn-type", row.CloudFormationTypeName,
		"-probe-schema", tmp.Name(),
		"-probe-prefix", cfg.prefix,
		"-probe-cache-dir", cfg.cacheDir,
		"-probe-services-path", cfg.servicesPath,
		"-probe-out-manifest", manifestPath,
	}
	// This mode always wants the real artifact bytes, so the row's real
	// suppress_* flags are always forwarded (unlike probeArtifact's
	// recheck-style use, which deliberately ignores them).
	if row.SuppressResourceGeneration {
		args = append(args, "-probe-suppress-resource")
	}
	if row.SuppressSingularDataSourceGeneration {
		args = append(args, "-probe-suppress-singular")
	}
	if row.SuppressPluralDataSourceGeneration {
		args = append(args, "-probe-suppress-plural")
	}
	if row.PathAwareAttributeNames {
		args = append(args, "-probe-path-aware-names")
	}

	ctx, cancel := context.WithTimeout(context.Background(), healProbeTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOMEMLIMIT=%s", healProbeMemLimit))
	out, runErr := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded || runErr != nil {
		// No manifest to read (timeout, or the subprocess died outright) —
		// every artifact in this candidate is reported failed with the same
		// signal.
		var crashMsg string
		switch {
		case ctx.Err() == context.DeadlineExceeded:
			crashMsg = fmt.Sprintf("probe timed out after %s (possible runaway/recursive schema)", healProbeTimeout)
		default:
			trimmed := strings.TrimSpace(string(out))
			if trimmed == "" {
				trimmed = "no output captured"
			}
			// Not truncated — a real stack overflow's self-report can be
			// tens of KB, but every overlay-write site truncates to
			// firstLine(err.Error()) before persisting, so this only
			// affects an in-memory value a caller might log.
			crashMsg = fmt.Sprintf("probe crashed (signal/OOM/no self-reported error), output: %s", trimmed)
		}
		results := make([]genResult, len(artifacts))
		for i, a := range artifacts {
			results[i] = genResult{a: a, err: errors.New(crashMsg)}
		}
		return results, nil
	}

	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("reading probe manifest: %w", err)
	}
	var manifest probeCandidateManifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return nil, fmt.Errorf("decoding probe manifest: %w", err)
	}
	byKind := make(map[artifactKind]probeManifestArtifact, len(manifest.Artifacts))
	for _, ma := range manifest.Artifacts {
		byKind[ma.Kind] = ma
	}
	results := make([]genResult, len(artifacts))
	for i, a := range artifacts {
		ma, ok := byKind[a.kind]
		if !ok {
			results[i] = genResult{a: a, err: fmt.Errorf("probe manifest has no entry for artifact %s", a.kind)}
			continue
		}
		if ma.Err != "" {
			results[i] = genResult{a: a, err: errors.New(ma.Err)}
			continue
		}
		// An unreadable code/test file (vanishingly unlikely — just written
		// by the subprocess) is folded into a generation failure below
		// rather than a separate infra-error class.
		code, err := os.ReadFile(ma.CodePath)
		if err != nil {
			results[i] = genResult{a: a, err: fmt.Errorf("reading manifest code output: %w", err)}
			continue
		}
		test, err := os.ReadFile(ma.TestPath)
		if err != nil {
			results[i] = genResult{a: a, err: fmt.Errorf("reading manifest test output: %w", err)}
			continue
		}
		results[i] = genResult{a: a, code: code, test: test}
	}
	return results, nil
}

// runIsolatedProbe is the shared core of probeArtifactWithBinary and
// generateArtifactIsolatedWithBinary: stage the schema to a temp file, re-exec
// bin into -recheck-probe-artifact under a timeout and memory cap, and
// classify the result. wantBytes selects whether the two -probe-out-* flags
// are set and the resulting files read back (generateArtifactIsolated's
// case) or omitted entirely (probeArtifact's read-only case, which never
// needs the generated bytes and would otherwise pay for two needless temp
// files on every one of -recheck's backlog probes).
func runIsolatedProbe(bin string, cfg config, row resourceRow, kind artifactKind, schema []byte, wantBytes bool) (code, test []byte, err error) {
	tmp, err := os.CreateTemp("", "bigdiffer-recheck-*.json")
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = os.Remove(tmp.Name()) }()
	if _, err := tmp.Write(schema); err != nil {
		return nil, nil, err
	}
	if err := tmp.Close(); err != nil {
		return nil, nil, err
	}

	args := []string{
		"-recheck-probe-artifact",
		"-probe-tf-type", row.ResourceTypeName,
		"-probe-cfn-type", row.CloudFormationTypeName,
		"-probe-kind", string(kind),
		"-probe-schema", tmp.Name(),
		"-probe-prefix", cfg.prefix,
		"-probe-cache-dir", cfg.cacheDir,
		"-probe-services-path", cfg.servicesPath,
		"-probe-repo-root", cfg.repoRoot,
		"-probe-output-root", cfg.outputRoot,
	}
	// The row's real suppress_*/path_aware_attribute_names flags are only
	// passed through for generateArtifactIsolated's case (wantBytes: it
	// needs the exact same artifact shape refreshCandidate's own plan
	// already computed). probeArtifact's read-only recheck case deliberately
	// does not: -recheck's whole purpose is asking "would this kind still
	// fail if attempted" for a row that may already be suppressed for that
	// very kind — passing the real suppress flag through would make
	// generationPlan drop the artifact entirely and fail every re-probe of
	// an already-suppressed row with "not derivable from this row" instead
	// of actually re-attempting generation.
	if wantBytes {
		if row.SuppressResourceGeneration {
			args = append(args, "-probe-suppress-resource")
		}
		if row.SuppressSingularDataSourceGeneration {
			args = append(args, "-probe-suppress-singular")
		}
		if row.SuppressPluralDataSourceGeneration {
			args = append(args, "-probe-suppress-plural")
		}
		if row.PathAwareAttributeNames {
			args = append(args, "-probe-path-aware-names")
		}
	}

	var codePath, testPath string
	if wantBytes {
		codeTmp, err := os.CreateTemp("", "bigdiffer-probe-code-*")
		if err != nil {
			return nil, nil, err
		}
		codePath = codeTmp.Name()
		_ = codeTmp.Close()
		defer func() { _ = os.Remove(codePath) }()

		testTmp, err := os.CreateTemp("", "bigdiffer-probe-test-*")
		if err != nil {
			return nil, nil, err
		}
		testPath = testTmp.Name()
		_ = testTmp.Close()
		defer func() { _ = os.Remove(testPath) }()

		args = append(args, "-probe-out-code", codePath, "-probe-out-test", testPath)
	}

	ctx, cancel := context.WithTimeout(context.Background(), healProbeTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOMEMLIMIT=%s", healProbeMemLimit))
	out, runErr := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return nil, nil, fmt.Errorf("probe timed out after %s (possible runaway/recursive schema)", healProbeTimeout)
	}
	if runErr != nil {
		trimmed := strings.TrimSpace(string(out))
		var exitErr *exec.ExitError
		if errors.As(runErr, &exitErr) && exitErr.ExitCode() == exitCodeBuildGateFailed {
			// The probe generated cleanly but the compile gate rejected it —
			// runRecheckProbeArtifact's message is exactly buildGateFailure's
			// Error() text, printed verbatim to stdout/stderr. Reconstruct
			// the typed error here so healArtifact can distinguish this from
			// a plain generation failure and tag the proposal build_failed
			// instead of generation_failed.
			return nil, nil, &buildGateFailure{detail: strings.TrimPrefix(trimmed, "generates cleanly but fails the compile gate: ")}
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
			return nil, nil, errors.New(trimmed)
		}
		if trimmed == "" {
			trimmed = "no output captured"
		}
		return nil, nil, fmt.Errorf("probe crashed (signal/OOM/no self-reported error), output: %s: %w", trimmed, runErr)
	}
	if !wantBytes {
		return nil, nil, nil
	}
	code, err = os.ReadFile(codePath)
	if err != nil {
		return nil, nil, fmt.Errorf("reading probe code output: %w", err)
	}
	test, err = os.ReadFile(testPath)
	if err != nil {
		return nil, nil, fmt.Errorf("reading probe test output: %w", err)
	}
	return code, test, nil
}

// healProbeTimeout and healProbeMemLimit bound one isolated -recheck probe.
const (
	healProbeTimeout  = 30 * time.Second
	healProbeMemLimit = "512MiB"
)

// exitCodeBuildGateFailed is runRecheckProbeArtifact's exit code for a
// buildGateFailure specifically — distinct from the generic exit 1 a plain
// generation error uses. probeArtifactWithBinary keys on this to reconstruct
// a *buildGateFailure across the process boundary, since the parent can't
// see the child's Go error values directly, only its exit code and output.
const exitCodeBuildGateFailed = 3

// runRecheckProbeArtifact is the hidden subprocess entrypoint probeArtifact
// re-execs into. It rebuilds the plan for exactly one artifact from flags,
// generates it (writing nothing to the real output tree — only the compile
// gate check below touches it, transiently, via buildOnce's own
// overlay-then-revert), and, if generation succeeds, runs that one artifact
// through the compile gate (bigdiffer-design.md §6) before reporting
// success — so a "lift" proposal is trustworthy against both stages a real
// -sync run would have to pass, not just generation. repoRoot/outputRoot
// are required for the compile gate step; if either is empty (a caller that
// only wants the generation-only check, or an older binary's flag surface),
// the compile gate step is skipped rather than erroring, so probing a type
// with a kind that has no real destination concept yet degrades safely.
// Reports success/failure via exit code — never printed as a bigdiffer
// report, since this process only exists to be probed by its parent.
//
// suppressResource/suppressSingular/suppressPlural/pathAwareNames mirror the
// same-named resourceRow fields so the child's generationPlan matches the
// parent's plan exactly (file names, ListResource or not). outCodePath/
// outTestPath, when non-empty, write the generated code/test bytes on
// success (probeArtifact's own recheck use only needs pass/fail, so leaves
// these empty).
//
// outManifestPath, when non-empty, switches to whole-candidate mode:
// kindFlag/outCodePath/outTestPath are ignored, every artifact in the row's
// plan is generated in-band in this process, and results are described in a
// JSON probeCandidateManifest written to outManifestPath. This lets
// generateCandidateArtifacts isolate a whole candidate (2-3 artifacts) in
// one subprocess instead of one per artifact.
func runRecheckProbeArtifact(tfType, cfnType, kindFlag, schemaPath, prefix, cacheDir, servicesPath, repoRoot, outputRoot string, suppressResource, suppressSingular, suppressPlural, pathAwareNames bool, outCodePath, outTestPath, outManifestPath string) error {
	row := resourceRow{
		ResourceTypeName:                     tfType,
		CloudFormationTypeName:               cfnType,
		CloudFormationSchemaPath:             schemaPath,
		SuppressResourceGeneration:           suppressResource,
		SuppressSingularDataSourceGeneration: suppressSingular,
		SuppressPluralDataSourceGeneration:   suppressPlural,
		PathAwareAttributeNames:              pathAwareNames,
	}
	cfg := config{prefix: prefix, cacheDir: cacheDir, servicesPath: servicesPath, repoRoot: repoRoot, outputRoot: outputRoot}

	p, err := generationPlan(row, cfg.prefix, cfg.cacheDir)
	if err != nil {
		return fmt.Errorf("plan: %w", err)
	}

	if outManifestPath != "" {
		return runProbeCandidate(cfg, p, outManifestPath)
	}

	kind := artifactKind(kindFlag)
	for _, a := range p.artifacts {
		if a.kind != kind {
			continue
		}
		ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
		code, test, genErr := generateArtifact(ui, cfg, p, a)
		if genErr != nil {
			return genErr
		}
		if cfg.repoRoot != "" && cfg.outputRoot != "" {
			dest := filepath.Join(cfg.outputRoot, a.pathSuffix, a.codeFile)
			// context.Background(), not a derived timeout: this whole process is
			// already killed wholesale by the parent's exec.CommandContext once
			// healProbeTimeout elapses (probeArtifact), which also tears down
			// this one buildOnce call along with everything else in the process.
			ok, buildErrs, buildErr := buildOnce(context.Background(), cfg.repoRoot, map[string][]byte{dest: code})
			if buildErr != nil {
				return fmt.Errorf("compile gate: %w", buildErr)
			}
			if !ok {
				return &buildGateFailure{detail: formatBuildErrors(relativizeBuildErrors(buildErrs, cfg.repoRoot))}
			}
		}
		if outCodePath != "" {
			if err := os.WriteFile(outCodePath, code, filePerm); err != nil {
				return fmt.Errorf("writing probe code output: %w", err)
			}
		}
		if outTestPath != "" {
			if err := os.WriteFile(outTestPath, test, filePerm); err != nil {
				return fmt.Errorf("writing probe test output: %w", err)
			}
		}
		return nil
	}
	return fmt.Errorf("artifact %s not derivable from this row", kind)
}

// probeManifestArtifact is one artifact's outcome in a probeCandidateManifest
// — the whole-candidate mode's wire format back to the parent process.
// CodePath/TestPath are only meaningful when Err == "". Err carries only
// firstLine(original error), not a raw crash dump.
type probeManifestArtifact struct {
	Kind     artifactKind `json:"kind"`
	CodePath string       `json:"codePath,omitempty"`
	TestPath string       `json:"testPath,omitempty"`
	Err      string       `json:"err,omitempty"`
}

// probeCandidateManifest is runProbeCandidate's JSON report, one entry per
// artifact in the row's plan. A whole-process crash means there is no
// manifest to read at all; the parent (generateCandidateArtifacts) treats
// that as every artifact failed.
type probeCandidateManifest struct {
	Artifacts []probeManifestArtifact `json:"artifacts"`
}

// runProbeCandidate is runRecheckProbeArtifact's whole-candidate mode:
// generate every artifact in p in this process, one artifact's error
// captured in-band and never stopping the others, and write the result to
// manifestPath. Never runs the compile gate — that's pipeline.go's job over
// the whole staged batch.
func runProbeCandidate(cfg config, p plan, manifestPath string) error {
	manifest := probeCandidateManifest{Artifacts: make([]probeManifestArtifact, len(p.artifacts))}
	ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
	for i, a := range p.artifacts {
		code, test, genErr := generateArtifact(ui, cfg, p, a)
		if genErr != nil {
			manifest.Artifacts[i] = probeManifestArtifact{Kind: a.kind, Err: firstLine(genErr.Error())}
			continue
		}
		codePath := manifestPath + "." + string(a.kind) + ".code"
		testPath := manifestPath + "." + string(a.kind) + ".test"
		if err := os.WriteFile(codePath, code, filePerm); err != nil {
			return fmt.Errorf("writing manifest code output for %s: %w", a.kind, err)
		}
		if err := os.WriteFile(testPath, test, filePerm); err != nil {
			return fmt.Errorf("writing manifest test output for %s: %w", a.kind, err)
		}
		manifest.Artifacts[i] = probeManifestArtifact{Kind: a.kind, CodePath: codePath, TestPath: testPath}
	}
	out, err := json.Marshal(manifest)
	if err != nil {
		return fmt.Errorf("encoding probe manifest: %w", err)
	}
	if err := os.WriteFile(manifestPath, out, filePerm); err != nil {
		return fmt.Errorf("writing probe manifest: %w", err)
	}
	return nil
}

// buildGateFailure is runRecheckProbeArtifact's distinct signal that an artifact
// generated cleanly but was rejected by the compile gate — as opposed to a
// plain error, which means generation itself failed. probeArtifactWithBinary
// maps this to exitCodeBuildGateFailed so the parent process (running in a
// separate binary, so it cannot see this type directly) can still tell the
// two stages apart and healArtifact can tag the proposal's reason
// accordingly (build_failed vs generation_failed) — the whole point of
// wiring the compile gate into -recheck.
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

// freezeProposal handles a frozen_since fact — the freeze is an orthogonal,
// schema-level fact (item 9b), never derived from an artifact's suppression
// reason, so it always falls through to the same comment-migration/unknown
// fallback every other in-scope fact uses (commentOrUnknown), carrying its
// own existing reason (f.reason) through so a real, already-recorded reason
// is never overwritten with a guess when there is no schema to re-probe with
// (only reachable under -recheck-all, since the default scope excludes
// facts with a real reason already).
func freezeProposal(row resourceRow, f healFact, comment string, multiPending bool) healProposal {
	base := healProposal{cfn: row.CloudFormationTypeName, label: row.ResourceTypeName, kind: "", field: attrFrozenReason}
	return commentOrUnknown(base, f.reason, comment, multiPending)
}

// commentOrUnknown is step 4 of the recheck probe (suppressed-and-frozen.md): if
// the fact already carries a real, specific reason (only reachable at all
// under -recheck-all, since the default scope excludes these), that reason
// is kept — there is no schema to re-probe with, so there is nothing to
// weigh against it, and a comment-migration guess must never overwrite a
// real, already-recorded answer. The appended note differs for the freeze
// (base.field == attrFrozenReason): a freeze is a schema-level fact, never
// re-probed regardless of cache state (that is what "frozen" means), so it
// gets its own wording rather than the artifact-oriented "re-run once the
// schema cache has this type" — that phrasing would misleadingly imply a
// future run could resolve a freeze differently, when nothing does without
// a human lifting it by hand. Otherwise: if a free-form
// "# Suppression Reason:" comment exists, migrate its text into a manual:
// reason; otherwise the row still needs a human look.
//
// When multiPending is true (more than one of the row's facts is in scope
// this run), the same row-level comment text is offered identically to
// each — it is not auto-assigned as a confirmed per-fact reason, since a
// single comment written for one artifact does not necessarily explain a
// different artifact's suppression or the freeze
// (contributing/docs/suppressed-and-frozen.md, "-recheck: re-probe and fill
// gaps"). The proposal text is worded as a shared candidate for a human
// to assign, edit, or reject per field, rather than a confirmed fact.
func commentOrUnknown(base healProposal, existingReason, comment string, multiPending bool) healProposal {
	base.action = "reason"
	if existingReason != "" && !strings.HasPrefix(existingReason, string(reasonUnknown)+":") {
		if base.field == attrFrozenReason {
			// The freeze is a schema-level fact, not a per-artifact
			// generation attempt (healFactsFor) — there is no schema cache
			// state that would ever change this note's applicability, unlike
			// an artifact fact where a schema simply hasn't been cached yet.
			// A frozen row's schema is deliberately never refreshed (that is
			// what "frozen" means), so "re-run once the cache has this type"
			// would be actively misleading here: it implies a future run
			// might resolve this differently, when nothing about a freeze
			// changes without a human lifting it by hand.
			base.reason = existingReason + " (no schema re-probe applies to a freeze; existing reason kept as-is — verify by hand)"
			return base
		}
		base.reason = existingReason + " (no cached schema to re-probe against; existing reason kept as-is — re-run once the schema cache has this type, or verify by hand)"
		return base
	}
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

func writeHealReport(needsReason int, proposals []healProposal, all bool) {
	fmt.Fprintf(os.Stderr, "== bigdiffer -recheck report ==\n")
	if all {
		fmt.Fprintf(os.Stderr, "facts re-probed (every active fact, -recheck-all): %d\n", needsReason)
	} else {
		fmt.Fprintf(os.Stderr, "facts needing a reason: %d\n", needsReason)
	}
	fmt.Fprintf(os.Stderr, "proposals: %d\n", len(proposals))
	for _, p := range proposals {
		switch p.action {
		case "lift":
			fmt.Fprintf(os.Stderr, "  ~ %s (%s) [%s]: %s\n", p.cfn, p.label, p.field, p.reason)
		default:
			fmt.Fprintf(os.Stderr, "  + %s (%s) [%s]: %s = %q\n", p.cfn, p.label, p.field, p.field, p.reason)
			if isIssueWorthy(p.category) {
				fmt.Fprintf(os.Stderr, "    consider opening a GitHub issue for this %s (%s)\n", p.category, issueTitle(p))
			}
		}
	}
	fmt.Fprintln(os.Stderr, "Nothing above was written; review and apply by hand.")
}

// isIssueWorthy reports whether a proposal's reason category warrants
// recommending a GitHub issue (contributing/docs/suppressed-and-frozen.md,
// "GitHub issues: when to file, and what to say"): a real defect worth
// tracking (generation_failed, build_failed), not upstream/structural, not a
// human's own manual call, and not an unresolved unknown still needing
// triage.
func isIssueWorthy(category reasonCategory) bool {
	return category == reasonGenerationFailed || category == reasonBuildFailed
}

// issueTitle renders a short "<type> <artifact>" label for the recommended
// issue's title — the captured error detail is already printed on the line
// above (part of p.reason), so this only adds what that line doesn't already
// say. No URL, no stub, no auto-filing: bigdiffer only says an issue is
// warranted; the human files it and records the URL by hand (appending
// " (issue: <URL>)" to the reason value already proposed above).
func issueTitle(p healProposal) string {
	if p.kind != "" {
		return p.cfn + " " + p.kind
	}
	return p.cfn
}
