package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	provider := provider.New()

	resources, _ := provider.GetResources(context.Background())

	for resource := range resources {
		GenerateExample(resource, ui)
	}
}

func GenerateExample(resourceName string, ui *cli.BasicUi) {

	templateData := TemplateData{
		ResourceType: resourceName,
	}

	tmpl, err := template.New("function").Parse(importExampleTemplateBody)

	if err != nil {
		ui.Error(fmt.Sprintf("error parsing function template: %s", err))
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		ui.Error(fmt.Sprintf("error executing template: %s", err))
	}

	examplesPath := "./examples/resources"

	dirname := fmt.Sprintf("%s/%s", examplesPath, resourceName)
	err = os.MkdirAll(dirname, shared.DirPerm)

	if err != nil {
		ui.Error(fmt.Sprintf("creating target directory %s: %s", dirname, err))
	}

	filename := fmt.Sprintf("%s/import.sh", dirname)

	f, err := os.Create(filename)

	if err != nil {
		ui.Error(fmt.Sprintf("error creating file (%s): %s", filename, err))
	}

	defer f.Close()

	_, err = f.Write(buffer.Bytes())

	if err != nil {
		ui.Error(fmt.Sprintf("error writing to file (%s): %s", filename, err))
	}

}

type TemplateData struct {
	ResourceType string
}

var importExampleTemplateBody = `$ terraform import {{ .ResourceType }}.example <resource ID>`
