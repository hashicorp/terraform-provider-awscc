package generic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
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
	provider       tfcloudcontrol.Provider
	dataSourceType *singularDataSourceType
}

func newGenericSingularDataSource(provider tfsdk.Provider, singularDataSourceType *singularDataSourceType) tfsdk.DataSource {
	return &singularDataSource{
		provider:       provider.(tfcloudcontrol.Provider),
		dataSourceType: singularDataSourceType,
	}
}

func (sd *singularDataSource) Read(ctx context.Context, request tfsdk.ReadDataSourceRequest, response *tfsdk.ReadDataSourceResponse) {
	ctx = sd.cfnTypeContext(ctx)

	traceEntry(ctx, "SingularDataSource.Read")

	conn := sd.provider.CloudControlApiClient(ctx)

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
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))

		return
	}

	translator := toTerraform{cfToTfNameMap: sd.dataSourceType.cfToTfNameMap}
	schema := &currentConfig.Schema
	val, err := translator.FromString(ctx, schema, aws.ToString(description.Properties))

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to create a Terraform State value from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
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

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "SingularDataSource.Read")
}

// describe returns the live state of the specified resource.
func (sd *singularDataSource) describe(ctx context.Context, conn *cloudcontrol.Client, id string) (*cctypes.ResourceDescription, error) {
	return tfcloudcontrol.FindResourceByTypeNameAndID(ctx, conn, sd.provider.RoleARN(ctx), sd.dataSourceType.cfTypeName, id)
}

// getId returns the data source's primary identifier value from Config.
func (sd *singularDataSource) getId(ctx context.Context, config *tfsdk.Config) (string, error) {
	var val string
	diags := config.GetAttribute(ctx, idAttributePath, &val)

	if diags.HasError() {
		return "", tfresource.DiagsError(diags)
	}

	return val, nil
}

// setId sets the data source's primary identifier value in State.
func (sd *singularDataSource) setId(ctx context.Context, val string, state *tfsdk.State) error {
	diags := state.SetAttribute(ctx, idAttributePath, val)

	if diags.HasError() {
		return tfresource.DiagsError(diags)
	}

	return nil
}

// cfnTypeContext injects the CloudFormation type name into logger contexts.
func (sd *singularDataSource) cfnTypeContext(ctx context.Context) context.Context {
	ctx = tflog.With(ctx, LoggingKeyCFNType, sd.dataSourceType.cfTypeName)

	return ctx
}
