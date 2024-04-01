// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shared

import (
	"bytes"
	"fmt"
	"strings"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared/codegen"
	"golang.org/x/exp/slices"
)

const (
	DataSourceType = "DataSource"
	ResourceType   = "Resource"
)

// GenerateTemplateData generates the template body from the Resource
// constructed from a CloudFormation type's Schema file.
// This method can be applied to both singular data source and resource types.
func GenerateTemplateData(ui cli.Ui, cfTypeSchemaFile, resType, tfResourceType, packageName string) (*TemplateData, error) {
	resource, err := NewResource(tfResourceType, cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("reading CloudFormation resource schema for %s: %w", tfResourceType, err)
	}

	cfTypeName := *resource.CfResource.TypeName
	org, svc, res, err := naming.ParseCloudFormationTypeName(cfTypeName)

	if err != nil {
		return nil, fmt.Errorf("incorrect format for CloudFormation Resource Provider Schema type name: %s", cfTypeName)
	}

	// e.g. "logGroupResource" or "logGroupDataSource"
	factoryFunctionName := string(bytes.ToLower([]byte(res[:1]))) + res[1:] + resType

	// e.g. "TestAccAWSLogsLogGroup"
	acceptanceTestFunctionPrefix := fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, res)

	sb := strings.Builder{}
	codeEmitter := codegen.Emitter{
		CfResource:   resource.CfResource,
		IsDataSource: resType == DataSourceType,
		Ui:           ui,
		Writer:       &sb,
	}

	// Generate code for the CloudFormation root properties schema.
	attributeNameMap := make(map[string]string) // Terraform attribute name to CloudFormation property name.
	codeFeatures, err := codeEmitter.EmitRootPropertiesSchema(resource.TfType, attributeNameMap)

	if err != nil {
		return nil, fmt.Errorf("emitting schema code: %w", err)
	}

	rootPropertiesSchema := sb.String()
	sb.Reset()

	templateData := &TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		AttributeNameMap:             attributeNameMap,
		CloudFormationTypeName:       cfTypeName,
		FactoryFunctionName:          factoryFunctionName,
		HasRequiredAttribute:         true,
		PackageName:                  packageName,
		RootPropertiesSchema:         rootPropertiesSchema,
		SchemaVersion:                1,
		TerraformTypeName:            resource.TfType,
	}

	if !codeFeatures.HasRequiredRootProperty {
		templateData.HasRequiredAttribute = false
	}
	if codeFeatures.UsesFrameworkTypes {
		templateData.ImportFrameworkTypes = true
	}
	if codeFeatures.UsesFrameworkJSONTypes {
		templateData.ImportFrameworkJSONTypes = true
	}
	if codeFeatures.UsesFrameworkTimeTypes {
		templateData.ImportFrameworkTimeTypes = true
	}
	if codeFeatures.HasValidator {
		templateData.ImportFrameworkValidator = true
	}

	if resType == DataSourceType {
		templateData.SchemaDescription = fmt.Sprintf("Data Source schema for %s", cfTypeName)

		return templateData, nil
	}

	templateData.HasUpdateMethod = true
	templateData.SyntheticIDAttribute = true

	if !codeFeatures.HasUpdatableProperty {
		templateData.HasUpdateMethod = false
	}
	if codeFeatures.UsesRegexpInValidation {
		templateData.ImportRegexp = true
	}
	if codeFeatures.HasIDRootProperty {
		templateData.SyntheticIDAttribute = false
	}

	if description := resource.CfResource.Description; description != nil {
		templateData.SchemaDescription = *description
	}

	for _, path := range resource.CfResource.WriteOnlyProperties {
		templateData.WriteOnlyPropertyPaths = append(templateData.WriteOnlyPropertyPaths, string(path))
	}

	if v, ok := resource.CfResource.Handlers[cfschema.HandlerTypeCreate]; ok {
		templateData.CreateTimeoutInMinutes = v.TimeoutInMinutes
	}
	if v, ok := resource.CfResource.Handlers[cfschema.HandlerTypeUpdate]; ok {
		templateData.UpdateTimeoutInMinutes = v.TimeoutInMinutes
	}
	if v, ok := resource.CfResource.Handlers[cfschema.HandlerTypeDelete]; ok {
		templateData.DeleteTimeoutInMinutes = v.TimeoutInMinutes
	}

	if templateData.SyntheticIDAttribute {
		templateData.FrameworkPlanModifierPackages = []string{"stringplanmodifier"}
	}
	for _, v := range codeFeatures.FrameworkPlanModifierPackages {
		if !slices.Contains(templateData.FrameworkPlanModifierPackages, v) {
			templateData.FrameworkPlanModifierPackages = append(templateData.FrameworkPlanModifierPackages, v)
		}
	}
	for _, v := range codeFeatures.FrameworkValidatorsPackages {
		if !slices.Contains(templateData.FrameworkValidatorsPackages, v) {
			templateData.FrameworkValidatorsPackages = append(templateData.FrameworkValidatorsPackages, v)
		}
	}

	return templateData, nil
}

type TemplateData struct {
	AcceptanceTestFunctionPrefix  string
	AttributeNameMap              map[string]string
	CloudFormationTypeName        string
	CreateTimeoutInMinutes        int
	DeleteTimeoutInMinutes        int
	FactoryFunctionName           string
	FrameworkPlanModifierPackages []string
	FrameworkValidatorsPackages   []string
	HasRequiredAttribute          bool
	HasUpdateMethod               bool
	ImportFrameworkTypes          bool
	ImportFrameworkJSONTypes      bool
	ImportFrameworkTimeTypes      bool
	ImportFrameworkValidator      bool
	ImportRegexp                  bool
	PackageName                   string
	RootPropertiesSchema          string
	SchemaDescription             string
	SchemaVersion                 int64
	SyntheticIDAttribute          bool
	TerraformTypeName             string
	UpdateTimeoutInMinutes        int
	WriteOnlyPropertyPaths        []string
}

type Resource struct {
	CfResource *cfschema.Resource
	TfType     string
}

// NewResource creates a Resource type
// from the corresponding resource's CloudFormation Schema file
func NewResource(resourceType, cfTypeSchemaFile string) (*Resource, error) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("reading CloudFormation Resource Type Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return nil, fmt.Errorf("parsing CloudFormation Resource Type Schema: %w", err)
	}

	if err := resource.Expand(); err != nil {
		return nil, fmt.Errorf("expanding JSON Pointer references: %w", err)
	}

	return &Resource{
		CfResource: resource,
		TfType:     resourceType,
	}, nil
}
