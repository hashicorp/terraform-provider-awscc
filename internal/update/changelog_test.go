// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteChangelog_NewResourceAndDataSource(t *testing.T) {
	dir := t.TempDir()
	allSchemasPath := filepath.Join(dir, "all_schemas.hcl")

	// Write a minimal AllSchemas file with TestResource and no suppression
	allSchemasContent := `

defaults {
  schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

resource_schema "aws_testresource" {
    cloudformation_type_name = "AWS::TestResource"
    suppress_plural_data_source_generation = false
    suppress_singular_data_source_generation = false
    suppress_resource_generation = false
  }
`
	if err := os.WriteFile(allSchemasPath, []byte(allSchemasContent), 0644); err != nil {
		t.Fatalf("failed to write all_schemas.hcl: %v", err)
	}

	filePaths := &UpdateFilePaths{
		AllSchemasHCL: allSchemasPath,
	}

	changes := []string{"* **New Resource** `aws_testresource`"}
	// Call generateDataSourceChanges to expand changes
	expanded, err := generateDataSourceChanges(changes, filePaths)
	log.Printf("Expanded changes: %v", expanded)
	if err != nil {
		t.Fatalf("generateDataSourceChanges failed: %v", err)
	}

	original := "## 1.47.0 (June 26, 2025)\n\nFEATURES:\n\n* Initial release\n"
	result := writeChangelog(original, expanded)

	if !strings.Contains(result, "aws_testresource") {
		t.Errorf("writeChangelog did not include new resource entry")
	}
	if !strings.Contains(result, "* **New Data Source:**") {
		t.Errorf("writeChangelog did not include new data source entry")
	}
	if !strings.Contains(result, "1.48.0") {
		t.Errorf("writeChangelog did not increment version correctly")
	}
}

func TestParseAndIncrementChangelogVersion_Valid(t *testing.T) {
	changelog := "## 1.47.0 (June 26, 2025)\n\nFEATURES:\n\n* Initial release\n"
	version, err := parseAndIncrementChangelogVersion(changelog)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if version != "1.48.0" {
		t.Errorf("expected version 1.48.0, got %s", version)
	}
}

func TestUpdateVersionFile(t *testing.T) {
	tempDir := t.TempDir()
	versionPath := filepath.Join(tempDir, "VERSION")

	// Write an initial version file
	if err := os.WriteFile(versionPath, []byte("1.49.1"), 0644); err != nil {
		t.Fatalf("failed to write version file: %v", err)
	}

	filePaths := &UpdateFilePaths{Version: versionPath}
	if err := updateVersionFile(filePaths); err != nil {
		t.Fatalf("updateVersionFile failed: %v", err)
	}

	// Read back the version file
	updated, err := os.ReadFile(versionPath)
	if err != nil {
		t.Fatalf("failed to read updated version file: %v", err)
	}
	if string(updated) != "1.50.0" {
		t.Errorf("expected version to be incremented to 1.50.0, got %q", string(updated))
	}
}
