// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// The shared candidate-settling pipeline (contributing/docs/held-artifacts-design.md
// §2, §6 step 2): candidate-build -> refreshCandidate -> compileFixpoint ->
// decide(), extracted out of runSync's AWS-specific version so an offline row
// source (every overlay row, no discover() call) can drive the identical
// generate-compile-decide sequence for -reconcile and -check (§6 steps 4/5).
// runSync, -reconcile, and -check differ only in how candidates are built
// (AWS-diffed vs. every overlay row) and what happens once the batch has
// settled (promote vs. report) — settleBatch is the one piece both share
// unchanged.
//
// This step is a pure extraction: decide()'s signature and behavior are
// unchanged. The freeze-vs-held routing fix decide() needs (§3) is item 3's
// scope, not this one's — settleBatch calls decide() exactly as runSync always
// has, so -reconcile/-check cannot safely call it yet (they would mis-freeze
// on their first failure, per §3's routing bug) until item 3 lands.

// settledBatch is what settleBatch produces once every candidate has been
// generated, staged, and compile-gated: the final per-type policy decisions,
// the staged artifacts that survived the compile gate (ready to promote), and
// the gateResults of every candidate that staged at least one artifact with
// GenerateListResource baked in (for the caller's checkListResourceCoupling
// check, which must run after the fixpoint has settled).
type settledBatch struct {
	stagingDir             string // caller must os.RemoveAll this when done
	decisions              map[string]policyDecision
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
		stagedByDest:           stagedByDest,
		listResourceCandidates: listResourceCandidates,
		okN:                    okN,
		brokeN:                 brokeN,
	}, nil
}

// todayString returns the current date in dateLayout (YYYY-MM-DD), the form
// frozen_since and every date-stamped decision use. A tiny wrapper so callers
// don't repeat time.Now().Format(dateLayout) at each call site.
func todayString() string {
	return time.Now().Format(dateLayout)
}
