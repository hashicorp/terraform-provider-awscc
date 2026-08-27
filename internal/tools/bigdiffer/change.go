// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// byteStatus is the result of comparing a type's freshly discovered schema bytes
// against the on-disk cache. Only New and Changed need gating downstream.
type byteStatus string

const (
	statusNew       byteStatus = "new"       // no cached copy exists yet
	statusChanged   byteStatus = "changed"   // cached bytes differ from AWS
	statusUnchanged byteStatus = "unchanged" // cached bytes identical — skip the gate
	statusFrozen    byteStatus = "frozen"    // frozen_since set — pinned bytes are authoritative
	statusMissing   byteStatus = "missing"   // discovery failed, no bytes to compare
)

type changeResult struct {
	cfType string
	status byteStatus
}

// schemaCachePath returns the cached JSON path for a CloudFormation type name,
// e.g. "AWS::Logs::LogGroup" -> "<cacheDir>/AWS_Logs_LogGroup.json".
func schemaCachePath(cacheDir, cfn string) string {
	return filepath.Join(cacheDir, strings.ReplaceAll(cfn, "::", "_")+".json")
}

// classifyChange decides, from freshly discovered bytes and the cached copy,
// whether a type needs gating. Frozen types are skipped (their pinned bytes are
// authoritative, §4); a discovery failure (nil bytes) surfaces as missing. Both
// inputs are cfschema.Sanitize output, so a plain byte comparison is valid.
func classifyChange(discovered, cached []byte, cachedExists, frozen bool) byteStatus {
	switch {
	case frozen:
		return statusFrozen
	case discovered == nil:
		return statusMissing
	case !cachedExists:
		return statusNew
	case bytes.Equal(discovered, cached):
		return statusUnchanged
	default:
		return statusChanged
	}
}

// detectChanges compares the discovered set against the on-disk schema cache and
// the overlay's frozen set, returning one result per discovered type. Frozen
// types and unchanged types are skipped by the gate; New and Changed proceed.
func detectChanges(disc []discovered, frozenCFN map[string]bool, cacheDir string) ([]changeResult, error) {
	results := make([]changeResult, 0, len(disc))
	for _, d := range disc {
		cfn := d.row.CloudFormationTypeName
		frozen := frozenCFN[cfn]

		var cached []byte
		cachedExists := false
		// Only read the cache when we have fresh bytes to compare and the type
		// is not frozen (frozen bytes are never re-evaluated during refresh).
		if !frozen && d.schema != nil {
			b, err := os.ReadFile(schemaCachePath(cacheDir, cfn))
			switch {
			case err == nil:
				cached, cachedExists = b, true
			case os.IsNotExist(err):
				// no cached copy yet
			default:
				return nil, fmt.Errorf("reading cache for %s: %w", cfn, err)
			}
		}

		results = append(results, changeResult{
			cfType: cfn,
			status: classifyChange(d.schema, cached, cachedExists, frozen),
		})
	}
	return results, nil
}

// changeSummary counts results by status for the change report.
type changeSummary struct {
	New, Changed, Unchanged, Frozen, Missing int
}

func summarize(results []changeResult) changeSummary {
	var s changeSummary
	for _, r := range results {
		switch r.status {
		case statusNew:
			s.New++
		case statusChanged:
			s.Changed++
		case statusUnchanged:
			s.Unchanged++
		case statusFrozen:
			s.Frozen++
		case statusMissing:
			s.Missing++
		}
	}
	return s
}

// gateCandidates returns the CloudFormation type names that need gating
// (New or Changed) — the small set Brick 4 will actually run generation on.
func gateCandidates(results []changeResult) []string {
	var out []string
	for _, r := range results {
		if r.status == statusNew || r.status == statusChanged {
			out = append(out, r.cfType)
		}
	}
	return out
}
