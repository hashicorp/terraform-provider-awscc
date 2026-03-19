// Copyright IBM Corp. 2021, 2026
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

// tagsPathPrefix is the JSON Pointer path for the Tags property in Cloud Control desired state.
// CloudFormation does not preserve tag array ordering, so index-based patch operations
// (e.g. replace /Tags/5) target the wrong tag. We replace any Tags-related operations
// with a single full replacement of the entire Tags array.
const tagsPathPrefix = "/Tags"

// patchDocument returns a JSON Patch document describing the difference between `old` and `new`.
// It sorts remove operations to ensure they are applied in reverse order to avoid index out of bounds errors.
// For the Tags property, it uses full-array replacement instead of index-based patches to avoid
// tag corruption when CloudFormation returns tags in a different order than Terraform state.
func patchDocument(old, new string) (string, error) {
	patch, err := jsonpatch.CreatePatch([]byte(old), []byte(new))
	if err != nil {
		return "", err
	}

	patch = replaceTagsPatchWithFullReplace(patch, new)

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

// replaceTagsPatchWithFullReplace replaces any index-based patch operations targeting the Tags
// array with a single full replacement (or remove). This avoids tag corruption because
// CloudFormation does not preserve tag array ordering, so positional patches target wrong keys.
func replaceTagsPatchWithFullReplace(patch []jsonpatch.JsonPatchOperation, newState string) []jsonpatch.JsonPatchOperation {
	var filtered []jsonpatch.JsonPatchOperation
	var hasTagsOps bool
	for _, op := range patch {
		if op.Path == tagsPathPrefix || strings.HasPrefix(op.Path, tagsPathPrefix+"/") {
			hasTagsOps = true
			continue
		}
		filtered = append(filtered, op)
	}
	if !hasTagsOps {
		return patch
	}

	var newDoc map[string]any
	if err := json.Unmarshal([]byte(newState), &newDoc); err != nil {
		// If we can't parse new state, return filtered patch (Tags ops already removed)
		return filtered
	}
	if tagsVal, ok := newDoc["Tags"]; ok {
		filtered = append(filtered, jsonpatch.NewPatch("replace", tagsPathPrefix, tagsVal))
	} else {
		filtered = append(filtered, jsonpatch.NewPatch("remove", tagsPathPrefix, nil))
	}
	return filtered
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
