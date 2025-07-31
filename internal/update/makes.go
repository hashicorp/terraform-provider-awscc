// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

// Build type constants define the different types of make operations that can be performed.
const (
	BuildTypeSchemas             = "schemas"               // Generate CloudFormation schemas
	BuildTypeResources           = "resources"             // Generate Terraform resources
	BuildTypeSingularDataSources = "singular-data-sources" // Generate singular data sources
	BuildTypePluralDataSources   = "plural-data-sources"   // Generate plural data sources

	// makesErrorFileMode represents the file permissions for error log files
	makesErrorFileMode = 0600
)

// makeBuild serves as the main api to orchestrate different types of Terraform provider components.
// It is run in the order make schemas, make resources, make singular-data-sources, make plural-data-sources
// Each make command is run and then errors are handled accordingly to the sub functions
// The current make command with a grep for "error" piped to an output file until there are no more errors
// After a failed make-singular-data-sources or make-plural-data-sources, it runs make schemas again to generate schemas
//
// Parameters:
//   - ctx: Context for cancellation and logging
//   - config: GitHub configuration for issue creation
//   - currentSchemas: Current schema definitions
//   - buildType: Type of build to perform (schemas, resources, data-sources)
//   - changes: Slice to append change descriptions to
//   - filePaths: Configuration paths for various files
//   - isNewMap: Map tracking which resources are new vs existing
func makeBuild(ctx context.Context, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	// Validate buildType parameter
	if buildType != BuildTypeSchemas && buildType != BuildTypeResources && buildType != BuildTypeSingularDataSources && buildType != BuildTypePluralDataSources {
		return fmt.Errorf("invalid build type: %s, must be '%s', '%s', '%s', or '%s'", buildType, BuildTypeSchemas, BuildTypeResources, BuildTypeSingularDataSources, BuildTypePluralDataSources)
	}

	// Open error log file for capturing make command errors
	file, err := os.OpenFile(filePaths.RunMakesErrors, os.O_RDWR|os.O_CREATE|os.O_TRUNC, makesErrorFileMode)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %w", err)
	}
	defer file.Close()

	hasErrors := true

	// Process errors iteratively until all are handled
	for hasErrors {
		log.Println("Running make command for", buildType)

		// Ensure schemas are up to date before building
		if err := checkoutSchemas(filePaths.SuppressionCheckout); err != nil {
			return fmt.Errorf("failed to checkout schemas: %w", err)
		}

		// Clear previous error log
		err = os.Truncate(filePaths.RunMakesErrors, 0)
		if err != nil {
			return fmt.Errorf("failed to clear makes_errors.txt: %w", err)
		}
		file.Close()

		// Execute make command with error filtering
		command := fmt.Sprintf("make %s 2>&1 | grep \"error\" > %s", buildType, filePaths.RunMakesErrors)
		if err := execCommand("sh", "-c", command); err != nil {
			fmt.Fprintf(os.Stderr, "Make command failed: %v\nSee makes_output.txt for full output.\n", err)
		}

		// Read and process errors from the error log
		runMakesErrorData, err := os.ReadFile(filePaths.RunMakesErrors)
		if err != nil {
			return fmt.Errorf("failed to read error log file: %w", err)
		}
		err = os.Truncate(filePaths.RunMakesErrors, 0)
		if err != nil {
			return fmt.Errorf("error occurred after reading error log file: %w", err)
		}

		// Process each error line individually
		makesErrors := strings.Split(string(runMakesErrorData), "\n")
		errorCount := 0
		for _, errorLine := range makesErrors {
			if errorLine == "" {
				continue
			}
			errorCount++
			err := processErrorLine(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap)
			if err != nil {
				log.Printf("Error processing line: %v", err)
				return fmt.Errorf("failed to process error line: %w. manual intervention is needed to complete release", err)
			}
		}

		// Continue looping if we found errors to process, otherwise exit
		hasErrors = errorCount > 0
		for idx, errorLine := range makesErrors {
			log.Printf("lines[%d]: %q\n", idx, errorLine)
		}
		print("Processed ", errorCount, " lines from error log file.\n")
		for _, l := range makesErrors {
			log.Println(l)
		}
	}

	return nil
}

