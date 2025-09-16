// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"iter"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
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
	var diags diag.Diagnostics

	diags.Append(diag.NewWarningDiagnostic(
		"Not Implemented",
		"This list resource is not implemented."))

	stream.Results = list.ListResultsStreamDiagnostics(diags)
}

func (r *genericResource) stream(ctx context.Context, conn *cloudcontrol.Client) iter.Seq2[cctypes.ResourceDescription, error] {
	return tfcloudcontrol.StreamResourcesByTypeName(ctx, conn, r.provider.RoleARN(ctx), r.cfTypeName)
}
