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

	g := NewGenerator()

	if err := g.Generate(destinationPackage, schemaFilename, acctestsFilename); err != nil {
		g.Fatalf("error generating Terraform %s data source: %s", *tfDataSourceType, err)
	}
}

type Generator struct {
	*shared.Generator
	cfType           string
	tfDataSourceType string
}

func NewGenerator() *Generator {
	return &Generator{
		Generator:        shared.NewGenerator(),
		cfType:           *cfType,
		tfDataSourceType: *tfDataSourceType,
	}
}

// Generate generates the plural data source type's factory into the specified file.
func (g *Generator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	g.Infof("generating Terraform data source code for %[1]q into %[2]q and %[3]q", g.tfDataSourceType, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, shared.DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	org, svc, res, err := naming.ParseCloudFormationTypeName(g.cfType)

	if err != nil {
		return fmt.Errorf("incorrect format for CloudFormation Resource Provider Schema type name: %s", g.cfType)
	}

	ds := naming.PluralizeWithCustomNameSuffix(res, "Plural")

	factoryFunctionName := string(bytes.ToLower([]byte(ds[:1]))) + ds[1:] + shared.DataSourceType

	acceptanceTestFunctionPrefix := fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, ds)

	schemaDescription := fmt.Sprintf("Plural Data Source schema for %s", g.cfType)

	templateData := &shared.TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		CloudFormationTypeName:       g.cfType,
		FactoryFunctionName:          factoryFunctionName,
		PackageName:                  packageName,
		SchemaDescription:            schemaDescription,
		SchemaVersion:                1,
		TerraformTypeName:            g.tfDataSourceType,
	}

	d := g.NewGoFileDestination(schemaFilename)

	if err := d.WriteTemplate("data-source", dataSourceSchemaTemplateBody, templateData); err != nil {
		return err
	}

	if err := d.Write(); err != nil {
		return err
	}

	d = g.NewGoFileDestination(acctestsFilename)

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
var dataSourceSchemaTemplateBody string

// Terraform acceptance tests.
//
//go:embed tests.tmpl
var acceptanceTestsTemplateBody string
