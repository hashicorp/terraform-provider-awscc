package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"

	//"path/filepath"
	"strings"
	"time"

	"github.com/google/go-github/v72/github"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	diags := checkEnv(ctx)
	if diags.HasError() {
		fmt.Println("Environment variable check failed:")
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

	branchName := fmt.Sprintf("update-schemas-%d", rand.Intn(1000000))
	fmt.Printf("Generated branch name: %s\n", branchName)
	execGit("checkout", "-b", branchName)

	var client *github.Client = nil
	filePaths, err := parseSchemaToStruct("internal/update/update_filepaths.hcl", UpdateFilePaths{})
	if err != nil {
		return fmt.Errorf("failed to parse update file paths: %w", err)
	}

	matches, err := filepath.Glob(filePaths.AwsSchemas)
	if err != nil {
		return fmt.Errorf("failed to glob for old CloudFormation schemas: %w", err)
	}
	for _, file := range matches {
		if removeErr := os.Remove(file); removeErr != nil && !os.IsNotExist(removeErr) {
			return fmt.Errorf("failed to remove old CloudFormation schema %s: %w", file, removeErr)
		}
	}

	// open file and get to suppressionData

	checkoutSchemas(ctx, filePaths.SuppressionCheckout)
	currAllSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
	if err != nil {
		return fmt.Errorf("failed to parse current schemas: %w", err)
	}
	err = makeBuild(ctx, client, currAllSchemas, "schemas", filePaths)
	if err != nil {
		return fmt.Errorf("failed to make schemas: %w", err)
	}
	execGit("add", "-A")

	currentDate := time.Now().Format("2006-01-02")
	execGit("commit", "-m", fmt.Sprintf("%s CloudFormation schemas in us-east-1; Refresh existing schemas.", currentDate))
	// go run internal/provider/generators/allschemas/main.go > internal/provider/generators/allschemas/available_schemas.year-month-day.hcl

	// Diff Step Start

	lastDate, err := getLastDate()
	if err != nil {
		return fmt.Errorf("no previous schema file found")
	}
	tflog.Info(ctx, fmt.Sprintf("Last schema date: %s", lastDate))

	newSchemas := allschemas.NewSchemaGeneration()
	err = writeSchemasToHCLFile(newSchemas, "internal/provider/generators/allschemas/available_schemas."+currentDate+".hcl")
	if err != nil {
		return fmt.Errorf("failed to write new schemas to HCL file: %w", err)
	}
	// Parse schema from previous run
	lastSchemas, err := parseSchemaToStruct("internal/provider/generators/allschemas/available_schemas."+lastDate+".hcl", allschemas.AvailableSchemas{})
	if err != nil {
		return fmt.Errorf("failed to parse last schemas: %w", err)
	}

	currAllSchemas, err = diffSchemas(ctx, lastSchemas, newSchemas, filePaths.AllSchemasHCL)

	// Diff Step Stop

	if err != nil {
		return fmt.Errorf("failed to diff schemas: %w", err)
	}
	execGit("add", "-A")

	execGit("commit", "-m", fmt.Sprintf("%s CloudFormation schemas in us-east-1; New schemas.", currentDate))

	// Execute make resources command

	err = makeBuild(ctx, client, currAllSchemas, "schemas", filePaths)
	if err != nil {
		return fmt.Errorf("failed to make new schemas: %w", err)
	}
	err = makeBuild(ctx, client, currAllSchemas, "resources", filePaths)
	if err != nil {
		return fmt.Errorf("failed to execute make resources: %w", err)
	}
	execGit("add", "-A")

	execGit("commit", "-m", fmt.Sprintf("%s CloudFormation schemas in us-east-1; Generate Terraform resource schemas.", currentDate))

	// Run make singular-data-sources plural-data-sources
	err = makeBuild(ctx, client, currAllSchemas, "singular-data-sources", filePaths)
	if err != nil {
		return fmt.Errorf("failed to update singular data sources: %w", err)
	}

	err = makeBuild(ctx, client, currAllSchemas, "plural-data-sources", filePaths)
	if err != nil {
		return fmt.Errorf("failed to update plural data sources: %w", err)
	}

	// Commit data source schema changes
	err = execGit("add", "-A")
	if err != nil {
		return fmt.Errorf("failed to git add data source files after updating data sources: %w", err)
	}
	err = execGit("commit", "-m", fmt.Sprintf("%s CloudFormation schemas in us-east-1; Generate Terraform data source schemas.", currentDate))
	if err != nil {
		return fmt.Errorf("failed to commit schemas after updating data sources: %w", err)
	}

	// Validate the provider
	err = execCommand("make", "build")
	if err != nil {
		return fmt.Errorf("failed to build provider: %w", err)
	}

	err = execCommand("make", "testacc", "PKG_NAME=internal/aws/logs", "TESTARGS=-run=TestAccAWSLogsLogGroup_\\|TestAccAWSLogsLogGroupDataSource_", "ACCTEST_PARALLELISM=3")
	if err != nil {
		return fmt.Errorf("failed to run acceptance tests: %w", err)
	}

	err = execCommand("make", "docs-all")
	if err != nil {
		return fmt.Errorf("failed to generate documentation: %w", err)
	}

	// Commit documentation changes
	err = execGit("add", "-A")
	if err != nil {
		return fmt.Errorf("failed to git add documentation files: %w", err)
	}
	err = execGit("commit", "-m", fmt.Sprintf("%s Run 'make docs-all'.", currentDate))
	if err != nil {
		return fmt.Errorf("failed to commit documentation: %w", err)
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
	_ = cmd.Run()
	// Ignore all errors from git commands
	return nil
}

func newGithubClient() (*github.Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable is not set")
	}
	client := github.NewClient(nil).WithAuthToken(token)
	return client, nil
}

func getLastDate() (string, error) {
	files, err := os.ReadDir("internal/provider/generators/allschemas")
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	var lastDate string
	for _, file := range files {
		name := file.Name()
		// Check if file matches the pattern "available_schemas.yyyy-mm-dd.hcl"
		if strings.HasPrefix(name, "available_schemas.") && strings.HasSuffix(name, ".hcl") {
			datePart := strings.TrimPrefix(name, "available_schemas.")
			datePart = strings.TrimSuffix(datePart, ".hcl")

			// Validate that it looks like a date
			if _, err := time.Parse("2006-01-02", datePart); err == nil {
				if datePart > lastDate {
					lastDate = datePart
				}
			}
		}
	}

	return lastDate, nil
}

func parseUpdateConfig() (*UpdateFilePaths, error) {
	var filename = "internal/update/update_files.hcl"
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config UpdateFilePaths
	err = hclsimple.Decode(filename, data, nil, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

type UpdateFilePaths struct {
	RunMakesResourceLog string `hcl:"run_makes_resource_log"`
	RunMakesOutput      string `hcl:"run_makes_output"`
	RunMakesErrors      string `hcl:"run_makes_errors"`
	SuppressionCheckout string `hcl:"suppression_checkout"`
	AwsSchemas          string `hcl:"aws_schemas"`
	AllSchemasHCL       string `hcl:"all_schemas_hcl"`
}
