// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shared

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared/codegen"
	"github.com/hashicorp/cli"
	"golang.org/x/exp/slices"
)

const (
	DataSourceType = "DataSource"
	ResourceType   = "Resource"

	DirPerm = 0755
)

type Generator struct {
	UI cli.Ui
}

// ApplyAndWriteTemplate applies the template body to the specified data and writes it to file.
func (g *Generator) ApplyAndWriteTemplate(filename, templateBody string, templateData *TemplateData) error {
	tmpl, err := template.New("function").Parse(templateBody)

	if err != nil {
		return fmt.Errorf("parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		g.Infof("%s", buffer.String())
		return fmt.Errorf("formatting generated source code: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("writing to file (%s): %w", filename, err)
	}

	return nil
}

// GenerateTemplateData generates the template body from the Resource
// constructed from a CloudFormation type's Schema file.
// This method can be applied to both singular data source and resource types.
func (g *Generator) GenerateTemplateData(cfTypeSchemaFile, resType, tfResourceType, packageName string) (*TemplateData, error) {
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
		Ui:           g.UI,
		Writer:       &sb,
	}

	// Generate code for the CloudFormation root properties schema.
	attributeNameMap := make(map[string]string) // Terraform attribute name to CloudFormation property name.
	codeFeatures, err := codeEmitter.EmitRootPropertiesSchema(attributeNameMap)

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
	if codeFeatures.UsesInternalValidate {
		templateData.ImportInternalValidate = true
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

func (g *Generator) Infof(format string, a ...interface{}) {
	g.UI.Info(fmt.Sprintf(format, a...))
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
	ImportFrameworkValidator      bool
	ImportInternalValidate        bool
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
