// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

// Command bigdiffer keeps internal/provider/all_schemas.hcl in sync with what
// AWS offers, and owns generating the provider from it. Five modes:
//
//   - -check: parse the overlay, verify it is normalized and anomaly-free.
//     Offline, writes nothing; safe on every PR. Flags any suppressed/frozen
//     row with no suppression_reason as an advisory anomaly (never a failure).
//   - -update: the live weekly incremental. One AWS crawl (ListTypes +
//     DescribeType, us-east-1) feeds both overlay reconciliation (add new rows,
//     report retained/anomalous ones) and change detection (byte-compare against
//     the schema cache). Only New/Changed types are regenerated from their fresh
//     bytes; generation success/failure drives policy (frozen_since / suppress_*)
//     applied to the overlay in one pass. Never regresses: a type's files + cache
//     are promoted only on clean generation.
//   - -generate: a full, offline, parallel regeneration of the whole provider
//     from the committed overlay + schema cache. Does not touch AWS.
//   - -docs: owns import-example docs generation from import_examples_gen.json,
//     then orchestrates terraform fmt + tfplugindocs generate.
//   - -heal: re-probes every suppressed/frozen row with no recorded reason (or
//     one tagged unknown) and proposes a category + detail for it — structural
//     check, then a real regeneration attempt, then a free-form comment
//     migration — printed as a report, never written. Offline except reading
//     the committed schema cache.
//
// There is no separate discovery or snapshot-diff mode: bigdiffer owns discovery
// itself, so there is no need to generate and diff checked-in
// available_schemas.<date>.hcl snapshots, and no need for git or dated files.
// Reconciliation, on the CloudFormation type name:
//   - Add rows present live but missing from the overlay (recovers backlog and
//     genuinely new resources), copying the resource type name and suppress_plural
//     flag from the discovered row.
//   - Preserve every existing overlay block byte-for-byte, including
//     free-form "# Suppression Reason" comments and suppression attributes.
//   - Re-sort all blocks by CloudFormation type name.
//   - Update the header count to the number of available schemas.
//   - Report anomalies. A retained row (in the overlay, gone from live AWS) is
//     benign when explained by frozen_since or non_provisionable.
//
// bigdiffer is self-contained: it defines all_schemas.hcl's shape (model.go),
// its own discovery (discover.go), its own naming (naming/naming.go), and owns a
// copy of the generation engine (codegen/) rather than importing the legacy
// generator or update packages, which are slated for removal.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
)

const (
	defaultAllSchemas = "internal/provider/all_schemas.hcl"
	defaultCheckout   = "internal/update/suppressions_checkout.txt"

	dateLayout = "2006-01-02"

	pluralAttr = "suppress_plural_data_source_generation"
	cfnAttr    = "cloudformation_type_name"
)

var countLineRE = regexp.MustCompile(`(?m)^# \d+ CloudFormation resource types schemas are available[^\n]*$`)

