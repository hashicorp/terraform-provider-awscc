// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

// runReconcile regenerates the whole corpus offline, from the committed
// overlay + schema cache, through the same generate -> compile -> decide
// pipeline runSync uses, and promotes the clean result
// (contributing/docs/bigdiffer-design.md §6, "One engine, three callers").
// Unlike a plain offline generator, it compile-gates: any failure aborts the
// whole run rather than holding broken artifacts and shipping the rest, so
// it can back "land a machinery fix, no AWS crawl" (§0 story 2). reconcile
// has no AWS, no discovery diff, and no New/Absent concept at all — every
// committed row is simply "attempt it."
func runReconcile(ctx context.Context, allSchemasPath, checkoutPath string) error {
	cfg, overlayRows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}
	checkout, err := parseCheckout(checkoutPath)
	if err != nil {
		return fmt.Errorf("reading checkout %s: %w", checkoutPath, err)
	}

	overlayByCFN := make(map[string]resourceRow, len(overlayRows))
	for _, r := range overlayRows {
		overlayByCFN[r.CloudFormationTypeName] = r
	}

	cands, frozenSkipped, cacheMissSkipped := reconcileCandidates(overlayRows, cfg.cacheDir)
	stepf("Regenerating %d type(s) offline from the committed overlay + cache (%d frozen, %d cache-miss, skipped)…", len(cands), frozenSkipped, cacheMissSkipped)
	today := todayString()

	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}

	// overlayRows doubles as the offline row source's own "base": settleBatch
	// /compileFixpoint need it to project what the overlay would look like if
	// promoted right now (projectRows, compile.go), and offline there is
	// nothing fresher than what is already committed.
	settled, err := settleBatch(ctx, cfg, cands, nil, string(overlayContent), overlayRows, checkout, today)
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(settled.stagingDir) }()

	// Every candidate here is classPresentUnchanged (reconcileCandidates'
	// own convention, pinned in the design doc and punchlist ahead of this
	// implementation): reconcile has no discovery diff at all, so there is
	// no "the schema moved" case to distinguish, and every failure it
	// produces is, by construction, a machinery failure — never the
	// freeze/suppress branch. See runSync's identical check for why this
	// must be checked explicitly rather than trusting compileFixpoint's own
	// error return (§3, "Why all-or-nothing, precisely"): the fixpoint can
	// reach green while a machineryFailure decision still sits unpromoted,
	// and reconcile aborts on that exactly like sync does — promoting
	// nothing, not even the types that compiled cleanly this run.
	if failures := machineryFailures(settled); len(failures) > 0 {
		var b strings.Builder
		fmt.Fprintf(&b, "%d type(s) failed to regenerate or compile — machinery regression, promoting nothing this run:\n", len(failures))
		for _, f := range failures {
			fmt.Fprintf(&b, "  %s\n", f.cfType)
			for _, a := range f.artifacts {
				fmt.Fprintf(&b, "    %s\n", a)
			}
		}
		return errors.New(b.String())
	}

	decisions := settled.decisions
	stagedByDest := settled.stagedByDest

	// CHANGELOG delta, identical convention to runSync (changelogEntries,
	// changelog.go): every candidate here is classPresentUnchanged, so the
	// classNew "every promoted artifact is new" branch never fires — a
	// promoted artifact is only a New entry if reconcile just lifted a
	// previously-suppressed one back into compiling status (a genuine,
	// changelog-worthy backlog lift), never a routine re-promotion of an
	// already-working type.
	changelog := changelogEntries(stagedByDest, overlayByCFN)

	stepf("Promoting staged output (%d artifact(s) refreshed)…", settled.okN+settled.brokeN)
	if err := promoteStaged(cfg, settled.stagingDir); err != nil {
		return err
	}

	stepf("Reconciling all_schemas.hcl and applying policy…")
	out, report, err := normalizeWithDecisions(string(overlayContent), overlayRows, nil, checkout, decisions)
	if err != nil {
		return err
	}
	// The count-header line ("# N CloudFormation resource types schemas are
	// available...") means "AWS says N are available" — a number reconcile,
	// offline and never calling discover(), has no authoritative value for.
	// normalizeWithDecisions computed it from len(overlayRows) (reconcile's
	// own base, standing in for a live AWS set it does not have), which is a
	// different, larger count than "available from AWS" whenever any row is
	// frozen/retained/non-provisionable — not a no-op rewrite, a wrong one.
	// runLint already treats this exact line as unknowable offline (its own
	// count-header comparison excludes it); reconcile does the same by
	// restoring the committed line verbatim rather than trusting the
	// recomputed one.
	if committedCountLine := countLineRE.FindString(string(overlayContent)); committedCountLine != "" {
		out = countLineRE.ReplaceAllLiteralString(out, committedCountLine)
	}
	if err := os.WriteFile(allSchemasPath, []byte(out), filePerm); err != nil {
		return fmt.Errorf("writing %s: %w", allSchemasPath, err)
	}

	// Emit import_examples_gen.json from the promoted state, same ordering
	// rationale as runSync: it reads each resource's cached schema bytes, so
	// it must run after the cache is promoted.
	stepf("Emitting import examples…")
	_, finalRows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}
	ie, err := emitImportExamples(cfg, finalRows)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfg.importExamplesPath, ie, filePerm); err != nil {
		return fmt.Errorf("writing %s: %w", cfg.importExamplesPath, err)
	}

	if err := writeChangelogFragment(cfg.changelogPath, changelog); err != nil {
		return err
	}
	if len(changelog) > 0 {
		stepf("Added %d CHANGELOG entries to %s.", len(changelog), cfg.changelogPath)
	}

	report.write()
	// Unlike runSync's equivalent tally, this is never split into "OK" vs
	// "frozen/suppressed": reconcile never freezes or suppresses at all (its
	// only failure disposition is machineryFailure, and machineryFailures
	// already aborted above before reaching here if any candidate had one),
	// and frozen rows were excluded from cands entirely
	// (reconcileCandidates). So every decision still standing here is, by
	// construction, a clean generation — reporting a second, always-zero
	// number would be vacuous, not just uninformative.
	stepf("Done: reconciled %d type(s) — all generated OK.", len(cands))
	infof("Review `git status`/`git diff`, then: `make build`, `make smoke`, `go run ./internal/tools/bigdiffer -docs`.")
	return nil
}

