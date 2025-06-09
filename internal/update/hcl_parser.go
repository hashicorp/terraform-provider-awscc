package update

import (
	"fmt"
	"os"
	"sort"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	allschemas "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/allschemas"
)

func parseSchemaToStruct(filePath string) (*allschemas.AllSchemas, error) {
	allSchemas := &allschemas.AllSchemas{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	parser := hclparse.NewParser()
	fileHCL, diag := parser.ParseHCL(file, filePath)
	if diag.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL from %s: %s", filePath, diag.Error())
	}
	if diag := gohcl.DecodeBody(fileHCL.Body, nil, allSchemas); diag.HasErrors() {
		return nil, fmt.Errorf("failed to decode HCL from %s: %s", filePath, diag.Error())
	}

	return allSchemas, nil
} // good

func mapSchemasFiles(filePath1 string, filePath2 string) (map[string]allschemas.ResourceSchema, map[string]allschemas.ResourceSchema, error) {
	allSchemas1, err := parseSchemaToStruct(filePath1)
	if err != nil {
		return nil, nil, err
	}

	allSchemas2, err := parseSchemaToStruct(filePath2)
	if err != nil {
		return nil, nil, err
	}

	allSchemas1map := make(map[string]allschemas.ResourceSchema)
	for _, resource := range allSchemas1.Resources {
		allSchemas1map[resource.CloudFormationTypeName] = resource
	}

	allSchemas2map := make(map[string]allschemas.ResourceSchema)
	for _, resource := range allSchemas2.Resources {
		allSchemas2map[resource.CloudFormationTypeName] = resource
	}

	return allSchemas1map, allSchemas2map, nil
}

func diffAllSchemas(newResources []allschemas.ResourceSchema, allSchemasPath string) error {
	allSchemasStruct := &allschemas.AllSchemas{}
	if len(newResources) == 0 {
		fmt.Println("No new resources to add.")
		return nil
	}
	fmt.Printf("Found %d new resources to add.\n", len(newResources))
	allSchemasFile, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", allSchemasPath, err)
	}
	parser := hclparse.NewParser()
	allSchemaHCL, diag := parser.ParseHCL(allSchemasFile, allSchemasPath)
	if diag.HasErrors() {
		return fmt.Errorf("failed to parse HCL from %s: %s", allSchemasPath, diag.Error())
	}
	if diag := gohcl.DecodeBody(allSchemaHCL.Body, nil, allSchemasStruct); diag.HasErrors() {
		return fmt.Errorf("failed to decode HCL from %s: %s", allSchemasPath, diag.Error())
	}

	allSchemasMap := make(map[string]allschemas.ResourceSchema)
	for _, resource := range allSchemasStruct.Resources {
		allSchemasMap[resource.CloudFormationTypeName] = resource
	}

	for _, newResource := range newResources {
		if _, exists := allSchemasMap[newResource.CloudFormationTypeName]; exists {
			fmt.Printf("Resource %s already exists in all_schemas.hcl, skipping.\n", newResource.CloudFormationTypeName)
			continue
		}
		allSchemasStruct.Resources = append(allSchemasStruct.Resources, newResource)
	}
	outputFile, err := os.Create(allSchemasPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s for writing: %w", allSchemasPath, err)
	}
	defer outputFile.Close()

	hclFile := hclwrite.NewEmptyFile()
	body := hclFile.Body()
	gohcl.EncodeIntoBody(allSchemasStruct, body)

	if _, err := outputFile.Write(hclFile.Bytes()); err != nil {
		return fmt.Errorf("failed to write HCL to file %s: %w", allSchemasPath, err)
	}

	fmt.Printf("Updated %s with new resources.\n", allSchemasPath)
	return nil
}

