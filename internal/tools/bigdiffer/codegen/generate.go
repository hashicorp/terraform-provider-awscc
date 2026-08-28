// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	_ "embed"
	"fmt"
	"go/format"
	"strings"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-provider-awscc/internal/identity"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
)

//go:embed resource_schema.tmpl
var resourceSchemaTemplate string

//go:embed resource_tests.tmpl
var resourceTestsTemplate string

//go:embed singular_data_source_schema.tmpl
var singularDataSourceSchemaTemplate string

//go:embed singular_data_source_tests.tmpl
var singularDataSourceTestsTemplate string

//go:embed plural_data_source_schema.tmpl
var pluralDataSourceSchemaTemplate string

//go:embed plural_data_source_tests.tmpl
var pluralDataSourceTestsTemplate string

//go:embed imports.tmpl
var importExamplesTemplate string

// ImportExample is one resource's entry in import_examples_gen.json.
type ImportExample struct {
	ResourceName         string
	GenerateListResource bool
	Identifier           []string
}

// GenerateImportExamples renders the import_examples_gen.json aggregate from the
// given resource entries. The output is JSON (not Go), so it is not
// gofmt-formatted — the template's exact layout is authoritative.
func GenerateImportExamples(examples []ImportExample) ([]byte, error) {
	data := struct{ Resources []ImportExample }{Resources: examples}
	return parseTemplate("imports", importExamplesTemplate, data)
}

// GenerateResource renders a resource's generated code and acceptance-test source
// in-process, returning the gofmt-formatted bytes. It is the owned, in-process
// equivalent of generators/resource/main.go: build the template data, snake-case
// the primary-identifier names, honor the list-resource flag, render, and format.
// servicesPath supplies the global/identity flags (§ services.hcl).
func GenerateResource(ui cli.Ui, schemaFile, tfType, packageName, servicesPath string, listResource bool) (code, test []byte, err error) {
	td, err := GenerateTemplateData(ui, schemaFile, ResourceType, tfType, packageName, servicesPath)
	if err != nil {
		return nil, nil, err
	}

	// Mirror resource/main.go: snake_case the primary-identifier names.
	primaryIdentifier := make([]identity.Identifier, len(td.PrimaryIdentifier))
	for i, v := range td.PrimaryIdentifier {
		primaryIdentifier[i] = identity.Identifier{
			Name:        naming.SnakeCase(v.Name),
			Description: v.Description,
		}
	}
	td.PrimaryIdentifier = primaryIdentifier
	td.GenerateListResource = listResource

	if code, err = renderGo("resource", resourceSchemaTemplate, td); err != nil {
		return nil, nil, err
	}
	if test, err = renderGo("acctest", resourceTestsTemplate, td); err != nil {
		return nil, nil, err
	}
	return code, test, nil
}

// renderGo renders a template and gofmt-formats the result, so a template that
// emits invalid Go surfaces as an error here rather than at compile time.
func renderGo(name, body string, data any) ([]byte, error) {
	b, err := parseTemplate(name, body, data)
	if err != nil {
		return nil, err
	}
	formatted, err := format.Source(b)
	if err != nil {
		return nil, fmt.Errorf("formatting generated %s:\n%s\n%w", name, b, err)
	}
	return formatted, nil
}

// GenerateSingularDataSource renders a singular data source's code and test
// in-process. servicesPath is unused for data sources (GenerateTemplateData reads
// services.hcl only for resources) but kept for signature symmetry.
func GenerateSingularDataSource(ui cli.Ui, schemaFile, tfType, packageName, servicesPath string) (code, test []byte, err error) {
	td, err := GenerateTemplateData(ui, schemaFile, DataSourceType, tfType, packageName, servicesPath)
	if err != nil {
		return nil, nil, err
	}
	if code, err = renderGo("data-source", singularDataSourceSchemaTemplate, td); err != nil {
		return nil, nil, err
	}
	if test, err = renderGo("acctest", singularDataSourceTestsTemplate, td); err != nil {
		return nil, nil, err
	}
	return code, test, nil
}

// GeneratePluralDataSource renders a plural data source's code and test
// in-process. Unlike the others it does no schema emission — the template data is
// built from the CloudFormation type name alone (mirroring
// generators/plural-data-source/main.go). pluralTFType is the pluralized
// Terraform type name (computed safely upstream, e.g. by the plan).
func GeneratePluralDataSource(cfType, pluralTFType, packageName string) (code, test []byte, err error) {
	org, svc, res, err := naming.ParseCloudFormationTypeName(cfType)
	if err != nil {
		return nil, nil, fmt.Errorf("incorrect format for CloudFormation type name %q: %w", cfType, err)
	}

	ds := naming.PluralizeWithCustomNameSuffix(res, "Plural")
	factoryFunctionName := strings.ToLower(ds[:1]) + ds[1:] + DataSourceType

	td := &TemplateData{
		AcceptanceTestFunctionPrefix: fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, ds),
		CloudFormationTypeName:       cfType,
		FactoryFunctionName:          factoryFunctionName,
		PackageName:                  packageName,
		SchemaDescription:            fmt.Sprintf("Plural Data Source schema for %s", cfType),
		SchemaVersion:                1,
		TerraformTypeName:            pluralTFType,
	}

	if code, err = renderGo("data-source", pluralDataSourceSchemaTemplate, td); err != nil {
		return nil, nil, err
	}
	if test, err = renderGo("acctest", pluralDataSourceTestsTemplate, td); err != nil {
		return nil, nil, err
	}
	return code, test, nil
}
