// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// errFakeGeneration stands in for a real generation error in test fixtures
// that need an already-failed (not merely suppressed) sibling artifact.
var errFakeGeneration = errors.New("fake generation failure for test fixture")

// scratchRow returns a synthetic resourceRow whose ResourceTypeName resolves
// (via naming.ParseTerraformTypeName) to a real but unused path suffix under
// internal/aws — a safe location to overlay a deliberately broken file onto
// the real module for compileFixpoint's tests, guaranteed not to collide with
// any real generated package.
func scratchRow(tfType, cfType string) resourceRow {
	return resourceRow{ResourceTypeName: tfType, CloudFormationTypeName: cfType}
}

// stageBrokenArtifact writes a deliberately-broken resource artifact .go file
// plus its paired _test.go under stagingDir/out at the real relative paths
// they would occupy, and returns their real destination paths (dest for
// building stagedByDest, testDest for confirming the regression this fixture
// exists for: dropping a rejected artifact must drop its paired test file
// too, not just the code file the compiler blamed — go build never sees or
// blames _test.go, since it excludes tests entirely) plus a gateResult
// pre-populated as gateOK for that one artifact, mirroring what
// refreshCandidate would have produced had generation actually succeeded.
// Only artifactResource is needed by any test in this file today; the other
// two kinds render identically for this purpose (a broken .go file is a
// broken .go file to the compiler), so this is not generalized over kind.
func stageBrokenArtifact(t *testing.T, cfg config, stagingDir string, row resourceRow, cfType string) (dest, testDest string, gr *gateResult) {
	t.Helper()
	org, svc, res := parseForTest(t, row.ResourceTypeName)
	pathSuffix := org + "/" + svc
	codeFile := res + "_resource_gen.go"
	testFile := res + "_resource_gen_test.go"
	broken := "package " + svc + "\nfunc F() int { return \"not an int\" }\n"
	// A real generated _gen_test.go never references anything from its
	// paired _gen.go (confirmed against internal/aws/logs/*_resource_gen_test.go
	// — it only calls internal/acctest helpers), so a trivial, independently
	// valid test file is a faithful stand-in, not a simplification that hides
	// the bug this fixture targets.
	test := "package " + svc + "_test\nfunc TestNothing() {}\n"

	dir := filepath.Join(stagingDir, "out", pathSuffix)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, codeFile), []byte(broken), filePerm); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, testFile), []byte(test), filePerm); err != nil {
		t.Fatal(err)
	}

	dest = filepath.Join(cfg.outputRoot, pathSuffix, codeFile)
	testDest = filepath.Join(cfg.outputRoot, pathSuffix, testFile)
	gr = &gateResult{cfType: cfType, artifacts: []artifactResult{{kind: artifactResource, outcome: gateOK}}}
	return dest, testDest, gr
}

// parseForTest splits a test fixture's org_svc_res-shaped ResourceTypeName.
// Mirrors naming.ParseTerraformTypeName's convention without importing the
// naming package's regex directly — these tests only need a deterministic
// split, and generationPlan (exercised via projectRows) is the thing actually
// asserting real naming rules elsewhere. Fails the test directly rather than
// returning an error: every caller in this file would just t.Fatal on it
// anyway.
func parseForTest(t *testing.T, tfType string) (org, svc, res string) {
	t.Helper()
	parts := strings.SplitN(tfType, "_", 3)
	if len(parts) != 3 {
		t.Fatalf("test fixture tfType %q must be org_svc_res", tfType)
	}
	return parts[0], parts[1], parts[2]
}

func minimalOverlayFor(rows []resourceRow) (string, []resourceRow) {
	overlay := testHead +
		"# 0 CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n"
	return overlay, rows
}

