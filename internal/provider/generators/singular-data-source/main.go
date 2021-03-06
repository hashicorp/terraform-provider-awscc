// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
	"github.com/mitchellh/cli"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfDataSourceType = flag.String("data-source", "", "Terraform data source type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -data-source <TF-data-source-type> -cfschema <CF-type-schema-file> <generated-schema-file> <generated-acctests-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 || *tfDataSourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	schemaFilename := args[0]
	acctestsFilename := args[1]

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	generator := &SingularDataSourceGenerator{
		cfTypeSchemaFile: *cfTypeSchemaFile,
		tfDataSourceType: *tfDataSourceType,
		Generator: shared.Generator{
			UI: ui,
		},
	}

	if err := generator.Generate(destinationPackage, schemaFilename, acctestsFilename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform %s data source: %s", *tfDataSourceType, err))
		os.Exit(1)
	}
}

type SingularDataSourceGenerator struct {
	cfTypeSchemaFile string
	tfDataSourceType string
	shared.Generator
}

// Generate generates the singular data source's type factory into the specified file.
func (s *SingularDataSourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	s.Infof("generating Terraform data source code for %[1]q from %[2]q into %[3]q and %[4]q", s.tfDataSourceType, s.cfTypeSchemaFile, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, shared.DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	templateData, err := s.GenerateTemplateData(s.cfTypeSchemaFile, shared.DataSourceType, s.tfDataSourceType, packageName)

	if err != nil {
		return err
	}

	err = s.ApplyAndWriteTemplate(schemaFilename, dataSourceSchemaTemplateBody, templateData)

	if err != nil {
		return err
	}

	err = s.ApplyAndWriteTemplate(acctestsFilename, acceptanceTestsTemplateBody, templateData)

	if err != nil {
		return err
	}

	return nil
}

// Terraform data source schema definition.
var dataSourceSchemaTemplateBody = `
// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package {{ .PackageName }}

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("{{ .TerraformTypeName }}", {{ .FactoryFunctionName }})
}

// {{ .FactoryFunctionName }} returns the Terraform {{ .TerraformTypeName }} data source type.
// This Terraform data source type corresponds to the CloudFormation {{ .CloudFormationTypeName }} resource type.
func {{ .FactoryFunctionName }}(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := {{ .RootPropertiesSchema }}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "{{ .SchemaDescription }}",
		Version:     {{ .SchemaVersion }},
		Attributes:  attributes,
	}

    var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("{{ .CloudFormationTypeName }}").WithTerraformTypeName("{{ .TerraformTypeName }}")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
{{- range $key, $value := .AttributeNameMap }}
		"{{ $key }}": "{{ $value }}",
{{- end }}
	})

    singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
`

// Terraform acceptance tests.
var acceptanceTestsTemplateBody = `
// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package {{ .PackageName }}_test

import (
    {{ if not .HasRequiredAttribute }}"fmt"{{- end }}
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

{{ if .HasRequiredAttribute }}
func {{ .AcceptanceTestFunctionPrefix }}DataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
{{- else }}
func {{ .AcceptanceTestFunctionPrefix }}DataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.DataSourceWithEmptyResourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "id", td.ResourceName, "id"), 
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "arn", td.ResourceName, "arn"),
			),
		},
	})
}
{{- end }}

func {{ .AcceptanceTestFunctionPrefix }}DataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
`
