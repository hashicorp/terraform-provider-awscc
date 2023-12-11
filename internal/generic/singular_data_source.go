// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// NewSingularDataSource returns a new singular DataSource from the specified variadic list of functional options.
// It's public as it's called from generated code.
func NewSingularDataSource(_ context.Context, optFns ...DataSourceOptionsFunc) (datasource.DataSource, error) {
	v := &genericSingularDataSource{}

	for _, optFn := range optFns {
		err := optFn(&v.genericDataSource)

		if err != nil {
			return nil, err
		}
	}

	if v.cfTypeName == "" {
		return nil, fmt.Errorf("no CloudFormation type name specified")
	}

	if v.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}

	return v, nil
}

// Implements datasource.DataSource
type genericSingularDataSource struct {
	genericDataSource
	provider tfcloudcontrol.Provider
}

func (sd *genericSingularDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = sd.tfTypeName
}

func (sd *genericSingularDataSource) Schema(_ context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = sd.tfSchema
}

func (sd *genericSingularDataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) { //nolint:unparam
	if v := request.ProviderData; v != nil {
		sd.provider = v.(tfcloudcontrol.Provider)
	}
}

func (sd *genericSingularDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	ctx = sd.cfnTypeContext(ctx)

	traceEntry(ctx, "SingularDataSource.Read")

	conn := sd.provider.CloudControlAPIClient(ctx)

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

	translator := toTerraform{cfToTfNameMap: sd.cfToTfNameMap}
	schema := currentConfig.Schema
	val, err := translator.FromString(ctx, schema, aws.ToString(description.Properties))

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to create a Terraform State value from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	response.State = tfsdk.State{
		Schema: schema,
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
func (sd *genericSingularDataSource) describe(ctx context.Context, conn *cloudcontrol.Client, id string) (*cctypes.ResourceDescription, error) {
	return tfcloudcontrol.FindResourceByTypeNameAndID(ctx, conn, sd.provider.RoleARN(ctx), sd.cfTypeName, id)
}

// getId returns the data source's primary identifier value from Config.
func (sd *genericSingularDataSource) getId(ctx context.Context, config *tfsdk.Config) (string, error) {
	var val string
	diags := config.GetAttribute(ctx, idAttributePath, &val)

	if diags.HasError() {
		return "", tfresource.DiagnosticsError(diags)
	}

	return val, nil
}

// setId sets the data source's primary identifier value in State.
func (sd *genericSingularDataSource) setId(ctx context.Context, val string, state *tfsdk.State) error {
	diags := state.SetAttribute(ctx, idAttributePath, val)

	if diags.HasError() {
		return tfresource.DiagnosticsError(diags)
	}

	return nil
}

// cfnTypeContext injects the CloudFormation type name into logger contexts.
func (sd *genericSingularDataSource) cfnTypeContext(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, LoggingKeyCFNType, sd.cfTypeName)

	return ctx
}
