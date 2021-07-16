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

// Features of the resource type.
type ResourceTypeFeatures int

const (
	ResourceTypeHasUpdatableAttribute ResourceTypeFeatures = 1 << iota // At least one attribute can be updated.
)

// Implements tfsdk.ResourceType.
type resourceType struct {
	cfTypeName string               // CloudFormation type name for the resource type
	tfSchema   schema.Schema        // Terraform schema for the resource type
	tfTypeName string               // Terraform type name for resource type
	features   ResourceTypeFeatures // Resource type features
}

// NewResourceType returns a new ResourceType representing the specified CloudFormation type.
// It's public as it's called from generated code.
func NewResourceType(cfTypeName, tfTypeName string, tfSchema schema.Schema, features ResourceTypeFeatures) tfsdk.ResourceType {
	return &resourceType{
		features:   features,
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
		response.Diagnostics = append(response.Diagnostics, ClientNotFoundDiag(err))

		return
	}

	log.Printf("[DEBUG] Resource.Create(%s/%s)\nRaw plan: %v", r.resourceType.cfTypeName, r.resourceType.tfTypeName, request.Plan.Raw)

	desiredState, err := GetCloudFormationDesiredState(ctx, request.Plan.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, DesiredStateErrorDiag("Plan", err))

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
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "CreateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("CloudFormation", "CreateResource"))

		return
	}

	output.ProgressEvent, err = waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), 5*time.Minute)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationWaiterErrorDiag("CloudFormation", "CreateResource", err))

		return
	}

	identifier := aws.StringValue(output.ProgressEvent.Identifier)
	description, err := r.describe(ctx, conn, identifier)

	if NotFound(err) {
		response.Diagnostics = append(response.Diagnostics, ResourceNotFoundAfterCreationDiag(err))

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "GetResource", err))

		return
	}

	if description == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("CloudFormation", "GetResource"))

		return
	}

	log.Printf("[DEBUG] ResourceModel: %s", aws.StringValue(description.ResourceModel))

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	// Set the well-known "identifier" attribute.
	err = SetIdentifier(ctx, &response.State, identifier)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierSetErrorDiag(err))

		return
	}

	err = SetUnknownValuesFromCloudFormationResourceModel(ctx, &response.State, aws.StringValue(description.ResourceModel))

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Creation Of Terraform State Unsuccessful",
			Detail:   fmt.Sprintf("Unable to set Terraform State Unknown values from a CloudFormation Resource Model. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		})

		return
	}

	log.Printf("[DEBUG] Resource.Create(%s/%s)\nRaw state: %v", r.resourceType.cfTypeName, r.resourceType.tfTypeName, response.State.Raw)
}

func (r *resource) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	tflog.Debug(ctx, "Resource.Read(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	conn, err := r.clientProvider.CloudFormationClient(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ClientNotFoundDiag(err))

		return
	}

	currentState := &request.State
	schema := &currentState.Schema
	identifier, err := GetIdentifier(ctx, currentState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	description, err := r.describe(ctx, conn, identifier)

	if NotFound(err) {
		response.Diagnostics = append(response.Diagnostics, ResourceNotFoundWarningDiag(err))
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "GetResource", err))

		return
	}

	val, err := GetCloudFormationResourceModelValue(ctx, schema, aws.StringValue(description.ResourceModel))

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Creation Of Terraform State Unsuccessful",
			Detail:   fmt.Sprintf("Unable to create a Terraform State value from a CloudFormation Resource Model. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
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
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierSetErrorDiag(err))

		return
	}
}

func (r *resource) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	tflog.Debug(ctx, "Resource.Update(%s/%s) enter", r.resourceType.cfTypeName, r.resourceType.tfTypeName)

	conn, err := r.clientProvider.CloudFormationClient(ctx)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ClientNotFoundDiag(err))

		return
	}

	currentState := &request.State
	identifier, err := GetIdentifier(ctx, currentState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	currentDesiredState, err := GetCloudFormationDesiredState(ctx, currentState.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, DesiredStateErrorDiag("Prior State", err))

		return
	}

	plannedDesiredState, err := GetCloudFormationDesiredState(ctx, request.Plan.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, DesiredStateErrorDiag("Plan", err))

		return
	}

	patchOperations, err := cfjsonpatch.PatchOperations(currentDesiredState, plannedDesiredState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Creation Of JSON Patch Unsuccessful",
			Detail:   fmt.Sprintf("Unable to create a JSON Patch for resource update. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
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
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "UpdateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("CloudFormation", "UpdateResource"))

		return
	}

	output.ProgressEvent, err = waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), 5*time.Minute)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationWaiterErrorDiag("CloudFormation", "UpdateResource", err))

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
		response.Diagnostics = append(response.Diagnostics, ClientNotFoundDiag(err))

		return
	}

	identifier, err := GetIdentifier(ctx, &request.State)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	input := &cloudformation.DeleteResourceInput{
		ClientToken: aws.String(UniqueId()),
		Identifier:  aws.String(identifier),
		TypeName:    aws.String(r.resourceType.cfTypeName),
	}

	output, err := conn.DeleteResourceWithContext(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "DeleteResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("CloudFormation", "DeleteResource"))

		return
	}

	progressEvent, err := waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), 5*time.Minute)

	if progressEvent != nil && aws.StringValue(progressEvent.ErrorCode) == cloudformation.HandlerErrorCodeNotFound {
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationWaiterErrorDiag("CloudFormation", "DeleteResource", err))

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
