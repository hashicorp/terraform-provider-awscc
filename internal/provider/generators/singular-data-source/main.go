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
	tfDataSourceType = flag.String("data-source", "", "Terraform data source type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -data-source <TF-data-source-type> -cfschema <CF-type-schema-file> <generated-schema-file> <generated-acctests-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 || *tfDataSourceType == "" || *cfTypeSchemaFile == "" {
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

	generator := &SingularDataSourceGenerator{
		cfTypeSchemaFile: *cfTypeSchemaFile,
		tfDataSourceType: *tfDataSourceType,
		Generator: shared.Generator{
			UI: ui,
		},
	}

	if err := generator.Generate(destinationPackage, schemaFilename, acctestsFilename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform %s data source: %s", *tfDataSourceType, err))
		os.Exit(1)
	}
}

type SingularDataSourceGenerator struct {
	cfTypeSchemaFile string
	tfDataSourceType string
	shared.Generator
}

// Generate generates the singular data source's type factory into the specified file.
func (s *SingularDataSourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	s.Infof("generating Terraform data source code for %[1]q from %[2]q into %[3]q and %[4]q", s.tfDataSourceType, s.cfTypeSchemaFile, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, shared.DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	templateData, err := s.GenerateTemplateData(s.cfTypeSchemaFile, shared.DataSourceType, s.tfDataSourceType, packageName)

	if err != nil {
		return err
	}

	err = s.ApplyAndWriteTemplate(schemaFilename, dataSourceSchemaTemplateBody, templateData)

	if err != nil {
		return err
	}

	err = s.ApplyAndWriteTemplate(acctestsFilename, acceptanceTestsTemplateBody, templateData)

	if err != nil {
		return err
	}

	return nil
}

// Terraform data source schema definition.
//
//go:embed schema.tmpl
var dataSourceSchemaTemplateBody string

// Terraform acceptance tests.
//
//go:embed tests.tmpl
var acceptanceTestsTemplateBody string
