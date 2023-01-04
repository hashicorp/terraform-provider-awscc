package generic

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// DataSourceOptionsFunc is a type alias for a data source type functional option.
type DataSourceOptionsFunc func(*genericDataSource) error

// dataSourceWithAttributeNameMap is a helper function to construct functional options
// that set a data source type's attribute name maps.
// If multiple dataSourceWithAttributeNameMap calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithAttributeNameMap(v map[string]string) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		if _, ok := v["id"]; !ok {
			// Synthesize a mapping for the reserved top-level "id" attribute.
			v["id"] = "ID"
		}

		cfToTfNameMap := make(map[string]string, len(v))

		for tfName, cfName := range v {
			_, ok := cfToTfNameMap[cfName]
			if ok {
				return fmt.Errorf("duplicate attribute name mapping for CloudFormation property %s", cfName)
			}
			cfToTfNameMap[cfName] = tfName
		}

		o.cfToTfNameMap = cfToTfNameMap

		return nil
	}
}

// dataSourceWithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name.
// If multiple dataSourceWithCloudFormationTypeName calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithCloudFormationTypeName(v string) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		o.cfTypeName = v

		return nil
	}
}

// dataSourceWithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema.
// If multiple dataSourceWithTerraformSchema calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithTerraformSchema(v schema.Schema) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		o.tfSchema = v

		return nil
	}
}

// dataSourceWithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name.
// If multiple dataSourceWithTerraformTypeName calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithTerraformTypeName(v string) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		o.tfTypeName = v

		return nil
	}
}

type DataSourceOptions []DataSourceOptionsFunc

// WithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name map, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithAttributeNameMap(v map[string]string) DataSourceOptions {
	return append(opts, dataSourceWithAttributeNameMap(v))
}

// WithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithCloudFormationTypeName(v string) DataSourceOptions {
	return append(opts, dataSourceWithCloudFormationTypeName(v))
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithTerraformSchema(v schema.Schema) DataSourceOptions {
	return append(opts, dataSourceWithTerraformSchema(v))
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithTerraformTypeName(v string) DataSourceOptions {
	return append(opts, dataSourceWithTerraformTypeName(v))
}

type genericDataSource struct {
	cfToTfNameMap map[string]string // Map of CloudFormation property name to Terraform attribute name
	cfTypeName    string            // CloudFormation type name for the resource type
	tfSchema      schema.Schema     // Terraform schema for the data source type
	tfTypeName    string            // Terraform type name for data source type
}