func main() {
	var (
		allSchemasPath = flag.String("all-schemas", defaultAllSchemas, "path to the curated all_schemas.hcl overlay")
		checkoutPath   = flag.String("checkout", defaultCheckout, "path to suppressions_checkout.txt (cross-referenced, never modified)")
		check          = flag.Bool("check", false, "verify all_schemas.hcl is normalized and anomaly-free (offline; writes nothing)")
		generate       = flag.Bool("generate", false, "regenerate the whole provider offline from the committed overlay + schema cache (writes *_gen.go, registrations_gen.go, import_examples_gen.json)")
		update         = flag.Bool("update", false, "live weekly incremental: discover, refresh only changed types from fresh bytes, apply policy to the overlay, write cache + aggregates (needs AWS, us-east-1)")
		docs           = flag.Bool("docs", false, "regenerate documentation: own import-example docs from import_examples_gen.json, then orchestrate terraform fmt + tfplugindocs")
		heal           = flag.Bool("heal", false, "re-probe suppressed/frozen rows with no recorded reason and propose a reclassification (offline except reading the schema cache; never writes all_schemas.hcl)")

		// Hidden: the hidden -heal-probe-artifact mode heal.go re-execs into,
		// so a crashy regeneration (e.g. a recursive schema) kills only this
		// subprocess. Not part of the documented CLI surface.
		healProbeArtifact = flag.Bool("heal-probe-artifact", false, "internal: probe one artifact's regeneration in isolation")
		probeTFType       = flag.String("probe-tf-type", "", "internal")
		probeCFNType      = flag.String("probe-cfn-type", "", "internal")
		probeKind         = flag.String("probe-kind", "", "internal")
		probeSchema       = flag.String("probe-schema", "", "internal")
		probePrefix       = flag.String("probe-prefix", "", "internal")
		probeCacheDir     = flag.String("probe-cache-dir", "", "internal")
		probeServicesPath = flag.String("probe-services-path", "", "internal")
		probeRepoRoot     = flag.String("probe-repo-root", "", "internal")
		probeOutputRoot   = flag.String("probe-output-root", "", "internal")
	)
	flag.Parse()

	if *healProbeArtifact {
		err := runHealProbeArtifact(*probeTFType, *probeCFNType, *probeKind, *probeSchema, *probePrefix, *probeCacheDir, *probeServicesPath, *probeRepoRoot, *probeOutputRoot)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			var gateFailure *buildGateFailure
			if errors.As(err, &gateFailure) {
				os.Exit(exitCodeBuildGateFailed)
			}
			os.Exit(1)
		}
		return
	}

	if err := run(*allSchemasPath, *checkoutPath, *check, *generate, *update, *docs, *heal); err != nil {
		fmt.Fprintf(os.Stderr, "bigdiffer: %v\n", err)
		os.Exit(1)
	}
}

// run dispatches to exactly one command. bigdiffer is the generator, not a manual
// snapshot-diff helper: the weekly workflow is -update, offline rebuilds are
// -generate, docs are -docs, and -check is an offline hygiene guard for CI.
func run(allSchemasPath, checkoutPath string, check, generate, update, docs, heal bool) error {
	switch {
	case update:
		return runUpdate(context.Background(), allSchemasPath, checkoutPath)
	case generate:
		cfg, rows, err := loadOverlay(allSchemasPath)
		if err != nil {
			return err
		}
		return runGenerate(cfg, rows)
	case docs:
		cfg, _, err := loadOverlay(allSchemasPath)
		if err != nil {
			return err
		}
		return runDocs(cfg)
	case check:
		return runCheck(allSchemasPath, checkoutPath)
	case heal:
		return runHeal(allSchemasPath)
	default:
		flag.Usage()
		return fmt.Errorf("no command given; use one of -update, -generate, -docs, -check, -heal")
	}
}

// runCheck verifies all_schemas.hcl is normalized (sorted, canonical formatting,
// correct count header) and anomaly-free, using the overlay as its own base so no
// AWS query or snapshot file is needed. It writes nothing and is suitable for CI.
func runCheck(allSchemasPath, checkoutPath string) error {
	_, rows, err := loadOverlay(allSchemasPath)
	if err != nil {
		return err
	}
	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}
	checkout, err := parseCheckout(checkoutPath)
	if err != nil {
		return fmt.Errorf("reading checkout %s: %w", checkoutPath, err)
	}

	// Reconcile the overlay against itself: no rows are added, so this only
	// re-sorts and re-formats, surfacing any hand-edit that left it un-normalized.
	out, report, err := normalize(string(overlayContent), rows, nil, checkout)
	if err != nil {
		return err
	}
	report.write()

	problems := report.anomalyProblems()
	// The count-header value counts schemas *available from AWS*, which an offline
	// run cannot know, so compare everything except that line: sorting, canonical
	// formatting, and byte-preservation of blocks.
	if countLineRE.ReplaceAllString(out, "#") != countLineRE.ReplaceAllString(string(overlayContent), "#") {
		problems = append([]string{"not normalized (sorting/formatting; re-run `-update`, or fix by hand)"}, problems...)
	}
	if len(problems) > 0 {
		return fmt.Errorf("all_schemas.hcl check failed: %s", strings.Join(problems, "; "))
	}
	if n := len(report.UnexplainedRetained) + len(report.ReasonlessSuppressed); n > 0 {
		fmt.Fprintf(os.Stderr, "bigdiffer: all_schemas.hcl is normalized; %d advisory anomaly line(s) reported above (not a check failure).\n", n)
		return nil
	}
	fmt.Fprintln(os.Stderr, "bigdiffer: all_schemas.hcl is normalized and anomaly-free.")
	return nil
}

