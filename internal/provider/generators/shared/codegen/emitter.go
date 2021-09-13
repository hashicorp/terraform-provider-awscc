package codegen

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/mitchellh/cli"
)

// Features of the emitted code.
type Features int

const (
	HasUpdatableProperty    Features = 1 << iota // At least one property can be updated.
	UsesInternalTypes                            // Uses a type from the internal/types package.
	HasRequiredRootProperty                      // At least one root property is required.
	UsesValidation                               // Uses a type from the internal/validate package.
	HasIDRootProperty                            // Has a root property named "id"
)

var (
	tfMetaArguments = []string{
		"count",
		"depends_on",
		"for_each",
		"lifecycle",
		"provider",
	}
)

type Emitter struct {
	CfResource *cfschema.Resource
	Ui         cli.Ui
	Writer     io.Writer
}

type parent struct {
	path []string
	reqd interface {
		IsRequired(name string) bool
	}
}

// EmitResourceSchemaRequiredAttributesValidator generates any resource schema-level required Attributes validators.
func (e Emitter) EmitResourceSchemaRequiredAttributesValidator() {
	e.printf(resourceRequiredAttributesValidator(e.CfResource))
}

// EmitRootPropertiesSchema generates the Terraform Plugin SDK code for a CloudFormation root schema
// and emits the generated code to the emitter's Writer. Code features are returned.
// The root schema is the map of root property names to Attributes.
// To generate all code features as "Computed", e.g. to be used in a singular data source, set the computedOnly argument to true.
func (e Emitter) EmitRootPropertiesSchema(attributeNameMap map[string]string, computedOnly bool) (Features, error) {
	var features Features

	cfResource := e.CfResource
	features, err := e.emitSchema(attributeNameMap, parent{reqd: cfResource}, cfResource.Properties, computedOnly)

	if err != nil {
		return 0, err
	}

	for name, property := range cfResource.Properties {
		if naming.CloudFormationPropertyToTerraformAttribute(name) == "id" {
			// Ensure that any schema-declared top-level ID property is of type String and is the primary identifier.
			if propertyType := property.Type.String(); propertyType != cfschema.PropertyTypeString {
				return 0, fmt.Errorf("top-level property %s has type: %s", name, propertyType)
			}

			if !cfResource.PrimaryIdentifier.ContainsPath([]string{name}) {
				return 0, fmt.Errorf("top-level property %s is not a primary identifier", name)
			}

			features |= HasIDRootProperty
		}

		for _, tfMetaArgument := range tfMetaArguments {
			if naming.CloudFormationPropertyToTerraformAttribute(name) == tfMetaArgument {
				return 0, fmt.Errorf("top-level property %s conflicts with Terraform meta-argument: %s", name, tfMetaArgument)
			}
		}

		if cfResource.IsRequired(name) {
			features |= HasRequiredRootProperty
		}
	}

	return features, nil
}

