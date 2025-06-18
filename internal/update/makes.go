package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/google/go-github/v72/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

const (
	BuildTypeSchemas             = "schemas"
	BuildTypeResources           = "resources"
	BuildTypeSingularDataSources = "singular-data-sources"
	BuildTypePluralDataSources   = "plural-data-sources"
	CloudFormationSchemasDir     = "internal/service/cloudformation/schemas"
)

func makeBuild(ctx context.Context, client *github.Client, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths) error {
	if buildType != BuildTypeSchemas && buildType != BuildTypeResources && buildType != BuildTypeSingularDataSources && buildType != BuildTypePluralDataSources {
		return fmt.Errorf("invalid build type: %s, must be '%s', '%s', '%s', or '%s'", buildType, BuildTypeSchemas, BuildTypeResources, BuildTypeSingularDataSources, BuildTypePluralDataSources)
	}

	file, err := os.OpenFile(filePaths.RunMakesErrors, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %w", err)
	}
	defer file.Close()

	var loopCount int = 1

	for i := 0; i < loopCount; i++ {
		log.Println("Running make command for", buildType)
		checkoutSchemas(ctx, filePaths.SuppressionCheckout)

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
		execCommand("sh", "-c", command)

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
			err := processErrorLine(ctx, errorLine, client, currentSchemas, buildType, filePaths)
			if err != nil {
				tflog.Debug(ctx, fmt.Sprintf("Error processing line: %v", err))
			}
		}
		i = 0
		loopCount = len(makesErrors)
		for idx, errorLine := range makesErrors {
			fmt.Printf("lines[%d]: %q\n", idx, errorLine)
		}
		print("Processed ", len(makesErrors), " lines from error log file.\n")
		for _, l := range makesErrors {
			fmt.Println(l)
		}
	}

	return nil
}
func processErrorLine(ctx context.Context, errorLine string, client *github.Client, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths) error {
	if errorLine == "" {
		return nil // Skip empty lines
	}
	var resourceName string
	fmt.Printf("Found an entry in the error log: %s during make %s \n", errorLine, buildType)
	if strings.Contains(errorLine, "stack overflow") {
		fmt.Println("Detected stack overflow error, attempting to extract resource name from logs.")
		var resourceName string = ""
		// Try to extract resource name from stack overflow error using emit_attribute_last_tftype.txt
		lastResourceFile := "internal/provider/last_resource.txt"
		data, err := os.ReadFile(lastResourceFile)
		if err != nil {
			fmt.Printf("Failed to read %s: %v\n", lastResourceFile, err)
			return fmt.Errorf("failed to read %s: %w", lastResourceFile, err)
		}
		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		if len(lines) > 0 {
			resourceName = strings.TrimSpace(lines[0])
		} else {
			fmt.Println("Could not extract resource name from last_resource.txt")
			return fmt.Errorf("could not extract resource name from last_resource.txt")
		}
		if resourceName == "" {
			fmt.Println("Resource name not found for stack overflow:", resourceName)
			return fmt.Errorf("resource name not found for stack overflow: %s", resourceName)
		}
		new := isNew(resourceName, currentSchemas)
		err = suppress(ctx, resourceName, errorLine, client, new, buildType, filePaths, currentSchemas)
		fmt.Print("Suppression result: ", err)
		return err
	} else if strings.Contains(errorLine, "AWS_") {
		// "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:"
		errorLineParts := strings.Split(errorLine, " ")
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
		new := isNew(resourceName, currentSchemas)
		return suppress(ctx, resourceName, errorLine, client, new, buildType, filePaths, currentSchemas)
	} else if strings.Contains(errorLine, "AWS::") {
		// Deleted Resource
		/* error loading CloudFormation Resource Provider Schema for aws_datasync_storage_system: describing CloudFormation type: operation error CloudFormation: DescribeType, https response error StatusCode: 404, RequestID: b41adbc2-cb4f-4e06-93c0-b6cb2bbae150, TypeNotFoundException: The type 'AWS::DataSync::StorageSystem' cannot be found. */
		errorParts := strings.Split(errorLine, " ")
		if len(errorParts) < 2 {
			return fmt.Errorf("failed to parse 404 error line: %s", errorLine)
		}
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
		new := isNew(resourceName, currentSchemas)
		return suppress(ctx, resourceName, errorLine, client, new, buildType, filePaths, currentSchemas)
	} else if strings.Contains(errorLine, "aws_") {
		var resourceName string
		/*
			Example error: "error loading CloudFormation Resource Provider Schema for aws_nimblestudio_studio: describing CloudFormation type: operation error CloudFormation: DescribeType, exceeded maximum number of attempts, 3, https response error StatusCode: 400, ..."
		*/
		words := strings.Split(errorLine, " ")
		for _, word := range words {
			if strings.HasPrefix(word, "aws_") {
				// Look for a matching file in internal/service/cloudformation/schemas
				schemasDir := CloudFormationSchemasDir
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
		new := isNew(resourceName, currentSchemas)
		return suppress(ctx, resourceName, errorLine, client, new, buildType, filePaths, currentSchemas)
	} else if strings.Contains(errorLine, "StatusCode: 403,") {
		return fmt.Errorf("authentication failed: no valid AWS credentials")
	} else {
		return fmt.Errorf("unhandled schema error: %s", errorLine)
	}
}
func suppress(ctx context.Context, cloudFormationTypeName, schemaError string, client *github.Client, new bool, buildType string, filePaths *UpdateFilePaths, allSchemas *allschemas.AllSchemas) error {
	log.Println("Suppressing resource:", cloudFormationTypeName)
	// Create Issue - temporarily commented out to avoid GitHub API calls
	// issueURL, err := createIssue(resource, schemaError, client)
	// if err != nil {
	//     return fmt.Errorf("failed to create GitHub issue: %w", err)
	// }

	// Use empty issue URL instead
	issueURL := ""
	var err error
	// Add to all_schemas.hcl
	if buildType != BuildTypeSchemas {
		allSchemas, err = parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
		if err != nil {
			return fmt.Errorf("failed to parse all_schemas.hcl: %w", err)
		}
		terraformResourceName := strings.ToLower(cloudFormationTypeName)
		for i := range allSchemas.Resources {
			if terraformResourceName == allSchemas.Resources[i].CloudFormationTypeName {
				switch buildType {
				case BuildTypeResources:
					allSchemas.Resources[i].SuppressResourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing resource generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
				case BuildTypeSingularDataSources:
					allSchemas.Resources[i].SuppressSingularDataSourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing singular data source generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
				case "plural-data-source":
					allSchemas.Resources[i].SuppressPluralDataSourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing plural data source generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
				default:
					if allSchemas.Resources[i].SuppressionReason == "" {
						allSchemas.Resources[i].SuppressionReason = fmt.Sprintf("%s %s", schemaError, issueURL)
					} else {
						allSchemas.Resources[i].SuppressionReason = fmt.Sprintf("%s, %s", allSchemas.Resources[i].SuppressionReason, schemaError)
					}
				}
			}
		}
		err = writeSchemasToHCLFile(allSchemas, "internal/provider/all_schemas.hcl")
		if err != nil {
			return fmt.Errorf("failed to write schemas to HCL file: %w", err)
		}
	} else {
		log.Println("Skipping suppression for schemas mode")
		tflog.Debug(ctx, fmt.Sprintf("Skipping suppression for %s in mode %s", cloudFormationTypeName, buildType))
		if !new {
			err := addSchemaToCheckout(cloudFormationTypeName, filePaths)
			if err != nil {
				return fmt.Errorf("failed to add resource to checkout file: %w", err)
			}
			return nil
		} else { // This effectively deletes the resource from the provider
			temp := &allschemas.ResourceAllSchema{
				ResourceTypeName:           strings.ToLower(cloudFormationTypeName),
				CloudFormationTypeName:     cloudFormationTypeName,
				SuppressResourceGeneration: true,
			}
			allSchemas.Resources = append(allSchemas.Resources, *temp)
			sort.Slice(allSchemas.Resources, func(i, j int) bool {
				return allSchemas.Resources[i].ResourceTypeName < allSchemas.Resources[j].ResourceTypeName
			})
			tflog.Debug(ctx, fmt.Sprintf("Suppressing resource generation for %s in all_schemas.hcl", cloudFormationTypeName))
			err := writeSchemasToHCLFile(allSchemas, "internal/provider/all_schemas.hcl")
			if err != nil {
				return fmt.Errorf("failed to write schemas to HCL file: %w", err)
			}
			return nil

		}
	}

	return nil
}
func addSchemaToCheckout(resource string, filePaths *UpdateFilePaths) error {
	log.Println("Adding resource to checkout:", resource)

	log.Println("Opening file:", filePaths.SuppressionCheckout)

	file, err := os.OpenFile(filePaths.SuppressionCheckout, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return fmt.Errorf("failed to open checkout file for writing: %w", err)
	}
	defer file.Close()

	writeContent := fmt.Sprintf("%s/%s.json \n", CloudFormationSchemasDir, resource)
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
					tflog.Debug(ctx, fmt.Sprintf("Failed to checkout %s: %v", path, err))
				}
			}
		}
	}
	return nil
}

