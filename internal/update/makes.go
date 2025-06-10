package update

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v72/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func makeSchemas(ctx context.Context, client *github.Client, currentSchemas allschemas.AllSchemas) error {
	/*
		var waitTime int = 0
		if waitTime > 300 {
			return fmt.Errorf("wait time exceeded maximum of 300 seconds")
		}
	*/

	file, err := os.OpenFile("makes_errors.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("temp")
	if err != nil {
		return fmt.Errorf("failed to write to error log file: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %w", err)
	}
	defer file.Close()
	for stat.Size() > 0 {
		err = execCommand("sh", "-c", `make schemas | grep "error" > makes_errors.txt`)
		if err != nil {
			return fmt.Errorf("failed to execute make schemas command: %w", err)
		}
		data, err := os.ReadFile("makes_errors.txt")
		if err != nil {
			return fmt.Errorf("failed to read error log file: %w", err)
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line == "fatal error: stack overflow" { // Recursive Data Type Stack Overflow
				resourceName, err := GetResourceFromLog()
				if err != nil {
					return fmt.Errorf("failed to get resource from log: %w", err)
				}
				resourceName = strings.ReplaceAll(resourceName, "::", "_")
				new := isNew(strings.ToLower(resourceName), currentSchemas)

				suppress(ctx, resourceName, line, client, new, "schemas")
			} else if strings.Contains(line, "AWS_") {
				fmt.Printf("Found a resource in the error log: %s\n", line)
				var resourceName string = ""
				// Extract resource name from error message like:
				// "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:"
				words := strings.Split(line, " ")
				for _, word := range words {
					if strings.HasPrefix(word, "../service/cloudformation/schemas/AWS_") && strings.HasSuffix(word, ".json:") {
						resourceName = strings.ToLower(strings.TrimSuffix(strings.TrimPrefix(word, "../service/cloudformation/schemas/"), ".json:"))
						break
					}

				}
				if resourceName == "" {
					return fmt.Errorf("failed to extract resource name from error line: %s", line)
				}
				new := isNew(resourceName, currentSchemas)
				suppress(ctx, resourceName, line, client, new, "schemas")
			} else if strings.Contains(line, "https response error StatusCode: 400") { // Deleted resource with spun down service (?)
				var lowerCaseResourceName string
				var resourceName string

				// Example error: "error loading CloudFormation Resource Provider Schema for aws_nimblestudio_studio: describing CloudFormation type: operation error CloudFormation: DescribeType, exceeded maximum number of attempts, 3, https response error StatusCode: 400, ..."
				words := strings.Split(line, " ")
				for _, word := range words {
					if strings.HasPrefix(word, "aws_") && strings.HasSuffix(word, ":") {
						lowerCaseResourceName = strings.TrimSuffix(word, ":")
						// Look for a matching file in internal/service/cloudformation/schemas
						schemasDir := "internal/service/cloudformation/schemas"
						files, err := os.ReadDir(schemasDir)
						if err != nil {
							return fmt.Errorf("failed to read schemas directory: %w", err)
						}
						for _, file := range files {
							if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
								name := strings.TrimSuffix(file.Name(), ".json")
								if name == resourceName {
									resourceName = name
								}
							}
						}
						if resourceName == "" {
							return fmt.Errorf("resource schema file not found for: %s", lowerCaseResourceName)
						}
					}
				}
				if resourceName == "" {
					return fmt.Errorf("failed to extract resource name from 400 error line: %s", line)
				}
				new := isNew(resourceName, currentSchemas)
				suppress(ctx, resourceName, line, client, new, "schemas")
			} else if strings.Contains(line, "https response error StatusCode: 404,") { // Deleted Resource
				/* error loading CloudFormation Resource Provider Schema for aws_datasync_storage_system: describing CloudFormation type: operation error CloudFormation: DescribeType, https response error StatusCode: 404, RequestID: b41adbc2-cb4f-4e06-93c0-b6cb2bbae150, TypeNotFoundException: The type 'AWS::DataSync::StorageSystem' cannot be found. */

				errorParts := strings.Split(line, " ")
				if len(errorParts) < 2 {
					return fmt.Errorf("failed to parse 404 error line: %s", line)
				}
				var resourceName string
				for _, part := range errorParts {
					if strings.HasPrefix(part, "AWS:") && strings.HasSuffix(part, ":") {
						resourceName = strings.ReplaceAll(resourceName, "::", "_")
						break
					}
				}
				if resourceName == "" {
					return fmt.Errorf("failed to extract resource name from 404 error line: %s", line)
				}
				new := isNew(resourceName, currentSchemas)
				suppress(ctx, resourceName, line, client, new, "schemas")
			} else if strings.Contains(line, "StatusCode: 403,") {
				return fmt.Errorf("authentication failed: no valid AWS credentials")
			} else {
				return fmt.Errorf("unhandled schema error: %s", line)
			}
		}
		stat, err = file.Stat()
		if err != nil {
			return fmt.Errorf("failed to get file stats: %w", err)
		}
	}

	return nil
}

