package codegen

import (
	"fmt"
	"hash/fnv"
	"io"
	"sort"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/iancoleman/strcase"
)

// Features of the emitted code.
type Features int

const (
	HasUpdatableProperty Features = 1 << iota // At least one property can be updated.
)

type Emitter struct {
	CfResource *cfschema.Resource
	Writer     io.Writer
}

// EmitRootPropertyAttribute generates the Terraform Plugin SDK code for a CloudFormation root property
// and emits the generated code to the emitter's Writer. Code features are returned.
func (e *Emitter) EmitRootPropertyAttribute(name string) (Features, error) {
	property, ok := e.CfResource.Properties[name]

	if !ok || property == nil {
		return 0, fmt.Errorf("root property not found")
	}

	return e.EmitPropertySchema([]string{}, name, property)
}

// EmitPropertySchema generates the Terraform Plugin SDK code for a CloudFormation property
// and emits the generated code to the emitter's Writer. Code features are returned.
func (e *Emitter) EmitPropertySchema(pathPrefix []string, name string, property *cfschema.Property) (Features, error) {
	var features Features

	path := append(pathPrefix, name)

	if name != "" {
		e.printf("// Property: %s\n", name)
		if e.CfResource.PrimaryIdentifier.ContainsPath(path) {
			e.printf("// PrimaryIdentifier: %t\n", true)
		}
		e.printf("// CloudFormation resource type schema:\n")
		e.printf("/*\n")
		e.printf("%v\n", property)
		e.printf("*/\n")
		e.printf("%q: {\n", strcase.ToSnake(name))
	}

	if description := property.Description; description != nil {
		e.printf("Description: `%s`,\n", *description)
	}

	switch propertyType := property.Type.String(); propertyType {
	case cfschema.PropertyTypeBoolean:
		e.printf("Type: types.BoolType,\n")
	case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
		e.printf("Type: types.NumberType,\n")
	case cfschema.PropertyTypeString:
		e.printf("Type: types.StringType,\n")
	default:
		return 0, fmt.Errorf("%s is of unsupported type: %s", name, propertyType)
	}

	createOnly := e.CfResource.CreateOnlyProperties.ContainsPath(path)
	readOnly := e.CfResource.ReadOnlyProperties.ContainsPath(path)
	required := e.CfResource.IsRequired(name)
	writeOnly := e.CfResource.WriteOnlyProperties.ContainsPath(path)

	if required {
		e.printf("Required: true,\n")
	} else if !readOnly {
		e.printf("Optional: true,\n")
	}

	if readOnly && !required {
		e.printf("Computed: true,\n")
	}

	if createOnly {
		e.printf("// %s is a force-new attribute.\n", name)
	}

	if writeOnly {
		e.printf("// %s is a write-only attribute.\n", name)
	}

	if name != "" {
		e.printf("},\n")
	}

	return features, nil
}

// AppendCfDefinition generates the Terraform Plugin SDK code for a CloudFormation definition
// and appends the generated to code to the generator's Writer.
func (g *Emitter) AppendCfDefinition(definitionName string, properties map[string]*cfschema.Property) error {
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
func (g *Emitter) appendCfPropertyReferences(definitionName string, propertyNames []string) {
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
func (g *Emitter) appendCfProperty(definitionName, propertyName string, property *cfschema.Property) error {
	attributeVariableName := CfPropertyTfAttributeVariableName(definitionName, propertyName)
	path := []string{propertyName}

	g.printf("\n")
	g.printf("// Definition: %s\n", definitionName)
	g.printf("// Property: %s\n", propertyName)
	if g.CfResource.PrimaryIdentifier.ContainsPath(path) {
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

			// TODO
			// TODO Worry about InsertionOrder and Sets.
			// TODO

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
				n := 0
				for pattern, property := range patternProperties {
					g.printf("// Pattern: %q\n", pattern)
					if n == 0 {
						if ref := property.Ref; ref != nil {
							return fmt.Errorf("%s/%s is of unsupported type: key-value map of complex type (%s)", definitionName, propertyName, ref.Field())
						}

						switch propertyType := property.Type.String(); propertyType {
						case cfschema.PropertyTypeBoolean:
							g.printf("%s.Type = types.MapType{ElemType:types.BoolType}\n", attributeVariableName)
						case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
							g.printf("%s.Type = types.MapType{ElemType:types.NumberType}\n", attributeVariableName)
						case cfschema.PropertyTypeString:
							g.printf("%s.Type = types.MapType{ElemType:types.StringType}\n", attributeVariableName)
						default:
							return fmt.Errorf("%s/%s is of unsupported type: key-value map of %s", definitionName, propertyName, propertyType)
						}
					} else {
						g.printf("// Ignored.\n")
					}
					n++
				}

				break
			} else if len(property.Properties) > 0 {
				return fmt.Errorf("%s/%s has unsupported inline subproperties", definitionName, propertyName)
			}
			fallthrough
		default:
			return fmt.Errorf("%s/%s is of unsupported type: %s", definitionName, propertyName, propertyType)
		}
	}

	createOnly := g.CfResource.CreateOnlyProperties.ContainsPath(path)
	readOnly := g.CfResource.ReadOnlyProperties.ContainsPath(path)
	required := g.CfResource.IsRequired(propertyName)
	writeOnly := g.CfResource.WriteOnlyProperties.ContainsPath(path)

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

	if createOnly {
		g.printf("// %s.ForceNew = true\n", attributeVariableName)
	}

	if writeOnly {
		g.printf("// %s is a write-only attribute.\n", propertyName)
	}

	return nil
}

// printf writes a formatted string to the underlying writer.
func (g *Emitter) printf(format string, a ...interface{}) (int, error) {
	return io.WriteString(g.Writer, fmt.Sprintf(format, a...))
}

// CfDefinitionTfAttributesVariableName returns a CloudFormation definition's Terraform map[string]Attribute variable name.
func CfDefinitionTfAttributesVariableName(definitionName string) string {
	h := fnv.New32a()
	h.Write([]byte(definitionName))
	return fmt.Sprintf("attrs%d", h.Sum32())
}

// CfPropertyTfAttributeVariableName returns a CloudFormation property's Terraform Attribute variable name.
func CfPropertyTfAttributeVariableName(definitionName, propertyName string) string {
	h := fnv.New32a()
	h.Write([]byte(definitionName))
	h.Write([]byte(propertyName))
	return fmt.Sprintf("attr%d", h.Sum32())
}

// TfPropertyAttributeName returns a CloudFormation property's Terraform Attribute name.
func CfPropertyTfAttributeName(propertyName string) string {
	return strcase.ToSnake(propertyName)
}
