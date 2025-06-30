// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func TestDiffSchemas(t *testing.T) {
	tests := map[string]struct {
		newSchemas         *allschemas.AvailableSchemas
		lastSchemas        *allschemas.AvailableSchemas
		existingAllSchemas *allschemas.AllSchemas
		expectedChanges    map[string]bool
		expectChanges      bool
		expectedResources  int
	}{
		"no_changes": {
			newSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			lastSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			existingAllSchemas: &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{
					{
						CloudFormationTypeName:               "AWS::S3::Bucket",
						ResourceTypeName:                     "aws_s3_bucket",
						SuppressPluralDataSourceGeneration:   false,
						SuppressSingularDataSourceGeneration: false,
						SuppressResourceGeneration:           false,
					},
				},
			},
			expectedChanges:   map[string]bool{},
			expectChanges:     false,
			expectedResources: 1,
		},
		"new_resource": {
			newSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
					{
						CloudFormationTypeName:             "AWS::EC2::Instance",
						ResourceTypeName:                   "aws_ec2_instance",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			lastSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			existingAllSchemas: &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			expectedChanges:   map[string]bool{"New Resource: AWS::EC2::Instance": true},
			expectChanges:     true,
			expectedResources: 2,
		},
		"changed_resource": {
			newSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: true, // Changed from false
					},
				},
			},
			lastSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			existingAllSchemas: &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			expectedChanges: map[string]bool{
				"Changed Resource: AWS::S3::Bucket": true,
				"AWS::S3::Bucket - update":          true,
			},
			expectChanges:     true,
			expectedResources: 1,
		},
		"mixed_changes": {
			newSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: true, // Changed
					},
					{
						CloudFormationTypeName:             "AWS::EC2::Instance",
						ResourceTypeName:                   "aws_ec2_instance",
						SuppressPluralDataSourceGeneration: false, // New
					},
					{
						CloudFormationTypeName:             "AWS::RDS::DBInstance",
						ResourceTypeName:                   "aws_rds_db_instance",
						SuppressPluralDataSourceGeneration: false, // Unchanged
					},
				},
			},
			lastSchemas: &allschemas.AvailableSchemas{
				Resources: []allschemas.ResourceSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
					{
						CloudFormationTypeName:             "AWS::RDS::DBInstance",
						ResourceTypeName:                   "aws_rds_db_instance",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			existingAllSchemas: &allschemas.AllSchemas{
				Resources: []allschemas.ResourceAllSchema{
					{
						CloudFormationTypeName:             "AWS::S3::Bucket",
						ResourceTypeName:                   "aws_s3_bucket",
						SuppressPluralDataSourceGeneration: false,
					},
					{
						CloudFormationTypeName:             "AWS::RDS::DBInstance",
						ResourceTypeName:                   "aws_rds_db_instance",
						SuppressPluralDataSourceGeneration: false,
					},
				},
			},
			expectedChanges: map[string]bool{
				"Changed Resource: AWS::S3::Bucket": true,
				"New Resource: AWS::EC2::Instance":  true,
				"AWS::S3::Bucket - update":          true,
			},
			expectChanges:     true,
			expectedResources: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Create temporary directory for test files
			tempDir := t.TempDir()

			// Mock file paths
			filePaths := &UpdateFilePaths{
				AllSchemasHCL: filepath.Join(tempDir, "all_schemas.hcl"),
			}

			// Write existing AllSchemas to HCL file for the function to read
			if err := writeSchemasToHCLFile(test.existingAllSchemas, filePaths.AllSchemasHCL); err != nil {
				t.Fatalf("failed to write existing AllSchemas to HCL file: %v", err)
			}

			changes := []string{}

			// Call diffSchemas function
			result, err := diffSchemas(test.newSchemas, test.lastSchemas, &changes, filePaths)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("expected non-nil result")
			}

			// Verify changes array
			if len(changes) != len(test.expectedChanges) {
				t.Fatalf("expected %d changes, but got %d: %v", len(test.expectedChanges), len(changes), changes)
			}

			// Check that all expected changes are present (order-independent)
			// Make a copy of expected changes to track which ones we've seen
			expectedSet := make(map[string]bool)
			for expectedChange := range test.expectedChanges {
				expectedSet[expectedChange] = true
			}

			for _, change := range changes {
				if !expectedSet[change] {
					t.Fatalf("unexpected change found: %q. Expected changes: %v, Got changes: %v", change, test.expectedChanges, changes)
				}
				delete(expectedSet, change)
			}

			// Check that all expected changes were found
			if len(expectedSet) > 0 {
				var missing []string
				for change := range expectedSet {
					missing = append(missing, change)
				}
				t.Fatalf("missing expected changes: %v. Got changes: %v", missing, changes)
			}

			// Verify the updated AllSchemas structure
			if len(result.Resources) != test.expectedResources {
				t.Fatalf("expected %d resources in result, but got %d", test.expectedResources, len(result.Resources))
			}

			// For cases with changes, verify the resources were properly updated
			if test.expectChanges {
				// Verify resources are sorted by ResourceTypeName
				for i := 1; i < len(result.Resources); i++ {
					if result.Resources[i-1].ResourceTypeName >= result.Resources[i].ResourceTypeName {
						t.Fatalf("resources are not sorted by ResourceTypeName")
					}
				}

				// Check specific updates based on test case
				switch name {
				case "new_resource":
					// Should have both S3 bucket and EC2 instance
					found := make(map[string]bool)
					for _, resource := range result.Resources {
						found[resource.CloudFormationTypeName] = true
					}
					if !found["AWS::S3::Bucket"] || !found["AWS::EC2::Instance"] {
						t.Fatal("expected both AWS::S3::Bucket and AWS::EC2::Instance in result")
					}

				case "changed_resource":
					// Should have updated SuppressPluralDataSourceGeneration to true
					var s3Resource *allschemas.ResourceAllSchema
					for i := range result.Resources {
						if result.Resources[i].CloudFormationTypeName == "AWS::S3::Bucket" {
							s3Resource = &result.Resources[i]
							break
						}
					}
					if s3Resource == nil {
						t.Fatal("expected AWS::S3::Bucket in result")
					}
					if !s3Resource.SuppressPluralDataSourceGeneration {
						t.Fatal("expected SuppressPluralDataSourceGeneration to be true for AWS::S3::Bucket")
					}

				case "mixed_changes":
					// Verify all three resources exist and S3 bucket is updated
					found := make(map[string]*allschemas.ResourceAllSchema)
					for i := range result.Resources {
						found[result.Resources[i].CloudFormationTypeName] = &result.Resources[i]
					}

					if len(found) != 3 {
						t.Fatalf("expected 3 resources, got %d", len(found))
					}

					s3Resource := found["AWS::S3::Bucket"]
					if s3Resource == nil || !s3Resource.SuppressPluralDataSourceGeneration {
						t.Fatal("expected AWS::S3::Bucket with SuppressPluralDataSourceGeneration=true")
					}

					ec2Resource := found["AWS::EC2::Instance"]
					if ec2Resource == nil {
						t.Fatal("expected AWS::EC2::Instance in result")
					}

					rdsResource := found["AWS::RDS::DBInstance"]
					if rdsResource == nil {
						t.Fatal("expected AWS::RDS::DBInstance in result")
					}
				}
			}

			// For no changes case, verify structure remains the same
			if !test.expectChanges {
				if !reflect.DeepEqual(result.Resources, test.existingAllSchemas.Resources) {
					t.Fatal("expected AllSchemas to remain unchanged when no changes detected")
				}
			}
		})
	}
}

