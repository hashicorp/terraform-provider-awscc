// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// The compile gate (generation-punchlist.md item 1, design:
// bigdiffer-design.md §6, "The compile gate"). Generated code only type-checks
// against the real module graph, so the gate builds in place: it overlays
// staged files onto their real destinations, runs `go build ./...` once from
// cfg.repoRoot, and unconditionally reverts the overlay — regardless of the
// build's outcome — before returning. The real tree is bit-for-bit unchanged
// by the gate; only compileFixpoint's caller (runUpdate) promotes for real,
// and only after the gate has gone green.

// buildOverlay records, for each real destination path touched by one
// overlay/build/revert cycle, the bytes to restore on revert — or nil if the
// path did not exist beforehand, meaning revert should remove it rather than
// restore it.
type buildOverlay struct {
	prior   map[string][]byte // path -> prior bytes; a nil value with true `ok` (see existed) means "existed but empty" is impossible to distinguish from "didn't exist" using a plain map, so existed is tracked separately
	existed map[string]bool
}

// overlayFiles writes each entry of files (destination path -> new bytes) to
// disk, first recording whatever was there before (bytes, or nothing) so
// revert can restore it exactly. If a write fails partway through, the files
// written so far are still recorded and will be reverted by the caller's
// defer — overlayFiles itself does not roll back on a partial failure; it
// returns the overlay so far plus the error, and the caller's defer must
// still call revert on it.
func overlayFiles(files map[string][]byte) (buildOverlay, error) {
	o := buildOverlay{prior: make(map[string][]byte), existed: make(map[string]bool)}
	for path, data := range files {
		prior, err := os.ReadFile(path)
		switch {
		case err == nil:
			o.prior[path] = prior
			o.existed[path] = true
		case os.IsNotExist(err):
			o.existed[path] = false
		default:
			return o, fmt.Errorf("reading prior %s: %w", path, err)
		}
		if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
			return o, fmt.Errorf("creating %s: %w", filepath.Dir(path), err)
		}
		if err := os.WriteFile(path, data, filePerm); err != nil {
			return o, fmt.Errorf("overlaying %s: %w", path, err)
		}
	}
	return o, nil
}

// revert restores every path the overlay touched to its prior state —
// rewriting its prior bytes, or removing it if it did not exist before —
// regardless of the build outcome that follows overlayFiles. Errors are
// joined so one failed revert does not abandon the rest.
func (o buildOverlay) revert() error {
	var errs []error
	for path, existed := range o.existed {
		if existed {
			if err := os.WriteFile(path, o.prior[path], filePerm); err != nil {
				errs = append(errs, fmt.Errorf("restoring %s: %w", path, err))
			}
			continue
		}
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			errs = append(errs, fmt.Errorf("removing %s: %w", path, err))
		}
	}
	if len(errs) == 0 {
		return nil
	}
	joined := errs[0]
	for _, e := range errs[1:] {
		joined = fmt.Errorf("%w; %w", joined, e)
	}
	return joined
}

// buildErrorRE matches one Go compiler error line: "path/to/file.go:line:col:
// message". Go's error format is stable across versions; this deliberately
// does not try to parse the message, only the leading file reference — the
// `# <package path>` header lines `go build` prints per broken package do not
// match this and are ignored.
var buildErrorRE = regexp.MustCompile(`(?m)^([^:\s][^:\n]*\.go):(\d+):(\d+): (.+)$`)

// buildError is one parsed compiler error line, with its file path resolved
// to absolute (go build reports paths relative to the module root it was run
// from).
type buildError struct {
	file    string // absolute path
	line    int
	col     int
	message string
}

// parseBuildErrors extracts every file:line:col error from go build's
// combined output, resolving each file path against repoRoot (go build
// reports module-relative paths when run from the module root).
func parseBuildErrors(output, repoRoot string) []buildError {
	matches := buildErrorRE.FindAllStringSubmatch(output, -1)
	out := make([]buildError, 0, len(matches))
	for _, m := range matches {
		file := m[1]
		if !filepath.IsAbs(file) {
			file = filepath.Join(repoRoot, file)
		}
		var line, col int
		_, _ = fmt.Sscanf(m[2], "%d", &line)
		_, _ = fmt.Sscanf(m[3], "%d", &col)
		out = append(out, buildError{file: file, line: line, col: col, message: m[4]})
	}
	return out
}

