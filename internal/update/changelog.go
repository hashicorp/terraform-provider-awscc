package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

// generateDataSourceChanges takes changes and filepaths, parses allSchemas, and adds data source entries
func generateDataSourceChanges(changes []string, filePaths *UpdateFilePaths) ([]string, error) {
	fmt.Printf("generateDataSourceChanges called with %d changes\n", len(changes))

	// Parse the AllSchemasHCL file to get all schemas
	allSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
	if err != nil {
		return changes, fmt.Errorf("failed to parse AllSchemasHCL: %w", err)
	}

	fmt.Printf("Parsed %d resources from AllSchemasHCL\n", len(allSchemas.Resources))

	// Create a map of CloudFormationTypeName to ResourceAllSchema for quick lookup
	cfTypeToResource := make(map[string]allschemas.ResourceAllSchema)
	for _, resource := range allSchemas.Resources {
		cfTypeToResource[resource.CloudFormationTypeName] = resource
		fmt.Printf("  Found resource: %s\n", resource.CloudFormationTypeName)
	}
	// Process each change entry
	var newChanges []string
	for _, change := range changes {
		fmt.Printf("Processing change: %s\n", change)

		// Separate each entry by " - " into resource and message
		parts := strings.SplitN(change, " - ", 2)
		if len(parts) != 2 {
			fmt.Printf("  Skipping malformed change: %s\n", change)
			newChanges = append(newChanges, change)
			continue
		}

		resource := parts[0]
		message := parts[1]

		fmt.Printf("  Resource: %s, Message: %s\n", resource, message)

		// Add the original change
		newChanges = append(newChanges, change)

		// Check if this is a suppression message - if so, skip data source generation
		if isSuppressionMessage(message) {
			fmt.Printf("  Skipping data source generation for suppressed resource: %s\n", resource)
			continue
		}

		// Check if this entry is in allSchemas.Resources
		if resourceSchema, exists := cfTypeToResource[resource]; exists {
			fmt.Printf("  Found resource %s in allSchemas\n", resource)

			// Check if the resource has a suppression reason set
			if resourceSchema.SuppressionReason != "" {
				fmt.Printf("  Resource %s has suppression reason: %s - skipping data source generation\n", resource, resourceSchema.SuppressionReason)
				continue
			}

			// Check if the resource's plural and singular data sources are suppressed
			if !resourceSchema.SuppressPluralDataSourceGeneration {
				plural := naming.Pluralize(resource)
				pluralChange := fmt.Sprintf("%s - New Plural Data Source", plural)
				newChanges = append(newChanges, pluralChange)
				fmt.Printf("  Added plural data source: %s\n", pluralChange)
			} else {
				fmt.Printf("  Plural data source suppressed for %s\n", resource)
			}

			if !resourceSchema.SuppressSingularDataSourceGeneration {
				singularChange := fmt.Sprintf("%s - New Singular Data Source", resource)
				newChanges = append(newChanges, singularChange)
				fmt.Printf("  Added singular data source: %s\n", singularChange)
			} else {
				fmt.Printf("  Singular data source suppressed for %s\n", resource)
			}
		} else {
			fmt.Printf("  Resource %s not found in allSchemas\n", resource)
		}
	}

	fmt.Printf("generateDataSourceChanges returning %d changes (was %d)\n", len(newChanges), len(changes))
	return newChanges, nil
}

func makeChangelog(changes *[]string, filePaths *UpdateFilePaths) error {
	// Parse the CHANGELOG.md file to get existing change

	// Generate additional changes like plural and singular data sources

	// Read the current CHANGELOG.md file

	newChanges, err := generateDataSourceChanges(*changes, filePaths)
	if err != nil {
		return fmt.Errorf("failed to generate data source changes: %w", err)
	}
	originalContent, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		return fmt.Errorf("failed to read CHANGELOG.md: %w", err)
	}

	// Generate the new changelog content
	newContent, err := writeChangelog(string(originalContent), newChanges)
	if err != nil {
		return fmt.Errorf("failed to format changelog: %w", err)
	}

	// Open the file in truncate mode and write the updated content
	file, err := os.OpenFile("CHANGELOG.md", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CHANGELOG.md for writing: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(newContent)
	if err != nil {
		return fmt.Errorf("failed to write to CHANGELOG.md: %w", err)
	}

	return nil
}