// processErrorLine analyzes individual error lines and applies appropriate handling strategies.
// It identifies different error patterns and delegates to specific handler functions.
func processErrorLine(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	if errorLine == "" {
		return nil // Skip empty lines
	}

	log.Printf("Found an entry in the error log: %s during make %s \n", errorLine, buildType)

	// Dispatch to appropriate error handler based on error pattern
	if strings.Contains(errorLine, "StatusCode: 403,") {
		return fmt.Errorf("failed to handle StatusCode 403 error: %s", errorLine)
	} else if strings.Contains(errorLine, "stack overflow") {
		if err := handleStackOverflowError(ctx, errorLine, config, currentSchemas, buildType, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle stack overflow error: %w", err)
		}
	} else if strings.Contains(errorLine, "AWS_") {
		if err := handleAWS_Error(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle AWS_ error: %w", err)
		}
	} else if strings.Contains(errorLine, "AWS::") {
		if err := handleAWSColonError(ctx, errorLine, config, currentSchemas, buildType, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle AWS:: error: %w", err)
		}
	} else if strings.Contains(errorLine, "aws_") {
		if err := handleAWS_UnderscoreError(ctx, errorLine, config, currentSchemas, buildType, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle aws_ error: %w", err)
		}
	} else if strings.Contains(errorLine, "awscc_") {
		if err := handleAWSCC_Error(ctx, errorLine, config, currentSchemas, buildType, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle awscc_ error: %w", err)
		}
	} else if errorLine == "" || strings.TrimSpace(errorLine) == "" {
		return nil
	} else {
		return fmt.Errorf("failed to handle unhandled error: %s", errorLine)
	}

	// For data source builds, regenerate schemas to ensure consistency
	if buildType == BuildTypeSingularDataSources || buildType == BuildTypePluralDataSources {
		err := makeBuild(ctx, config, currentSchemas, BuildTypeSchemas, changes, filePaths, isNewMap)
		if err != nil {
			return fmt.Errorf("failed to run make build for schemas: %w", err)
		}
		log.Println("Successfully ran make build for schemas after processing error line.")
	}

	return nil
}

// handleStackOverflowError processes stack overflow errors by extracting the resource name
// from the last processed resource file and applying suppression.
func handleStackOverflowError(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	log.Println("Detected stack overflow error, attempting to extract resource name from logs.")

	// Extract resource name from the last resource tracking file
	data, err := os.ReadFile(filePaths.LastResource)
	if err != nil {
		log.Printf("Failed to read %s: %v\n", filePaths.LastResource, err)
		return fmt.Errorf("failed to read %s: %w", filePaths.LastResource, err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var resourceName string
	if len(lines) > 0 {
		resourceName = strings.TrimSpace(lines[0])
	} else {
		log.Println("Could not extract resource name from last_resource.txt")
		return fmt.Errorf("could not extract resource name from last_resource.txt")
	}
	if resourceName == "" {
		log.Println("Resource name not found for stack overflow:", resourceName)
		return fmt.Errorf("resource name not found for stack overflow: %s", resourceName)
	}

	// Apply suppression for the problematic resource
	new := isNew(resourceName, isNewMap)
	err = suppress(ctx, resourceName, errorLine, config, new, buildType, filePaths, currentSchemas)
	log.Printf("Suppression result: %v", err)
	return err
}

// handleAWS_Error processes errors related to AWS CloudFormation schema files.
// It extracts the resource name from error lines containing AWS_ prefixed schema paths.
func handleAWS_Error(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	// Parse error line to extract resource name from schema file path
	// Example: "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:"
	errorLineParts := strings.Split(errorLine, " ")
	var resourceName string
	for _, errorLinePart := range errorLineParts {
		if strings.HasPrefix(errorLinePart, "../service/cloudformation/schemas/AWS_") && strings.HasSuffix(errorLinePart, ".json:") {
			trimmed := strings.TrimPrefix(errorLinePart, "../service/cloudformation/schemas/")
			resourceName = strings.TrimSuffix(trimmed, ".json:")
			break
		}
	}
	if resourceName == "" {
		return fmt.Errorf("failed to extract resource name from error line: %s", errorLine)
	}

	new := isNew(resourceName, isNewMap)
	print(resourceName, new, buildType, changes, filePaths, currentSchemas)
	return suppress(ctx, resourceName, errorLine, config, new, buildType, filePaths, currentSchemas)
}

// handleAWSColonError processes errors related to AWS CloudFormation type descriptions.
// It handles cases where CloudFormation types are not found or are no longer available.
func handleAWSColonError(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	// Parse error lines containing AWS:: type references
	/* Example: error loading CloudFormation Resource Provider Schema for aws_datasync_storage_system:
	   describing CloudFormation type: operation error CloudFormation: DescribeType,
	   https response error StatusCode: 404, RequestID: b41adbc2-cb4f-4e06-93c0-b6cb2bbae150,
	   TypeNotFoundException: The type 'AWS::DataSync::StorageSystem' cannot be found. */
	errorParts := strings.Split(errorLine, " ")
	if len(errorParts) < 2 {
		return fmt.Errorf("failed to parse 404 error line: %s", errorLine)
	}
	var resourceName string
	for _, part := range errorParts {
		if strings.Contains(part, "AWS::") {
			resourceName = part
			resourceName = strings.ReplaceAll(resourceName, "::", "_")
			// Trim from start until we find any letter
			for len(resourceName) > 0 && !isLetter(resourceName[0]) {
				resourceName = resourceName[1:]
			}
			// Trim from end until we find any letter
			for len(resourceName) > 0 && !isLetter(resourceName[len(resourceName)-1]) {
				resourceName = resourceName[:len(resourceName)-1]
			}
			break
		}
	}
	if resourceName == "" {
		return fmt.Errorf("failed to extract resource name from 404 error line: %s", errorLine)
	}
	new := isNew(resourceName, isNewMap)
	return suppress(ctx, resourceName, errorLine, config, new, buildType, filePaths, currentSchemas)
}

func handleAWS_UnderscoreError(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	var resourceName string
	/*
		Example error: "error loading CloudFormation Resource Provider Schema for aws_nimblestudio_studio: describing CloudFormation type: operation error CloudFormation: DescribeType, exceeded maximum number of attempts, 3, https response error StatusCode: 400, ..."
	*/
	words := strings.Split(errorLine, " ")
	for _, word := range words {
		if strings.HasPrefix(word, "aws_") {
			// Look for a matching file in internal/service/cloudformation/schemas
			schemasDir := filePaths.CloudFormationSchemasDir
			files, err := os.ReadDir(schemasDir)
			if err != nil {
				return fmt.Errorf("failed to read schemas directory: %w", err)
			}
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
					resourceName = strings.TrimSuffix(file.Name(), ".json")
				}
			}
			if resourceName == "" {
				return fmt.Errorf("resource schema file not found for: %s", word)
			}
		}
	}
	if resourceName == "" {
		return fmt.Errorf("failed to extract resource name from 400 error line: %s", errorLine)
	}
	new := isNew(resourceName, isNewMap)
	return suppress(ctx, resourceName, errorLine, config, new, buildType, filePaths, currentSchemas)
}