// emitAttribute generates the Terraform Plugin SDK code for a CloudFormation property's Attributes
// and emits the generated code to the emitter's Writer. Code features are returned.
// To generate all code features as "Computed", e.g. to be used in a singular data source, set the computedOnly argument to true.
func (e Emitter) emitAttribute(attributeNameMap map[string]string, path []string, name string, property *cfschema.Property, required, computedOnly bool) (Features, error) {
	var features Features
	var validators []string

	e.printf("{\n")
	e.printf("// Property: %s\n", name)

	// Only dump top-level property schemas as nested properties have been expanded here.
	if len(path) == 1 {
		e.printf("// CloudFormation resource type schema:\n")
		// Comment out each line.
		e.printf("%s\n", regexp.MustCompile(`(?m)^`).ReplaceAllString(fmt.Sprintf("%v", property), "// "))
	}

	if description := property.Description; description != nil {
		e.printf("Description:%q,\n", *description)
	}

	switch propertyType := property.Type.String(); propertyType {
	//
	// Primitive types.
	//
	case cfschema.PropertyTypeBoolean:
		e.printf("Type:types.BoolType,\n")

	case cfschema.PropertyTypeInteger:
		e.printf("Type:types.NumberType,\n")

		if property.Minimum == nil && property.Maximum != nil {
			return 0, fmt.Errorf("%s has Maximum but no Minimum", strings.Join(path, "/"))
		}

		if property.Minimum != nil && property.Maximum == nil {
			validators = append(validators, fmt.Sprintf("validate.IntAtLeast(%d)", *property.Minimum))
		}
		if property.Minimum != nil && property.Maximum != nil {
			validators = append(validators, fmt.Sprintf("validate.IntBetween(%d,%d)", *property.Minimum, *property.Maximum))
		}

		if property.Format != nil {
			if format := *property.Format; format != "int64" {
				return 0, fmt.Errorf("%s has unsupported format :%s", strings.Join(path, "/"), format)
			}
		}

		if len(property.Enum) > 0 {
			sb := strings.Builder{}
			sb.WriteString("validate.IntInSlice([]int{\n")
			for _, enum := range property.Enum {
				sb.WriteString(fmt.Sprintf("%d", int(enum.(float64))))
				sb.WriteString(",\n")
			}
			sb.WriteString("})")
			validators = append(validators, sb.String())
		}

	case cfschema.PropertyTypeNumber:
		e.printf("Type:types.NumberType,\n")

		if property.Minimum == nil && property.Maximum != nil {
			return 0, fmt.Errorf("%s has Maximum but no Minimum", strings.Join(path, "/"))
		}

		if property.Minimum != nil && property.Maximum == nil {
			validators = append(validators, fmt.Sprintf("validate.FloatAtLeast(%f)", float64(*property.Minimum)))
		}
		if property.Minimum != nil && property.Maximum != nil {
			validators = append(validators, fmt.Sprintf("validate.FloatBetween(%f,%f)", float64(*property.Minimum), float64(*property.Maximum)))
		}

		if property.Format != nil {
			if format := *property.Format; format != "double" {
				return 0, fmt.Errorf("%s has unsupported format :%s", strings.Join(path, "/"), format)
			}
		}

		if len(property.Enum) > 0 {
			return 0, fmt.Errorf("%s has enumerated values", strings.Join(path, "/"))
		}

	case cfschema.PropertyTypeString:
		e.printf("Type:types.StringType,\n")

		if property.MinLength != nil && property.MaxLength == nil {
			validators = append(validators, fmt.Sprintf("validate.StringLenAtLeast(%d)", *property.MinLength))
		}
		if property.MaxLength != nil {
			minLength := 0
			if property.MinLength != nil {
				minLength = *property.MinLength
			}
			validators = append(validators, fmt.Sprintf("validate.StringLenBetween(%d,%d)", minLength, *property.MaxLength))
		}

		if property.Format != nil {
			switch format := *property.Format; format {
			case "date-time":
				validators = append(validators, "validate.IsRFC3339Time()")
			case "string":
			default:
				return 0, fmt.Errorf("%s has unsupported format :%s", strings.Join(path, "/"), format)
			}
		}

		if len(property.Enum) > 0 {
			sb := strings.Builder{}
			sb.WriteString("validate.StringInSlice([]string{\n")
			for _, enum := range property.Enum {
				sb.WriteString("\"")
				sb.WriteString(enum.(string))
				sb.WriteString("\",\n")
			}
			sb.WriteString("})")
			validators = append(validators, sb.String())
		}

	//
	// Complex types.
	//
	case cfschema.PropertyTypeArray:
		arrayType := aggregateType(property)

		if arrayType == aggregateSet {
			//
			// Set.
			//
			switch itemType := property.Items.Type.String(); itemType {
			// Sets of primitive types use provider-local Set type until tfsdk support is available.
			case cfschema.PropertyTypeBoolean:
				features |= UsesInternalTypes
				e.printf("Type:providertypes.SetType{ElemType:types.BoolType},\n")

			case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
				features |= UsesInternalTypes
				e.printf("Type:providertypes.SetType{ElemType:types.NumberType},\n")

			case cfschema.PropertyTypeString:
				features |= UsesInternalTypes
				e.printf("Type:providertypes.SetType{ElemType:types.StringType},\n")

			case cfschema.PropertyTypeObject:
				if len(property.Items.PatternProperties) > 0 {
					return 0, unsupportedTypeError(path, "set of key-value map")
				}

				if len(property.Items.Properties) == 0 {
					return 0, unsupportedTypeError(path, "set of undefined schema")
				}

				features |= UsesInternalTypes
				e.printf("Attributes:providertypes.SetNestedAttributes(\n")

				f, err := e.emitSchema(
					attributeNameMap,
					parent{
						path: path,
						reqd: property.Items,
					},
					property.Items.Properties,
					computedOnly)

				if err != nil {
					return 0, err
				}

				features |= f

				e.printf(",\n")
				e.printf("providertypes.SetNestedAttributesOptions{\n")

				if !computedOnly {
					if property.MinItems != nil {
						e.printf("MinItems:%d,\n", *property.MinItems)
					}
					if property.MaxItems != nil {
						e.printf("MaxItems:%d,\n", *property.MaxItems)
					}
				}

				e.printf("},\n")
				e.printf("),\n")

				if validator := propertyRequiredAttributesValidator(property.Items); validator != "" {
					validators = append(validators, validator)
				}

			default:
				return 0, unsupportedTypeError(path, fmt.Sprintf("set of %s", itemType))
			}
		} else {
			//
			// List.
			//
			var elementType string

			switch itemType := property.Items.Type.String(); itemType {
			case cfschema.PropertyTypeBoolean:
				elementType = "types.BoolType"

			case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
				elementType = "types.NumberType"

			case cfschema.PropertyTypeString:
				elementType = "types.StringType"

			case cfschema.PropertyTypeObject:
				if len(property.Items.PatternProperties) > 0 {
					return 0, unsupportedTypeError(path, "list of key-value map")
				}

				if len(property.Items.Properties) == 0 {
					return 0, unsupportedTypeError(path, "list of undefined schema")
				}

				e.printf("Attributes:tfsdk.ListNestedAttributes(\n")

				f, err := e.emitSchema(
					attributeNameMap,
					parent{
						path: path,
						reqd: property.Items,
					},
					property.Items.Properties,
					computedOnly)

				if err != nil {
					return 0, err
				}

				features |= f

				e.printf(",\n")
				e.printf("tfsdk.ListNestedAttributesOptions{\n")

				if !computedOnly {
					if property.MinItems != nil {
						e.printf("MinItems:%d,\n", *property.MinItems)
					}
					if property.MaxItems != nil {
						e.printf("MaxItems:%d,\n", *property.MaxItems)
					}
				}

				e.printf("},\n")
				e.printf("),\n")

				if arrayType == aggregateOrderedSet {
					validators = append(validators, "validate.UniqueItems()")
				}

				if validator := propertyRequiredAttributesValidator(property.Items); validator != "" {
					validators = append(validators, validator)
				}

			default:
				return 0, unsupportedTypeError(path, fmt.Sprintf("list of %s", itemType))
			}

			if elementType != "" {
				e.printf("Type:types.ListType{ElemType:%s},\n", elementType)

				if property.MinItems != nil && property.MaxItems == nil {
					validators = append(validators, fmt.Sprintf("validate.ArrayLenAtLeast(%d)", *property.MinItems))
				}
				if property.MaxItems != nil {
					minItems := 0
					if property.MinItems != nil {
						minItems = *property.MinItems
					}
					validators = append(validators, fmt.Sprintf("validate.ArrayLenBetween(%d,%d)", minItems, *property.MaxItems))
				}

				if arrayType == aggregateOrderedSet {
					validators = append(validators, "validate.UniqueItems()")
				}
			}
		}

	case cfschema.PropertyTypeObject:
		if patternProperties := property.PatternProperties; len(patternProperties) > 0 {
			//
			// Map.
			//
			if len(property.Properties) > 0 {
				return 0, fmt.Errorf("%s has both Properties and PatternProperties", strings.Join(path, "/"))
			}

			// Sort the patterns to reduce diffs.
			patterns := make([]string, 0)
			for pattern := range patternProperties {
				patterns = append(patterns, pattern)
			}
			sort.Strings(patterns)

			// Ignore all but the first pattern.
			pattern := patterns[0]
			patternProperty := patternProperties[pattern]

			e.printf("// Pattern: %q\n", pattern)
			switch propertyType := patternProperty.Type.String(); propertyType {
			//
			// Primitive types.
			//
			case cfschema.PropertyTypeBoolean:
				e.printf("Type:types.MapType{ElemType:types.BoolType},\n")

			case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
				e.printf("Type:types.MapType{ElemType:types.NumberType},\n")

			case cfschema.PropertyTypeString:
				e.printf("Type:types.MapType{ElemType:types.StringType},\n")

			//
			// Complex types.
			//
			case cfschema.PropertyTypeArray:
				if aggregateType(patternProperty) == aggregateSet {
					switch itemType := patternProperty.Items.Type.String(); itemType {
					// Sets of primitive types use provider-local Set type until tfsdk support is available.
					case cfschema.PropertyTypeBoolean:
						features |= UsesInternalTypes
						e.printf("Type: types.MapType{ElemType:providertypes.SetType{ElemType:types.BoolType}},\n")

					case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
						features |= UsesInternalTypes
						e.printf("Type: types.MapType{ElemType:providertypes.SetType{ElemType:types.NumberType}},\n")

					case cfschema.PropertyTypeString:
						features |= UsesInternalTypes
						e.printf("Type: types.MapType{ElemType:providertypes.SetType{ElemType:types.StringType}},\n")

					default:
						return 0, unsupportedTypeError(path, fmt.Sprintf("key-value map of set of %s", itemType))
					}
				} else {
					switch itemType := patternProperty.Items.Type.String(); itemType {
					case cfschema.PropertyTypeBoolean:
						e.printf("Type:types.MapType{ElemType:types.ListType{ElemType:types.BoolType}},\n")

					case cfschema.PropertyTypeInteger, cfschema.PropertyTypeNumber:
						e.printf("Type:types.MapType{ElemType:types.ListType{ElemType:types.NumberType}},\n")

					case cfschema.PropertyTypeString:
						e.printf("Type:types.MapType{ElemType:types.ListType{ElemType:types.StringType}},\n")

					default:
						return 0, unsupportedTypeError(path, fmt.Sprintf("key-value map of list of %s", itemType))
					}
				}

			case cfschema.PropertyTypeObject:
				if len(patternProperty.PatternProperties) > 0 {
					return 0, unsupportedTypeError(path, "key-value map of key-value map")
				}

				if len(patternProperty.Properties) == 0 {
					return 0, unsupportedTypeError(path, "key-value map of undefined schema")
				}

				e.printf("Attributes:tfsdk.MapNestedAttributes(\n")

				f, err := e.emitSchema(
					attributeNameMap,
					parent{
						path: path,
						reqd: property.Items,
					},
					patternProperty.Properties,
					computedOnly)

				if err != nil {
					return 0, err
				}

				features |= f

				e.printf(",\n")
				e.printf("tfsdk.MapNestedAttributesOptions{\n")

				if !computedOnly {
					if patternProperty.MinItems != nil {
						e.printf("MinItems:%d,\n", *patternProperty.MinItems)
					}
					if patternProperty.MaxItems != nil {
						e.printf("MaxItems:%d,\n", *patternProperty.MaxItems)
					}
				}

				e.printf("},\n")
				e.printf("),\n")

			default:
				return 0, unsupportedTypeError(path, fmt.Sprintf("key-value map of %s", propertyType))
			}

			for _, pattern := range patterns[1:] {
				e.printf("// Pattern %q ignored.\n", pattern)
			}

			break
		}

		//
		// Object.
		//
		if len(property.Properties) == 0 {
			// Schemaless object => key-value map of string.
			e.printf("Type:types.MapType{ElemType:types.StringType},\n")

			break
		}

		e.printf("Attributes:tfsdk.SingleNestedAttributes(\n")
		f, err := e.emitSchema(
			attributeNameMap,
			parent{
				path: path,
				reqd: property,
			},
			property.Properties,
			computedOnly)

		if err != nil {
			return 0, err
		}

		features |= f

		e.printf(",\n")
		e.printf("),\n")

		if validator := propertyRequiredAttributesValidator(property); validator != "" {
			validators = append(validators, validator)
		}

	default:
		return 0, unsupportedTypeError(path, propertyType)
	}

	// Return early as attribute validations are not required
	// and additional configurations are not supported when an attribute is Computed-only.
	if computedOnly {
		e.printf("Computed:true,\n")
		e.printf("}")

		return features, nil
	}

	createOnly := e.CfResource.CreateOnlyProperties.ContainsPath(path)
	readOnly := e.CfResource.ReadOnlyProperties.ContainsPath(path)
	writeOnly := e.CfResource.WriteOnlyProperties.ContainsPath(path)

	if readOnly && required {
		e.warnf("%s is ReadOnly and Required", strings.Join(path, "/"))
	}
	if readOnly && writeOnly {
		e.warnf("%s is ReadOnly and WriteOnly", strings.Join(path, "/"))
	}

	var optional, computed bool

	if required {
		e.printf("Required:true,\n")
	} else if !readOnly {
		optional = true
		e.printf("Optional:true,\n")
	}

	if (readOnly || createOnly) && !required {
		computed = true
		e.printf("Computed:true,\n")
	}

	// Don't emit validators for Computed-only attributes.
	if !computed || optional {
		if len(validators) > 0 {
			features |= UsesValidation
			e.printf("Validators:[]tfsdk.AttributeValidator{\n")
			for _, validator := range validators {
				e.printf("%s,\n", validator)
			}
			e.printf("},\n")
		}
	}

	if createOnly {
		e.printf("PlanModifiers:[]tfsdk.AttributePlanModifier{\n")
		e.printf("tfsdk.RequiresReplace(),// %s is a force-new property.\n", name)
		e.printf("},\n")
	}

	if writeOnly {
		e.printf("// %s is a write-only property.\n", name)
	}

	if !createOnly && !readOnly {
		features |= HasUpdatableProperty
	}

	e.printf("}")

	return features, nil
}

