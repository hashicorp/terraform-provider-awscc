package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func parseChangeLog() ([]string, error) {
	// Open CHANGELOG.md file
	file, err := os.Open("CHANGELOG.md")
	if err != nil {
		return nil, fmt.Errorf("failed to open CHANGELOG.md: %w", err)
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
		return nil, fmt.Errorf("error reading CHANGELOG.md: %w", err)
	}

	return changes, nil
}

func makeChangelog(newSchemas *allschemas.AvailableSchemas, lastSchemas *allschemas.AvailableSchemas, filePaths *UpdateFilePaths) error {
	// Parse the CHANGELOG.md file to get existing changes
	existingChanges, err := parseChangeLog()
	if err != nil {
		return fmt.Errorf("failed to parse CHANGELOG.md: %w", err)
	}

	// Find new resources by comparing schemas
	var schemaChanges []string
	for _, newResource := range newSchemas.Resources {
		// Check if resource exists in lastSchemas
		found := false
		for _, lastResource := range lastSchemas.Resources {
			if newResource.CloudFormationTypeName == lastResource.CloudFormationTypeName {
				found = true
				break
			}
		}

		// If not found, it's a new resource
		if !found {
			schemaChanges = append(schemaChanges, fmt.Sprintf("%s - New Resource", newResource.CloudFormationTypeName))
		}
	}

	// Combine existing changes with new schema changes
	allChanges := append(existingChanges, schemaChanges...)

	// Generate additional changes like plural and singular data sources
	err = generateChanges(&allChanges, filePaths)
	if err != nil {
		return fmt.Errorf("failed to generate changes: %w", err)
	}

	// Read the current CHANGELOG.md file
	originalContent, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		return fmt.Errorf("failed to read CHANGELOG.md: %w", err)
	}

	// Generate the new changelog content
	newContent, err := formatChangelog(string(originalContent), allChanges)
	if err != nil {
		return fmt.Errorf("failed to format changelog: %w", err)
	}

	// Write the updated content back to CHANGELOG.md
	err = os.WriteFile("CHANGELOG.md", []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to CHANGELOG.md: %w", err)
	}

	return nil
}

func generateChanges(changes *[]string, filePaths *UpdateFilePaths) error {
	if changes == nil || len(*changes) == 0 {
		return nil
	}

	type changelogParts struct {
		Change   string
		Resource string
	}

	temp := make([]changelogParts, 0)

	for _, change := range *changes {
		parts := strings.SplitN(change, " - ", 2)
		if len(parts) == 2 {
			temp = append(temp, changelogParts{
				Change:   parts[1],
				Resource: parts[0],
			})
		} else {
			continue
		}
	}

	v, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
	if err != nil {
		return fmt.Errorf("failed to parse existing allSchemas: %w", err)
	}
	// Create a map of cfTypeName to allschemas.Resource
	cfTypeToResource := make(map[string]allschemas.ResourceAllSchema)
	for _, resource := range v.Resources {
		cfTypeToResource[resource.CloudFormationTypeName] = resource
	}

	for _, part := range temp {
		if part.Change == "New Resource" {
			if t, exists := cfTypeToResource[part.Resource]; exists {
				if !t.SuppressPluralDataSourceGeneration {
					*changes = append(*changes, "New Plural Data Source: "+t.CloudFormationTypeName)
				}
				if !t.SuppressSingularDataSourceGeneration {
					*changes = append(*changes, "New Singular Data Source: "+t.CloudFormationTypeName)
				}

			}
		}
	}

	sort.Strings(*changes)
	return nil
}

// formatChangelog updates the existing changelog content with new changes
func formatChangelog(originalContent string, changes []string) (string, error) {
	// Sort changes by type
	var (
		newResources   []string
		newDataSources []string
	)

	// Categorize changes
	for _, change := range changes {
		parts := strings.SplitN(change, " - ", 2)
		if len(parts) != 2 {
			continue
		}

		resource := parts[0]
		changeType := parts[1]

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

	// Find the "Unreleased" section in the original content
	scanner := bufio.NewScanner(strings.NewReader(originalContent))
	var (
		lines          []string
		unreleaseFound bool
		featuresFound  bool
		insertPosition int
	)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if !unreleaseFound && strings.Contains(line, "## 1.") && strings.Contains(line, "(Unreleased)") {
			unreleaseFound = true
			continue
		}

		if unreleaseFound && !featuresFound && strings.TrimSpace(line) == "FEATURES:" {
			featuresFound = true
			insertPosition = len(lines)
			continue
		}

		if unreleaseFound && featuresFound && strings.TrimSpace(line) != "" && !strings.HasPrefix(strings.TrimSpace(line), "*") {
			insertPosition = len(lines)
			break
		}
	}

	// If we found the Unreleased section but not FEATURES, add it
	if unreleaseFound && !featuresFound {
		lines = append(lines[:insertPosition], append([]string{"", "FEATURES:", ""}, lines[insertPosition:]...)...)
		insertPosition = insertPosition + 3
	}

	// If we didn't find an Unreleased section, create one
	if !unreleaseFound {
		// Get the current version from the first line that contains a version
		var currentVersion string
		for _, line := range lines {
			if strings.HasPrefix(line, "## ") && strings.Contains(line, ".") {
				parts := strings.Split(line, " ")
				if len(parts) >= 2 {
					currentVersion = parts[1]
					break
				}
			}
		}

		// Calculate the next version
		if currentVersion != "" {
			parts := strings.Split(currentVersion, ".")
			if len(parts) >= 3 {
				major, _ := strconv.Atoi(parts[0])
				minor, _ := strconv.Atoi(parts[1])
				// patch value is not used for incrementing the version

				minor++ // Increment the minor version
				nextVersion := fmt.Sprintf("%d.%d.0", major, minor)

				// Insert the new Unreleased section at the beginning
				lines = append([]string{fmt.Sprintf("## %s (Unreleased)", nextVersion), "", "FEATURES:", ""}, lines...)
				insertPosition = 4
			}
		}
	}

	// Create new changelog content
	var newLines []string

	// Add data sources
	for _, ds := range newDataSources {
		newLines = append(newLines, fmt.Sprintf("* **New Data Source:** `%s`", ds))
	}

	// Add resources
	for _, res := range newResources {
		newLines = append(newLines, fmt.Sprintf("* **New Resource:** `%s`", res))
	}

	// If we have changes, add them to the changelog
	if len(newLines) > 0 {
		if insertPosition > 0 && insertPosition < len(lines) {
			lines = append(lines[:insertPosition], append(newLines, lines[insertPosition:]...)...)
		} else {
			lines = append(lines, newLines...)
		}
	}

	return strings.Join(lines, "\n"), nil
}
