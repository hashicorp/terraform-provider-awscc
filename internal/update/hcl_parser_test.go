// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func TestDiffSchemas_WithoutValidation(t *testing.T) {

	tempDir := t.TempDir()

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
			filePaths.AllSchemasHCL = tempAllSchemasPath

			// Call diffSchemas with the temporary file path
			result, err := diffSchemas(tc.newSchemas, tc.lastSchemas, &changes, filePaths)

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
