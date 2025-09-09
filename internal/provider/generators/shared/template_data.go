// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shared

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-provider-awscc/internal/identity"
	identitynames "github.com/hashicorp/terraform-provider-awscc/internal/identity/names"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared/codegen"
	tfslices "github.com/hashicorp/terraform-provider-awscc/internal/slices"
)

const (
	DataSourceType = "DataSource"
	ResourceType   = "Resource"
	FilePermission = 0600
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

	resourceName := fmt.Sprintf("%s_%s_%s\n", org, svc, res)
	_, err = os.Getwd()
	if err != nil {
		ui.Warn(fmt.Sprintf("Failed to get current working directory: %s", err))
	}
	if writeErr := os.WriteFile("last_resource.txt", []byte(resourceName), FilePermission); writeErr != nil {
		// Log but don't fail if writing debug file fails
		ui.Warn(fmt.Sprintf("Failed to write to last_resource.txt: %s", writeErr))
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
	if codeFeatures.UsesInternalDefaultsPackage {
		templateData.ImportInternalDefaults = true
	}
	if codeFeatures.HasValidator {
		templateData.ImportFrameworkValidator = true
	}
	if codeFeatures.UsesInternalValidatorsPackage {
		templateData.ImportInternalValidators = true
	}

	if resType == DataSourceType {
		templateData.SchemaDescription = fmt.Sprintf("Data Source schema for %s", cfTypeName)

		return templateData, nil
	}

	templateData.HasUpdateMethod = true

	if !codeFeatures.HasUpdatableProperty {
		templateData.HasUpdateMethod = false
	}
	if codeFeatures.UsesRegexpInValidation {
		templateData.ImportRegexp = true
	}

	if description := resource.CfResource.Description; description != nil {
		templateData.SchemaDescription = *description
	}

	for _, path := range resource.CfResource.WriteOnlyProperties {
		templateData.WriteOnlyPropertyPaths = append(templateData.WriteOnlyPropertyPaths, string(path))
	}

	var identifiers []identity.Identifier
	for _, path := range resource.CfResource.PrimaryIdentifier {
		id := strings.TrimPrefix(string(path), "/properties/")

		pID := strings.Split(id, "/")
		identifier := identity.Identifier{}
		if len(pID) != 1 {
			id = strings.Join(pID, "")
		}

		if id == "Provider" {
			id = fmt.Sprintf("%sId", id)
		}
		identifier.Name = id

		if v, ok := resource.CfResource.Properties[id]; ok {
			if v.Description != nil {
				identifier.Description = strings.Split(*v.Description, ".")[0]
			}
		}
		identifiers = append(identifiers, identifier)
	}
	templateData.PrimaryIdentifier = identifiers

	if v, ok := resource.CfResource.Handlers[cfschema.HandlerTypeCreate]; ok {
		templateData.CreateTimeoutInMinutes = v.TimeoutInMinutes
	}
	if v, ok := resource.CfResource.Handlers[cfschema.HandlerTypeUpdate]; ok {
		templateData.UpdateTimeoutInMinutes = v.TimeoutInMinutes
	}
	if v, ok := resource.CfResource.Handlers[cfschema.HandlerTypeDelete]; ok {
		templateData.DeleteTimeoutInMinutes = v.TimeoutInMinutes
	}

	templateData.FrameworkDefaultsPackages = tfslices.AppendUnique(templateData.FrameworkDefaultsPackages, codeFeatures.FrameworkDefaultsPackages...)
	templateData.FrameworkPlanModifierPackages = []string{"stringplanmodifier"} // For the 'id' attribute.
	templateData.FrameworkPlanModifierPackages = tfslices.AppendUnique(templateData.FrameworkPlanModifierPackages, codeFeatures.FrameworkPlanModifierPackages...)
	templateData.FrameworkValidatorsPackages = tfslices.AppendUnique(templateData.FrameworkValidatorsPackages, codeFeatures.FrameworkValidatorsPackages...)

	// add global flag for resources only
	if resType == ResourceType {
		services, err := identitynames.ParseServicesFile("../identity/names/services.hcl")
		if err != nil {
			return nil, err
		}

		if services != nil {
			serviceName := identitynames.GetServiceName(templateData.CloudFormationTypeName)
			if serviceName != "" {
				t := slices.IndexFunc(services.Services, func(s identitynames.Service) bool {
					return s.ServiceName == serviceName
				})

				if t != -1 {
					templateData.IsGlobal = services.Services[t].IsGlobal

					if s := services.Services[t].Resources; s != nil {
						t := slices.IndexFunc(s, func(r identitynames.Resource) bool {
							return r.TFResourceName == templateData.TerraformTypeName
						})

						if t != -1 {
							templateData.HasMutableIdentity = s[t].HasMutableIdentity
						}
					}
				}
			}
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
	FrameworkDefaultsPackages     []string
	FrameworkPlanModifierPackages []string
	FrameworkValidatorsPackages   []string
	HasRequiredAttribute          bool
	HasUpdateMethod               bool
	ImportFrameworkTypes          bool
	ImportFrameworkJSONTypes      bool
	ImportFrameworkTimeTypes      bool
	ImportFrameworkValidator      bool
	ImportInternalDefaults        bool
	ImportInternalValidators      bool
	ImportRegexp                  bool
	PackageName                   string
	PrimaryIdentifier             []identity.Identifier
	RootPropertiesSchema          string
	SchemaDescription             string
	SchemaVersion                 int64
	TerraformTypeName             string
	UpdateTimeoutInMinutes        int
	WriteOnlyPropertyPaths        []string

	IsGlobal           bool
	HasMutableIdentity bool
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
