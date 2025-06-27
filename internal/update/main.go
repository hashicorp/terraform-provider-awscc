package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/go-github/v72/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

// Global variable to store test results for PR creation
var AcceptanceTestResults string

// Constants for file paths and patterns
const (
	// File names and patterns
	UpdateFilePathsHCL          = "internal/update/update_filepaths.hcl"
	AvailableSchemasPrefix      = "available_schemas."
	HCLExtension                = ".hcl"
	AvailableSchemasFilePattern = "available_schemas.%s.hcl"

	// Commit message components
	CloudFormationRegion = "CloudFormation schemas in us-east-1"

	// Git commit message formats
	CommitMsgRefreshSchemas    = "%s " + CloudFormationRegion + "; Refresh existing schemas."
	CommitMsgNewSchemas        = "%s " + CloudFormationRegion + "; New schemas."
	CommitMsgResourceSchemas   = "%s " + CloudFormationRegion + "; Generate Terraform resource schemas."
	CommitMsgDataSourceSchemas = "%s " + CloudFormationRegion + "; Generate Terraform data source schemas."
	CommitMsgDocs              = "%s Run 'make docs-all'."

	// Branch name format
	BranchNameFormat    = "update-schemas-%d"
	BranchNameMaxRandom = 1000000

	// Date format
	DateFormat = "2006-01-02"

	// Environment variables
	GithubTokenEnv = "GITHUB_TOKEN"
	TestModeEnv    = "TEST_MODE"

	// Make commands
	MakeBuildCmd   = "build"
	MakeTestAccCmd = "testacc"
	MakeDocsAllCmd = "docs-all"

	// Make targets
	TargetSchemas             = "schemas"
	TargetResources           = "resources"
	TargetSingularDataSources = "singular-data-sources"
	TargetPluralDataSources   = "plural-data-sources"

	// Test arguments
	PKGNameArg            = "PKG_NAME=internal/aws/logs"
	TestArgsArg           = "TESTARGS=-run=TestAccAWSLogsLogGroup_\\|TestAccAWSLogsLogGroupDataSource_"
	AccTestParallelismArg = "ACCTEST_PARALLELISM=3"

	// File permissions
	FilePermission = 0600
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	changes := []string{}

	// Check if we're in test mode
	testMode := os.Getenv(TestModeEnv) == "true"
	if testMode {
		fmt.Println("Running in TEST MODE - no Git operations or GitHub API calls will be made")
	}

	diags := checkEnv(ctx)
	if diags.HasError() {
		log.Println("Environment variable check failed:")
		for _, diag := range diags {
			fmt.Printf("Error: %s - %s\n", diag.Summary(), diag.Detail())
		}
		return fmt.Errorf("environment variable check failed")
	}

	// do we have to do anything about diags?

	// Comment out GitHub client creation to avoid requiring GitHub token
	// client, err := newGithubClient()
	// if err != nil {
	//     return fmt.Errorf("failed to create GitHub client: %w", err)
	// }

	// Use nil client instead

	branchName := fmt.Sprintf(BranchNameFormat, rand.Intn(BranchNameMaxRandom))
	fmt.Printf("Generated branch name: %s\n", branchName)

	var client *github.Client = nil
	filePaths, err := parseSchemaToStruct(UpdateFilePathsHCL, UpdateFilePaths{})
	if err != nil {
		return fmt.Errorf("failed to parse update file paths: %w", err)
	}

	isNewMap := make(map[string]bool)

	if !testMode {
		if err := execGit("checkout", "-b", branchName); err != nil {
			return fmt.Errorf("failed to create and checkout branch %s: %w", branchName, err)
		}
	} else {
		fmt.Printf("TEST MODE: Would create and checkout branch: %s\n", branchName)
	}

	/*
		matches, err := filepath.Glob(filePaths.AwsSchemas)
		if err != nil {
			return fmt.Errorf("failed to glob for old CloudFormation schemas: %w", err)
		}
		for _, file := range matches {
			if removeErr := os.Remove(file); removeErr != nil && !os.IsNotExist(removeErr) {
				return fmt.Errorf("failed to remove old CloudFormation schema %s: %w", file, removeErr)
			}
		}
	*/

	// open file and get to suppressionData

	if err := checkoutSchemas(ctx, filePaths.SuppressionCheckout); err != nil {
		return fmt.Errorf("failed to checkout schemas: %w", err)
	}
	currAllSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
	if err != nil {
		return fmt.Errorf("failed to parse current schemas: %w", err)
	}
	for i := range currAllSchemas.Resources {
		isNewMap[currAllSchemas.Resources[i].ResourceTypeName] = true
	}

	err = makeBuild(ctx, client, currAllSchemas, TargetSchemas, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to make schemas: %w", err)
	}
	if err := execGit("add", "-A"); err != nil {
		return fmt.Errorf("failed to git add files after schema refresh: %w", err)
	}

	currentDate := time.Now().Format(DateFormat)
	if err := execGit("commit", "-m", fmt.Sprintf(CommitMsgRefreshSchemas, currentDate)); err != nil {
		return fmt.Errorf("failed to commit schema refresh: %w", err)
	}
	// go run internal/provider/generators/allschemas/main.go > internal/provider/generators/allschemas/available_schemas.year-month-day.hcl

	// Diff Step Start

	lastDate, err := getLastDate()
	if err != nil {
		return fmt.Errorf("no previous schema file found")
	}
	tflog.Info(ctx, fmt.Sprintf("Last schema date: %s", lastDate))

	currentDate = time.Now().Format(DateFormat)
	newSchemas := allschemas.NewSchemaGeneration()
	err = writeSchemasToHCLFile(newSchemas, fmt.Sprintf("%s/%s%s%s", filePaths.AllSchemasDir, AvailableSchemasPrefix, currentDate, HCLExtension))
	if err != nil {
		return fmt.Errorf("failed to write new schemas to HCL file: %w", err)
	}
	// Parse schema from previous run
	lastSchemas, err := parseSchemaToStruct(fmt.Sprintf("%s/%s%s%s", filePaths.AllSchemasDir, AvailableSchemasPrefix, lastDate, HCLExtension), allschemas.AvailableSchemas{})
	if err != nil {
		return fmt.Errorf("failed to parse last schemas: %w", err)
	}

	currAllSchemas, err = diffSchemas(newSchemas, lastSchemas, &changes, filePaths)
	// Diff Step Stop

	if err != nil {
		return fmt.Errorf("failed to diff schemas: %w", err)
	}

	// Since we've disabled validation in diffSchemas, we should do it here
	err = validateResources(ctx, currAllSchemas, client, filePaths)
	if err != nil {
		return fmt.Errorf("failed to validate resources: %w", err)
	}

	if err := execGit("add", "-A"); err != nil {
		return fmt.Errorf("failed to git add files after schema diff: %w", err)
	}

	if err := execGit("commit", "-m", fmt.Sprintf(CommitMsgNewSchemas, currentDate)); err != nil {
		return fmt.Errorf("failed to commit new schemas: %w", err)
	}

	// Execute make resources command

	err = makeBuild(ctx, client, currAllSchemas, TargetSchemas, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to make new schemas: %w", err)
	}
	err = makeBuild(ctx, client, currAllSchemas, TargetResources, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to execute make resources: %w", err)
	}
	if err := execGit("add", "-A"); err != nil {
		return fmt.Errorf("failed to git add files after generating resource schemas: %w", err)
	}

	if err := execGit("commit", "-m", fmt.Sprintf(CommitMsgResourceSchemas, currentDate)); err != nil {
		return fmt.Errorf("failed to commit resource schemas: %w", err)
	}

	// Run make singular-data-sources plural-data-sources
	err = makeBuild(ctx, client, currAllSchemas, TargetSingularDataSources, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to update singular data sources: %w", err)
	}

	err = makeBuild(ctx, client, currAllSchemas, TargetPluralDataSources, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to update plural data sources: %w", err)
	}

	// Commit data source schema changes
	err = execGit("add", "-A")
	if err != nil {
		return fmt.Errorf("failed to git add data source files after updating data sources: %w", err)
	}
	err = execGit("commit", "-m", fmt.Sprintf(CommitMsgDataSourceSchemas, currentDate))
	if err != nil {
		return fmt.Errorf("failed to commit schemas after updating data sources: %w", err)
	}

	// Validate the provider
	err = execCommand("make", MakeBuildCmd)
	if err != nil {
		return fmt.Errorf("failed to build provider: %w", err)
	}

	// Run acceptance tests and capture output for PR description
	AcceptanceTestResults, err = RunAcceptanceTests()
	if err != nil {
		log.Printf("Warning: Acceptance tests had issues: %v", err)
		// We continue even if there are test failures to include results in PR
	}

	err = execCommand("make", MakeDocsAllCmd)
	if err != nil {
		return fmt.Errorf("failed to generate documentation: %w", err)
	}

	// Commit documentation changes
	err = execGit("add", "-A")
	if err != nil {
		return fmt.Errorf("failed to git add documentation files: %w", err)
	}
	err = execGit("commit", "-m", fmt.Sprintf(CommitMsgDocs, currentDate))
	if err != nil {
		return fmt.Errorf("failed to commit documentation: %w", err)
	}

	// Print all suppressions that happened during the process
	suppressions := strings.Builder{}
	for _, change := range changes {
		suppressions.WriteString(change)
		suppressions.WriteString("\n")
	}
	log.Println("Suppressions during process:\n" + suppressions.String())

	// If we have a GitHub client, create a PR with the test results

	fmt.Fprintf(os.Stdout, "env_token=production\n")
	fmt.Fprintf(os.Stdout, "suppressions<<EOF\n%sEOF\n", suppressions.String())
	// return nil

	// Create GitHubConfig
	config := NewGitHubConfig(client, filePaths.RepositoryLink, GetCurrentDate())

	if !testMode {
		_, err = submitOnGit(client, &changes, filePaths, AcceptanceTestResults, config.RepoOwner, config.RepoName)
		if err != nil {
			return fmt.Errorf("failed to submit PR: %w", err)
		}
	} else {
		fmt.Printf("TEST MODE: Would submit PR with %d changes to %s/%s\n", len(changes), config.RepoOwner, config.RepoName)
		fmt.Printf("TEST MODE: Changes would be: %v\n", changes)
	}

	err = makeChangelog(&changes, filePaths)
	if err != nil {
		return fmt.Errorf("failed to update changelog: %w", err)
	}
	return nil
}

