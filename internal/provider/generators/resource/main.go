// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build ignore
// +build ignore

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
	"github.com/hashicorp/cli"
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

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	generator := &ResourceGenerator{
		cfTypeSchemaFile: *cfTypeSchemaFile,
		tfResourceType:   *tfResourceType,
		Generator: shared.Generator{
			UI: ui,
		},
	}

	if err := generator.Generate(destinationPackage, schemaFilename, acctestsFilename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform %s resource: %s", *tfResourceType, err))
		os.Exit(1)
	}
}

type ResourceGenerator struct {
	cfTypeSchemaFile string
	tfResourceType   string
	shared.Generator
}

// Generate generates the resource's type factory into the specified file.
func (r *ResourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	r.Infof("generating Terraform resource code for %[1]q from %[2]q into %[3]q and %[4]q", r.tfResourceType, r.cfTypeSchemaFile, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, shared.DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	templateData, err := r.GenerateTemplateData(r.cfTypeSchemaFile, shared.ResourceType, r.tfResourceType, packageName)

	if err != nil {
		return err
	}

	err = r.ApplyAndWriteTemplate(schemaFilename, resourceSchemaTemplateBody, templateData)

	if err != nil {
		return err
	}

	err = r.ApplyAndWriteTemplate(acctestsFilename, acceptanceTestsTemplateBody, templateData)

	if err != nil {
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
