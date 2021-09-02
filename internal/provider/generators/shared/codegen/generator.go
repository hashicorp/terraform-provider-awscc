package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/mitchellh/cli"
)

const (
	DataSourceType = "DataSourceType"
	ResourceType   = "ResourceType"

	DirPerm = 0755
)

type Generator struct {
	acceptanceTestsTemplateBody string
	schemaTemplateBody          string
	ui                          cli.Ui
}

type PluralDataSourceGenerator struct {
	cfType           string
	tfDataSourceType string
	Generator
}

type ResourceGenerator struct {
	cfTypeSchemaFile string
	tfResourceType   string
	Generator
}

func NewPluralDataSourceGenerator(ui cli.Ui, acceptanceTestsTemplateBody, schemaTemplateBody, cfType, tfDataSourceType string) *PluralDataSourceGenerator {
	return &PluralDataSourceGenerator{
		cfType:           cfType,
		tfDataSourceType: tfDataSourceType,
		Generator: Generator{
			acceptanceTestsTemplateBody: acceptanceTestsTemplateBody,
			schemaTemplateBody:          schemaTemplateBody,
			ui:                          ui,
		},
	}
}

func NewResourceGenerator(ui cli.Ui, acceptanceTestsTemplateBody, schemaTemplateBody, cfTypeSchemaFile, tfResourceType string) *ResourceGenerator {
	return &ResourceGenerator{
		cfTypeSchemaFile: cfTypeSchemaFile,
		tfResourceType:   tfResourceType,
		Generator: Generator{
			acceptanceTestsTemplateBody: acceptanceTestsTemplateBody,
			schemaTemplateBody:          schemaTemplateBody,
			ui:                          ui,
		},
	}
}

// Generate generates the plural data source type's factory into the specified file.
func (p *PluralDataSourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	p.infof("generating Terraform data source code for %[1]q into %[2]q and %[3]q", p.tfDataSourceType, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	org, svc, res, err := naming.ParseCloudFormationTypeName(p.cfType)

	if err != nil {
		return fmt.Errorf("incorrect format for CloudFormation Resource Provider Schema type name: %s", p.cfType)
	}

	ds := naming.PluralizeWithCustomNameSuffix(res, "Plural")

	factoryFunctionName := string(bytes.ToLower([]byte(ds[:1]))) + ds[1:] + DataSourceType

	acceptanceTestFunctionPrefix := fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, ds)

	schemaDescription := fmt.Sprintf("Plural Data Source schema for %s", p.cfType)

	templateData := TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		CloudFormationTypeName:       p.cfType,
		FactoryFunctionName:          factoryFunctionName,
		PackageName:                  packageName,
		SchemaDescription:            schemaDescription,
		SchemaVersion:                1,
		TerraformTypeName:            p.tfDataSourceType,
	}

	err = p.applyAndWriteTemplate(schemaFilename, p.schemaTemplateBody, &templateData)

	if err != nil {
		return err
	}

	err = p.applyAndWriteTemplate(acctestsFilename, p.acceptanceTestsTemplateBody, &templateData)

	if err != nil {
		return err
	}

	return nil
}

// Generate generates the resource's type factory into the specified file.
func (r *ResourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	r.infof("generating Terraform resource code for %[1]q from %[2]q into %[3]q and %[4]q", r.tfResourceType, r.cfTypeSchemaFile, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	templateData, err := r.generateTemplateData(r.cfTypeSchemaFile, ResourceType, r.tfResourceType, packageName)

	if err != nil {
		return err
	}

	err = r.applyAndWriteTemplate(schemaFilename, r.schemaTemplateBody, templateData)

	if err != nil {
		return err
	}

	err = r.applyAndWriteTemplate(acctestsFilename, r.acceptanceTestsTemplateBody, templateData)

	if err != nil {
		return err
	}

	return nil
}

// applyAndWriteTemplate applies the template body to the specified data and writes it to file.
func (g *Generator) applyAndWriteTemplate(filename, templateBody string, templateData *TemplateData) error {
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
		g.infof("%s", buffer.String())
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

// generateTemplateData generates the template body from the Resource
// constructed from a CloudFormation type's Schema file.
// This method can be applied to both singular data source and resource types.
func (g *Generator) generateTemplateData(cfTypeSchemaFile, resType, tfResourceType, packageName string) (*TemplateData, error) {
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
	codeEmitter := Emitter{
		CfResource: resource.CfResource,
		Ui:         g.ui,
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

	codeEmitter.EmitResourceSchemaRequiredAttributesValidator()
	requiredAttributesValidator := sb.String()
	sb.Reset()

	templateData := &TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		AttributeNameMap:             attributeNameMap,
		CloudFormationTypeName:       cfTypeName,
		FactoryFunctionName:          factoryFunctionName,
		HasRequiredAttribute:         true,
		HasUpdateMethod:              true,
		PackageName:                  packageName,
		RequiredAttributesValidator:  requiredAttributesValidator,
		RootPropertiesSchema:         rootPropertiesSchema,
		SchemaVersion:                1,
		SyntheticIDAttribute:         true,
		TerraformTypeName:            resource.TfType,
	}

	if codeFeatures&HasRequiredRootProperty == 0 {
		templateData.HasRequiredAttribute = false
	}
	if codeFeatures&HasUpdatableProperty == 0 {
		templateData.HasUpdateMethod = false
	}
	if codeFeatures&UsesInternalTypes > 0 {
		templateData.ImportInternalTypes = true
	}
	if codeFeatures&UsesValidation > 0 || requiredAttributesValidator != "" {
		templateData.ImportValidate = true
	}
	if codeFeatures&HasIDRootProperty > 0 {
		templateData.SyntheticIDAttribute = false
	}

	if resType == DataSourceType {
		templateData.SchemaDescription = fmt.Sprintf("Data Source schema for %s", cfTypeName)
	} else if description := resource.CfResource.Description; description != nil {
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

func (g *Generator) infof(format string, a ...interface{}) {
	g.ui.Info(fmt.Sprintf(format, a...))
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
	ImportInternalTypes          bool
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
