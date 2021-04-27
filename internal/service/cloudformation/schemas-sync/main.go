package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	getter "github.com/hashicorp/go-getter"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Schemas []Schema `hcl:"schema,block"`
}

type Schema struct {
	ResourceType string `hcl:"resource_type"`
	Source       string `hcl:"source"`
}

func main() {
	var configFilename string

	flag.StringVar(&configFilename, "config", "config.hcl", "configuration file name")
	flag.Parse()

	var config Config

	err := hclsimple.DecodeFile(configFilename, nil, &config)
	if err != nil {
		log.Printf("error loading configuration: %s", err)
		os.Exit(1)
	}

	tempDirectory, err := ioutil.TempDir("", "*")
	if err != nil {
		log.Printf("error creating temporary directory: %s", err)
		os.Exit(1)
	}

	defer os.RemoveAll(tempDirectory)

	for _, schema := range config.Schemas {
		src := schema.Source
		dst := filepath.Join(tempDirectory, schema.ResourceType+".clouformation-resource.json")
		log.Printf("downloading %s to %s\n", src, dst)

		if err := getter.GetFile(dst, src); err != nil {
			log.Printf("error downloading file: %s", err)
		}

		// Validate
	}

}
