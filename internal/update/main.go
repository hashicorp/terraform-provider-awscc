// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

// AcceptanceTestResults stores the output of acceptance tests for inclusion in pull request descriptions.
// This global variable is populated by RunAcceptanceTests and used when creating pull requests.

// Constants for file paths, patterns, and configuration values used throughout the update process.
const (
	// Configuration file paths
	UpdateFilePathsHCL = "internal/update/update_filepaths.hcl"

	// Schema file patterns and naming
	AvailableSchemasPrefix      = "available_schemas."
	HCLExtension                = ".hcl"
	AvailableSchemasFilePattern = "available_schemas.%s.hcl"

	// Git commit message templates
	CloudFormationRegion       = "CloudFormation schemas in " + AWSRegion
	CommitMsgRefreshSchemas    = "%s " + CloudFormationRegion + "; Refresh existing schemas."
	CommitMsgNewSchemas        = "%s " + CloudFormationRegion + "; New schemas."
	CommitMsgResourceSchemas   = "%s " + CloudFormationRegion + "; Generate Terraform resource schemas."
	CommitMsgDataSourceSchemas = "%s " + CloudFormationRegion + "; Generate Terraform data source schemas."
	CommitMsgDocs              = "%s Run 'make docs-all'."

	// Branch naming configuration
	BranchNameFormat    = "update-schemas-%d"
	BranchNameMaxRandom = 1000000

	// Time and date formatting
	DateFormat = "2006-01-02"

	// Environment variable names
	GithubTokenEnv = "GITHUB_TOKEN"

	// Make command names
	MakeBuildCmd   = "build"
	MakeTestAccCmd = "testacc"
	MakeDocsAllCmd = "docs-all"

	// Make target names for different build types
	TargetSchemas             = "schemas"
	TargetResources           = "resources"
	TargetSingularDataSources = "singular-data-sources"
	TargetPluralDataSources   = "plural-data-sources"

	// Test execution parameters
	PKGNameArg            = "PKG_NAME=internal/aws/logs"
	TestArgsArg           = "TESTARGS=-run=TestAccAWSLogsLogGroup_\\|TestAccAWSLogsLogGroupDataSource_"
	AccTestParallelismArg = "ACCTEST_PARALLELISM=3"

	// File system permissions
	FilePermission = 0600

	// GitHub repository configuration
	DefaultRepoOwner = "HashiCorp"
	DefaultRepoName  = "terraform-provider-awscc"
	GitHubURLPrefix  = "https://github.com/"

	// AWS configuration
	AWSRegion = "us-east-1"
)

