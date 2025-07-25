// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-github/v72/github"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func TestMakeBuild(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		buildType           string
		expectError         bool
		expectedErrorString string
	}{
		"valid_schemas_build_type": {
			buildType:   BuildTypeSchemas,
			expectError: false,
		},
		"valid_resources_build_type": {
			buildType:   BuildTypeResources,
			expectError: false,
		},
		"valid_singular_data_sources_build_type": {
			buildType:   BuildTypeSingularDataSources,
			expectError: false,
		},
		"valid_plural_data_sources_build_type": {
			buildType:   BuildTypePluralDataSources,
			expectError: false,
		},
		"invalid_build_type": {
			buildType:           "invalid-type",
			expectError:         true,
			expectedErrorString: "invalid build type",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			// Create temporary directory for test files
			tempDir := t.TempDir()

			// Mock file paths
			filePaths := &UpdateFilePaths{
				RunMakesErrors:           filepath.Join(tempDir, "makes_errors.txt"),
				SuppressionCheckout:      filepath.Join(tempDir, "suppression_checkout.txt"),
				LastResource:             filepath.Join(tempDir, "last_resource.txt"),
				CloudFormationSchemasDir: filepath.Join(tempDir, "schemas"),
			}

			// Create schemas directory
			if err := os.MkdirAll(filePaths.CloudFormationSchemasDir, 0755); err != nil {
				t.Fatalf("failed to create schemas directory: %v", err)
			}

			// Mock GitHub config
			config := &GitHubConfig{
				Client:      &github.Client{},
				Repository:  "test-repo",
				RepoOwner:   "test-owner",
				RepoName:    "test-name",
				CurrentDate: "2023-01-01",
			}

			// Mock schemas
			currentSchemas := &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{},
			}

			changes := []string{}
			isNewMap := map[string]bool{}

			err := makeBuild(ctx, config, currentSchemas, test.buildType, &changes, filePaths, isNewMap)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if !strings.Contains(err.Error(), test.expectedErrorString) {
					t.Fatalf("expected error to contain %q, but got: %v", test.expectedErrorString, err)
				}
			} else {
				if err != nil && !strings.Contains(err.Error(), "failed to execute make") {
					// We expect make command to fail in test environment, but not validation errors
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestProcessErrorLine(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		errorLine           string
		buildType           string
		expectError         bool
		expectedErrorString string
	}{
		"empty_line": {
			errorLine:   "",
			buildType:   BuildTypeSchemas,
			expectError: false,
		},
		"stack_overflow_error": {
			errorLine:   "stack overflow detected in resource processing",
			buildType:   BuildTypeSchemas,
			expectError: false, // Will succeed when last_resource.txt is available
		},
		"aws_underscore_error": {
			errorLine:   "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:",
			buildType:   BuildTypeSchemas,
			expectError: false, // Will succeed with suppression
		},
		"aws_colon_error": {
			errorLine:   "The type 'AWS::DataSync::StorageSystem' cannot be found",
			buildType:   BuildTypeSchemas,
			expectError: false, // Will succeed with suppression
		},
		"aws_underscore_lowercase_error": {
			errorLine:   "error loading CloudFormation Resource Provider Schema for aws_nimblestudio_studio:",
			buildType:   BuildTypeSchemas,
			expectError: true, // Will fail due to missing schemas in directory
		},
		"awscc_error": {
			errorLine:   "error loading CloudFormation Resource Provider Schema for awscc_s3_bucket:",
			buildType:   BuildTypeSchemas,
			expectError: true, // Will fail due to parsing error
		},
		"statuscode_403_error": {
			errorLine:           "StatusCode: 403, authentication failed",
			buildType:           BuildTypeSchemas,
			expectError:         true,
			expectedErrorString: "authentication failed: no valid AWS credentials",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			// Create temporary directory for test files
			tempDir := t.TempDir()

			// Mock file paths
			filePaths := &UpdateFilePaths{
				RunMakesErrors:           filepath.Join(tempDir, "makes_errors.txt"),
				SuppressionCheckout:      filepath.Join(tempDir, "suppression_checkout.txt"),
				LastResource:             filepath.Join(tempDir, "last_resource.txt"),
				CloudFormationSchemasDir: filepath.Join(tempDir, "schemas"),
				AllSchemasHCL:            filepath.Join(tempDir, "all_schemas.hcl"),
			}

			// Create schemas directory
			if err := os.MkdirAll(filePaths.CloudFormationSchemasDir, 0755); err != nil {
				t.Fatalf("failed to create schemas directory: %v", err)
			}

			// Create last_resource.txt for stack overflow test
			if strings.Contains(test.errorLine, "stack overflow") {
				if err := os.WriteFile(filePaths.LastResource, []byte("AWS_Test_Resource"), 0644); err != nil {
					t.Fatalf("failed to create last_resource.txt: %v", err)
				}
			}

			// Mock GitHub config
			config := &GitHubConfig{
				Client:      &github.Client{},
				Repository:  "test-repo",
				RepoOwner:   "test-owner",
				RepoName:    "test-name",
				CurrentDate: "2023-01-01",
			}

			// Mock schemas
			currentSchemas := &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{},
			}

			changes := []string{}
			isNewMap := map[string]bool{}

			err := processErrorLine(ctx, test.errorLine, config, currentSchemas, test.buildType, &changes, filePaths, isNewMap)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if test.expectedErrorString != "" && !strings.Contains(err.Error(), test.expectedErrorString) {
					t.Fatalf("expected error to contain %q, but got: %v", test.expectedErrorString, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestHandleStackOverflowError(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		setupLastResource   bool
		lastResourceContent string
		expectError         bool
		expectedErrorString string
	}{
		"valid_resource_name": {
			setupLastResource:   true,
			lastResourceContent: "AWS_Test_Resource\n",
			expectError:         false, // Will succeed with proper resource name
		},
		"empty_resource_name": {
			setupLastResource:   true,
			lastResourceContent: "\n",
			expectError:         true,
			expectedErrorString: "resource name not found for stack overflow",
		},
		"missing_last_resource_file": {
			setupLastResource:   false,
			expectError:         true,
			expectedErrorString: "failed to read",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			// Create temporary directory for test files
			tempDir := t.TempDir()

			// Mock file paths
			filePaths := &UpdateFilePaths{
				LastResource:             filepath.Join(tempDir, "last_resource.txt"),
				CloudFormationSchemasDir: filepath.Join(tempDir, "schemas"),
				AllSchemasHCL:            filepath.Join(tempDir, "all_schemas.hcl"),
				SuppressionCheckout:      filepath.Join(tempDir, "suppression_checkout.txt"),
			}

			// Create schemas directory
			if err := os.MkdirAll(filePaths.CloudFormationSchemasDir, 0755); err != nil {
				t.Fatalf("failed to create schemas directory: %v", err)
			}

			// Setup last_resource.txt if needed
			if test.setupLastResource {
				if err := os.WriteFile(filePaths.LastResource, []byte(test.lastResourceContent), 0644); err != nil {
					t.Fatalf("failed to create last_resource.txt: %v", err)
				}
			}

			// Mock GitHub config
			config := &GitHubConfig{
				Client:      &github.Client{},
				Repository:  "test-repo",
				RepoOwner:   "test-owner",
				RepoName:    "test-name",
				CurrentDate: "2023-01-01",
			}

			// Mock schemas
			currentSchemas := &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{},
			}

			isNewMap := map[string]bool{}

			err := handleStackOverflowError(ctx, "stack overflow error", config, currentSchemas, BuildTypeSchemas, filePaths, isNewMap)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if test.expectedErrorString != "" && !strings.Contains(err.Error(), test.expectedErrorString) {
					t.Fatalf("expected error to contain %q, but got: %v", test.expectedErrorString, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestHandleAWS_Error(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		errorLine           string
		expectError         bool
		expectedErrorString string
	}{
		"valid_aws_error": {
			errorLine:   "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:",
			expectError: false, // Will succeed with proper suppression
		},
		"invalid_aws_error_format": {
			errorLine:           "some error without AWS_ pattern",
			expectError:         true,
			expectedErrorString: "failed to extract resource name",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			// Create temporary directory for test files
			tempDir := t.TempDir()

			// Mock file paths
			filePaths := &UpdateFilePaths{
				CloudFormationSchemasDir: filepath.Join(tempDir, "schemas"),
				AllSchemasHCL:            filepath.Join(tempDir, "all_schemas.hcl"),
				SuppressionCheckout:      filepath.Join(tempDir, "suppression_checkout.txt"),
			}

			// Create schemas directory
			if err := os.MkdirAll(filePaths.CloudFormationSchemasDir, 0755); err != nil {
				t.Fatalf("failed to create schemas directory: %v", err)
			}

			// Mock GitHub config
			config := &GitHubConfig{
				Client:      &github.Client{},
				Repository:  "test-repo",
				RepoOwner:   "test-owner",
				RepoName:    "test-name",
				CurrentDate: "2023-01-01",
			}

			// Mock schemas
			currentSchemas := &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{},
			}

			changes := []string{}
			isNewMap := map[string]bool{}

			err := handleAWS_Error(ctx, test.errorLine, config, currentSchemas, BuildTypeSchemas, &changes, filePaths, isNewMap)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if test.expectedErrorString != "" && !strings.Contains(err.Error(), test.expectedErrorString) {
					t.Fatalf("expected error to contain %q, but got: %v", test.expectedErrorString, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestHandleAWSColonError(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		errorLine           string
		expectError         bool
		expectedErrorString string
	}{
		"valid_aws_colon_error": {
			errorLine:   "The type 'AWS::DataSync::StorageSystem' cannot be found",
			expectError: false, // Will succeed with proper suppression
		},
		"invalid_aws_colon_error_format": {
			errorLine:   "some error without AWS:: pattern",
			expectError: false, // Will succeed as it extracts "AWS" and suppresses it
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			// Create temporary directory for test files
			tempDir := t.TempDir()

			// Mock file paths
			filePaths := &UpdateFilePaths{
				CloudFormationSchemasDir: filepath.Join(tempDir, "schemas"),
				AllSchemasHCL:            filepath.Join(tempDir, "all_schemas.hcl"),
				SuppressionCheckout:      filepath.Join(tempDir, "suppression_checkout.txt"),
			}

			// Create schemas directory
			if err := os.MkdirAll(filePaths.CloudFormationSchemasDir, 0755); err != nil {
				t.Fatalf("failed to create schemas directory: %v", err)
			}

			// Mock GitHub config
			config := &GitHubConfig{
				Client:      &github.Client{},
				Repository:  "test-repo",
				RepoOwner:   "test-owner",
				RepoName:    "test-name",
				CurrentDate: "2023-01-01",
			}

			// Mock schemas
			currentSchemas := &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{},
			}

			isNewMap := map[string]bool{}

			err := handleAWSColonError(ctx, test.errorLine, config, currentSchemas, BuildTypeSchemas, filePaths, isNewMap)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if test.expectedErrorString != "" && !strings.Contains(err.Error(), test.expectedErrorString) {
					t.Fatalf("expected error to contain %q, but got: %v", test.expectedErrorString, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestHandleStatusCode403Error(t *testing.T) {
	t.Parallel()

	err := handleStatusCode403Error()

	if err == nil {
		t.Fatal("expected error but got none")
	}

	expectedErrorString := "authentication failed: no valid AWS credentials"
	if !strings.Contains(err.Error(), expectedErrorString) {
		t.Fatalf("expected error to contain %q, but got: %v", expectedErrorString, err)
	}
}

func TestNormalizeNames(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		cfTypeName     string
		tfTypeName     string
		expectedCfType string
		expectedTfType string
	}{
		"basic_normalization": {
			cfTypeName:     "AWS::S3::Bucket",
			tfTypeName:     "aws_s3_bucket",
			expectedCfType: "awss3bucket",
			expectedTfType: "awss3bucket",
		},
		"complex_normalization": {
			cfTypeName:     "AWS::EC2::VPC::Endpoint",
			tfTypeName:     "aws_ec2_vpc_endpoint",
			expectedCfType: "awsec2vpcendpoint",
			expectedTfType: "awsec2vpcendpoint",
		},
		"mixed_separators": {
			cfTypeName:     "AWS::DynamoDB::Table",
			tfTypeName:     "aws_dynamodb_table",
			expectedCfType: "awsdynamodbtable",
			expectedTfType: "awsdynamodbtable",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualCfType, actualTfType := normalizeNames(test.cfTypeName, test.tfTypeName)

			if actualCfType != test.expectedCfType {
				t.Fatalf("expected normalized cfTypeName %q, but got: %q", test.expectedCfType, actualCfType)
			}

			if actualTfType != test.expectedTfType {
				t.Fatalf("expected normalized tfTypeName %q, but got: %q", test.expectedTfType, actualTfType)
			}
		})
	}
}

func TestAddSchemaToCheckout(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		resource        string
		existingContent string
		expectError     bool
	}{
		"add_to_empty_file": {
			resource:        "AWS_S3_Bucket",
			existingContent: "",
			expectError:     false,
		},
		"add_to_existing_file": {
			resource:        "AWS_EC2_Instance",
			existingContent: "\nsome/existing/path.json",
			expectError:     false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create temporary directory for test files
			tempDir := t.TempDir()

			// Mock file paths
			filePaths := &UpdateFilePaths{
				SuppressionCheckout:      filepath.Join(tempDir, "suppression_checkout.txt"),
				CloudFormationSchemasDir: filepath.Join(tempDir, "schemas"),
			}

			// Create schemas directory
			if err := os.MkdirAll(filePaths.CloudFormationSchemasDir, 0755); err != nil {
				t.Fatalf("failed to create schemas directory: %v", err)
			}

			// Setup existing content if any
			if test.existingContent != "" {
				if err := os.WriteFile(filePaths.SuppressionCheckout, []byte(test.existingContent), 0644); err != nil {
					t.Fatalf("failed to create suppression checkout file: %v", err)
				}
			}

			err := addSchemaToCheckout(test.resource, filePaths)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				// Verify the content was added
				content, readErr := os.ReadFile(filePaths.SuppressionCheckout)
				if readErr != nil {
					t.Fatalf("failed to read suppression checkout file: %v", readErr)
				}

				expectedContent := fmt.Sprintf("\n%s/%s.json", filePaths.CloudFormationSchemasDir, test.resource)
				if !strings.Contains(string(content), expectedContent) {
					t.Fatalf("expected file to contain %q, but got: %q", expectedContent, string(content))
				}
			}
		})
	}
}

func TestCfTypeNameToTerraformTypeName(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		cfTypeName          string
		expectedTfTypeName  string
		expectError         bool
		expectedErrorString string
	}{
		"valid_s3_bucket": {
			cfTypeName:         "AWS::S3::Bucket",
			expectedTfTypeName: "aws_s3_bucket",
			expectError:        false,
		},
		"valid_ec2_instance": {
			cfTypeName:         "AWS::EC2::Instance",
			expectedTfTypeName: "aws_ec2_instance",
			expectError:        false,
		},
		"with_underscores": {
			cfTypeName:         "AWS_DynamoDB_Table",
			expectedTfTypeName: "aws_dynamodb_table",
			expectError:        false,
		},
		"invalid_format": {
			cfTypeName:          "InvalidFormat",
			expectError:         true,
			expectedErrorString: "parsing CloudFormation type name",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tfTypeName, err := cfTypeNameToTerraformTypeName(test.cfTypeName)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if test.expectedErrorString != "" && !strings.Contains(err.Error(), test.expectedErrorString) {
					t.Fatalf("expected error to contain %q, but got: %v", test.expectedErrorString, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if tfTypeName != test.expectedTfTypeName {
					t.Fatalf("expected tfTypeName %q, but got: %q", test.expectedTfTypeName, tfTypeName)
				}
			}
		})
	}
}

func TestIsNew(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		cloudFormationTypeName string
		isNewMap               map[string]bool
		expectedResult         bool
	}{
		"existing_resource": {
			cloudFormationTypeName: "AWS::S3::Bucket",
			isNewMap:               map[string]bool{"aws_s3_bucket": true},
			expectedResult:         false,
		},
		"new_resource": {
			cloudFormationTypeName: "AWS::EC2::Instance",
			isNewMap:               map[string]bool{"aws_s3_bucket": true},
			expectedResult:         true,
		},
		"empty_map": {
			cloudFormationTypeName: "AWS::S3::Bucket",
			isNewMap:               map[string]bool{},
			expectedResult:         true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := isNew(test.cloudFormationTypeName, test.isNewMap)

			if result != test.expectedResult {
				t.Fatalf("expected result %v, but got: %v", test.expectedResult, result)
			}
		})
	}
}

func TestGetPluralResourceNames(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input               string
		isNewMap            map[string]bool
		expectedMatch       string
		expectedOriginal    string
		expectError         bool
		expectedErrorString string
	}{
		"valid_plural_match": {
			input:            "awscc_dynamodb_tables",
			isNewMap:         map[string]bool{"aws_dynamodb": true},
			expectedMatch:    "aws_dynamodb",
			expectedOriginal: "aws_dynamodb_tables",
			expectError:      false,
		},
		"trimmed_match": {
			input:            "awscc_s3_storage_lenses",
			isNewMap:         map[string]bool{"aws_s3_storage": true},
			expectedMatch:    "aws_s3_storage",
			expectedOriginal: "aws_s3_storage_lenses",
			expectError:      false,
		},
		"invalid_prefix": {
			input:               "invalid_prefix_tables",
			isNewMap:            map[string]bool{},
			expectError:         true,
			expectedErrorString: "input does not start with awscc_",
		},
		"no_match_found": {
			input:               "awscc_nonexistent_resource",
			isNewMap:            map[string]bool{"aws_other_resource": true},
			expectError:         true,
			expectedErrorString: "no match found",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			match, original, err := getPluralResourceNames(test.input, test.isNewMap)

			if test.expectError {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if test.expectedErrorString != "" && !strings.Contains(err.Error(), test.expectedErrorString) {
					t.Fatalf("expected error to contain %q, but got: %v", test.expectedErrorString, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if match != test.expectedMatch {
					t.Fatalf("expected match %q, but got: %q", test.expectedMatch, match)
				}
				if original != test.expectedOriginal {
					t.Fatalf("expected original %q, but got: %q", test.expectedOriginal, original)
				}
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    byte
		expected bool
	}{
		"lowercase_a": {
			input:    'a',
			expected: true,
		},
		"uppercase_z": {
			input:    'Z',
			expected: true,
		},
		"digit": {
			input:    '5',
			expected: false,
		},
		"underscore": {
			input:    '_',
			expected: false,
		},
		"colon": {
			input:    ':',
			expected: false,
		},
		"space": {
			input:    ' ',
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := isLetter(test.input)

			if result != test.expected {
				t.Fatalf("expected %v for input %q, but got: %v", test.expected, test.input, result)
			}
		})
	}
}
