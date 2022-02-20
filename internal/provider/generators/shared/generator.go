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
	"github.com/mitchellh/cli"
)

const (
	DataSourceType = "DataSourceType"
	ResourceType   = "ResourceType"

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

	// e.g. "logGroupResourceType" or "logGroupDataSourceType"
	factoryFunctionName := string(bytes.ToLower([]byte(res[:1]))) + res[1:] + resType

	// e.g. "TestAccAWSLogsLogGroup"
	acceptanceTestFunctionPrefix := fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, res)

	sb := strings.Builder{}
	codeEmitter := codegen.Emitter{
		CfResource: resource.CfResource,
		Ui:         g.UI,
		Writer:     &sb,
	}

	// Generate code for the CloudFormation root properties schema.
	attributeNameMap := make(map[string]string) // Terraform attribute name to CloudFormation property name.
	computedOnly := resType == DataSourceType
	codeFeatures, err := codeEmitter.EmitRootPropertiesSchema(attributeNameMap, computedOnly)

	if err != nil {
		return nil, fmt.Errorf("emitting schema code: %w", err)
	}

	rootPropertiesSchema := sb.String()
	sb.Reset()

	err = codeEmitter.EmitResourceSchemaRequiredAttributesValidator()

	if err != nil {
		return nil, fmt.Errorf("emitting schema required attributes validator: %w", err)
	}

	requiredAttributesValidator := sb.String()
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

	if codeFeatures&codegen.HasRequiredRootProperty == 0 {
		templateData.HasRequiredAttribute = false
	}
	if codeFeatures&codegen.UsesFrameworkAttr > 0 {
		templateData.ImportFrameworkAttr = true
	}

	if resType == DataSourceType {
		templateData.SchemaDescription = fmt.Sprintf("Data Source schema for %s", cfTypeName)

		return templateData, nil
	}

	templateData.HasUpdateMethod = true
	templateData.RequiredAttributesValidator = requiredAttributesValidator
	templateData.SyntheticIDAttribute = true

	if codeFeatures&codegen.HasUpdatableProperty == 0 {
		templateData.HasUpdateMethod = false
	}
	if codeFeatures&codegen.UsesValidation > 0 && codeFeatures&codegen.UsesRegexpInValidation > 0 {
		templateData.ImportRegexp = true
	}
	if codeFeatures&codegen.UsesValidation > 0 || requiredAttributesValidator != "" {
		templateData.ImportValidate = true
	}
	if codeFeatures&codegen.HasIDRootProperty > 0 {
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

	return templateData, nil
}

func (g *Generator) Infof(format string, a ...interface{}) {
	g.UI.Info(fmt.Sprintf(format, a...))
}

type TemplateData struct {
	AcceptanceTestFunctionPrefix string
	AttributeNameMap             map[string]string
	CloudFormationTypeName       string
	CreateTimeoutInMinutes       int
	DeleteTimeoutInMinutes       int
	FactoryFunctionName          string
	HasRequiredAttribute         bool
	HasUpdateMethod              bool
	ImportFrameworkAttr          bool
	ImportRegexp                 bool
	ImportValidate               bool
	PackageName                  string
	RequiredAttributesValidator  string
	RootPropertiesSchema         string
	SchemaDescription            string
	SchemaVersion                int64
	SyntheticIDAttribute         bool
	TerraformTypeName            string
	UpdateTimeoutInMinutes       int
	WriteOnlyPropertyPaths       []string
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