// main is the entry point for the schema update process.
// It calls run() and handles any errors by printing them to stderr and exiting with status 1.
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// run orchestrates the complete schema update workflow:
// 1. Parses configuration and initializes GitHub setup
// 2. Creates a new branch for changes
// 3. Refreshes existing schemas
// 4. Identifies and processes new schemas
// 5. Generates resources and data sources
// 6. Runs tests and documentation generation
// 7. Creates a pull request with the changes
func run() error {
	ctx := context.Background()
	changes := []string{}

	// Parse configuration file to get file paths and repository information
	filePaths, err := parseSchemaToStruct(UpdateFilePathsHCL, UpdateFilePaths{})
	if err != nil {
		return fmt.Errorf("failed to parse update file paths: %w", err)
	}

	// Initialize GitHub configuration with all GitHub-related setup
	currentDate := GetCurrentDate()
	config, err := NewGitHubConfig(filePaths.RepositoryLink, currentDate)
	if err != nil {
		return fmt.Errorf("failed to initialize GitHub configuration: %w", err)
	}

	// Create a unique branch name for this update run
	branchName := fmt.Sprintf(BranchNameFormat, rand.Intn(BranchNameMaxRandom))

	//Update version file
	err = updateVersionFile(filePaths)
	if err != nil {
		return fmt.Errorf("failed to update version file: %w", err)
	}

	// Run make tools for tool dependencies
	log.Printf("Running make tools")
	err = execCommand("make", "tools")
	if err != nil {
		return fmt.Errorf("failed to run 'make tools': %w", err)
	}

	// Track which resources are new for suppression logic
	isNewMap := make(map[string]bool)

	// Create and checkout a new feature branch
	if err := execGit("checkout", "-b", branchName); err != nil {
		return fmt.Errorf("failed to create and checkout branch %s: %w", branchName, err)
	}

	// Remove existing CloudFormation schema files to start fresh
	matches, err := filepath.Glob(filePaths.AwsSchemas)
	if err != nil {
		return fmt.Errorf("failed to glob for old CloudFormation schemas: %w", err)
	}
	for _, file := range matches {
		if removeErr := os.Remove(file); removeErr != nil && !os.IsNotExist(removeErr) {
			return fmt.Errorf("failed to remove old CloudFormation schema %s: %w", file, removeErr)
		}
	}

	// Checkout fresh schemas and load current schema configuration
	if err := checkoutSchemas(filePaths.SuppressionCheckout); err != nil && strings.Contains(err.Error(), "not found") {
		return fmt.Errorf("failed to checkout schemas: %w", err)
	}
	currAllSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
	if err != nil {
		return fmt.Errorf("failed to parse current schemas: %w", err)
	}

	// Mark existing resources in the isNewMap for suppression logic
	for i := range currAllSchemas.Resources {
		isNewMap[currAllSchemas.Resources[i].ResourceTypeName] = true
	}

	// Step 1: Refresh existing schemas
	err = makeBuild(ctx, config, currAllSchemas, TargetSchemas, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to make schemas: %w", err)
	}
	if err := execGit("add", "-A"); err != nil {
		return fmt.Errorf("failed to git add files after schema refresh: %w", err)
	}

	currentDate = time.Now().Format(DateFormat)
	if err := execGit("commit", "-m", fmt.Sprintf(CommitMsgRefreshSchemas, currentDate)); err != nil {
		return fmt.Errorf("failed to commit schema refresh: %w", err)
	}

	// Step 2: Generate and compare schemas to identify new/changed resources
	// Find the most recent schema file to compare against
	lastDate, err := getLastDate()
	if err != nil {
		return fmt.Errorf("no previous schema file found")
	}
	tflog.Info(ctx, fmt.Sprintf("Last schema date: %s", lastDate))

	// Generate current schemas and write to dated file
	currentDate = time.Now().Format(DateFormat)
	newSchemas := allschemas.NewSchemaGeneration()
	err = writeSchemasToHCLFile(newSchemas, fmt.Sprintf("%s/%s%s%s", filePaths.AllSchemasDir, AvailableSchemasPrefix, currentDate, HCLExtension))
	if err != nil {
		return fmt.Errorf("failed to write new schemas to HCL file: %w", err)
	}

	// Parse and compare with previous schemas to identify changes
	lastSchemas, err := parseSchemaToStruct(fmt.Sprintf("%s/%s%s%s", filePaths.AllSchemasDir, AvailableSchemasPrefix, lastDate, HCLExtension), allschemas.AvailableSchemas{})
	if err != nil {
		return fmt.Errorf("failed to parse last schemas: %w", err)
	}

	currAllSchemas, err = diffSchemas(newSchemas, lastSchemas, &changes, filePaths)
	if err != nil {
		return fmt.Errorf("failed to diff schemas: %w", err)
	}

	// Step 3: Validate resources and handle suppressions
	err = validateResources(ctx, currAllSchemas, config, filePaths)
	if err != nil {
		log.Println(fmt.Errorf("failed to validate resources: %w", err))
		log.Println("continuing with schema update despite validation errors. please review the logs for details.")
	}

	// Commit the new schema changes
	if err := execGit("add", "-A"); err != nil {
		return fmt.Errorf("failed to git add files after schema diff: %w", err)
	}

	if err := execGit("commit", "-m", fmt.Sprintf(CommitMsgNewSchemas, currentDate)); err != nil {
		return fmt.Errorf("failed to commit new schemas: %w", err)
	}

	// Step 4: Generate Terraform resources and data sources
	err = makeBuild(ctx, config, currAllSchemas, TargetSchemas, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to make new schemas: %w", err)
	}
	err = makeBuild(ctx, config, currAllSchemas, TargetResources, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to execute make resources: %w", err)
	}
	if err := execGit("add", "-A"); err != nil {
		return fmt.Errorf("failed to git add files after generating resource schemas: %w", err)
	}

	if err := execGit("commit", "-m", fmt.Sprintf(CommitMsgResourceSchemas, currentDate)); err != nil {
		return fmt.Errorf("failed to commit resource schemas: %w", err)
	}

	// Generate data sources (both singular and plural)
	err = makeBuild(ctx, config, currAllSchemas, TargetSingularDataSources, &changes, filePaths, isNewMap)
	if err != nil {
		return fmt.Errorf("failed to update singular data sources: %w", err)
	}

	err = makeBuild(ctx, config, currAllSchemas, TargetPluralDataSources, &changes, filePaths, isNewMap)
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

	// Step 5: Build and test the provider
	// Validate the provider builds successfully
	log.Printf("Building provider with 'make %s'...", MakeBuildCmd)
	err = execCommand("make", MakeBuildCmd)
	if err != nil {
		return fmt.Errorf("failed to build provider: %w", err)
	}

	// Generate updated documentation]
	log.Printf("Generating documentation with 'make %s'...", MakeDocsAllCmd)
	err = execCommand("make", MakeDocsAllCmd)
	if err != nil {
		return fmt.Errorf("failed to generate documentation: %w", err)
	}

	err = trimAllSchemas(filePaths)
	if err != nil {
		return fmt.Errorf("failed to trim allSchemas.hcl: %w", err)
	}

	// Run acceptance tests and capture output for PR description
	log.Printf("Running acceptance tests with 'make %s'...", MakeTestAccCmd)
	AcceptanceTestResults, err := RunAcceptanceTests()
	if err != nil {
		log.Printf("Warning: Acceptance tests had issues: %v", err)
		// We continue even if there are test failures to include results in PR
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

	// Step 6: Prepare output and create pull request
	// Collect and log all suppressions that occurred during the process
	suppressions := strings.Builder{}
	for _, change := range changes {
		suppressions.WriteString(change)
		suppressions.WriteString("\n")
	}
	log.Println("Suppressions during process:\n" + suppressions.String())

	// Update the configuration with current date and submit pull request
	config.CurrentDate = GetCurrentDate()

	// Update the changelog with the changes
	fullChanges, err := makeChangelog(&changes, filePaths)
	if err != nil {
		return fmt.Errorf("failed to update changelog: %w", err)
	}

	// Commit the changelog changes
	if err := execGit("add", "CHANGELOG.md"); err != nil {
		return fmt.Errorf("failed to git add changelog: %w", err)
	}
	if err := execGit("commit", "-m", fmt.Sprintf("Update changelog for %s", currentDate)); err != nil {
		return fmt.Errorf("failed to commit changelog: %w", err)
	}

	_, err = submitOnGit(config, fullChanges, filePaths, AcceptanceTestResults, config.RepoOwner, config.RepoName, branchName)
	if err != nil {
		return fmt.Errorf("failed to submit PR: %w", err)
	}
	return nil
}

// execCommand executes a non-git command with standardized output handling.
// It redirects both stdout and stderr to the current process's output streams.
func execCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// execGit executes a git command with the provided arguments.
// It redirects both stdout and stderr to the current process's output streams.
func execGit(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// getLastDate finds the most recent dated schema file in the schemas directory.
// It looks for files matching the pattern "available_schemas.yyyy-mm-dd.hcl" and
// returns the date string from the most recent file.
func getLastDate() (string, error) {
	// Parse configuration to get the schemas directory path
	filePaths, err := parseSchemaToStruct(UpdateFilePathsHCL, UpdateFilePaths{})
	if err != nil {
		return "", fmt.Errorf("failed to parse update file paths: %w", err)
	}

	// Read the schemas directory to find available schema files
	files, err := os.ReadDir(filePaths.AllSchemasDir)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	var lastDate string
	for _, file := range files {
		name := file.Name()
		// Check if file matches the expected schema file pattern
		if strings.HasPrefix(name, AvailableSchemasPrefix) && strings.HasSuffix(name, HCLExtension) {
			datePart := strings.TrimPrefix(name, AvailableSchemasPrefix)
			datePart = strings.TrimSuffix(datePart, HCLExtension)

			// Validate that the extracted part is a valid date
			if _, err := time.Parse(DateFormat, datePart); err == nil {
				if datePart > lastDate {
					lastDate = datePart
				}
			}
		}
	}

	return lastDate, nil
}

// UpdateFilePaths defines the configuration structure for file paths used during the update process.
// This struct is populated by parsing the HCL configuration file.
type UpdateFilePaths struct {
	RunMakesResourceLog      string `hcl:"run_makes_resource_log"`     // Path to resource generation log file
	RunMakesOutput           string `hcl:"run_makes_output"`           // Path to make command output file
	RunMakesErrors           string `hcl:"run_makes_errors"`           // Path to make command error log file
	SuppressionCheckout      string `hcl:"suppression_checkout"`       // Path to suppression configuration file
	AwsSchemas               string `hcl:"aws_schemas"`                // Glob pattern for AWS schema files
	AllSchemasHCL            string `hcl:"all_schemas_hcl"`            // Path to all schemas HCL file
	AllSchemasDir            string `hcl:"all_schemas_dir"`            // Directory containing schema files
	LastResource             string `hcl:"lastresource"`               // Path to file tracking last processed resource
	CloudFormationSchemasDir string `hcl:"cloudformation_schemas_dir"` // Directory for CloudFormation schemas
	RepositoryLink           string `hcl:"repository_link"`            // GitHub repository URL
	Version                  string `hcl:"version_file"`               // Version file path
}

// validateResources checks if each resource in the schema is provisionable through CloudFormation.
// For resources that are not provisionable, it marks them for suppression to avoid generation failures.
// It also creates GitHub issues to track non-provisionable resources for future investigation.
//
// Parameters:
//   - ctx: Context for logging and API calls
//   - currAllSchemas: Schema configuration containing resources to validate
//   - config: GitHub configuration for issue creation (can be nil)
//   - filePaths: Configuration containing repository information
//
// Returns an error if validation fails for any resource.
func validateResources(ctx context.Context, currAllSchemas *allschemas.AllSchemas, config *GitHubConfig, filePaths *UpdateFilePaths) error {
	isSuppressed := parseCheckoutList(filePaths)
	timer := 2
	for i := 0; i < len(currAllSchemas.Resources); i++ {
		if currAllSchemas.Resources[i].SuppressResourceGeneration || isSuppressed[currAllSchemas.Resources[i].CloudFormationTypeName] {
			log.Printf("Skipping validation for suppressed resource %s", currAllSchemas.Resources[i].CloudFormationTypeName)
			continue
		}
		// Check if the resource type can be provisioned via CloudFormation
		flag, err := validateResourceType(ctx, currAllSchemas.Resources[i].CloudFormationTypeName)
		if err != nil && !strings.Contains(err.Error(), "TypeNotFoundException") {
			if strings.Contains(err.Error(), "api error Throttling: Rate exceeded") {
				log.Printf("Throttling error encountered, retrying in %d seconds...", timer)
				time.Sleep(time.Duration(timer) * time.Second)
				timer *= 2
				i-- // Retry the same resource
				continue
			}
			return fmt.Errorf("failed to check if resource %s is provisionable: %w", currAllSchemas.Resources[i].CloudFormationTypeName, err)
		}

		timer = 2

		// Suppress resources that are not provisionable
		if !flag || (err != nil && strings.Contains(err.Error(), "TypeNotFoundException")) {
			log.Printf("Resource %s is not provisionable via CloudFormation, marking for suppression via checkout", currAllSchemas.Resources[i].CloudFormationTypeName)
			err := addSchemaToCheckout(currAllSchemas.Resources[i].CloudFormationTypeName, filePaths)
			if err != nil {
				return fmt.Errorf("failed to add resource to checkout file: %w", err)
			}

			// Create GitHub issue for tracking if client is available
			if config != nil && config.Client != nil {
				link, err := createIssue(ctx, currAllSchemas.Resources[i].CloudFormationTypeName, "Resource is not provisionable", config, filePaths.RepositoryLink)
				if err != nil {
					log.Printf("Failed to create GitHub issue for resource %s: %v", currAllSchemas.Resources[i].CloudFormationTypeName, err)
					log.Printf("Please create an issue manually for resource %s not being provisionable", currAllSchemas.Resources[i].CloudFormationTypeName)
				}
				log.Printf("Created GitHub issue for resource %s: %s", currAllSchemas.Resources[i].CloudFormationTypeName, link)
			} else {
				tflog.Info(ctx, "Skipping GitHub issue creation (no client)", map[string]any{
					"resource": currAllSchemas.Resources[i].CloudFormationTypeName,
				})
			}
		}
	}
	return nil
}

func parseCheckoutList(filePaths *UpdateFilePaths) map[string]bool {
	result := make(map[string]bool)
	data, err := os.ReadFile(filePaths.SuppressionCheckout)
	if err != nil {
		log.Printf("Failed to read suppression checkout file: %v", err)
		return result
	}
	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			resource := convertJSONResourceToCloudFormationTypeName(line)
			result[resource] = true
		}
	}
	return result
}

func convertJSONResourceToCloudFormationTypeName(line string) string {
	// Convert JSON resource name to Terraform type name
	base := filepath.Base(line)
	resource := strings.TrimSuffix(base, filepath.Ext(base))
	resource = strings.TrimSuffix(resource, ".json")
	resource = strings.ReplaceAll(resource, "_", "::")
	return resource
}
