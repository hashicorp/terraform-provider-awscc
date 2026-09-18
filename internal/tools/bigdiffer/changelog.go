// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// changelogKind is one of the three CHANGELOG.md bullet kinds a promoted
// artifact can produce (contributing/docs/bigdiffer-design.md §8).
type changelogKind string

const (
	changelogNewResource     changelogKind = "New Resource"
	changelogNewDataSource   changelogKind = "New Data Source"
	changelogNewListResource changelogKind = "New List Resource"
)

// changelogEntry is one CHANGELOG.md bullet: "* **<kind>:** `<tfType>`".
type changelogEntry struct {
	kind   changelogKind
	tfType string // awscc_… (already pluralized for the plural data source)
}

// changelogEntries classifies the post-gate promoted artifacts into
// CHANGELOG.md bullets. promoted is -sync's stagedByDest *after*
// compileFixpoint and checkListResourceCoupling have settled — every
// remaining entry is, by construction, an artifact that survived the compile
// gate and will actually be promoted (compileFixpoint deletes a rejected
// artifact's entry outright), so no separate "was this dropped" check is
// needed here. preRunOverlay is the committed overlay's rows, keyed by
// CloudFormation type name, read before this run's regeneration — the "was
// this artifact already user-visible" baseline.
//
// An artifact is a changelog New entry iff it is promoted now and was not
// user-visible in the pre-run overlay:
//   - classNew (the CFN type is absent from preRunOverlay) → every promoted
//     artifact of that type is new;
//   - classPresent (or any other class where the type already has an overlay
//     row) → a promoted artifact is new only if its own Suppress*Generation
//     flag was true in the pre-run row — it was suppressed before and is
//     registering for the first time now (a backlog lift, or the "AWS added
//     the plural list operation later" case that motivated item 9b).
//
// A promoted resource with listRes true additionally yields a New List
// Resource entry, but only when that list capability is itself newly
// user-visible: either the type is new, or the plural data source
// specifically was newly unsuppressed this run (the plural is what backs the
// list resource — reconcileListResource, plan.go). A classPresent resource
// that was already shipped with a working plural DS regenerates with
// listRes true every run; that must not re-announce a list resource that has
// existed all along.
//
// Pure: no I/O, no AWS, no live tree — a direct function of two maps already
// held in memory by the time -sync reaches this point.
func changelogEntries(promoted map[string]stagedArtifact, preRunOverlay map[string]resourceRow) []changelogEntry {
	var entries []changelogEntry
	seen := make(map[changelogEntry]bool)
	add := func(e changelogEntry) {
		if !seen[e] {
			seen[e] = true
			entries = append(entries, e)
		}
	}

	// Precompute, once, which CFN types had their plural data source promoted
	// this run after being suppressed on the pre-run row — the one condition
	// under which an already-shipped resource's List Resource becomes newly
	// user-visible (reconcileListResource, plan.go: ListResource generation
	// tracks the plural DS's current suppression state every run, so an
	// already-working list resource must not be re-announced just because the
	// resource itself was regenerated).
	pluralNewlyUnsuppressed := make(map[string]bool)
	for _, sa := range promoted {
		if sa.kind != artifactPluralDataSource || sa.gr == nil {
			continue
		}
		cfn := sa.gr.cfType
		if preRow, hadRow := preRunOverlay[cfn]; hadRow && preRow.SuppressPluralDataSourceGeneration {
			pluralNewlyUnsuppressed[cfn] = true
		}
	}

	for _, sa := range promoted {
		cfn := ""
		if sa.gr != nil {
			cfn = sa.gr.cfType
		}
		preRow, hadRow := preRunOverlay[cfn]

		if !hadRow || wasSuppressed(preRow, sa.kind) {
			switch sa.kind {
			case artifactResource:
				add(changelogEntry{kind: changelogNewResource, tfType: sa.tfType})
			case artifactSingularDataSource, artifactPluralDataSource:
				add(changelogEntry{kind: changelogNewDataSource, tfType: sa.tfType})
			}
		}

		if sa.kind == artifactResource && sa.listRes && (!hadRow || pluralNewlyUnsuppressed[cfn]) {
			add(changelogEntry{kind: changelogNewListResource, tfType: sa.tfType})
		}
	}

	sortChangelogEntries(entries)
	return entries
}

