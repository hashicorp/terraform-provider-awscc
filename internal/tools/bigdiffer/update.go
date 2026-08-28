// SPDX-License-Identifier: MPL-2.0

package main

// The incremental weekly pipeline: from a single discovery crawl, refresh only
// the types whose sanitized schema bytes changed (New or Changed), generate them
// for real from the fresh bytes, and let generation success/failure drive the
// policy decision written back to the overlay. Generation is the gate: if a type
// still generates, it is refreshed and its cache promoted; if it breaks, its
// last-good output and cache are kept (never regress) and the policy freezes or
// suppresses it so the release is never blocked.

// candidate is a type selected for regeneration because its discovered bytes
// differ from the cache (New or Changed). It carries the fresh bytes plus the
// row and class that drive generation and the policy decision.
type candidate struct {
	cfType string
	class  changeClass
	row    resourceRow
	schema []byte // freshly discovered, sanitized bytes to generate from
}

// buildCandidates selects the New+Changed types and attaches, for each, its
// policy class (new to the overlay vs already present), the row to generate from
// (the overlay row when present so its suppress_* flags are honored, else the
// discovered row), and the fresh discovered bytes.
func buildCandidates(results []changeResult, discByCFN map[string]discovered, overlayByCFN map[string]resourceRow) []candidate {
	var out []candidate
	for _, r := range results {
		if r.status != statusNew && r.status != statusChanged {
			continue
		}
		d := discByCFN[r.cfType]
		c := candidate{cfType: r.cfType, schema: d.schema}
		if row, ok := overlayByCFN[r.cfType]; ok {
			c.class = classPresent
			c.row = row
		} else {
			c.class = classNew
			c.row = d.row
		}
		out = append(out, c)
	}
	return out
}

// gateResultFromGenResults converts a candidate's real generation outcomes into
// a gateResult for the policy engine: any artifact that failed to generate marks
// the type as broken, so decide() freezes (Present) or suppresses (New) it.
// Plan-level errors (empty artifact kind) count as failures too.
func gateResultFromGenResults(cfType string, results []genResult) gateResult {
	gr := gateResult{cfType: cfType}
	for _, r := range results {
		outcome := gateOK
		if r.err != nil {
			outcome = gateFailedGeneration
		}
		gr.artifacts = append(gr.artifacts, artifactResult{
			kind:    r.a.kind,
			outcome: outcome,
			err:     r.err,
		})
	}
	return gr
}