func handleAWSCC_Error(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	// Example error: "error loading CloudFormation Resource Provider Schema for awscc_aws
	words := strings.Split(errorLine, " ")
	var foundWord string

	for _, word := range words {
		if strings.HasPrefix(word, "awscc_") {
			foundWord = word
			break
		}
	}

	if foundWord == "" {
		return fmt.Errorf("failed to extract awscc_ prefixed word from error line: %s", errorLine)
	}

	var resourceName string
	// Check if we're dealing with plural or singular data sources
	switch buildType {
	case BuildTypePluralDataSources:
		// For plural data sources
		log.Printf("Processing plural data source: %s", foundWord)
		// Use getPluralResourceNames function
		matchedResource, _, err := getPluralResourceNames(foundWord, isNewMap)
		if err != nil {
			return fmt.Errorf("failed to get plural resource name: %w", err)
		}
		t1, t2, t3, err := naming.ParseTerraformTypeName(matchedResource)
		if err != nil {
			return fmt.Errorf("failed to parse Terraform type name for singular data source: %w", err)
		}
		resourceName = fmt.Sprintf("%s_%s_%s", t1, t2, t3)
	case BuildTypeSingularDataSources:
		// For singular data sources - existing logic
		word := strings.ReplaceAll(foundWord, "awscc_", "aws_")
		t1, t2, t3, err := naming.ParseTerraformTypeName(word)
		if err != nil {
			return fmt.Errorf("failed to parse Terraform type name for singular data source: %w", err)
		}
		resourceName = fmt.Sprintf("%s_%s_%s", t1, t2, t3)
	default:
		// For other build types
		word := strings.ReplaceAll(foundWord, "awscc_", "aws_")
		t1, t2, t3, err := naming.ParseTerraformTypeName(word)
		if err != nil {
			return fmt.Errorf("failed to parse Terraform type name: %w", err)
		}
		resourceName = fmt.Sprintf("%s_%s_%s", t1, t2, t3)
	}

	log.Printf("Word found in error line data source: %s", resourceName)

	if resourceName == "" {
		return fmt.Errorf("failed to extract resource name from error line: %s", errorLine)
	}
	return suppress(ctx, resourceName, errorLine, config, true, buildType, filePaths, currentSchemas)
}

