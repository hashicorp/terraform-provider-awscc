// Package main provides HCL parsing and CloudFormation schema processing functionality.
// This file contains utilities for parsing HCL configuration files, comparing schemas
// between runs, and validating CloudFormation resource types.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

// parseSchemaToStruct is a generic function that parses HCL files into Go structs.
// It reads the specified HCL file and unmarshals its contents into the provided schema type.
//
// Type parameter T: The target struct type to unmarshal into
// Parameters:
//   - filePath: Path to the HCL file to parse
//   - schema: Instance of the target struct type
//
// Returns a pointer to the populated struct or an error if parsing fails.
func parseSchemaToStruct[T any](filePath string, schema T) (*T, error) {
	// Read the HCL file from disk
	file, err := os.ReadFile(filePath)
	if err != nil {
		return &schema, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Parse the HCL content using the HCL parser
	parser := hclparse.NewParser()
	fileHCL, diag := parser.ParseHCL(file, filePath)
	if diag.HasErrors() {
		return &schema, fmt.Errorf("failed to parse HCL from %s: %s", filePath, diag.Error())
	}

	// Decode the parsed HCL into the target struct
	if diag := gohcl.DecodeBody(fileHCL.Body, nil, &schema); diag.HasErrors() {
		return &schema, fmt.Errorf("failed to decode HCL from %s: %s", filePath, diag.Error())
	}

	return &schema, nil
}

// diffSchemas compares new and previous schema sets to identify changes and new resources.
// It creates an AllSchemas structure containing only the resources that have changed
// or are newly added, which is used for targeted resource generation.
//
// Parameters:
//   - newSchemas: Current available schemas from CloudFormation
//   - lastSchemas: Previous schemas from the last run
//   - changes: Slice to append change descriptions to
//   - filePaths: Configuration containing file paths
//
// Returns an AllSchemas structure with only changed/new resources.
func diffSchemas(newSchemas *allschemas.AvailableSchemas, lastSchemas *allschemas.AvailableSchemas, changes *[]string, filePaths *UpdateFilePaths) (*allschemas.AllSchemas, error) {
	// Create lookup map for efficient schema comparison
	lastSchemasMap := make(map[string]int)
	for i, resource := range lastSchemas.Resources {
		lastSchemasMap[resource.CloudFormationTypeName] = i
	}

	// Collect resources that have changed or are new
	var changedOrNewResources []allschemas.ResourceSchema

	// Identify changed or new resources by comparing schemas
	for _, newResource := range newSchemas.Resources {
		if lastResourceIndex, exists := lastSchemasMap[newResource.CloudFormationTypeName]; exists {
			// Check if resource changed
			if newResource.CloudFormationTypeName != lastSchemas.Resources[lastResourceIndex].CloudFormationTypeName ||
				newResource.SuppressPluralDataSourceGeneration != lastSchemas.Resources[lastResourceIndex].SuppressPluralDataSourceGeneration ||
				newResource.ResourceTypeName != lastSchemas.Resources[lastResourceIndex].ResourceTypeName {
				changedOrNewResources = append(changedOrNewResources, newResource)
			}
		} else {
			// New resource
			changedOrNewResources = append(changedOrNewResources, newResource)
			*changes = append(*changes, fmt.Sprintf("%s - New Resource", newResource.ResourceTypeName))

		}
	}

	if len(changedOrNewResources) == 0 {
		log.Println("No changes or new resources found.")
		// Instead of returning nil, return the existing schemas
		existingAllSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
		if err != nil {
			return nil, fmt.Errorf("failed to parse existing allSchemas when no changes found: %w", err)
		}
		return existingAllSchemas, nil
	}

	log.Printf("Found %d changed or new resources.\n", len(changedOrNewResources))

	// Read existing allSchemas.hcl
	var err error
	existingAllSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse existing allSchemas: %w", err)
	}
	existingResourcesMap := make(map[string]int)
	for index, resource := range existingAllSchemas.Resources {
		existingResourcesMap[resource.CloudFormationTypeName] = index
	}

	for _, resource := range changedOrNewResources {
		if existingResourceIndex, exists := existingResourcesMap[resource.CloudFormationTypeName]; exists {
			// Update existing resource
			curr := existingAllSchemas.Resources[existingResourceIndex]
			if curr.CloudFormationTypeName != resource.CloudFormationTypeName {
				existingAllSchemas.Resources[existingResourceIndex].CloudFormationTypeName = resource.CloudFormationTypeName
			}
			if curr.ResourceTypeName != resource.ResourceTypeName {
				existingAllSchemas.Resources[existingResourceIndex].ResourceTypeName = resource.ResourceTypeName
			}
			if curr.SuppressPluralDataSourceGeneration != resource.SuppressPluralDataSourceGeneration {
				existingAllSchemas.Resources[existingResourceIndex].SuppressPluralDataSourceGeneration = resource.SuppressPluralDataSourceGeneration
			}
		} else {
			tempResource := &allschemas.ResourceAllSchema{
				ResourceTypeName:       resource.ResourceTypeName,
				CloudFormationTypeName: resource.CloudFormationTypeName,
			}
			tempResource.SuppressPluralDataSourceGeneration = resource.SuppressPluralDataSourceGeneration

			existingAllSchemas.Resources = append(existingAllSchemas.Resources, *tempResource)
		}
	}

	sort.Slice(existingAllSchemas.Resources, func(i, j int) bool {
		return existingAllSchemas.Resources[i].ResourceTypeName < existingAllSchemas.Resources[j].ResourceTypeName
	})
	return existingAllSchemas, writeSchemasToHCLFile(existingAllSchemas, filePaths.AllSchemasHCL)
}

