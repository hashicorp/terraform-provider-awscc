package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

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
	configFile        = flag.String("config", "", "configuration file; required")
	generatedCodeRoot = flag.String("generated-code-root", "", "directory root for generated resource code")
	importPathRoot    = flag.String("import-path-root", "", "import path root for generated resource code; required")
	packageName       = flag.String("package", "", "override package name for generated code")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -config <configuration-file> -import-path-root <import-path-root> <generated-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || *configFile == "" || *importPathRoot == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	filename := args[0]

	generatedCodeRootDirectoryName := "."
	if *generatedCodeRoot != "" {
		generatedCodeRootDirectoryName = *generatedCodeRoot
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
		baseDir:       filepath.Dir(*configFile),
		tempDirectory: tempDirectory,
		ui:            ui,
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

	resources, err := downloader.ResourceSchemas()

	if err != nil {
		ui.Error(fmt.Sprintf("error processing CloudFormation Resource Provider Schemas: %s", err))
		os.Exit(1)
	}

	generator := &Generator{
		ui: ui,
	}

	if err := generator.Generate(destinationPackage, filename, generatedCodeRootDirectoryName, *importPathRoot, resources); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform resource generation instructions: %s", err))
		os.Exit(1)
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

type Downloader struct {
	baseDir       string
	config        Config
	metaSchema    *cfschema.MetaJsonSchema
	tempDirectory string
	ui            cli.Ui
}

func (d *Downloader) MetaSchema() error {
	metaSchemaFilename, err := filepath.Abs(filepath.Join(d.baseDir, d.config.MetaSchema.Local))

	if err != nil {
		return fmt.Errorf("error making absolute path: %w", err)
	}

	metaSchemaFileExists := fileExists(metaSchemaFilename)

	if !metaSchemaFileExists || d.config.MetaSchema.Refresh {
		src := d.config.MetaSchema.Source.Url
		d.Infof("downloading CloudFormation Resource Provider Definition Schema %q to %q", src, metaSchemaFilename)

		if err := getter.GetFile(metaSchemaFilename, src); err != nil {
			return fmt.Errorf("error downloading: %w", err)
		}
	} else {
		d.Infof("using cached CloudFormation Resource Provider Definition Schema %q", metaSchemaFilename)
	}

	d.metaSchema, err = cfschema.NewMetaJsonSchemaPath(metaSchemaFilename)

	if err != nil {
		return fmt.Errorf("error loading CloudFormation Resource Provider Definition Schema: %w", err)
	}

	return nil
}

func (d *Downloader) ResourceSchemas() ([]*ResourceData, error) {
	if d.metaSchema == nil {
		return nil, fmt.Errorf("CloudFormation Resource Provider Definition Schema not loaded")
	}

	resources := []*ResourceData{}

	for _, schema := range d.config.ResourceSchemas {
		resourceSchemaFilename, resourceTypeName, err := d.ResourceSchema(schema)

		if err != nil {
			d.ui.Warn(fmt.Sprintf("error loading CloudFormation Resource Provider Schema for %s: %s", schema.ResourceName, err))
			continue
		}

		// https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-typeName
		parts := strings.Split(resourceTypeName, "::")

		if len(parts) != 3 {
			d.ui.Warn(fmt.Sprintf("incorrect format for CloudFormation Resource Provider Schema type name: %s", resourceTypeName))
			continue
		}

		resources = append(resources, &ResourceData{
			CloudFormationTypeSchemaFile: resourceSchemaFilename,
			GeneratedCodePackageName:     strings.ToLower(parts[1]),                                 // e.g. "logs"
			GeneratedCodePathSuffix:      strings.ToLower(fmt.Sprintf("%s/%s", parts[0], parts[1])), // e.g. "aws/logs"
			TerraformResourceType:        schema.ResourceName,
		})
	}

	return resources, nil
}

// ResourceSchema returns the local resource schema file name and type name.
func (d *Downloader) ResourceSchema(schema ResourceSchema) (string, string, error) {
	resourceSchemaFilename, err := filepath.Abs(filepath.Join(d.baseDir, schema.Local))

	if err != nil {
		return "", "", fmt.Errorf("error making absolute path: %w", err)
	}

	resourceSchemaFileExists := fileExists(resourceSchemaFilename)

	if !resourceSchemaFileExists || schema.Refresh {
		src := schema.Source.Url
		dst := filepath.Join(d.tempDirectory, filepath.Base(resourceSchemaFilename))

		d.Infof("downloading CloudFormation Resource Provider Schema %q to %q", src, dst)

		if err := getter.GetFile(dst, src); err != nil {
			return "", "", fmt.Errorf("error downloading: %w", err)
		}

		resourceSchema, err := cfschema.NewResourceJsonSchemaPath(dst)

		if err != nil {
			return "", "", fmt.Errorf("error loading %s: %w", dst, err)
		}

		if err := d.metaSchema.ValidateResourceJsonSchema(resourceSchema); err != nil {
			return "", "", fmt.Errorf("error validating %s: %w", dst, err)
		}

		if err := copyFile(resourceSchemaFilename, dst); err != nil {
			return "", "", fmt.Errorf("error copying: %w", err)
		}
	} else {
		d.Infof("using cached CloudFormation Resource Provider Schema %q", resourceSchemaFilename)
	}

	// Read the resource type name from the schema.
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(resourceSchemaFilename)

	if err != nil {
		return "", "", fmt.Errorf("error loading %s: %w", resourceSchemaFilename, err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return "", "", fmt.Errorf("error parsing %s: %w", resourceSchemaFilename, err)
	}

	return resourceSchemaFilename, *resource.TypeName, nil
}

func (d *Downloader) Infof(format string, a ...interface{}) {
	d.ui.Info(fmt.Sprintf(format, a...))
}

type ResourceData struct {
	CloudFormationTypeSchemaFile string
	GeneratedCodePackageName     string
	GeneratedCodePathSuffix      string
	TerraformResourceType        string
}

type Generator struct {
	ui cli.Ui
}

func (g *Generator) Infof(format string, a ...interface{}) {
	g.ui.Info(fmt.Sprintf(format, a...))
}

func (g *Generator) Generate(packageName, filename, generatedCodeRootDirectoryName, importPathRoot string, resources []*ResourceData) error {
	g.Infof("generating Terraform resource generation instructions into %q", filename)

	importPaths := make(map[string]struct{}) // Set of strings.

	for _, resource := range resources {
		if _, ok := importPaths[resource.GeneratedCodePathSuffix]; !ok {
			importPaths[resource.GeneratedCodePathSuffix] = struct{}{}
		}
	}

	importPathSuffixes := make([]string, 0)

	for importPathSuffix := range importPaths {
		importPathSuffixes = append(importPathSuffixes, importPathSuffix)
	}

	sort.Strings(importPathSuffixes)

	templateData := TemplateData{
		GeneratedCodeRootDirectoryName: generatedCodeRootDirectoryName,
		ImportPathRoot:                 importPathRoot,
		ImportPathSuffixes:             importPathSuffixes,
		PackageName:                    packageName,
		Resources:                      resources,
	}

	tmpl, err := template.New("function").Parse(templateBody)

	if err != nil {
		return fmt.Errorf("error parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		return fmt.Errorf("error formatting generated file: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("error creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("error writing to file (%s): %w", filename, err)
	}

	return nil
}

var templateBody = `
// Code generated by generators/schema/main.go; DO NOT EDIT.

{{- range .Resources }}
//go:generate go run generators/resource/main.go -resource {{ .TerraformResourceType }} -cfschema {{ .CloudFormationTypeSchemaFile }} -package {{ .GeneratedCodePackageName }} -- {{ $.GeneratedCodeRootDirectoryName }}/{{ .GeneratedCodePathSuffix }}/{{ .TerraformResourceType }}_gen.go
{{- end }}

package {{ .PackageName }}

import (
{{- range .ImportPathSuffixes }}
	_ "{{ $.ImportPathRoot }}/{{ . }}"
{{- end }}
)
`

type TemplateData struct {
	GeneratedCodeRootDirectoryName string
	ImportPathRoot                 string
	ImportPathSuffixes             []string
	PackageName                    string
	Resources                      []*ResourceData
}