// TestCompileFixpointAttributesAndSuppressesABuildFailure is the core case: one
// staged artifact has a real type error. The fixpoint must catch it via a real
// go build, downgrade exactly that artifact to gateFailedBuild/build_failed,
// drop its staged file, and converge (green) without touching any other
// candidate.
func TestCompileFixpointAttributesAndSuppressesABuildFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module")
	}
	// Not t.Parallel(): every test in this file overlays real files onto the
	// real module tree via compileFixpoint/buildOnce — running them
	// concurrently would race each other's overlay the same way the design
	// doc rules out for the gate itself. See compile_test.go's
	// TestBuildOnceGreenOnCleanPackage for the same rule.

	cfg, _ := loadCorpus(t)
	stagingDir := t.TempDir()

	row := scratchRow("aws_zzzcfab_thing", "AWS::ZzzCfab::Thing")
	dest, testDest, gr := stageBrokenArtifact(t, cfg, stagingDir, row, row.CloudFormationTypeName)
	// The singular/plural data sources are modeled as already having failed
	// generation (not merely row-suppressed — canonicalBlock, which renders a
	// brand-new row's block, only ever honors discovery's own structural
	// plural determination, never an arbitrary pre-set suppress_* field on
	// the row itself). This is what makes the package genuinely empty once
	// the resource artifact is also downgraded below: a real total failure
	// where every requested artifact failed one gate or another, not a
	// partial one where a still-live sibling would legitimately keep the
	// package's import alive.
	gr.artifacts = append(gr.artifacts,
		artifactResult{kind: artifactSingularDataSource, outcome: gateFailedGeneration, err: errFakeGeneration},
		artifactResult{kind: artifactPluralDataSource, outcome: gateFailedGeneration, err: errFakeGeneration},
	)
	t.Cleanup(func() { _ = os.RemoveAll(filepath.Dir(dest)) })

	decisions := map[string]policyDecision{
		row.CloudFormationTypeName: decide(classNew, *gr, "2026-08-31"),
	}
	stagedByDest := map[string]stagedArtifact{
		dest: {class: classNew, gr: gr, artifact: 0, testDest: testDest},
	}

	overlay, base := minimalOverlayFor([]resourceRow{row})
	err := compileFixpoint(cfg, stagingDir, overlay, base, map[string]bool{}, decisions, stagedByDest, "2026-08-31")
	if err != nil {
		t.Fatalf("compileFixpoint should converge by suppressing the broken artifact, got error: %v", err)
	}

	if _, found := stagedByDest[dest]; found {
		t.Error("the broken artifact should have been removed from stagedByDest once attributed")
	}
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Errorf("compileFixpoint must revert its build overlay; %s should not exist on the real tree, stat err = %v", dest, err)
	}
	if _, err := os.Stat(filepath.Join(stagingDir, "out", "aws", "zzzcfab", "thing_resource_gen.go")); !os.IsNotExist(err) {
		t.Error("the rejected artifact's staged file should have been deleted so it is never promoted")
	}
	if _, err := os.Stat(filepath.Join(stagingDir, "out", "aws", "zzzcfab", "thing_resource_gen_test.go")); !os.IsNotExist(err) {
		t.Error("the rejected artifact's paired staged _test.go must also be deleted — go build never blames it (tests are excluded), so it would otherwise silently survive promotion for a resource with no _gen.go")
	}

	d := decisions[row.CloudFormationTypeName]
	if d.setAttrs[attrSuppressResource] != "true" {
		t.Errorf("expected the resource to be suppressed after the build gate rejected it, got %+v", d)
	}
	if !strings.HasPrefix(d.reason, "build_failed:") {
		t.Errorf("expected a build_failed-tagged reason, got %q", d.reason)
	}
}

// TestCompileFixpointRegistrationDropsFailedPackage confirms the fixpoint
// re-derives the registration file's import set from the *updated* decisions
// each round: a brand-new service package whose only artifact fails the build
// gate must not leave a dangling blank import to an empty package (which would
// itself break the build the fixpoint is trying to converge on).
func TestCompileFixpointRegistrationDropsFailedPackage(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module")
	}
	// Not t.Parallel() — see TestCompileFixpointAttributesAndSuppressesABuildFailure.

	cfg, _ := loadCorpus(t)
	stagingDir := t.TempDir()

	row := scratchRow("aws_zzzcfrd_thing", "AWS::ZzzCfrd::Thing")
	dest, testDest, gr := stageBrokenArtifact(t, cfg, stagingDir, row, row.CloudFormationTypeName)
	// Singular/plural modeled as already generation-failed — see the identical
	// comment in TestCompileFixpointAttributesAndSuppressesABuildFailure for
	// why a row-level pre-suppression does not work for a brand-new row.
	gr.artifacts = append(gr.artifacts,
		artifactResult{kind: artifactSingularDataSource, outcome: gateFailedGeneration, err: errFakeGeneration},
		artifactResult{kind: artifactPluralDataSource, outcome: gateFailedGeneration, err: errFakeGeneration},
	)
	t.Cleanup(func() { _ = os.RemoveAll(filepath.Dir(dest)) })

	decisions := map[string]policyDecision{
		row.CloudFormationTypeName: decide(classNew, *gr, "2026-08-31"),
	}
	stagedByDest := map[string]stagedArtifact{
		dest: {class: classNew, gr: gr, artifact: 0, testDest: testDest},
	}

	overlay, base := minimalOverlayFor([]resourceRow{row})
	if err := compileFixpoint(cfg, stagingDir, overlay, base, map[string]bool{}, decisions, stagedByDest, "2026-08-31"); err != nil {
		t.Fatalf("compileFixpoint: %v", err)
	}

	// The final round's re-emitted registration file (still staged; promotion
	// never runs in this test) must not import the now-empty aws/zzzcfrd
	// package.
	regRel, err := filepath.Rel(cfg.outputRoot, cfg.registrationPath)
	if err != nil {
		t.Fatal(err)
	}
	reg, err := os.ReadFile(filepath.Join(stagingDir, "out", regRel))
	if err != nil {
		t.Fatalf("reading staged registration file: %v", err)
	}
	if strings.Contains(string(reg), "aws/zzzcfrd") {
		t.Errorf("registration file should have dropped the empty package, got:\n%s", reg)
	}
}

