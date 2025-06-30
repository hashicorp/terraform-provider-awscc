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

const (
	BuildTypeSchemas             = "schemas"
	BuildTypeResources           = "resources"
	BuildTypeSingularDataSources = "singular-data-sources"
	BuildTypePluralDataSources   = "plural-data-sources"
)

func makeBuild(ctx context.Context, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	if buildType != BuildTypeSchemas && buildType != BuildTypeResources && buildType != BuildTypeSingularDataSources && buildType != BuildTypePluralDataSources {
		return fmt.Errorf("invalid build type: %s, must be '%s', '%s', '%s', or '%s'", buildType, BuildTypeSchemas, BuildTypeResources, BuildTypeSingularDataSources, BuildTypePluralDataSources)
	}

	file, err := os.OpenFile(filePaths.RunMakesErrors, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %w", err)
	}
	defer file.Close()

	var loopCount = 1

	for i := 0; i < loopCount; i++ {
		log.Println("Running make command for", buildType)
		if err := checkoutSchemas(ctx, filePaths.SuppressionCheckout); err != nil {
			return fmt.Errorf("failed to checkout schemas: %w", err)
		}

		err = os.Truncate(filePaths.RunMakesErrors, 0)
		if err != nil {
			return fmt.Errorf("failed to clear makes_errors.txt: %w", err)
		}
		file.Close()

		// Generate a random filename for the tee output
		tmpFile, err := os.CreateTemp("", "makes_output_*.txt")
		if err != nil {
			return fmt.Errorf("failed to create temporary output file: %w", err)
		}
		tmpFileName := tmpFile.Name()
		tmpFile.Close()

		command := fmt.Sprintf("make %s 2>&1 | tee %s | grep \"error\" > %s", buildType, tmpFileName, filePaths.RunMakesErrors)
		if err := execCommand("sh", "-c", command); err != nil {
			fmt.Fprintf(os.Stderr, "Make command failed: %v\nSee makes_output.txt for full output.\n", err)
		}

		// Optionally, you can clean up the temp file after use if needed:
		// defer os.Remove(tmpFileName)
		/* I think the grep is erroring out so no returns
		if err != nil {
			fmt.Fprintf(os.Stderr, "Make command failed: %v\nSee makes_output.txt for full output.\n", err)
			return fmt.Errorf("failed to execute make %s command: %w", buildType, err)
		}
		*/
		runMakesErrorData, err := os.ReadFile(filePaths.RunMakesErrors)
		if err != nil {
			return fmt.Errorf("failed to read error log file: %w", err)
		}
		err = os.Truncate(filePaths.RunMakesErrors, 0)
		if err != nil {
			return fmt.Errorf("error occurred after reading error log file: %w", err)
		}

		makesErrors := strings.Split(string(runMakesErrorData), "\n")
		for _, errorLine := range makesErrors {
			err := processErrorLine(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap)
			if err != nil {
				log.Printf("Error processing line: %v", err)
			}
		}
		i = 0
		loopCount = len(makesErrors)
		for idx, errorLine := range makesErrors {
			fmt.Printf("lines[%d]: %q\n", idx, errorLine)
		}
		print("Processed ", len(makesErrors), " lines from error log file.\n")
		for _, l := range makesErrors {
			log.Println(l)
		}
	}

	return nil
}
func processErrorLine(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	if errorLine == "" {
		return nil // Skip empty lines
	}

	fmt.Printf("Found an entry in the error log: %s during make %s \n", errorLine, buildType)

	// Check for different error patterns and handle them
	if strings.Contains(errorLine, "stack overflow") {
		if err := handleStackOverflowError(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle stack overflow error: %w", err)
		}
	} else if strings.Contains(errorLine, "AWS_") {
		if err := handleAWS_Error(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle AWS_ error: %w", err)
		}
	} else if strings.Contains(errorLine, "AWS::") {
		if err := handleAWSColonError(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle AWS:: error: %w", err)
		}
	} else if strings.Contains(errorLine, "aws_") {
		if err := handleAWS_UnderscoreError(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle aws_ error: %w", err)
		}
	} else if strings.Contains(errorLine, "awscc_") {
		if err := handleAWSCC_Error(ctx, errorLine, config, currentSchemas, buildType, changes, filePaths, isNewMap); err != nil {
			return fmt.Errorf("failed to handle awscc_ error: %w", err)
		}
	} else if strings.Contains(errorLine, "StatusCode: 403,") {
		if err := handleStatusCode403Error(errorLine); err != nil {
			return fmt.Errorf("failed to handle StatusCode 403 error: %w", err)
		}
	} else {
		if err := handleUnhandledError(errorLine); err != nil {
			return fmt.Errorf("failed to handle unhandled error: %w", err)
		}
	}

	if buildType == BuildTypeSingularDataSources || buildType == BuildTypePluralDataSources {
		err := makeBuild(ctx, config, currentSchemas, BuildTypeSchemas, changes, filePaths, isNewMap)
		if err != nil {
			return fmt.Errorf("failed to run make build for schemas: %w", err)
		}
		log.Println("Successfully ran make build for schemas after processing error line.")
	}

	return nil
}

