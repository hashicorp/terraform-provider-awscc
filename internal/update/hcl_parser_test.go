// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func TestParseSchemaToStruct(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	testCases := []struct {
		name          string
		fileContent   string
		expectedError bool
		validateFunc  func(*testing.T, *allschemas.AvailableSchemas)
	}{
		{
			name: "valid available schemas file",
			fileContent: `# Test available schemas
resource_schema "aws_test_service" {
  cloudformation_type_name = "AWS::Test::Service"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_another_service" {
  cloudformation_type_name = "AWS::Another::Service"
}`,
			expectedError: false,
			validateFunc: func(t *testing.T, schemas *allschemas.AvailableSchemas) {
				if len(schemas.Resources) != 2 {
					t.Errorf("Expected 2 resources, got %d", len(schemas.Resources))
				}
				if schemas.Resources[0].ResourceTypeName != "aws_test_service" {
					t.Errorf("Expected first resource name 'aws_test_service', got '%s'", schemas.Resources[0].ResourceTypeName)
				}
				if schemas.Resources[0].CloudFormationTypeName != "AWS::Test::Service" {
					t.Errorf("Expected first CF type 'AWS::Test::Service', got '%s'", schemas.Resources[0].CloudFormationTypeName)
				}
				if !schemas.Resources[0].SuppressPluralDataSourceGeneration {
					t.Error("Expected first resource to suppress plural data source generation")
				}
				if schemas.Resources[1].SuppressPluralDataSourceGeneration {
					t.Error("Expected second resource to not suppress plural data source generation")
				}
			},
		},
		{
			name: "valid minimal schemas file",
			fileContent: `resource_schema "aws_test_resource" {
  cloudformation_type_name = "AWS::Test::Resource"
}`,
			expectedError: false,
			validateFunc: func(t *testing.T, schemas *allschemas.AvailableSchemas) {
				if len(schemas.Resources) != 1 {
					t.Errorf("Expected 1 resource, got %d", len(schemas.Resources))
				}
				if schemas.Resources[0].ResourceTypeName != "aws_test_resource" {
					t.Errorf("Expected resource name 'aws_test_resource', got '%s'", schemas.Resources[0].ResourceTypeName)
				}
			},
		},
		{
			name:          "invalid HCL syntax",
			fileContent:   `resource_schema "invalid" { missing_closing_brace = true`,
			expectedError: true,
			validateFunc:  nil,
		},
		{
			name:          "empty file",
			fileContent:   "",
			expectedError: false,
			validateFunc: func(t *testing.T, schemas *allschemas.AvailableSchemas) {
				if len(schemas.Resources) != 0 {
					t.Errorf("Expected 0 resources for empty file, got %d", len(schemas.Resources))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(tempDir, tc.name+".hcl")
			err := os.WriteFile(testFile, []byte(tc.fileContent), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			// Test parseSchemaToStruct
			result, err := parseSchemaToStruct(testFile, allschemas.AvailableSchemas{})

			// Check error expectation
			if tc.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			// Run validation if provided and no error expected
			if !tc.expectedError && tc.validateFunc != nil {
				tc.validateFunc(t, result)
			}
		})
	}
}

func TestParseSchemaToStruct_AllSchemas(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	testCases := []struct {
		name          string
		fileContent   string
		expectedError bool
		validateFunc  func(*testing.T, *allschemas.AllSchemas)
	}{
		{
			name: "valid all schemas file",
			fileContent: `defaults {
  schema_cache_directory = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

resource_schema "aws_test_resource" {
  cloudformation_type_name = "AWS::Test::Resource"
  suppress_resource_generation = true
  suppression_reason = "Test suppression"
}`,
			expectedError: false,
			validateFunc: func(t *testing.T, schemas *allschemas.AllSchemas) {
				if schemas.Defaults.SchemaCacheDirectory != "../service/cloudformation/schemas" {
					t.Errorf("Expected schema cache directory '../service/cloudformation/schemas', got '%s'", schemas.Defaults.SchemaCacheDirectory)
				}
				if len(schemas.Resources) != 1 {
					t.Errorf("Expected 1 resource, got %d", len(schemas.Resources))
				}
				if schemas.Resources[0].ResourceTypeName != "aws_test_resource" {
					t.Errorf("Expected resource name 'aws_test_resource', got '%s'", schemas.Resources[0].ResourceTypeName)
				}
				if !schemas.Resources[0].SuppressResourceGeneration {
					t.Error("Expected resource to have suppress_resource_generation = true")
				}
				if schemas.Resources[0].SuppressionReason != "Test suppression" {
					t.Errorf("Expected suppression reason 'Test suppression', got '%s'", schemas.Resources[0].SuppressionReason)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(tempDir, tc.name+".hcl")
			err := os.WriteFile(testFile, []byte(tc.fileContent), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			// Test parseSchemaToStruct with AllSchemas type
			result, err := parseSchemaToStruct(testFile, allschemas.AllSchemas{})

			// Check error expectation
			if tc.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			// Run validation if provided and no error expected
			if !tc.expectedError && tc.validateFunc != nil {
				tc.validateFunc(t, result)
			}
		})
	}
}

func TestParseSchemaToStruct_FileNotFound(t *testing.T) {
	t.Parallel()

	nonExistentFile := "/path/that/does/not/exist.hcl"
	_, err := parseSchemaToStruct(nonExistentFile, allschemas.AvailableSchemas{})

	if err == nil {
		t.Error("Expected error for non-existent file but got none")
	}

	if !strings.Contains(err.Error(), "failed to read file") {
		t.Errorf("Expected error message to contain 'failed to read file', got: %v", err)
	}
}

func TestDiffSchemas_WithoutValidation(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	ctx := context.Background()

	// Create test file paths
	allSchemasPath := filepath.Join(tempDir, "all_schemas.hcl")
	filePaths := &UpdateFilePaths{
		AllSchemasHCL:            allSchemasPath,
		LastResource:             filepath.Join(tempDir, "internal/provider/last_resource.txt"),
		CloudFormationSchemasDir: filepath.Join(tempDir, "internal/service/cloudformation/schemas"),
	}

	// Create initial all_schemas.hcl file
	initialAllSchemas := `defaults {
  schema_cache_directory = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

resource_schema "aws_existing_resource" {
  cloudformation_type_name = "AWS::Existing::Resource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_unchanged_resource" {
  cloudformation_type_name = "AWS::Unchanged::Resource"
}`

	err := os.WriteFile(allSchemasPath, []byte(initialAllSchemas), 0644)
	if err != nil {
		t.Fatalf("Failed to create initial all_schemas.hcl: %v", err)
	}

	testCases := []struct {
		name            string
		newSchemas      *allschemas.AvailableSchemas
		lastSchemas     *allschemas.AvailableSchemas
		expectedChanges []string
		expectedError   bool
	}{
		{
			name: "new resource added",
			newSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						ResourceTypeName:       "aws_new_resource",
						CloudFormationTypeName: "AWS::New::Resource",
					},
					{
						ResourceTypeName:       "aws_unchanged_resource",
						CloudFormationTypeName: "AWS::Unchanged::Resource",
					},
				},
			},
			lastSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						ResourceTypeName:       "aws_unchanged_resource",
						CloudFormationTypeName: "AWS::Unchanged::Resource",
					},
				},
			},
			expectedChanges: []string{"AWS::New::Resource - new"},
			expectedError:   false,
		},
		{
			name: "no changes",
			newSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						ResourceTypeName:       "aws_unchanged_resource",
						CloudFormationTypeName: "AWS::Unchanged::Resource",
					},
				},
			},
			lastSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						ResourceTypeName:       "aws_unchanged_resource",
						CloudFormationTypeName: "AWS::Unchanged::Resource",
					},
				},
			},
			expectedChanges: []string{},
			expectedError:   false,
		},
		{
			name: "resource changed",
			newSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						ResourceTypeName:       "aws_unchanged_resource",
						CloudFormationTypeName: "AWS::Unchanged::Resource",
					},
					{
						ResourceTypeName:                   "aws_existing_resource",
						CloudFormationTypeName:             "AWS::Existing::Resource",
						SuppressPluralDataSourceGeneration: false, // changed from true to false
					},
				},
			},
			lastSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						ResourceTypeName:       "aws_unchanged_resource",
						CloudFormationTypeName: "AWS::Unchanged::Resource",
					},
					{
						ResourceTypeName:                   "aws_existing_resource",
						CloudFormationTypeName:             "AWS::Existing::Resource",
						SuppressPluralDataSourceGeneration: true, // original value
					},
				},
			},
			expectedChanges: []string{"AWS::Existing::Resource - changed"},
			expectedError:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset changes slice
			changes := make([]string, 0)

			v := &allschemas.AllSchemas{}

			v.Defaults = allschemas.Defaults{
				SchemaCacheDirectory:    "../service/cloudformation/schemas",
				TerraformTypeNamePrefix: "awscc",
			}
			v.Meta = allschemas.MetaSchema{
				Path: "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json",
			}
			v.Resources = []allschemas.ResourceAllSchema{
				{
					ResourceTypeName:           "aws_test_resource",
					CloudFormationTypeName:     "AWS::Test::Resource",
					SuppressResourceGeneration: true,
				},
			}

			// Create a temporary allschemas.hcl file with the v object
			tempAllSchemasPath := filepath.Join(tempDir, "temp_all_schemas.hcl")
			err := writeSchemasToHCLFile(v, tempAllSchemasPath)
			if err != nil {
				t.Fatalf("Failed to write temporary all_schemas.hcl: %v", err)
			}

			// Call diffSchemas with the temporary file path
			result, err := diffSchemas(ctx, tc.newSchemas, tc.lastSchemas, tempAllSchemasPath, &changes, filePaths)

			// For no changes case, result should be nil
			if len(tc.expectedChanges) == 0 {
				if result != nil {
					t.Error("Expected nil result for no changes case")
				}
				return
			}

			// Check that we don't get parsing errors (AWS validation errors are expected in tests)
			if err != nil && strings.Contains(err.Error(), "failed to parse") {
				t.Errorf("Unexpected parsing error: %v", err)
			}

			// Validate that new resources are detected in changes (before AWS validation)
			for _, expectedChange := range tc.expectedChanges {
				found := false
				for _, change := range changes {
					if strings.Contains(change, expectedChange) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected change '%s' not found in changes: %v", expectedChange, changes)
				}
			}
		})
	}
}