// item is a blank-line-delimited chunk of the resource region, retained as
// verbatim text so that comments, manual attributes, and even fully
// commented-out blocks survive untouched. An item is keyed by the
// CloudFormation type name it carries (from a live attribute if present, else
// from a commented-out one) so it can be joined against the base and sorted.
type item struct {
	key   string // CloudFormation type name used for joining and sorting
	label string // Terraform resource type name, for live blocks only
	live  bool   // true if the item contains an uncommented resource_schema block
	text  string // exact bytes of the chunk
}

// rowRef is a lightweight reference used in the report.
type rowRef struct {
	cfn   string
	label string
}

// Report summarizes what the normalization did and what a human should review.
type Report struct {
	AddedNew             []rowRef // in base, absent from overlay, absent from previous base
	AddedBacklog         []rowRef // in base, absent from overlay, present in previous base
	Retained             []rowRef // live in overlay, absent from base
	UnexplainedRetained  []rowRef // Retained rows not explained by frozen_since, non_provisionable, or a checkout pin
	Duplicates           []string // CloudFormation type names with more than one live block
	NamingViolate        []string
	ReasonlessSuppressed []reasonlessFact // a suppressed artifact or a freeze with its own reason field empty (suppressed-and-frozen.md); one entry per missing fact, not per row
	Total                int
}

// reasonlessFact is one missing reason: a specific suppressed artifact or the
// freeze itself, named by which attribute should carry the explanation. A row
// with, say, its resource suppressed and no frozen reason yields two entries,
// not one — the four reason facts (suppression_reason_resource, _singular_data_source,
// _plural_data_source, frozen_reason) are independent
// (contributing/docs/suppressed-and-frozen.md).
type reasonlessFact struct {
	cfn   string
	label string
	field string // the empty attribute, e.g. "suppression_reason_plural_data_source"
}

func normalize(overlay string, base, previous []resourceRow, checkout map[string]bool) (string, Report, error) {
	return normalizeWithDecisions(overlay, base, previous, checkout, nil)
}

