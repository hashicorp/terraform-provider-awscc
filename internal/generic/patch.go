// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mattbaird/jsonpatch"
)

// patchDocument returns a JSON Patch document describing the difference between `old` and `new`.
// It sorts remove operations to ensure they are applied in reverse order to avoid index out of bounds errors.
func patchDocument(old, new string) (string, error) {
	patch, err := jsonpatch.CreatePatch([]byte(old), []byte(new))
	if err != nil {
		return "", err
	}

	// Sort the patch operations to ensure remove operations are applied in reverse order
	sortedPatch := sortPatchOperations(patch)

	// Ensure we always have a valid JSON array, even if empty
	if len(sortedPatch) == 0 {
		return "[]", nil
	}

	b, err := json.Marshal(sortedPatch)
	if err != nil {
		return "", err
	}

	// Verify that the marshaled JSON starts with '[' to ensure it's a valid JSON array
	result := string(b)
	if !strings.HasPrefix(result, "[") {
		return "[]", fmt.Errorf("generated patch document is not a valid JSON array: %s", result)
	}

	return result, nil
}

// sortPatchOperations sorts the patch operations to ensure that remove operations
// are applied in reverse order (highest index first) to avoid index out of bounds errors.
func sortPatchOperations(patch []jsonpatch.JsonPatchOperation) []jsonpatch.JsonPatchOperation {
	// First, separate remove operations from other operations
	var removeOps []jsonpatch.JsonPatchOperation
	var otherOps []jsonpatch.JsonPatchOperation

	for _, op := range patch {
		if op.Operation == "remove" {
			removeOps = append(removeOps, op)
		} else {
			otherOps = append(otherOps, op)
		}
	}

	// Sort remove operations by path in reverse order, handling numeric indices correctly
	sort.Slice(removeOps, func(i, j int) bool {
		return comparePathsNumerically(removeOps[i].Path, removeOps[j].Path)
	})

	// Combine the operations back together with remove operations first
	return append(removeOps, otherOps...)
}

// comparePathsNumerically compares two JSON patch paths, treating array indices as numbers
//
// Returns true if path1 should come before path2 (higher indices first for removal).
func comparePathsNumerically(path1, path2 string) bool {
	parts1 := strings.Split(path1, "/")
	parts2 := strings.Split(path2, "/")

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		if idx1, err1 := strconv.Atoi(parts1[i]); err1 == nil {
			if idx2, err2 := strconv.Atoi(parts2[i]); err2 == nil {
				if idx1 != idx2 {
					return idx1 > idx2
				}
				continue
			}
		}
		if parts1[i] != parts2[i] {
			return parts1[i] > parts2[i]
		}
	}
	return len(parts1) > len(parts2)
}