// emitSchema generates the Terraform Plugin SDK code for a CloudFormation property's schema.
// and emits the generated code to the emitter's Writer. Code features are returned.
// A schema is a map of property names to Attributes.
// Property names are sorted prior to code generation to reduce diffs.
// To generate all code features as "Computed", e.g. to be used in a singular data source, set the computedOnly argument to true.
func (e Emitter) emitSchema(attributeNameMap map[string]string, parent parent, properties map[string]*cfschema.Property, computedOnly bool) (Features, error) {
	names := make([]string, 0)
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)

	var features Features

	e.printf("map[string]tfsdk.Attribute{\n")
	for _, name := range names {
		tfAttributeName := naming.CloudFormationPropertyToTerraformAttribute(name)
		cfPropertyName, ok := attributeNameMap[tfAttributeName]
		if ok {
			if cfPropertyName != name {
				return 0, fmt.Errorf("%s overwrites %s for Terraform attribute %s", name, cfPropertyName, tfAttributeName)
			}
		} else {
			attributeNameMap[tfAttributeName] = name
		}

		e.printf("%q:", tfAttributeName)

		f, err := e.emitAttribute(
			attributeNameMap,
			append(parent.path, name),
			name,
			properties[name],
			parent.reqd.IsRequired(name),
			computedOnly)

		if err != nil {
			return 0, err
		}

		features |= f

		e.printf(",\n")
	}
	e.printf("}")

	return features, nil
}

