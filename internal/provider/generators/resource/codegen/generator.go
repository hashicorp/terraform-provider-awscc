package codegen

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
func (g *Generator) AppendCfDefinition(definitionName string, properties map[string]*cfschema.Property) error {
	propertyNames := make([]string, 0)

	for propertyName := range properties {
		propertyNames = append(propertyNames, propertyName)
	}

	// Sort the property names to reduce generated code diffs.
	sort.Strings(propertyNames)

	// Generate and append each property.
	for _, propertyName := range propertyNames {
		if err := g.appendCfProperty(definitionName, propertyName, properties[propertyName]); err != nil {
			return err
		}
	}

	g.appendCfPropertyReferences(definitionName, propertyNames)

	return nil
}

// appendCfPropertyReferences generates Go code that references the code generated for the
// specified CloudFormation properties and appends the generated to code to the generator's Writer.
func (g *Generator) appendCfPropertyReferences(definitionName string, propertyNames []string) {
	attributesVariableName := CfDefinitionTfAttributesVariableName(definitionName)

	g.printf("\n")
	g.printf("// Property references for %s:\n", definitionName)
	g.printf("%s := make(map[string]schema.Attribute, %d)\n", attributesVariableName, len(propertyNames))
	for _, propertyName := range propertyNames {
		g.printf("%s[%q] = %s\n", attributesVariableName, CfPropertyTfAttributeName(propertyName), CfPropertyTfAttributeVariableName(definitionName, propertyName))
	}
	g.printf("\n")
}

// appendCfProperty generates Go code for a single CloudFormation property
// and appends the generated to code to the generator's Writer.
func (g *Generator) appendCfProperty(definitionName, propertyName string, property *cfschema.Property) error {
	attributeVariableName := CfPropertyTfAttributeVariableName(definitionName, propertyName)

	g.printf("\n")
	g.printf("// Definition: %s\n", definitionName)
	g.printf("// Property: %s\n", propertyName)
	if g.CfResource.PrimaryIdentifier.ContainsPath([]string{propertyName}) {
		g.printf("// PrimaryIdentifier: %t\n", true)
	}
	g.printf("// CloudFormation resource type schema:\n")
	g.printf("/*\n")
	g.printf("%v\n", property)
	g.printf("*/\n")
	g.printf("%s := schema.Attribute{}\n", attributeVariableName)

	if ref := property.Ref; ref != nil {
		g.printf("%s.Attributes = schema.SingleNestedAttributes(%s)\n", attributeVariableName, CfDefinitionTfAttributesVariableName(ref.Field()))
	} else {
		switch propertyType := property.Type.String(); propertyType {
		case cfschema.PropertyTypeBoolean:
			g.printf("%s.Type = types.BoolType\n", attributeVariableName)
		case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
			g.printf("%s.Type = types.NumberType\n", attributeVariableName)
		case cfschema.PropertyTypeString:
			g.printf("%s.Type = types.StringType\n", attributeVariableName)
		case cfschema.PropertyTypeArray:
			isSet := property.UniqueItems != nil && *property.UniqueItems

			if ref := property.Items.Ref; ref != nil {
				nestedAttributesVariableName := CfDefinitionTfAttributesVariableName(ref.Field())
				nestedAttributesOptionsVariableName := attributeVariableName + "Options"

				if isSet {
					g.printf("%s := schema.SetNestedAttributesOptions{}\n", nestedAttributesOptionsVariableName)
					if property.MinItems != nil {
						g.printf("%s.MinItems = %d\n", nestedAttributesOptionsVariableName, *property.MinItems)
					}
					if property.MaxItems != nil {
						g.printf("%s.MaxItems = %d\n", nestedAttributesOptionsVariableName, *property.MaxItems)
					}
					g.printf("%s.Attributes = schema.SetNestedAttributes(%s, %s)\n", attributeVariableName, nestedAttributesVariableName, nestedAttributesOptionsVariableName)
				} else {
					g.printf("%s := schema.ListNestedAttributesOptions{}\n", nestedAttributesOptionsVariableName)
					if property.MinItems != nil {
						g.printf("%s.MinItems = %d\n", nestedAttributesOptionsVariableName, *property.MinItems)
					}
					if property.MaxItems != nil {
						g.printf("%s.MaxItems = %d\n", nestedAttributesOptionsVariableName, *property.MaxItems)
					}
					g.printf("%s.Attributes = schema.ListNestedAttributes(%s, %s)\n", attributeVariableName, nestedAttributesVariableName, nestedAttributesOptionsVariableName)
				}
			} else {
				if isSet {
					return fmt.Errorf("%s/%s is of unsupported type: set of primitive", definitionName, propertyName)
				} else {
					switch itemType := property.Items.Type.String(); itemType {
					case cfschema.PropertyTypeBoolean:
						g.printf("%s.Type = types.ListType{ElemType:types.BoolType}\n", attributeVariableName)
					case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
						g.printf("%s.Type = types.ListType{ElemType:types.NumberType}\n", attributeVariableName)
					case cfschema.PropertyTypeString:
						g.printf("%s.Type = types.ListType{ElemType:types.StringType}\n", attributeVariableName)
					default:
						return fmt.Errorf("%s/%s is of unsupported type: list of %s", definitionName, propertyName, itemType)
					}
				}
			}
		case cfschema.PropertyTypeObject:
			if patternProperties := property.PatternProperties; len(patternProperties) > 0 {
				return fmt.Errorf("%s/%s is of unsupported type: key-value map", definitionName, propertyName)
			} else if len(property.Properties) > 0 {
				return fmt.Errorf("%s/%s has unsupported inline subproperties", definitionName, propertyName)
			}
			fallthrough
		default:
			return fmt.Errorf("%s/%s is of unsupported type: %s", definitionName, propertyName, propertyType)
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

	return nil
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
