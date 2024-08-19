// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
)

var (
	identifier   = flag.String("identifier", "", "Primary identifier for the resource")
	resourceName = flag.String("resource", "", "Resource name")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -identifier <primary-identifier> -resource <resource-name> -- filename\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	filename := args[0]

	if *resourceName == "" || filename == "" {
		flag.Usage()
		os.Exit(2)
	}

	g := NewGenerator()

	if err := g.GenerateExample(*resourceName, *identifier, filename); err != nil {
		g.Fatalf("error generating Terraform %s import example: %s", *resourceName, err)
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

func (g *Generator) GenerateExample(resourceName, identifier, filename string) error {
	g.Infof("generating Terraform import code for %[1]q ", resourceName)
	templateData := &TemplateData{
		ResourceType: resourceName,
		Identifier:   "<resource ID>",
	}

	if identifier != "" {
		ident := strings.Split(identifier, ",")
		var out []string
		for _, i := range ident {
			out = append(out, toSnake(i))
		}

		templateData.Identifier = fmt.Sprintf("\"%s\"", strings.Join(out, "|"))
	}

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

func toSnake(s string) string {
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

type TemplateData struct {
	ResourceType string
	Identifier   string
}

var importExampleTemplateBody = `$ terraform import {{ .ResourceType }}.example {{ .Identifier }}`
