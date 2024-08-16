// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
)

var (
	identifier   = flag.String("identifier", "", "Primary identifier for the resource")
	resourceName = flag.String("resource", "", "Resource name")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -identifier <primary-identifier>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *resourceName == "" {
		flag.Usage()
		os.Exit(2)
	}

	g := NewGenerator()
	//resources := provider.New().Resources(context.Background())

	//for _, v := range resources {
	//	resp := resource.MetadataResponse{}
	//	v().Metadata(context.Background(), resource.MetadataRequest{}, &resp)
	//	if err := g.GenerateExample(resp.TypeName, *identifier); err != nil {
	//		g.Fatalf("error generating Terraform %s import example: %s", resp.TypeName, err)
	//	}
	//}

	if err := g.GenerateExample(*resourceName, *identifier); err != nil {
		g.Fatalf("error generating Terraform %s import example: %s", resourceName, err)
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

func (g *Generator) GenerateExample(resourceName, identifier string) error {
	templateData := &TemplateData{
		ResourceType: resourceName,
		Identifier:   identifier,
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
	Identifier   string
}

var importExampleTemplateBody = `$ terraform import {{ .ResourceType }}.example {{ .Identifier }}`
