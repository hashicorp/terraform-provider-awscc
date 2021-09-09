package shared

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/cli"
)

type ResourceGenerator struct {
	cfTypeSchemaFile string
	tfResourceType   string
	Generator
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