func TestDiffSchemasFileOperations(t *testing.T) {
	t.Run("missing_all_schemas_file", func(t *testing.T) {
		// Create temporary directory
		tempDir := t.TempDir()

		// Mock file paths with non-existent file
		filePaths := &UpdateFilePaths{
			AllSchemasHCL: filepath.Join(tempDir, "nonexistent.hcl"),
		}

		newSchemas := &allschemas.AvailableSchemas{
			Resources: []allschemas.ResourceSchema{
				{
					CloudFormationTypeName:             "AWS::S3::Bucket",
					ResourceTypeName:                   "aws_s3_bucket",
					SuppressPluralDataSourceGeneration: false,
				},
			},
		}

		lastSchemas := &allschemas.AvailableSchemas{Resources: []allschemas.ResourceSchema{}}
		changes := []string{}

		// This should fail because the AllSchemas file doesn't exist
		_, err := diffSchemas(newSchemas, lastSchemas, &changes, filePaths)

		if err == nil {
			t.Fatal("expected error when AllSchemas file doesn't exist")
		}

		if !strings.Contains(err.Error(), "failed to parse existing allSchemas") {
			t.Fatalf("expected error about parsing AllSchemas, but got: %v", err)
		}
	})

	t.Run("file_write_verification", func(t *testing.T) {
		// Create temporary directory
		tempDir := t.TempDir()

		// Mock file paths
		filePaths := &UpdateFilePaths{
			AllSchemasHCL: filepath.Join(tempDir, "all_schemas.hcl"),
		}

		// Create initial AllSchemas
		existingAllSchemas := &allschemas.AllSchemas{
			Resources: []allschemas.ResourceAllSchema{
				{
					CloudFormationTypeName:             "AWS::S3::Bucket",
					ResourceTypeName:                   "aws_s3_bucket",
					SuppressPluralDataSourceGeneration: false,
				},
			},
		}

		// Write initial file
		if err := writeSchemasToHCLFile(existingAllSchemas, filePaths.AllSchemasHCL); err != nil {
			t.Fatalf("failed to write initial AllSchemas: %v", err)
		}

		newSchemas := &allschemas.AvailableSchemas{
			Resources: []allschemas.ResourceSchema{
				{
					CloudFormationTypeName:             "AWS::S3::Bucket",
					ResourceTypeName:                   "aws_s3_bucket",
					SuppressPluralDataSourceGeneration: false,
				},
				{
					CloudFormationTypeName:             "AWS::EC2::Instance",
					ResourceTypeName:                   "aws_ec2_instance",
					SuppressPluralDataSourceGeneration: false,
				},
			},
		}

		lastSchemas := &allschemas.AvailableSchemas{
			Resources: []allschemas.ResourceSchema{
				{
					CloudFormationTypeName:             "AWS::S3::Bucket",
					ResourceTypeName:                   "aws_s3_bucket",
					SuppressPluralDataSourceGeneration: false,
				},
			},
		}

		changes := []string{}

		// Call diffSchemas
		result, err := diffSchemas(newSchemas, lastSchemas, &changes, filePaths)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify file was written
		if _, err := os.Stat(filePaths.AllSchemasHCL); os.IsNotExist(err) {
			t.Fatal("expected AllSchemas file to be written")
		}

		// Verify we can read back the file and get the same result
		readBackSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
		if err != nil {
			t.Fatalf("failed to read back AllSchemas file: %v", err)
		}

		if len(readBackSchemas.Resources) != len(result.Resources) {
			t.Fatalf("expected %d resources in file, but got %d", len(result.Resources), len(readBackSchemas.Resources))
		}

		// Check that both S3 and EC2 resources are present
		found := make(map[string]bool)
		for _, resource := range readBackSchemas.Resources {
			found[resource.CloudFormationTypeName] = true
		}

		if !found["AWS::S3::Bucket"] || !found["AWS::EC2::Instance"] {
			t.Fatal("expected both AWS::S3::Bucket and AWS::EC2::Instance in written file")
		}
	})
}
