// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"

	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
)

// artifactKind enumerates the generated code artifacts for a resource type.
type artifactKind string

const (
	artifactResource           artifactKind = "resource"
	artifactSingularDataSource artifactKind = "singular_data_source"
	artifactPluralDataSource   artifactKind = "plural_data_source"
)

// genArtifact is one code target to generate for a resource type. It carries
// everything the generators need: the Terraform type, package, output paths, and
// (for the resource) whether to also emit a ListResource.
type genArtifact struct {
	kind         artifactKind
	tfType       string // awscc_… (pluralized for the plural data source)
	packageName  string // e.g. "logs"
	pathSuffix   string // e.g. "aws/logs"
	codeFile     string // e.g. "log_group_resource_gen.go"
	testFile     string // e.g. "log_group_resource_gen_test.go"
	listResource bool   // resource only: also emit a ListResource
}

// plan is the fully-derived generation recipe for one resource_schema row.
type plan struct {
	cfType     string // e.g. "AWS::Logs::LogGroup"
	tfType     string // e.g. "awscc_logs_log_group"
	schemaFile string // cache path to the CloudFormation schema JSON
	artifacts  []genArtifact
}

// generationPlan derives, purely from an overlay row (no AWS, no I/O), the set
// of Terraform artifacts to generate plus their names and paths. It is a modern,
// tested reimplementation of the legacy Downloader.Schemas recipe logic:
// artifact selection follows the row's suppress_* flags, and a resource that has
// a plural data source also gets a ListResource.
func generationPlan(row resourceRow, prefix, cacheDir string) (plan, error) {
	org, svc, res, err := naming.ParseTerraformTypeName(row.ResourceTypeName)
	if err != nil {
		return plan{}, fmt.Errorf("parsing Terraform type name %q: %w", row.ResourceTypeName, err)
	}

	tfType := row.ResourceTypeName
	if prefix != "" {
		tfType = naming.CreateTerraformTypeName(prefix, svc, res)
	}

	schemaFile := row.CloudFormationSchemaPath
	if schemaFile == "" {
		schemaFile = schemaCachePath(cacheDir, row.CloudFormationTypeName)
	}

	pathSuffix := org + "/" + svc
	p := plan{cfType: row.CloudFormationTypeName, tfType: tfType, schemaFile: schemaFile}

	if !row.SuppressResourceGeneration {
		p.artifacts = append(p.artifacts, genArtifact{
			kind:         artifactResource,
			tfType:       tfType,
			packageName:  svc,
			pathSuffix:   pathSuffix,
			codeFile:     res + "_resource_gen.go",
			testFile:     res + "_resource_gen_test.go",
			listResource: !row.SuppressPluralDataSourceGeneration,
		})
	}
	if !row.SuppressSingularDataSourceGeneration {
		p.artifacts = append(p.artifacts, genArtifact{
			kind:        artifactSingularDataSource,
			tfType:      tfType,
			packageName: svc,
			pathSuffix:  pathSuffix,
			codeFile:    res + "_singular_data_source_gen.go",
			testFile:    res + "_singular_data_source_gen_test.go",
		})
	}
	if !row.SuppressPluralDataSourceGeneration {
		p.artifacts = append(p.artifacts, genArtifact{
			kind:        artifactPluralDataSource,
			tfType:      naming.Pluralize(tfType),
			packageName: svc,
			pathSuffix:  pathSuffix,
			codeFile:    res + "_plural_data_source_gen.go",
			testFile:    res + "_plural_data_source_gen_test.go",
		})
	}
	return p, nil
}
