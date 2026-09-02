// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
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
// CHANGELOG.md bullets. promoted is -update's stagedByDest *after*
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
// held in memory by the time -update reaches this point.
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

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].kind != entries[j].kind {
			return entries[i].kind < entries[j].kind
		}
		return entries[i].tfType < entries[j].tfType
	})
	return entries
}

// writeChangelogFragment inserts entries as a FEATURES: section into path's
// (CHANGELOG.md) top version block — the in-progress release entry (e.g.
// "## 1.100.0 (Unreleased)"; a release-prep commit retitles it and starts a
// fresh one immediately after cutting a release, confirmed against real
// history: "Add changelog entry for v1.100.0" landed right after "Bumped
// product version to 1.99.1"). Assumes that block's FEATURES: section is
// empty or absent when -update runs — bigdiffer normally runs once and
// drafts the whole entry in one pass, so there is nothing to merge with. An
// empty entries list is a no-op.
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

// insertChangelogFeatures inserts entries as a FEATURES: section into
// content's first version block, returning the whole updated content.
// Section order within a block, verified against every block in the real
// file's history, is always NOTES: (optional), then FEATURES: (optional),
// then BUG FIXES: (optional), then the next "## " heading or EOF. The new
// section goes directly after NOTES: (or directly after the "## " heading
// line if there is no NOTES:) and before BUG FIXES:/the next heading.
//
// Errors if the block already has a non-empty FEATURES: section — bigdiffer
// normally runs once per week and drafts the whole entry in a single pass,
// so an existing populated section means either -update ran twice without an
// intervening release (merge by hand) or the block structure isn't what was
// expected; guessing how to merge silently risks corrupting a human-owned
// file.
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
		if !allBlank(block[idx+1 : end]) {
			return "", fmt.Errorf("top version block already has a non-empty FEATURES: section " +
				"(-update normally runs once per release cycle; merge the new entries by hand)")
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
