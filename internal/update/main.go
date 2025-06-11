package update

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/go-github/v72/github"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func main() error {
	// Check environment variables
	diags := checkEnv()
	if diags.HasError() {
		ccdiag.DiagnosticsError(diags)
		os.Exit(1)
	}
	// do we have to do anything about diags?
	ctx := context.Background()
	client, err := newGithubClient()
	if err != nil {
		return fmt.Errorf("failed to create GitHub client: %w", err)
	}

	// open file and get to suppressionData
	var suppressionData string
	err = execGit("checkout", suppressionData)
	if err != nil {
		return fmt.Errorf("failed to checkout suppression data: %w", err)
	}
	currAllSchemas, err := parseSchemaToStruct("internal/provider/allschemas.hcl") // this should be
	if err != nil {
		return fmt.Errorf("failed to parse current schemas: %w", err)
	}
	err = makeSchemas(ctx, client, *currAllSchemas) // First Make Schema
	if err != nil {
		return fmt.Errorf("failed to make schemas: %w", err)
	}
	err = execGit("add", "-A")
	if err != nil {
		return fmt.Errorf("failed to git add files: %w", err)
	}
	currentDate := time.Now().Format("2006-01-02")
	err = execGit("commit", "-m", fmt.Sprintf("%s CloudFormation schemas in us-east-1; Refresh existing schemas.", currentDate))
	if err != nil {
		return fmt.Errorf("failed to commit schema refresh: %w", err)
	}
	// go run internal/provider/generators/allschemas/main.go > internal/provider/generators/allschemas/available_schemas.year-month-day.hcl

	// Diff Step Start

	lastDate, err := getLastDate()
	if err != nil {
		return fmt.Errorf("no previous schema file found")
	}
	fmt.Println("Last date found:", lastDate)

	newSchemas := allschemas.NewSchemaGeneration()
	err = writeSchemasToHCLFile(newSchemas, "internal/provider/generators/allschemas/available_schemas."+currentDate+".hcl")
	if err != nil {
		return fmt.Errorf("failed to write new schemas to HCL file: %w", err)
	}
	// Parse schema from previous run
	lastSchemas, err := parseSchemaToStruct("internal/provider/generators/allschemas/available_schemas." + lastDate + ".hcl")
	if err != nil {
		return fmt.Errorf("failed to parse last schemas: %w", err)
	}

	currAllSchemas, err = diffSchemas(lastSchemas, newSchemas, "internal/provider/allschemas.hcl")

	// Diff Step Stop

	if err != nil {
		return fmt.Errorf("failed to diff schemas: %w", err)
	}
	err = execGit("add", "-A")
	if err != nil {
		return fmt.Errorf("failed to git add new schema files: %w", err)
	}

	err = execGit("commit", "-m", fmt.Sprintf("%s CloudFormation schemas in us-east-1; New schemas.", currentDate))
	if err != nil {
		return fmt.Errorf("failed to commit new schemas: %w", err)
	}

	// Execute make resources command

	err = makeSchemas(ctx, client, *currAllSchemas)
	if err != nil {
		return fmt.Errorf("failed to make new schemas: %w", err)
	}
	err = makeResources(ctx, client, *currAllSchemas)
	if err != nil {
		return fmt.Errorf("failed to execute make resources: %w", err)
	}
	err = execGit("add", "-A")
	if err != nil {
		return fmt.Errorf("failed to git add resource schema files: %w", err)
	}
	err = execGit("commit", "-m", fmt.Sprintf("%s CloudFormation schemas in us-east-1; Generate Terraform resource schemas.", currentDate))
	if err != nil {
		return fmt.Errorf("failed to commit resource schemas: %w", err)
	}

	// Run make singular-data-sources plural-data-sources
	err = makeBuild(ctx, client, *currAllSchemas, "singular-data-sources")
	if err != nil {
		return fmt.Errorf("failed to update singular data sources: %w", err)
	}

	err = makeBuild(ctx, client, *currAllSchemas, "plural-data-sources")
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
	return cmd.Run()
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
