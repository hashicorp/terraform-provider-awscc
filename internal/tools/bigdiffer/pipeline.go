// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// The shared candidate-settling pipeline (contributing/docs/held-artifacts-design.md
// §1, §5 step 2): candidate-build -> refreshCandidate -> compileFixpoint ->
// decide(), extracted out of runSync's AWS-specific version so an offline row
// source (every overlay row, no discover() call) can drive the identical
// generate-compile-decide sequence for -reconcile and -check (§5 steps 4/5).
// runSync, -reconcile, and -check differ only in how candidates are built
// (AWS-diffed vs. every overlay row) and what happens once the batch has
// settled (promote vs. report) — settleBatch is the one piece both share
// unchanged.
//
// settleBatch is deliberately caller-agnostic about machineryFailure (§3):
// it surfaces the flag on each affected decision (via decide()) but never
// aborts or promotes on its own — that disposition differs per caller
// (sync/reconcile abort; check only needs to report and exit non-zero, with
// no promote step to skip). Callers scan the returned decisions with
// machineryFailures (below) and act on it themselves.

// settledBatch is what settleBatch produces once every candidate has been
// generated, staged, and compile-gated: the final per-type policy decisions,
// the staged artifacts that survived the compile gate (ready to promote), and
// the gateResults of every candidate that staged at least one artifact with
// GenerateListResource baked in (for the caller's checkListResourceCoupling
// check, which must run after the fixpoint has settled).
type settledBatch struct {
	stagingDir             string // caller must os.RemoveAll this when done
	decisions              map[string]policyDecision
	gateResults            map[string]*gateResult // every candidate's final gateResult, keyed by cfType — not just staged artifacts' (stagedByDest only covers what survived the gate); machineryFailures uses this for per-artifact blame
	stagedByDest           map[string]stagedArtifact
	listResourceCandidates []*gateResult
	okN, brokeN            int // generation-time tally (pre-compile-gate); see runSync's comment on why the final count is recomputed from decisions instead
}

// settleBatch runs the shared generate -> stage -> compile-gate sequence for
// every candidate in cands, then checks the list-resource/plural-data-source
// coupling invariant once the fixpoint has settled. decisions starts as
// baseDecisions (already-resolved rows outside cands, e.g. runSync's
// absent-row probe results) and gains one entry per candidate. It never
// promotes or touches the overlay — the caller decides what to do with the
// returned settledBatch (runSync promotes; -reconcile/-check, once they exist,
// will promote or report respectively).
//
// The caller must os.RemoveAll(result.stagingDir) once done with it (runSync
// does this via defer, exactly as before this extraction).
func settleBatch(ctx context.Context, cfg config, cands []candidate, baseDecisions map[string]policyDecision, overlayContent string, base []resourceRow, checkout map[string]bool, today string) (settledBatch, error) {
	stagingDir, err := os.MkdirTemp("", "bigdiffer-staging-")
	if err != nil {
		return settledBatch{}, fmt.Errorf("creating staging dir: %w", err)
	}

	decisions := baseDecisions
	if decisions == nil {
		decisions = make(map[string]policyDecision)
	}
	stagedByDest := make(map[string]stagedArtifact)
	// listResourceCandidates tracks every candidate whose staged resource
	// artifact was rendered with listResource: true (baked into the file's
	// AddListResourceFactory call at generation time — reconcileListResource
	// already reconciled this against generation's own plural DS outcome, but
	// only generation's; the compile gate can still drop the plural artifact
	// afterward, and nothing re-runs the reconciliation for that). Checked
	// once, after the fixpoint settles, so a plural artifact the gate drops
	// can never leave a promoted resource advertising a list capability with
	// no backing data source (design doc invariant; bigdiffer-design.md §6).
	var listResourceCandidates []*gateResult
	var okN, brokeN int
	gateResults := make(map[string]*gateResult, len(cands))
	candBar := newBar(len(cands), "regenerate")
	for _, c := range cands {
		gr, staged, err := refreshCandidate(cfg, stagingDir, c)
		if err != nil {
			_ = os.RemoveAll(stagingDir)
			return settledBatch{}, err
		}
		decisions[c.cfType] = decide(c.class, gr, today)
		if gr.ok() {
			okN++
		} else {
			brokeN++
			infof("%s: %s", c.cfType, decisions[c.cfType].summary)
		}
		// gr is shared by pointer with every staged artifact of this
		// candidate: the compile gate (below) downgrades one artifact's
		// outcome in place and recomputes decide() from the same gateResult,
		// so a build-gate rejection of, say, the plural data source does not
		// lose the fact that the resource and singular data source are still
		// fine.
		grPtr := gr
		gateResults[c.cfType] = &grPtr
		for _, r := range staged {
			dest := filepath.Join(cfg.outputRoot, r.a.pathSuffix, r.a.codeFile)
			testDest := filepath.Join(cfg.outputRoot, r.a.pathSuffix, r.a.testFile)
			// Find this artifact's index within grPtr.artifacts (staged only
			// contains artifacts that generated OK, i.e. gateOK in grPtr).
			for ai := range grPtr.artifacts {
				if grPtr.artifacts[ai].kind == r.a.kind {
					stagedByDest[dest] = stagedArtifact{
						class:    c.class,
						gr:       &grPtr,
						artifact: ai,
						testDest: testDest,
						kind:     r.a.kind,
						tfType:   r.a.tfType,
						listRes:  r.a.listResource,
					}
					break
				}
			}
			if r.a.kind == artifactResource && r.a.listResource {
				listResourceCandidates = append(listResourceCandidates, &grPtr)
			}
		}
		_ = candBar.Add(1)
	}
	_ = candBar.Finish()

	// Compile gate (bigdiffer-design.md §6): build the staged code plus
	// the registration file it implies against the real module until it
	// compiles clean, downgrading whatever the compiler rejects. Runs before
	// promotion so a build failure changes what gets promoted rather than
	// promoting something broken; never blocks the release (design doc
	// bigdiffer-design.md §6, "The compile gate").
	stepf("Compile-gating %d staged artifact(s)…", len(stagedByDest))
	if err := compileFixpoint(ctx, cfg, stagingDir, overlayContent, base, checkout, decisions, stagedByDest, today); err != nil {
		_ = os.RemoveAll(stagingDir)
		return settledBatch{}, fmt.Errorf("compile gate: %w", err)
	}

	// The compile gate can drop a plural data source that generated fine
	// (reconcileListResource only reconciles against generation's outcome,
	// which ran before the gate). Refuse to promote a resource still
	// advertising a list resource with no working plural data source behind
	// it, rather than silently promote the inconsistency.
	if err := checkListResourceCoupling(listResourceCandidates); err != nil {
		_ = os.RemoveAll(stagingDir)
		return settledBatch{}, fmt.Errorf("compile gate: list resource coupling: %w", err)
	}

	return settledBatch{
		stagingDir:             stagingDir,
		decisions:              decisions,
		gateResults:            gateResults,
		stagedByDest:           stagedByDest,
		listResourceCandidates: listResourceCandidates,
		okN:                    okN,
		brokeN:                 brokeN,
	}, nil
}

