package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	getter "github.com/hashicorp/go-getter"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/mitchellh/cli"
)

type Config struct {
	MetaSchema      MetaSchema       `hcl:"meta_schema,block"`
	ResourceSchemas []ResourceSchema `hcl:"resource_schema,block"`
}

type MetaSchema struct {
	Local   string `hcl:"local"`
	Refresh bool   `hcl:"refresh,optional"`
	Source  Source `hcl:"source,block"`
}

type ResourceSchema struct {
	Local        string `hcl:"local"`
	Refresh      bool   `hcl:"refresh,optional"`
	ResourceName string `hcl:"resource_name,label"`
	Source       Source `hcl:"source,block"`
}

type Source struct {
	Url string `hcl:"url"`
}

var (
	configFile = flag.String("config", "", "configuration file; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -config <configuration-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *configFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	tempDirectory, err := ioutil.TempDir("", "*")

	if err != nil {
		ui.Error(fmt.Sprintf("error creating temporary directory: %s", err))
		os.Exit(1)
	}

	defer os.RemoveAll(tempDirectory)

	downloader := &Downloader{
		baseDir: filepath.Dir(*configFile),
		ui:      ui,
	}

	err = hclsimple.DecodeFile(*configFile, nil, &downloader.config)
	if err != nil {
		ui.Error(fmt.Sprintf("error loading configuration: %s", err))
		os.Exit(1)
	}

	if err := downloader.MetaSchema(); err != nil {
		ui.Error(fmt.Sprintf("error loading CloudFormation Resource Provider Definition Schema: %s", err))
		os.Exit(1)
	}

	/*
		metaSchemaFilename, err := filepath.Abs(filepath.Join(baseDir, config.MetaSchema.Local))
		if err != nil {
			log.Printf("error making absolute path: %s", err)
			os.Exit(1)
		}
		metaSchemaFileExists := fileExists(metaSchemaFilename)

		if !metaSchemaFileExists || config.MetaSchema.Refresh {
			src := config.MetaSchema.Source.Url
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
			resourceSchemaFilename, err := filepath.Abs(filepath.Join(baseDir, schema.Local))
			if err != nil {
				log.Printf("error making absolute path: %s", err)
				continue
			}
			resourceSchemaFileExists := fileExists(resourceSchemaFilename)

			if !resourceSchemaFileExists || schema.Refresh {
				src := schema.Source.Url
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
	*/
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

type Downloader struct {
	baseDir    string
	config     Config
	metaSchema *cfschema.MetaJsonSchema
	ui         cli.Ui
}

func (d *Downloader) MetaSchema() error {
	metaSchemaFilename, err := filepath.Abs(filepath.Join(d.baseDir, d.config.MetaSchema.Local))

	if err != nil {
		return fmt.Errorf("error making absolute path: %w", err)
	}

	metaSchemaFileExists := fileExists(metaSchemaFilename)

	if !metaSchemaFileExists || d.config.MetaSchema.Refresh {
		src := d.config.MetaSchema.Source.Url
		d.Infof("downloading CloudFormation Resource Provider Definition Schema %s to %s", src, metaSchemaFilename)

		if err := getter.GetFile(metaSchemaFilename, src); err != nil {
			return fmt.Errorf("error downloading: %w", err)
		}
	}

	d.metaSchema, err = cfschema.NewMetaJsonSchemaPath(metaSchemaFilename)

	if err != nil {
		return fmt.Errorf("error loading CloudFormation Resource Provider Definition Schema: %w", err)
	}

	return nil
}

func (d *Downloader) Infof(format string, a ...interface{}) {
	d.ui.Info(fmt.Sprintf(format, a...))
}
