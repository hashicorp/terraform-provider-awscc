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
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
)

// NewPluralDataSource returns a new plural DataSource from the specified variadic list of functional options.
// It's public as it's called from generated code.
func NewPluralDataSource(_ context.Context, optFns ...DataSourceOptionsFunc) (datasource.DataSource, error) {
	v := &genericPluralDataSource{}

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
type genericPluralDataSource struct {
	genericDataSource
	provider tfcloudcontrol.Provider
}

func (pd *genericPluralDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = pd.tfTypeName
}

func (pd *genericPluralDataSource) Schema(_ context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = pd.tfSchema
}

func (pd *genericPluralDataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) { //nolint:unparam
	if v := request.ProviderData; v != nil {
		pd.provider = v.(tfcloudcontrol.Provider)
	}
}

func (pd *genericPluralDataSource) Read(ctx context.Context, _ datasource.ReadRequest, response *datasource.ReadResponse) {
	ctx = pd.bootstrapContext(ctx)

	traceEntry(ctx, "PluralDataSource.Read")

	conn := pd.provider.CloudControlAPIClient(ctx)

	descriptions, err := pd.list(ctx, conn)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("CloudControl", "ListResources", err))

		return
	}

	val := GetCloudControlResourceDescriptionsValue(pd.provider.Region(ctx), descriptions)

	response.State = tfsdk.State{
		Schema: pd.tfSchema,
		Raw:    val,
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "PluralDataSource.Read")
}

// list returns the ResourceDescriptions of the specified CloudFormation type.
func (pd *genericPluralDataSource) list(ctx context.Context, conn *cloudcontrol.Client) ([]cctypes.ResourceDescription, error) {
	return tfcloudcontrol.ListResourcesByTypeName(ctx, conn, pd.provider.RoleARN(ctx), pd.cfTypeName)
}

// bootstrapContext injects the CloudFormation type name into logger contexts.
func (pd *genericPluralDataSource) bootstrapContext(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, LoggingKeyCFNType, pd.cfTypeName)
	ctx = pd.provider.RegisterLogger(ctx)

	return ctx
}

// GetCloudControlResourceDescriptionsValue returns the Terraform Value for the specified Cloud Control API ResourceDescriptions.
func GetCloudControlResourceDescriptionsValue(id string, descriptions []cctypes.ResourceDescription) tftypes.Value {
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
