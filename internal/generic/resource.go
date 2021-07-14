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
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation/cfjsonpatch"
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

	desiredState, err := GetCloudFormationDesiredState(ctx, request.Plan.Raw)

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

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	// Set the well-known "identifier" attribute.
	err = SetIdentifier(ctx, &response.State, identifier)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting identifier",
			Detail:   fmt.Sprintf("Error setting resource identifier in state.\n%s\n", err),
		})

		return
	}

	err = SetUnknownValuesFromCloudFormationResourceModel(ctx, &response.State, aws.StringValue(description.ResourceModel))

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting CloudFormation resource model",
			Detail:   fmt.Sprintf("Error setting AWS CloudFormation resource model.\n%s\n", err),
		})

		return
	}

	log.Printf("[DEBUG] Resource.Create(%s/%s)\nRaw state: %v", r.resourceType.cfTypeName, r.resourceType.tfTypeName, response.State.Raw)
}

func (r *resource) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	tflog.Debug(ctx, "Resource.Read(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	conn, err := r.clientProvider.CloudFormationClient(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting CloudFormation client",
			Detail:   fmt.Sprintf("Error getting AWS CloudFormation client.\n%s\n", err),
		})

		return
	}

	currentState := &request.State
	schema := &currentState.Schema
	identifier, err := GetIdentifier(ctx, currentState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting identifier",
			Detail:   fmt.Sprintf("Error getting resource identifier from state.\n%s\n", err),
		})

		return
	}

	description, err := r.describe(ctx, conn, identifier)

	if NotFound(err) {
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error describing CloudFormation resource",
			Detail:   fmt.Sprintf("Error describing AWS CloudFormation resource.\n%s\n", err),
		})

		return
	}

	val, err := GetCloudFormationResourceModelValue(ctx, schema, aws.StringValue(description.ResourceModel))

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting Value from CloudFormation ResourceModel",
			Detail:   fmt.Sprintf("Error getting Terraform Value from AWS CloudFormation ResourceModel.\n%s\n", err),
		})

		return
	}

	// TODO
	// TODO Consider write-only values. They can only be in the current state.
	// TODO

	response.State = tfsdk.State{
		Schema: *schema,
		Raw:    val,
	}

	err = SetIdentifier(ctx, &response.State, identifier)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error setting identifier",
			Detail:   fmt.Sprintf("Error setting resource identifier in state.\n%s\n", err),
		})

		return
	}
}

func (r *resource) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	tflog.Debug(ctx, "Resource.Update(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	conn, err := r.clientProvider.CloudFormationClient(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting CloudFormation client",
			Detail:   fmt.Sprintf("Error getting AWS CloudFormation client.\n%s\n", err),
		})

		return
	}

	currentState := &request.State
	identifier, err := GetIdentifier(ctx, currentState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting identifier",
			Detail:   fmt.Sprintf("Error getting resource identifier from state.\n%s\n", err),
		})

		return
	}

	oldDesiredState, err := GetCloudFormationDesiredState(ctx, currentState.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting CloudFormation old desired state",
			Detail:   fmt.Sprintf("Error getting AWS CloudFormation old desired state.\n%s\n", err),
		})

		return
	}

	newDesiredState, err := GetCloudFormationDesiredState(ctx, request.Plan.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error getting CloudFormation new desired state",
			Detail:   fmt.Sprintf("Error getting AWS CloudFormation new desired state.\n%s\n", err),
		})

		return
	}

	patchOperations, err := cfjsonpatch.PatchOperations(oldDesiredState, newDesiredState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "JSON Patch Creation Unsuccessful",
			Detail:   fmt.Sprintf("Creating JSON Patch failed.\n%s\n", err),
		})

		return
	}

	input := &cloudformation.UpdateResourceInput{
		ClientToken:     aws.String(UniqueId()),
		Identifier:      aws.String(identifier),
		PatchOperations: patchOperations,
		TypeName:        aws.String(r.resourceType.cfTypeName),
	}

	output, err := conn.UpdateResourceWithContext(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error updating CloudFormation resource",
			Detail:   fmt.Sprintf("Error updating AWS CloudFormation resource.\n%s\n", err),
		})

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error updating CloudFormation resource",
			Detail:   "Error updating AWS CloudFormation resource.\nEmpty response\n",
		})

		return
	}

	output.ProgressEvent, err = waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), 5*time.Minute)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error waiting for CloudFormation resource update",
			Detail:   fmt.Sprintf("Error waiting for AWS CloudFormation resource update.\n%s\n", err),
		})

		return
	}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	// On Update there should be nothing unknown in the planned state...
	response.State.Raw = request.Plan.Raw
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

	identifier, err := GetIdentifier(ctx, &request.State)

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
func (r *resource) describe(ctx context.Context, conn *cloudformation.CloudFormation, identifier string) (*cloudformation.ResourceDescription, error) {
	input := &cloudformation.GetResourceInput{
		Identifier: aws.String(identifier),
		TypeName:   aws.String(r.resourceType.cfTypeName),
	}

	output, err := conn.GetResourceWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, cloudformation.ErrCodeResourceNotFoundException) {
		return nil, &NotFoundError{LastError: err}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.ResourceDescription == nil {
		return nil, &NotFoundError{Message: "Empty result"}
	}

	return output.ResourceDescription, nil
}