// printf emits a formatted string to the underlying writer.
func (e Emitter) printf(format string, a ...interface{}) (int, error) {
	return wprintf(e.Writer, format, a...)
}

// warnf emits a formatted warning message to the UI.
func (e Emitter) warnf(format string, a ...interface{}) {
	e.Ui.Warn(fmt.Sprintf(format, a...))
}

// wprintf writes a formatted string to a Writer.
func wprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	return io.WriteString(w, fmt.Sprintf(format, a...))
}

type aggregate int

const (
	aggregateNone aggregate = iota
	aggregateList
	aggregateSet
	aggregateMultiset
	aggregateOrderedSet
)

// aggregate returns the type of a Property.
func aggregateType(property *cfschema.Property) aggregate {
	if property.Type.String() != cfschema.PropertyTypeArray {
		return aggregateNone
	}

	// https://github.com/aws-cloudformation/cloudformation-resource-schema#insertionorder
	insertionOrder := true
	if property.InsertionOrder != nil {
		insertionOrder = *property.InsertionOrder
	}
	uniqueItems := false
	if property.UniqueItems != nil {
		uniqueItems = *property.UniqueItems
	}

	if uniqueItems && !insertionOrder {
		return aggregateSet
	}

	if uniqueItems && insertionOrder {
		return aggregateOrderedSet
	}

	if !uniqueItems && !insertionOrder {
		return aggregateMultiset
	}

	return aggregateList
}