// blamedFiles reduces a set of buildErrors to the distinct absolute file paths
// they blame.
func blamedFiles(errs []buildError) map[string]struct{} {
	out := make(map[string]struct{}, len(errs))
	for _, e := range errs {
		out[e.file] = struct{}{}
	}
	return out
}

// buildOnce overlays files (destination path -> content) onto the real tree,
// runs `go build ./...` from repoRoot once, unconditionally reverts the
// overlay, and returns whether the build was green plus every parsed compiler
// error (empty when green). A non-nil error is a mechanical failure (I/O,
// failing to even invoke go build) distinct from a normal red build, which is
// reported via the returned errors slice, not err.
func buildOnce(repoRoot string, files map[string][]byte) (ok bool, errs []buildError, err error) {
	overlay, oerr := overlayFiles(files)
	defer func() {
		if rerr := overlay.revert(); rerr != nil && err == nil {
			err = fmt.Errorf("reverting build overlay: %w", rerr)
		}
	}()
	if oerr != nil {
		return false, nil, fmt.Errorf("staging build overlay: %w", oerr)
	}

	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = repoRoot
	out, runErr := cmd.CombinedOutput()
	if runErr == nil {
		return true, nil, nil
	}
	parsed := parseBuildErrors(string(out), repoRoot)
	if len(parsed) == 0 {
		// go build failed but produced nothing matching the expected error
		// format — a mechanical problem (bad module state, go itself missing),
		// not an attributable compile error.
		return false, nil, fmt.Errorf("go build ./... failed with unparseable output: %w\n%s", runErr, out)
	}
	return false, parsed, nil
}

// checkListResourceCoupling is the cheap guard for a gap the compile gate's
// fixpoint does not close on its own: reconcileListResource (update.go)
// already reconciles a resource's ListResource registration against its
// plural data source's outcome, but only generation's outcome — it runs
// before compileFixpoint, off generateCorpus's results alone. If the plural
// artifact generates fine (so the resource is staged with
// GenerateListResource: true, baked into its rendered code's
// AddListResourceFactory call) but is later rejected by the compile gate
// itself, nothing re-runs the reconciliation: the resource would promote
// still advertising a list resource with no working plural data source
// behind it — the exact "list resource with no backing data source" the
// design (bigdiffer-design.md §6) rules out, and, since the candidate is
// unchanged next cycle, a persistent one.
//
// In practice this requires the plural data source template itself to
// regress (it renders from the CloudFormation type name alone, with no
// schema dependency, so it essentially always compiles) — rare, but a real
// gap against a stated invariant, not a hypothetical. Rather than the fuller
// fix (re-running reconcileListResource mid-fixpoint, which would mean
// regenerating and re-staging the resource and triggering another build
// round), this is the cheap version: detect the inconsistent combination
// after the fixpoint has settled and refuse to promote it, matching the same
// "promote nothing rather than something known-inconsistent" posture as the
// fixpoint's own unattributable-failure fallback.
func checkListResourceCoupling(candidates []*gateResult) error {
	for _, gr := range candidates {
		pluralOK := false
		pluralAttempted := false
		for _, a := range gr.artifacts {
			if a.kind == artifactPluralDataSource {
				pluralAttempted = true
				pluralOK = a.outcome == gateOK
			}
		}
		if pluralAttempted && !pluralOK {
			return fmt.Errorf("%s: resource was staged advertising a list resource (GenerateListResource) but its plural data source was rejected by the compile gate after generation succeeded — reconcileListResource only reconciles generation's outcome, not the compile gate's; refusing to promote a list resource with no backing data source", gr.cfType)
		}
	}
	return nil
}

// projectRows derives the row set the overlay would contain if promoted right
// now, given the current per-candidate decisions — the same rows
// registrationPackages needs to compute the registration file's import set for
// this fixpoint round, without writing to disk. It is a thin wrapper around the
// same normalizeWithDecisions the real promotion path uses (never a
// hand-rolled, parallel reimplementation of "how a decision changes a row"),
// round-tripped back into []resourceRow via the same HCL decoder the real
// overlay is loaded with. Cheap relative to the fixpoint's dominant cost (one
// go build per round): this is in-memory string handling, not I/O.
func projectRows(overlayContent string, base []resourceRow, checkout map[string]bool, decisions map[string]policyDecision) ([]resourceRow, error) {
	out, _, err := normalizeWithDecisions(overlayContent, base, nil, checkout, decisions)
	if err != nil {
		return nil, fmt.Errorf("projecting rows from decisions: %w", err)
	}
	var f allSchemasFile
	if err := hclsimple.Decode("projected.hcl", []byte(out), nil, &f); err != nil {
		return nil, fmt.Errorf("decoding projected overlay: %w", err)
	}
	return f.Resources, nil
}

