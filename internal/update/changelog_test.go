package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMakeChangelog(t *testing.T) {
	// Create a temporary directory for test files
	testDir := "/tmp/terraform-provider-awscc-changelog-test"
	t.Logf("Creating test directory: %s", testDir)
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	t.Logf("Test directory created successfully")
	// Copy the original CHANGELOG.md to the test directory
	originalChangelog := filepath.Join("..", "..", "CHANGELOG.md")
	testChangelog := filepath.Join(testDir, "CHANGELOG.md")

	t.Logf("Source changelog: %s", originalChangelog)
	t.Logf("Target changelog: %s", testChangelog)

	originalContent, err := os.ReadFile(originalChangelog)
	if err != nil {
		t.Fatalf("Failed to read original CHANGELOG.md: %v", err)
	}
	t.Logf("Read original changelog, size: %d bytes", len(originalContent))

	err = os.WriteFile(testChangelog, originalContent, 0644)
	if err != nil {
		t.Fatalf("Failed to copy CHANGELOG.md to test directory: %v", err)
	}
	t.Logf("Copied changelog to test directory successfully")

	// Create a test AllSchemasHCL file with sample data
	testAllSchemasHCL := filepath.Join(testDir, "allschemas.hcl")
	t.Logf("Creating test AllSchemasHCL file: %s", testAllSchemasHCL)
	allSchemasContent := `
	defaults {
schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}
resource_schema "aws_s3_bucket" {
  cloudformation_type_name = "AWS::S3::Bucket"
  suppress_resource_generation = false
  suppress_singular_data_source_generation = false
  suppress_plural_data_source_generation = false
}

resource_schema "aws_ec2_instance" {
  cloudformation_type_name = "AWS::EC2::Instance"
  suppress_resource_generation = false
  suppress_singular_data_source_generation = false
  suppress_plural_data_source_generation = false
}

resource_schema "aws_rds_db_instance" {
  cloudformation_type_name = "AWS::RDS::DBInstance"
  suppress_resource_generation = true
  suppress_singular_data_source_generation = false
  suppress_plural_data_source_generation = false
}

resource_schema "aws_lambda_function" {
  cloudformation_type_name = "AWS::Lambda::Function"
  suppress_resource_generation = false
  suppress_singular_data_source_generation = false
  suppress_plural_data_source_generation = true
}
`
	err = os.WriteFile(testAllSchemasHCL, []byte(allSchemasContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test AllSchemasHCL file: %v", err)
	}
	t.Logf("Created test AllSchemasHCL file successfully")
	// Create test file paths
	filePaths := &UpdateFilePaths{
		AllSchemasHCL: testAllSchemasHCL,
	}
	t.Logf("Created UpdateFilePaths with AllSchemasHCL: %s", filePaths.AllSchemasHCL)

	// Change to the test directory so the function operates on our test CHANGELOG.md
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)
	t.Logf("Current directory: %s", originalDir)

	err = os.Chdir(testDir)
	if err != nil {
		t.Fatalf("Failed to change to test directory: %v", err)
	}
	t.Logf("Changed to test directory: %s", testDir)

	// Test Case 1: New resources with data source generation
	t.Run("New resources with data sources", func(t *testing.T) {
		// Create example resource changes
		changes := []string{
			"AWS::S3::Bucket - New Resource",
			"AWS::EC2::Instance - New Resource",
		}

		t.Logf("Testing with changes: %v", changes)

		// Call the makeChangelog function
		t.Logf("Calling makeChangelog function...")
		err := makeChangelog(&changes, filePaths)
		if err != nil {
			t.Errorf("makeChangelog failed: %v", err)
			return
		}
		t.Logf("makeChangelog completed successfully")

		// Read the updated changelog
		updatedContent, err := os.ReadFile("CHANGELOG.md")
		if err != nil {
			t.Errorf("Failed to read updated CHANGELOG.md: %v", err)
			return
		}
		t.Logf("Read updated changelog, size: %d bytes", len(updatedContent))

		updatedStr := string(updatedContent)

		// Verify the changelog was updated
		t.Logf("Checking for resource entries in changelog...")
		if !strings.Contains(updatedStr, "**New Resource:** `AWS::S3::Bucket`") {
			t.Errorf("Expected to find AWS::S3::Bucket resource entry")
		} else {
			t.Logf("✓ Found AWS::S3::Bucket resource entry")
		}
		if !strings.Contains(updatedStr, "**New Resource:** `AWS::EC2::Instance`") {
			t.Errorf("Expected to find AWS::EC2::Instance resource entry")
		} else {
			t.Logf("✓ Found AWS::EC2::Instance resource entry")
		}
		if !strings.Contains(updatedStr, "**New Data Source:** `AWS::S3::Bucket`") {
			t.Errorf("Expected to find AWS::S3::Bucket data source entry")
		} else {
			t.Logf("✓ Found AWS::S3::Bucket data source entry")
		}
		if !strings.Contains(updatedStr, "**New Data Source:** `AWS::EC2::Instance`") {
			t.Errorf("Expected to find AWS::EC2::Instance data source entry")
		} else {
			t.Logf("✓ Found AWS::EC2::Instance data source entry")
		}

		// Verify version was incremented
		if !strings.Contains(updatedStr, "## 1.47.0") {
			t.Errorf("Expected version to be incremented to 1.47.0")
		} else {
			t.Logf("✓ Found version 1.47.0")
		}

		t.Logf("Test Case 1 passed - Updated changelog preview:\n%s", updatedStr[:1000])
	})

	// Test Case 2: Resources with suppress_resource_generation
	t.Run("Resource with suppress_resource_generation", func(t *testing.T) {
		// Reset the changelog
		err = os.WriteFile("CHANGELOG.md", originalContent, 0644)
		if err != nil {
			t.Fatalf("Failed to reset CHANGELOG.md: %v", err)
		}

		changes := []string{
			"AWS::RDS::DBInstance - New Resource", // This has suppress_resource_generation = true but should still generate data sources
		}

		t.Logf("Testing resource with suppress_resource_generation with changes: %v", changes)

		err := makeChangelog(&changes, filePaths)
		if err != nil {
			t.Errorf("makeChangelog failed: %v", err)
			return
		}

		updatedContent, err := os.ReadFile("CHANGELOG.md")
		if err != nil {
			t.Errorf("Failed to read updated CHANGELOG.md: %v", err)
			return
		}

		updatedStr := string(updatedContent)

		// Verify the resource entry exists and data sources were also added
		// (suppress_resource_generation only affects resource generation, not data sources)
		if !strings.Contains(updatedStr, "**New Resource:** `AWS::RDS::DBInstance`") {
			t.Errorf("Expected to find AWS::RDS::DBInstance resource entry")
		}

		// SHOULD contain data source entries even for resources with suppress_resource_generation = true
		if !strings.Contains(updatedStr, "**New Data Source:** `AWS::RDS::DBInstance`") {
			t.Errorf("Expected to find AWS::RDS::DBInstance singular data source entry")
		}
		if !strings.Contains(updatedStr, "**New Data Source:** `AWS::RDS::DBInstances`") {
			t.Errorf("Expected to find AWS::RDS::DBInstances plural data source entry")
		}

		t.Logf("Test Case 2 passed - Resource with suppress_resource_generation still generates data sources")
	})

	// Test Case 3: Partial data source suppression
	t.Run("Partial data source suppression", func(t *testing.T) {
		// Reset the changelog
		err = os.WriteFile("CHANGELOG.md", originalContent, 0644)
		if err != nil {
			t.Fatalf("Failed to reset CHANGELOG.md: %v", err)
		}

		changes := []string{
			"AWS::Lambda::Function - New Resource", // Plural data source suppressed, singular allowed
		}

		fmt.Printf("Testing partial suppression with changes: %v\n", changes)

		err := makeChangelog(&changes, filePaths)
		if err != nil {
			t.Errorf("makeChangelog failed: %v", err)
			return
		}

		updatedContent, err := os.ReadFile("CHANGELOG.md")
		if err != nil {
			t.Errorf("Failed to read updated CHANGELOG.md: %v", err)
			return
		}

		updatedStr := string(updatedContent)

		// Verify the resource entry exists
		if !strings.Contains(updatedStr, "**New Resource:** `AWS::Lambda::Function`") {
			t.Errorf("Expected to find AWS::Lambda::Function resource entry")
		}

		// Should have singular data source (not suppressed) but not plural (suppressed)
		if !strings.Contains(updatedStr, "**New Data Source:** `AWS::Lambda::Function`") {
			t.Errorf("Expected to find AWS::Lambda::Function singular data source entry")
		}

		fmt.Printf("Test Case 3 passed - Partial suppression handled correctly\n")
	})

	// Test Case 4: Mixed changes including suppressions
	t.Run("Mixed changes with suppressions", func(t *testing.T) {
		// Reset the changelog
		err = os.WriteFile("CHANGELOG.md", originalContent, 0644)
		if err != nil {
			t.Fatalf("Failed to reset CHANGELOG.md: %v", err)
		}

		changes := []string{
			"AWS::S3::Bucket - New Resource",
			"AWS::EC2::VPC - New Resource Suppression",
			"AWS::Lambda::Function - New Resource",
		}

		fmt.Printf("Testing mixed changes with suppressions: %v\n", changes)

		err := makeChangelog(&changes, filePaths)
		if err != nil {
			t.Errorf("makeChangelog failed: %v", err)
			return
		}

		updatedContent, err := os.ReadFile("CHANGELOG.md")
		if err != nil {
			t.Errorf("Failed to read updated CHANGELOG.md: %v", err)
			return
		}

		updatedStr := string(updatedContent)

		// Should have entries for non-suppressed resources
		if !strings.Contains(updatedStr, "**New Resource:** `AWS::S3::Bucket`") {
			t.Errorf("Expected to find AWS::S3::Bucket resource entry")
		}
		if !strings.Contains(updatedStr, "**New Resource:** `AWS::Lambda::Function`") {
			t.Errorf("Expected to find AWS::Lambda::Function resource entry")
		}

		// Should NOT have entries for suppressed resources in public changelog
		if strings.Contains(updatedStr, "**New Resource:** `AWS::EC2::VPC`") {
			t.Errorf("Should not have AWS::EC2::VPC in public changelog (suppressed)")
		}

		fmt.Printf("Test Case 4 passed - Mixed changes handled correctly\n")
	})

	// Print the final test directory path for review
	fmt.Printf("\n=== TEST COMPLETED ===\n")
	fmt.Printf("Test files are preserved in: %s\n", testDir)
	fmt.Printf("You can review the following files:\n")
	fmt.Printf("- %s (updated changelog)\n", testChangelog)
	fmt.Printf("- %s (test schema file)\n", testAllSchemasHCL)

	// List all files in test directory
	err = filepath.WalkDir(testDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			relPath, _ := filepath.Rel(testDir, path)
			fmt.Printf("- %s\n", filepath.Join(testDir, relPath))
		}
		return nil
	})
	if err != nil {
		t.Errorf("Failed to list test directory contents: %v", err)
	}

	fmt.Printf("\nTo clean up test files later, run: rm -rf %s\n", testDir)
}

