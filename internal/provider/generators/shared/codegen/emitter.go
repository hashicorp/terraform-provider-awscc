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
	HasRequiredRootProperty                      // At least one root property is required.
	UsesFrameworkAttr                            // Uses a type from the terraform-plugin-framework/attr package.
	UsesRegexpInValidation                       // Uses a type from the Go standard regexp package for attribute validation.
	UsesValidation                               // Uses a type or function from the internal/validate package.
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
func (e Emitter) EmitResourceSchemaRequiredAttributesValidator() error {
	v, err := resourceRequiredAttributesValidator(e.CfResource)

	if err != nil {
		return err
	}

	e.printf(v)

	return nil
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
	var planModifiers []string

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
		e.printf("Type:types.Int64Type,\n")

		if f, v, err := integerValidators(path, property); err != nil {
			return 0, err
		} else if len(v) > 0 {
			features |= f
			validators = append(validators, v...)
		}

	case cfschema.PropertyTypeNumber:
		e.printf("Type:types.Float64Type,\n")

		if f, v, err := numberValidators(path, property); err != nil {
			return 0, err
		} else if len(v) > 0 {
			features |= f
			validators = append(validators, v...)
		}

	case cfschema.PropertyTypeString:
		e.printf("Type:types.StringType,\n")

		if f, v, err := stringValidators(path, property); err != nil {
			return 0, err
		} else if len(v) > 0 {
			features |= f
			validators = append(validators, v...)
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
			var elementType string
			var validatorsGenerator primitiveValidatorsGenerator

			switch itemType := property.Items.Type.String(); itemType {
			case cfschema.PropertyTypeBoolean:
				elementType = "types.BoolType"

			case cfschema.PropertyTypeInteger:
				elementType = "types.Int64Type"
				validatorsGenerator = integerValidators

			case cfschema.PropertyTypeNumber:
				elementType = "types.Float64Type"
				validatorsGenerator = numberValidators

			case cfschema.PropertyTypeString:
				elementType = "types.StringType"
				validatorsGenerator = stringValidators

			case cfschema.PropertyTypeObject:
				if len(property.Items.PatternProperties) > 0 {
					return 0, unsupportedTypeError(path, "set of key-value map")
				}

				if len(property.Items.Properties) == 0 {
					return 0, unsupportedTypeError(path, "set of undefined schema")
				}

				e.printf("Attributes:tfsdk.SetNestedAttributes(\n")

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
				e.printf("tfsdk.SetNestedAttributesOptions{},\n")

				e.printf("),\n")

				if v, err := arrayLengthValidator(path, property); err != nil {
					return 0, err
				} else if v != "" {
					validators = append(validators, v)
				}

				if validator, err := propertyRequiredAttributesValidator(property.Items); err != nil {
					return 0, err
				} else if validator != "" {
					validators = append(validators, validator)
				}

			default:
				return 0, unsupportedTypeError(path, fmt.Sprintf("set of %s", itemType))
			}

			if elementType != "" {
				e.printf("Type:types.SetType{ElemType:%s},\n", elementType)

				if v, err := arrayLengthValidator(path, property); err != nil {
					return 0, err
				} else if v != "" {
					validators = append(validators, v)
				}

				if validatorsGenerator != nil {
					if f, v, err := validatorsGenerator(path, property.Items); err != nil {
						return 0, err
					} else if len(v) > 0 {
						features |= f
						for _, v := range v {
							validators = append(validators, fmt.Sprintf("validate.ArrayForEach(%s)", v))
						}
					}
				}
			}
		} else {
			//
			// List.
			//
			var elementType string
			var validatorsGenerator primitiveValidatorsGenerator

			switch itemType := property.Items.Type.String(); itemType {
			case cfschema.PropertyTypeBoolean:
				elementType = "types.BoolType"

			case cfschema.PropertyTypeInteger:
				elementType = "types.Int64Type"
				validatorsGenerator = integerValidators

			case cfschema.PropertyTypeNumber:
				elementType = "types.Float64Type"
				validatorsGenerator = numberValidators

			case cfschema.PropertyTypeString:
				elementType = "types.StringType"
				validatorsGenerator = stringValidators

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
				e.printf("tfsdk.ListNestedAttributesOptions{},\n")

				e.printf("),\n")

				if v, err := arrayLengthValidator(path, property); err != nil {
					return 0, err
				} else if v != "" {
					validators = append(validators, v)
				}

				switch arrayType {
				case aggregateOrderedSet:
					validators = append(validators, "validate.UniqueItems()")
				case aggregateMultiset:
					planModifiers = append(planModifiers, "Multiset()")
				}

				if validator, err := propertyRequiredAttributesValidator(property.Items); err != nil {
					return 0, err
				} else if validator != "" {
					validators = append(validators, validator)
				}

			default:
				return 0, unsupportedTypeError(path, fmt.Sprintf("list of %s", itemType))
			}

			if elementType != "" {
				e.printf("Type:types.ListType{ElemType:%s},\n", elementType)

				if v, err := arrayLengthValidator(path, property); err != nil {
					return 0, err
				} else if v != "" {
					validators = append(validators, v)
				}

				switch arrayType {
				case aggregateOrderedSet:
					validators = append(validators, "validate.UniqueItems()")
				case aggregateMultiset:
					planModifiers = append(planModifiers, "Multiset()")
				}

				if validatorsGenerator != nil {
					if f, v, err := validatorsGenerator(path, property.Items); err != nil {
						return 0, err
					} else if len(v) > 0 {
						features |= f
						for _, v := range v {
							validators = append(validators, fmt.Sprintf("validate.ArrayForEach(%s)", v))
						}
					}
				}
			}
		}

	case "":
		//
		// If the property has no specified type but has properties then assume it's an object.
		//
		if len(property.PatternProperties) > 0 || len(property.Properties) == 0 {
			return 0, unsupportedTypeError(path, propertyType)
		}
		fallthrough

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

			case cfschema.PropertyTypeInteger:
				e.printf("Type:types.MapType{ElemType:types.Int64Type},\n")

			case cfschema.PropertyTypeNumber:
				e.printf("Type:types.MapType{ElemType:types.Float64Type},\n")

			case cfschema.PropertyTypeString:
				e.printf("Type:types.MapType{ElemType:types.StringType},\n")

			//
			// Complex types.
			//
			case cfschema.PropertyTypeArray:
				if aggregateType(patternProperty) == aggregateSet {
					switch itemType := patternProperty.Items.Type.String(); itemType {
					case cfschema.PropertyTypeBoolean:
						e.printf("Type: types.MapType{ElemType:types.SetType{ElemType:types.BoolType}},\n")

					case cfschema.PropertyTypeInteger:
						e.printf("Type: types.MapType{ElemType:types.SetType{ElemType:types.Int64Type}},\n")

					case cfschema.PropertyTypeNumber:
						e.printf("Type: types.MapType{ElemType:types.SetType{ElemType:types.Float64Type}},\n")

					case cfschema.PropertyTypeString:
						e.printf("Type: types.MapType{ElemType:types.SetType{ElemType:types.StringType}},\n")

					default:
						return 0, unsupportedTypeError(path, fmt.Sprintf("key-value map of set of %s", itemType))
					}
				} else {
					switch itemType := patternProperty.Items.Type.String(); itemType {
					case cfschema.PropertyTypeBoolean:
						e.printf("Type:types.MapType{ElemType:types.ListType{ElemType:types.BoolType}},\n")

					case cfschema.PropertyTypeInteger:
						e.printf("Type:types.MapType{ElemType:types.ListType{ElemType:types.Int64Type}},\n")

					case cfschema.PropertyTypeNumber:
						e.printf("Type:types.MapType{ElemType:types.ListType{ElemType:types.Float64Type}},\n")

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
				e.printf("tfsdk.MapNestedAttributesOptions{},\n")

				if !computedOnly {
					if patternProperty.MinItems != nil {
						return 0, fmt.Errorf("%s has unsupported MinItems", strings.Join(path, "/"))
					}
					if patternProperty.MaxItems != nil {
						return 0, fmt.Errorf("%s has unsupported MaxItems", strings.Join(path, "/"))
					}
				}

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
			e.warnf("%s is of type %s but has no schema", strings.Join(path, "/"), propertyType)
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

		if validator, err := propertyRequiredAttributesValidator(property); err != nil {
			return 0, err
		} else if validator != "" {
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

	// Handle any default value.
	var hasDefaultValue bool

	if f, planModifier, err := defaultValueAttributePlanModifier(path, property); err != nil {
		return 0, err
	} else if planModifier != "" {
		if required {
			e.warnf("%s is Required and has a default value. Emitting as Computed,Optional", strings.Join(path, "/"))
		}

		hasDefaultValue = true
		features |= f
		planModifiers = append(planModifiers, planModifier)
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

	if required && hasDefaultValue {
		required = false
		optional = true
	}

	if !required && !readOnly {
		optional = true
	}

	if (readOnly || createOnly) && !required {
		computed = true
	}

	if hasDefaultValue && !computed {
		computed = true
	}

	if required {
		e.printf("Required:true,\n")
	}
	if optional {
		e.printf("Optional:true,\n")
	}
	if computed {
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
	} else {
		features &^= UsesRegexpInValidation
	}

	if computed {
		// Computed.
		planModifiers = append(planModifiers, "tfsdk.UseStateForUnknown()")
	}

	if createOnly {
		// ForceNew.
		planModifiers = append(planModifiers, "tfsdk.RequiresReplace()")
	}

	if len(planModifiers) > 0 {
		e.printf("PlanModifiers:[]tfsdk.AttributePlanModifier{\n")
		for _, planModifier := range planModifiers {
			e.printf("%s,\n", planModifier)
		}
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
	return fprintf(e.Writer, format, a...)
}

// warnf emits a formatted warning message to the UI.
func (e Emitter) warnf(format string, a ...interface{}) {
	e.Ui.Warn(fmt.Sprintf(format, a...))
}

// fprintf writes a formatted string to a Writer.
func fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
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

// arrayLengthValidator returns any array length AttributeValidator for the specified Property.
func arrayLengthValidator(path []string, property *cfschema.Property) (string, error) { //nolint:unparam
	if property.MinItems != nil && property.MaxItems == nil {
		return fmt.Sprintf("validate.ArrayLenAtLeast(%d)", *property.MinItems), nil
	} else if property.MinItems == nil && property.MaxItems != nil {
		return fmt.Sprintf("validate.ArrayLenAtMost(%d)", *property.MaxItems), nil
	} else if property.MinItems != nil && property.MaxItems != nil {
		return fmt.Sprintf("validate.ArrayLenBetween(%d,%d)", *property.MinItems, *property.MaxItems), nil
	}

	return "", nil
}

// defaultValueAttributePlanModifier returns any AttributePlanModifier for the specified Property.
func defaultValueAttributePlanModifier(path []string, property *cfschema.Property) (Features, string, error) {
	var features Features

	switch v := property.Default.(type) {
	case nil:
		return features, "", nil
	//
	// Primitive types.
	//
	case bool:
		return features, fmt.Sprintf("DefaultValue(types.Bool{Value: %t})", v), nil
	case float64:
		switch propertyType := property.Type.String(); propertyType {
		case cfschema.PropertyTypeInteger:
			return features, fmt.Sprintf("DefaultValue(types.Int64{Value: %d})", int64(v)), nil
		case cfschema.PropertyTypeNumber:
			return features, fmt.Sprintf("DefaultValue(types.Float64{Value: %f})", v), nil
		default:
			return 0, "", fmt.Errorf("%s has invalid default value element type: %T", strings.Join(path, "/"), v)
		}
	case string:
		return features, fmt.Sprintf("DefaultValue(types.String{Value: %q})", v), nil

	//
	// Complex types.
	//
	case []interface{}:
		switch arrayType := aggregateType(property); arrayType {
		case aggregateNone:
			return 0, "", fmt.Errorf("%s has invalid default value type: %T", strings.Join(path, "/"), v)
		case aggregateSet:
			features |= UsesFrameworkAttr

			w := &strings.Builder{}
			fprintf(w, "DefaultValue(types.Set{ElemType:types.StringType, Elems: []attr.Value{\n")
			for _, elem := range v {
				switch v := elem.(type) {
				case string:
					fprintf(w, "types.String{Value: %q},\n", v)
				default:
					return 0, "", fmt.Errorf("%s has invalid default value element type: %T", strings.Join(path, "/"), v)
				}
			}
			fprintf(w, "}})")
			return features, w.String(), nil
		default:
			features |= UsesFrameworkAttr

			w := &strings.Builder{}
			fprintf(w, "DefaultValue(types.List{ElemType:types.StringType, Elems: []attr.Value{\n")
			for _, elem := range v {
				switch v := elem.(type) {
				case string:
					fprintf(w, "types.String{Value: %q},\n", v)
				default:
					return 0, "", fmt.Errorf("%s has invalid default value element type: %T", strings.Join(path, "/"), v)
				}
			}
			fprintf(w, "}})")
			return features, w.String(), nil
		}

	case map[string]interface{}:
		features |= UsesFrameworkAttr

		w := &strings.Builder{}
		fprintf(w, "DefaultValue(types.Object{\nAttrTypes: map[string]attr.Type{\n")
		for key1, v := range v {
			switch v := v.(type) {
			case bool:
				fprintf(w, "%q: types.BoolType,\n", naming.CloudFormationPropertyToTerraformAttribute(key1))
			case string:
				fprintf(w, "%q: types.StringType,\n", naming.CloudFormationPropertyToTerraformAttribute(key1))
			case map[string]interface{}:
				fprintf(w, "%q: types.ObjectType{\nAttrTypes: map[string]attr.Type{\n", naming.CloudFormationPropertyToTerraformAttribute(key1))
				for key2, v := range v {
					switch v := v.(type) {
					case bool:
						fprintf(w, "%q: types.BoolType,\n", naming.CloudFormationPropertyToTerraformAttribute(key2))
					case string:
						fprintf(w, "%q: types.StringType,\n", naming.CloudFormationPropertyToTerraformAttribute(key2))
					default:
						return 0, "", fmt.Errorf("%s has invalid default value element type: %T", strings.Join(append(path, key1, key2), "/"), v)
					}
				}
				fprintf(w, "},\n")
				fprintf(w, "},\n")
			default:
				return 0, "", fmt.Errorf("%s has invalid default value element type: %T", strings.Join(append(path, key1), "/"), v)
			}
		}
		fprintf(w, "},\n")
		fprintf(w, "Attrs: map[string]attr.Value{\n")
		for key1, v := range v {
			switch v := v.(type) {
			case bool:
				fprintf(w, "%q: types.Bool{Value: %t},\n", naming.CloudFormationPropertyToTerraformAttribute(key1), v)
			case string:
				fprintf(w, "%q: types.String{Value: %q},\n", naming.CloudFormationPropertyToTerraformAttribute(key1), v)
			case map[string]interface{}:
				fprintf(w, "%q: types.Object{\nAttrTypes: map[string]attr.Type{\n", naming.CloudFormationPropertyToTerraformAttribute(key1))
				for key2, v := range v {
					switch v.(type) {
					case bool:
						fprintf(w, "%q: types.BoolType,\n", naming.CloudFormationPropertyToTerraformAttribute(key2))
					case string:
						fprintf(w, "%q: types.StringType,\n", naming.CloudFormationPropertyToTerraformAttribute(key2))
					}
				}
				fprintf(w, "},\n")
				fprintf(w, "Attrs: map[string]attr.Value{\n")
				for key2, v := range v {
					switch v := v.(type) {
					case bool:
						fprintf(w, "%q: types.Bool{Value: %t},\n", naming.CloudFormationPropertyToTerraformAttribute(key2), v)
					case string:
						fprintf(w, "%q: types.String{Value: %q},\n", naming.CloudFormationPropertyToTerraformAttribute(key2), v)
					}
				}
				fprintf(w, "},\n")
				fprintf(w, "},\n")
			}
		}
		fprintf(w, "},\n")
		fprintf(w, "},\n")
		fprintf(w, ")")
		return features, w.String(), nil

	default:
		return 0, "", fmt.Errorf("%s has unsupported default value type: %T", strings.Join(path, "/"), v)
	}
}

type primitiveValidatorsGenerator func([]string, *cfschema.Property) (Features, []string, error)

// integerValidators returns any validators for the specified integer Property.
func integerValidators(path []string, property *cfschema.Property) (Features, []string, error) {
	var features Features

	if propertyType := property.Type.String(); propertyType != cfschema.PropertyTypeInteger {
		return features, nil, fmt.Errorf("invalid property type: %s", propertyType)
	}

	var validators []string

	if property.Minimum != nil && property.Maximum == nil {
		validators = append(validators, fmt.Sprintf("validate.IntAtLeast(%d)", *property.Minimum))
	} else if property.Minimum == nil && property.Maximum != nil {
		validators = append(validators, fmt.Sprintf("validate.IntAtMost(%d)", *property.Maximum))
	} else if property.Minimum != nil && property.Maximum != nil {
		validators = append(validators, fmt.Sprintf("validate.IntBetween(%d,%d)", *property.Minimum, *property.Maximum))
	}

	if property.Format != nil {
		if format := *property.Format; format != "int64" {
			return features, nil, fmt.Errorf("%s has unsupported format: %s", strings.Join(path, "/"), format)
		}
	}

	if len(property.Enum) > 0 {
		w := &strings.Builder{}
		fprintf(w, "validate.IntInSlice([]int{\n")
		for _, enum := range property.Enum {
			fprintf(w, "%d", int(enum.(float64)))
			fprintf(w, ",\n")
		}
		fprintf(w, "})")
		validators = append(validators, w.String())
	}

	return features, validators, nil
}

// numberValidators returns any validators for the specified number Property.
func numberValidators(path []string, property *cfschema.Property) (Features, []string, error) {
	var features Features

	if propertyType := property.Type.String(); propertyType != cfschema.PropertyTypeNumber {
		return features, nil, fmt.Errorf("invalid property type: %s", propertyType)
	}

	var validators []string

	if property.Minimum != nil && property.Maximum == nil {
		validators = append(validators, fmt.Sprintf("validate.FloatAtLeast(%f)", float64(*property.Minimum)))
	} else if property.Minimum == nil && property.Maximum != nil {
		validators = append(validators, fmt.Sprintf("validate.FloatAtMost(%f)", float64(*property.Maximum)))
	} else if property.Minimum != nil && property.Maximum != nil {
		validators = append(validators, fmt.Sprintf("validate.FloatBetween(%f,%f)", float64(*property.Minimum), float64(*property.Maximum)))
	}

	if property.Format != nil {
		if format := *property.Format; format != "double" {
			return features, nil, fmt.Errorf("%s has unsupported format: %s", strings.Join(path, "/"), format)
		}
	}

	if len(property.Enum) > 0 {
		return features, nil, fmt.Errorf("%s has enumerated values", strings.Join(path, "/"))
	}

	return features, validators, nil
}

// stringValidators returns any validators for the specified string Property.
func stringValidators(path []string, property *cfschema.Property) (Features, []string, error) {
	var features Features

	if propertyType := property.Type.String(); propertyType != cfschema.PropertyTypeString {
		return features, nil, fmt.Errorf("invalid property type: %s", propertyType)
	}

	var validators []string

	if property.MinLength != nil && property.MaxLength == nil {
		validators = append(validators, fmt.Sprintf("validate.StringLenAtLeast(%d)", *property.MinLength))
	} else if property.MinLength == nil && property.MaxLength != nil {
		validators = append(validators, fmt.Sprintf("validate.StringLenAtMost(%d)", *property.MaxLength))
	} else if property.MinLength != nil && property.MaxLength != nil {
		validators = append(validators, fmt.Sprintf("validate.StringLenBetween(%d,%d)", *property.MinLength, *property.MaxLength))
	}

	if property.Pattern != nil && *property.Pattern != "" {
		features |= UsesRegexpInValidation
		validators = append(validators, fmt.Sprintf("validate.StringMatch(regexp.MustCompile(%q), \"\")", *property.Pattern))
	}

	if property.Format != nil {
		switch format := *property.Format; format {
		case "date-time":
			validators = append(validators, "validate.IsRFC3339Time()")
		case "string":
		case "uri":
			validators = append(validators, "validate.IsURI()")
		default:
			// TODO
			// return nil, fmt.Errorf("%s has unsupported format: %s", strings.Join(path, "/"), format)
		}
	}

	if len(property.Enum) > 0 {
		w := &strings.Builder{}
		fprintf(w, "validate.StringInSlice([]string{\n")
		for _, enum := range property.Enum {
			fprintf(w, "\"")
			fprintf(w, enum.(string))
			fprintf(w, "\",\n")
		}
		fprintf(w, "})")
		validators = append(validators, w.String())
	}

	return features, validators, nil
}

func addPropertyRequiredAttributes(writer io.Writer, p *cfschema.PropertySubschema) int {
	var nRequired int

	if len(p.AllOf) > 0 {
		var n int
		w := &strings.Builder{}

		fprintf(w, "validate.AllOfRequired(\n")
		for _, a := range p.AllOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		fprintf(w, "),\n")

		if n > 0 {
			fprintf(writer, w.String())
		}

		nRequired += n
	}
	if len(p.AnyOf) > 0 {
		var n int
		w := &strings.Builder{}

		fprintf(w, "validate.AnyOfRequired(\n")
		for _, a := range p.AnyOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		fprintf(w, "),\n")

		if n > 0 {
			fprintf(writer, w.String())
		}

		nRequired += n
	}
	if len(p.OneOf) > 0 {
		var n int
		w := &strings.Builder{}

		fprintf(w, "validate.OneOfRequired(\n")
		for _, a := range p.OneOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		fprintf(w, "),\n")

		if n > 0 {
			fprintf(writer, w.String())
		}

		nRequired += n
	}
	if len(p.Required) > 0 {
		fprintf(writer, "validate.Required(\n")
		for _, r := range p.Required {
			fprintf(writer, "%q,\n", naming.CloudFormationPropertyToTerraformAttribute(r))
			nRequired++
		}
		fprintf(writer, "),\n")
	}

	return nRequired
}

func addSchemaCompositionRequiredAttributes(writer io.Writer, r schemaComposition) int {
	var nRequired int

	if allOf := r.All(); len(allOf) > 0 {
		var n int
		w := &strings.Builder{}

		fprintf(w, "validate.AllOfRequired(\n")
		for _, a := range allOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		fprintf(w, "),\n")

		if n > 0 {
			fprintf(writer, w.String())
		}

		nRequired += n
	}
	if anyOf := r.Any(); len(anyOf) > 0 {
		var n int
		w := &strings.Builder{}

		fprintf(w, "validate.AnyOfRequired(\n")
		for _, a := range anyOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		fprintf(w, "),\n")

		if n > 0 {
			fprintf(writer, w.String())
		}

		nRequired += n
	}
	if oneOf := r.One(); len(oneOf) > 0 {
		var n int
		w := &strings.Builder{}

		fprintf(w, "validate.OneOfRequired(\n")
		for _, a := range oneOf {
			n += addPropertyRequiredAttributes(w, a)
		}
		fprintf(w, "),\n")

		if n > 0 {
			fprintf(writer, w.String())
		}

		nRequired += n
	}

	return nRequired
}

func propertyRequiredAttributesValidator(p *cfschema.Property) (string, error) { //nolint:unparam
	if p == nil {
		return "", nil
	}

	w := &strings.Builder{}
	fprintf(w, "validate.RequiredAttributes(\n")
	nRequired := addSchemaCompositionRequiredAttributes(w, property(*p))
	fprintf(w, ")")

	if nRequired == 0 {
		return "", nil
	}

	return w.String(), nil
}

func resourceRequiredAttributesValidator(r *cfschema.Resource) (string, error) { //nolint:unparam
	if r == nil {
		return "", nil
	}

	w := &strings.Builder{}
	nRequired := addSchemaCompositionRequiredAttributes(w, resource(*r))

	if nRequired == 0 {
		return "", nil
	}

	return w.String(), nil
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
