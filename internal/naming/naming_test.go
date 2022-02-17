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
		{
			TestName:      "including replacement CloudWatch",
			Value:         "CloudWatchLogsLogGroupArn",
			ExpectedValue: "cloudwatch_logs_log_group_arn",
		},
		{
			TestName:      "including replacement CNAMEs",
			Value:         "CNAMEs",
			ExpectedValue: "cnames",
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

func TestPluralize(t *testing.T) {
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
			TestName:      "name ending in s",
			Value:         "aws_cloudwatch_event_bus",
			ExpectedValue: "aws_cloudwatch_event_buses",
		},
		{
			TestName:      "name ending in capital s",
			Value:         "locationNFS",
			ExpectedValue: "locationNFS_plural",
		},
		{
			TestName:      "name ending in number",
			Value:         "aws_datasync_location_s3",
			ExpectedValue: "aws_datasync_location_s3s",
		},
		{
			TestName:      "name ending in 'efs'",
			Value:         "aws_example_efs",
			ExpectedValue: "aws_example_efs_plural",
		},
		{
			TestName:      "name ending in 'nfs'",
			Value:         "aws_example_nfs",
			ExpectedValue: "aws_example_nfs_plural",
		},
		{
			TestName:      "name ending in 'xfs'",
			Value:         "aws_example_xfs",
			ExpectedValue: "aws_example_xfs",
		},
		{
			TestName:      "name ending in 'tion'",
			Value:         "aws_datasync_location",
			ExpectedValue: "aws_datasync_locations",
		},
		{
			TestName:      "name ending in 'tions'",
			Value:         "aws_datasync_locations",
			ExpectedValue: "aws_datasync_locations_plural",
		},
		{
			TestName:      "name ending in 'window'",
			Value:         "aws_datasync_window",
			ExpectedValue: "aws_datasync_windows",
		},
		{
			TestName:      "name ending in 'windows'",
			Value:         "aws_datasync_windows",
			ExpectedValue: "aws_datasync_windows_plural",
		},
		{
			TestName:      "singular name",
			Value:         "aws_wafv2_web_acl",
			ExpectedValue: "aws_wafv2_web_acls",
		},
		{
			TestName:      "custom rule for lens",
			Value:         "awscc_example_lens",
			ExpectedValue: "awscc_example_lenses",
		},
		{
			TestName:      "name ending in 'hdfs'",
			Value:         "awscc_datasync_location_hdfs",
			ExpectedValue: "awscc_datasync_location_hdfs_plural",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got := naming.Pluralize(testCase.Value)

			if got != testCase.ExpectedValue {
				t.Errorf("expected: %s, got: %s", testCase.ExpectedValue, got)
			}
		})
	}
}

func TestPluralizeWithCustomNameSuffix(t *testing.T) {
	testCases := []struct {
		TestName      string
		Name          string
		Suffix        string
		ExpectedValue string
	}{
		{
			TestName:      "empty string",
			Name:          "",
			Suffix:        "",
			ExpectedValue: "",
		},
		{
			TestName:      "non custom name with suffix",
			Name:          "aws_example_association",
			Suffix:        "_plural",
			ExpectedValue: "aws_example_associations",
		},
		{
			TestName:      "custom underscored name ending in 'tions'",
			Name:          "aws_example_associations",
			Suffix:        "_plural",
			ExpectedValue: "aws_example_associations_plural",
		},
		{
			TestName:      "custom name ending in 'tions'",
			Name:          "Associations",
			Suffix:        "Plural",
			ExpectedValue: "AssociationsPlural",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got := naming.PluralizeWithCustomNameSuffix(testCase.Name, testCase.Suffix)

			if got != testCase.ExpectedValue {
				t.Errorf("expected: %s, got: %s", testCase.ExpectedValue, got)
			}
		})
	}
}