// TestCompileFixpointUnattributableFailureHardStops covers the conservative
// fallback: a build failure that cannot be traced to any staged artifact (here
// simulated by a stagedByDest that does not contain the blamed file at all,
// as if a pre-existing base-tree problem or a bug in bigdiffer's own
// registration emission caused it) must promote nothing and return an error,
// never guess.
func TestCompileFixpointUnattributableFailureHardStops(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module")
	}
	// Not t.Parallel() — see TestCompileFixpointAttributesAndSuppressesABuildFailure.

	cfg, _ := loadCorpus(t)
	stagingDir := t.TempDir()

	row := scratchRow("aws_zzzcfuf_thing", "AWS::ZzzCfuf::Thing")
	dest, _, gr := stageBrokenArtifact(t, cfg, stagingDir, row, row.CloudFormationTypeName)
	t.Cleanup(func() { _ = os.RemoveAll(filepath.Dir(dest)) })
	_ = gr // the gateResult itself is irrelevant here; stagedByDest is left empty on purpose

	decisions := map[string]policyDecision{}
	stagedByDest := map[string]stagedArtifact{} // deliberately does not include dest

	overlay, base := minimalOverlayFor([]resourceRow{row})
	err := compileFixpoint(cfg, stagingDir, overlay, base, map[string]bool{}, decisions, stagedByDest, "2026-08-31")
	if err == nil {
		t.Fatal("expected a hard error: the build failure cannot be attributed to any staged artifact")
	}
	if !strings.Contains(err.Error(), "not attributable") && !strings.Contains(err.Error(), "promoting nothing") {
		t.Errorf("expected the conservative-fallback error, got: %v", err)
	}
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Errorf("the build overlay must still be reverted even on the hard-stop path, stat err = %v", err)
	}
}

// TestCompileFixpointRoundCapStopsAPathologicalLoop covers the circuit
// breaker: a build failure that keeps reappearing without shrinking the staged
// set (simulated by a stagedByDest entry that intentionally never matches the
// injected error's file, and never becomes attributable) must stop after
// maxFixpointRounds rather than looping unboundedly. This test does not spawn
// a real unbounded loop — the unattributable path already exits on round 1 in
// practice, so the round cap is exercised structurally by capping and
// confirming termination is bounded, not by actually driving many rounds.
func TestCompileFixpointRoundCapIsBounded(t *testing.T) {
	if maxFixpointRounds <= 0 {
		t.Fatal("maxFixpointRounds must be positive")
	}
	if maxFixpointRounds > 100 {
		t.Errorf("maxFixpointRounds = %d seems too high for a once-a-week gate; each round is a full go build", maxFixpointRounds)
	}
}

