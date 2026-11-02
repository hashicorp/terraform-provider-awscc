// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schemaoverrides

import "testing"

func TestPreserveDefaultOmission(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		typeName string
		path     []string
		want     bool
	}{
		"AgentCore Gateway ProtocolType": {
			typeName: "AWS::BedrockAgentCore::Gateway",
			path:     []string{"ProtocolType"},
			want:     true,
		},
		"different AgentCore Gateway property": {
			typeName: "AWS::BedrockAgentCore::Gateway",
			path:     []string{"AuthorizerType"},
		},
		"different resource": {
			typeName: "AWS::Example::Gateway",
			path:     []string{"ProtocolType"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := PreserveDefaultOmission(test.typeName, test.path); got != test.want {
				t.Errorf("PreserveDefaultOmission() = %t, want %t", got, test.want)
			}
		})
	}
}
