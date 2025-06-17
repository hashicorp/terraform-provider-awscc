package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/google/go-github/v72/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func makeBuild(ctx context.Context, client *github.Client, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths) error {
	if buildType != "schemas" && buildType != "resources" && buildType != "singular-data-sources" && buildType != "plural-data-sources" {
		return fmt.Errorf("invalid build type: %s, must be 'schemas', 'resources', 'singular-data-sources', or 'plural-data-sources'", buildType)
	}

	file, err := os.OpenFile(filePaths.RunMakesErrors, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %w", err)
	}
	defer file.Close()
	awsServices, err := listResourceTypes(ctx)
	if err != nil {
		return fmt.Errorf("failed to list AWS resource types: %w", err)
	}

	var loopCount int = 1

	for i := 0; i < loopCount; i++ {
		log.Println("Running make command for", buildType)
		checkoutSchemas(ctx, filePaths.SuppressionCheckout)

		err = os.Truncate(filePaths.RunMakesErrors, 0)
		if err != nil {
			return fmt.Errorf("failed to clear makes_errors.txt: %w", err)
		}
		file.Close()
		time.Sleep(2 * time.Second) // Ensure file is closed before reopening in command
		command := fmt.Sprintf("make %s 2>&1 | tee %s | grep \"error\" > %s", buildType, filePaths.RunMakesOutput, filePaths.RunMakesErrors)
		execCommand("sh", "-c", command)
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

		lines := strings.Split(string(runMakesErrorData), "\n")
		for _, line := range lines {
			if err := processErrorLine(ctx, line, client, currentSchemas, buildType, filePaths, awsServices); err != nil {
				tflog.Debug(ctx, fmt.Sprintf("Error processing line: %v", err))
			}
		}
		i = 0
		loopCount = len(lines)
		print("Processed ", len(lines), " lines from error log file.\n")
		runMakesErrorFile, err := os.OpenFile(filePaths.RunMakesErrors, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to reopen error log file: %w", err)
		}
		err = runMakesErrorFile.Truncate(0)
		if err != nil {
			return fmt.Errorf("failed to clear error log file: %w", err)
		}
		runMakesErrorFile.Close()
	}

	return nil
}
func processErrorLine(ctx context.Context, line string, client *github.Client, currentSchemas *allschemas.AllSchemas, buildType string, filePaths *UpdateFilePaths, awsServices map[string]string) error {
	if line == "" {
		return nil // Skip empty lines
	}
	fmt.Printf("Found an entry in the error log: %s during make %s \n", line, buildType)
	if strings.Contains(line, "stack overflow") {
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
		err = suppress(ctx, resourceName, line, client, new, buildType, filePaths)
		fmt.Print("Suppression result: ", err)
		return err
	} else if strings.Contains(line, "AWS_") {
		var resourceName string = ""
		// Example error:
		// "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:"
		words := strings.Split(line, " ")
		for _, word := range words {
			if strings.HasPrefix(word, "../service/cloudformation/schemas/AWS_") && strings.HasSuffix(word, ".json:") {
				trimmed := strings.TrimPrefix(word, "../service/cloudformation/schemas/")
				resourceName = strings.TrimSuffix(trimmed, ".json:")
				break
			}
		}
		if resourceName == "" {
			return fmt.Errorf("failed to extract resource name from error line: %s", line)
		}
		new := isNew(resourceName, currentSchemas)
		return suppress(ctx, resourceName, line, client, new, buildType, filePaths)
	} else if strings.Contains(line, "AWS::") {
		// Deleted Resource
		/* error loading CloudFormation Resource Provider Schema for aws_datasync_storage_system: describing CloudFormation type: operation error CloudFormation: DescribeType, https response error StatusCode: 404, RequestID: b41adbc2-cb4f-4e06-93c0-b6cb2bbae150, TypeNotFoundException: The type 'AWS::DataSync::StorageSystem' cannot be found. */

		errorParts := strings.Split(line, " ")
		if len(errorParts) < 2 {
			return fmt.Errorf("failed to parse 404 error line: %s", line)
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
			return fmt.Errorf("failed to extract resource name from 404 error line: %s", line)
		}
		new := isNew(resourceName, currentSchemas)
		return suppress(ctx, resourceName, line, client, new, buildType, filePaths)
	} else if strings.Contains(line, "aws_") {
		var resourceName string
		/*
			Example error: "error loading CloudFormation Resource Provider Schema for aws_nimblestudio_studio: describing CloudFormation type: operation error CloudFormation: DescribeType, exceeded maximum number of attempts, 3, https response error StatusCode: 400, ..."
		*/
		words := strings.Split(line, " ")
		for _, word := range words {
			if strings.HasPrefix(word, "aws_") {
				// Look for a matching file in internal/service/cloudformation/schemas
				schemasDir := "internal/service/cloudformation/schemas"
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
			return fmt.Errorf("failed to extract resource name from 400 error line: %s", line)
		}
		new := isNew(resourceName, currentSchemas)
		return suppress(ctx, resourceName, line, client, new, buildType, filePaths)
	} else if strings.Contains(line, "StatusCode: 403,") {
		return fmt.Errorf("authentication failed: no valid AWS credentials")
	} else {
		return fmt.Errorf("unhandled schema error: %s", line)
	}
}
func suppress(ctx context.Context, resource, schemaError string, client *github.Client, new bool, mode string, filePaths *UpdateFilePaths) error {
	log.Println("Suppressing resource:", resource)
	// Create Issue - temporarily commented out to avoid GitHub API calls
	// issueURL, err := createIssue(resource, schemaError, client)
	// if err != nil {
	//     return fmt.Errorf("failed to create GitHub issue: %w", err)
	// }

	// Use empty issue URL instead
	issueURL := ""
	// Add to all_schemas.hcl
	if mode != "schemas" {
		allSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
		// we can probably optimize this but dont premature optimize
		if err != nil {
			return fmt.Errorf("failed to parse all_schemas.hcl: %w", err)
		}
		terraformResourceName := strings.ToLower(resource)
		for i := range allSchemas.Resources {
			if terraformResourceName == allSchemas.Resources[i].CloudFormationTypeName {
				switch mode {
				case "schemas": // This deletes the resource
					allSchemas.Resources[i].SuppressResourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing resource generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
				case "resources":
					allSchemas.Resources[i].SuppressResourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing resource generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
				case "singular-data-sources":
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
		tflog.Debug(ctx, fmt.Sprintf("Skipping suppression for %s in mode %s", resource, mode))
		err := addSchemaToCheckout(resource, filePaths)
		if err != nil {
			return fmt.Errorf("failed to add resource to checkout file: %w", err)
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
		file.Close()
		return fmt.Errorf("failed to open checkout file for writing: %w", err)
	}
	defer file.Close()

	writeContent := fmt.Sprintf("internal/service/cloudformation/schemas/%s.json \n", resource)
	log.Println("Writing to file:", writeContent)

	_, err = file.WriteString(writeContent)
	if err != nil {
		log.Println("Error writing to file:", err)
		return fmt.Errorf("failed to write to checkout file: %w", err)
	}

	log.Println("Successfully wrote to file")
	time.Sleep(1 * time.Second)
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

func isNew(resourceName string, currentSchemas *allschemas.AllSchemas) bool {
	resourceName = strings.ToLower(resourceName)
	for _, r := range currentSchemas.Resources {
		if r.ResourceTypeName == resourceName {
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

// Get config with specific region
func getConfigWithRegion(ctx context.Context, region string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

// List CloudFormation resource types using SDK v2
func listResourceTypes(ctx context.Context) (map[string]string, error) {
	// Get AWS configuration
	cfg, err := getConfigWithRegion(ctx, "us-east-1")
	if err != nil {
		return nil, err
	}

	// Create CloudFormation client
	client := cloudformation.NewFromConfig(cfg)

	// Create the input for ListTypes
	input := &cloudformation.ListTypesInput{
		Type:       "RESOURCE",
		Visibility: "PUBLIC",
	}

	resourceTypes := make(map[string]string)
	paginator := cloudformation.NewListTypesPaginator(client, input)

	// Paginate through results
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, summary := range page.TypeSummaries {
			if summary.TypeName != nil {
				tfccString := aws.ToString(summary.TypeName)

				tfccString = strings.ToLower(tfccString)
				tfccString = strings.ReplaceAll(tfccString, "::", "_")

				aws_string := aws.ToString(summary.TypeName)
				aws_string = strings.ReplaceAll(aws_string, "::", "_")

				resourceTypes[tfccString] = aws_string
			}
		}
	}

	return resourceTypes, nil
}
