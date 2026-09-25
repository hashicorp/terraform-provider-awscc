// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Output routing. structOut carries the structural, human-facing lines
// (stepf/infof and the report) to the console and, while a run log is open, to
// the log as well. logOnly carries verbose detail (external-tool output) to the
// log only. Progress bars write straight to os.Stderr and are never sent to the
// log — their carriage-return redraws would corrupt the file.
var (
	structOut io.Writer = os.Stderr
	logOnly   io.Writer = io.Discard
)

// startRunLog opens a gitignored .bigdiffer-<command>.log in the working
// directory (the repo root, where bigdiffer runs), tees structural output to it
// and routes verbose external-tool output to it only. It returns a closer that
// restores the console-only defaults. On failure it warns and routes logOnly
// to stderr too, so the run degrades to genuinely console-only — nothing an
// external tool prints is silently dropped — rather than discarding
// everything logOnly would have carried.
func startRunLog(command string) func() {
	name := ".bigdiffer-" + command + ".log"
	f, err := os.Create(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "    (could not open %s: %v — console only)\n", name, err)
		logOnly = os.Stderr
		return func() { logOnly = io.Discard }
	}
	structOut = io.MultiWriter(os.Stderr, f)
	logOnly = f
	stepf("Full output logged to %s", name)
	return func() {
		structOut = os.Stderr
		logOnly = io.Discard
		_ = f.Close()
	}
}

// commandName maps the selected command flag to the log-file suffix, or "" when
// no command was chosen (so no log is opened for a usage error).
func commandName(sync, reconcile, check, docs, lint, recheck bool) string {
	switch {
	case sync:
		return "sync"
	case reconcile:
		return "reconcile"
	case check:
		return "check"
	case docs:
		return "docs"
	case lint:
		return "lint"
	case recheck:
		return "recheck"
	default:
		return ""
	}
}

// printActionItems prints the end-of-run to-do list: every issue-worthy
// suppression this run (a real generation or build failure worth a GitHub
// issue), each with the artifact kind and one-line reason. It surfaces
// at the bottom of the console, where it survives however long the run's
// output was. Prints nothing when there is nothing to act on.
//
// A partial failure (policy.go's classPresent, anyOK branch) records reasons
// for both the broken artifact(s) (suppression_reason_resource/
// _singular_data_source/_plural_data_source) and the shared-schema freeze
// itself (frozen_reason) — the freeze's rationale is the generic "one JSON
// backs three artifacts, so a break in one pins the whole schema" invariant,
// not any specific artifact's bug. The action-item worth surfacing is the
// specific one: every qualifying artifact reason, not just the first found.
// frozen_reason is shown only when it's the sole qualifying reason for that
// type (a total failure, policy.go's non-anyOK branch), where it's the only
// explanation there is.
func printActionItems(decisions map[string]policyDecision) {
	issueWorthy := func(reason string) bool {
		return strings.HasPrefix(reason, string(reasonGenerationFailed)+":") ||
			strings.HasPrefix(reason, string(reasonBuildFailed)+":")
	}
	type item struct{ cfType, kind, reason string }
	artifactKinds := []struct{ attr, label string }{
		{attrSuppressionReasonResource, "resource"},
		{attrSuppressionReasonSingular, "singular data source"},
		{attrSuppressionReasonPlural, "plural data source"},
	}
	var items []item
	for cfType, d := range decisions {
		var artifact, frozen []item
		for _, k := range artifactKinds {
			if reason, ok := d.reasons[k.attr]; ok && issueWorthy(reason) {
				artifact = append(artifact, item{cfType: cfType, kind: k.label, reason: reason})
			}
		}
		if reason, ok := d.reasons[attrFrozenReason]; ok && issueWorthy(reason) {
			frozen = append(frozen, item{cfType: cfType, kind: "shared schema", reason: reason})
		}
		// A partial failure lists the specific broken artifact(s); a total
		// failure has only the shared-schema freeze to point at.
		chosen := artifact
		if len(chosen) == 0 {
			chosen = frozen
		}
		items = append(items, chosen...)
	}
	if len(items) == 0 {
		return
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].cfType != items[j].cfType {
			return items[i].cfType < items[j].cfType
		}
		if items[i].kind != items[j].kind {
			return items[i].kind < items[j].kind
		}
		return items[i].reason < items[j].reason
	})
	stepf("Action items — open a GitHub issue for each failed generation or build:")
	for _, it := range items {
		infof("%s (%s) — %s", it.cfType, it.kind, it.reason)
	}
}