// normalizeWithDecisions is normalize plus a policy pass: as each block is
// written, any policy decision keyed by its CloudFormation type name is applied
// with setBlockAttributes (comment-preserving), so suppression flags,
// frozen_since, and suppression_reason land on exactly the right blocks. A nil
// decisions map reduces to a plain reconcile.
func normalizeWithDecisions(overlay string, base, previous []resourceRow, checkout map[string]bool, decisions map[string]policyDecision) (string, Report, error) {
	loc := countLineRE.FindStringIndex(overlay)
	if loc == nil {
		return "", Report{}, fmt.Errorf("could not find the count header comment in all_schemas.hcl")
	}
	head := overlay[:loc[0]] // copyright, defaults, meta_schema, trailing blank line(s)

	items := extractItems(overlay[loc[1]:])

	// Cross-validate the scanner against a real HCL decode: the set of live
	// resource_schema blocks must match exactly. This catches any scanner drift
	// or unusual formatting before we rewrite the file. The decoded resources
	// also give us per-block attributes (frozen_since, non_provisionable) that
	// explain why a row may legitimately be absent from the base.
	overlayResources, err := crossValidate(overlay, items)
	if err != nil {
		return "", Report{}, err
	}
	overlayByCFN := make(map[string]resourceRow, len(overlayResources))
	for _, r := range overlayResources {
		overlayByCFN[r.CloudFormationTypeName] = r
	}

	// Live blocks indexed by CloudFormation type name.
	liveByCFN := make(map[string]struct{})
	liveCount := make(map[string]int)
	for _, it := range items {
		if it.live {
			liveByCFN[it.key] = struct{}{}
			liveCount[it.key]++
		}
	}

	baseByCFN := make(map[string]resourceRow, len(base))
	for _, r := range base {
		baseByCFN[r.CloudFormationTypeName] = r
	}
	prevByCFN := make(map[string]struct{}, len(previous))
	for _, r := range previous {
		prevByCFN[r.CloudFormationTypeName] = struct{}{}
	}

	var report Report

	// Add base rows missing from the overlay's live blocks.
	for _, r := range base {
		if _, ok := liveByCFN[r.CloudFormationTypeName]; ok {
			continue
		}
		items = append(items, item{
			key:   r.CloudFormationTypeName,
			label: r.ResourceTypeName,
			live:  true,
			text:  canonicalBlock(r.ResourceTypeName, r.CloudFormationTypeName, r.SuppressPluralDataSourceGeneration),
		})
		ref := rowRef{cfn: r.CloudFormationTypeName, label: r.ResourceTypeName}
		if _, seen := prevByCFN[r.CloudFormationTypeName]; seen {
			report.AddedBacklog = append(report.AddedBacklog, ref)
		} else {
			report.AddedNew = append(report.AddedNew, ref)
		}
	}

	// Retained rows: live in overlay, gone from the current available base.
	for _, it := range items {
		if !it.live || it.key == "" {
			continue
		}
		if _, ok := baseByCFN[it.key]; ok {
			continue
		}
		ref := rowRef{cfn: it.key, label: it.label}
		report.Retained = append(report.Retained, ref)
		r := overlayByCFN[it.key]
		explained := checkout[it.key] || r.FrozenSince != "" || r.NonProvisionable
		if !explained {
			report.UnexplainedRetained = append(report.UnexplainedRetained, ref)
		}
	}

	// Naming-invariant check across every live row.
	for _, it := range items {
		if !it.live {
			continue
		}
		if want, err := expectedLabel(it.key); err == nil && want != it.label {
			report.NamingViolate = append(report.NamingViolate,
				fmt.Sprintf("%s: label %q, expected %q", it.key, it.label, want))
		}
	}
	sort.Strings(report.NamingViolate)

	// Reason-less suppression/freeze check (generation-punchlist.md item 8;
	// contributing/docs/suppressed-and-frozen.md): every artifact bigdiffer
	// itself suppresses, and every freeze it sets, carries its own reason
	// (policy.go), but a row can also be suppressed or frozen by direct
	// hand-edit, or predate the taxonomy. Checked per-fact, not per-row (item
	// 9b): a row's resource can have a real reason while its plural DS is
	// still reason-less, or vice versa. Advisory only, like every other
	// anomaly here — never a hard failure, and -heal (not this check) is what
	// fills the gap.
	for _, it := range items {
		if !it.live || it.key == "" {
			continue
		}
		r := overlayByCFN[it.key]
		for _, f := range []struct {
			suppressed bool
			reason     string
			field      string
		}{
			{r.SuppressResourceGeneration, r.SuppressionReasonResource, attrSuppressionReasonResource},
			{r.SuppressSingularDataSourceGeneration, r.SuppressionReasonSingularDataSource, attrSuppressionReasonSingular},
			{r.SuppressPluralDataSourceGeneration, r.SuppressionReasonPluralDataSource, attrSuppressionReasonPlural},
			{r.FrozenSince != "", r.FrozenReason, attrFrozenReason},
		} {
			if f.suppressed && f.reason == "" {
				report.ReasonlessSuppressed = append(report.ReasonlessSuppressed, reasonlessFact{cfn: it.key, label: it.label, field: f.field})
			}
		}
	}
	sort.Slice(report.ReasonlessSuppressed, func(i, j int) bool {
		if report.ReasonlessSuppressed[i].cfn != report.ReasonlessSuppressed[j].cfn {
			return report.ReasonlessSuppressed[i].cfn < report.ReasonlessSuppressed[j].cfn
		}
		return report.ReasonlessSuppressed[i].field < report.ReasonlessSuppressed[j].field
	})

	for cfn, n := range liveCount {
		if n > 1 {
			report.Duplicates = append(report.Duplicates, fmt.Sprintf("%s (%d live blocks)", cfn, n))
		}
	}
	sort.Strings(report.Duplicates)

	// Sort by CloudFormation type name (matches the generator's sort.Strings).
	// Stable so that any (unexpected) key collisions keep their relative order.
	sort.SliceStable(items, func(i, j int) bool { return items[i].key < items[j].key })

	report.Total = len(base)

	var sb strings.Builder
	sb.WriteString(head)
	fmt.Fprintf(&sb, "# %d CloudFormation resource types schemas are available for use with the Cloud Control API.\n\n", len(base))
	for i, it := range items {
		text := it.text
		// Only mutate live blocks. A commented-out block can still carry a
		// matching it.key (classifyItem's comment fallback), and setBlockAttributes
		// requires exactly one resource_schema block, which a commented-out block
		// has none of — mutating it would abort the whole run.
		if it.live {
			if d, ok := decisions[it.key]; ok {
				if attrs := decisionAttrs(d); len(attrs) > 0 {
					mutated, err := setBlockAttributes(text, attrs)
					if err != nil {
						return "", Report{}, fmt.Errorf("applying policy to %s: %w", it.key, err)
					}
					text = mutated
				}
			}
		}
		sb.WriteString(text)
		if i < len(items)-1 {
			sb.WriteString("\n\n")
		} else {
			sb.WriteString("\n")
		}
	}
	return sb.String(), report, nil
}

