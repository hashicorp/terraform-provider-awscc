package generic

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// DataSourceTypeOptionsFunc is a type alias for a data source type functional option.
type DataSourceTypeOptionsFunc func(*genericDataSourceType) error

// dataSourceWithAttributeNameMap is a helper function to construct functional options
// that set a data source type's attribute name maps.
// If multiple dataSourceWithAttributeNameMap calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithAttributeNameMap(v map[string]string) DataSourceTypeOptionsFunc {
	return func(o *genericDataSourceType) error {
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
func dataSourceWithCloudFormationTypeName(v string) DataSourceTypeOptionsFunc {
	return func(o *genericDataSourceType) error {
		o.cfTypeName = v

		return nil
	}
}

// dataSourceWithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema.
// If multiple dataSourceWithTerraformSchema calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithTerraformSchema(v tfsdk.Schema) DataSourceTypeOptionsFunc {
	return func(o *genericDataSourceType) error {
		o.tfSchema = v

		return nil
	}
}

// dataSourceWithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name.
// If multiple dataSourceWithTerraformTypeName calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithTerraformTypeName(v string) DataSourceTypeOptionsFunc {
	return func(o *genericDataSourceType) error {
		o.tfTypeName = v

		return nil
	}
}

type DataSourceTypeOptions []DataSourceTypeOptionsFunc

// WithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name map, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceTypeOptions) WithAttributeNameMap(v map[string]string) DataSourceTypeOptions {
	return append(opts, dataSourceWithAttributeNameMap(v))
}

// WithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceTypeOptions) WithCloudFormationTypeName(v string) DataSourceTypeOptions {
	return append(opts, dataSourceWithCloudFormationTypeName(v))
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceTypeOptions) WithTerraformSchema(v tfsdk.Schema) DataSourceTypeOptions {
	return append(opts, dataSourceWithTerraformSchema(v))
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceTypeOptions) WithTerraformTypeName(v string) DataSourceTypeOptions {
	return append(opts, dataSourceWithTerraformTypeName(v))
}

// genericDataSourceType implements provider.DataSourceType
type genericDataSourceType struct {
	cfToTfNameMap map[string]string // Map of CloudFormation property name to Terraform attribute name
	cfTypeName    string            // CloudFormation type name for the resource type
	tfSchema      tfsdk.Schema      // Terraform schema for the data source type
	tfTypeName    string            // Terraform type name for data source type
}
