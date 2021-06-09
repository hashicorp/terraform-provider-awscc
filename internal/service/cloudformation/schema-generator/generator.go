package generator

import (
	"fmt"
	"io"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/iancoleman/strcase"
)

type Generator struct {
	CfResource *cfschema.Resource
	Writer     io.Writer
}

// AppendCfDefinition appends the code generated for a CloudFormation definition.
func (g *Generator) AppendCfDefinition(name string, property *cfschema.Property) {
	g.appendCfProperty(name, CfDefinitionTfAttributeVariableName(name), property)
}

// AppendCfRootProperty appends the code generated for a CloudFormation root property.
func (g *Generator) AppendCfRootProperty(name string, property *cfschema.Property) {
	g.appendCfProperty(name, CfPropertyTfAttributeVariableName(name), property)
}

func (g *Generator) appendCfProperty(name, attributeVariableName string, property *cfschema.Property) {
	g.printf("\n")
	g.printf("// %s\n", name)
	g.printf("/*\n")
	g.printf("%v\n", property)
	g.printf("*/\n")
	g.printf("%s := schema.Attribute{}\n", attributeVariableName)
}

// printf writes a formatted string to the underlying writer.
func (g *Generator) printf(format string, a ...interface{}) (int, error) {
	return io.WriteString(g.Writer, fmt.Sprintf(format, a...))
}

// CfDefinitionTfAttributeVariableName returns a CloudFormation property's Terraform Attribute variable name.
func CfDefinitionTfAttributeVariableName(name string) string {
	return fmt.Sprintf("defn%sAttribute", name)
}

// CfPropertyTfAttributeVariableName returns a CloudFormation property's Terraform Attribute variable name.
func CfPropertyTfAttributeVariableName(name string) string {
	return fmt.Sprintf("prop%sAttribute", name)
}

// TfPropertyAttributeName returns a CloudFormation property's Terraform Attribute name.
func CfPropertyTfAttributeName(name string) string {
	return strcase.ToSnake(name)
}
