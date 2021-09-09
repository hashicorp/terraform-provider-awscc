package shared

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/mitchellh/cli"
)

type PluralDataSourceGenerator struct {
	cfType           string
	tfDataSourceType string
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