func makeResources(ctx context.Context, client *github.Client, currentSchemas allschemas.AllSchemas) error {
	/*
		var waitTime int = 0
		if waitTime > 300 {
			return fmt.Errorf("wait time exceeded maximum of 300 seconds")
		}
	*/

	file, err := os.OpenFile("makes_errors.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("temp")
	if err != nil {
		return fmt.Errorf("failed to write to error log file: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %w", err)
	}
	defer file.Close()
	for stat.Size() > 0 {
		err = execCommand("sh", "-c", `make resources | grep "error" > makes_errors.txt`)
		if err != nil {
			return fmt.Errorf("failed to execute make schemas command: %w", err)
		}
		data, err := os.ReadFile("makes_errors.txt")
		if err != nil {
			return fmt.Errorf("failed to read error log file: %w", err)
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line == "fatal error: stack overflow" { // Recursive Data Type Stack Overflow
				resourceName, err := GetResourceFromLog()
				if err != nil {
					return fmt.Errorf("failed to get resource from log: %w", err)
				}
				resourceName = strings.ReplaceAll(resourceName, "::", "_")
				new := isNew(strings.ToLower(resourceName), currentSchemas)

				suppress(ctx, resourceName, line, client, new, "resources")
			} else if strings.Contains(line, "AWS_") {
				fmt.Printf("Found a resource in the error log: %s\n", line)
				var resourceName string = ""
				// Extract resource name from error message like:
				// "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:"
				words := strings.Split(line, " ")
				for _, word := range words {
					if strings.HasPrefix(word, "../service/cloudformation/schemas/AWS_") && strings.HasSuffix(word, ".json:") {
						resourceName = strings.ToLower(strings.TrimSuffix(strings.TrimPrefix(word, "../service/cloudformation/schemas/"), ".json:"))
						break
					}
				}
				if resourceName == "" {
					return fmt.Errorf("failed to extract resource name from error line: %s", line)
				}
				new := isNew(resourceName, currentSchemas)
				suppress(ctx, resourceName, line, client, new, "resources")
			} else if strings.Contains(line, "https response error StatusCode: 400") { // Deleted resource with spun down service (?)
				var lowerCaseResourceName string
				var resourceName string

				// Example error: "error loading CloudFormation Resource Provider Schema for aws_nimblestudio_studio: describing CloudFormation type: operation error CloudFormation: DescribeType, exceeded maximum number of attempts, 3, https response error StatusCode: 400, ..."
				words := strings.Split(line, " ")
				for _, word := range words {
					if strings.HasPrefix(word, "aws_") && strings.HasSuffix(word, ":") {
						lowerCaseResourceName = strings.TrimSuffix(word, ":")
						// Look for a matching file in internal/service/cloudformation/schemas
						schemasDir := "internal/service/cloudformation/schemas"
						files, err := os.ReadDir(schemasDir)
						if err != nil {
							return fmt.Errorf("failed to read schemas directory: %w", err)
						}
						for _, file := range files {
							if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
								name := strings.TrimSuffix(file.Name(), ".json")
								if name == resourceName {
									resourceName = name
								}
							}
						}
						if resourceName == "" {
							return fmt.Errorf("resource schema file not found for: %s", lowerCaseResourceName)
						}
					}
				}
				if resourceName == "" {
					return fmt.Errorf("failed to extract resource name from 400 error line: %s", line)
				}
				new := isNew(resourceName, currentSchemas)
				suppress(ctx, resourceName, line, client, new, "resources")
			} else if strings.Contains(line, "https response error StatusCode: 404,") { // Deleted Resource
				/* error loading CloudFormation Resource Provider Schema for aws_datasync_storage_system: describing CloudFormation type: operation error CloudFormation: DescribeType, https response error StatusCode: 404, RequestID: b41adbc2-cb4f-4e06-93c0-b6cb2bbae150, TypeNotFoundException: The type 'AWS::DataSync::StorageSystem' cannot be found. */

				errorParts := strings.Split(line, " ")
				if len(errorParts) < 2 {
					return fmt.Errorf("failed to parse 404 error line: %s", line)
				}
				var resourceName string
				for _, part := range errorParts {
					if strings.HasPrefix(part, "AWS:") && strings.HasSuffix(part, ":") {
						resourceName = strings.ReplaceAll(resourceName, "::", "_")
						break
					}
				}
				if resourceName == "" {
					return fmt.Errorf("failed to extract resource name from 404 error line: %s", line)
				}
				new := isNew(resourceName, currentSchemas)
				suppress(ctx, resourceName, line, client, new, "resources")
			} else if strings.Contains(line, "StatusCode: 403,") {
				return fmt.Errorf("authentication failed: no valid AWS credentials")
			} else {
				return fmt.Errorf("unhandled schema error: %s", line)
			}
		}
		stat, err = file.Stat()
		if err != nil {
			return fmt.Errorf("failed to get file stats: %w", err)
		}
	}

	return nil
}