// writeChangelog updates the existing changelog content with new changes
func writeChangelog(originalContent string, changes []string) (string, error) {
	fmt.Printf("writeChangelog called with %d changes\n", len(changes))

	if len(changes) == 0 {
		fmt.Println("No changes provided, returning original content")
		return originalContent, nil
	}

	// Sort changes by type
	var (
		newResources   []string
		newDataSources []string
		suppressions   []string
	)

	fmt.Println("Categorizing changes:")
	// Categorize changes
	for _, change := range changes {
		parts := strings.SplitN(change, " - ", 2)
		if len(parts) != 2 {
			fmt.Printf("  Skipping malformed change: %s\n", change)
			continue
		}

		resource := parts[0]
		changeType := parts[1]

		fmt.Printf("  Processing: %s -> %s\n", resource, changeType)

		// Check if this is a suppression message
		if isSuppressionMessage(changeType) {
			suppressions = append(suppressions, resource)
			fmt.Printf("  Categorized as suppression: %s\n", resource)
			continue
		}

		switch changeType {
		case "New Resource":
			newResources = append(newResources, resource)
		case "New Singular Data Source", "New Plural Data Source":
			newDataSources = append(newDataSources, resource)
		}
	}

	// Sort each category alphabetically
	sort.Strings(newResources)
	sort.Strings(newDataSources)
	sort.Strings(suppressions)

	fmt.Printf("After sorting - Data Sources: %d, Resources: %d, Suppressions: %d\n", len(newDataSources), len(newResources), len(suppressions))

	// Build the new changelog entries
	var newEntries []string

	// Add data sources first
	for _, ds := range newDataSources {
		entry := fmt.Sprintf("* **New Data Source:** `%s`", ds)
		newEntries = append(newEntries, entry)
		fmt.Printf("  Added data source entry: %s\n", entry)
	}

	// Add resources
	for _, res := range newResources {
		entry := fmt.Sprintf("* **New Resource:** `%s`", res)
		newEntries = append(newEntries, entry)
		fmt.Printf("  Added resource entry: %s\n", entry)
	}

	// Note: Suppressions are tracked but not added to the public changelog
	if len(suppressions) > 0 {
		fmt.Printf("  Tracked %d suppressions (not added to changelog): %v\n", len(suppressions), suppressions)
	}

	if len(newEntries) == 0 {
		fmt.Println("No valid entries to add, returning original content")
		return originalContent, nil
	}

	fmt.Printf("Total entries to add: %d\n", len(newEntries))

	// Parse and increment version
	newVersion, err := parseAndIncrementVersion(originalContent)
	if err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	// Create new version header with current date
	currentDate := time.Now().Format("January 2, 2006")
	newVersionHeader := fmt.Sprintf("## %s (%s)", newVersion, currentDate)
	fmt.Printf("Creating new version header: %s\n", newVersionHeader)

	// Build the complete new section
	var newSection []string
	newSection = append(newSection, newVersionHeader)
	newSection = append(newSection, "")
	newSection = append(newSection, "FEATURES:")
	newSection = append(newSection, "")
	newSection = append(newSection, newEntries...)
	newSection = append(newSection, "")

	// Add the new section to the top of the changelog
	lines := strings.Split(originalContent, "\n")
	result := make([]string, 0, len(lines)+len(newSection))
	result = append(result, newSection...)
	result = append(result, lines...)

	fmt.Printf("Final result has %d lines (original: %d)\n", len(result), len(lines))
	return strings.Join(result, "\n"), nil
}

// parseAndIncrementVersion finds the latest version in changelog content and increments the minor version
func parseAndIncrementVersion(changelogContent string) (string, error) {
	lines := strings.Split(changelogContent, "\n")

	for _, line := range lines {
		// Look for version headers like "## 1.47.0 (June 26, 2025)"
		if strings.HasPrefix(strings.TrimSpace(line), "## ") && strings.Contains(line, ".") {
			fmt.Printf("Found version line: %s\n", line)

			// Extract the version part
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				versionStr := parts[1] // Should be something like "1.47.0"

				// Split version into major.minor.patch
				versionParts := strings.Split(versionStr, ".")
				if len(versionParts) >= 2 {
					major, err := strconv.Atoi(versionParts[0])
					if err != nil {
						return "", fmt.Errorf("failed to parse major version: %w", err)
					}

					minor, err := strconv.Atoi(versionParts[1])
					if err != nil {
						return "", fmt.Errorf("failed to parse minor version: %w", err)
					}

					// Increment minor version and reset patch to 0
					newVersion := fmt.Sprintf("%d.%d.0", major, minor+1)
					fmt.Printf("Incremented version from %s to %s\n", versionStr, newVersion)
					return newVersion, nil
				}
			}
		}
	}

	// If no version found, start with a default
	return "1.48.0", fmt.Errorf("no version found in changelog, using default")
}

// isSuppressionMessage checks if a message indicates suppression
func isSuppressionMessage(message string) bool {
	suppressionKeywords := []string{
		"New Resource Suppression",
		"New Singular Data Source Suppression",
		"New Plural Data Source Suppression",
		"Suppressed Resource",
		"Suppression",
	}

	for _, keyword := range suppressionKeywords {
		if strings.Contains(message, keyword) {
			return true
		}
	}

	return false
}
