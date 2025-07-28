// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/v72/github"
)

// Define test constants for GitHub API tests.
// These constants configure the test environment for interacting with GitHub.
const (
	// Repository owner for API tests
	testRepoOwner = ""

	// Repository name for API tests
	testRepoName = ""

	// Test date for commit messages and PR creation
	testCurrentDate = "2025-06-26"

	// Resource name for issue creation tests
	testResourceName = "AWS::Test::Resource"
)

// Variables that can't be constants
var (
	// GitHub API token (required for API tests)
	// Do not Commit :)
	testGitHubToken = ""

	// Test changes to include in pull requests
	testChangeList = []string{"Schema updated for AWS::S3::Bucket", "Added new property to AWS::Lambda::Function"}
)

// Setup function to create a GitHub client for testing.
// This function initializes a GitHub API client with authentication for tests that interact with GitHub.
// If the GitHub token is not set, the tests requiring API access will be skipped.
func setupGitHubClient(t *testing.T) *github.Client {
	if testGitHubToken == "" {
		t.Skip("GITHUB_TOKEN environment variable not set, skipping GitHub API tests")
	}

	// Create a GitHub client using the token
	client := github.NewClient(nil).WithAuthToken(testGitHubToken)
	if client == nil {
		t.Fatal("GitHub client should not be nil")
	}

	return client
}

// Define UpdateFilePaths for testing
func getTestFilePaths() *UpdateFilePaths {
	return &UpdateFilePaths{
		RunMakesResourceLog:      "test_resource.log",
		RunMakesOutput:           "test_output.txt",
		RunMakesErrors:           "test_errors.txt",
		SuppressionCheckout:      "test_checkout.txt",
		AwsSchemas:               "test_schemas.json",
		AllSchemasHCL:            "test_schemas.hcl",
		AllSchemasDir:            "./test_schemas",
		LastResource:             "test_last_resource.txt",
		CloudFormationSchemasDir: "./test_cf_schemas",
		RepositoryLink:           fmt.Sprintf("https://github.com/%s/%s", testRepoOwner, testRepoName),
	}
}

// TestGenerateCompletePRTemplate tests the GenerateCompletePRTemplate function
// which creates a standardized pull request template with all required sections.
// This test verifies that the template includes critical sections like the community note,
// rollback plan, and properly incorporates version information and test results.
// It tests both with and without version/test results to ensure flexibility.
func TestGenerateCompletePRTemplate(t *testing.T) {
	// Helper function to check if a string contains another
	checkContains := func(s, substr string) bool {
		return strings.Contains(s, substr)
	}

	// Test with test results and version
	testResults := "All tests passed successfully"
	version := "1.2.3"
	template := GenerateCompletePRTemplate(testResults, version)

	if !checkContains(template, "Community Note") {
		t.Error("Expected template to contain 'Community Note'")
	}
	if !checkContains(template, "Rollback Plan") {
		t.Error("Expected template to contain 'Rollback Plan'")
	}
	if !checkContains(template, "v1.2.3") {
		t.Error("Expected template to contain 'v1.2.3'")
	}
	if !checkContains(template, "All tests passed successfully") {
		t.Error("Expected template to contain 'All tests passed successfully'")
	}

	// Test without test results and version
	template = GenerateCompletePRTemplate("", "")
	if !checkContains(template, "Community Note") {
		t.Error("Expected template to contain 'Community Note'")
	}
	if !checkContains(template, "Rollback Plan") {
		t.Error("Expected template to contain 'Rollback Plan'")
	}
	if checkContains(template, "v1.2.3") {
		t.Error("Expected template to NOT contain 'v1.2.3'")
	}
}