// normalizeNames normalizes CloudFormation and Terraform type names for comparison.
// It removes separators and converts to lowercase to enable case-insensitive matching.
//
// Parameters:
//   - cfTypeName: CloudFormation type name (e.g., "AWS::S3::Bucket")
//   - tfTypeName: Terraform type name (e.g., "aws_s3_bucket")
//
// Returns normalized versions of both names.
func normalizeNames(cfTypeName string, tfTypeName string) (string, string) {
	// Remove all separators and convert to lowercase for comparison
	normalize := func(s string) string {
		s = strings.ReplaceAll(s, ":", "")
		s = strings.ReplaceAll(s, "_", "")
		return strings.ToLower(s)
	}
	return normalize(cfTypeName), normalize(tfTypeName)
}

func suppress(ctx context.Context, cfTypeName, schemaError string, config *GitHubConfig, new bool, buildType string, filePaths *UpdateFilePaths, allSchemas *allschemas.AllSchemas) error {
	// Create a GitHub issue for the schema error

	issueURL, err := createIssue(ctx, cfTypeName, schemaError, config, filePaths.RepositoryLink)
	if err != nil {
		log.Printf("Warning: Failed to create GitHub issue: %v", err)
		issueURL = "" // Use empty string if issue creation fails
	}

	// Add to all_schemas.hcl
	log.Println("Suppressing schema generation for", cfTypeName, "with error:", schemaError, "issue URL:", issueURL)
	if buildType != BuildTypeSchemas || new || strings.Contains(schemaError, "TypeNotFoundException") {
		tfTypeName, err := cfTypeNameToTerraformTypeName(cfTypeName)
		if tfTypeName == "" && cfTypeName != "" {
			err = nil
			tfTypeName = strings.ReplaceAll(cfTypeName, "::", "_")
		}
		if err != nil {
			return fmt.Errorf("failed to convert CloudFormation type name to Terraform type name: %w", err)
		}
		cfTypeName, tfTypeName := normalizeNames(cfTypeName, tfTypeName)
		for i := range allSchemas.Resources {
			resourceCfTypeName, resourceTypeName := normalizeNames(allSchemas.Resources[i].CloudFormationTypeName, allSchemas.Resources[i].ResourceTypeName)

			if tfTypeName == resourceTypeName || cfTypeName == resourceCfTypeName {
				switch buildType {
				case BuildTypeSchemas:
					allSchemas.Resources[i].SuppressResourceGeneration = true
					log.Printf("Suppressing schema generation for %s", allSchemas.Resources[i].CloudFormationTypeName)
				case BuildTypeResources:
					allSchemas.Resources[i].SuppressResourceGeneration = true
					log.Printf("Suppressing resource generation for %s", allSchemas.Resources[i].CloudFormationTypeName)
				case BuildTypeSingularDataSources:
					allSchemas.Resources[i].SuppressSingularDataSourceGeneration = true
					allSchemas.Resources[i].SuppressPluralDataSourceGeneration = true
					log.Printf("Suppressing singular data source generation for %s", allSchemas.Resources[i].CloudFormationTypeName)
				case BuildTypePluralDataSources:
					allSchemas.Resources[i].SuppressPluralDataSourceGeneration = true
					log.Printf("Suppressing plural data source generation for %s", allSchemas.Resources[i].CloudFormationTypeName)
				default:
					if allSchemas.Resources[i].SuppressionReason == "" {
						allSchemas.Resources[i].SuppressionReason = fmt.Sprintf("%s %s", schemaError, issueURL)
					} else {
						allSchemas.Resources[i].SuppressionReason = fmt.Sprintf("%s, %s", allSchemas.Resources[i].SuppressionReason, schemaError)
					}
				}
			}
		}
	} else {
		log.Println("Suppressing for build type:", buildType, "new:", new, "resource:", cfTypeName)
		if !new {
			err := addSchemaToCheckout(cfTypeName, filePaths)
			if err != nil {
				return fmt.Errorf("failed to add resource to checkout file: %w", err)
			}
			return nil
		}
	}

	err = writeSchemasToHCLFile(allSchemas, filePaths.AllSchemasHCL)
	if err != nil {
		return fmt.Errorf("failed to write schemas to HCL file: %w", err)
	}
	return nil
}

// adds a schema file path to the suppression checkout file for git operations after a resource fails during make schemas
func addSchemaToCheckout(resource string, filePaths *UpdateFilePaths) error {
	log.Println("Adding resource to checkout:", resource)

	log.Println("Opening file:", filePaths.SuppressionCheckout)

	file, err := os.OpenFile(filePaths.SuppressionCheckout, os.O_APPEND|os.O_CREATE|os.O_WRONLY, FilePermission)
	if err != nil {
		log.Println("Error opening file:", err)
		return fmt.Errorf("failed to open checkout file for writing: %w", err)
	}
	defer file.Close()

	writeContent := fmt.Sprintf("\n%s/%s.json", filePaths.CloudFormationSchemasDir, resource)
	log.Println("Writing to file:", writeContent)

	_, err = file.WriteString(writeContent)
	if err != nil {
		log.Println("Error writing to file:", err)
		return fmt.Errorf("failed to write to checkout file: %w", err)
	}

	log.Println("Successfully wrote to file")
	return nil
}

