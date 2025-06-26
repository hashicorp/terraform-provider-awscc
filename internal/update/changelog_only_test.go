package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestChangelogFunctions(t *testing.T) {
	// Sample original content
	originalContent := `## 1.46.0 (Unreleased)

## 1.45.0 (June 12, 2025)

FEATURES:

* provider: Updated resource schemas
`

	// Sample changes
	changes := []string{
		"AWS::ApiGateway::Resource - New Resource",
		"awscc_apigateway_resource - New Singular Data Source",
	}

	// Format the changelog
	result, err := formatChangelog(originalContent, changes)
	if err != nil {
		t.Fatalf("formatChangelog failed: %v", err)
	}

	// Check that the formatted content contains our changes
	if !contains(result, "* **New Resource:** `AWS::ApiGateway::Resource`") {
		t.Error("New resource not found in formatted changelog")
	}

	if !contains(result, "* **New Data Source:** `awscc_apigateway_resource`") {
		t.Error("New data source not found in formatted changelog")
	}
}

func TestParseChangeLogSample(t *testing.T) {
	// We can't really test against the actual CHANGELOG.md in a unit test
	// But we can verify that the function works with a sample changelog
	sampleChangelog := "## 1.46.0 (Unreleased)\n\n" +
		"## 1.45.0 (June 12, 2025)\n\n" +
		"FEATURES:\n\n" +
		"* **New Data Source:** `awscc_apigateway_resource`\n" +
		"* **New Data Source:** `awscc_apigateway_resources`\n" +
		"* **New Resource:** `awscc_apigateway_method`\n\n" +
		"## 1.44.0 (June 5, 2025)\n\n" +
		"FEATURES:\n\n" +
		"* **New Resource:** `awscc_lambda_function`\n"

	// Create a temporary file with the sample content
	tmpfile, err := os.CreateTemp("", "sample-changelog-*.md")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(sampleChangelog)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Create a function to parse the sample changelog
	parseFunction := func() ([]string, error) {
		// Open CHANGELOG.md file
		file, err := os.Open(tmpfile.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// Variables to track parsing state
		var (
			changes          []string
			inCurrentRelease = false
			inFeatures       = false
		)

		// Parse the file line by line
		for scanner.Scan() {
			line := scanner.Text()
			trimmedLine := strings.TrimSpace(line)

			// Look for version headers (e.g., "## 1.45.0 (June 12, 2025)")
			if strings.HasPrefix(trimmedLine, "## ") {
				// If we find a version header, we're entering a new release section
				if strings.Contains(trimmedLine, "(Unreleased)") {
					// Skip the unreleased section
					inCurrentRelease = false
					continue
				}

				// Extract version number
				versionParts := strings.Split(trimmedLine, " ")
				if len(versionParts) >= 2 {
					inCurrentRelease = true
					inFeatures = false
				}
				continue
			}

			// Check if we're in the FEATURES section
			if inCurrentRelease && trimmedLine == "FEATURES:" {
				inFeatures = true
				continue
			}

			// If we hit another section header, we're no longer in FEATURES
			if inCurrentRelease && inFeatures && trimmedLine != "" && !strings.HasPrefix(trimmedLine, "*") && !strings.HasPrefix(trimmedLine, "**") {
				inFeatures = false
				continue
			}

			// Process feature lines
			if inCurrentRelease && inFeatures && strings.HasPrefix(trimmedLine, "**New Resource:**") {
				resourceName := strings.TrimSpace(strings.TrimPrefix(trimmedLine, "**New Resource:**"))
				resourceName = strings.Trim(resourceName, "`")
				changes = append(changes, fmt.Sprintf("%s - New Resource", resourceName))
				continue
			}

			if inCurrentRelease && inFeatures && strings.HasPrefix(trimmedLine, "**New Data Source:**") {
				dataSourceName := strings.TrimSpace(strings.TrimPrefix(trimmedLine, "**New Data Source:**"))
				dataSourceName = strings.Trim(dataSourceName, "`")

				// Check if it's a singular or plural data source based on naming convention
				if strings.HasSuffix(dataSourceName, "s") && !strings.HasSuffix(dataSourceName, "ss") && !strings.HasSuffix(dataSourceName, "status") {
					changes = append(changes, fmt.Sprintf("%s - New Plural Data Source", dataSourceName))
				} else {
					changes = append(changes, fmt.Sprintf("%s - New Singular Data Source", dataSourceName))
				}
				continue
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}

		return changes, nil
	}

	// Call the function
	changes, err := parseFunction()
	if err != nil {
		t.Fatalf("parseFunction failed: %v", err)
	}

	// Verify the results
	expectedChanges := []string{
		"awscc_apigateway_resource - New Singular Data Source",
		"awscc_apigateway_resources - New Plural Data Source",
		"awscc_apigateway_method - New Resource",
		"awscc_lambda_function - New Resource",
	}

	if len(changes) != len(expectedChanges) {
		t.Errorf("Expected %d changes, got %d", len(expectedChanges), len(changes))
	}

	// Check that all expected changes are present
	for _, expected := range expectedChanges {
		found := false
		for _, actual := range changes {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected change not found: %s", expected)
		}
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
