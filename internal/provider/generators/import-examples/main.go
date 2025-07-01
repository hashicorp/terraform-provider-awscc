// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build generate
// +build generate

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
	json "github.com/json-iterator/go"
)

var (
	file = flag.String("file", "", "File containing import data in JSON list format")
)

type FileData struct {
	Resource   string
	Identifier []string
	Path       string
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -file <file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *file == "" {
		flag.Usage()
		os.Exit(2)
	}

	f, err := os.Open(*file)
	if err != nil {
		os.Exit(2)
	}
	defer f.Close()

	var data []FileData
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	g := NewGenerator()

	for _, v := range data {
		if err := g.GenerateExample(v.Resource, v.Path, v.Identifier); err != nil {
			g.Fatalf("error generating Terraform %s import example: %s", v.Resource, err)
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

func (g *Generator) GenerateExample(resourceName, directory string, identifier []string) error {
	g.Infof("generating Terraform import code for %[1]q ", resourceName)
	templateData := &TemplateData{
		ResourceType: resourceName,
		Identifier:   formatIdentifier(identifier),
	}

	for _, v := range filesData {
		if err := createFile(g, v.filename(directory), v.templateBody, templateData); err != nil {
			return err
		}
	}

	return nil
}

func formatIdentifier(identifier []string) string {
	if len(identifier) != 0 {
		var out []string
		for _, i := range identifier {
			out = append(out, toSnake(i))
		}

		return fmt.Sprintf("\"%s\"", strings.Join(out, "|"))
	}

	return "<resource ID>"
}

type fileData struct {
	filename     func(string) string
	templateBody string
}

var filesData = []fileData{
	{
		filename:     func(directory string) string { return fmt.Sprintf("%simport.sh", directory) },
		templateBody: importExampleTemplateBody,
	},
	{
		filename:     func(directory string) string { return fmt.Sprintf("%simport-by-string-id.tf", directory) },
		templateBody: importExampleTemplateByStringIDBody,
	},
}

func createFile(g *Generator, filename, templateBody string, templateData *TemplateData) error {
	d := g.NewUnformattedFileDestination(filename)

	if err := d.CreateDirectories(); err != nil {
		return err
	}

	if err := d.WriteTemplate("resource", templateBody, templateData); err != nil {
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

var importExampleTemplateByStringIDBody = `import {
  to = {{ .ResourceType }}.example
  id = {{ .Identifier }}
}`
