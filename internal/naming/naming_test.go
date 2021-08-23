package naming_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
)

func TestCloudFormationPropertyToTerraformAttribute(t *testing.T) {
	testCases := []struct {
		TestName      string
		Value         string
		ExpectedValue string
	}{
		{
			TestName:      "empty string",
			Value:         "",
			ExpectedValue: "",
		},
		{
			TestName:      "whitespace string",
			Value:         "  ",
			ExpectedValue: "",
		},
		{
			TestName:      "short property name",
			Value:         "Arn",
			ExpectedValue: "arn",
		},
		{
			TestName:      "long property name",
			Value:         "GlobalReplicationGroupDescription",
			ExpectedValue: "global_replication_group_description",
		},
		{
			TestName:      "including digit",
			Value:         "S3Bucket",
			ExpectedValue: "s3_bucket",
		},
		{
			TestName:      "including multiple digits",
			Value:         "AWS99Thing",
			ExpectedValue: "aws99_thing",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got := naming.CloudFormationPropertyToTerraformAttribute(testCase.Value)

			if got != testCase.ExpectedValue {
				t.Errorf("expected: %s, got: %s", testCase.ExpectedValue, got)
			}
		})
	}
}