// execCommand standardizes execution of non-git commands
func execCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func execGit(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getLastDate() (string, error) {
	// First get the file paths configuration
	filePaths, err := parseSchemaToStruct(UpdateFilePathsHCL, UpdateFilePaths{})
	if err != nil {
		return "", fmt.Errorf("failed to parse update file paths: %w", err)
	}

	files, err := os.ReadDir(filePaths.AllSchemasDir)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	var lastDate string
	for _, file := range files {
		name := file.Name()
		// Check if file matches the pattern "available_schemas.yyyy-mm-dd.hcl"
		if strings.HasPrefix(name, AvailableSchemasPrefix) && strings.HasSuffix(name, HCLExtension) {
			datePart := strings.TrimPrefix(name, AvailableSchemasPrefix)
			datePart = strings.TrimSuffix(datePart, HCLExtension)

			// Validate that it looks like a date
			if _, err := time.Parse(DateFormat, datePart); err == nil {
				if datePart > lastDate {
					lastDate = datePart
				}
			}
		}
	}

	return lastDate, nil
}

type UpdateFilePaths struct {
	RunMakesResourceLog      string `hcl:"run_makes_resource_log"`
	RunMakesOutput           string `hcl:"run_makes_output"`
	RunMakesErrors           string `hcl:"run_makes_errors"`
	SuppressionCheckout      string `hcl:"suppression_checkout"`
	AwsSchemas               string `hcl:"aws_schemas"`
	AllSchemasHCL            string `hcl:"all_schemas_hcl"`
	AllSchemasDir            string `hcl:"all_schemas_dir"`
	LastResource             string `hcl:"lastresource"`
	CloudFormationSchemasDir string `hcl:"cloudformation_schemas_dir"`
	RepositoryLink           string `hcl:"repository_link"`
}

// GetAcceptanceTestResults returns the captured acceptance test results.
// If no tests have been run yet, it returns an empty string.
func GetAcceptanceTestResults() string {
	return AcceptanceTestResults
}

// validateResources checks if each resource in the schema is provisionable and suppresses
// generation for those that are not. It also creates GitHub issues for non-provisionable resources.
func validateResources(ctx context.Context, currAllSchemas *allschemas.AllSchemas, client *github.Client, filePaths *UpdateFilePaths) error {
	// Create GitHubConfig
	config := NewGitHubConfig(client, filePaths.RepositoryLink, GetCurrentDate())

	for i := range currAllSchemas.Resources {
		flag, err := validateResourceType(ctx, currAllSchemas.Resources[i].CloudFormationTypeName)
		if err != nil {
			return fmt.Errorf("failed to check if resource %s is provisionable: %w", currAllSchemas.Resources[i].CloudFormationTypeName, err)
		}
		if !flag {
			currAllSchemas.Resources[i].SuppressResourceGeneration = true
			createIssue(ctx, currAllSchemas.Resources[i].CloudFormationTypeName, "Resource is not provisionable", config, filePaths.RepositoryLink)
		}
	}
	return nil
}
