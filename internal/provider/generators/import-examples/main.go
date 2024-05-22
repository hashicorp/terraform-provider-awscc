// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
)

func main() {
	g := NewGenerator()
	resources := provider.New().Resources(context.Background())

	for _, v := range resources {
		resp := resource.MetadataResponse{}
		v().Metadata(context.Background(), resource.MetadataRequest{}, &resp)
		if err := g.GenerateExample(resp.TypeName); err != nil {
			g.Fatalf("error generating Terraform %s import example: %s", resp.TypeName, err)
		}
	}
}

type Generator struct {
	*common.Generator
}

func NewGenerator() *Generator {
	return &Generator{
		Generator: common.NewGenerator(),
	}
}

func (g *Generator) GenerateExample(resourceName string) error {
	templateData := &TemplateData{
		ResourceType: resourceName,
	}

	filename := fmt.Sprintf("./examples/resources/%s/import.sh", resourceName)
	d := g.NewUnformattedFileDestination(filename)

	if err := d.CreateDirectories(); err != nil {
		return err
	}

	if err := d.WriteTemplate("resource", importExampleTemplateBody, templateData); err != nil {
		return err
	}

	if err := d.Write(); err != nil {
		return err
	}

	return nil
}

type TemplateData struct {
	ResourceType string
}

var importExampleTemplateBody = `$ terraform import {{ .ResourceType }}.example <resource ID>`