// decisionAttrs flattens a policy decision into the attribute set to write on
// its block: the suppression/freeze flags plus any reason attributes
// (suppression_reason_resource, _singular_data_source, _plural_data_source,
// frozen_reason), already keyed by attribute name in d.reasons.
func decisionAttrs(d policyDecision) map[string]string {
	if len(d.setAttrs) == 0 && len(d.reasons) == 0 {
		return nil
	}
	attrs := make(map[string]string, len(d.setAttrs)+len(d.reasons))
	maps.Copy(attrs, d.setAttrs)
	maps.Copy(attrs, d.reasons)
	return attrs
}

var (
	labelRE = regexp.MustCompile(`resource_schema\s+"([^"]+)"`)
	cfnRE   = regexp.MustCompile(cfnAttr + `\s*=\s*"([^"]+)"`)
)

// extractItems walks the resource region, delimiting live blocks by their
// braces (so internal blank lines and comments inside a block are kept intact),
// attaching any contiguous leading comments to the following block, and
// treating a blank-terminated comment run (e.g. a fully commented-out block) as
// its own preserved item keyed by any commented-out CloudFormation type name.
func extractItems(region string) []item {
	region = strings.Trim(region, "\n")
	if region == "" {
		return nil
	}
	lines := strings.Split(region, "\n")

	var items []item
	var lead []string // contiguous comment lines with no intervening blank

	flushLead := func() {
		if len(lead) > 0 {
			items = append(items, classifyItem(strings.Join(lead, "\n")))
			lead = nil
		}
	}

	for i := 0; i < len(lines); {
		line := lines[i]
		switch {
		case strings.TrimSpace(line) == "":
			// A blank line ends a standalone comment run.
			flushLead()
			i++
		case strings.HasPrefix(line, `resource_schema "`):
			// Consume a live block through its closing brace at column 0.
			start := i
			i++
			for i < len(lines) && lines[i] != "}" {
				i++
			}
			if i < len(lines) {
				i++ // include the closing "}"
			}
			text := strings.Join(lines[start:i], "\n")
			if len(lead) > 0 {
				text = strings.Join(lead, "\n") + "\n" + text
				lead = nil
			}
			items = append(items, classifyItem(text))
		default:
			// Comment lines (and, defensively, anything else) buffer as lead.
			lead = append(lead, line)
			i++
		}
	}
	flushLead()

	return attachKeyless(items)
}

