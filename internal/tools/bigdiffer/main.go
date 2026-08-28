// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

// Command bigdiffer updates internal/provider/all_schemas.hcl.
//
// The mandate is simple: keep all_schemas.hcl in sync with what AWS offers.
//   - See where we are: parse the curated overlay (all_schemas.hcl).
//   - See what's new: get the available base either by querying AWS live
//     (-discover, us-east-1) or by reading a snapshot file (-available). bigdiffer
//     owns discovery itself (ListTypes + DescribeType), so there is no need to
//     generate and diff checked-in available_schemas.<date>.hcl snapshots.
//   - Reconcile against the base on the CloudFormation type name:
//   - Add rows present in the base but missing from the overlay (recovers
//     backlog and genuinely new resources), copying the resource type name
//     and suppress_plural flag from the base.
//   - Preserve every existing overlay block byte-for-byte, including
//     free-form "# Suppression Reason" comments and suppression attributes.
//   - Re-sort all blocks by CloudFormation type name.
//   - Update the header count to the number of available schemas.
//   - Report anomalies. A retained row (in the overlay, gone from the base) is
//     benign when explained by frozen_since or non_provisionable.
//
// bigdiffer is self-contained: it defines all_schemas.hcl's shape (model.go) and
// its own discovery (discover.go) rather than importing the legacy generator or
// update packages, which are slated for removal.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

const (
	defaultAllSchemas    = "internal/provider/all_schemas.hcl"
	defaultAllSchemasDir = "internal/provider/generators/allschemas"
	defaultCheckout      = "internal/update/suppressions_checkout.txt"

	availablePrefix = "available_schemas."
	availableSuffix = ".hcl"
	dateLayout      = "2006-01-02"

	pluralAttr = "suppress_plural_data_source_generation"
	cfnAttr    = "cloudformation_type_name"
)

var countLineRE = regexp.MustCompile(`(?m)^# \d+ CloudFormation resource types schemas are available[^\n]*$`)

func main() {
	var (
		allSchemasPath = flag.String("all-schemas", defaultAllSchemas, "path to the curated all_schemas.hcl overlay")
		availablePath  = flag.String("available", "", "path to the generated available_schemas.<date>.hcl base (default: newest in -allschemas-dir)")
		allSchemasDir  = flag.String("allschemas-dir", defaultAllSchemasDir, "directory containing available_schemas.<date>.hcl files")
		checkoutPath   = flag.String("checkout", defaultCheckout, "path to suppressions_checkout.txt (cross-referenced, never modified)")
		discoverLive   = flag.Bool("discover", false, "query AWS live (us-east-1) for the available base instead of reading a snapshot file")
		check          = flag.Bool("check", false, "verify all_schemas.hcl is already normalized; exit non-zero if not (writes nothing)")
		generate       = flag.Bool("generate", false, "regenerate the whole provider offline from the committed overlay + schema cache (writes *_gen.go, registrations_gen.go, import_examples_gen.json)")
		update         = flag.Bool("update", false, "live weekly incremental: discover, refresh only changed types from fresh bytes, apply policy to the overlay, write cache + aggregates (needs AWS, us-east-1)")
	)
	flag.Parse()

	if err := run(*allSchemasPath, *availablePath, *allSchemasDir, *checkoutPath, *discoverLive, *check, *generate, *update); err != nil {
		fmt.Fprintf(os.Stderr, "bigdiffer: %v\n", err)
		os.Exit(1)
	}
}

