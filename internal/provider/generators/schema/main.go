package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/mitchellh/cli"
)

type Config struct {
	Defaults        Defaults         `hcl:"defaults,block"`
	MetaSchema      MetaSchema       `hcl:"meta_schema,block"`
	ResourceSchemas []ResourceSchema `hcl:"resource_schema,block"`
}

type Defaults struct {
	SchemaCacheDirectory    string `hcl:"schema_cache_directory"`
	TerraformTypeNamePrefix string `hcl:"terraform_type_name_prefix,optional"`
}

type MetaSchema struct {
	Path string `hcl:"path"`
}

type ResourceSchema struct {
	CloudFormationSchemaPath             string `hcl:"cloudformation_schema_path,optional"`
	CloudFormationTypeName               string `hcl:"cloudformation_type_name"`
	ResourceTypeName                     string `hcl:"resource_type_name,label"`
	SuppressPluralDataSourceGeneration   bool   `hcl:"suppress_plural_data_source_generation,optional"`
	SuppressResourceGeneration           bool   `hcl:"suppress_resource_generation,optional"`
	SuppressSingularDataSourceGeneration bool   `hcl:"suppress_singular_data_source_generation,optional"`
}

var (
	configFile        = flag.String("config", "", "configuration file; required")
	generatedCodeRoot = flag.String("generated-code-root", "", "directory root for generated resource code")
	importPathRoot    = flag.String("import-path-root", "", "import path root for generated resource code; required")
	packageName       = flag.String("package", "", "override package name for generated code")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -config <configuration-file> -import-path-root <import-path-root> <generated-resources-file> <generated-singular-data-sources-file> <generated-plural-data-sources-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 3 || *configFile == "" || *importPathRoot == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	resourcesFilename := args[0]
	singularDatasourcesFilename := args[1]
	pluralDatasourcesFilename := args[2]

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

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		ui.Error(fmt.Sprintf("error loading AWS SDK config: %s", err))
		os.Exit(1)
	}

	client := cloudformation.NewFromConfig(cfg)

	downloader := &Downloader{
		client:        client,
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

	resources, dataSources, err := downloader.Schemas()

	if err != nil {
		ui.Error(fmt.Sprintf("error processing CloudFormation Resource Provider Schemas: %s", err))
		os.Exit(1)
	}

	generator := &Generator{
		ui: ui,
	}

	if err := generator.GenerateResources(destinationPackage, resourcesFilename, generatedCodeRootDirectoryName, *importPathRoot, resources); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform resource generation instructions: %s", err))
		os.Exit(1)
	}

	if err := generator.GenerateDataSources(destinationPackage, singularDatasourcesFilename, generatedCodeRootDirectoryName, *importPathRoot, dataSources.Singular); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform singular data-source generation instructions: %s", err))
		os.Exit(1)
	}

	if err := generator.GenerateDataSources(destinationPackage, pluralDatasourcesFilename, generatedCodeRootDirectoryName, *importPathRoot, dataSources.Plural); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform plural data-source generation instructions: %s", err))
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
	client        *cloudformation.Client
	config        Config
	metaSchema    *cfschema.MetaJsonSchema
	tempDirectory string
	ui            cli.Ui
}

func (d *Downloader) MetaSchema() error {
	d.infof("using CloudFormation Resource Provider Definition Schema %q", d.config.MetaSchema.Path)

	metaSchema, err := cfschema.NewMetaJsonSchemaPath(d.config.MetaSchema.Path)

	if err != nil {
		return fmt.Errorf("loading CloudFormation Resource Provider Definition Schema: %w", err)
	}

	d.metaSchema = metaSchema

	return nil
}

