// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	_ "embed"
	"fmt"
	"go/format"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-provider-awscc/internal/identity"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
)

//go:embed resource_schema.tmpl
var resourceSchemaTemplate string

//go:embed resource_tests.tmpl
var resourceTestsTemplate string

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
