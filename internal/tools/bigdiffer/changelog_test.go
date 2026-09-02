// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// staged is a small helper for building a synthetic promoted map in tests:
// one stagedArtifact per artifact, each wired to its own gateResult (mirroring
// how runUpdate shares one *gateResult per candidate across all of that
// candidate's staged artifacts).
func staged(cfType string, kind artifactKind, tfType string, listRes bool) stagedArtifact {
	gr := &gateResult{cfType: cfType}
	return stagedArtifact{gr: gr, kind: kind, tfType: tfType, listRes: listRes}
}

// TestChangelogEntriesNewType is the classNew case: a brand-new type with all
// three artifacts plus a list resource promotes four bullets, one per kind
// per artifact (the plural DS gets its own New Data Source line distinct from
// the singular's, matching the real CHANGELOG.md's convention verified
// against awscc_appstream_user/_users).
func TestChangelogEntriesNewType(t *testing.T) {
	t.Parallel()

	promoted := map[string]stagedArtifact{
		"resource": staged("AWS::Foo::Bar", artifactResource, "awscc_foo_bar", true),
		"singular": staged("AWS::Foo::Bar", artifactSingularDataSource, "awscc_foo_bar", false),
		"plural":   staged("AWS::Foo::Bar", artifactPluralDataSource, "awscc_foo_bars", false),
	}
	preRunOverlay := map[string]resourceRow{} // absent from the pre-run overlay: classNew

	got := changelogEntries(promoted, preRunOverlay)
	want := []changelogEntry{
		{kind: changelogNewDataSource, tfType: "awscc_foo_bar"},
		{kind: changelogNewDataSource, tfType: "awscc_foo_bars"},
		{kind: changelogNewListResource, tfType: "awscc_foo_bar"},
		{kind: changelogNewResource, tfType: "awscc_foo_bar"},
	}
	assertEntriesEqual(t, got, want)
}

// TestChangelogEntriesNewlyUnsuppressedPlural is the design doc's motivating
// classPresent case (the "AWS added the plural list operation later" goal
// from item 9b): the resource and singular DS were already shipped and are
// simply refreshed (not new), but the plural DS was suppressed on the pre-run
// row and is promoted for the first time this run — it, and the resource's
// now-newly-backed List Resource, are the only two bullets. The resource and
// singular DS must NOT get New Resource / New Data Source bullets just
// because they were regenerated.
func TestChangelogEntriesNewlyUnsuppressedPlural(t *testing.T) {
	t.Parallel()

	promoted := map[string]stagedArtifact{
		"resource": staged("AWS::Foo::Bar", artifactResource, "awscc_foo_bar", true),
		"singular": staged("AWS::Foo::Bar", artifactSingularDataSource, "awscc_foo_bar", false),
		"plural":   staged("AWS::Foo::Bar", artifactPluralDataSource, "awscc_foo_bars", false),
	}
	preRunOverlay := map[string]resourceRow{
		"AWS::Foo::Bar": {
			CloudFormationTypeName:             "AWS::Foo::Bar",
			SuppressPluralDataSourceGeneration: true, // suppressed before this run
		},
	}

	got := changelogEntries(promoted, preRunOverlay)
	want := []changelogEntry{
		{kind: changelogNewDataSource, tfType: "awscc_foo_bars"},
		{kind: changelogNewListResource, tfType: "awscc_foo_bar"},
	}
	assertEntriesEqual(t, got, want)
}

// TestChangelogEntriesOrdinaryRefreshYieldsNothing is the negative case for
// the above: a classPresent type where nothing was previously suppressed,
// refreshed against unchanged or newer schema bytes, must not produce any
// bullet at all — its artifacts (and its list resource) were already
// user-visible.
func TestChangelogEntriesOrdinaryRefreshYieldsNothing(t *testing.T) {
	t.Parallel()

	promoted := map[string]stagedArtifact{
		"resource": staged("AWS::Foo::Bar", artifactResource, "awscc_foo_bar", true),
		"singular": staged("AWS::Foo::Bar", artifactSingularDataSource, "awscc_foo_bar", false),
		"plural":   staged("AWS::Foo::Bar", artifactPluralDataSource, "awscc_foo_bars", false),
	}
	preRunOverlay := map[string]resourceRow{
		"AWS::Foo::Bar": {CloudFormationTypeName: "AWS::Foo::Bar"}, // nothing suppressed
	}

	got := changelogEntries(promoted, preRunOverlay)
	if len(got) != 0 {
		t.Errorf("got %+v, want no entries for an ordinary refresh of an already-shipped type", got)
	}
}