func classifyItem(text string) item {
	it := item{text: text}
	for ln := range strings.SplitSeq(text, "\n") {
		trimmed := strings.TrimSpace(ln)
		if strings.HasPrefix(trimmed, `resource_schema "`) {
			it.live = true
			if m := labelRE.FindStringSubmatch(ln); m != nil {
				it.label = m[1]
			}
		}
		// A live (uncommented) CloudFormation type name is authoritative.
		if !strings.HasPrefix(trimmed, "#") && strings.HasPrefix(trimmed, cfnAttr) {
			if m := cfnRE.FindStringSubmatch(ln); m != nil {
				it.key = m[1]
			}
		}
	}
	// Fall back to a commented-out CloudFormation type name (e.g. a fully
	// commented-out block retained for context).
	if it.key == "" {
		if m := cfnRE.FindStringSubmatch(text); m != nil {
			it.key = m[1]
		}
	}
	return it
}

// attachKeyless folds any item with no discoverable CloudFormation type name
// into the following item (or the previous one if it is last), so free-standing
// comments never sort away from the block they describe or get dropped.
func attachKeyless(items []item) []item {
	out := make([]item, 0, len(items))
	var pending []string
	for _, it := range items {
		if it.key == "" {
			pending = append(pending, it.text)
			continue
		}
		if len(pending) > 0 {
			it.text = strings.Join(append(pending, it.text), "\n\n")
			pending = nil
		}
		out = append(out, it)
	}
	if len(pending) > 0 {
		if len(out) > 0 {
			out[len(out)-1].text = strings.Join(append([]string{out[len(out)-1].text}, pending...), "\n\n")
		} else {
			out = append(out, item{text: strings.Join(pending, "\n\n")})
		}
	}
	return out
}

// canonicalBlock renders a brand-new resource_schema block for a row present in
// the discovered base but absent from the overlay. The only suppression it ever
// carries is the plural data source's structural determination (discover.go's
// pluralSupported check, run before any generation is attempted), so it always
// stamps that with the structural reason category — never a bare, unexplained
// suppress_plural_data_source_generation (contributing/docs/suppressed-and-frozen.md).
func canonicalBlock(label, cfn string, suppressPlural bool) string {
	if suppressPlural {
		reason := formatReason(reasonStructural, "no list handler with zero required arguments")
		width := max(len(pluralAttr), len(attrSuppressionReasonPlural))
		return fmt.Sprintf("resource_schema %q {\n  %-*s = %q\n  %-*s = true\n  %-*s = %q\n}",
			label, width, cfnAttr, cfn,
			width, pluralAttr,
			width, attrSuppressionReasonPlural, reason)
	}
	return fmt.Sprintf("resource_schema %q {\n  %s = %q\n}", label, cfnAttr, cfn)
}

func expectedLabel(cfn string) (string, error) {
	org, svc, res, err := naming.ParseCloudFormationTypeName(cfn)
	if err != nil {
		return "", err
	}
	return strings.ToLower(org) + "_" + strings.ToLower(svc) + "_" + naming.CloudFormationPropertyToTerraformAttribute(res), nil
}

// crossValidate decodes the overlay with a real HCL parser and asserts that the
// set of live resource_schema blocks equals what the text scanner found. This
// guards the verbatim, comment-preserving text path with a structural check.
func crossValidate(overlay string, items []item) ([]resourceRow, error) {
	parser := hclparse.NewParser()
	f, diag := parser.ParseHCL([]byte(overlay), "all_schemas.hcl")
	if diag.HasErrors() {
		return nil, fmt.Errorf("HCL parse: %s", diag.Error())
	}
	var decoded allSchemasFile
	if diag := gohcl.DecodeBody(f.Body, nil, &decoded); diag.HasErrors() {
		return nil, fmt.Errorf("HCL decode: %s", diag.Error())
	}

	decodedCFN := make(map[string]string, len(decoded.Resources)) // cfn -> label
	for _, r := range decoded.Resources {
		decodedCFN[r.CloudFormationTypeName] = r.ResourceTypeName
	}

	scannedCFN := make(map[string]string)
	for _, it := range items {
		if it.live {
			scannedCFN[it.key] = it.label
		}
	}

	for cfn, label := range decodedCFN {
		if got, ok := scannedCFN[cfn]; !ok {
			return nil, fmt.Errorf("HCL decode found live block %q that the scanner missed", cfn)
		} else if got != label {
			return nil, fmt.Errorf("live block %q: scanner label %q != HCL label %q", cfn, got, label)
		}
	}
	for cfn := range scannedCFN {
		if _, ok := decodedCFN[cfn]; !ok {
			return nil, fmt.Errorf("scanner found live block %q that the HCL decode did not", cfn)
		}
	}
	return decoded.Resources, nil
}