func handleStackOverflowError(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	log.Println("Detected stack overflow error, attempting to extract resource name from logs.")
	// Try to extract resource name from stack overflow error using emit_attribute_last_tftype.txt
	data, err := os.ReadFile(filePaths.LastResource)
	if err != nil {
		fmt.Printf("Failed to read %s: %v\n", filePaths.LastResource, err)
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
	new := isNew(resourceName, isNewMap)
	err = suppress(ctx, resourceName, errorLine, config, new, buildType, changes, filePaths, currentSchemas)
	fmt.Print("Suppression result: ", err)
	return err
}

func handleAWS_Error(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	// "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:"
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
	return suppress(ctx, resourceName, errorLine, config, new, buildType, changes, filePaths, currentSchemas)
}

func handleAWSColonError(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
	// Deleted Resource
	/* error loading CloudFormation Resource Provider Schema for aws_datasync_storage_system: describing CloudFormation type: operation error CloudFormation: DescribeType, https response error StatusCode: 404, RequestID: b41adbc2-cb4f-4e06-93c0-b6cb2bbae150, TypeNotFoundException: The type 'AWS::DataSync::StorageSystem' cannot be found. */
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
	return suppress(ctx, resourceName, errorLine, config, new, buildType, changes, filePaths, currentSchemas)
}

func handleAWS_UnderscoreError(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
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
	return suppress(ctx, resourceName, errorLine, config, new, buildType, changes, filePaths, currentSchemas)
}

func handleAWSCC_Error(ctx context.Context, errorLine string, config *GitHubConfig, currentSchemas *allschemas.AllSchemas, buildType string, changes *[]string, filePaths *UpdateFilePaths, isNewMap map[string]bool) error {
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
		fmt.Println("Processing plural data source:", foundWord)
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

	fmt.Println("Word found in error line data source:", resourceName)

	if resourceName == "" {
		return fmt.Errorf("failed to extract resource name from error line: %s", errorLine)
	}
	return suppress(ctx, resourceName, errorLine, config, true, buildType, changes, filePaths, currentSchemas)
}

func handleStatusCode403Error(errorLine string) error {
	return fmt.Errorf("authentication failed: no valid AWS credentials")
}

func handleUnhandledError(errorLine string) error {
	return fmt.Errorf("unhandled schema error: %s", errorLine)
}

func normalizeNames(cfTypeName string, tfTypeName string) (string, string) {
	// Remove all ':' and '_' and concatenate the parts, then lowercase
	normalize := func(s string) string {
		s = strings.ReplaceAll(s, ":", "")
		s = strings.ReplaceAll(s, "_", "")
		return strings.ToLower(s)
	}
	return normalize(cfTypeName), normalize(tfTypeName)
}

func suppress(ctx context.Context, cfTypeName, schemaError string, config *GitHubConfig, new bool, buildType string, changes *[]string, filePaths *UpdateFilePaths, allSchemas *allschemas.AllSchemas) error {

	log.Println("Suppressing resource:", cfTypeName)
	// Create Issue - temporarily commented out to avoid GitHub API calls
	// issueURL, err := createIssue(resource, schemaError, client)
	// if err != nil {
	//     return fmt.Errorf("failed to create GitHub issue: %w", err)
	// }

	// Record this resource change with the appropriate type
	var reason string
	switch buildType {
	case BuildTypeResources:
		reason = "New Resource Suppression"
	case BuildTypeSingularDataSources:
		reason = "New Singular Data Source Suppression"
	case BuildTypePluralDataSources:
		reason = "New Plural Data Source Suppression"
	case BuildTypeSchemas:
		if new {
			reason = "Suppressed Resource"
		}
	}

	// Store the change data as a string for later use
	*changes = append(*changes, fmt.Sprintf("%s - %s", cfTypeName, reason))

	// Use empty issue URL instead
	issueURL := ""
	var err error
	// Add to all_schemas.hcl
	if buildType != BuildTypeSchemas || new {
		tfTypeName, err := cfTypeNameToTerraformTypeName(cfTypeName)
		fmt.Println("Converting CloudFormation type name to Terraform type name:", cfTypeName, "->", tfTypeName)
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
		/*
			else { // This effectively deletes the resource from the provider
				temp := &allschemas.ResourceAllSchema{
					ResourceTypeName:           strings.ToLower(cfTypeName),
					CloudFormationTypeName:     cfTypeName,
					SuppressResourceGeneration: true,
				}
				allSchemas.Resources = append(allSchemas.Resources, *temp)
				sort.Slice(allSchemas.Resources, func(i, j int) bool {
					return allSchemas.Resources[i].ResourceTypeName < allSchemas.Resources[j].ResourceTypeName
				})
				fmt.Printf("Suppressing resource generation for %s in all_schemas.hcl\n", cfTypeName)
		*/

	}

	err = writeSchemasToHCLFile(allSchemas, filePaths.AllSchemasHCL)
	if err != nil {
		return fmt.Errorf("failed to write schemas to HCL file: %w", err)
	}
	return nil
}
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
func checkoutSchemas(ctx context.Context, suppressionData string) error {
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

func GetResourceFromLog(filePaths *UpdateFilePaths) (string, error) {
	var resourceName string
	logData, err := os.ReadFile(filePaths.RunMakesResourceLog)
	log.Println("Reading log file:", filePaths.RunMakesResourceLog)
	if err != nil {
		return "", fmt.Errorf("failed to read logs file: %w", err)
	}
	logLines := strings.Split(strings.TrimSpace(string(logData)), "\n")
	if len(logLines) > 0 {
		log.Println("Log lines found:", logLines)
		resourceName = logLines[len(logLines)-1]
	} else {
		return "", fmt.Errorf("no resource name found in logs")
	}
	return resourceName, nil
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
