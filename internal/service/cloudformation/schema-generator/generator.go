package generator

import (
	"fmt"
	"io"
	"sort"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/iancoleman/strcase"
)

type Generator struct {
	CfResource *cfschema.Resource
	Writer     io.Writer
}

// AppendCfDefinition generates the Terraform Plugin SDK code for a CloudFormation definition
// and appends the generated to code to the generator's Writer.
func (g *Generator) AppendCfDefinition(definitionName string, properties map[string]*cfschema.Property) {
	propertyNames := make([]string, 0)

	for propertyName := range properties {
		propertyNames = append(propertyNames, propertyName)
	}

	// Sort the property names to reduce generated code diffs.
	sort.Strings(propertyNames)

	// Generate and append each property.
	for _, propertyName := range propertyNames {
		g.appendCfProperty(definitionName, propertyName, properties[propertyName])
	}

	g.appendCfPropertyReferences(definitionName, propertyNames)
}

// appendCfPropertyReferences generates Go code that references the code generated for the
// specified CloudFormation properties and appends the generated to code to the generator's Writer.
func (g *Generator) appendCfPropertyReferences(definitionName string, propertyNames []string) {
	attributesVariableName := CfDefinitionTfAttributesVariableName(definitionName)

	g.printf("\n")
	g.printf("%s := make(map[string]schema.Attribute, %d)\n", attributesVariableName, len(propertyNames))
	for _, propertyName := range propertyNames {
		g.printf("%s[%q] = %s\n", attributesVariableName, CfPropertyTfAttributeName(propertyName), CfPropertyTfAttributeVariableName(definitionName, propertyName))
	}
}

// appendCfProperty generates Go code for a single CloudFormation property
// and appends the generated to code to the generator's Writer.
func (g *Generator) appendCfProperty(definitionName, propertyName string, property *cfschema.Property) {
	attributeVariableName := CfPropertyTfAttributeVariableName(definitionName, propertyName)

	g.printf("\n")
	g.printf("// Definition: %s\n", definitionName)
	g.printf("// Property: %s\n", propertyName)
	g.printf("// CloudFormation Resource Type Schema:\n")
	g.printf("/*\n")
	g.printf("%v\n", property)
	g.printf("*/\n")
	g.printf("%s := schema.Attribute{}\n", attributeVariableName)

	switch propertyType := property.Type.String(); propertyType {
	case cfschema.PropertyTypeBoolean:
		g.printf("%s.Type = types.BoolType\n", attributeVariableName)
	case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
		g.printf("%s.Type = types.NumberType\n", attributeVariableName)
	case cfschema.PropertyTypeString:
		g.printf("%s.Type = types.StringType\n", attributeVariableName)
	default:
		if ref := property.Ref; ref != nil {
			g.printf("%s.Attributes = schema.SingleNestedAttributes(%s)\n", attributeVariableName, CfDefinitionTfAttributesVariableName(ref.Field()))
		} else {
			g.printf("// Unsupported property type: %s\n", propertyType)
			return
		}
	}

	readOnly := g.CfResource.ReadOnlyProperties.ContainsPath([]string{propertyName})
	required := g.CfResource.IsRequired(propertyName)

	if required {
		g.printf("%s.Required = true\n", attributeVariableName)
	} else if !readOnly {
		g.printf("%s.Optional = true\n", attributeVariableName)
	}

	if readOnly && !required {
		g.printf("%s.Computed = true\n", attributeVariableName)
	}

	if description := property.Description; description != nil {
		g.printf("%s.Description = `%s`\n", attributeVariableName, *description)
	}
}

// printf writes a formatted string to the underlying writer.
func (g *Generator) printf(format string, a ...interface{}) (int, error) {
	return io.WriteString(g.Writer, fmt.Sprintf(format, a...))
}

// CfDefinitionTfAttributesVariableName returns a CloudFormation definition's Terraform map[string]Attribute variable name.
func CfDefinitionTfAttributesVariableName(definitionName string) string {
	return strcase.ToLowerCamel(fmt.Sprintf("%sAttributes", definitionName))
}

// CfPropertyTfAttributeVariableName returns a CloudFormation property's Terraform Attribute variable name.
func CfPropertyTfAttributeVariableName(definitionName, propertyName string) string {
	return strcase.ToLowerCamel(fmt.Sprintf("%s%sAttribute", definitionName, propertyName))
}

// TfPropertyAttributeName returns a CloudFormation property's Terraform Attribute name.
func CfPropertyTfAttributeName(propertyName string) string {
	return strcase.ToSnake(propertyName)
}
