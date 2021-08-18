package naming_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
)

func TestParseCloudFormationTypeName(t *testing.T) {
	testCases := []struct {
		TestName             string
		Value                string
		ExpectError          bool
		ExpectedOrganization string
		ExpectedService      string
		ExpectedResource     string
	}{
		{
			TestName:    "empty string",
			Value:       "",
			ExpectError: true,
		},
		{
			TestName:    "whitespace string",
			Value:       "  ",
			ExpectError: true,
		},
		{
			TestName:    "incorrect number of segments",
			Value:       "AWS::EC2",
			ExpectError: true,
		},
		{
			TestName:    "invalid type name",
			Value:       "AWS::KMS::WayTooLongAResourceName000000000000000000000000000000000000000012",
			ExpectError: true,
		},
		{
			TestName:    "Terraform type name",
			Value:       "aws_kms_key",
			ExpectError: true,
		},
		{
			TestName:             "valid type name",
			Value:                "AWS::KMS::Key",
			ExpectedOrganization: "AWS",
			ExpectedService:      "KMS",
			ExpectedResource:     "Key",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			gotOrganization, gotService, gotResource, err := naming.ParseCloudFormationTypeName(testCase.Value)

			if err == nil && testCase.ExpectError {
				t.Fatalf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Fatalf("got unexpected error: %s", err)
			}

			if gotOrganization != testCase.ExpectedOrganization {
				t.Errorf("expected Organization: %s, got: %s", testCase.ExpectedOrganization, gotOrganization)
			}
			if gotService != testCase.ExpectedService {
				t.Errorf("expected Service: %s, got: %s", testCase.ExpectedService, gotService)
			}
			if gotResource != testCase.ExpectedResource {
				t.Errorf("expected Resource: %s, got: %s", testCase.ExpectedResource, gotResource)
			}
		})
	}
}

func TestCreateTerraformTypeName(t *testing.T) {
	testCases := []struct {
		TestName      string
		Organization  string
		Service       string
		Resource      string
		ExpectedValue string
	}{
		{
			TestName:      "valid type name",
			Organization:  "aws",
			Service:       "kms",
			Resource:      "key",
			ExpectedValue: "aws_kms_key",
		},
		{
			TestName:      "valid type name multiple underscores in resource",
			Organization:  "aws",
			Service:       "logs",
			Resource:      "log_group",
			ExpectedValue: "aws_logs_log_group",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			gotValue := naming.CreateTerraformTypeName(testCase.Organization, testCase.Service, testCase.Resource)

			if gotValue != testCase.ExpectedValue {
				t.Errorf("expected type name: %s, got: %s", testCase.ExpectedValue, gotValue)
			}
		})
	}
}

func TestParseTerraformTypeName(t *testing.T) {
	testCases := []struct {
		TestName             string
		Value                string
		ExpectError          bool
		ExpectedOrganization string
		ExpectedService      string
		ExpectedResource     string
	}{
		{
			TestName:    "empty string",
			Value:       "",
			ExpectError: true,
		},
		{
			TestName:    "whitespace string",
			Value:       "  ",
			ExpectError: true,
		},
		{
			TestName:    "incorrect number of segments",
			Value:       "aws_ec2",
			ExpectError: true,
		},
		{
			TestName:    "invalid type name",
			Value:       "aws_kms_k-y",
			ExpectError: true,
		},
		{
			TestName:    "CloudFormation type name",
			Value:       "AWS::KMS::Key",
			ExpectError: true,
		},
		{
			TestName:             "valid type name",
			Value:                "aws_kms_key",
			ExpectedOrganization: "aws",
			ExpectedService:      "kms",
			ExpectedResource:     "key",
		},
		{
			TestName:             "valid type name multiple underscores in resource",
			Value:                "aws_logs_log_group",
			ExpectedOrganization: "aws",
			ExpectedService:      "logs",
			ExpectedResource:     "log_group",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			gotOrganization, gotService, gotResource, err := naming.ParseTerraformTypeName(testCase.Value)

			if err == nil && testCase.ExpectError {
				t.Fatalf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Fatalf("got unexpected error: %s", err)
			}

			if gotOrganization != testCase.ExpectedOrganization {
				t.Errorf("expected Organization: %s, got: %s", testCase.ExpectedOrganization, gotOrganization)
			}
			if gotService != testCase.ExpectedService {
				t.Errorf("expected Service: %s, got: %s", testCase.ExpectedService, gotService)
			}
			if gotResource != testCase.ExpectedResource {
				t.Errorf("expected Resource: %s, got: %s", testCase.ExpectedResource, gotResource)
			}
		})
	}
}
