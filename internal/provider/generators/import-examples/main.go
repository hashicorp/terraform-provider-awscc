package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider"
)

// TODO: Split up function
// Code drawn from ../schema/main.go func GenerateDataSources
func main() {
	provider := provider.New()

	resources, _ := provider.GetResources(context.Background())

	for resource := range resources {

		templateData := TemplateData{
			ResourceType: resource,
		}

		tmpl, err := template.New("function").Parse(importExampleTemplateBody)

		// TODO: Actual error handling as opposed to just printing out
		if err != nil {
			fmt.Println(err)
			fmt.Errorf("error parsing function template: %w", err)
		}

		var buffer bytes.Buffer
		err = tmpl.Execute(&buffer, templateData)

		if err != nil {
			fmt.Println(err)
			fmt.Errorf("error executing template: %w", err)
		}

		filename := fmt.Sprintf("%s.sh", resource)

		// TODO: Write to /examples/resources/resourceType/import.sh instead of same directory
		f, err := os.Create(filename)

		if err != nil {
			fmt.Println(err)
			fmt.Errorf("error creating file (%s): %w", filename, err)
		}

		defer f.Close()

		_, err = f.Write(buffer.Bytes())

		if err != nil {
			fmt.Println(err)
			fmt.Errorf("error writing to file (%s): %w", filename, err)
		}
	}
}

type TemplateData struct {
	ResourceType string
}

// TODO: Template currently renders with an unwanted empty line at top
var importExampleTemplateBody = `
terraform import {{ .ResourceType }}.example <resource ID>
`
