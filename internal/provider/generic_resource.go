package provider

import (
	"context"

	tfsdk "github.com/hashicorp/terraform-plugin-framework"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// Implements tfsdk.ResourceType.
type genericResourceType struct {
	schema schema.Schema
}

func NewGenericResourceType(schema schema.Schema) tfsdk.ResourceType {
	return &genericResourceType{
		schema: schema,
	}
}

func (rt *genericResourceType) GetSchema(_ context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return rt.schema, nil
}

func (rt *genericResourceType) NewResource(provider tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	return newGenericResource(provider.(*awsProvider), rt), nil
}

// Implements tfsdk.Resource.
type genericResource struct {
	client *awsClient
	schema schema.Schema
}

func newGenericResource(provider *awsProvider, resourceType *genericResourceType) tfsdk.Resource {
	return &genericResource{
		client: provider.Client,
		schema: resourceType.schema,
	}
}

func (r *genericResource) Create(ctx context.Context, input *tfsdk.CreateResourceRequest, output *tfsdk.CreateResourceResponse) {
}

func (r *genericResource) Read(ctx context.Context, input *tfsdk.ReadResourceRequest, output *tfsdk.ReadResourceResponse) {
}

func (r *genericResource) Update(ctx context.Context, input *tfsdk.UpdateResourceRequest, output *tfsdk.UpdateResourceResponse) {
}

func (r *genericResource) Delete(ctx context.Context, input *tfsdk.DeleteResourceRequest, output *tfsdk.DeleteResourceResponse) {
}
