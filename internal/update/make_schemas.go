package update

import (
	"context"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v72/github"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func makeSchemas(client *github.Client, currentSchemas allschemas.AllSchemas) error {
	var waitTime int = 0
	if waitTime > 300 {
		return fmt.Errorf("wait time exceeded maximum of 300 seconds")
	}

	file, err := os.OpenFile("make_schema_errors.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

	for stat.Size() > 0 {
		time.Sleep(time.Duration(waitTime) * time.Second)
		err = execCommand("sh", "-c", `make schemas | grep "error" > make_schema_errors.txt`)
		if err != nil {
			return fmt.Errorf("failed to execute make schemas command: %w", err)
		}
		data, err := os.ReadFile("make_schema_errors.txt")
		if err != nil {
			return fmt.Errorf("failed to read error log file: %w", err)
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line == "fatal error: stack overflow" { // Recursive Data Type
				var resourceName string
				var new bool = true
				logData, err := os.ReadFile("logs.txt")
				if err != nil {
					return fmt.Errorf("failed to read logs file: %w", err)
				}
				logLines := strings.Split(strings.TrimSpace(string(logData)), "\n")
				if len(logLines) > 0 {
					resourceName = logLines[len(logLines)-1]
				} else {
					return fmt.Errorf("no resource name found in logs")
				}
				defer file.Close()
				resourceName = strings.ToLower(strings.ReplaceAll(resourceName, "::", "_"))
				for _, r := range currentSchemas.Resources {
					if r.CloudFormationTypeName == resourceName {
						new = false
						break
					}
				}

				suppress(resourceName, line, client, new)
			}
			if strings.Contains(line, "is of unsupported type:") { // Unsupported Data Type
				var resourceName string
				var new bool = true

				// Extract resource name from error message like:
				// "../service/cloudformation/schemas/AWS_AccessAnalyzer_Analyzer.json: emitting schema code:"
				parts := strings.Split(line, "/")
				if len(parts) > 0 {
					for _, part := range parts {
						if strings.HasPrefix(part, "AWS") && strings.Contains(part, ".json") {
							resourceName = strings.ToLower(strings.TrimSuffix(part, ".json"))
							break
						}
					}
				}
				for _, r := range currentSchemas.Resources {
					if r.CloudFormationTypeName == resourceName {
						new = false
						break
					}
				}

				suppress(resourceName, line, client, new)

			} else if strings.Contains(line, "for Terraform attribute") { // Overwriting Terraform Attribute
				var resourceName string
				var new bool = true

				parts := strings.Split(line, "/")
				if len(parts) > 0 {
					for _, part := range parts {
						if strings.HasPrefix(part, "AWS") && strings.Contains(part, ".json") {
							resourceName = strings.ToLower(strings.TrimSuffix(part, ".json"))
							break
						}
					}
				}
				for _, r := range currentSchemas.Resources {
					if r.CloudFormationTypeName == resourceName {
						new = false
						break
					}
				}

				suppress(resourceName, line, client, new)
			} else if strings.Contains(line, "404") { // Existing Resource Missing or Rate limit | Can a resource disappear because of a rate limit?
				var rateLimitError bool = true
				var resourceName string = "False"

				for _, word := range strings.Fields(line) {
					if strings.HasPrefix(word, "/'AWS") {
						resourceName = strings.ReplaceAll(word, "'", "")
						rateLimitError = false
						break
					}
					if rateLimitError { // Rate limit error
						waitTime = int(math.Max(1, math.Floor(math.Pow(float64(waitTime), 1.2))))
						break
					} else { // Existing Resource Missing
						resourceName = strings.ToLower((resourceName))
						suppress(resourceName, line, client, false)
					}
				}
			} else if strings.Contains(line, "403") {
				return fmt.Errorf("authentication failed: no valid AWS credentials")
			} else {
				return fmt.Errorf("unhandled schema error: %s", line)
			}
		}
	}

	return nil
}

func suppress(resource, schemaError string, client *github.Client, new bool) error {
	// Create Issue
	issueURL, err := createIssue(resource, schemaError, client)
	if err != nil {
		return fmt.Errorf("failed to create GitHub issue: %w", err)
	}
	// Add to all_schemas.hcl
	allSchemas, err := parseSchemaToStruct("internal/provider/allschemas.hcl")
	if err != nil {
		return fmt.Errorf("failed to parse allschemas.hcl: %w", err)
	}
	for i := range allSchemas.Resources {
		if resource == allSchemas.Resources[i].CloudFormationTypeName {
			if allSchemas.Resources[i].SuppressionReason == "" {
				allSchemas.Resources[i].SuppressionReason = schemaError + issueURL
			} else {
				allSchemas.Resources[i].SuppressionReason += ", " + schemaError
			}
			allSchemas.Resources[i].SuppressResourceGeneration = true
			// Log this
			break
		}
	}

	err = writeSchemasToHCLFile(allSchemas, "internal/provider/allschemas.hcl")
	if err != nil {
		return fmt.Errorf("failed to write schemas to HCL file: %w", err)
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
