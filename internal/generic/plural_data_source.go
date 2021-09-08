package generic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	tflog "github.com/hashicorp/terraform-plugin-log"
	tfcloudformation "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudformation"
)

type PluralDataSourceType DataSourceType

// NewPluralDataSourceType returns a new PluralDataSourceType from the specified variadic list of functional options.
// It's public as it's called from generated code.
func NewPluralDataSourceType(_ context.Context, optFns ...DataSourceTypeOptionsFunc) (tfsdk.DataSourceType, error) {
	dataSourceType := &DataSourceType{}

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

	var pdt *PluralDataSourceType

	return pdt.New(dataSourceType), nil
}

func (pdt *PluralDataSourceType) New(dst *DataSourceType) *PluralDataSourceType {
	return &PluralDataSourceType{
		cfTypeName: dst.cfTypeName,
		tfSchema:   dst.tfSchema,
		tfTypeName: dst.tfTypeName,
	}
}

func (pdt *PluralDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return pdt.tfSchema, nil
}

func (pdt *PluralDataSourceType) NewDataSource(ctx context.Context, provider tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return newGenericPluralDataSource(provider, pdt), nil
}

// Implements tfsdk.DataSource
type pluralDataSource struct {
	provider       tfcloudformation.Provider
	dataSourceType *PluralDataSourceType
}

func newGenericPluralDataSource(provider tfsdk.Provider, pluralDataSourceType *PluralDataSourceType) tfsdk.DataSource {
	return &pluralDataSource{
		provider:       provider.(tfcloudformation.Provider),
		dataSourceType: pluralDataSourceType,
	}
}

func (pd *pluralDataSource) Read(ctx context.Context, _ tfsdk.ReadDataSourceRequest, response *tfsdk.ReadDataSourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := pd.dataSourceType.cfTypeName
	tfTypeName := pd.dataSourceType.tfTypeName

	tflog.Debug(ctx, "DataSource.Read enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := pd.provider.CloudFormationClient(ctx)

	descriptions, err := pd.list(ctx, conn)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("CloudFormation", "ListResources", err))

		return
	}

	val := GetCloudFormationResourceDescriptionsValue(pd.provider.Region(ctx), descriptions)

	response.State = tfsdk.State{
		Schema: pd.dataSourceType.tfSchema,
		Raw:    val,
	}

	tflog.Debug(ctx, "Response.State.Raw", "value", hclog.Fmt("%v", response.State.Raw))

	tflog.Debug(ctx, "DataSource.Read exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

// list returns the ResourceDescriptions of the specified CloudFormation type.
func (pd *pluralDataSource) list(ctx context.Context, conn *cloudformation.Client) ([]cftypes.ResourceDescription, error) {
	return tfcloudformation.ListResourcesByTypeName(ctx, conn, pd.provider.RoleARN(ctx), pd.dataSourceType.cfTypeName)
}

// GetCloudFormationResourceDescriptionsValue returns the Terraform Value for the specified CloudFormation ResourceDescriptions.
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
