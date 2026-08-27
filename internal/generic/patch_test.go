// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"encoding/json"
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

func Test_resolveEmptyReadArtifacts(t *testing.T) {
	resourceModel := `{
		"Description": "d",
		"Rules": [
			{
				"Name": "r1",
				"Action": {
					"Block": {
						"CustomResponse": {
							"ResponseCode": 403,
							"ResponseHeaders": []
						}
					}
				}
			}
		]
	}`
	oldState := `{
		"Description": "d",
		"Rules": [
			{
				"Name": "r1",
				"Action": {
					"Block": {
						"CustomResponse": {
							"ResponseCode": 403
						}
					}
				}
			}
		]
	}`
	newState := `{
		"Description": "d",
		"Rules": [
			{
				"Name": "r1",
				"Action": {
					"Block": {
						"CustomResponse": {
							"ResponseCode": 404
						}
					}
				}
			}
		]
	}`

	tests := []struct {
		name          string
		patchDocument string
		oldState      string
		newState      string
		resourceModel string
		isMutable     func(tokens []string) bool
		want          []jsonpatch.JsonPatchOperation
	}{
		{
			name:          "read artifact is removed",
			patchDocument: `[{"op":"replace","path":"/Rules/0/Action/Block/CustomResponse/ResponseCode","value":404}]`,
			oldState:      oldState,
			newState:      newState,
			resourceModel: resourceModel,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/Rules/0/Action/Block/CustomResponse/ResponseHeaders"},
				{Operation: "replace", Path: "/Rules/0/Action/Block/CustomResponse/ResponseCode", Value: float64(404)},
			},
		},
		{
			name:          "empty array present in prior state is kept",
			patchDocument: `[{"op":"replace","path":"/Description","value":"e"}]`,
			oldState:      `{"Description":"d","Rules":[{"Name":"r1","Action":{"Block":{"CustomResponse":{"ResponseCode":403,"ResponseHeaders":[]}}}}]}`,
			newState:      newState,
			resourceModel: resourceModel,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "e"},
			},
		},
		{
			name:          "empty array present in planned state is kept",
			patchDocument: `[{"op":"replace","path":"/Description","value":"e"}]`,
			oldState:      oldState,
			newState:      `{"Description":"e","Rules":[{"Name":"r1","Action":{"Block":{"CustomResponse":{"ResponseCode":403,"ResponseHeaders":[]}}}}]}`,
			resourceModel: resourceModel,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "e"},
			},
		},
		{
			name:          "model without empty arrays leaves patch unchanged",
			patchDocument: `[{"op":"replace","path":"/Description","value":"e"}]`,
			oldState:      oldState,
			newState:      newState,
			resourceModel: `{"Description":"d","Rules":[{"Name":"r1","Action":{"Block":{"CustomResponse":{"ResponseCode":403,"ResponseHeaders":[{"Name":"n","Value":"v"}]}}}}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "e"},
			},
		},
		{
			name:          "multiple artifacts are removed in reverse path order",
			patchDocument: `[{"op":"replace","path":"/Description","value":"e"}]`,
			oldState:      `{"Description":"d"}`,
			newState:      `{"Description":"e"}`,
			resourceModel: `{"Description":"d","A":[],"B":{"C":[]}}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/B/C"},
				{Operation: "remove", Path: "/A"},
				{Operation: "replace", Path: "/Description", Value: "e"},
			},
		},
		{
			name:          "string and array artifacts as returned by a live WAF web ACL",
			patchDocument: `[{"op":"replace","path":"/Rules/0/Action/Block/CustomResponse/ResponseCode","value":404}]`,
			oldState:      `{"Rules":[{"Name":"r1","Action":{"Block":{"CustomResponse":{"ResponseCode":403}}}}]}`,
			newState:      `{"Rules":[{"Name":"r1","Action":{"Block":{"CustomResponse":{"ResponseCode":404}}}}]}`,
			resourceModel: `{"Description":"","Rules":[{"Name":"r1","RuleLabels":[],"Action":{"Block":{"CustomResponse":{"ResponseCode":403,"ResponseHeaders":[]}}}}]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/Rules/0/RuleLabels"},
				{Operation: "remove", Path: "/Rules/0/Action/Block/CustomResponse/ResponseHeaders"},
				{Operation: "remove", Path: "/Description"},
				{Operation: "replace", Path: "/Rules/0/Action/Block/CustomResponse/ResponseCode", Value: float64(404)},
			},
		},
		{
			name:          "empty string present in planned state is kept",
			patchDocument: `[{"op":"add","path":"/Description","value":""}]`,
			oldState:      `{"Rules":[]}`,
			newState:      `{"Description":"","Rules":[]}`,
			resourceModel: `{"Description":"","Rules":[]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "add", Path: "/Description", Value: ""},
			},
		},
		{
			name:          "empty string array elements are never removed",
			patchDocument: `[{"op":"replace","path":"/Name","value":"n"}]`,
			oldState:      `{"Name":"m"}`,
			newState:      `{"Name":"n"}`,
			resourceModel: `{"Name":"m","A":["","x"]}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "n"},
			},
		},
		{
			name:          "immutable properties are never removed",
			patchDocument: `[{"op":"replace","path":"/Name","value":"n"}]`,
			oldState:      `{"Name":"m"}`,
			newState:      `{"Name":"n"}`,
			resourceModel: `{"Name":"m","FailureReasons":[],"Detail":""}`,
			isMutable:     func([]string) bool { return false },
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "n"},
			},
		},
		{
			name:          "keys containing JSON Pointer special characters are looked up correctly",
			patchDocument: `[{"op":"replace","path":"/Name","value":"n"}]`,
			oldState:      `{"Name":"m","Attributes":{"a/b":""}}`,
			newState:      `{"Name":"n","Attributes":{"a/b":""}}`,
			resourceModel: `{"Name":"m","Attributes":{"a/b":""}}`,
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Name", Value: "n"},
			},
		},
		{
			name:          "real schema predicate: mutable empty array removed, read-only empty string kept",
			patchDocument: `[{"op":"replace","path":"/Name","value":"n"}]`,
			oldState:      `{"Name":"m"}`,
			newState:      `{"Name":"n"}`,
			resourceModel: `{"Name":"m","Ports":[],"Arn":""}`,
			isMutable: func(tokens []string) bool {
				r := &genericResource{tfSchema: testSimpleSchemaWithList, cfToTfNameMap: simpleCfToTfNameMap}
				return r.isMutablePropertyPath(context.TODO(), tokens)
			},
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "remove", Path: "/Ports"},
				{Operation: "replace", Path: "/Name", Value: "n"},
			},
		},
		{
			name:          "invalid resource model is a no-op",
			patchDocument: `[{"op":"replace","path":"/Description","value":"e"}]`,
			oldState:      oldState,
			newState:      newState,
			resourceModel: "not json",
			want: []jsonpatch.JsonPatchOperation{
				{Operation: "replace", Path: "/Description", Value: "e"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isMutable := tt.isMutable
			if isMutable == nil {
				isMutable = func([]string) bool { return true }
			}
			got, err := resolveEmptyReadArtifacts(tt.patchDocument, tt.oldState, tt.newState, tt.resourceModel, isMutable)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			var gotOps []jsonpatch.JsonPatchOperation
			if err := json.Unmarshal([]byte(got), &gotOps); err != nil {
				t.Fatalf("unmarshaling result %q: %s", got, err)
			}
			if !reflect.DeepEqual(gotOps, tt.want) {
				t.Errorf("resolveEmptyReadArtifacts() = %v, want %v", gotOps, tt.want)
			}
		})
	}
}

func Test_isMutablePropertyPath(t *testing.T) {
	tests := []struct {
		name     string
		resource *genericResource
		tokens   []string
		want     bool
	}{
		{
			name:     "required attribute",
			resource: &genericResource{tfSchema: testSimpleSchema, cfToTfNameMap: simpleCfToTfNameMap},
			tokens:   []string{"Name"},
			want:     true,
		},
		{
			name:     "optional attribute",
			resource: &genericResource{tfSchema: testSimpleSchema, cfToTfNameMap: simpleCfToTfNameMap},
			tokens:   []string{"Number"},
			want:     true,
		},
		{
			name:     "read-only (Computed-only) attribute",
			resource: &genericResource{tfSchema: testSimpleSchema, cfToTfNameMap: simpleCfToTfNameMap},
			tokens:   []string{"Arn"},
			want:     false,
		},
		{
			name:     "unknown property",
			resource: &genericResource{tfSchema: testSimpleSchema, cfToTfNameMap: simpleCfToTfNameMap},
			tokens:   []string{"Bogus"},
			want:     false,
		},
		{
			name:     "nested attribute in a list element",
			resource: &genericResource{tfSchema: testComplexSchema, cfToTfNameMap: complexCfToTfNameMap},
			tokens:   []string{"Disks", "0", "DeleteWithInstance"},
			want:     true,
		},
		{
			name:     "nested attribute in a single-nested attribute",
			resource: &genericResource{tfSchema: testComplexSchema, cfToTfNameMap: complexCfToTfNameMap},
			tokens:   []string{"BootDisk", "Id"},
			want:     true,
		},
		{
			name:     "non-numeric index into a list",
			resource: &genericResource{tfSchema: testComplexSchema, cfToTfNameMap: complexCfToTfNameMap},
			tokens:   []string{"Disks", "x", "Id"},
			want:     false,
		},
		{
			name:     "attribute in a map-nested attribute",
			resource: &genericResource{tfSchema: testMapsSchema, cfToTfNameMap: mapsCfToTfNameMap},
			tokens:   []string{"ComplexMap", "x", "Flags"},
			want:     true,
		},
		{
			name:     "structure inside a JSON-string-typed attribute",
			resource: &genericResource{tfSchema: testMapsSchema, cfToTfNameMap: mapsCfToTfNameMap},
			tokens:   []string{"JsonString", "Inner"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.resource.isMutablePropertyPath(context.TODO(), tt.tokens); got != tt.want {
				t.Errorf("isMutablePropertyPath(%v) = %v, want %v", tt.tokens, got, tt.want)
			}
		})
	}
}