// TestChangelogEntriesGateDroppedOrSuppressedYieldsNothing confirms the
// design's core claim: since promoted is read *after* the compile gate
// settles, an artifact the gate rejected (deleted from stagedByDest by
// compileFixpoint) or one that stayed suppressed simply never appears in the
// map at all — there is nothing for changelogEntries to filter out, because
// it was never staged. This test exercises that by omission: a promoted map
// with only two of three artifacts (as if the plural were dropped/suppressed)
// yields bullets only for the two present.
func TestChangelogEntriesGateDroppedOrSuppressedYieldsNothing(t *testing.T) {
	t.Parallel()

	promoted := map[string]stagedArtifact{
		"resource": staged("AWS::Foo::Bar", artifactResource, "awscc_foo_bar", false),
		"singular": staged("AWS::Foo::Bar", artifactSingularDataSource, "awscc_foo_bar", false),
		// plural intentionally absent: gate-dropped or suppressed
	}
	preRunOverlay := map[string]resourceRow{} // classNew

	got := changelogEntries(promoted, preRunOverlay)
	want := []changelogEntry{
		{kind: changelogNewDataSource, tfType: "awscc_foo_bar"},
		{kind: changelogNewResource, tfType: "awscc_foo_bar"},
	}
	assertEntriesEqual(t, got, want)
	for _, e := range got {
		if e.tfType == "awscc_foo_bars" {
			t.Errorf("the dropped/suppressed plural must not appear, got %+v", got)
		}
	}
}

// TestChangelogEntriesOrderingAndDedup asserts the real ordering
// (alphabetical by kind string: Data Source, then List Resource, then
// Resource — verified directly against CHANGELOG.md's committed history, not
// the design doc's own illustrative example, which showed Resource first) and
// that a duplicate (kind, tfType) pair — e.g. two artifacts of the same type
// mapping to the same bullet, which cannot happen in practice but the
// dedup guard exists for defensively — is never emitted twice.
func TestChangelogEntriesOrderingAndDedup(t *testing.T) {
	t.Parallel()

	promoted := map[string]stagedArtifact{
		"a-resource": staged("AWS::Zzz::Last", artifactResource, "awscc_zzz_last", false),
		"a-singular": staged("AWS::Zzz::Last", artifactSingularDataSource, "awscc_zzz_last", false),
		"a-plural":   staged("AWS::Zzz::Last", artifactPluralDataSource, "awscc_zzz_lasts", false),
		"b-resource": staged("AWS::Aaa::First", artifactResource, "awscc_aaa_first", true),
	}
	preRunOverlay := map[string]resourceRow{}

	got := changelogEntries(promoted, preRunOverlay)
	want := []changelogEntry{
		{kind: changelogNewDataSource, tfType: "awscc_zzz_last"},
		{kind: changelogNewDataSource, tfType: "awscc_zzz_lasts"},
		{kind: changelogNewListResource, tfType: "awscc_aaa_first"},
		{kind: changelogNewResource, tfType: "awscc_aaa_first"},
		{kind: changelogNewResource, tfType: "awscc_zzz_last"},
	}
	assertEntriesEqual(t, got, want)

	// Verify dedup directly: calling with a map that (by construction, since
	// map keys are unique) cannot literally duplicate a stagedArtifact, but
	// confirms the seen-set guard compiles and behaves as a no-op when there
	// is nothing to dedup — the real protection is structural (one map entry
	// per destination path).
	again := changelogEntries(promoted, preRunOverlay)
	assertEntriesEqual(t, again, want)
}

