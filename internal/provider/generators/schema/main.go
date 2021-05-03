// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	generator "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/clouformation/schema-generator"
)

var (
	resourceType     = flag.String("resource", "", "Terraform resource type; required")
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	outputDirectory  = flag.String("output", ".", "output directory")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -resource <resource-type> -cfschema <CF-type-schema-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *resourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	resource, err := NewResourcePath(*resourceType, *cfTypeSchemaFile)

	if err != nil {
		log.Fatalf("error reading CloudFormation resource schema for %s: %s", *resourceType, err)
	}

	filename := filepath.Join(*outputDirectory, resource.TfType+"_schema_gen.go")

	f, err := os.Create(filename)

	if err != nil {
		log.Fatalf("error creating file (%s): %s", filename, err)
	}

	defer f.Close()

	for propertyName := range resource.CfResource.Properties {
		log.Printf("%s\n\n%s\n\n", propertyName, generator.RootPropertySchema(resource.CfResource, propertyName))
	}
}

type Resource struct {
	CfResource *cfschema.Resource
	TfType     string
}

func NewResourcePath(resourceType, cfTypeSchemaFile string) (*Resource, error) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("error reading CloudFormation Resource Type Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return nil, fmt.Errorf("error parsing CloudFormation Resource Type Schema: %w", err)
	}

	if err := resource.Expand(); err != nil {
		return nil, fmt.Errorf("error expanding JSON Pointer references: %w", err)
	}

	return &Resource{CfResource: resource, TfType: resourceType}, nil
}