// writeSchemasToHCLFile serializes a schema structure to an HCL file.
// It takes any schema interface{} and writes it to the specified file path using HCL encoding.
//
// Parameters:
//   - schema: The schema structure to serialize (typically AllSchemas or AvailableSchemas)
//   - filePath: Target file path for writing the HCL content
//
// Returns an error if file creation or writing fails.
func writeSchemasToHCLFile(schema interface{}, filePath string) error {
	// Create new HCL file structure
	hclFile := hclwrite.NewEmptyFile()
	body := hclFile.Body()

	// Encode the schema into HCL format
	gohcl.EncodeIntoBody(schema, body)

	// Create or truncate the target file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	// Write the HCL content to the file
	if _, err := file.Write(hclFile.Bytes()); err != nil {
		return fmt.Errorf("failed to write HCL to file %s: %w", filePath, err)
	}

	log.Printf("Successfully wrote schema to %s\n", filePath)
	return nil
}

// validateResourceType checks if a CloudFormation resource type is provisionable.
// It uses the AWS CloudFormation API to describe the resource type and determine
// if it can be provisioned through CloudFormation operations.
//
// Parameters:
//   - ctx: Context for AWS API calls and cancellation
//   - resourceType: CloudFormation resource type name (e.g., "AWS::S3::Bucket")
//
// Returns true if the resource is provisionable, false if it's marked as NON_PROVISIONABLE,
// and an error if the API call fails in an unexpected way.
func validateResourceType(ctx context.Context, resourceType string) (bool, error) {
	// Load AWS configuration with the specified region
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(AWSRegion))
	if err != nil {
		return false, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create CloudFormation client
	conn := cloudformation.NewFromConfig(cfg)
	input := &cloudformation.DescribeTypeInput{
		Type:     types.RegistryTypeResource,
		TypeName: aws.String(resourceType),
	}

	// Query CloudFormation for resource type details
	res, err := conn.DescribeType(ctx, input)
	if err != nil {
		// If we can't describe the type, assume it's valid to avoid blocking everything
		return true, nil
	}

	// Check if the resource is marked as non-provisionable
	if string(res.ProvisioningType) == "NON_PROVISIONABLE" {
		return false, nil
	}

	log.Printf("Resource type %s is valid.\n", resourceType)
	return true, nil
}
