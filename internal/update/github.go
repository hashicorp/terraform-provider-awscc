package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/v72/github"
)

func createRemoteBranch(currentData string) (string, error) {
	name := fmt.Sprintf("f-%s-schema-updates", currentData)
	cmd := exec.Command("git", "push", "-u", "origin", name)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to push branch: %w", err)
	}
	return name, nil

}

func createPullRequest(ctx context.Context, client *github.Client, changes *[]string, currentData string, pullRequest string, filepaths *UpdateFilePaths, repository string) (string, error) {
	repoOwner := "hashicorp"
	repoName := "terraform-provider-awscc"

	prTitle := fmt.Sprintf("Schema Updates for %s", currentData)

	prBody := ""

	for _, change := range *changes {
		prBody += fmt.Sprintf("- %s\n", change)
	}
	if prBody == "" {
		prBody = "No schema changes detected."
	}
	prRequest := &github.NewPullRequest{
		Title:               &prTitle,
		Body:                &prBody,
		Base:                github.Ptr("main"),
		Head:                github.Ptr(pullRequest),
		HeadRepo:            github.Ptr(repository),
		MaintainerCanModify: github.Ptr(true),
	}
	pr, _, err := client.PullRequests.Create(ctx, repoOwner, repoName, prRequest)
	if err != nil {
		return "", fmt.Errorf("failed to create pull request: %w", err)
	}
	if pr == nil {
		return "", fmt.Errorf("pull request is nil")
	}
	if pr.GetHTMLURL() == "" {
		return "", fmt.Errorf("pull request URL is empty")
	}

	log.Printf("Pull request created: %s\n", pr.GetHTMLURL())

	// Add labels to the pull request
	labels := []string{"bug", "resource-suppression"}
	_, _, err = client.Issues.AddLabelsToIssue(ctx, repoOwner, repoName, pr.GetNumber(), labels)
	if err != nil {
		return "", fmt.Errorf("failed to add labels to pull request: %w", err)
	}
	log.Printf("Labels added to pull request: %v\n", labels)

	return pr.GetHTMLURL(), nil
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
