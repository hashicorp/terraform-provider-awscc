// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schemaoverrides

import "strings"

var preserveDefaultOmissionPaths = map[string]map[string]struct{}{
	"AWS::BedrockAgentCore::Gateway": {
		"ProtocolType": {},
	},
}

// PreserveDefaultOmission reports whether a CloudFormation schema default must
// not be materialized in a Terraform create plan. These overrides are for
// properties whose API distinguishes omission from the schema default.
func PreserveDefaultOmission(typeName string, path []string) bool {
	paths, ok := preserveDefaultOmissionPaths[typeName]
	if !ok {
		return false
	}

	_, ok = paths[strings.Join(path, "/")]
	return ok
}