func unsupportedTypeError(path []string, typ string) error {
	return fmt.Errorf("%s is of unsupported type: %s", strings.Join(path, "/"), typ)
}

func addPropertyRequiredAttributes(writer io.Writer, p *cfschema.PropertySubschema) int {
	var nRequired int

	if len(p.AllOf) > 0 {
		var n int
		w := &strings.Builder{}

		wprintf(w, "validate.AllOfRequired(\n")
		for _, a := range p.AllOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		wprintf(w, "),\n")

		if n > 0 {
			wprintf(writer, w.String())
		}

		nRequired += n
	}
	if len(p.AnyOf) > 0 {
		var n int
		w := &strings.Builder{}

		wprintf(w, "validate.AnyOfRequired(\n")
		for _, a := range p.AnyOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		wprintf(w, "),\n")

		if n > 0 {
			wprintf(writer, w.String())
		}

		nRequired += n
	}
	if len(p.OneOf) > 0 {
		var n int
		w := &strings.Builder{}

		wprintf(w, "validate.OneOfRequired(\n")
		for _, a := range p.OneOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		wprintf(w, "),\n")

		if n > 0 {
			wprintf(writer, w.String())
		}

		nRequired += n
	}
	if len(p.Required) > 0 {
		wprintf(writer, "validate.Required(\n")
		for _, r := range p.Required {
			wprintf(writer, "%q,\n", naming.CloudFormationPropertyToTerraformAttribute(r))
			nRequired++
		}
		wprintf(writer, "),\n")
	}

	return nRequired
}