// maxFixpointRounds bounds the compile-gate fixpoint (design doc "The
// fixpoint: build until clean"). Each non-green round is already guaranteed to
// remove at least one staged artifact, so the loop cannot exceed the staged
// count regardless — this cap exists purely as a clean, conservative
// terminator for the one situation that should never arise (attribution
// failing to shrink the staged set some other way), not as runaway protection.
const maxFixpointRounds = 25

// stagedArtifact identifies which candidate and artifact kind produced one
// staged .go file, so the fixpoint can turn a blamed real file path back into
// a gateResult downgrade and a suppress_* decision. gr points at the
// candidate's full gateResult (every artifact, from generation) shared across
// all of that candidate's staged files, so downgrading one artifact's outcome
// leaves its siblings' outcomes in the same gateResult intact — decide() must
// see the whole picture (e.g. "plural failed the build gate, resource and
// singular are still fine"), not a fabricated single-artifact result.
type stagedArtifact struct {
	class    changeClass
	gr       *gateResult
	artifact int    // index into gr.artifacts for this specific staged file
	testDest string // real destination path of the paired _test.go file
}

// collectStagedGoFiles walks stagingDir/out and returns every non-test .go
// file found, keyed by its real destination path (stagingDir/out mirrors
// cfg.outputRoot exactly, so the relative path under each is identical). Test
// files are excluded: go build ./... never compiles them (design doc Scope).
func collectStagedGoFiles(stagingOut, outputRoot string) (map[string][]byte, error) {
	files := make(map[string][]byte)
	if _, err := os.Stat(stagingOut); os.IsNotExist(err) {
		return files, nil
	}
	err := filepath.Walk(stagingOut, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		rel, err := filepath.Rel(stagingOut, path)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading staged %s: %w", path, err)
		}
		files[filepath.Join(outputRoot, rel)] = data
		return nil
	})
	return files, err
}

