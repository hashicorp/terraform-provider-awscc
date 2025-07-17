// Package main provides changelog generation and management functionality.
// This file handles creating and updating changelog entries for new resources,
// data sources, and schema changes in the Terraform provider.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

const (
	// changelogSplitParts represents the expected number of parts when splitting a change entry by " - "
	changelogSplitParts = 2
	// changelogFileMode represents the file permissions for CHANGELOG.md (readable by all, writable by owner)
	changelogFileMode = 0644
)

// generateDataSourceChanges expands resource changes to include corresponding data source entries.
// For each new resource, it automatically generates entries for both singular and plural data sources,
// taking into account suppression settings from the schema configuration.
//
// Parameters:
//   - changes: Slice of existing change descriptions
//   - filePaths: Configuration containing paths to schema files
//
// Returns an expanded slice including original changes plus auto-generated data source entries.
func generateDataSourceChanges(changes []string, filePaths *UpdateFilePaths) ([]string, error) {
	log.Printf("generateDataSourceChanges called with %d changes\n", len(changes))

	// Load all schema configurations to check suppression settings
	allSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
	if err != nil {
		return changes, fmt.Errorf("failed to parse AllSchemasHCL: %w", err)
	}

	log.Printf("Parsed %d resources from AllSchemasHCL\n", len(allSchemas.Resources))

	// Create lookup map for quick resource schema access by CloudFormation type name
	cfTypeToResource := make(map[string]allschemas.ResourceAllSchema)
	for _, resource := range allSchemas.Resources {
		cfTypeToResource[resource.CloudFormationTypeName] = resource
	}

	// Process each change and generate corresponding data source entries
	var newChanges []string
	for _, change := range changes {
		log.Printf("Processing change: %s\n", change)

		// Parse change entry format: "Description - Message"
		parts := strings.SplitN(change, " - ", changelogSplitParts)
		if len(parts) != changelogSplitParts {
			log.Printf("  Skipping malformed change: %s\n", change)
			newChanges = append(newChanges, change)
			continue
		}

		message := parts[0]
		resource := parts[1]

		log.Printf("  Resource: %s, Message: %s\n", resource, message)

		// Always include the original resource change
		newChanges = append(newChanges, change)

		// Generate data source entries if the resource exists in schemas and isn't suppressed
		if resourceSchema, exists := cfTypeToResource[resource]; exists {
			// Generate plural data source entry if not suppressed
			if !resourceSchema.SuppressPluralDataSourceGeneration {
				plural := naming.Pluralize(resource)
				pluralChange := fmt.Sprintf("%s - New Plural Data Source", plural)
				newChanges = append(newChanges, pluralChange)
				log.Printf("  Added plural data source: %s\n", pluralChange)
			} else {
				log.Printf("  Plural data source suppressed for %s\n", resource)
			}

			// Generate singular data source entry if not suppressed
			if !resourceSchema.SuppressSingularDataSourceGeneration {
				singularChange := fmt.Sprintf("%s - New Singular Data Source", resource)
				newChanges = append(newChanges, singularChange)
				log.Printf("  Added singular data source: %s\n", singularChange)
			} else {
				log.Printf("  Singular data source suppressed for %s\n", resource)
			}
		} else {
			newChanges = append(newChanges, change)
		}
	}

	log.Printf("generateDataSourceChanges returning %d changes (was %d)\n", len(newChanges), len(changes))
	return newChanges, nil
}

