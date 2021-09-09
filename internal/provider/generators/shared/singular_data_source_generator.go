package shared

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/cli"
)

type SingularDataSourceGenerator struct {
	cfTypeSchemaFile string
	tfDataSourceType string
	Generator
}

func NewSingularDataSourceGenerator(ui cli.Ui, acceptanceTestsTemplateBody, schemaTemplateBody, cfTypeSchemaFile, tfDataSourceType string) *SingularDataSourceGenerator {
	return &SingularDataSourceGenerator{
		cfTypeSchemaFile: cfTypeSchemaFile,
		tfDataSourceType: tfDataSourceType,
		Generator: Generator{
			acceptanceTestsTemplateBody: acceptanceTestsTemplateBody,
			schemaTemplateBody:          schemaTemplateBody,
			ui:                          ui,
		},
	}
}

// Generate generates the singular data source's type factory into the specified file.
func (s *SingularDataSourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	s.infof("generating Terraform data source code for %[1]q from %[2]q into %[3]q and %[4]q", s.tfDataSourceType, s.cfTypeSchemaFile, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	templateData, err := s.generateTemplateData(s.cfTypeSchemaFile, DataSourceType, s.tfDataSourceType, packageName)

	if err != nil {
		return err
	}

	err = s.applyAndWriteTemplate(schemaFilename, s.schemaTemplateBody, templateData)

	if err != nil {
		return err
	}

	err = s.applyAndWriteTemplate(acctestsFilename, s.acceptanceTestsTemplateBody, templateData)

	if err != nil {
		return err
	}

	return nil
}