func (d *Downloader) Schemas() ([]*ResourceData, *DataSources, error) {
	if d.metaSchema == nil {
		return nil, nil, fmt.Errorf("CloudFormation Resource Provider Definition Schema not loaded")
	}

	terraformTypeNamePrefix := d.config.Defaults.TerraformTypeNamePrefix

	resources := []*ResourceData{}
	var singularDataSources, pluralDataSources []*DataSourceData

	for _, schema := range d.config.ResourceSchemas {
		cfResourceSchemaFilename, cfResourceTypeName, err := d.ResourceSchema(schema)

		if err != nil {
			d.ui.Warn(fmt.Sprintf("error loading CloudFormation Resource Provider Schema for %s: %s", schema.ResourceTypeName, err))
			continue
		}

		_, _, _, err = naming.ParseCloudFormationTypeName(cfResourceTypeName)

		if err != nil {
			d.ui.Warn(fmt.Sprintf("incorrect format for CloudFormation Resource Provider Schema type name: %s", cfResourceTypeName))
			continue
		}

		tfResourceTypeName := schema.ResourceTypeName
		org, svc, res, err := naming.ParseTerraformTypeName(tfResourceTypeName)

		if err != nil {
			d.ui.Warn(fmt.Sprintf("incorrect format for Terraform resource type name: %s", tfResourceTypeName))
			continue
		}

		if terraformTypeNamePrefix != "" {
			tfResourceTypeName = naming.CreateTerraformTypeName(terraformTypeNamePrefix, svc, res)
		}

		if schema.SuppressSingularDataSourceGeneration {
			d.ui.Info(fmt.Sprintf("generation of a Terraform singular data source schema for %s has been suppressed", tfResourceTypeName))
		} else {
			singularDataSources = append(singularDataSources, &DataSourceData{
				CloudFormationTypeSchemaFile: cfResourceSchemaFilename,
				GeneratedAccTestsFileName:    res + "_singular_data_source_gen_test",
				GeneratedCodeFileName:        res + "_singular_data_source_gen",
				GeneratedCodePackageName:     svc,
				GeneratedCodePathSuffix:      fmt.Sprintf("%s/%s", org, svc),
				TerraformResourceType:        tfResourceTypeName,
			})
		}

		if schema.SuppressPluralDataSourceGeneration {
			d.ui.Info(fmt.Sprintf("generation of a Terraform plural data source schema for %s has been suppressed", tfResourceTypeName))
		} else {
			pluralTfResourceTypeName := naming.Pluralize(tfResourceTypeName)

			pluralDataSources = append(pluralDataSources, &DataSourceData{
				CloudFormationType:        cfResourceTypeName,
				GeneratedAccTestsFileName: res + "_plural_data_source_gen_test",
				GeneratedCodeFileName:     res + "_plural_data_source_gen",
				GeneratedCodePackageName:  svc,
				GeneratedCodePathSuffix:   fmt.Sprintf("%s/%s", org, svc),
				TerraformResourceType:     pluralTfResourceTypeName,
			})
		}

		if schema.SuppressResourceGeneration {
			d.ui.Info(fmt.Sprintf("generation of a Terraform resource schema for %s has been suppressed", tfResourceTypeName))
			continue
		}

		resources = append(resources, &ResourceData{
			CloudFormationTypeSchemaFile: cfResourceSchemaFilename,
			GeneratedAccTestsFileName:    res + "_resource_gen_test",     // e.g. "log_group_resource_gen_test"
			GeneratedCodeFileName:        res + "_resource_gen",          // e.g. "log_group_resource_gen"
			GeneratedCodePackageName:     svc,                            // e.g. "logs"
			GeneratedCodePathSuffix:      fmt.Sprintf("%s/%s", org, svc), // e.g. "aws/logs"
			TerraformResourceType:        tfResourceTypeName,
		})
	}

	dataSources := &DataSources{
		Singular: singularDataSources,
		Plural:   pluralDataSources,
	}

	return resources, dataSources, nil
}

