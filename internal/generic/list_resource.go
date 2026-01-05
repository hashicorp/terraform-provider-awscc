// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"errors"
	"fmt"
	"iter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
)

var _ list.ListResource = &genericResource{}

// NewListResource returns a new generic ListResource
func NewListResource(resource func(context.Context) (resource.Resource, error)) func(context.Context) (list.ListResource, error) {
	return func(ctx context.Context) (list.ListResource, error) {
		res, err := resource(ctx)
		if err != nil {
			return nil, err
		}

		if v, ok := res.(*genericResource); ok {
			return v, nil
		}

		return nil, errors.New("list resource does not implement generic resource")
	}
}

func (r *genericResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, response *list.ListResourceSchemaResponse) {
	response.Schema = listschema.Schema{
		MarkdownDescription: "List all resources for `" + r.cfTypeName + "` resource type.",
		Attributes:          map[string]listschema.Attribute{},
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

			result.Diagnostics.Append(r.populateIdentityTypes(ctx, request.ResourceIdentitySchema.Type(), result)...)
			if result.Diagnostics.HasError() {
				result = list.ListResult{Diagnostics: result.Diagnostics}
				yield(result)
				return
			}

			translator := toTerraform{cfToTfNameMap: r.cfToTfNameMap}
			schema := request.ResourceSchema
			val, err := translator.FromString(ctx, schema, aws.ToString(description.Properties))
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

			pi := r.primaryIdentifier.AppendDefaults(r.isGlobal)

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
	return func(yield func(cctypes.ResourceDescription, error) bool) {
		for description, err := range tfcloudcontrol.StreamResourcesByTypeName(ctx, conn, r.provider.RoleARN(ctx), r.cfTypeName) {
			if err != nil {
				yield(cctypes.ResourceDescription{}, err)
				return
			}

			output, err := r.describe(ctx, conn, aws.ToString(description.Identifier))
			if err != nil {
				yield(cctypes.ResourceDescription{}, err)
				return
			}

			if !yield(*output, nil) {
				return
			}
		}
	}
}

func (r *genericResource) populateIdentityTypes(ctx context.Context, schemaType attr.Type, result list.ListResult) diag.Diagnostics {
	var diags diag.Diagnostics
	obj, d := newEmptyObject(schemaType)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	diags.Append(result.Identity.Set(ctx, obj)...)
	if diags.HasError() {
		return diags
	}

	return diags
}

func newEmptyObject(typ attr.Type) (obj basetypes.ObjectValue, diags diag.Diagnostics) {
	i, ok := typ.(attr.TypeWithAttributeTypes)
	if !ok {
		diags.AddError(
			"Internal Error",
			"An unexpected error occurred. "+
				"This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+
				fmt.Sprintf("Expected value type to implement attr.TypeWithAttributeTypes, got: %T", typ),
		)
		return
	}

	attrTypes := i.AttributeTypes()
	attrValues := make(map[string]attr.Value, len(attrTypes))
	for attrName := range attrTypes {
		attrValues[attrName] = types.StringNull()
	}
	obj, d := basetypes.NewObjectValue(attrTypes, attrValues)
	diags.Append(d...)
	if d.HasError() {
		return basetypes.ObjectValue{}, diags
	}

	return obj, diags
}
