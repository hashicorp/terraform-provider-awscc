package main

import (
	"context"
	"fmt"
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

func parseSchemaToStruct[T any](filePath string, schema T) (*T, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return &schema, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	parser := hclparse.NewParser()
	fileHCL, diag := parser.ParseHCL(file, filePath)
	if diag.HasErrors() {
		return &schema, fmt.Errorf("failed to parse HCL from %s: %s", filePath, diag.Error())
	}
	if diag := gohcl.DecodeBody(fileHCL.Body, nil, &schema); diag.HasErrors() {
		return &schema, fmt.Errorf("failed to decode HCL from %s: %s", filePath, diag.Error())
	}

	return &schema, nil
}

func diffSchemas(ctx context.Context, newSchemas *allschemas.AvailableSchemas, lastSchemas *allschemas.AvailableSchemas, allSchemasPath string, filePaths *UpdateFilePaths) (*allschemas.AllSchemas, error) {
	// Create a map from lastSchemas for
	// schema name to index
	// use index to get the resource from lastSchemas
	lastSchemasMap := make(map[string]int)
	for i, resource := range lastSchemas.Resources {
		lastSchemasMap[resource.CloudFormationTypeName] = i
	}

	// Array to hold changed and new resources
	var changedOrNewResources []allschemas.ResourceSchema

	// Find changed or new resources
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
		}
	}

	if len(changedOrNewResources) == 0 {
		fmt.Println("No changes or new resources found.")
		return nil, nil
	}

	fmt.Printf("Found %d changed or new resources.\n", len(changedOrNewResources))

	// Read existing allSchemas.hcl
	var err error
	existingAllSchemas, err := parseSchemaToStruct(filePaths.AllSchemasHCL, allschemas.AllSchemas{}) // replace this with file
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
	for i := range existingAllSchemas.Resources {
		err = validateResourceType(ctx, existingAllSchemas.Resources[i].CloudFormationTypeName)
		if err != nil {
			// Remove the invalid resource from the slice
			existingAllSchemas.Resources = append(existingAllSchemas.Resources[:i], existingAllSchemas.Resources[i+1:]...)
			i--
		}
	}
	// Write updated schema back to file
	return existingAllSchemas, writeSchemasToHCLFile(existingAllSchemas, allSchemasPath)
}

func writeSchemasToHCLFile(schema interface{}, filePath string) error {
	hclFile := hclwrite.NewEmptyFile()
	body := hclFile.Body()

	gohcl.EncodeIntoBody(schema, body)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err := file.Write(hclFile.Bytes()); err != nil {
		return fmt.Errorf("failed to write HCL to file %s: %w", filePath, err)
	}

	fmt.Printf("Successfully wrote schema to %s\n", filePath)
	return nil
}

func validateResourceType(ctx context.Context, resourceType string) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}
	conn := cloudformation.NewFromConfig(cfg)

	input := &cloudformation.DescribeTypeInput{
		Type:     types.RegistryTypeResource,
		TypeName: aws.String(resourceType),
	}

	res, err := conn.DescribeType(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to describe resource type %s: %w", resourceType, err)
	}
	if string(res.ProvisioningType) == "NON_PROVISIONABLE" {
		return fmt.Errorf("resource type %s is not provisionable", resourceType)
	}
	fmt.Printf("Resource type %s is valid.\n", resourceType)
	return nil
}
