// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/provider/generators"
)

var (
	resourcePattern = flag.String("resource", `^.*$`, "regular expression matching the resources to generate")
	configFilename  = flag.String("config", "config.hcl", "name of the resource schema configuration file")
	outputDirectory = flag.String("output", ".", "output directory")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags]\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	resourceRegexp, err := regexp.Compile(*resourcePattern)

	if err != nil {
		log.Fatalf("error compiling resource matching regular expression: %s", err)
	}

	config, err := generators.NewConfigPath(*configFilename)

	if err != nil {
		log.Fatalf("error reading configuration file: %s", err)
	}

	for _, resourceSchema := range config.ResourceSchemas {
		resourceName := resourceSchema.ResourceName
		if !resourceRegexp.MatchString(resourceName) {
			continue
		}

		_, err := NewResourcePath(resourceSchema.Local)

		if err != nil {
			log.Fatalf("error reading resource schema (%s): %s", resourceName, err)
		}

	}
}

type Resource struct {
	inner *cfschema.Resource
}

func NewResourcePath(path string) (*Resource, error) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(path)

	if err != nil {
		return nil, fmt.Errorf("error reading CloudFormation Resource Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return nil, fmt.Errorf("error parsing CloudFormation Resource Schema: %w", err)
	}

	if err := resource.Expand(); err != nil {
		return nil, fmt.Errorf("error expanding JSON Pointer references: %w", err)
	}

	return &Resource{inner: resource}, nil
}