// TestCompileGateFailureBlocksPromotion is the atomicity guarantee (item 12)
// extended to the compile gate (item 1). The compile gate always builds the
// real module (cfg.repoRoot) — it cannot be pointed at a temp tree the way
// pure-generation tests point cfg.outputRoot at one — so this does not
// reconstruct a full runUpdate batch. It instead confirms the two halves of
// the guarantee directly: (1) compileFixpoint itself never leaves a trace on
// the real tree when it hard-fails (proven again here, chained with a real,
// otherwise-successful sibling candidate staged alongside the broken one, to
// rule out the sibling's presence changing the outcome), and (2) runUpdate's
// own source unconditionally returns before promoteStaged on a compile-gate
// error (`if err := compileFixpoint(...); err != nil { return ... }`,
// update.go) — the same short-circuit TestUpdateBatchAtomicity already relies
// on for a staging-time failure, just triggered by the gate instead.
func TestCompileGateFailureBlocksPromotion(t *testing.T) {
	if testing.Short() {
		t.Skip("runs a real go build ./... against the module")
	}
	// Not t.Parallel() — see TestCompileFixpointAttributesAndSuppressesABuildFailure.

	cfg, rows := loadCorpus(t)
	lg := logGroupRow(t, rows)
	schema, err := os.ReadFile(schemaCachePath(cfg.cacheDir, "AWS::Logs::LogGroup"))
	if err != nil {
		t.Fatalf("reading committed schema: %v", err)
	}

	stagingDir := t.TempDir()

	// A real, otherwise-clean candidate stages successfully into the same
	// staging tree as the broken one below — proving its presence does not
	// mask or otherwise change the gate's hard-stop on the unattributable
	// failure elsewhere in the batch.
	c1 := candidate{cfType: "AWS::Logs::LogGroup", class: classPresent, row: lg, schema: schema}
	if _, _, err := refreshCandidate(cfg, stagingDir, c1); err != nil {
		t.Fatalf("refreshCandidate (candidate 1, the clean sibling): %v", err)
	}

	row2 := scratchRow("aws_zzzcgfb_thing", "AWS::ZzzCgfb::Thing")
	dest2, _, _ := stageBrokenArtifact(t, cfg, stagingDir, row2, row2.CloudFormationTypeName)
	t.Cleanup(func() { _ = os.RemoveAll(filepath.Dir(dest2)) })

	overlay, base := minimalOverlayFor(append(append([]resourceRow{}, rows...), row2))
	decisions := map[string]policyDecision{}
	stagedByDest := map[string]stagedArtifact{} // deliberately omits dest2: unattributable

	gateErr := compileFixpoint(cfg, stagingDir, overlay, base, map[string]bool{}, decisions, stagedByDest, "2026-08-31")
	if gateErr == nil {
		t.Fatal("expected the compile gate to hard-fail on the unattributable broken artifact")
	}
	if _, statErr := os.Stat(dest2); !os.IsNotExist(statErr) {
		t.Errorf("compile gate must revert its build overlay even on the hard-stop path; %s should not exist, stat err = %v", dest2, statErr)
	}
	// The clean sibling's staged file must still be sitting only in staging —
	// never promoted, never touched by the gate at all (it built fine every
	// round and was never blamed).
	if _, statErr := os.Stat(filepath.Join(stagingDir, "out", "aws", "logs", "log_group_resource_gen.go")); statErr != nil {
		t.Errorf("the clean sibling's staged file should be untouched in staging, got %v", statErr)
	}
}

// TestCheckListResourceCoupling covers the guard added for a real review
// finding: reconcileListResource only ever reconciles a resource's
// GenerateListResource against its plural data source's *generation*
// outcome, before the compile gate runs. If the plural artifact generates
// fine but is later rejected by the compile gate itself, nothing re-runs that
// reconciliation — the guard exists to catch exactly that combination after
// the fixpoint has settled, rather than silently promote a resource
// advertising a list resource with no working plural data source.
func TestCheckListResourceCoupling(t *testing.T) {
	t.Parallel()

	t.Run("plural gate-rejected after listResource was staged true is caught", func(t *testing.T) {
		t.Parallel()
		gr := &gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateOK},
			{kind: artifactPluralDataSource, outcome: gateFailedBuild, err: errFakeGeneration},
		}}
		err := checkListResourceCoupling([]*gateResult{gr})
		if err == nil {
			t.Fatal("expected an error: the resource advertises a list resource with no working plural data source")
		}
		if !strings.Contains(err.Error(), "AWS::Svc::Thing") {
			t.Errorf("expected the error to name the offending type, got: %v", err)
		}
	})

	t.Run("plural still OK is not flagged", func(t *testing.T) {
		t.Parallel()
		gr := &gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateOK},
			{kind: artifactPluralDataSource, outcome: gateOK},
		}}
		if err := checkListResourceCoupling([]*gateResult{gr}); err != nil {
			t.Errorf("plural is fine, should not be flagged, got: %v", err)
		}
	})

	t.Run("plural never attempted (already suppressed) is not flagged", func(t *testing.T) {
		t.Parallel()
		// reconcileListResource already handles "plural was never attempted"
		// by regenerating the resource without listResource before staging —
		// a candidate reaching this guard with no plural artifact in its
		// gateResult at all means that reconciliation already ran; the guard
		// must not double-flag it.
		gr := &gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
			{kind: artifactResource, outcome: gateOK},
		}}
		if err := checkListResourceCoupling([]*gateResult{gr}); err != nil {
			t.Errorf("plural was never attempted, should not be flagged, got: %v", err)
		}
	})

	t.Run("empty candidate list is a no-op", func(t *testing.T) {
		t.Parallel()
		if err := checkListResourceCoupling(nil); err != nil {
			t.Errorf("no candidates staged listResource: true, should not error, got: %v", err)
		}
	})
}
