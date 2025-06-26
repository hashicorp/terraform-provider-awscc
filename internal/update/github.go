package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/google/go-github/v72/github"
)

// GitHubConfig encapsulates all GitHub-related configuration and client information
type GitHubConfig struct {
	Client      *github.Client
	Repository  string
	RepoOwner   string
	RepoName    string
	CurrentDate string
}

// NewGitHubConfig creates a new GitHubConfig with the given parameters
func NewGitHubConfig(client *github.Client, repositoryLink string, date string) *GitHubConfig {
	config := &GitHubConfig{
		Client:      client,
		CurrentDate: date,
		RepoOwner:   "hashicorp",
		RepoName:    "terraform-provider-awscc",
	}

	// If repository link is provided, use it to extract owner and name
	if repositoryLink != "" {
		parts := strings.Split(strings.TrimPrefix(repositoryLink, "https://github.com/"), "/")
		if len(parts) >= 2 {
			config.RepoOwner = parts[0]
			config.RepoName = parts[1]
		}
	}

	// Set full repository path
	config.Repository = config.RepoOwner + "/" + config.RepoName

	return config
}

func createRemoteBranch(currentData string) (string, error) {
	// Add a random component to avoid branch name collisions
	randomSuffix := fmt.Sprintf("%d", time.Now().UnixNano()%10000)
	name := fmt.Sprintf("f-%s-schema-updates-%s", currentData, randomSuffix)
	cmd := exec.Command("git", "push", "-u", "origin", name)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to push branch: %w", err)
	}
	return name, nil
}

func createPullRequest(ctx context.Context, config *GitHubConfig, changes *[]string, pullRequest string, filepaths *UpdateFilePaths) (string, error) {
	repoOwner := config.RepoOwner
	repoName := config.RepoName
	client := config.Client
	currentData := config.CurrentDate

	// If repository link is provided in filepaths, use it to extract owner and name
	if filepaths.RepositoryLink != "" {
		parts := strings.Split(strings.TrimPrefix(filepaths.RepositoryLink, "https://github.com/"), "/")
		if len(parts) >= 2 {
			repoOwner = parts[0]
			repoName = parts[1]
		}
	}

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
		HeadRepo:            github.Ptr(config.Repository),
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
	labels := []string{"bug", "resource-suppression", "prioritized"}
	_, _, err = client.Issues.AddLabelsToIssue(ctx, repoOwner, repoName, pr.GetNumber(), labels)
	if err != nil {
		return "", fmt.Errorf("failed to add labels to pull request: %w", err)
	}
	log.Printf("Labels added to pull request: %v\n", labels)

	return pr.GetHTMLURL(), nil
}