// compileFixpoint is the compile gate (generation-punchlist.md item 1): it
// builds the exact set of code the batch is about to promote — every staged
// artifact plus the registration file computed from the current decisions —
// against the real module, downgrading and dropping whatever the compiler
// rejects, until the build goes green or the situation is unattributable, in
// which case it promotes nothing and returns a hard error (design doc
// "Attribution fallback").
//
// decisions and stagedByDest are mutated in place: a downgraded artifact's
// outcome moves from gateOK to gateFailedBuild within its candidate's full
// gateResult (tagged reasonBuildFailed via decide()'s reasonsForFailures, the
// per-artifact reason-map builder that replaced the old single-string
// gateFailureReason — policy.go), that
// candidate's decide() is recomputed from the updated gateResult, its staged
// .go file is deleted from stagingDir/out so it is never promoted, and it is
// removed from stagedByDest so a later round does not try to attribute the
// same file twice. The staged registration file (regPath, mirroring
// cfg.registrationPath under stagingDir/out) is re-rendered from the updated
// decisions before every rebuild, since dropping an artifact can change its
// import set.
func compileFixpoint(cfg config, stagingDir, overlayContent string, base []resourceRow, checkout map[string]bool, decisions map[string]policyDecision, stagedByDest map[string]stagedArtifact, today string) error {
	stagingOut := filepath.Join(stagingDir, "out")
	regRel, err := filepath.Rel(cfg.outputRoot, cfg.registrationPath)
	if err != nil {
		return fmt.Errorf("compile gate: resolving registration path: %w", err)
	}
	regPath := filepath.Join(stagingOut, regRel)

	for round := 0; ; round++ {
		if round >= maxFixpointRounds {
			return fmt.Errorf("compile gate: exceeded %d rounds without a green build (attribution is not shrinking the staged set); promoting nothing", maxFixpointRounds)
		}

		rows, err := projectRows(overlayContent, base, checkout, decisions)
		if err != nil {
			return fmt.Errorf("compile gate: %w", err)
		}
		reg, err := emitRegistration(cfg, rows)
		if err != nil {
			return fmt.Errorf("compile gate: rendering registration file: %w", err)
		}
		if err := os.MkdirAll(filepath.Dir(regPath), dirPerm); err != nil {
			return fmt.Errorf("compile gate: %w", err)
		}
		if err := os.WriteFile(regPath, reg, filePerm); err != nil {
			return fmt.Errorf("compile gate: staging registration file: %w", err)
		}

		files, err := collectStagedGoFiles(stagingOut, cfg.outputRoot)
		if err != nil {
			return fmt.Errorf("compile gate: collecting staged files: %w", err)
		}
		files[cfg.registrationPath] = reg

		ok, buildErrs, err := buildOnce(cfg.repoRoot, files)
		if err != nil {
			return fmt.Errorf("compile gate: %w", err)
		}
		if ok {
			return nil // green: the staged set (and decisions) is exactly what promotes
		}

		blamed := blamedFiles(buildErrs)
		downgraded := 0
		for path := range blamed {
			art, found := stagedByDest[path]
			if !found {
				continue // not one of this batch's staged files (see fallback below)
			}
			delete(stagedByDest, path)
			downgraded++

			art.gr.artifacts[art.artifact].outcome = gateFailedBuild
			art.gr.artifacts[art.artifact].err = fmt.Errorf("go build: %s", firstBuildMessage(buildErrs, path))
			decisions[art.gr.cfType] = decide(art.class, *art.gr, today)

			// Delete the *staged* file (stagingDir/out/<relative path>), not
			// the real destination path: path is the real tree location
			// buildOnce's overlay used for this round's build and will revert
			// on its own return. The staged copy under stagingOut is what a
			// later round's collectStagedGoFiles would otherwise re-collect
			// and re-overlay, reproducing the identical, already-attributed
			// error and starving the loop's progress.
			rel, relErr := filepath.Rel(cfg.outputRoot, path)
			if relErr != nil {
				return fmt.Errorf("compile gate: resolving staged path for rejected artifact %s: %w", path, relErr)
			}
			stagedPath := filepath.Join(stagingOut, rel)
			if err := os.Remove(stagedPath); err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("compile gate: dropping rejected artifact %s: %w", stagedPath, err)
			}

			// Drop the paired staged _test.go too. go build never compiles
			// or blames it (collectStagedGoFiles excludes _test.go, and the
			// compiler doesn't build tests at all), so it would otherwise
			// silently survive promotion alongside a now-suppressed artifact
			// — a spurious committed test exercising a resource that no
			// longer has a _gen.go, invisible to `make build` (which also
			// excludes tests) and never cleaned up by a later run since the
			// type is unchanged. Caught in review, not by the fixpoint tests
			// (none of which stage a paired test file).
			if art.testDest != "" {
				testRel, testRelErr := filepath.Rel(cfg.outputRoot, art.testDest)
				if testRelErr != nil {
					return fmt.Errorf("compile gate: resolving staged test path for rejected artifact %s: %w", art.testDest, testRelErr)
				}
				stagedTestPath := filepath.Join(stagingOut, testRel)
				if err := os.Remove(stagedTestPath); err != nil && !os.IsNotExist(err) {
					return fmt.Errorf("compile gate: dropping rejected artifact's test file %s: %w", stagedTestPath, err)
				}
			}
		}

		if downgraded == 0 {
			// Nothing blamed maps to a staged file: either the base tree was
			// already broken before this run, or bigdiffer's own
			// registrations_gen.go has a bug. Neither is safe to guess at —
			// promote nothing (design doc "Attribution fallback").
			return fmt.Errorf("compile gate: build failed but no blamed file is attributable to this batch (base tree or registration bug?); promoting nothing:\n%s", formatBuildErrors(buildErrs))
		}
	}
}

// firstBuildMessage returns the first parsed error message blaming path, for
// a compact, single-line reason detail.
func firstBuildMessage(errs []buildError, path string) string {
	for _, e := range errs {
		if e.file == path {
			return e.message
		}
	}
	return "rejected by go build"
}

// formatBuildErrors renders every parsed error as one line each, for a hard
// error's detail text.
func formatBuildErrors(errs []buildError) string {
	var b strings.Builder
	for _, e := range errs {
		fmt.Fprintf(&b, "%s:%d:%d: %s\n", e.file, e.line, e.col, e.message)
	}
	return b.String()
}