// sortChangelogEntries sorts entries in place by the real CHANGELOG.md
// ordering convention (verified against committed history, not the design
// doc's own illustrative example): alphabetical by kind string — Data
// Source, then List Resource, then Resource — then alphabetical by tfType
// within each kind. Shared by changelogEntries and mergeChangelogBullets so
// a merged section sorts identically to a single-pass one.
func sortChangelogEntries(entries []changelogEntry) {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].kind != entries[j].kind {
			return entries[i].kind < entries[j].kind
		}
		return entries[i].tfType < entries[j].tfType
	})
}

// writeChangelogFragment inserts entries as a FEATURES: section into path's
// (CHANGELOG.md) top version block — the in-progress release entry (e.g.
// "## 1.100.0 (Unreleased)"; a release-prep commit retitles it and starts a
// fresh one immediately after cutting a release, confirmed against real
// history: "Add changelog entry for v1.100.0" landed right after "Bumped
// product version to 1.99.1"). If that block's FEATURES: section already has
// content from an earlier -sync this same still-open cycle, entries are
// merged with it rather than requiring a release to cut between every run
// (see insertChangelogFeatures/mergeChangelogBullets); an empty entries list
// is a no-op.
func writeChangelogFragment(path string, entries []changelogEntry) error {
	if len(entries) == 0 {
		return nil
	}
	orig, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}
	updated, err := insertChangelogFeatures(string(orig), entries)
	if err != nil {
		return fmt.Errorf("updating %s: %w", path, err)
	}
	if err := os.WriteFile(path, []byte(updated), filePerm); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}
	return nil
}

// changelogBulletRE matches one recognized FEATURES: bullet line exactly —
// the same syntax formatChangelogFragment emits, and the only shape
// insertChangelogFeatures' merge (below) knows how to parse back into a
// changelogEntry safely.
var changelogBulletRE = regexp.MustCompile("^\\* \\*\\*(New Data Source|New List Resource|New Resource):\\*\\* `([^`]+)`$")

// parseChangelogBullet parses one existing FEATURES: line back into a
// changelogEntry, or reports ok=false for anything that is not byte-for-byte
// the exact bullet syntax bigdiffer itself emits (a hand-written note, an
// older release's prose-style entry, a blank separator line). ok=false is
// the caller's signal to fall back to erroring rather than merging — a
// section mixing recognized bullets with anything else is not safe to
// reconstruct verbatim from parsed parts alone.
func parseChangelogBullet(line string) (changelogEntry, bool) {
	m := changelogBulletRE.FindStringSubmatch(line)
	if m == nil {
		return changelogEntry{}, false
	}
	return changelogEntry{kind: changelogKind(m[1]), tfType: m[2]}, true
}

// changelogBulletRun locates the maximal contiguous run of recognized
// bullets (parseChangelogBullet) within existing (the current FEATURES:
// section's lines, sans the "FEATURES:" label itself), tolerating blank
// lines inside the run. ok is false if existing has no recognized bullets
// at all, or has recognized bullets in more than one separate run (i.e.
// non-recognized content — a hand-written note, an older-style entry —
// appears both before and after some recognized bullets, or recognized runs
// are separated by non-recognized content in between): with the bullets
// scattered, it is ambiguous which run represents "the new-artifacts list"
// bigdiffer should merge into, so the caller must refuse rather than guess.
// start and end are indices into existing such that existing[start:end] is
// exactly the run (blank lines at its edges are trimmed out of the range).
func changelogBulletRun(existing []string) (start, end int, ok bool) {
	start, end = -1, -1
	inRun := false
	for i, line := range existing {
		switch line {
		case "":
			continue // blank lines don't end a run; handled by the next non-blank check
		default:
			if _, bulletOK := parseChangelogBullet(line); bulletOK {
				if start == -1 {
					start = i
				} else if !inRun {
					// A recognized bullet after a prior run had already
					// closed (by hitting unrecognized content) means two
					// separate runs — ambiguous.
					return 0, 0, false
				}
				end = i + 1
				inRun = true
			} else if start != -1 {
				// Unrecognized content after a run has started: closes
				// that run. A further recognized bullet after this point
				// triggers the ambiguity check above.
				inRun = false
			}
		}
	}
	if start == -1 {
		return 0, 0, false
	}
	return start, end, true
}