// TestFormatChangelogFragmentGolden is the format golden check: the emitted
// block must match CHANGELOG.md's real bullet syntax verbatim.
func TestFormatChangelogFragmentGolden(t *testing.T) {
	t.Parallel()

	entries := []changelogEntry{
		{kind: changelogNewDataSource, tfType: "awscc_foo_bar"},
		{kind: changelogNewDataSource, tfType: "awscc_foo_bars"},
		{kind: changelogNewListResource, tfType: "awscc_foo_bar"},
		{kind: changelogNewResource, tfType: "awscc_foo_bar"},
	}
	got := formatChangelogFragment(entries)
	want := "FEATURES:\n\n" +
		"* **New Data Source:** `awscc_foo_bar`\n" +
		"* **New Data Source:** `awscc_foo_bars`\n" +
		"* **New List Resource:** `awscc_foo_bar`\n" +
		"* **New Resource:** `awscc_foo_bar`\n"
	if got != want {
		t.Errorf("formatChangelogFragment mismatch:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

// TestFormatChangelogFragmentEmpty confirms an empty entry set produces an
// empty fragment (no bare "FEATURES:" header with nothing under it), so
// runUpdate's caller can use an empty string as "don't write the file."
func TestFormatChangelogFragmentEmpty(t *testing.T) {
	t.Parallel()
	if got := formatChangelogFragment(nil); got != "" {
		t.Errorf("got %q, want empty string for no entries", got)
	}
}

func assertEntriesEqual(t *testing.T, got, want []changelogEntry) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %d entries %+v, want %d %+v", len(got), got, len(want), want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("entry %d: got %+v, want %+v (full got=%+v want=%+v)", i, got[i], want[i], got, want)
		}
	}
}

// TestFormatChangelogFragmentUsesRealBulletSyntax is a narrower, string-level
// sanity check independent of the golden test above, guarding specifically
// against a future refactor accidentally changing the bold-marker or
// backtick placement.
func TestFormatChangelogFragmentUsesRealBulletSyntax(t *testing.T) {
	t.Parallel()
	got := formatChangelogFragment([]changelogEntry{{kind: changelogNewResource, tfType: "awscc_foo_bar"}})
	if !strings.Contains(got, "* **New Resource:** `awscc_foo_bar`") {
		t.Errorf("got %q, missing the exact CHANGELOG.md bullet syntax", got)
	}
}

