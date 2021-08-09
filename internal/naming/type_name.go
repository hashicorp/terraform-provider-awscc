package naming

import (
	"fmt"
	"regexp"
)

var cloudFormationTypeNameRegexp = regexp.MustCompile(`^([a-zA-Z0-9]{2,64})::([a-zA-Z0-9]{2,64})::([a-zA-Z0-9]{2,64})$`)

// ParseCloudFormationTypeName parses a CloudFormation resource type name into 3 parts - Organization, Service and Resource.
func ParseCloudFormationTypeName(typeName string) (string, string, string, error) {
	matches := cloudFormationTypeNameRegexp.FindStringSubmatch(typeName)

	if got, expected := len(matches), 4; got != expected {
		return "", "", "", fmt.Errorf("matching CloudFormation type name returned %d matches; expected %d", got, expected)
	}

	return matches[1], matches[2], matches[3], nil
}
