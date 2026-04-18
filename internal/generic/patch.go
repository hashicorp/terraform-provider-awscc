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

// patchDocument returns a JSON Patch document describing the difference between `old` and `new`.
// It sorts remove operations to ensure they are applied in reverse order to avoid index out of bounds errors.
// For key-value arrays (Tags, LoadBalancerAttributes, etc.), it uses full-array replacement instead
// of index-based patches to avoid corruption when CloudFormation returns arrays in different order.
func patchDocument(old, new string) (string, error) {
	patch, err := jsonpatch.CreatePatch([]byte(old), []byte(new))
	if err != nil {
		return "", err
	}

	patch = replaceKeyValueArrayPatchesWithFullReplace(patch, new)

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

// replaceKeyValueArrayPatchesWithFullReplace replaces index-based patch operations targeting
// key-value arrays with full array replacements. CloudFormation does not preserve array ordering
// for key-value structures (objects with "Key"/"key" field), so positional patches target wrong elements.
func replaceKeyValueArrayPatchesWithFullReplace(patch []jsonpatch.JsonPatchOperation, newState string) []jsonpatch.JsonPatchOperation {
	var newDoc map[string]any
	if err := json.Unmarshal([]byte(newState), &newDoc); err != nil {
		return patch
	}

	// Find all array paths that have index-based operations
	arrayPaths := make(map[string]bool)
	for _, op := range patch {
		if arrayPath := extractArrayPath(op.Path); arrayPath != "" {
			arrayPaths[arrayPath] = true
		}
	}

	// Check which arrays are key-value arrays (need to check old state for removals)
	keyValueArrays := make(map[string]bool)
	for path := range arrayPaths {
		// Check if it's a key-value array in new state, or if all ops are removes (array deleted)
		if isKeyValueArray(newDoc, path) {
			keyValueArrays[path] = true
		} else if allOpsAreRemoves(patch, path) {
			// For removed arrays, assume they were key-value if being removed
			keyValueArrays[path] = true
		}
	}

	if len(keyValueArrays) == 0 {
		return patch
	}

	// Filter out index-based operations for key-value arrays
	var filtered []jsonpatch.JsonPatchOperation
	for _, op := range patch {
		arrayPath := extractArrayPath(op.Path)
		if arrayPath != "" && keyValueArrays[arrayPath] {
			continue
		}
		filtered = append(filtered, op)
	}

	// Add full replacements for key-value arrays (sorted for deterministic order)
	paths := make([]string, 0, len(keyValueArrays))
	for path := range keyValueArrays {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		if val := getValueAtPath(newDoc, path); val != nil {
			filtered = append(filtered, jsonpatch.NewPatch("replace", path, val))
		} else {
			filtered = append(filtered, jsonpatch.NewPatch("remove", path, nil))
		}
	}

	return filtered
}

// allOpsAreRemoves checks if all operations for a given array path are remove operations
func allOpsAreRemoves(patch []jsonpatch.JsonPatchOperation, arrayPath string) bool {
	hasOps := false
	for _, op := range patch {
		if strings.HasPrefix(op.Path, arrayPath+"/") || op.Path == arrayPath {
			hasOps = true
			if op.Operation != "remove" {
				return false
			}
		}
	}
	return hasOps
}

// extractArrayPath extracts the array path from an indexed path (e.g., "/Tags/0" -> "/Tags")
func extractArrayPath(path string) string {
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash <= 0 {
		return ""
	}
	lastPart := path[lastSlash+1:]
	if _, err := strconv.Atoi(lastPart); err == nil {
		return path[:lastSlash]
	}
	return ""
}

// isKeyValueArray checks if the array at the given path contains key-value objects
func isKeyValueArray(doc map[string]any, path string) bool {
	val := getValueAtPath(doc, path)
	arr, ok := val.([]any)
	if !ok || len(arr) == 0 {
		return false
	}
	// Check if first element is a map with "Key" or "key" field
	m, ok := arr[0].(map[string]any)
	if !ok {
		return false
	}
	_, hasKey := m["Key"]
	_, haskey := m["key"]
	return hasKey || haskey
}

// getValueAtPath retrieves the value at a JSON pointer path
func getValueAtPath(doc map[string]any, path string) any {
	if path == "" || path == "/" {
		return doc
	}
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	var current any = doc
	for _, part := range parts {
		switch v := current.(type) {
		case map[string]any:
			current = v[part]
		case []any:
			if idx, err := strconv.Atoi(part); err == nil && idx >= 0 && idx < len(v) {
				current = v[idx]
			} else {
				return nil
			}
		default:
			return nil
		}
	}
	return current
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