// machineryFailure is one type's machinery-failure detail: its CloudFormation
// type name and a one-line-per-artifact blame (kind: error), for a caller's
// abort/report message. artifacts is sorted by kind for deterministic output.
type machineryFailure struct {
	cfType    string
	artifacts []string // "<kind>: <first line of error>", sorted
}

// machineryFailures returns every decision with machineryFailure set, sorted
// by CloudFormation type name for deterministic reporting, each with its
// gateResult's per-artifact blame attached (baseDecisions entries — e.g.
// runSync's absent-row probe results — are never machineryFailure and have no
// gateResults entry, so they are skipped, not a lookup bug). It must be
// checked after every settleBatch call, before any promotion step
// (contributing/docs/held-artifacts-design.md §3, "Why all-or-nothing,
// precisely"): compileFixpoint can legitimately reach a green build while
// still leaving one or more classPresentUnchanged decisions flagged — a
// schema-unchanged type whose broken new output got dropped and reverted to
// its still-compiling committed file settles the fixpoint cleanly, but the
// decision decide() returned for it is still machineryFailure, not "refreshed
// OK." A caller that only checks compileFixpoint's error would miss this
// entirely and silently promote everything else while that one type sits at
// last-good — exactly the silent keep-last-good behavior the whole-corpus
// gate exists to end. Every caller (sync, reconcile; check has no promote
// step but still reports this) must call this and abort/report on any
// non-empty result; settleBatch itself deliberately never does (its own doc
// comment).
func machineryFailures(b settledBatch) []machineryFailure {
	var out []machineryFailure
	for cfType, d := range b.decisions {
		if !d.machineryFailure {
			continue
		}
		mf := machineryFailure{cfType: cfType}
		if gr, ok := b.gateResults[cfType]; ok {
			for _, a := range gr.artifacts {
				if a.outcome == gateOK {
					continue
				}
				detail := "rejected"
				if a.err != nil {
					detail = firstLine(a.err.Error())
				}
				mf.artifacts = append(mf.artifacts, string(a.kind)+": "+detail)
			}
			sort.Strings(mf.artifacts)
		}
		out = append(out, mf)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].cfType < out[j].cfType })
	return out
}

// todayString returns the current date in dateLayout (YYYY-MM-DD), the form
// frozen_since and every date-stamped decision use. A tiny wrapper so callers
// don't repeat time.Now().Format(dateLayout) at each call site.
func todayString() string {
	return time.Now().Format(dateLayout)
}
