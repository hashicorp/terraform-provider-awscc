// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfResourceType   = flag.String("resource", "", "Terraform resource type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -resource <TF-resource-type> -cfschema <CF-type-schema-file> <generated-schema-file> <generated-acctests-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 || *tfResourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	schemaFilename := args[0]
	acctestsFilename := args[1]

	g := NewGenerator()

	if err := g.Generate(destinationPackage, schemaFilename, acctestsFilename); err != nil {
		g.Fatalf("error generating Terraform %s resource: %s", *tfResourceType, err)
	}
}

type Generator struct {
	*common.Generator
	cfTypeSchemaFile string
	tfResourceType   string
}

func NewGenerator() *Generator {
	return &Generator{
		Generator:        common.NewGenerator(),
		cfTypeSchemaFile: *cfTypeSchemaFile,
		tfResourceType:   *tfResourceType,
	}
}

// Generate generates the resource's type factory into the specified file.
func (g *Generator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	g.Infof("generating Terraform resource code for %[1]q from %[2]q into %[3]q and %[4]q", g.tfResourceType, g.cfTypeSchemaFile, schemaFilename, acctestsFilename)

	templateData, err := shared.GenerateTemplateData(g.UI(), g.cfTypeSchemaFile, shared.ResourceType, g.tfResourceType, packageName)

	if err != nil {
		return err
	}

	d := g.NewGoFileDestination(schemaFilename)

	if err := d.CreateDirectories(); err != nil {
		return err
	}

	if err := d.WriteTemplate("resource", resourceSchemaTemplateBody, templateData); err != nil {
		return err
	}

	if err := d.Write(); err != nil {
		return err
	}

	d = g.NewGoFileDestination(acctestsFilename)

	if err := d.CreateDirectories(); err != nil {
		return err
	}

	if err := d.WriteTemplate("acctest", acceptanceTestsTemplateBody, templateData); err != nil {
		return err
	}

	if err := d.Write(); err != nil {
		return err
	}

	return nil
}

// Terraform resource schema definition.
//
//go:embed schema.tmpl
var resourceSchemaTemplateBody string

// Terraform acceptance tests.
//
//go:embed tests.tmpl
var acceptanceTestsTemplateBody string