func GetResourceFromLog(filePaths *UpdateFilePaths, buildType string) (string, error) {
	var resourceName string
	logData, err := os.ReadFile(filePaths.RunMakesResourceLog)
	fmt.Println("Reading log file:", filePaths.RunMakesResourceLog)
	if err != nil {
		return "", fmt.Errorf("failed to read logs file: %w", err)
	}
	logLines := strings.Split(strings.TrimSpace(string(logData)), "\n")
	if len(logLines) > 0 {
		fmt.Println("Log lines found:", logLines)
		resourceName = logLines[len(logLines)-1]
	} else {
		return "", fmt.Errorf("no resource name found in logs")
	}
	return resourceName, nil
}

func isNew(cloudFormationTypeName string, currentSchemas *allschemas.AllSchemas) bool {
	cloudFormationTypeName = strings.ReplaceAll(cloudFormationTypeName, "_", "::")
	for _, r := range currentSchemas.Resources {
		if r.CloudFormationTypeName == cloudFormationTypeName {
			return false
		}
	}
	return true
}

func createIssue(resource, schemaError string, client *github.Client) (string, error) {
	// testing
	return "", nil
	ctx := context.Background()
	repoOwner := "hashicorp"
	repoName := "terraform-provider-awscc"

	issueTitle := fmt.Sprintf("Resource Suppression: %s", resource)
	issueBody := fmt.Sprintf(`
	<!--- Please keep this note for the community --->

	### Community Note

	* Please vote on this issue by adding a 👍 [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/) to the original issue to help the community and maintainers prioritize this request
	* Please do not leave "+1" or other comments that do not add relevant new information or questions, they generate extra noise for issue followers and do not help prioritize the request
	* If you are interested in working on this issue or have submitted a pull request, please leave a comment
	* The resources and data sources in this provider are generated from the CloudFormation schema, so they can only support the actions that the underlying schema supports. For this reason submitted bugs should be limited to defects in the generation and runtime code of the provider. Customizing behavior of the resource, or noting a gap in behavior are not valid bugs and should be submitted as enhancements to AWS via the CloudFormation Open Coverage Roadmap.

	<!--- Thank you for keeping this note for the community --->

	### Description

	<!--- Please leave a helpful description of the feature request here. --->

	%s

	### Affected Resource(s)

	<!--- Please list the new or affected resources and data sources. --->

	* %s
	
	### Schema Definition

	https://github.com/hashicorp/terraform-provider-awscc/blob/d1f668deabc299d8ef5c8bdfe50bfa9cb98bbeee/internal/service/cloudformation/schemas/AWS_IoTFleetWise_DecoderManifest.json#L378-L393

	### References

	- #156 `, resource, schemaError)

	issueRequest := &github.IssueRequest{
		Title: &issueTitle,
		Body:  &issueBody,
	}

	issue, _, err := client.Issues.Create(
		ctx,
		repoOwner,
		repoName,
		issueRequest,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create GitHub issue: %w", err)
	}
	_, _, err = client.Issues.AddLabelsToIssue(
		ctx,
		repoOwner,
		repoName,
		*issue.Number,
		[]string{"bug", "resource-suppression"},
	)

	if err != nil {
		return "", fmt.Errorf("failed to add labels to GitHub issue: %w", err)
	}
	return *issue.URL, nil
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
