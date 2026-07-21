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
		name                string
		patch               []jsonpatch.JsonPatchOperation
		newState            string
		unorderedArrayPaths map[string]bool
		want                []jsonpatch.JsonPatchOperation
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
		{
			// Reproduces issue #3260: Config elements have no "Key"/"key" field, so
			// isKeyValueArray alone would miss them. unorderedArrayPaths (derived from
			// the attribute's Set type in the Terraform schema) must catch this instead.
			name: "handles Set-typed non-key-value array via unorderedArrayPaths",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Config/0/StartTime/Hours", Value: float64(8)},
			},
			newState:            `{"Config":[{"Day":"MONDAY","StartTime":{"Hours":8,"Minutes":0}},{"Day":"TUESDAY","StartTime":{"Hours":10,"Minutes":0}},{"Day":"WEDNESDAY","StartTime":{"Hours":11,"Minutes":0}}]}`,
			unorderedArrayPaths: map[string]bool{"/Config": true},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Config", Value: []any{
					map[string]any{"Day": "MONDAY", "StartTime": map[string]any{"Hours": float64(8), "Minutes": float64(0)}},
					map[string]any{"Day": "TUESDAY", "StartTime": map[string]any{"Hours": float64(10), "Minutes": float64(0)}},
					map[string]any{"Day": "WEDNESDAY", "StartTime": map[string]any{"Hours": float64(11), "Minutes": float64(0)}},
				}},
			},
		},
		{
			// Without a matching unorderedArrayPaths entry, non-key-value arrays are
			// left as positional patches -- confirms the new signal is additive, not a
			// blanket full-replace for every array.
			name: "leaves non-key-value array unchanged when path not in unorderedArrayPaths",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Config/0/StartTime/Hours", Value: float64(8)},
			},
			newState: `{"Config":[{"Day":"MONDAY","StartTime":{"Hours":8,"Minutes":0}}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Config/0/StartTime/Hours", Value: float64(8)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceKeyValueArrayPatchesWithFullReplace(tt.patch, tt.newState, tt.unorderedArrayPaths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceKeyValueArrayPatchesWithFullReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_replaceKeyValueArrayPatchesWithFullReplace_Nested covers arrays nested inside
// another array's elements (e.g. AWS::EC2::LaunchTemplate's
// TagSpecifications[].Tags, or a hypothetical Set nested inside an ordered list). A
// patch operation into such a structure crosses two array boundaries, and it's the
// inner one -- not the outer, ordinary-list one -- that needs full-replace protection.
// This guards against a regression where only the outermost (or only the innermost)
// index in a path was considered, which would either miss the real order-unstable
// array or misidentify an unrelated ordered container as needing full replacement.
func Test_replaceKeyValueArrayPatchesWithFullReplace_Nested(t *testing.T) {
	tests := []struct {
		name                string
		patch               []jsonpatch.JsonPatchOperation
		newState            string
		unorderedArrayPaths map[string]bool
		want                []jsonpatch.JsonPatchOperation
	}{
		{
			name: "key-value array nested inside an ordered list gets full replace, not the outer list",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "add", Path: "/TagSpecifications/0/Tags/1", Value: map[string]any{"Key": "team", "Value": "a"}},
			},
			newState: `{"TagSpecifications":[{"ResourceType":"instance","Tags":[{"Key":"env","Value":"prod"},{"Key":"team","Value":"a"}]}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/TagSpecifications/0/Tags", Value: []any{
					map[string]any{"Key": "env", "Value": "prod"},
					map[string]any{"Key": "team", "Value": "a"},
				}},
			},
		},
		{
			name: "removing one nested tag doesn't spuriously full-replace the outer list too",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/TagSpecifications/0/Tags/1"},
			},
			newState: `{"TagSpecifications":[{"ResourceType":"instance","Tags":[{"Key":"env","Value":"prod"},{"Key":"owner","Value":"x"}]}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/TagSpecifications/0/Tags", Value: []any{
					map[string]any{"Key": "env", "Value": "prod"},
					map[string]any{"Key": "owner", "Value": "x"},
				}},
			},
		},
		{
			name: "Set-typed attribute nested inside an ordered list's elements gets full replace",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "add", Path: "/QueueConfigs/0/Tags/1", Value: map[string]any{"Foo": "b"}},
			},
			newState:            `{"QueueConfigs":[{"Name":"q1","Tags":[{"Foo":"a"},{"Foo":"b"}]}]}`,
			unorderedArrayPaths: map[string]bool{"/QueueConfigs/Tags": true},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/QueueConfigs/0/Tags", Value: []any{
					map[string]any{"Foo": "a"},
					map[string]any{"Foo": "b"},
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceKeyValueArrayPatchesWithFullReplace(tt.patch, tt.newState, tt.unorderedArrayPaths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceKeyValueArrayPatchesWithFullReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolveMutuallyExclusiveProperties(t *testing.T) {
	oldState := `{
		"Description": "old",
		"Rules": [
			{
				"Name": "r1",
				"Statement": {
					"ByteMatchStatement": {
						"SearchString": "marker",
						"SearchStringBase64": "bWFya2Vy",
						"PositionalConstraint": "CONTAINS"
					}
				}
			}
		]
	}`
	oldStateSingle := `{
		"Rules": [
			{
				"Statement": {
					"ByteMatchStatement": {
						"SearchStringBase64": "bWFya2Vy"
					}
				}
			}
		]
	}`

	tests := []struct {
		name     string
		oldState string
		patch    []jsonpatch.JsonPatchOperation
		want     []jsonpatch.JsonPatchOperation
	}{
		{
			name:     "unrelated change drops derived member",
			oldState: oldState,
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "new"},
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "new"},
				{Operation: "remove", Path: "/Rules/0/Statement/ByteMatchStatement/SearchString"},
			},
		},
		{
			name:     "changing derived member drops the other",
			oldState: oldState,
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Rules/0/Statement/ByteMatchStatement/SearchString", Value: "changed"},
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Rules/0/Statement/ByteMatchStatement/SearchString", Value: "changed"},
				{Operation: "remove", Path: "/Rules/0/Statement/ByteMatchStatement/SearchStringBase64"},
			},
		},
		{
			name:     "both changed is left alone",
			oldState: oldState,
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Rules/0/Statement/ByteMatchStatement/SearchString", Value: "a"},
				{Operation: "replace", Path: "/Rules/0/Statement/ByteMatchStatement/SearchStringBase64", Value: "YQ=="},
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Rules/0/Statement/ByteMatchStatement/SearchString", Value: "a"},
				{Operation: "replace", Path: "/Rules/0/Statement/ByteMatchStatement/SearchStringBase64", Value: "YQ=="},
			},
		},
		{
			name:     "single member present is untouched",
			oldState: oldStateSingle,
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "new"},
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "new"},
			},
		},
		{
			name:     "invalid old state is a no-op",
			oldState: "not json",
			patch: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "new"},
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "new"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveMutuallyExclusiveProperties(tt.patch, tt.oldState)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveMutuallyExclusiveProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendMutuallyExclusiveResolutionsForModel(t *testing.T) {
	// Remote model contains a pair inside content the patch does not touch
	// (e.g. a statement subtree beyond a depth-limited schema's maximum depth).
	model := `{
		"Description": "old",
		"Rules": [
			{
				"Statement": {
					"NotStatement": {
						"Statement": {
							"OrStatement": {
								"Statements": [
									{
										"ByteMatchStatement": {
											"SearchString": "marker",
											"SearchStringBase64": "bWFya2Vy"
										}
									}
								]
							}
						}
					}
				}
			}
		]
	}`

	got, err := appendMutuallyExclusiveResolutionsForModel(`[{"op":"replace","path":"/Description","value":"new"}]`, model)
	if err != nil {
		t.Fatal(err)
	}
	want := `[{"op":"replace","path":"/Description","value":"new"},{"op":"remove","path":"/Rules/0/Statement/NotStatement/Statement/OrStatement/Statements/0/ByteMatchStatement/SearchString"}]`
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// Pair already resolved by the state-based pass: no duplicate remove.
	resolved := `[{"op":"remove","path":"/Rules/0/Statement/NotStatement/Statement/OrStatement/Statements/0/ByteMatchStatement/SearchString"}]`
	got, err = appendMutuallyExclusiveResolutionsForModel(resolved, model)
	if err != nil {
		t.Fatal(err)
	}
	if got != resolved {
		t.Errorf("expected unchanged document, got %s", got)
	}

	// Invalid inputs are a no-op.
	got, err = appendMutuallyExclusiveResolutionsForModel("[]", "not json")
	if err != nil || got != "[]" {
		t.Errorf("invalid model: got %s, err %v", got, err)
	}
}
