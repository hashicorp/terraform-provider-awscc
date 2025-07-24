// Package main provides environment variable validation functionality.
// This file contains functions to verify that required AWS and GitHub
// credentials are properly configured before running the update process.
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CheckAWSEnv validates that all required AWS environment variables are set and configured correctly.
// It checks for AWS credentials, session token, and ensures the region is set to the expected value.
//
// Returns an error if any required environment variable is missing or has an invalid value.
func CheckAWSEnv() error {
	// List of required AWS environment variables for CloudFormation API access
	requiredVars := []string{
		"AWS_REGION",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
	}

	// Verify all required variables are set
	for _, v := range requiredVars {
		val := os.Getenv(v)
		if val == "" {
			return fmt.Errorf("environment variable %s is not set", v)
		}
	}

	// Ensure AWS region is set to the expected value for CloudFormation schema access
	if os.Getenv("AWS_REGION") != AWSRegion {
		return fmt.Errorf("AWS_REGION must be set to %s", AWSRegion)
	}

	return nil
}

// checkGithubToken validates the GitHub token environment variable.
// Currently commented out to allow running without GitHub integration during development.
//
// Returns an error if the GitHub token is missing or invalid.
func checkGithubToken() error {
	// GitHub token validation is currently disabled to allow development without GitHub integration
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return fmt.Errorf("GITHUB_TOKEN environment variable is not set")
	}
	if len(githubToken) < 40 {
		return fmt.Errorf("GITHUB_TOKEN must be at least 40 characters long")
	}

	return nil
}

// checkEnv performs comprehensive environment validation by checking both AWS and GitHub configurations.
// It logs validation results and aggregates any errors found during the check process.
//
// Parameters:
//   - ctx: Context for logging and cancellation
//
// Returns an error describing all validation failures, or nil if validation passes.
func checkEnv(ctx context.Context) error {
	var errors []string

	// Validate AWS environment configuration
	if err := CheckAWSEnv(); err != nil {
		errors = append(errors, fmt.Sprintf("Environment Variable Check Failed: %s", err.Error()))
	}

	// Validate GitHub token configuration (currently disabled)
	if err := checkGithubToken(); err != nil {
		errors = append(errors, fmt.Sprintf("GitHub Token Check Failed: %s", err.Error()))
	}

	// Return aggregated errors if any validation failed
	if len(errors) > 0 {
		return fmt.Errorf("environment checks failed:\n- %s", strings.Join(errors, "\n- "))
	}

	// Log successful validation
	tflog.Info(ctx, "Environment variable check passed")
	return nil
}
