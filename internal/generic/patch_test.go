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
