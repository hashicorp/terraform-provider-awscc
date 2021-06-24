package provider

import (
	"context"

	tfsdk "github.com/hashicorp/terraform-plugin-framework"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tflog "github.com/hashicorp/terraform-plugin-log"
)

// Implements tfsdk.ResourceType.
type genericResourceType struct {
	cfTypeName string        // CloudFormation type name for the resource type
	tfSchema   schema.Schema // Terraform schema for the resource type
	tfTypeName string        // Terraform type name for resource type
}

// NewGenericResourceType returns a new ResourceType representing the specified CloudFormation type.
// It's public as it's called from generated code.
func NewGenericResourceType(cfTypeName, tfTypeName string, tfSchema schema.Schema) tfsdk.ResourceType {
	return &genericResourceType{
		cfTypeName: cfTypeName,
		tfSchema:   tfSchema,
		tfTypeName: tfTypeName,
	}
}

func (rt *genericResourceType) GetSchema(ctx context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	tflog.Trace(ctx, "GenericResourceType.GetSchema(%s/%s)", rt.cfTypeName, rt.tfTypeName)

	return rt.tfSchema, nil
}

func (rt *genericResourceType) NewResource(ctx context.Context, provider tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	tflog.Trace(ctx, "GenericResourceType.NewResource(%s/%s)", rt.cfTypeName, rt.tfTypeName)

	return newGenericResource(provider.(*awsProvider), rt), nil
}

// Implements tfsdk.Resource.
type genericResource struct {
	client       *awsClient
	resourceType *genericResourceType
}

func newGenericResource(provider *awsProvider, resourceType *genericResourceType) tfsdk.Resource {
	return &genericResource{
		client:       provider.Client,
		resourceType: resourceType,
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
