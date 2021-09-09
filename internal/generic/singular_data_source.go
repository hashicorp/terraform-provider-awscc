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
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	tfcloudformation "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// singularDataSourceType is a type alias for a data source type.
type singularDataSourceType dataSourceType

// NewSingularDataSourceType returns a new singularDataSourceType from the specified variadic list of functional options.
// It's public as it's called from generated code.
func NewSingularDataSourceType(_ context.Context, optFns ...DataSourceTypeOptionsFunc) (tfsdk.DataSourceType, error) {
	dataSourceType := &dataSourceType{}

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

	singularDataSourceType := singularDataSourceType(*dataSourceType)

	return &singularDataSourceType, nil
}

func (sdt *singularDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return sdt.tfSchema, nil
}

func (sdt *singularDataSourceType) NewDataSource(ctx context.Context, provider tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return newGenericSingularDataSource(provider, sdt), nil
}

// Implements tfsdk.DataSource
type singularDataSource struct {
	provider       tfcloudformation.Provider
	dataSourceType *singularDataSourceType
}

func newGenericSingularDataSource(provider tfsdk.Provider, singularDataSourceType *singularDataSourceType) tfsdk.DataSource {
	return &singularDataSource{
		provider:       provider.(tfcloudformation.Provider),
		dataSourceType: singularDataSourceType,
	}
}

func (sd *singularDataSource) Read(ctx context.Context, request tfsdk.ReadDataSourceRequest, response *tfsdk.ReadDataSourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := sd.dataSourceType.cfTypeName
	tfTypeName := sd.dataSourceType.tfTypeName

	tflog.Debug(ctx, "DataSource.Read enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := sd.provider.CloudFormationClient(ctx)

	currentConfig := &request.Config

	id, err := sd.getId(ctx, currentConfig)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	description, err := sd.describe(ctx, conn, id)

	if tfresource.NotFound(err) {
		response.Diagnostics = append(response.Diagnostics, DataSourceNotFoundDiag(err))

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "GetResource", err))

		return
	}

	translator := toTerraform{cfToTfNameMap: sd.dataSourceType.cfToTfNameMap}
	schema := &currentConfig.Schema
	val, err := translator.FromString(ctx, schema, aws.ToString(description.ResourceModel))

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to create a Terraform State value from a CloudFormation Resource Model. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	response.State = tfsdk.State{
		Schema: *schema,
		Raw:    val,
	}

	err = sd.setId(ctx, aws.ToString(description.Identifier), &response.State)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotSetDiag(err))

		return
	}

	tflog.Debug(ctx, "Response.State.Raw", "value", hclog.Fmt("%v", response.State.Raw))

	tflog.Debug(ctx, "DataSource.Read exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

// describe returns the live state of the specified resource.
func (sd *singularDataSource) describe(ctx context.Context, conn *cloudformation.Client, id string) (*cftypes.ResourceDescription, error) {
	return tfcloudformation.FindResourceByTypeNameAndID(ctx, conn, sd.provider.RoleARN(ctx), sd.dataSourceType.cfTypeName, id)
}

// getId returns the data source's primary identifier value from Config.
func (sd *singularDataSource) getId(ctx context.Context, config *tfsdk.Config) (string, error) {
	val, diags := config.GetAttribute(ctx, idAttributePath)

	if diags.HasError() {
		return "", tfresource.DiagsError(diags)
	}

	if val, ok := val.(types.String); ok {
		return val.Value, nil
	}

	return "", fmt.Errorf("invalid identifier type %T", val)
}

// setId sets the data source's primary identifier value in State.
func (sd *singularDataSource) setId(ctx context.Context, val string, state *tfsdk.State) error {
	diags := state.SetAttribute(ctx, idAttributePath, val)

	if diags.HasError() {
		return tfresource.DiagsError(diags)
	}

	return nil
}