// reconcileCandidates builds reconcile's offline candidate set: every
// committed overlay row except frozen ones, each carrying its own cached
// schema bytes (there is no discovery diff to attach fresher ones from —
// generating from the same bytes every row already has is exactly what
// "redo everything the new way, don't ask AWS" means, §0 story 2). It
// returns the candidates plus a count of frozen rows skipped, for the
// caller's progress line.
//
// Two conventions pinned ahead of this implementation
// (contributing/docs/bigdiffer-design.md §6, "Routing a failure to the
// right branch"), both load-bearing:
//   - Every candidate is classed classPresentUnchanged, never classPresent.
//     reconcile has no discovery diff at all (it never calls discover()), so
//     there is no "the schema moved" case classPresent exists to
//     distinguish; every failure a candidate produces must route through
//     machineryFailure. Reusing classPresent here by copy-paste convenience
//     from buildCandidates (sync.go) would silently reintroduce the exact
//     mis-freeze bug item 3 fixed.
//   - Frozen rows are excluded entirely, mirroring the exclusion sync gets
//     for free via classifyChange's statusFrozen short-circuit (change.go),
//     which reconcile has no equivalent call to inherit it from. A frozen
//     type is frozen precisely because it fails to generate/compile at its
//     pinned bytes; feeding it through classPresentUnchanged would hit
//     machineryFailure on that row every single run, aborting reconcile
//     unconditionally with no path to ever completing.
//
// cacheMissSkipped counts rows with no cached schema bytes at all (a row
// that predates the cache, or a genuine cache miss) — rare, but counted and
// surfaced to the caller for the same reason frozenSkipped is: a silent skip
// with no visible count would make "why did reconcile only touch N of the
// M rows I expected" a mystery a maintainer has to go dig for by hand.
func reconcileCandidates(rows []resourceRow, cacheDir string) (cands []candidate, frozenSkipped, cacheMissSkipped int) {
	for _, row := range rows {
		if row.FrozenSince != "" {
			frozenSkipped++
			continue
		}
		schema, err := os.ReadFile(schemaCachePath(cacheDir, row.CloudFormationTypeName))
		if err != nil {
			// No cached bytes to regenerate from at all — nothing offline
			// can do here; skip it exactly as an absent-row probe result
			// would leave an unresolvable row alone for a later run, rather
			// than manufacturing a failure out of missing input.
			cacheMissSkipped++
			continue
		}
		cands = append(cands, candidate{
			cfType: row.CloudFormationTypeName,
			class:  classPresentUnchanged,
			row:    row,
			schema: schema,
		})
	}
	return cands, frozenSkipped, cacheMissSkipped
}
