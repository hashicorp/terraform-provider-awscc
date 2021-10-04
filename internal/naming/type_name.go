package naming

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	OrganizationNameAWS = "AWS"
)

var cloudFormationTypeNameRegexp = regexp.MustCompile(`^([a-zA-Z0-9]{2,64})::([a-zA-Z0-9]{2,64})::([a-zA-Z0-9]{2,64})$`)

// ParseCloudFormationTypeName parses a CloudFormation resource type name into 3 parts - Organization, Service and Resource.
// See https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-typeName.
func ParseCloudFormationTypeName(typeName string) (string, string, string, error) {
	matches := cloudFormationTypeNameRegexp.FindStringSubmatch(typeName)

	if got, expected := len(matches), 4; got != expected {
		return "", "", "", fmt.Errorf("matching CloudFormation type name returned %d matches; expected %d", got, expected)
	}

	return matches[1], matches[2], matches[3], nil
}

const terraformTypeNameSeparator = "_"

// CreateTerraformTypeName creates a Terraform resource type name from 3 parts - Organization, Service and Resource.
func CreateTerraformTypeName(org, svc, res string) string {
	parts := []string{org, svc, res}
	id := strings.Join(parts, terraformTypeNameSeparator)

	return id
}

var terraformTypeNameRegexp = regexp.MustCompile(`^([a-zA-Z0-9]{2,64})` + terraformTypeNameSeparator + `([a-zA-Z0-9]{2,64})` + terraformTypeNameSeparator + `([a-zA-Z0-9_]{2,})$`)

// ParseTerraformTypeName parses a Terraform resource type name into 3 parts - Organization, Service and Resource.
func ParseTerraformTypeName(typeName string) (string, string, string, error) {
	matches := terraformTypeNameRegexp.FindStringSubmatch(typeName)

	if got, expected := len(matches), 4; got != expected {
		return "", "", "", fmt.Errorf("matching Terraform type name returned %d matches; expected %d", got, expected)
	}

	return matches[1], matches[2], matches[3], nil
}