// parseCheckout reads suppressions_checkout.txt and returns the set of pinned
// CloudFormation type names (e.g. AWS_IoTFleetWise_DecoderManifest.json ->
// AWS::IoTFleetWise::DecoderManifest).
func parseCheckout(path string) (map[string]bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]bool{}, nil
		}
		return nil, err
	}
	out := make(map[string]bool)
	for line := range strings.SplitSeq(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		base := filepath.Base(line)
		base = strings.TrimSuffix(base, ".json")
		out[strings.ReplaceAll(base, "_", "::")] = true
	}
	return out, nil
}

// anomalyProblems returns human-readable descriptions of anomalies that should
// fail a -check run. Retained-but-unpinned rows are intentionally excluded: they
// are advisory and tracked for later.
func (r Report) anomalyProblems() []string {
	var problems []string
	if len(r.Duplicates) > 0 {
		problems = append(problems, fmt.Sprintf("%d duplicate live block(s)", len(r.Duplicates)))
	}
	if len(r.NamingViolate) > 0 {
		problems = append(problems, fmt.Sprintf("%d naming-invariant violation(s)", len(r.NamingViolate)))
	}
	return problems
}

func (r Report) write() {
	fmt.Fprintf(os.Stderr, "== bigdiffer report ==\n")
	fmt.Fprintf(os.Stderr, "available (base) resources: %d\n", r.Total)
	fmt.Fprintf(os.Stderr, "added (new this week):      %d\n", len(r.AddedNew))
	for _, b := range r.AddedNew {
		fmt.Fprintf(os.Stderr, "  + %s  (%s)\n", b.cfn, b.label)
	}
	fmt.Fprintf(os.Stderr, "added (recovered backlog):  %d\n", len(r.AddedBacklog))
	for _, b := range r.AddedBacklog {
		fmt.Fprintf(os.Stderr, "  + %s  (%s)  [available previously but never added]\n", b.cfn, b.label)
	}
	fmt.Fprintf(os.Stderr, "retained (gone from AWS):    %d\n", len(r.Retained))
	if len(r.UnexplainedRetained) > 0 {
		fmt.Fprintf(os.Stderr, "ANOMALY - retained but unexplained (no frozen_since / non_provisionable / checkout pin): %d\n", len(r.UnexplainedRetained))
		for _, b := range r.UnexplainedRetained {
			fmt.Fprintf(os.Stderr, "  ! %s  (%s)\n", b.cfn, b.label)
		}
	}
	if len(r.Duplicates) > 0 {
		fmt.Fprintf(os.Stderr, "ANOMALY - duplicate live blocks for the same CloudFormation type: %d\n", len(r.Duplicates))
		for _, d := range r.Duplicates {
			fmt.Fprintf(os.Stderr, "  ! %s\n", d)
		}
	}
	if len(r.NamingViolate) > 0 {
		fmt.Fprintf(os.Stderr, "ANOMALY - naming invariant violations (resource_type_name != transform(cfn)): %d\n", len(r.NamingViolate))
		for _, v := range r.NamingViolate {
			fmt.Fprintf(os.Stderr, "  ! %s\n", v)
		}
	}
	if len(r.ReasonlessSuppressed) > 0 {
		fmt.Fprintf(os.Stderr, "ANOMALY - suppressed/frozen with no reason recorded: %d (run -heal)\n", len(r.ReasonlessSuppressed))
		for _, b := range r.ReasonlessSuppressed {
			fmt.Fprintf(os.Stderr, "  ! %s  (%s)  [%s]\n", b.cfn, b.label, b.field)
		}
	}
}