func T_CreateRemoteBranch(client *github.Client) (string, error) {
	ctx := context.Background()
	repoOwner := testRepoOwner
	repoName := testRepoName

	// Verify repository exists
	r, resp, err := client.Repositories.Get(ctx, repoOwner, repoName)
	if err != nil {
		return "", fmt.Errorf("Failed to get repository: %v", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Expected status code 200, got: %d", resp.StatusCode)
	}

	// Get the SHA of the default branch to use as the base for our new branch
	defaultBranch := r.GetDefaultBranch()
	if defaultBranch == "" {
		defaultBranch = "main" // Fallback to main if default branch isn't set
	}

	// Get the reference to the default branch
	ref, resp, err := client.Git.GetRef(ctx, repoOwner, repoName, "refs/heads/"+defaultBranch)
	if err != nil {
		return "", fmt.Errorf("Failed to get reference for default branch: %v", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Expected status code 200 for ref, got: %d", resp.StatusCode)
	}

	// Generate a unique branch name for testing
	currentDate := GetCurrentDate()
	randomSuffix := fmt.Sprintf("%d", os.Getpid())
	branchName := fmt.Sprintf("test-%s-branch-%s", currentDate, randomSuffix)

	log.Printf("Attempting to create branch: %s\n", branchName)
	log.Printf("Using repository: %s/%s\n", repoOwner, repoName)
	log.Printf("Base SHA: %s\n", ref.Object.GetSHA())

	// Check if branch already exists and delete it if it does
	_, checkResp, checkErr := client.Git.GetRef(ctx, repoOwner, repoName, "refs/heads/"+branchName)
	if checkErr == nil && checkResp.StatusCode == 200 {
		log.Printf("Branch %s already exists, deleting it first\n", branchName)
		delResp, delErr := client.Git.DeleteRef(ctx, repoOwner, repoName, "heads/"+branchName)
		if delErr != nil {
			log.Printf("Warning: Failed to delete existing branch: %v (status: %d)\n", delErr, delResp.StatusCode)
			// Try with a more unique name
			randomSuffix = fmt.Sprintf("%d-%d", os.Getpid(), time.Now().UnixNano()%10000)
			branchName = fmt.Sprintf("test-%s-branch-%s", currentDate, randomSuffix)
			log.Printf("Trying with new branch name: %s\n", branchName)
		} else {
			log.Printf("Successfully deleted existing branch\n")
		}
	} else if checkResp != nil && checkResp.StatusCode != 404 {
		log.Printf("Unexpected response when checking for existing branch: %d\n", checkResp.StatusCode)
	}

	// Create a new reference (branch) using the SHA from the default branch
	newRef := &github.Reference{
		Ref: github.Ptr("refs/heads/" + branchName),
		Object: &github.GitObject{
			SHA: ref.Object.SHA,
		},
	}

	// Create the branch on the remote repository
	log.Printf("Creating new branch with ref: %s\n", newRef.GetRef())
	createdRef, resp, err := client.Git.CreateRef(ctx, repoOwner, repoName, newRef)
	if err != nil {
		return "", fmt.Errorf("Failed to create new branch '%s': %v (status: %d)", branchName, err, resp.StatusCode)
	}
	if resp.StatusCode != 201 {
		return "", fmt.Errorf("Expected status code 201 for branch creation, got: %d", resp.StatusCode)
	}

	// Verify the branch was created successfully
	if createdRef == nil || createdRef.GetRef() != "refs/heads/"+branchName {
		return "", fmt.Errorf("Created branch reference is invalid")
	}

	// Create a test commit on the new branch so there are changes for the pull request
	log.Printf("Creating test commit on branch: %s\n", branchName)

	// Create a test file content
	testContent := fmt.Sprintf("# Test File\n\nThis is a test file created on %s for testing purposes.\nBranch: %s\nTimestamp: %s\n",
		currentDate, branchName, time.Now().Format(time.RFC3339))

	// Create a blob with the test content
	blob := &github.Blob{
		Content:  github.Ptr(testContent),
		Encoding: github.Ptr("utf-8"),
	}

	createdBlob, resp, err := client.Git.CreateBlob(ctx, repoOwner, repoName, blob)
	if err != nil {
		return "", fmt.Errorf("Failed to create blob: %v", err)
	}
	if resp.StatusCode != 201 {
		return "", fmt.Errorf("Expected status code 201 for blob creation, got: %d", resp.StatusCode)
	}

	// Get the current tree from the base commit
	baseCommit, _, err := client.Git.GetCommit(ctx, repoOwner, repoName, ref.Object.GetSHA())
	if err != nil {
		return "", fmt.Errorf("Failed to get base commit: %v", err)
	}

	// Create a new tree with the test file
	testFileName := fmt.Sprintf("test-file-%s.md", randomSuffix)
	entries := []*github.TreeEntry{
		{
			Path: github.Ptr(testFileName),
			Mode: github.Ptr("100644"),
			Type: github.Ptr("blob"),
			SHA:  createdBlob.SHA,
		},
	}

	createdTree, resp, err := client.Git.CreateTree(ctx, repoOwner, repoName, baseCommit.Tree.GetSHA(), entries)
	if err != nil {
		return "", fmt.Errorf("Failed to create tree: %v", err)
	}
	if resp.StatusCode != 201 {
		return "", fmt.Errorf("Expected status code 201 for tree creation, got: %d", resp.StatusCode)
	}

	// Create a commit with the new tree
	commitMessage := fmt.Sprintf("Test commit for branch %s\n\nThis commit was created for testing purposes on %s", branchName, currentDate)
	commit := &github.Commit{
		Message: github.Ptr(commitMessage),
		Tree:    createdTree,
		Parents: []*github.Commit{{SHA: ref.Object.SHA}},
	}

	createdCommit, resp, err := client.Git.CreateCommit(ctx, repoOwner, repoName, commit, &github.CreateCommitOptions{})
	if err != nil {
		return "", fmt.Errorf("Failed to create commit: %v", err)
	}
	if resp.StatusCode != 201 {
		return "", fmt.Errorf("Expected status code 201 for commit creation, got: %d", resp.StatusCode)
	}

	// Update the branch reference to point to the new commit
	updateRef := &github.Reference{
		Ref: github.Ptr("refs/heads/" + branchName),
		Object: &github.GitObject{
			SHA: createdCommit.SHA,
		},
	}

	_, resp, err = client.Git.UpdateRef(ctx, repoOwner, repoName, updateRef, false)
	if err != nil {
		return "", fmt.Errorf("Failed to update branch reference: %v", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Expected status code 200 for reference update, got: %d", resp.StatusCode)
	}

	log.Printf("Successfully created branch %s with test commit %s\n", branchName, createdCommit.GetSHA())
	return branchName, nil
}

// TestCreatePullRequest tests the createPullRequest function using the GitHub API.
// This test verifies that the function can create a pull request with the specified changes
// and correctly return the URL of the created PR. It also validates that the PR
// is created with the expected title and body content based on the change list.
// This test requires a valid GitHub token and is skipped by default unless explicitly enabled.
func TestCreatePullRequest(t *testing.T) {
	// Only run this test if explicitly enabled

	client := setupGitHubClient(t)
	ctx := context.Background()
	// Set up test data
	changes := &testChangeList
	filepaths := getTestFilePaths()

	// Call the function
	repoOwner := testRepoOwner
	repoName := testRepoName

	config := &GitHubConfig{
		Client:      client,
		Repository:  testRepoOwner + "/" + testRepoName,
		RepoOwner:   repoOwner,
		RepoName:    repoName,
		CurrentDate: testCurrentDate,
	}

	// Generate a random branch name for testing
	randomBranchName, err := T_CreateRemoteBranch(client)
	if err != nil {
		t.Fatalf("Expected no error creating remote branch, got: %v", err)
	}
	t.Logf("Using random branch name: %s", randomBranchName)

	prURL, err := createPullRequest(ctx, config, changes, randomBranchName, filepaths, "test execution data")

	// Check the results
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if prURL == "" {
		t.Error("Expected non-empty PR URL")
	}
	if !strings.Contains(prURL, "github.com") {
		t.Errorf("Expected PR URL to contain github.com, got: %s", prURL)
	}
	if !strings.Contains(prURL, "pull") {
		t.Errorf("Expected PR URL to contain pull, got: %s", prURL)
	}

	t.Logf("Created pull request: %s", prURL)
}

// TestCreateIssue tests the createIssue function using the GitHub API.
// This function verifies that GitHub issues can be created for schema errors in resources.
// The current implementation is a placeholder returning empty string, but the test ensures
// the function signature and behavior is ready for actual implementation.
func TestCreateIssue(t *testing.T) {
	// Only run this test if explicitly enabled
	ctx := context.Background()
	client := setupGitHubClient(t)

	// Get test file paths with repository link
	filepaths := getTestFilePaths()

	// Call the function
	config := &GitHubConfig{
		Client:      client,
		Repository:  testRepoOwner + "/" + testRepoName,
		RepoOwner:   testRepoOwner,
		RepoName:    testRepoName,
		CurrentDate: testCurrentDate,
	}

	_, err := createIssue(ctx, testResourceName, "test error", config, filepaths.RepositoryLink)

	// The function currently returns an empty string and nil error
	// This is a placeholder for when the function is implemented
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

// TestCreateFormattedPullRequest tests the createFormattedPullRequest function.
// This test verifies that the function can create a pull request with our standardized
// template including all required sections such as community notes, rollback plans,
// and properly incorporating version information and test results.
func TestCreateFormattedPullRequest(t *testing.T) {
	// Only run this test if explicitly enabled
	client := setupGitHubClient(t)
	ctx := context.Background()

	// Set up test data
	testResults := "All tests passed successfully"
	version := "1.2.3"
	filepaths := getTestFilePaths()
	// Call the function using the repository link instead of just the name
	repoOwner := testRepoOwner
	repoName := testRepoName
	testCurrentDate := GetCurrentDate()

	config := &GitHubConfig{
		Client:      client,
		Repository:  filepaths.RepositoryLink,
		RepoOwner:   repoOwner,
		RepoName:    repoName,
		CurrentDate: testCurrentDate,
	}

	// Generate a random branch name for testing
	randomBranchName, err := T_CreateRemoteBranch(client)
	if err != nil {
		t.Fatalf("Expected no error creating remote branch, got: %v", err)
	}
	t.Logf("Using random branch name: %s", randomBranchName)

	prURL, err := createFormattedPullRequest(ctx, config, testResults, version, randomBranchName)

	// Check the results
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if prURL == "" {
		t.Error("Expected non-empty PR URL")
	}
	if !strings.Contains(prURL, "github.com") {
		t.Errorf("Expected PR URL to contain github.com, got: %s", prURL)
	}
	if !strings.Contains(prURL, "pull") {
		t.Errorf("Expected PR URL to contain pull, got: %s", prURL)
	}

	t.Logf("Created formatted pull request: %s", prURL)
}

// TestRunAcceptanceTests tests the RunAcceptanceTests function
// which executes acceptance tests and captures their output for PR descriptions.
// This test validates that the test runner can execute tests and capture their output
// successfully, regardless of whether the tests themselves pass or fail.
// This test is skipped by default as it would run potentially time-consuming acceptance tests.
func TestRunAcceptanceTests(t *testing.T) {
	// This test would actually run acceptance tests which could be time-consuming
	// For unit testing purposes, we'll skip it by default
	if os.Getenv("RUN_ACCEPTANCE_TESTS") != "true" {
		t.Skip("RUN_ACCEPTANCE_TESTS not set to true, skipping acceptance tests")
	}

	output, err := RunAcceptanceTests()

	// We don't necessarily expect the tests to pass, but we expect the function to execute
	if output == "" {
		t.Error("Expected non-empty output from acceptance tests")
	}

	t.Logf("RunAcceptanceTests output: %s", output)
	t.Logf("RunAcceptanceTests error: %v", err)
}

// TestGetCurrentDate tests the GetCurrentDate function
// which returns the current date in the standard format used across GitHub operations.
// This test ensures the function returns a date string in the expected format (YYYY-MM-DD).
func TestGetCurrentDate(t *testing.T) {
	// Get the current date
	date := GetCurrentDate()

	// Check that the date is in the expected format (YYYY-MM-DD)
	if len(date) != 10 {
		t.Errorf("Date should be 10 characters in format YYYY-MM-DD, got: %s", date)
	}

	// Check format with dashes in the right places
	if date[4] != '-' || date[7] != '-' {
		t.Errorf("Date should have dashes in positions 4 and 7, got: %s", date)
	}

	// Check that all other characters are digits
	for i, char := range date {
		if i != 4 && i != 7 && (char < '0' || char > '9') {
			t.Errorf("Expected digit at position %d, got: %c", i, char)
		}
	}
}

// TestParseRepositoryLink tests the repository link parsing functionality
// added to the GitHub API integration functions. It ensures the proper extraction
// of owner and repo name from a GitHub URL.
func TestParseRepositoryLink(t *testing.T) {
	// Test cases
	testCases := []struct {
		name      string
		link      string
		wantOwner string
		wantRepo  string
	}{
		{
			name:      "Standard GitHub URL",
			link:      "https://github.com/" + testRepoOwner + "/" + testRepoName,
			wantOwner: testRepoOwner,
			wantRepo:  testRepoName,
		},
		{
			name:      "URL with trailing slash",
			link:      "https://github.com/" + testRepoOwner + "/" + testRepoName + "/",
			wantOwner: testRepoOwner,
			wantRepo:  testRepoName,
		},
		{
			name:      "URL with additional path",
			link:      "https://github.com/" + testRepoOwner + "/" + testRepoName + "/tree/main",
			wantOwner: testRepoOwner,
			wantRepo:  testRepoName,
		},
		{
			name:      "Empty string",
			link:      "",
			wantOwner: testRepoOwner, // Default value
			wantRepo:  testRepoName,  // Default value
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // Create filepaths with the test repository link
			filepaths := &UpdateFilePaths{RepositoryLink: tc.link}

			// We'll use a custom function to validate the parsing without making API calls
			owner, repo := extractRepoInfo(filepaths.RepositoryLink)

			// Check results
			if owner != tc.wantOwner {
				t.Errorf("Owner = %q, want %q", owner, tc.wantOwner)
			}
			if repo != tc.wantRepo {
				t.Errorf("Repo = %q, want %q", repo, tc.wantRepo)
			}
		})
	}
}

// Helper function to extract repository owner and name from a link
func extractRepoInfo(link string) (string, string) {
	owner := testRepoOwner // Default values
	repo := testRepoName

	if link == "" {
		return owner, repo
	}

	parts := strings.Split(strings.TrimPrefix(link, "https://github.com/"), "/")
	if len(parts) >= 2 {
		owner = parts[0]
		repo = parts[1]
	}

	return owner, repo
}
