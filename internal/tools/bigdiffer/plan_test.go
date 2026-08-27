// SPDX-License-Identifier: MPL-2.0

package main

import (
	"testing"
)

func TestGenerationPlan(t *testing.T) {
	t.Parallel()

	const (
		prefix   = "awscc"
		cacheDir = "../service/cloudformation/schemas"
	)

	// artifact summary for compact assertions: kind -> tfType.
	kinds := func(p plan) map[artifactKind]string {
		m := make(map[artifactKind]string, len(p.artifacts))
		for _, a := range p.artifacts {
			m[a.kind] = a.tfType
		}
		return m
	}

	t.Run("full (resource + both data sources)", func(t *testing.T) {
		t.Parallel()
		p, err := generationPlan(resourceRow{
			ResourceTypeName:       "aws_logs_log_group",
			CloudFormationTypeName: "AWS::Logs::LogGroup",
		}, prefix, cacheDir)
		if err != nil {
			t.Fatalf("generationPlan: %v", err)
		}
		if p.tfType != "awscc_logs_log_group" {
			t.Errorf("tfType = %q, want awscc_logs_log_group", p.tfType)
		}
		if want := cacheDir + "/AWS_Logs_LogGroup.json"; p.schemaFile != want {
			t.Errorf("schemaFile = %q, want %q", p.schemaFile, want)
		}
		if len(p.artifacts) != 3 {
			t.Fatalf("artifacts = %d, want 3 (%+v)", len(p.artifacts), p.artifacts)
		}
		got := kinds(p)
		if got[artifactResource] != "awscc_logs_log_group" {
			t.Errorf("resource tfType = %q", got[artifactResource])
		}
		if got[artifactSingularDataSource] != "awscc_logs_log_group" {
			t.Errorf("singular tfType = %q", got[artifactSingularDataSource])
		}
		if got[artifactPluralDataSource] != "awscc_logs_log_groups" {
			t.Errorf("plural tfType = %q, want awscc_logs_log_groups", got[artifactPluralDataSource])
		}
		// File names + package/path derive from the CFN resource segment.
		for _, a := range p.artifacts {
			if a.packageName != "logs" || a.pathSuffix != "aws/logs" {
				t.Errorf("%s: package=%q pathSuffix=%q, want logs / aws/logs", a.kind, a.packageName, a.pathSuffix)
			}
			if a.kind == artifactResource {
				if a.codeFile != "log_group_resource_gen.go" || a.testFile != "log_group_resource_gen_test.go" {
					t.Errorf("resource files = %q/%q", a.codeFile, a.testFile)
				}
				if !a.listResource {
					t.Errorf("resource should generate a ListResource when a plural DS exists")
				}
			}
		}
	})

	t.Run("suppress_plural drops plural DS and ListResource", func(t *testing.T) {
		t.Parallel()
		p, err := generationPlan(resourceRow{
			ResourceTypeName:                   "aws_appstream_stack_fleet_association",
			CloudFormationTypeName:             "AWS::AppStream::StackFleetAssociation",
			SuppressPluralDataSourceGeneration: true,
		}, prefix, cacheDir)
		if err != nil {
			t.Fatalf("generationPlan: %v", err)
		}
		got := kinds(p)
		if _, ok := got[artifactPluralDataSource]; ok {
			t.Errorf("plural DS should be suppressed, got %+v", p.artifacts)
		}
		if _, ok := got[artifactResource]; !ok {
			t.Errorf("resource should still be generated")
		}
		for _, a := range p.artifacts {
			if a.kind == artifactResource && a.listResource {
				t.Errorf("no ListResource when plural DS is suppressed")
			}
		}
	})

	t.Run("all suppressed yields no artifacts", func(t *testing.T) {
		t.Parallel()
		p, err := generationPlan(resourceRow{
			ResourceTypeName:                     "aws_emr_instance_group_config",
			CloudFormationTypeName:               "AWS::EMR::InstanceGroupConfig",
			SuppressResourceGeneration:           true,
			SuppressSingularDataSourceGeneration: true,
			SuppressPluralDataSourceGeneration:   true,
		}, prefix, cacheDir)
		if err != nil {
			t.Fatalf("generationPlan: %v", err)
		}
		if len(p.artifacts) != 0 {
			t.Errorf("artifacts = %d, want 0 (%+v)", len(p.artifacts), p.artifacts)
		}
	})

	t.Run("explicit schema path overrides the cache path", func(t *testing.T) {
		t.Parallel()
		p, err := generationPlan(resourceRow{
			ResourceTypeName:         "aws_logs_log_group",
			CloudFormationTypeName:   "AWS::Logs::LogGroup",
			CloudFormationSchemaPath: "custom/path/schema.json",
		}, prefix, cacheDir)
		if err != nil {
			t.Fatalf("generationPlan: %v", err)
		}
		if p.schemaFile != "custom/path/schema.json" {
			t.Errorf("schemaFile = %q, want the explicit override", p.schemaFile)
		}
	})

	t.Run("malformed Terraform type name errors", func(t *testing.T) {
		t.Parallel()
		if _, err := generationPlan(resourceRow{
			ResourceTypeName:       "notavalidname",
			CloudFormationTypeName: "AWS::Foo::Bar",
		}, prefix, cacheDir); err == nil {
			t.Errorf("expected an error for a malformed Terraform type name")
		}
	})
}
