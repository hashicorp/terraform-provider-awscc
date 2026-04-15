// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"reflect"
	"testing"

	"github.com/mattbaird/jsonpatch"
)

func Test_sortPatchOperations(t *testing.T) {
	tests := []struct {
		name  string
		patch []jsonpatch.JsonPatchOperation
		want  []jsonpatch.JsonPatchOperation
	}{
		{
			name: "index sorted",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/QueueConfigs/6"},
				{Operation: "remove", Path: "/QueueConfigs/4"},
				{Operation: "remove", Path: "/QueueConfigs/11"},
				{Operation: "remove", Path: "/QueueConfigs/10"},
				{Operation: "add", Path: "/QueueConfigs/3", Value: map[string]any{"test": "value"}},
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/QueueConfigs/11"},
				{Operation: "remove", Path: "/QueueConfigs/10"},
				{Operation: "remove", Path: "/QueueConfigs/6"},
				{Operation: "remove", Path: "/QueueConfigs/4"},
				{Operation: "add", Path: "/QueueConfigs/3", Value: map[string]any{"test": "value"}},
			},
		},
		{
			name: "mixed paths",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/QueueConfigs/11"},
				{Operation: "remove", Path: "/QueueConfigs/10"},
				{Operation: "remove", Path: "/QueueConfigs/12"},
				{Operation: "remove", Path: "/OtherQueueConfigs/1"},
				{Operation: "remove", Path: "/OtherQueueConfigs/3"},
				{Operation: "remove", Path: "/ThirdQueueConfigs/2"},
				{Operation: "remove", Path: "/A"},
				{Operation: "remove", Path: "/B"},
				{Operation: "add", Path: "/QueueConfigs/3", Value: map[string]any{"test": "value"}},
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/ThirdQueueConfigs/2"},
				{Operation: "remove", Path: "/QueueConfigs/12"},
				{Operation: "remove", Path: "/QueueConfigs/11"},
				{Operation: "remove", Path: "/QueueConfigs/10"},
				{Operation: "remove", Path: "/OtherQueueConfigs/3"},
				{Operation: "remove", Path: "/OtherQueueConfigs/1"},
				{Operation: "remove", Path: "/B"},
				{Operation: "remove", Path: "/A"},
				{Operation: "add", Path: "/QueueConfigs/3", Value: map[string]any{"test": "value"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortPatchOperations(tt.patch)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPatchOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_comparePathsNumerically(t *testing.T) {
	tests := []struct {
		path1    string
		path2    string
		expected bool
		name     string
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := comparePathsNumerically(tt.path1, tt.path2)
			if result != tt.expected {
				t.Errorf("comparePathsNumerically(%s, %s) = %v, expected %v (%s)",
					tt.path1, tt.path2, result, tt.expected, tt.name)
			}
		})
	}
}

func Test_replaceKeyValueArrayPatchesWithFullReplace(t *testing.T) {
	tests := []struct {
		name     string
		patch    []jsonpatch.JsonPatchOperation
		newState string
		want     []jsonpatch.JsonPatchOperation
	}{
		{
			name: "replaces index-based tag operations with full replace",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Tags/0", Value: map[string]any{"Key": "Apple", "Value": "a"}},
				{Operation: "replace", Path: "/Tags/1", Value: map[string]any{"Key": "Zebra", "Value": "z"}},
				{Operation: "replace", Path: "/Name", Value: "test"},
			},
			newState: `{"Name":"test","Tags":[{"Key":"Apple","Value":"a"},{"Key":"Zebra","Value":"z"}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "test"},
				{Operation: "replace", Path: "/Tags", Value: []any{
					map[string]any{"Key": "Apple", "Value": "a"},
					map[string]any{"Key": "Zebra", "Value": "z"},
				}},
			},
		},
		{
			name: "handles LoadBalancerAttributes key-value array",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/LoadBalancerAttributes/0", Value: map[string]any{"Key": "idle_timeout", "Value": "60"}},
				{Operation: "replace", Path: "/LoadBalancerAttributes/1", Value: map[string]any{"Key": "deletion_protection", "Value": "true"}},
			},
			newState: `{"LoadBalancerAttributes":[{"Key":"idle_timeout","Value":"60"},{"Key":"deletion_protection","Value":"true"}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/LoadBalancerAttributes", Value: []any{
					map[string]any{"Key": "idle_timeout", "Value": "60"},
					map[string]any{"Key": "deletion_protection", "Value": "true"},
				}},
			},
		},
		{
			name: "leaves primitive arrays unchanged",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Ports/0", Value: float64(80)},
				{Operation: "replace", Path: "/Ports/1", Value: float64(443)},
			},
			newState: `{"Ports":[80,443]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Ports/0", Value: float64(80)},
				{Operation: "replace", Path: "/Ports/1", Value: float64(443)},
			},
		},
		{
			name: "removes Tags when not in new state",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/Tags/0"},
				{Operation: "remove", Path: "/Tags/1"},
			},
			newState: `{"Name":"test"}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/Tags", Value: nil},
			},
		},
		{
			name: "no-op when no key-value array operations",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "test"},
				{Operation: "replace", Path: "/Number", Value: float64(42)},
			},
			newState: `{"Name":"test","Number":42}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "test"},
				{Operation: "replace", Path: "/Number", Value: float64(42)},
			},
		},
		{
			name: "handles /Tags path exactly",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Tags", Value: []any{map[string]any{"Key": "a", "Value": "1"}}},
			},
			newState: `{"Tags":[{"Key":"a","Value":"1"}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Tags", Value: []any{map[string]any{"Key": "a", "Value": "1"}}},
			},
		},
		{
			name: "mixed operations with multiple key-value arrays",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "test"},
				{Operation: "add", Path: "/Tags/0", Value: map[string]any{"Key": "New", "Value": "n"}},
				{Operation: "replace", Path: "/Tags/1", Value: map[string]any{"Key": "Old", "Value": "o"}},
				{Operation: "replace", Path: "/Attributes/0", Value: map[string]any{"Key": "attr1", "Value": "val1"}},
				{Operation: "replace", Path: "/Number", Value: float64(42)},
			},
			newState: `{"Name":"test","Tags":[{"Key":"New","Value":"n"},{"Key":"Old","Value":"o"}],"Attributes":[{"Key":"attr1","Value":"val1"}],"Number":42}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "test"},
				{Operation: "replace", Path: "/Number", Value: float64(42)},
				{Operation: "replace", Path: "/Attributes", Value: []any{
					map[string]any{"Key": "attr1", "Value": "val1"},
				}},
				{Operation: "replace", Path: "/Tags", Value: []any{
					map[string]any{"Key": "New", "Value": "n"},
					map[string]any{"Key": "Old", "Value": "o"},
				}},
			},
		},
		{
			name: "handles invalid JSON gracefully",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "test"},
				{Operation: "replace", Path: "/Tags/0", Value: map[string]any{"Key": "a", "Value": "1"}},
			},
			newState: `{invalid json`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "test"},
				{Operation: "replace", Path: "/Tags/0", Value: map[string]any{"Key": "a", "Value": "1"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceKeyValueArrayPatchesWithFullReplace(tt.patch, tt.newState)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceKeyValueArrayPatchesWithFullReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}
