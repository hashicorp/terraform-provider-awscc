// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
	"github.com/hashicorp/cli"
)

var (
	cfType           = flag.String("cftype", "", "CloudFormation resource type; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfDataSourceType = flag.String("data-source", "", "Terraform data source type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -data-source <TF-data-source-type> -cftype <CF-type> <generated-schema-file> <generated-acctests-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 || *tfDataSourceType == "" {
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

	generator := &PluralDataSourceGenerator{
		cfType:           *cfType,
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

type PluralDataSourceGenerator struct {
	cfType           string
	tfDataSourceType string
	shared.Generator
}

// Generate generates the plural data source type's factory into the specified file.
func (p *PluralDataSourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	p.Infof("generating Terraform data source code for %[1]q into %[2]q and %[3]q", p.tfDataSourceType, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, shared.DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	org, svc, res, err := naming.ParseCloudFormationTypeName(p.cfType)

	if err != nil {
		return fmt.Errorf("incorrect format for CloudFormation Resource Provider Schema type name: %s", p.cfType)
	}

	ds := naming.PluralizeWithCustomNameSuffix(res, "Plural")

	factoryFunctionName := string(bytes.ToLower([]byte(ds[:1]))) + ds[1:] + shared.DataSourceType

	acceptanceTestFunctionPrefix := fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, ds)

	schemaDescription := fmt.Sprintf("Plural Data Source schema for %s", p.cfType)

	templateData := shared.TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		CloudFormationTypeName:       p.cfType,
		FactoryFunctionName:          factoryFunctionName,
		PackageName:                  packageName,
		SchemaDescription:            schemaDescription,
		SchemaVersion:                1,
		TerraformTypeName:            p.tfDataSourceType,
	}

	err = p.ApplyAndWriteTemplate(schemaFilename, dataSourceSchemaTemplateBody, &templateData)

	if err != nil {
		return err
	}

	err = p.ApplyAndWriteTemplate(acctestsFilename, acceptanceTestsTemplateBody, &templateData)

	if err != nil {
		return err
	}

	return nil
}

// Terraform resource schema definition.
//
//go:embed schema.tmpl
var dataSourceSchemaTemplateBody string

// Terraform acceptance tests.
//
//go:embed tests.tmpl
var acceptanceTestsTemplateBody string