// ResourceSchema returns the local resource schema file name and type name.
func (d *Downloader) ResourceSchema(schema ResourceSchema) (string, string, error) {
	resourceSchemaFilename := schema.CloudFormationSchemaPath
	if resourceSchemaFilename == "" {
		filename := fmt.Sprintf("%s.json", schema.CloudFormationTypeName)
		// Replace all '::'s in the filename.
		filename = strings.ReplaceAll(filename, "::", "_")
		resourceSchemaFilename = path.Join(d.config.Defaults.SchemaCacheDirectory, filename)
	}

	resourceSchemaFileExists := fileExists(resourceSchemaFilename)

	if !resourceSchemaFileExists {
		dst := filepath.Join(d.tempDirectory, filepath.Base(resourceSchemaFilename))

		d.infof("downloading CloudFormation Resource Provider Schema to %q", dst)

		input := &cloudformation.DescribeTypeInput{
			Type:     types.RegistryTypeResource,
			TypeName: aws.String(schema.CloudFormationTypeName),
		}
		output, err := d.client.DescribeType(context.TODO(), input)

		if err != nil {
			return "", "", fmt.Errorf("describing CloudFormation type: %w", err)
		}

		schema, err := cfschema.Sanitize(aws.ToString(output.Schema))

		if err != nil {
			return "", "", fmt.Errorf("sanitizing schema: %w", err)
		}

		err = ioutil.WriteFile(dst, []byte(schema), 0644) //nolint:gomnd

		if err != nil {
			return "", "", fmt.Errorf("writing schema to %q: %w", dst, err)
		}

		resourceSchema, err := cfschema.NewResourceJsonSchemaPath(dst)

		if err != nil {
			return "", "", fmt.Errorf("loading %s: %w", dst, err)
		}

		if err := d.metaSchema.ValidateResourceJsonSchema(resourceSchema); err != nil {
			return "", "", fmt.Errorf("validating %s: %w", dst, err)
		}

		if err := copyFile(resourceSchemaFilename, dst); err != nil {
			return "", "", fmt.Errorf("copying: %w", err)
		}
	} else {
		d.infof("using cached CloudFormation Resource Provider Schema %q", resourceSchemaFilename)
	}

	// Read the resource type name from the schema.
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(resourceSchemaFilename)

	if err != nil {
		return "", "", fmt.Errorf("loading %s: %w", resourceSchemaFilename, err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return "", "", fmt.Errorf("parsing %s: %w", resourceSchemaFilename, err)
	}

	return resourceSchemaFilename, *resource.TypeName, nil
}

func (d *Downloader) infof(format string, a ...interface{}) {
	d.ui.Info(fmt.Sprintf(format, a...))
}

type ResourceData struct {
	CloudFormationTypeSchemaFile string
	GeneratedAccTestsFileName    string
	GeneratedCodeFileName        string
	GeneratedCodePackageName     string
	GeneratedCodePathSuffix      string
	TerraformResourceType        string
}

type DataSourceData struct {
	CloudFormationType           string
	CloudFormationTypeSchemaFile string
	GeneratedAccTestsFileName    string
	GeneratedCodeFileName        string
	GeneratedCodePackageName     string
	GeneratedCodePathSuffix      string
	TerraformResourceType        string
}

type DataSources struct {
	Singular []*DataSourceData
	Plural   []*DataSourceData
}

type Generator struct {
	ui cli.Ui
}

func (g *Generator) GenerateResources(packageName, filename, generatedCodeRootDirectoryName, importPathRoot string, resources []*ResourceData) error {
	g.infof("generating Terraform resource generation instructions into %q", filename)

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

	tmpl, err := template.New("function").Parse(resourceTemplateBody)

	if err != nil {
		return fmt.Errorf("parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		return fmt.Errorf("formatting generated file: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("writing to file (%s): %w", filename, err)
	}

	return nil
}

func (g *Generator) GenerateDataSources(packageName, filename, generatedCodeRootDirectoryName, importPathRoot string, dataSources []*DataSourceData) error {
	g.infof("generating Terraform data-source generation instructions into %q", filename)

	importPaths := make(map[string]struct{}) // Set of strings.

	for _, dataSource := range dataSources {
		if _, ok := importPaths[dataSource.GeneratedCodePathSuffix]; !ok {
			importPaths[dataSource.GeneratedCodePathSuffix] = struct{}{}
		}
	}

	importPathSuffixes := make([]string, 0)

	for importPathSuffix := range importPaths {
		importPathSuffixes = append(importPathSuffixes, importPathSuffix)
	}

	sort.Strings(importPathSuffixes)

	templateData := TemplateData{
		DataSources:                    dataSources,
		GeneratedCodeRootDirectoryName: generatedCodeRootDirectoryName,
		ImportPathRoot:                 importPathRoot,
		ImportPathSuffixes:             importPathSuffixes,
		PackageName:                    packageName,
	}

	tmpl, err := template.New("function").Parse(dataSourceTemplateBody)

	if err != nil {
		return fmt.Errorf("parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		return fmt.Errorf("formatting generated file: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("writing to file (%s): %w", filename, err)
	}

	return nil
}

func (g *Generator) infof(format string, a ...interface{}) {
	g.ui.Info(fmt.Sprintf(format, a...))
}

var resourceTemplateBody = `
// Code generated by generators/schema/main.go; DO NOT EDIT.

{{- range .Resources }}
//go:generate go run generators/resource/main.go -resource {{ .TerraformResourceType }} -cfschema {{ .CloudFormationTypeSchemaFile }} -package {{ .GeneratedCodePackageName }} -- {{ $.GeneratedCodeRootDirectoryName }}/{{ .GeneratedCodePathSuffix }}/{{ .GeneratedCodeFileName }}.go {{ $.GeneratedCodeRootDirectoryName }}/{{ .GeneratedCodePathSuffix }}/{{ .GeneratedAccTestsFileName }}.go
{{- end }}

package {{ .PackageName }}

import (
{{- range .ImportPathSuffixes }}
	_ "{{ $.ImportPathRoot }}/{{ . }}"
{{- end }}
)
`

var dataSourceTemplateBody = `
// Code generated by generators/schema/main.go; DO NOT EDIT.

{{- range .DataSources }}
{{ if .CloudFormationType }}//go:generate go run generators/plural-data-source/main.go -data-source {{ .TerraformResourceType }} -cftype {{ .CloudFormationType }} -package {{ .GeneratedCodePackageName }} {{ $.GeneratedCodeRootDirectoryName }}/{{ .GeneratedCodePathSuffix }}/{{ .GeneratedCodeFileName }}.go {{ $.GeneratedCodeRootDirectoryName }}/{{ .GeneratedCodePathSuffix }}/{{ .GeneratedAccTestsFileName }}.go{{ else }}//go:generate go run generators/singular-data-source/main.go -data-source {{ .TerraformResourceType }} -cfschema {{ .CloudFormationTypeSchemaFile }} -package {{ .GeneratedCodePackageName }} {{ $.GeneratedCodeRootDirectoryName }}/{{ .GeneratedCodePathSuffix }}/{{ .GeneratedCodeFileName }}.go {{ $.GeneratedCodeRootDirectoryName }}/{{ .GeneratedCodePathSuffix }}/{{ .GeneratedAccTestsFileName }}.go{{- end }}
{{- end }}

package {{ .PackageName }}

import (
{{- range .ImportPathSuffixes }}
	_ "{{ $.ImportPathRoot }}/{{ . }}"
{{- end }}
)
`

type TemplateData struct {
	DataSources                    []*DataSourceData
	GeneratedCodeRootDirectoryName string
	ImportPathRoot                 string
	ImportPathSuffixes             []string
	PackageName                    string
	Resources                      []*ResourceData
}
