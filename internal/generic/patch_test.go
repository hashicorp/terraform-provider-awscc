// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"testing"

	"github.com/mattbaird/jsonpatch"
)

func TestSortPatchOperations(t *testing.T) {
	// Test case that reproduces the issue from the GitHub issue
	patch := []jsonpatch.JsonPatchOperation{
		{Operation: "remove", Path: "/QueueConfigs/6"},
		{Operation: "remove", Path: "/QueueConfigs/4"},
		{Operation: "remove", Path: "/QueueConfigs/3"},
		{Operation: "remove", Path: "/QueueConfigs/11"},
		{Operation: "remove", Path: "/QueueConfigs/10"},
		{Operation: "add", Path: "/QueueConfigs/3", Value: map[string]interface{}{"test": "value"}},
	}

	sorted := sortPatchOperations(patch)

	// Verify that remove operations are sorted numerically in descending order
	expectedRemoveOrder := []string{
		"/QueueConfigs/11",
		"/QueueConfigs/10",
		"/QueueConfigs/6",
		"/QueueConfigs/4",
		"/QueueConfigs/3",
	}

	removeOpIndex := 0
	for _, op := range sorted {
		if op.Operation == "remove" {
			if removeOpIndex >= len(expectedRemoveOrder) {
				t.Fatalf("More remove operations than expected")
			}
			if op.Path != expectedRemoveOrder[removeOpIndex] {
				t.Errorf("Expected remove operation %d to have path %s, got %s", 
					removeOpIndex, expectedRemoveOrder[removeOpIndex], op.Path)
			}
			removeOpIndex++
		}
	}

	if removeOpIndex != len(expectedRemoveOrder) {
		t.Errorf("Expected %d remove operations, got %d", len(expectedRemoveOrder), removeOpIndex)
	}
}

func TestComparePathsNumerically(t *testing.T) {
	tests := []struct {
		path1    string
		path2    string
		expected bool
		desc     string
	}{
		{"/QueueConfigs/11", "/QueueConfigs/10", true, "11 > 10"},
		{"/QueueConfigs/10", "/QueueConfigs/6", true, "10 > 6"},
		{"/QueueConfigs/6", "/QueueConfigs/4", true, "6 > 4"},
		{"/QueueConfigs/4", "/QueueConfigs/3", true, "4 > 3"},
		{"/QueueConfigs/3", "/QueueConfigs/11", false, "3 < 11"},
		{"/QueueConfigs/abc", "/QueueConfigs/def", false, "lexical fallback"},
		{"/QueueConfigs/20", "/QueueConfigs/9", true, "20 > 9"},
		{"/QueueConfigs/100", "/QueueConfigs/99", true, "100 > 99"},
		{"/A/1", "/B/2", false, "A < B lexically"},
		{"/QueueConfigs/1", "/OtherConfigs/2", true, "Q > O lexically"},
	}

	for _, test := range tests {
		result := comparePathsNumerically(test.path1, test.path2)
		if result != test.expected {
			t.Errorf("comparePathsNumerically(%s, %s) = %v, expected %v (%s)", 
				test.path1, test.path2, result, test.expected, test.desc)
		}
	}
}

func TestSortPatchOperationsWithLexicalIssue(t *testing.T) {
	// This test reproduces the exact issue where lexical sorting causes problems
	// with indices >= 10
	patch := []jsonpatch.JsonPatchOperation{
		{Operation: "remove", Path: "/QueueConfigs/6"},
		{Operation: "remove", Path: "/QueueConfigs/4"},
		{Operation: "remove", Path: "/QueueConfigs/3"},
		{Operation: "remove", Path: "/QueueConfigs/11"},
		{Operation: "remove", Path: "/QueueConfigs/10"},
		{Operation: "remove", Path: "/QueueConfigs/9"},
		{Operation: "remove", Path: "/QueueConfigs/12"},
		{Operation: "remove", Path: "/QueueConfigs/1"},
		{Operation: "add", Path: "/QueueConfigs/3", Value: map[string]interface{}{"test": "value"}},
	}

	sorted := sortPatchOperations(patch)

	// Expected order: 12, 11, 10, 9, 6, 4, 3, 1 (highest to lowest)
	expectedRemoveOrder := []string{
		"/QueueConfigs/12",
		"/QueueConfigs/11", 
		"/QueueConfigs/10",
		"/QueueConfigs/9",
		"/QueueConfigs/6",
		"/QueueConfigs/4",
		"/QueueConfigs/3",
		"/QueueConfigs/1",
	}

	removeOpIndex := 0
	for _, op := range sorted {
		if op.Operation == "remove" {
			if removeOpIndex >= len(expectedRemoveOrder) {
				t.Fatalf("More remove operations than expected")
			}
			if op.Path != expectedRemoveOrder[removeOpIndex] {
				t.Errorf("Expected remove operation %d to have path %s, got %s", 
					removeOpIndex, expectedRemoveOrder[removeOpIndex], op.Path)
			}
			removeOpIndex++
		}
	}

	if removeOpIndex != len(expectedRemoveOrder) {
		t.Errorf("Expected %d remove operations, got %d", len(expectedRemoveOrder), removeOpIndex)
	}
}
