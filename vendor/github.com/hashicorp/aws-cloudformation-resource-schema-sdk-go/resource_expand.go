package cfschema

import (
	"fmt"
)

// Expand replaces all Definition and Property JSON Pointer references with their content.
// This functionality removes the need for recursive logic when accessing Definitions and Properties.
// In unresolved form nested properties are not allowed, instead nested properties use a '$ref' JSON Pointer to reference a definition.
// See https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-properties.
func (r *Resource) Expand() error {
	if r == nil {
		return nil
	}

	err := r.ResolveProperties(r.Definitions)

	if err != nil {
		return fmt.Errorf("error expanding Resource (%s) Definitions: %w", *r.TypeName, err)
	}

	err = r.ResolveProperties(r.Properties)

	if err != nil {
		return fmt.Errorf("error expanding Resource (%s) Properties: %w", *r.TypeName, err)
	}

	return nil
}

// ResolveProperties resolves all References in a top-level name-to-property map.
// In unresolved form nested properties are not allowed so we don't need to recurse.
func (r *Resource) ResolveProperties(properties map[string]*Property) error {
	for propertyName, property := range properties {
		resolved, err := r.ResolveProperty(property)

		if err != nil {
			return fmt.Errorf("error resolving %s: %w", propertyName, err)
		}

		if resolved {
			// For example:
			// "Configuration": {
			//   "$ref": "#/definitions/ClusterConfiguration"
			// },
			continue
		}

		switch property.Type.String() {
		case PropertyTypeArray:
			// For example:
			// "DefaultCapacityProviderStrategy": {
			//   "type": "array",
			//   "items": {
			//     "$ref": "#/definitions/CapacityProviderStrategyItem"
			//   }
			// },
			_, err = r.ResolveProperty(property.Items)

			if err != nil {
				return fmt.Errorf("error resolving %s Items: %w", propertyName, err)
			}
		case PropertyTypeObject:
			for objPropertyName, objProperty := range property.Properties {
				resolved, err := r.ResolveProperty(objProperty)

				if err != nil {
					return fmt.Errorf("error resolving %s Property (%s): %w", propertyName, objPropertyName, err)
				}

				if resolved {
					// For example:
					// "ClusterConfiguration": {
					//   "type": "object",
					//   "properties": {
					//     "ExecuteCommandConfiguration": {
					//       "$ref": "#/definitions/ExecuteCommandConfiguration"
					//     }
					//   }
					// },
					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					// For example:
					// "LambdaContainerParams": {
					//   "type": "object",
					//   "properties": {
					//     "Volumes": {
					//       "type": "array",
					//       "items": {
					//         "$ref": "#/definitions/LambdaVolumeMount"
					//       }
					//     }
					//   }
					// },
					_, err = r.ResolveProperty(objProperty.Items)

					if err != nil {
						return fmt.Errorf("error resolving %s Property (%s) Items: %w", propertyName, objPropertyName, err)
					}
				case PropertyTypeObject:
					// Pragmatically resolve any References at this level even though they are not allowed.
					for objPropertyName2, objProperty2 := range objProperty.Properties {
						_, err := r.ResolveProperty(objProperty2)

						if err != nil {
							return fmt.Errorf("error resolving %s Property (%s) Property (%s): %w", propertyName, objPropertyName, objPropertyName2, err)
						}
					}

					for pattern, patternProperty := range objProperty.PatternProperties {
						// For example:
						// "LambdaFunctionRecipeSource": {
						//   "type": "object",
						//   "properties": {
						//     "ComponentDependencies": {
						//       "type": "object",
						//       "patternProperties": {
						//         "": {
						//           "$ref": "#/definitions/ComponentDependencyRequirement"
						//         }
						//       }
						//     }
						//   }
						// },
						_, err = r.ResolveProperty(patternProperty)

						if err != nil {
							return fmt.Errorf("error resolving %s Property (%s) Pattern(%s): %w", propertyName, objPropertyName, pattern, err)
						}
					}
				}
			}

			for patternName, objProperty := range property.PatternProperties {
				resolved, err := r.ResolveProperty(objProperty)

				if err != nil {
					return fmt.Errorf("error resolving %s pattern Property (%s): %w", propertyName, patternName, err)
				}

				if resolved {
					// For example:
					// "Tags": {
					//   "type": "object",
					//   "patternProperties": {
					//     "": {
					//       "$ref": "#/definitions/TagValue"
					//     }
					//   }
					// },
					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					// For example:
					// "Tags": {
					//   "type": "object",
					//   "patternProperties": {
					//     "": {
					//       "type": "array",
					//       "items": {
					//         "$ref": "#/definitions/TagValue"
					//       }
					//     }
					//   }
					// },
					_, err = r.ResolveProperty(objProperty.Items)

					if err != nil {
						return fmt.Errorf("error resolving %s Property (%s) Items: %w", propertyName, patternName, err)
					}
				}
			}
		}
	}

	return nil
}

// ResolveProperty resolves any Reference (JSON Pointer) in a Property.
// Returns whether a Reference was resolved.
func (r *Resource) ResolveProperty(property *Property) (bool, error) {
	if property != nil && property.Ref != nil {
		defaultValue := property.Default
		ref := property.Ref
		resolution, err := r.ResolveReference(*ref)

		if err != nil {
			return false, err
		}

		*property = *resolution

		// Ensure that any default value is not lost.
		if defaultValue != nil {
			property.Default = defaultValue
		}

		return true, nil
	}

	return false, nil
}
