// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"encoding/json"
	"testing"
)

func TestApplySchemaPatches(t *testing.T) {
	tests := []struct {
		name       string
		schemaJSON string
		operations []PatchOperation
		wantJSON   string
		wantErr    bool
	}{
		{
			name: "remove existing property default",
			schemaJSON: `{
				"properties": {
					"PlatformVersion": {
						"type": "string",
						"default": "LATEST"
					}
				}
			}`,
			operations: []PatchOperation{
				{Action: "remove", JSONPath: "/properties/PlatformVersion/default"},
			},
			wantJSON: `{
				"properties": {
					"PlatformVersion": {
						"type": "string"
					}
				}
			}`,
			wantErr: false,
		},
		{
			name: "remove nested property",
			schemaJSON: `{
				"properties": {
					"Config": {
						"type": "object",
						"properties": {
							"Settings": {
								"default": "value",
								"type": "string"
							}
						}
					}
				}
			}`,
			operations: []PatchOperation{
				{Action: "remove", JSONPath: "/properties/Config/properties/Settings/default"},
			},
			wantJSON: `{
				"properties": {
					"Config": {
						"type": "object",
						"properties": {
							"Settings": {
								"type": "string"
							}
						}
					}
				}
			}`,
			wantErr: false,
		},
		{
			name: "multiple remove operations",
			schemaJSON: `{
				"properties": {
					"Prop1": {
						"type": "string",
						"default": "value1"
					},
					"Prop2": {
						"type": "number",
						"default": 42
					}
				}
			}`,
			operations: []PatchOperation{
				{Action: "remove", JSONPath: "/properties/Prop1/default"},
				{Action: "remove", JSONPath: "/properties/Prop2/default"},
			},
			wantJSON: `{
				"properties": {
					"Prop1": {
						"type": "string"
					},
					"Prop2": {
						"type": "number"
					}
				}
			}`,
			wantErr: false,
		},
		{
			name:       "no operations returns original",
			schemaJSON: `{"properties": {"Foo": "bar"}}`,
			operations: []PatchOperation{},
			wantJSON:   `{"properties": {"Foo": "bar"}}`,
			wantErr:    false,
		},
		{
			name:       "nil operations returns original",
			schemaJSON: `{"properties": {"Foo": "bar"}}`,
			operations: nil,
			wantJSON:   `{"properties": {"Foo": "bar"}}`,
			wantErr:    false,
		},
		{
			name:       "remove non-existent path returns error",
			schemaJSON: `{"properties": {}}`,
			operations: []PatchOperation{
				{Action: "remove", JSONPath: "/properties/NonExistent/default"},
			},
			wantJSON: "",
			wantErr:  true,
		},
		{
			name:       "invalid JSON input returns error",
			schemaJSON: `{invalid json`,
			operations: []PatchOperation{
				{Action: "remove", JSONPath: "/properties/Foo"},
			},
			wantJSON: "",
			wantErr:  true,
		},
		{
			name: "remove entire property",
			schemaJSON: `{
				"properties": {
					"Keep": {"type": "string"},
					"Remove": {"type": "string"}
				}
			}`,
			operations: []PatchOperation{
				{Action: "remove", JSONPath: "/properties/Remove"},
			},
			wantJSON: `{
				"properties": {
					"Keep": {"type": "string"}
				}
			}`,
			wantErr: false,
		},
		{
			name: "remove array element",
			schemaJSON: `{
				"required": ["a", "b", "c"]
			}`,
			operations: []PatchOperation{
				{Action: "remove", JSONPath: "/required/1"},
			},
			wantJSON: `{
				"required": ["a", "c"]
			}`,
			wantErr: false,
		},
		{
			name:       "missing action returns error",
			schemaJSON: `{"properties": {"Foo": "bar"}}`,
			operations: []PatchOperation{
				{JSONPath: "/properties/Foo"},
			},
			wantJSON: "",
			wantErr:  true,
		},
		{
			name: "unsupported action returns error",
			schemaJSON: `{"properties": {}}`,
			operations: []PatchOperation{
				{Action: "add", JSONPath: "/properties/Foo"},
			},
			wantJSON: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplySchemaPatches(tt.schemaJSON, tt.operations)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplySchemaPatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			// Compare JSON semantically (ignore formatting differences)
			var gotJSON, wantJSON interface{}
			if err := json.Unmarshal([]byte(got), &gotJSON); err != nil {
				t.Errorf("ApplySchemaPatches() returned invalid JSON: %v", err)
				return
			}
			if err := json.Unmarshal([]byte(tt.wantJSON), &wantJSON); err != nil {
				t.Errorf("Test case has invalid wantJSON: %v", err)
				return
			}

			gotBytes, _ := json.Marshal(gotJSON)
			wantBytes, _ := json.Marshal(wantJSON)
			if string(gotBytes) != string(wantBytes) {
				t.Errorf("ApplySchemaPatches() = %s, want %s", string(gotBytes), string(wantBytes))
			}
		})
	}
}