// checkoutSchemas reads the suppression data file and checks out each listed schema file using git.
func checkoutSchemas(suppressionData string) error {
	info, err := os.Stat(suppressionData)
	if err != nil {
		if os.IsNotExist(err) {
			// Create an empty suppression data file if it does not exist
			f, createErr := os.Create(suppressionData)
			if createErr != nil {
				return fmt.Errorf("failed to create suppression data file: %w", createErr)
			}
			f.Close()
			return nil
		} else {
			return fmt.Errorf("failed to stat suppression data file: %w", err)
		}
	}
	if info.Size() != 0 {
		contents, readErr := os.ReadFile(suppressionData)
		if readErr != nil {
			return fmt.Errorf("failed to read suppression data file: %w", readErr)
		}

		// Split the contents by spaces to get individual file paths
		lines := strings.Split(string(contents), "\n")
		for _, line := range lines {
			path := strings.TrimSpace(line)
			if path != "" {
				err := execGit("checkout", path)
				if err != nil {
					log.Printf("Failed to checkout %s: %v", path, err)
				}
			}
		}
	}
	return nil
}

func cfTypeNameToTerraformTypeName(cfTypeName string) (string, error) {
	// Convert CloudFormation type name to Terraform type name
	cfTypeName = strings.ReplaceAll(cfTypeName, "_", "::")
	log.Println("Converting CloudFormation type name to Terraform type name:", cfTypeName)
	org, svc, res, err := naming.ParseCloudFormationTypeName(cfTypeName)
	log.Println("Parsed CloudFormation type name:", org, svc, res)
	if err != nil {
		return "", fmt.Errorf("parsing CloudFormation type name (%s): %w", cfTypeName, err)
	}
	tfTypeName := naming.CreateTerraformTypeName(strings.ToLower(org), strings.ToLower(svc), naming.CloudFormationPropertyToTerraformAttribute(res))
	return tfTypeName, nil
}

// isNew checks if a given CloudFormation type name corresponds to a new resource
// this is checked by the absence of the resource in the passed in isNewMap
func isNew(cloudFormationTypeName string, isNewMap map[string]bool) bool {
	log.Println("Checking if resource is new:", cloudFormationTypeName)
	// Convert CloudFormation type name to Terraform type name
	tfTypeName, err := cfTypeNameToTerraformTypeName(cloudFormationTypeName)
	if err != nil {
		log.Printf("Error converting CloudFormation type name to Terraform type name: %v", err)
		return false
	}
	// Check if the resource exists in the map
	if _, exists := isNewMap[tfTypeName]; exists {
		log.Println("Resource found in current schemas:", cloudFormationTypeName)
		return false
	}

	return true
}

// Returns plural resource name by trimming from the right until a match is found in isNewMap
// Returns the modified input with aws_ prefix and original modified input for logging
func getPluralResourceNames(input string, isNewMap map[string]bool) (string, string, error) {
	// awscc_dynamodb_tables
	// awscc_s3_storage_lenses
	// awscc_xray_resource_policies
	// awscc_workspacesweb_network_settings_plural

	// Trim awscc_ prefix and add aws_ prefix
	if !strings.HasPrefix(input, "awscc_") {
		return "", "", fmt.Errorf("input does not start with awscc_: %s", input)
	}

	// Remove awscc_ prefix and add aws_ prefix
	modifiedInput := "aws_" + strings.TrimPrefix(input, "awscc_")

	// Keep the original for return
	originalModified := modifiedInput

	// Gradually trim from right until we find a match in isNewMap
	for len(modifiedInput) > 4 { // At minimum we need "aws_"
		if _, exists := isNewMap[modifiedInput]; exists {
			return modifiedInput, originalModified, nil
		}

		// Find the last underscore
		lastUnderscoreIndex := strings.LastIndex(modifiedInput, "_")
		if lastUnderscoreIndex <= 3 { // If no underscore or only in "aws_" part
			break
		}

		// Trim everything after the last underscore
		modifiedInput = modifiedInput[:lastUnderscoreIndex]
	}

	return "", "", fmt.Errorf("no match found for %s after trimming", input)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
