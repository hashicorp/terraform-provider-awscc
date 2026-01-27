// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

// ApplySchemaPatches applies a list of patch operations to a JSON schema string.
// It uses RFC 6902 JSON Patch format to apply the patches.
// Currently only supports "remove" action (default).
// If no operations are provided, it returns the original schema unchanged.
func ApplySchemaPatches(schemaJSON string, operations []PatchOperation) (string, error) {
	if len(operations) == 0 {
		return schemaJSON, nil
	}

	// Create JSON Patch document
	patchOps := make([]map[string]interface{}, 0, len(operations))
	for _, op := range operations {
		if op.Action == "" {
			return "", fmt.Errorf("action is required for path %q", op.JSONPath)
		}

		// Currently only "remove" is supported
		if op.Action != "remove" {
			return "", fmt.Errorf("unsupported action %q for path %q: only \"remove\" is currently supported", op.Action, op.JSONPath)
		}

		patchOps = append(patchOps, map[string]interface{}{
			"op":   op.Action,
			"path": op.JSONPath,
		})
	}

	// Marshal patch operations to JSON
	patchBytes, err := json.Marshal(patchOps)
	if err != nil {
		return "", fmt.Errorf("failed to marshal patch operations: %w", err)
	}

	// Decode and apply the patch
	patch, err := jsonpatch.DecodePatch(patchBytes)
	if err != nil {
		return "", fmt.Errorf("failed to decode patch: %w", err)
	}

	// Apply patch to schema
	patchedBytes, err := patch.Apply([]byte(schemaJSON))
	if err != nil {
		return "", fmt.Errorf("failed to apply patch: %w", err)
	}

	return string(patchedBytes), nil
}