func TestWriteSchemasToHCLFile(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	testCases := []struct {
		name          string
		schema        interface{}
		expectedError bool
		validateFunc  func(*testing.T, string)
	}{
		{
			name: "write available schemas",
			schema: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						ResourceTypeName:                   "aws_test_resource",
						CloudFormationTypeName:             "AWS::Test::Resource",
						SuppressPluralDataSourceGeneration: true,
					},
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, content string) {
				if !strings.Contains(content, "aws_test_resource") {
					t.Error("Expected content to contain 'aws_test_resource'")
				}
				if !strings.Contains(content, "AWS::Test::Resource") {
					t.Error("Expected content to contain 'AWS::Test::Resource'")
				}
				if !strings.Contains(content, "suppress_plural_data_source_generation") {
					t.Error("Expected content to contain 'suppress_plural_data_source_generation'")
				}
			},
		},
		{
			name: "write all schemas",
			schema: &allschemas.AllSchemas{
				Defaults: allschemas.Defaults{
					SchemaCacheDirectory:    "../service/cloudformation/schemas",
					TerraformTypeNamePrefix: "awscc",
				},
				Meta: allschemas.MetaSchema{
					Path: "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json",
				},
				Resources: []allschemas.ResourceAllSchema{
					{
						ResourceTypeName:           "aws_test_resource",
						CloudFormationTypeName:     "AWS::Test::Resource",
						SuppressResourceGeneration: true,
						SuppressionReason:          "Test suppression",
					},
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, content string) {
				if !strings.Contains(content, "defaults") {
					t.Error("Expected content to contain 'defaults' block")
				}
				if !strings.Contains(content, "meta_schema") {
					t.Error("Expected content to contain 'meta_schema' block")
				}
				if !strings.Contains(content, "aws_test_resource") {
					t.Error("Expected content to contain 'aws_test_resource'")
				}
				if !strings.Contains(content, "suppress_resource_generation") {
					t.Error("Expected content to contain 'suppress_resource_generation'")
				}
				if !strings.Contains(content, "Test suppression") {
					t.Error("Expected content to contain 'Test suppression'")
				}
			},
		},
		{
			name:          "empty schema",
			schema:        &allschemas.AvailableSchemas{},
			expectedError: false,
			validateFunc: func(t *testing.T, content string) {
				// Empty schema might produce an empty file or minimal structure
				// The important thing is that it doesn't error out
				t.Logf("Empty schema produced content of length: %d", len(content))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testFile := filepath.Join(tempDir, tc.name+".hcl")

			// Call writeSchemasToHCLFile
			err := writeSchemasToHCLFile(tc.schema, testFile)

			// Check error expectation
			if tc.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tc.expectedError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			// Validate file content if no error expected
			if !tc.expectedError {
				content, readErr := os.ReadFile(testFile)
				if readErr != nil {
					t.Fatalf("Failed to read written file: %v", readErr)
				}

				if tc.validateFunc != nil {
					tc.validateFunc(t, string(content))
				}

				// Verify file was created
				if _, statErr := os.Stat(testFile); os.IsNotExist(statErr) {
					t.Error("Expected file to be created")
				}
			}
		})
	}
}