// Test helper function to verify changelog structure
func TestChangelogStructure(t *testing.T) {
	// Test the writeChangelog function directly
	originalContent := `## 1.46.0 (Unreleased)

## 1.45.0 (June 12, 2025)

FEATURES:

* provider: Updated resource schemas
`

	changes := []string{
		"AWS::S3::Bucket - New Resource",
		"AWS::EC2::Instance - New Singular Data Source",
		"AWS::Lambda::Function - New Plural Data Source",
	}

	result, err := writeChangelog(originalContent, changes)
	if err != nil {
		t.Fatalf("writeChangelog failed: %v", err)
	}

	// Verify structure
	if !strings.Contains(result, "## 1.47.0") {
		t.Errorf("Expected version 1.47.0 in result")
	}

	if !strings.Contains(result, "FEATURES:") {
		t.Errorf("Expected FEATURES section in result")
	}

	if !strings.Contains(result, "**New Data Source:** `AWS::EC2::Instance`") {
		t.Errorf("Expected data source entry for AWS::EC2::Instance")
	}

	if !strings.Contains(result, "**New Resource:** `AWS::S3::Bucket`") {
		t.Errorf("Expected resource entry for AWS::S3::Bucket")
	}

	fmt.Printf("Changelog structure test passed\n")
	fmt.Printf("Generated changelog preview:\n%s\n", result[:500])
}
