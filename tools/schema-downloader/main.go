package main

import (
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	getter "github.com/hashicorp/go-getter"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	MetaSchema      MetaSchema       `hcl:"meta_schema,block"`
	ResourceSchemas []ResourceSchema `hcl:"resource_schema,block"`
}

type MetaSchema struct {
	Destination string `hcl:"destination"`
	Refresh     bool   `hcl:"refresh,optional"`
	Source      string `hcl:"source"`
}

type ResourceSchema struct {
	Destination  string `hcl:"destination"`
	Refresh      bool   `hcl:"refresh,optional"`
	ResourceName string `hcl:"resource_name,label"`
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

	baseDir := filepath.Dir(configFilename)

	tempDirectory, err := ioutil.TempDir("", "*")
	if err != nil {
		log.Printf("error creating temporary directory: %s", err)
		os.Exit(1)
	}

	defer os.RemoveAll(tempDirectory)

	metaSchemaFilename, err := filepath.Abs(filepath.Join(baseDir, config.MetaSchema.Destination))
	if err != nil {
		log.Printf("error making absolute path: %s", err)
		os.Exit(1)
	}
	metaSchemaFileExists := fileExists(metaSchemaFilename)

	if !metaSchemaFileExists || config.MetaSchema.Refresh {
		src := config.MetaSchema.Source
		log.Printf("downloading meta-schema %s to %s", src, metaSchemaFilename)
		if err := getter.GetFile(metaSchemaFilename, src); err != nil {
			log.Printf("error downloading: %s", err)
			os.Exit(1)
		}
	}

	metaSchema, err := cfschema.NewMetaJsonSchemaPath(metaSchemaFilename)
	if err != nil {
		log.Printf("error loading meta-schema: %s", err)
		os.Exit(1)
	}

	for _, schema := range config.ResourceSchemas {
		resourceSchemaFilename, err := filepath.Abs(filepath.Join(baseDir, schema.Destination))
		if err != nil {
			log.Printf("error making absolute path: %s", err)
			continue
		}
		resourceSchemaFileExists := fileExists(resourceSchemaFilename)

		if !resourceSchemaFileExists || schema.Refresh {
			src := schema.Source
			dst := filepath.Join(tempDirectory, filepath.Base(resourceSchemaFilename))

			log.Printf("downloading resource schema %s to %s", src, dst)
			if err := getter.GetFile(dst, src); err != nil {
				log.Printf("error downloading: %s", err)
				continue
			}

			resourceSchema, err := cfschema.NewResourceJsonSchemaPath(dst)
			if err != nil {
				log.Printf("error loading %s: %s", dst, err)
				continue
			}

			if err := metaSchema.ValidateResourceJsonSchema(resourceSchema); err != nil {
				log.Printf("error validating %s: %s", dst, err)
				continue
			}

			if err := copyFile(resourceSchemaFilename, dst); err != nil {
				log.Printf("error copying: %s", err)
				continue
			}
		}
	}
}

var errCopyFileWithDir = errors.New("dir argument to CopyFile")

// copyFile copies the file with path src to dst. The new file must not exist.
// It is created with the same permissions as src.
func copyFile(dst, src string) error {
	rf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer rf.Close()
	rstat, err := rf.Stat()
	if err != nil {
		return err
	}
	if rstat.IsDir() {
		return errCopyFileWithDir
	}

	wf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, rstat.Mode())
	if err != nil {
		return err
	}
	if _, err := io.Copy(wf, rf); err != nil {
		wf.Close()
		return err
	}
	return wf.Close()
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
