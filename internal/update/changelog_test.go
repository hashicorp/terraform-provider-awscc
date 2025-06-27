package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseChangeLog(t *testing.T) {
	changes, err := parseChangeLogFromFile("../../CHANGELOG.md")
	if err != nil {
		t.Fatalf("Failed to parse CHANGELOG.md: %v", err)
	}

	// Print all the parsed changes
	fmt.Println("Parsed changes:")
	for _, change := range changes {
		fmt.Println("-", change)
	}

	// Check if we have any changes
	if len(changes) == 0 {
		t.Error("No changes were parsed from CHANGELOG.md")
	}
}

func TestFormatChangelog(t *testing.T) {
	// Sample original content
	originalContent := `## 1.46.0 (Unreleased)

## 1.45.0 (June 12, 2025)

FEATURES:

* provider: Updated resource schemas

## 1.44.0 (June 5, 2025)
`

	// Sample changes
	changes := []string{
		"AWS::Example::Resource - New Resource",
		"AWS::Example::Function - New Resource",
		"awscc_example_resource - New Singular Data Source",
		"awscc_example_resources - New Plural Data Source",
		"awscc_example_function - New Singular Data Source",
	}

	// Format the changelog
	result, err := writeChangelog(originalContent, changes)
	if err != nil {
		t.Fatalf("formatChangelog failed: %v", err)
	}

	// Verify the result contains the expected elements
	for _, change := range changes {
		parts := strings.SplitN(change, " - ", 2)
		if len(parts) != 2 {
			continue
		}

		resource := parts[0]
		changeType := parts[1]

		var expectedText string
		if changeType == "New Resource" {
			expectedText = fmt.Sprintf("* **New Resource:** `%s`", resource)
		} else {
			expectedText = fmt.Sprintf("* **New Data Source:** `%s`", resource)
		}

		if !strings.Contains(result, expectedText) {
			t.Errorf("formatChangelog output missing expected change: %s", expectedText)
		}
	}

	// Check the overall structure
	if !strings.Contains(result, "## 1.46.0 (Unreleased)") {
		t.Error("formatChangelog output missing Unreleased section")
	}

	if !strings.Contains(result, "FEATURES:") {
		t.Error("formatChangelog output missing FEATURES section")
	}
}