// TestWriteChangelogFragmentInsertsAfterNotes covers the real origin/main
// shape: a top version block with a NOTES: section and no FEATURES: section
// yet — the common case, since the top block is retitled fresh right after
// each release and NOTES: entries (if any) are usually added by hand before
// -update ever runs. The new FEATURES: section must land after NOTES:'s own
// bullets and before the next "## " heading, leaving everything else (the
// heading text, NOTES: content, the older release below) byte-for-byte
// untouched.
func TestWriteChangelogFragmentInsertsAfterNotes(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "CHANGELOG.md")
	original := `## 1.100.0 (Unreleased)

NOTES:

* provider: some prior note. ([#1111](https://example.com/1111))

## 1.99.0 (August 28, 2026)

NOTES:

* provider: an older release's note.

FEATURES:

* **New Resource:** ` + "`awscc_old_thing`" + `
`
	if err := os.WriteFile(path, []byte(original), filePerm); err != nil {
		t.Fatalf("seeding CHANGELOG.md: %v", err)
	}

	entries := []changelogEntry{
		{kind: changelogNewDataSource, tfType: "awscc_foo_bar"},
		{kind: changelogNewResource, tfType: "awscc_foo_bar"},
	}
	if err := writeChangelogFragment(path, entries); err != nil {
		t.Fatalf("writeChangelogFragment: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	want := `## 1.100.0 (Unreleased)

NOTES:

* provider: some prior note. ([#1111](https://example.com/1111))

FEATURES:

* **New Data Source:** ` + "`awscc_foo_bar`" + `
* **New Resource:** ` + "`awscc_foo_bar`" + `

## 1.99.0 (August 28, 2026)

NOTES:

* provider: an older release's note.

FEATURES:

* **New Resource:** ` + "`awscc_old_thing`" + `
`
	if string(got) != want {
		t.Errorf("mismatch:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

// TestWriteChangelogFragmentInsertsAfterHeadingNoNotes covers a top block
// with no NOTES: section at all — the new FEATURES: section must land
// directly after the heading's blank line.
func TestWriteChangelogFragmentInsertsAfterHeadingNoNotes(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "CHANGELOG.md")
	original := "## 1.100.0 (Unreleased)\n\n## 1.99.0 (August 28, 2026)\n\nFEATURES:\n\n* **New Resource:** `awscc_old_thing`\n"
	if err := os.WriteFile(path, []byte(original), filePerm); err != nil {
		t.Fatalf("seeding CHANGELOG.md: %v", err)
	}

	entries := []changelogEntry{{kind: changelogNewResource, tfType: "awscc_foo_bar"}}
	if err := writeChangelogFragment(path, entries); err != nil {
		t.Fatalf("writeChangelogFragment: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	want := "## 1.100.0 (Unreleased)\n\nFEATURES:\n\n* **New Resource:** `awscc_foo_bar`\n\n" +
		"## 1.99.0 (August 28, 2026)\n\nFEATURES:\n\n* **New Resource:** `awscc_old_thing`\n"
	if string(got) != want {
		t.Errorf("mismatch:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

// TestWriteChangelogFragmentEmptyIsNoop confirms an empty entries list never
// touches the file — including never erroring when the file doesn't even
// have a recognizable heading, since there is nothing to insert.
func TestWriteChangelogFragmentEmptyIsNoop(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "CHANGELOG.md")
	const original = "not even a changelog\n"
	if err := os.WriteFile(path, []byte(original), filePerm); err != nil {
		t.Fatalf("seeding CHANGELOG.md: %v", err)
	}

	if err := writeChangelogFragment(path, nil); err != nil {
		t.Fatalf("writeChangelogFragment: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != original {
		t.Errorf("file was modified on an empty entries list: got %q, want unchanged %q", got, original)
	}
}

// TestWriteChangelogFragmentRefusesNonEmptyExistingSection is the guard
// against silently corrupting a human-owned file: bigdiffer normally runs
// once per release cycle and drafts the whole entry in a single pass, so a
// top block that already has populated FEATURES: bullets means either
// -update ran twice without an intervening release, or the file's shape
// isn't what was expected — either way, this must error rather than guess
// how to merge.
func TestWriteChangelogFragmentRefusesNonEmptyExistingSection(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "CHANGELOG.md")
	original := "## 1.100.0 (Unreleased)\n\nFEATURES:\n\n* **New Resource:** `awscc_already_here`\n"
	if err := os.WriteFile(path, []byte(original), filePerm); err != nil {
		t.Fatalf("seeding CHANGELOG.md: %v", err)
	}

	entries := []changelogEntry{{kind: changelogNewResource, tfType: "awscc_foo_bar"}}
	if err := writeChangelogFragment(path, entries); err == nil {
		t.Fatal("want an error when the top block already has a non-empty FEATURES: section, got nil")
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != original {
		t.Errorf("file was modified despite the error: got %q, want unchanged %q", got, original)
	}
}

// TestWriteChangelogFragmentMissingHeadingErrors confirms a file with no
// "## " heading at all is an error (rather than silently doing nothing or
// panicking) when there is something to insert.
func TestWriteChangelogFragmentMissingHeadingErrors(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "CHANGELOG.md")
	if err := os.WriteFile(path, []byte("not even a changelog\n"), filePerm); err != nil {
		t.Fatalf("seeding CHANGELOG.md: %v", err)
	}

	entries := []changelogEntry{{kind: changelogNewResource, tfType: "awscc_foo_bar"}}
	if err := writeChangelogFragment(path, entries); err == nil {
		t.Fatal("want an error when no version heading is found, got nil")
	}
}
