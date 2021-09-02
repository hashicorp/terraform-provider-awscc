package generic

import "github.com/hashicorp/terraform-plugin-framework/tfsdk"

// DataSourceTypeOptionsFunc is a type alias for a DataSource type functional option.
type DataSourceTypeOptionsFunc func(*DataSourceType) error

type DataSourceTypeOptions []DataSourceTypeOptionsFunc

// DataSourceType implements tfsdk.DataSourceType
type DataSourceType struct {
	cfToTfNameMap map[string]string // Map of CloudFormation property name to Terraform attribute name
	cfTypeName    string            // CloudFormation type name for the resource type
	tfSchema      tfsdk.Schema      // Terraform schema for the data source type
	tfTypeName    string            // Terraform type name for data source type
}

// FromCloudFormationAndTerraform is a helper function to construct functional options
// that set a data source type's CloudFormation type name, Terraform type name, and Terraform schema.
// If multiple FromCloudFormationAndTerraform calls are made, the last call overrides
// the previous calls' values.
func FromCloudFormationAndTerraform(cfTypeName, tfTypeName string, schema tfsdk.Schema) DataSourceTypeOptionsFunc {
	return func(o *DataSourceType) error {
		o.cfTypeName = cfTypeName
		o.tfTypeName = tfTypeName
		o.tfSchema = schema
		return nil
	}
}

// FromCloudFormationAndTerraform is a helper function to construct functional options
// that set a data source type's CloudFormation type name, Terraform type name, and Terraform schema
// and append that function to the current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceTypeOptions) FromCloudFormationAndTerraform(cfTypeName, tfTypeName string, schema tfsdk.Schema) DataSourceTypeOptions {
	return append(opts, FromCloudFormationAndTerraform(cfTypeName, tfTypeName, schema))
}