func addSchemaCompositionRequiredAttributes(writer io.Writer, r schemaComposition) int {
	var nRequired int

	if allOf := r.All(); len(allOf) > 0 {
		var n int
		w := &strings.Builder{}

		wprintf(w, "validate.AllOfRequired(\n")
		for _, a := range allOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		wprintf(w, "),\n")

		if n > 0 {
			wprintf(writer, w.String())
		}

		nRequired += n
	}
	if anyOf := r.Any(); len(anyOf) > 0 {
		var n int
		w := &strings.Builder{}

		wprintf(w, "validate.AnyOfRequired(\n")
		for _, a := range anyOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		wprintf(w, "),\n")

		if n > 0 {
			wprintf(writer, w.String())
		}

		nRequired += n
	}
	if oneOf := r.One(); len(oneOf) > 0 {
		var n int
		w := &strings.Builder{}

		wprintf(w, "validate.OneOfRequired(\n")
		for _, a := range oneOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		wprintf(w, "),\n")

		if n > 0 {
			wprintf(writer, w.String())
		}

		nRequired += n
	}

	return nRequired
}

func propertyRequiredAttributesValidator(p *cfschema.Property) string {
	if p == nil {
		return ""
	}

	writer := &strings.Builder{}

	wprintf(writer, "validate.RequiredAttributes(\n")
	nRequired := addSchemaCompositionRequiredAttributes(writer, property(*p))
	wprintf(writer, ")")

	if nRequired == 0 {
		return ""
	}

	return writer.String()
}

func resourceRequiredAttributesValidator(r *cfschema.Resource) string {
	if r == nil {
		return ""
	}

	writer := &strings.Builder{}
	nRequired := addSchemaCompositionRequiredAttributes(writer, resource(*r))

	if nRequired == 0 {
		return ""
	}

	return writer.String()
}

// The schemaComposition interface can be implemented by Property and Resource.
type schemaComposition interface {
	All() []*cfschema.PropertySubschema
	Any() []*cfschema.PropertySubschema
	One() []*cfschema.PropertySubschema
}

type property cfschema.Property

func (p property) All() []*cfschema.PropertySubschema {
	return p.AllOf
}

func (p property) Any() []*cfschema.PropertySubschema {
	return p.AnyOf
}

func (p property) One() []*cfschema.PropertySubschema {
	return p.OneOf
}

type resource cfschema.Resource

func (r resource) All() []*cfschema.PropertySubschema {
	return r.AllOf
}

func (r resource) Any() []*cfschema.PropertySubschema {
	return r.AnyOf
}

func (r resource) One() []*cfschema.PropertySubschema {
	return r.OneOf
}