// makeChangelog updates the CHANGELOG.md file with new resource and data source entries.
// It expands resource changes to include corresponding data sources, increments the version,
// and writes the updated changelog back to disk.
//
// Parameters:
//   - changes: Slice of change descriptions to add to the changelog
//   - filePaths: Configuration containing file paths (currently unused but kept for consistency)
//
// Returns an error if changelog processing or file operations fail.
func makeChangelog(changes *[]string, filePaths *UpdateFilePaths) error {
	// Check if there are any changes to process
	if len(*changes) == 0 {
		return fmt.Errorf("no changes provided to changelog - cannot generate changelog entry")
	}

	// Generate additional entries for data sources based on resource changes
	newChanges, err := generateDataSourceChanges(*changes, filePaths)
	if err != nil {
		return fmt.Errorf("failed to generate data source changes: %w", err)
	}

	// Check if after processing there are still no valid changes
	if len(newChanges) == 0 {
		return fmt.Errorf("no valid changes found after processing - cannot generate changelog entry")
	}

	// Read the current changelog content
	originalContent, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		return fmt.Errorf("failed to read CHANGELOG.md: %w", err)
	}

	// Generate the new changelog content
	newContent := writeChangelog(string(originalContent), newChanges)

	// Check if the content actually changed (writeChangelog returns original if no entries)
	if newContent == string(originalContent) {
		return fmt.Errorf("no valid changelog entries generated - cannot update changelog")
	}

	// Open the file in truncate mode and write the updated content
	file, err := os.OpenFile("CHANGELOG.md", os.O_WRONLY|os.O_TRUNC, changelogFileMode)
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
func writeChangelog(originalContent string, changes []string) string {
	log.Printf("writeChangelog called with %d changes\n", len(changes))

	if len(changes) == 0 {
		log.Println("No changes provided, returning original content")
		return originalContent
	}

	// Sort changes by type
	var (
		newResources   []string
		newDataSources []string
		suppressions   []string
	)

	log.Println("Categorizing changes:")
	// Categorize changes
	for _, change := range changes {
		parts := strings.SplitN(change, " - ", changelogSplitParts)
		if len(parts) != changelogSplitParts {
			log.Printf("  Skipping malformed change: %s\n", change)
			continue
		}

		resource := parts[0]
		changeType := parts[1]

		if parts[0] != "" || parts[1] != "" {
			log.Printf("Processing change: %s - %s\n", resource, changeType)
			continue
		}

		log.Printf("  Processing: %s -> %s\n", resource, changeType)

		switch changeType {
		case "New Resource":
			newResources = append(newResources, resource)
		case "New Singular Data Source", "New Plural Data Source":
			newDataSources = append(newDataSources, resource)
		default:
			if strings.Contains(changeType, "Suppression") {
				suppressions = append(suppressions, resource)
			}
		}
	}

	// Sort each category alphabetically
	sort.Strings(newResources)
	sort.Strings(newDataSources)
	sort.Strings(suppressions)

	log.Printf("After sorting - Data Sources: %d, Resources: %d, Suppressions: %d\n", len(newDataSources), len(newResources), len(suppressions))

	// Build the new changelog entries
	var newEntries []string

	// Add data sources first
	for _, ds := range newDataSources {
		entry := fmt.Sprintf("* **New Data Source:** `%s`", ds)
		newEntries = append(newEntries, entry)
		log.Printf("  Added data source entry: %s\n", entry)
	}

	// Add resources
	for _, res := range newResources {
		entry := fmt.Sprintf("* **New Resource:** `%s`", res)
		newEntries = append(newEntries, entry)
		log.Printf("  Added resource entry: %s\n", entry)
	}

	for _, sup := range suppressions {
		log.Printf("  Tracked suppression: %s\n", sup)
		entry := fmt.Sprintf("* **Suppressed Resource:** `%s`", sup)
		newEntries = append(newEntries, entry)
	}

	if len(newEntries) == 0 {
		log.Println("No valid entries to add, returning original content")
		return originalContent
	}

	log.Printf("Total entries to add: %d\n", len(newEntries))

	// Parse and increment version
	newVersion, err := parseAndIncrementVersion(originalContent)
	if err != nil {
		log.Printf("Warning: %v\n", err)
	}

	// Create new version header with current date
	currentDate := time.Now().Format("January 2, 2006")
	newVersionHeader := fmt.Sprintf("## %s (%s)", newVersion, currentDate)
	log.Printf("Creating new version header: %s\n", newVersionHeader)

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

	log.Printf("Final result has %d lines (original: %d)\n", len(result), len(lines))
	return strings.Join(result, "\n")
}

// parseAndIncrementVersion finds the latest version in changelog content and increments the minor version
func parseAndIncrementVersion(changelogContent string) (string, error) {
	lines := strings.Split(changelogContent, "\n")

	for _, line := range lines {
		// Look for version headers like "## 1.47.0 (June 26, 2025)"
		if strings.HasPrefix(strings.TrimSpace(line), "## ") && strings.Contains(line, ".") {
			log.Printf("Found version line: %s\n", line)

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
					log.Printf("Incremented version from %s to %s\n", versionStr, newVersion)
					return newVersion, nil
				}
			}
		}
	}

	// If no version found, start with a default
	return "1.48.0", fmt.Errorf("no version found in changelog, using default")
}
