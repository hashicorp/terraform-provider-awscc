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
		cfTypeToResource[resource.ResourceTypeName] = resource
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

		resource := parts[0]
		message := parts[1]

		log.Printf("  Resource: %s, Message: %s\n", resource, message)

		// Always include the original resource change
		newChanges = append(newChanges, change)
		fmt.Println(newChanges)
		// Generate data source entries if the resource exists in schemas and isn't suppressed
		if resourceSchema, exists := cfTypeToResource[resource]; exists {
			// Generate plural data source entry if not suppressed
			if !resourceSchema.SuppressPluralDataSourceGeneration {
				plural := naming.Pluralize(resource)
				pluralChange := fmt.Sprintf("* **New Data Source:** `%s`", plural)
				newChanges = append(newChanges, pluralChange)
				log.Printf("Added plural data source: %s\n", pluralChange)
			} else {
				log.Printf("Plural data source suppressed for %s\n", resource)
			}

			// Generate singular data source entry if not suppressed
			if !resourceSchema.SuppressSingularDataSourceGeneration {
				singularChange := fmt.Sprintf("* **New Data Source:** `%s`", resource)
				newChanges = append(newChanges, singularChange)
				log.Printf("Added singular data source: %s\n", singularChange)
			} else {
				log.Printf("Singular data source suppressed for %s\n", resource)
			}
		} else {
			log.Printf("Resource %s not found in schema map\n", resource)
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

	sort.Slice(changes, func(i, j int) bool {
		return changes[i] < changes[j]
	})

	// Parse and increment version
	newVersion, err := parseAndIncrementChangelogVersion(originalContent)
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
	newSection = append(newSection, changes...)
	newSection = append(newSection, "")

	// Add the new section to the top of the changelog
	lines := strings.Split(originalContent, "\n")
	result := make([]string, 0, len(lines)+len(newSection))
	result = append(result, newSection...)
	result = append(result, lines...)

	log.Printf("Final result has %d lines (original: %d)\n", len(result), len(lines))
	return strings.Join(result, "\n")
}

// parseAndIncrementChangelogVersion finds the latest version in changelog content and increments the minor version
func parseAndIncrementChangelogVersion(changelogContent string) (string, error) {
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

func updateVersionFile(filePaths *UpdateFilePaths) error {
	// 1.49.1 -> 1.50.0
	versionBytes, err := os.ReadFile(filePaths.Version)
	if err != nil {
		return fmt.Errorf("failed to read version file %s: %w", filePaths.Version, err)
	}
	versionStr := strings.Split(strings.TrimSpace(string(versionBytes)), ".")
	if len(versionStr) == 0 || versionStr[0] == "" {
		return fmt.Errorf("version file %s is empty", filePaths.Version)
	}

	versionNumber, err := strconv.Atoi(versionStr[1])
	if err != nil {
		return fmt.Errorf("failed to parse version number from %s: %w", filePaths.Version, err)
	}
	versionNumber++ // Increment the version number
	if versionNumber > 999 {
		return fmt.Errorf("version number %d exceeds maximum allowed value", versionNumber)
	}
	versionNumberStr := strconv.Itoa(versionNumber)

	newVersionStr := fmt.Sprintf("%s.%s.%d", versionStr[0], versionNumberStr, 0)
	log.Printf("Updating version file %s to new version: %s\n", filePaths.Version, newVersionStr)
	if err := os.WriteFile(filePaths.Version, []byte(newVersionStr), 0644); err != nil {
		return fmt.Errorf("failed to write new version to %s: %w", filePaths.Version, err)
	}
	log.Printf("Successfully updated version file %s to %s\n", filePaths.Version, newVersionStr)
	return nil
}
