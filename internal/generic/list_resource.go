// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"fmt"
	"iter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
)

var _ list.ListResource = &genericResource{}

func NewListResource(resource func(context.Context) (resource.Resource, error)) func(context.Context) (list.ListResource, error) {
	return func(ctx context.Context) (list.ListResource, error) {
		res, err := resource(ctx)
		if err != nil {
			return nil, err
		}

		return res.(*genericResource), nil
	}
}

func (r *genericResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, response *list.ListResourceSchemaResponse) {
	response.Schema = listschema.Schema{
		Attributes: map[string]listschema.Attribute{},
	}
}

func (r *genericResource) List(ctx context.Context, request list.ListRequest, stream *list.ListResultsStream) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.List")

	conn := r.provider.CloudControlAPIClient(ctx)

	stream.Results = func(yield func(list.ListResult) bool) {
		result := request.NewListResult(ctx)
		for description, err := range r.stream(ctx, conn) {
			if err != nil {
				result = list.ListResult{
					Diagnostics: diag.Diagnostics{
						diag.NewErrorDiagnostic(
							"Error Listing Remote Resources",
							fmt.Sprintf("Error: %s", err),
						),
					},
				}
				yield(result)
				return
			}

			// TODO cleanup initializing identity
			result.Diagnostics.Append(result.Identity.SetAttribute(ctx, path.Root("bucket_name"), types.StringNull())...)
			result.Diagnostics.Append(result.Identity.SetAttribute(ctx, path.Root("account_id"), types.StringNull())...)
			if result.Diagnostics.HasError() {
				result = list.ListResult{Diagnostics: result.Diagnostics}
				yield(result)
				return
			}

			obj, err := r.describe(ctx, conn, aws.ToString(description.Identifier))
			if err != nil {
				result = list.ListResult{
					Diagnostics: diag.Diagnostics{
						diag.NewErrorDiagnostic(
							"Error Listing Remote Resources",
							fmt.Sprintf("Error: %s", err),
						),
					},
				}
				yield(result)
				return
			}

			translator := toTerraform{cfToTfNameMap: r.cfToTfNameMap}
			schema := request.ResourceSchema
			val, err := translator.FromString(ctx, schema, aws.ToString(obj.Properties))
			if err != nil {
				result = list.ListResult{
					Diagnostics: diag.Diagnostics{
						diag.NewErrorDiagnostic(
							"Error Listing Remote Resources",
							fmt.Sprintf("Error: %s", err),
						),
					},
				}
				yield(result)
				return
			}

			result.Resource = &tfsdk.Resource{
				Schema: schema,
				Raw:    val,
			}

			result.Diagnostics.Append(result.Resource.SetAttribute(ctx, idAttributePath, aws.ToString(description.Identifier))...)
			if result.Diagnostics.HasError() {
				result = list.ListResult{Diagnostics: result.Diagnostics}
				yield(result)
				return
			}

			pi := r.primaryIdentifier.AddAccountID()
			if !r.isGlobal {
				pi = pi.AddRegionID()
			}

			result.Diagnostics.Append(pi.SetIdentity(ctx, r.provider, result.Resource, result.Identity)...)
			if result.Diagnostics.HasError() {
				result = list.ListResult{Diagnostics: result.Diagnostics}
				yield(result)
				return
			}

			result.DisplayName = aws.ToString(description.Identifier)

			if !yield(result) {
				return
			}
		}
	}
}

func (r *genericResource) stream(ctx context.Context, conn *cloudcontrol.Client) iter.Seq2[cctypes.ResourceDescription, error] {
	return tfcloudcontrol.StreamResourcesByTypeName(ctx, conn, r.provider.RoleARN(ctx), r.cfTypeName)
}
