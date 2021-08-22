package generic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	tflog "github.com/hashicorp/terraform-plugin-log"
	tfcloudformation "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudformation"
)

// PluralDataSourceTypeOptionsFunc is a type alias for a dataSource type functional option.
type PluralDataSourceTypeOptionsFunc func(*PluralDataSourceType) error

type PluralDataSourceTypeOptions []PluralDataSourceTypeOptionsFunc

// PluralDataSourceType implements tfsdk.DataSourceType
type PluralDataSourceType struct {
	cfTypeName string       // CloudFormation type name for the resource type
	tfSchema   tfsdk.Schema // Terraform schema for the resource type
	tfTypeName string       // Terraform type name for resource type
}

func FromCloudFormationAndTerraform(cfTypeName, tfTypeName string, schema tfsdk.Schema) PluralDataSourceTypeOptionsFunc {
	return func(o *PluralDataSourceType) error {
		o.cfTypeName = cfTypeName
		o.tfTypeName = tfTypeName
		o.tfSchema = schema
		return nil
	}
}

func (opts PluralDataSourceTypeOptions) FromCloudFormationAndTerraform(cfTypeName, tfTypeName string, schema tfsdk.Schema) PluralDataSourceTypeOptions {
	return append(opts, FromCloudFormationAndTerraform(cfTypeName, tfTypeName, schema))
}

// NewPluralDataSourceType returns a new DataSourceType from the specified varidaic list of functional options.
// It's public as it's called from generated code.
func NewPluralDataSourceType(_ context.Context, optFns ...PluralDataSourceTypeOptionsFunc) (tfsdk.DataSourceType, error) {
	dataSourceType := &PluralDataSourceType{}

	for _, optFn := range optFns {
		err := optFn(dataSourceType)

		if err != nil {
			return nil, err
		}
	}

	if dataSourceType.cfTypeName == "" {
		return nil, fmt.Errorf("no CloudFormation type name specified")
	}
	if dataSourceType.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}

	return dataSourceType, nil
}

func (pdt *PluralDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, []*tfprotov6.Diagnostic) {
	return pdt.tfSchema, nil
}

func (pdt *PluralDataSourceType) NewDataSource(ctx context.Context, provider tfsdk.Provider) (tfsdk.DataSource, []*tfprotov6.Diagnostic) {
	return newGenericPluralDataSource(provider, pdt), nil
}

// Implements tfsdk.DataSource
type pluralDataSource struct {
	provider       tfcloudformation.Provider
	dataSourceType *PluralDataSourceType
}

func newGenericPluralDataSource(provider tfsdk.Provider, pluraldataSourceType *PluralDataSourceType) tfsdk.DataSource {
	return &pluralDataSource{
		provider:       provider.(tfcloudformation.Provider),
		dataSourceType: pluraldataSourceType,
	}
}

func (pd *pluralDataSource) Read(ctx context.Context, _ tfsdk.ReadDataSourceRequest, response *tfsdk.ReadDataSourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := pd.dataSourceType.cfTypeName
	tfTypeName := pd.dataSourceType.tfTypeName

	tflog.Debug(ctx, "DataSource.Read enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := pd.provider.CloudFormationClient(ctx)

	descriptions, err := pd.describe(ctx, conn)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "ListResources", err))

		return
	}

	val := GetCloudFormationResourceDescriptionsValue(pd.provider.RoleARN(ctx), descriptions)

	response.State = tfsdk.State{
		Schema: pd.dataSourceType.tfSchema,
		Raw:    val,
	}

	tflog.Debug(ctx, "Response.State.Raw", "value", hclog.Fmt("%v", response.State.Raw))

	tflog.Debug(ctx, "DataSource.Read exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

// describe returns the live state of the specified data source.
func (pd *pluralDataSource) describe(ctx context.Context, conn *cloudformation.Client) ([]cftypes.ResourceDescription, error) {
	return tfcloudformation.ListResourcesByTypeName(ctx, conn, pd.provider.RoleARN(ctx), pd.dataSourceType.cfTypeName)
}

// GetCloudFormationResourceModelValue returns the Terraform Value for the specified CloudFormation ResourceModel (string).
func GetCloudFormationResourceDescriptionsValue(id string, descriptions []cftypes.ResourceDescription) tftypes.Value {
	m := map[string]tftypes.Value{
		"id": tftypes.NewValue(tftypes.String, id),
	}

	ids := make([]tftypes.Value, 0, len(descriptions))

	for _, description := range descriptions {
		ids = append(ids, tftypes.NewValue(tftypes.String, aws.ToString(description.Identifier)))
	}

	m["ids"] = tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, ids)

	return tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"id":  tftypes.String,
			"ids": tftypes.Set{ElementType: tftypes.String},
		}}, m)
}