// mergeChangelogBullets locates the single contiguous run of recognized
// bullets (changelogBulletRun) within existing (the current FEATURES:
// section's lines, sans the "FEATURES:" label itself) and merges it with
// newEntries, deduplicating on (kind, tfType) and re-sorting by
// sortChangelogEntries' convention. ok is false, and merged/runStart/runEnd
// are zero, if existing has no recognized bullets or has them scattered
// across more than one run — the caller must fall back to erroring rather
// than guessing at a partial merge. runStart/runEnd (indices into existing)
// tell the caller exactly which lines the run occupied, so only that run —
// not the whole section — gets replaced; any non-bullet content before or
// after the run (e.g. a "provider: ..." note, always observed to precede
// the bullets in real CHANGELOG.md history, never interleaved with them) is
// left in place untouched.
func mergeChangelogBullets(existing []string, newEntries []changelogEntry) (merged []changelogEntry, runStart, runEnd int, ok bool) {
	runStart, runEnd, ok = changelogBulletRun(existing)
	if !ok {
		return nil, 0, 0, false
	}

	seen := make(map[changelogEntry]bool)
	add := func(e changelogEntry) {
		if !seen[e] {
			seen[e] = true
			merged = append(merged, e)
		}
	}
	for _, line := range existing[runStart:runEnd] {
		if line == "" {
			continue
		}
		entry, _ := parseChangelogBullet(line) // guaranteed ok by changelogBulletRun
		add(entry)
	}
	for _, e := range newEntries {
		add(e)
	}
	sortChangelogEntries(merged)
	return merged, runStart, runEnd, true
}

