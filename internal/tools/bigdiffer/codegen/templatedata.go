// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-provider-awscc/internal/identity"
	identitynames "github.com/hashicorp/terraform-provider-awscc/internal/identity/names"
	tfslices "github.com/hashicorp/terraform-provider-awscc/internal/slices"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
)

const (
	DataSourceType = "DataSource"
	ResourceType   = "Resource"
)

// GenerateTemplateData builds the TemplateData for a resource or singular data
// source from its CloudFormation schema file.
func GenerateTemplateData(ui cli.Ui, cfTypeSchemaFile, resType, tfResourceType, packageName, servicesPath string) (*TemplateData, error) {
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
	deduplicate, err := schemaIsDeRecursed(cfTypeSchemaFile)
	if err != nil {
		return nil, err
	}
	codeEmitter := Emitter{
		CfResource:   resource.CfResource,
		IsDataSource: resType == DataSourceType,
		Ui:           ui,
		Writer:       &sb,
		Deduplicate:  deduplicate,
	}

	// Generate code for the CloudFormation root properties schema.
	attributeNameMap := make(map[string]string) // Terraform attribute name to CloudFormation property name.
	codeFeatures, attributeFunctions, err := codeEmitter.EmitRootPropertiesSchema(resource.TfType, attributeNameMap)

	if err != nil {
		return nil, fmt.Errorf("emitting schema code: %w", err)
	}

	rootPropertiesSchema := sb.String()
	sb.Reset()

	templateData := &TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		AttributeFunctions:           attributeFunctions,
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

		// if identifier ls "Provider", rename to "ProviderName" to avoid conflict with provider keyword
		if id == "Provider" {
			id = fmt.Sprintf("%sName", id)
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
		services, err := identitynames.ParseServicesFile(servicesPath)
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
	AttributeFunctions            string
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

	IsGlobal             bool
	HasMutableIdentity   bool
	GenerateListResource bool
}

type Resource struct {
	CfResource *cfschema.Resource
	TfType     string
}

// schemaIsDeRecursed reports whether the CloudFormation schema file was
// depth-limited (de-recursed) by tools/derecurse-schema.py, indicated by the
// presence of the top-level "x-derecursed" key. Only such schemas need the
// deduplicating emitter (#3270); pristine schemas generate inline as before.
func schemaIsDeRecursed(cfTypeSchemaFile string) (bool, error) {
	data, err := os.ReadFile(cfTypeSchemaFile)
	if err != nil {
		return false, fmt.Errorf("reading CloudFormation resource schema %s: %w", cfTypeSchemaFile, err)
	}

	var doc struct {
		XDerecursed json.RawMessage `json:"x-derecursed"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		return false, fmt.Errorf("parsing CloudFormation resource schema %s: %w", cfTypeSchemaFile, err)
	}

	return len(doc.XDerecursed) > 0, nil
}

// NewResource creates a Resource type from the corresponding resource's
// CloudFormation Schema file. The schema is read into memory and parsed via
// NewResourceJsonSchemaDocument rather than NewResourceJsonSchemaPath: the
// path-based loader chdirs into the schema's directory (to resolve relative
// refs) and does not restore the working directory — harmless for the tool (it
// uses absolute paths) but a process-global side effect that corrupts CWD for
// anything relying on it, and a hazard under concurrent generation. Sanitized
// CloudFormation schemas are self-contained (internal "#/…" refs only), so the
// in-memory document is equivalent — proven by full-corpus parity.
func NewResource(resourceType, cfTypeSchemaFile string) (*Resource, error) {
	schemaBytes, err := os.ReadFile(cfTypeSchemaFile)
	if err != nil {
		return nil, fmt.Errorf("reading CloudFormation Resource Type Schema file: %w", err)
	}

	resourceSchema, err := cfschema.NewResourceJsonSchemaDocument(string(schemaBytes))

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