/*
go through the last schema and put it all in a map
go through current schema and for each resource
check if it exists in the last schema map
if it does, compare the two resources and if there is a change add to array a
if it doesn't, it's a new resource and add to array a
go through allSchemas.hcl and make a set to check if the resource exists in allSchemas
if it doesn't exist add to array of original allSchemas array
make a set s of all matching resources between array a and allSchemas set
go through allSchemas array and for each resource check if it exists in set s and then update the resource in a non destructive way (dont remove any fields)
write allSchemas array to allSchemas.hcl

make sure to check for errors and
*/
func diffSchemas(newSchemas, lastSchemas *allschemas.AllSchemas, allSchemasPath string) error {
	// Create a map from lastSchemas for lookup
	lastSchemasMap := make(map[string]allschemas.ResourceSchema)
	for _, resource := range lastSchemas.Resources {
		lastSchemasMap[resource.ResourceTypeName] = resource
	}

	// Array to hold changed and new resources
	var changedOrNewResources []allschemas.ResourceSchema

	// Find changed or new resources
	for _, newResource := range newSchemas.Resources {
		/*
					var currResourceSchema = ResourceSchema{
				ResourceTypeName:       tfTypeName,
				CloudFormationTypeName: cfTypeName,
			}

			if suppressPluralDataSource {
				currResourceSchema.SuppressPluralDataSourceGeneration = suppressPluralDataSource
			}


		*/
		if lastResource, exists := lastSchemasMap[newResource.CloudFormationTypeName]; exists {
			// Check if resource changed
			if newResource.ResourceTypeName != lastResource.ResourceTypeName ||
				newResource.CloudFormationTypeName != lastResource.CloudFormationTypeName ||
				newResource.SuppressPluralDataSourceGeneration != lastResource.SuppressPluralDataSourceGeneration {
				// newResource.SuppressSingularDataSourceGeneration != lastResource.SuppressSingularDataSourceGeneration  does this come iwth it
				changedOrNewResources = append(changedOrNewResources, newResource)
			}
		} else {
			// New resource
			changedOrNewResources = append(changedOrNewResources, newResource)
		}
	}

	if len(changedOrNewResources) == 0 {
		fmt.Println("No changes or new resources found.")
		return nil
	}

	fmt.Printf("Found %d changed or new resources.\n", len(changedOrNewResources))

	// Read existing allSchemas.hcl
	existingAllSchemas := &allschemas.AllSchemas{}
	file, err := os.ReadFile(allSchemasPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", allSchemasPath, err)
	}

	parser := hclparse.NewParser()
	fileHCL, diag := parser.ParseHCL(file, allSchemasPath)
	if diag.HasErrors() {
		return fmt.Errorf("failed to parse HCL from %s: %s", allSchemasPath, diag.Error())
	}
	if diag := gohcl.DecodeBody(fileHCL.Body, nil, existingAllSchemas); diag.HasErrors() {
		return fmt.Errorf("failed to decode HCL from %s: %s", allSchemasPath, diag.Error())
	}

	// Create a map of existing resources
	existingResourcesMap := make(map[string]bool)
	for _, resource := range existingAllSchemas.Resources {
		existingResourcesMap[resource.CloudFormationTypeName] = true
	}

	// Add new resources if they don't exist in allSchemas
	for _, resource := range changedOrNewResources {
		if !existingResourcesMap[resource.CloudFormationTypeName] {
			existingAllSchemas.Resources = append(existingAllSchemas.Resources, resource)
			fmt.Printf("Adding new resource: %s\n", resource.CloudFormationTypeName)
		}
	}

	// Create a set of resources that need to be updated
	updateMap := make(map[string]allschemas.ResourceSchema)
	for _, resource := range changedOrNewResources {
		if existingResourcesMap[resource.CloudFormationTypeName] {
			updateMap[resource.CloudFormationTypeName] = resource
		}
	}

	// Update existing resources
	for i, resource := range existingAllSchemas.Resources {
		if updatedResource, ok := updateMap[resource.CloudFormationTypeName]; ok {
			existingAllSchemas.Resources[i] = updatedResource
			fmt.Printf("Updating existing resource: %s\n", resource.CloudFormationTypeName)
		}
	}
	sort.Slice(existingAllSchemas.Resources, func(i, j int) bool {
		return existingAllSchemas.Resources[i].ResourceTypeName < existingAllSchemas.Resources[j].ResourceTypeName
	})
	// Write updated schema back to file
	return writeSchemasToHCLFile(existingAllSchemas, allSchemasPath)
}

func writeSchemasToHCLFile(schema *allschemas.AllSchemas, filePath string) error {
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
} // good