func TestWriteSchemasToHCLFile_InvalidPath(t *testing.T) {
	t.Parallel()

	// Try to write to a directory that doesn't exist and can't be created
	invalidPath := "/root/nonexistent/directory/file.hcl"
	schema := &allschemas.AvailableSchemas{}

	err := writeSchemasToHCLFile(schema, invalidPath)

	if err == nil {
		t.Error("Expected error for invalid path but got none")
	}

	if !strings.Contains(err.Error(), "failed to create file") {
		t.Errorf("Expected error message to contain 'failed to create file', got: %v", err)
	}
}

// Test helper function to create mock schemas
func createMockAvailableSchemas(resources ...allschemas.ResourceSchema) *allschemas.AvailableSchemas {
	return &allschemas.AvailableSchemas{
		Resources: resources,
	}
}

func createMockAllSchemas(resources ...allschemas.ResourceAllSchema) *allschemas.AllSchemas {
	return &allschemas.AllSchemas{
		Defaults: allschemas.Defaults{
			SchemaCacheDirectory:    "../service/cloudformation/schemas",
			TerraformTypeNamePrefix: "awscc",
		},
		Meta: allschemas.MetaSchema{
			Path: "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json",
		},
		Resources: resources,
	}
}

func TestDiffSchemas_EdgeCases(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	ctx := context.Background()

	testCases := []struct {
		name            string
		setupAllSchemas string
		newSchemas      *allschemas.AvailableSchemas
		lastSchemas     *allschemas.AvailableSchemas
		expectedError   bool
		description     string
	}{
		{
			name:            "missing all_schemas.hcl file",
			setupAllSchemas: "", // Don't create the file
			newSchemas: createMockAvailableSchemas(
				allschemas.ResourceSchema{
					ResourceTypeName:       "aws_test_resource",
					CloudFormationTypeName: "AWS::Test::Resource",
				},
			),
			lastSchemas:   createMockAvailableSchemas(),
			expectedError: true,
			description:   "Should fail when all_schemas.hcl doesn't exist",
		},
		{
			name:            "invalid all_schemas.hcl file",
			setupAllSchemas: `invalid hcl syntax {`,
			newSchemas: createMockAvailableSchemas(
				allschemas.ResourceSchema{
					ResourceTypeName:       "aws_test_resource",
					CloudFormationTypeName: "AWS::Test::Resource",
				},
			),
			lastSchemas:   createMockAvailableSchemas(),
			expectedError: true,
			description:   "Should fail when all_schemas.hcl has invalid syntax",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test directory for this case
			testDir := filepath.Join(tempDir, tc.name)
			err := os.MkdirAll(testDir, 0755)
			if err != nil {
				t.Fatalf("Failed to create test directory: %v", err)
			}

			allSchemasPath := filepath.Join(testDir, "all_schemas.hcl")
			filePaths := &UpdateFilePaths{
				AllSchemasHCL:            allSchemasPath,
				LastResource:             filepath.Join(testDir, "internal/provider/last_resource.txt"),
				CloudFormationSchemasDir: filepath.Join(testDir, "internal/service/cloudformation/schemas"),
			}

			// Setup all_schemas.hcl file if content provided
			if tc.setupAllSchemas != "" {
				err := os.WriteFile(allSchemasPath, []byte(tc.setupAllSchemas), 0644)
				if err != nil {
					t.Fatalf("Failed to create all_schemas.hcl: %v", err)
				}
			}

			changes := make([]string, 0)

			// Call diffSchemas
			_, err = diffSchemas(ctx, tc.newSchemas, tc.lastSchemas, allSchemasPath, &changes, filePaths)

			// Check error expectation
			if tc.expectedError && err == nil {
				t.Errorf("Expected error but got none for case: %s", tc.description)
			} else if !tc.expectedError && err != nil {
				t.Errorf("Expected no error but got: %v for case: %s", err, tc.description)
			}
		})
	}
}