func run(allSchemasPath, availablePath, allSchemasDir, checkoutPath string, discoverLive, check, generate, update bool) error {
	if generate {
		cfg, rows, err := loadOverlay(allSchemasPath)
		if err != nil {
			return err
		}
		return runGenerate(cfg, rows)
	}
	if update {
		return runUpdate(context.Background(), allSchemasPath, checkoutPath)
	}

	// Resolve the base: either query AWS live (bigdiffer owns discovery), or read
	// a snapshot file. Live discovery has no "previous" snapshot, so every base
	// row missing from the overlay is simply reported as newly added.
	var (
		base     []resourceRow
		previous []resourceRow
	)
	if discoverLive {
		disc, err := discover(context.Background())
		if err != nil {
			return fmt.Errorf("live discovery: %w", err)
		}
		base = make([]resourceRow, len(disc))
		for i, d := range disc {
			base[i] = d.row
		}
	} else {
		if availablePath == "" {
			newest, _, err := latestAvailable(allSchemasDir)
			if err != nil {
				return err
			}
			availablePath = newest
		}
		_, previousPath, _ := latestAvailable(allSchemasDir)

		b, err := parseAvailable(availablePath)
		if err != nil {
			return fmt.Errorf("parsing base %s: %w", availablePath, err)
		}
		base = b

		if previousPath != "" && previousPath != availablePath {
			if prev, err := parseAvailable(previousPath); err == nil {
				previous = prev
			}
		}
	}

	overlayContent, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("reading overlay %s: %w", allSchemasPath, err)
	}

	checkout, err := parseCheckout(checkoutPath)
	if err != nil {
		return fmt.Errorf("reading checkout %s: %w", checkoutPath, err)
	}

	out, report, err := normalize(string(overlayContent), base, previous, checkout)
	if err != nil {
		return err
	}

	report.write(os.Stderr)

	if check {
		problems := report.anomalyProblems()
		if out != string(overlayContent) {
			problems = append([]string{"not normalized (run `go run ./internal/tools/bigdiffer` to fix)"}, problems...)
		}
		if len(problems) == 0 {
			fmt.Fprintln(os.Stderr, "bigdiffer: all_schemas.hcl is normalized and anomaly-free.")
			return nil
		}
		return fmt.Errorf("all_schemas.hcl check failed: %s", strings.Join(problems, "; "))
	}

	if out == string(overlayContent) {
		fmt.Fprintln(os.Stderr, "bigdiffer: no changes.")
		return nil
	}
	if err := os.WriteFile(allSchemasPath, []byte(out), 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", allSchemasPath, err)
	}
	fmt.Fprintf(os.Stderr, "bigdiffer: wrote %s.\n", allSchemasPath)
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
	AddedNew            []rowRef // in base, absent from overlay, absent from previous base
	AddedBacklog        []rowRef // in base, absent from overlay, present in previous base
	Retained            []rowRef // live in overlay, absent from base
	UnexplainedRetained []rowRef // Retained rows not explained by frozen_since, non_provisionable, or a checkout pin
	Duplicates          []string // CloudFormation type names with more than one live block
	NamingViolate       []string
	Total               int
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
		if d, ok := decisions[it.key]; ok {
			if attrs := decisionAttrs(d); len(attrs) > 0 {
				mutated, err := setBlockAttributes(text, attrs)
				if err != nil {
					return "", Report{}, fmt.Errorf("applying policy to %s: %w", it.key, err)
				}
				text = mutated
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

// decisionAttrs flattens a policy decision into the attribute set to write on its
// block: the suppression/freeze flags plus suppression_reason when present.
func decisionAttrs(d policyDecision) map[string]string {
	if len(d.setAttrs) == 0 && d.reason == "" {
		return nil
	}
	attrs := make(map[string]string, len(d.setAttrs)+1)
	for k, v := range d.setAttrs {
		attrs[k] = v
	}
	if d.reason != "" {
		attrs[attrSuppressionReason] = d.reason
	}
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
	for _, ln := range strings.Split(text, "\n") {
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

// canonicalBlock renders a new row in the exact format used throughout
// all_schemas.hcl (aligned `=` when the plural-suppression flag is present).
func canonicalBlock(label, cfn string, suppressPlural bool) string {
	if suppressPlural {
		width := len(pluralAttr)
		return fmt.Sprintf("resource_schema %q {\n  %-*s = %q\n  %s = true\n}", label, width, cfnAttr, cfn, pluralAttr)
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

func parseAvailable(path string) ([]resourceRow, error) {
	var s availableFile
	if err := hclsimple.DecodeFile(path, nil, &s); err != nil {
		return nil, err
	}
	return s.Resources, nil
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

// latestAvailable returns the newest and second-newest available_schemas.<date>.hcl
// paths in dir, by embedded date.
func latestAvailable(dir string) (newest, previous string, err error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", "", err
	}
	var dates []string
	for _, e := range entries {
		name := e.Name()
		if !strings.HasPrefix(name, availablePrefix) || !strings.HasSuffix(name, availableSuffix) {
			continue
		}
		d := strings.TrimSuffix(strings.TrimPrefix(name, availablePrefix), availableSuffix)
		if _, perr := time.Parse(dateLayout, d); perr == nil {
			dates = append(dates, d)
		}
	}
	if len(dates) == 0 {
		return "", "", fmt.Errorf("no %s*%s files in %s", availablePrefix, availableSuffix, dir)
	}
	sort.Strings(dates)
	newest = filepath.Join(dir, availablePrefix+dates[len(dates)-1]+availableSuffix)
	if len(dates) >= 2 {
		previous = filepath.Join(dir, availablePrefix+dates[len(dates)-2]+availableSuffix)
	}
	return newest, previous, nil
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

func (r Report) write(w *os.File) {
	fmt.Fprintf(w, "== bigdiffer report ==\n")
	fmt.Fprintf(w, "available (base) resources: %d\n", r.Total)
	fmt.Fprintf(w, "added (new this week):      %d\n", len(r.AddedNew))
	for _, b := range r.AddedNew {
		fmt.Fprintf(w, "  + %s  (%s)\n", b.cfn, b.label)
	}
	fmt.Fprintf(w, "added (recovered backlog):  %d\n", len(r.AddedBacklog))
	for _, b := range r.AddedBacklog {
		fmt.Fprintf(w, "  + %s  (%s)  [available previously but never added]\n", b.cfn, b.label)
	}
	fmt.Fprintf(w, "retained (gone from AWS):    %d\n", len(r.Retained))
	if len(r.UnexplainedRetained) > 0 {
		fmt.Fprintf(w, "ANOMALY - retained but unexplained (no frozen_since / non_provisionable / checkout pin): %d\n", len(r.UnexplainedRetained))
		for _, b := range r.UnexplainedRetained {
			fmt.Fprintf(w, "  ! %s  (%s)\n", b.cfn, b.label)
		}
	}
	if len(r.Duplicates) > 0 {
		fmt.Fprintf(w, "ANOMALY - duplicate live blocks for the same CloudFormation type: %d\n", len(r.Duplicates))
		for _, d := range r.Duplicates {
			fmt.Fprintf(w, "  ! %s\n", d)
		}
	}
	if len(r.NamingViolate) > 0 {
		fmt.Fprintf(w, "ANOMALY - naming invariant violations (resource_type_name != transform(cfn)): %d\n", len(r.NamingViolate))
		for _, v := range r.NamingViolate {
			fmt.Fprintf(w, "  ! %s\n", v)
		}
	}
}