func suppress(ctx context.Context, resource, schemaError string, client *github.Client, new bool, mode string) error {
	// Create Issue
	issueURL, err := createIssue(resource, schemaError, client)
	if err != nil {
		return fmt.Errorf("failed to create GitHub issue: %w", err)
	}
	// Add to all_schemas.hcl
	if mode == "resources" {
		allSchemas, err := parseSchemaToStruct("internal/provider/allschemas.hcl")
		if err != nil {
			return fmt.Errorf("failed to parse allschemas.hcl: %w", err)
		}
		for i := range allSchemas.Resources {
			if resource == allSchemas.Resources[i].CloudFormationTypeName {
				switch mode {
				case "resources":
					allSchemas.Resources[i].SuppressResourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing resource generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
					break
				case "singular-data-sources":
					allSchemas.Resources[i].SuppressSingularDataSourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing singular data source generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
					break
				case "plural-data-source":
					allSchemas.Resources[i].SuppressPluralDataSourceGeneration = true
					tflog.Debug(ctx, fmt.Sprintf("Suppressing plural data source generation for %s", allSchemas.Resources[i].CloudFormationTypeName))
					break
				default:
					if allSchemas.Resources[i].SuppressionReason == "" {
						allSchemas.Resources[i].SuppressionReason = schemaError + issueURL
					} else {
						allSchemas.Resources[i].SuppressionReason += ", " + schemaError
					}
				}
				// Log this
				break
			}
		}
		err = writeSchemasToHCLFile(allSchemas, "internal/provider/allschemas.hcl")
		if err != nil {
			return fmt.Errorf("failed to write schemas to HCL file: %w", err)
		}
	}

	err = addCheckout(resource, new)
	if err != nil {
		return fmt.Errorf("failed to add checkout: %w", err)
	}
	return nil
}

func addCheckout(resource string, new bool) error {
	file, err := os.Open("./checkout.txt")
	if err != nil {
		return fmt.Errorf("failed to open checkout file: %w", err)
	}
	if !new {
		file, err = os.OpenFile("./checkout.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open checkout file for writing: %w", err)
		}
		_, err = file.WriteString(resource + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to checkout file: %w", err)
		}
	}
	return nil
}

func createIssue(resource, schemaError string, client *github.Client) (string, error) {
	ctx := context.Background()
	repoOwner := "hashicorp"
	repoName := "terraform-provider-awscc"

	issueTitle := fmt.Sprintf("Reasource Suppression: %s", resource)
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

func GetResourceFromLog() (string, error) {
	var resourceName string
	logData, err := os.ReadFile("logs.txt")
	if err != nil {
		return "", fmt.Errorf("failed to read logs file: %w", err)
	}
	logLines := strings.Split(strings.TrimSpace(string(logData)), "\n")
	if len(logLines) > 0 {
		resourceName = logLines[len(logLines)-1]
	} else {
		return "", fmt.Errorf("no resource name found in logs")
	}
	return resourceName, nil
}

func isNew(resourceName string, currentSchemas allschemas.AllSchemas) bool {
	resourceName = strings.ToLower(resourceName)
	for _, r := range currentSchemas.Resources {
		if r.ResourceTypeName == resourceName {
			return false
		}
	}
	return true
}
