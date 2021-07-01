package generic

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation/waiter"
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
	tflog.Debug(ctx, "ResourceType.GetSchema(%s/%s) enter", rt.cfTypeName, rt.tfTypeName)

	return rt.tfSchema, nil
}

func (rt *resourceType) NewResource(ctx context.Context, provider tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	tflog.Debug(ctx, "ResourceType.NewResource(%s/%s) enter", rt.cfTypeName, rt.tfTypeName)

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

func (r *resource) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	tflog.Debug(ctx, "Resource.Create(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	conn, err := r.clientProvider.CloudFormationClient(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting CloudFormation client",
			Detail:   fmt.Sprintf("Error getting AWS CloudFormation client.\n%s\n", err),
		})

		return
	}

	log.Printf("[DEBUG] Resource.Create(%s/%s)\nRaw plan: %v", r.resourceType.cfTypeName, r.resourceType.tfTypeName, request.Plan.Raw)

	plan := &Plan{inner: &request.Plan}
	desiredState, err := plan.GetCloudFormationDesiredState(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting CloudFormation desired state",
			Detail:   fmt.Sprintf("Error getting AWS CloudFormation desired state.\n%s\n", err),
		})

		return
	}

	log.Printf("[DEBUG] CloudFormation desired state: %s", desiredState)

	input := &cloudformation.CreateResourceInput{
		ClientToken:  aws.String(UniqueId()),
		DesiredState: aws.String(desiredState),
		TypeName:     aws.String(r.resourceType.cfTypeName),
	}

	output, err := conn.CreateResourceWithContext(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error creating CloudFormation resource",
			Detail:   fmt.Sprintf("Error creating AWS CloudFormation resource.\n%s\n", err),
		})

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error creating CloudFormation resource",
			Detail:   "Error creating AWS CloudFormation resource.\nEmpty response\n",
		})

		return
	}

	output.ProgressEvent, err = waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), 5*time.Minute)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error waiting for CloudFormation resource creation",
			Detail:   fmt.Sprintf("Error waiting for AWS CloudFormation resource creation.\n%s\n", err),
		})

		return
	}

	response.State.Raw = request.Plan.Raw

	identifier := aws.StringValue(output.ProgressEvent.Identifier)
	description, err := r.describe(ctx, conn, identifier)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error describing CloudFormation resource",
			Detail:   fmt.Sprintf("Error describing AWS CloudFormation resource.\n%s\n", err),
		})

		return
	}

	if description == nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error describing CloudFormation resource",
			Detail:   "Error describing AWS CloudFormation resource.\nEmpty response\n",
		})

		return
	}

	log.Printf("[DEBUG] ResourceModel: %s", aws.StringValue(description.ResourceModel))

	// TODO
	// TODO Populate rest of State.
	// TODO
	state := &State{inner: &response.State}

	err = state.SetCloudFormationResourceModel(ctx, aws.StringValue(description.ResourceModel))

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting CloudFormation resource model",
			Detail:   fmt.Sprintf("Error setting AWS CloudFormation resource model.\n%s\n", err),
		})

		return
	}

	err = state.SetIdentifier(ctx, identifier)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting identifier",
			Detail:   fmt.Sprintf("Error setting resource identifier in state.\n%s\n", err),
		})

		return
	}

	log.Printf("[DEBUG] Resource.Create(%s/%s)\nRaw state: %v", r.resourceType.cfTypeName, r.resourceType.tfTypeName, response.State.Raw)
}

func (r *resource) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	tflog.Debug(ctx, "Resource.Read(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Unimplemented Resource.Read",
	})
}

func (r *resource) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	tflog.Debug(ctx, "Resource.Update(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Unimplemented Resource.Update",
	})
}

func (r *resource) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	tflog.Debug(ctx, "Resource.Delete(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	conn, err := r.clientProvider.CloudFormationClient(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting CloudFormation client",
			Detail:   fmt.Sprintf("Error getting AWS CloudFormation client.\n%s\n", err),
		})

		return
	}

	log.Printf("[DEBUG] Resource.Delete(%s/%s)\nRaw state: %v", r.resourceType.cfTypeName, r.resourceType.tfTypeName, request.State.Raw)

	state := &State{inner: &request.State}
	identifier, err := state.GetIdentifier(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting identifier",
			Detail:   fmt.Sprintf("Error getting resource identifier from state.\n%s\n", err),
		})

		return
	}

	input := &cloudformation.DeleteResourceInput{
		ClientToken: aws.String(UniqueId()),
		Identifier:  aws.String(identifier),
		TypeName:    aws.String(r.resourceType.cfTypeName),
	}

	output, err := conn.DeleteResourceWithContext(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error deleting CloudFormation resource",
			Detail:   fmt.Sprintf("Error deleting AWS CloudFormation resource.\n%s\n", err),
		})

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error deleting CloudFormation resource",
			Detail:   "Error deleting AWS CloudFormation resource.\nEmpty response\n",
		})

		return
	}

	progressEvent, err := waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), 5*time.Minute)

	if progressEvent != nil && aws.StringValue(progressEvent.ErrorCode) == cloudformation.HandlerErrorCodeNotFound {
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error waiting for CloudFormation resource deletion",
			Detail:   fmt.Sprintf("Error waiting for AWS CloudFormation resource deletion.\n%s\n", err),
		})

		return
	}

	response.State.RemoveResource(ctx)
}

// describe returns the live state of the specified resource.
// TODO Return NotFoundError when not found.
func (r *resource) describe(ctx context.Context, conn *cloudformation.CloudFormation, identifier string) (*cloudformation.ResourceDescription, error) {
	input := &cloudformation.GetResourceInput{
		Identifier: aws.String(identifier),
		TypeName:   aws.String(r.resourceType.cfTypeName),
	}

	output, err := conn.GetResourceWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, cloudformation.ErrCodeResourceNotFoundException) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.ResourceDescription == nil {
		return nil, nil
	}

	return output.ResourceDescription, nil
}