// insertChangelogFeatures inserts entries as a FEATURES: section into
// content's first version block, returning the whole updated content.
// Section order within a block, verified against every block in the real
// file's history, is always NOTES: (optional), then FEATURES: (optional),
// then BUG FIXES: (optional), then the next "## " heading or EOF. The new
// section goes directly after NOTES: (or directly after the "## " heading
// line if there is no NOTES:) and before BUG FIXES:/the next heading.
//
// If the block already has a non-empty FEATURES: section, and that section
// contains a single contiguous run of recognized bullets
// (parseChangelogBullet/changelogBulletRun) — whether that run is the whole
// section or just part of it, e.g. following a "provider: ..." note, the
// real-world shape observed throughout CHANGELOG.md's history — that run's
// entries are merged with the new ones (deduplicated on (kind, tfType) via
// the identical seen-set changelogEntries itself uses, then re-sorted by the
// same convention) and only that run is replaced; anything else in the
// section (the note, blank lines) stays exactly where it was. This lets a
// second -sync within the same still-open release cycle (a genuinely common
// case: nothing requires a release to cut between every sync) accumulate
// correctly instead of blocking on a manual merge every time. This is safe
// specifically because every recognized bullet round-trips losslessly
// through changelogEntry — the replaced run is byte-for-byte what a single
// -sync run producing the union of both entry sets would have written, not
// an approximation.
//
// The merge is refused, and this errors exactly as the pre-merge behavior
// did, if the section's recognized bullets are not all contiguous (e.g. a
// note sits between two separate bullet groups, or one appears both before
// and after some bullets) — with the bullets scattered, it is ambiguous
// which run represents "the new-artifacts list," and guessing risks
// corrupting a human-owned file. A section with no recognized bullets at
// all (pure prose) is refused the same way.
func insertChangelogFeatures(content string, entries []changelogEntry) (string, error) {
	lines := strings.Split(content, "\n")

	headingIdx := slices.IndexFunc(lines, func(l string) bool { return strings.HasPrefix(l, "## ") })
	if headingIdx == -1 {
		return "", fmt.Errorf(`no "## " version heading found`)
	}

	blockEnd := len(lines)
	if next := slices.IndexFunc(lines[headingIdx+1:], func(l string) bool { return strings.HasPrefix(l, "## ") }); next != -1 {
		blockEnd = headingIdx + 1 + next
	}
	block := lines[headingIdx+1 : blockEnd]

	if idx := slices.Index(block, "FEATURES:"); idx != -1 {
		end := len(block)
		if next := slices.IndexFunc(block[idx+1:], func(l string) bool {
			return l == "BUG FIXES:" || strings.HasPrefix(l, "## ")
		}); next != -1 {
			end = idx + 1 + next
		}
		existing := block[idx+1 : end]
		if !allBlank(existing) {
			merged, runStart, runEnd, ok := mergeChangelogBullets(existing, entries)
			if !ok {
				return "", fmt.Errorf("top version block already has a non-empty FEATURES: section " +
					"with unrecognized content (not bigdiffer's own bullet syntax) that isn't a single " +
					"contiguous run of recognized bullets; merge the new entries by hand")
			}
			// Replace only the recognized-bullet run (existing[runStart:runEnd])
			// with the merged, re-sorted, re-rendered bullets — anything before
			// or after the run (e.g. a "provider: ..." note, or blank lines)
			// stays exactly where it was, byte-for-byte. Since this fully
			// satisfies the FEATURES: section already, return directly rather
			// than falling through to the fresh-section insertion path below.
			fragment := strings.Split(strings.TrimSuffix(strings.TrimPrefix(formatChangelogFragment(merged), "FEATURES:\n\n"), "\n"), "\n")
			absoluteStart := headingIdx + 1 + idx + 1 + runStart
			absoluteEnd := headingIdx + 1 + idx + 1 + runEnd
			out := make([]string, 0, len(lines)-(absoluteEnd-absoluteStart)+len(fragment))
			out = append(out, lines[:absoluteStart]...)
			out = append(out, fragment...)
			out = append(out, lines[absoluteEnd:]...)
			return strings.Join(out, "\n"), nil
		}
	}

	// Insert right after NOTES:'s section if present, else right after the
	// heading's own blank line.
	insertAt := headingIdx + 1
	for insertAt < blockEnd && lines[insertAt] == "" {
		insertAt++
	}
	if insertAt < blockEnd && lines[insertAt] == "NOTES:" {
		insertAt++
		for insertAt < blockEnd && lines[insertAt] == "" {
			insertAt++ // past the blank line between the label and its bullets
		}
		for insertAt < blockEnd && lines[insertAt] != "" {
			insertAt++ // past NOTES:'s own bullet/paragraph lines
		}
		if insertAt < blockEnd {
			insertAt++ // past the blank line separating NOTES: from what follows
		}
	}

	fragment := strings.Split(strings.TrimSuffix(formatChangelogFragment(entries), "\n"), "\n")
	out := make([]string, 0, len(lines)+len(fragment)+1)
	out = append(out, lines[:insertAt]...)
	out = append(out, fragment...)
	out = append(out, "")
	out = append(out, lines[insertAt:]...)
	return strings.Join(out, "\n"), nil
}

func allBlank(lines []string) bool {
	for _, l := range lines {
		if strings.TrimSpace(l) != "" {
			return false
		}
	}
	return true
}

// wasSuppressed reports whether kind's artifact was suppressed on the pre-run
// row — the backlog-lift / newly-unsuppressed condition for a type that
// already had an overlay row (classPresent and friends).
func wasSuppressed(row resourceRow, kind artifactKind) bool {
	switch kind {
	case artifactResource:
		return row.SuppressResourceGeneration
	case artifactSingularDataSource:
		return row.SuppressSingularDataSourceGeneration
	case artifactPluralDataSource:
		return row.SuppressPluralDataSourceGeneration
	default:
		return false
	}
}

// formatChangelogFragment renders entries as CHANGELOG.md's exact bullet
// syntax. changelogEntries already returns entries grouped by kind and
// alphabetical by tfType within each group; formatting does not re-sort.
// Grouping order is plain alphabetical-by-kind-string ("New Data Source" <
// "New List Resource" < "New Resource"), which is what CHANGELOG.md's real
// history actually uses (verified directly: every release's FEATURES block
// lists all Data Source bullets, then all List Resource bullets, then all
// Resource bullets) — not the Resource-first ordering the design doc's
// illustrative example showed, which does not match the committed file.
func formatChangelogFragment(entries []changelogEntry) string {
	if len(entries) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("FEATURES:\n\n")
	for _, e := range entries {
		fmt.Fprintf(&b, "* **%s:** `%s`\n", e.kind, e.tfType)
	}
	return b.String()
}