func createIssue(ctx context.Context, resource, schemaError string, config *GitHubConfig, repositoryLink string) (string, error) {
	repoOwner := config.RepoOwner
	repoName := config.RepoName
	client := config.Client

	// If repository link is provided, use it to extract owner and name
	if repositoryLink != "" {
		parts := strings.Split(strings.TrimPrefix(repositoryLink, "https://github.com/"), "/")
		if len(parts) >= 2 {
			repoOwner = parts[0]
			repoName = parts[1]
		}
	}

	issueTitle := fmt.Sprintf("Resource Suppression: %s", resource)
	issueBody := createFormattedIssue(
		fmt.Sprintf("Suppress generation of resource `%s` due to schema error.", resource),
		fmt.Sprintf("`%s`", resource),
	)

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

// FormatCommitMessage generates a formatted commit message for CloudFormation schema updates.
// It formats messages similar to those seen in PR #2347.
func FormatCommitMessage(date string, region string, updateType string, details string) string {
	var messageFormat string

	switch updateType {
	case "refresh":
		messageFormat = "%s CloudFormation schemas in %s; Refresh existing schemas."
	case "new":
		messageFormat = "%s CloudFormation schemas in %s; New schemas."
	case "terraform-resources":
		messageFormat = "%s CloudFormation schemas in %s; Generate Terraform resources."
	case "terraform-datasources":
		messageFormat = "%s CloudFormation schemas in %s; Generate Terraform data sources."
	case "docs":
		return fmt.Sprintf("%s Run 'make docs-all'.", date)
	case "changelog":
		return fmt.Sprintf("Add CHANGELOG entries.")
	case "release":
		return fmt.Sprintf("Prepare for v%s release.", details)
	default:
		messageFormat = "%s CloudFormation schemas in %s; %s"
		if details == "" {
			details = "Updates."
		}
	}

	if updateType == "default" {
		return fmt.Sprintf(messageFormat, date, region, details)
	} else if updateType != "docs" && updateType != "changelog" && updateType != "release" {
		return fmt.Sprintf(messageFormat, date, region)
	}

	return fmt.Sprintf(messageFormat, date, region)
}

// GenerateCompletePRTemplate creates a fully formatted PR description with all required sections
// including community note, rollback plan, security controls, and test results.
func GenerateCompletePRTemplate(testResults string, version string) string {
	template := `<!--- See what makes a good Pull Request at: https://github.com/hashicorp/terraform-provider-awscc/blob/main/contributing/CONTRIBUTING.md --->

<!--- Please keep this note for the community --->

### Community Note

* Please vote on this pull request by adding a 👍 [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/) to the original pull request comment to help the community and maintainers prioritize this request
* Please do not leave "+1" or other comments that do not add relevant new information or questions, they generate extra noise for pull request followers and do not help prioritize the request
* The resources and data sources in this provider are generated from the CloudFormation schema, so they can only support the actions that the underlying schema supports. For this reason submitted bugs should be limited to defects in the generation and runtime code of the provider. Customizing behavior of the resource, or noting a gap in behavior are not valid bugs and should be submitted as enhancements to AWS via the CloudFormation Open Coverage Roadmap.

<!--- Thank you for keeping this note for the community --->

<!-- heimdall_github_prtemplate:grc-pci_dss-2024-01-05 -->

## Rollback Plan

If a change needs to be reverted, we will publish an updated version of the library.

## Changes to Security Controls

Are there any changes to security controls (access controls, encryption, logging) in this pull request? If so, explain.

## Description

`

	if version != "" {
		template += fmt.Sprintf("Also prepares for the **v%s** release.\n\n", version)
	}

	if testResults != "" {
		template += fmt.Sprintf("```console\n%s\n```\n", testResults)
	}

	return template
}

// createFormattedPullRequest creates a pull request with our standardized template
func createFormattedPullRequest(ctx context.Context, config *GitHubConfig, testResults string, version string, pullRequest string) (string, error) {
	repoOwner := config.RepoOwner
	repoName := config.RepoName
	client := config.Client
	currentData := config.CurrentDate

	// If repository is a GitHub URL, parse it to extract owner and repo name
	if strings.HasPrefix(config.Repository, "https://github.com/") {
		parts := strings.Split(strings.TrimPrefix(config.Repository, "https://github.com/"), "/")
		if len(parts) >= 2 {
			repoOwner = parts[0]
			repoName = parts[1]
		}
	}

	prTitle := fmt.Sprintf("%s Schema Updates #%s", currentData, version)

	// Use our template generator for the PR body
	prBody := GenerateCompletePRTemplate(testResults, version)

	prRequest := &github.NewPullRequest{
		Title:               &prTitle,
		Body:                &prBody,
		Base:                github.Ptr("main"),
		Head:                github.Ptr(pullRequest),
		HeadRepo:            github.Ptr(config.Repository),
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

	log.Printf("Formatted pull request created: %s\n", pr.GetHTMLURL())

	// Add labels to the pull request
	labels := []string{"prioritized"}
	_, _, err = client.Issues.AddLabelsToIssue(ctx, repoOwner, repoName, pr.GetNumber(), labels)
	if err != nil {
		return "", fmt.Errorf("failed to add labels to pull request: %w", err)
	}
	log.Printf("Labels added to pull request: %v\n", labels)

	return pr.GetHTMLURL(), nil
}

// createFormattedPullRequestWithTests creates a pull request with our standardized template
func createFormattedPullRequestWithTests(ctx context.Context, config *GitHubConfig, pullRequest string, version string) (string, error) {
	// Get the stored test results from the global variable
	testResults := GetAcceptanceTestResults()

	// Create the formatted PR with the test results
	return createFormattedPullRequest(ctx, config, testResults, version, pullRequest)
}

// RunAcceptanceTests executes the acceptance tests and captures the output for the PR description
func RunAcceptanceTests() (string, error) {
	// Create a buffer to capture command output
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer

	// Prepare the command to run acceptance tests
	cmd := exec.Command("make", MakeTestAccCmd, PKGNameArg, TestArgsArg, AccTestParallelismArg)
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Display the command being run
	log.Printf("Running command: %s %s\n", cmd.Path, strings.Join(cmd.Args, " "))

	// Execute the command
	err := cmd.Run()

	// Combine stdout and stderr for the complete test output
	testOutput := fmt.Sprintf("%% TF_LOG=ERROR make testacc %s %s %s\n%s",
		PKGNameArg, TestArgsArg, AccTestParallelismArg,
		outBuffer.String())

	if errBuffer.Len() > 0 {
		testOutput += errBuffer.String()
	}

	// Return the test output even if there was an error
	return testOutput, err
}

func submitOnGit(client *github.Client, changes *[]string, filePaths *UpdateFilePaths, execData string, repoOwner string, repoName string) (string, error) {
	// Create a new branch and push it to remote
	currentDate := GetCurrentDate()
	branchName, err := createRemoteBranch(currentDate)
	if err != nil {
		return "", fmt.Errorf("failed to create and push remote branch: %w", err)
	}

	// Create GitHubConfig
	config := &GitHubConfig{
		Client:      client,
		Repository:  repoOwner + "/" + repoName,
		RepoOwner:   repoOwner,
		RepoName:    repoName,
		CurrentDate: currentDate,
	}

	// Create a pull request with the changes
	prURL, err := createPullRequest(context.Background(), config, changes, branchName, filePaths)
	if err != nil {
		return "", fmt.Errorf("failed to create pull request: %w", err)
	}

	return prURL, nil
}

// GetCurrentDate returns the current date in the format specified by DateFormat.
// This function centralizes date formatting for consistent use across GitHub operations.
func GetCurrentDate() string {
	return time.Now().Format(DateFormat)
}

// createFormattedIssue generates a GitHub issue template with the standardized community format
// and customizable description and affected resources sections.
// Returns the formatted text without creating an actual issue.
func createFormattedIssue(description, affectedResources string) string {
	issueBody := fmt.Sprintf(`<!--- Please keep this note for the community --->

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

%s

### References

- #156`, description, affectedResources)

	return issueBody
}
