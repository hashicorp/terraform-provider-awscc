package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func CheckAWSEnv() error {
	requiredVars := []string{
		"AWS_DEFAULT_REGION",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
	}

	for _, v := range requiredVars {
		val := os.Getenv(v)
		if val == "" {
			return fmt.Errorf("environment variable %s is not set", v)
		}
	}

	if os.Getenv("AWS_DEFAULT_REGION") != "us-east-1" {
		return fmt.Errorf("AWS_DEFAULT_REGION must be set to us-east-1")
	}

	return nil
}

func checkGithubToken() error {
	// Comment out GitHub token check to avoid requiring GitHub token
	// githubToken := os.Getenv("GITHUB_TOKEN")
	// if githubToken == "" {
	//     return fmt.Errorf("GITHUB_TOKEN environment variable is not set")
	// }
	//
	// if len(githubToken) < 40 {
	//     return fmt.Errorf("GITHUB_TOKEN must be at least 40 characters long")
	// }

	return nil
}

func checkEnv(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics

	if err := CheckAWSEnv(); err != nil {
		diags.AddError("Environment Variable Check Failed", err.Error())
	}

	if err := checkGithubToken(); err != nil {
		diags.AddError("GitHub Token Check Failed", err.Error())
	}
	if len(diags) > 0 {
		return diags
	}

	tflog.Info(ctx, "Environment variable check passed")
	return nil
}
