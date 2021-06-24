package generic

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	tfsdk "github.com/hashicorp/terraform-plugin-framework"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tflog "github.com/hashicorp/terraform-plugin-log"
)

// Implements tfsdk.ResourceType.
type resourceType struct {
	cfTypeName string        // CloudFormation type name for the resource type
	tfSchema   schema.Schema // Terraform schema for the resource type
	tfTypeName string        // Terraform type name for resource type
}

// NewResourceType returns a new ResourceType representing the specified CloudFormation type.
// It's public as it's called from generated code.
func NewResourceType(cfTypeName, tfTypeName string, tfSchema schema.Schema) tfsdk.ResourceType {
	return &resourceType{
		cfTypeName: cfTypeName,
		tfSchema:   tfSchema,
		tfTypeName: tfTypeName,
	}
}

func (rt *resourceType) GetSchema(ctx context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	tflog.Trace(ctx, "ResourceType.GetSchema(%s/%s) enter", rt.cfTypeName, rt.tfTypeName)

	return rt.tfSchema, nil
}

func (rt *resourceType) NewResource(ctx context.Context, provider tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	tflog.Trace(ctx, "ResourceType.NewResource(%s/%s) enter", rt.cfTypeName, rt.tfTypeName)

	return newGenericResource(provider, rt), nil
}

// CloudFormationClientProvider is the interface implemented by AWS CloudFormation client providers.
type CloudFormationClientProvider interface {
	// CloudFormationClient returns an AWS CloudFormation client.
	CloudFormationClient(context.Context) (*cloudformation.CloudFormation, error)
}

// Implements tfsdk.Resource.
type resource struct {
	clientProvider CloudFormationClientProvider
	resourceType   *resourceType
}

func newGenericResource(provider tfsdk.Provider, resourceType *resourceType) tfsdk.Resource {
	return &resource{
		clientProvider: provider.(CloudFormationClientProvider),
		resourceType:   resourceType,
	}
}

func (r *resource) Create(ctx context.Context, input *tfsdk.CreateResourceRequest, output *tfsdk.CreateResourceResponse) {
	tflog.Trace(ctx, "Resource.Create(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)
}

func (r *resource) Read(ctx context.Context, input *tfsdk.ReadResourceRequest, output *tfsdk.ReadResourceResponse) {
	tflog.Trace(ctx, "Resource.Read(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)
}

func (r *resource) Update(ctx context.Context, input *tfsdk.UpdateResourceRequest, output *tfsdk.UpdateResourceResponse) {
	tflog.Trace(ctx, "Resource.Update(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)
}

func (r *resource) Delete(ctx context.Context, input *tfsdk.DeleteResourceRequest, output *tfsdk.DeleteResourceResponse) {
	tflog.Trace(ctx, "Resource.Delete(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)
}
