package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/naming"
	tfcloudformation "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/tfresource"
	"github.com/mattbaird/jsonpatch"
)

// resourceTypeOptions are a discrete set of options that are valid for the resource type.
type resourceTypeOptions struct {
	cfTypeName              string                   // CloudFormation type name for the resource type
	tfSchema                schema.Schema            // Terraform schema for the resource type
	tfTypeName              string                   // Terraform type name for resource type
	isImmutableType         bool                     // Resources cannot be updated and must be recreated
	identifierAttributePath *tftypes.AttributePath   // Path to the resource's primary identifier attribute
	writeOnlyAttributePaths []*tftypes.AttributePath // Paths to any write-only attributes
}

// ResourceTypeOptionsFunc is a type alias for a resource type functional option.
type ResourceTypeOptionsFunc func(*resourceTypeOptions) error

// WithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name.
// If multiple WithCloudFormationTypeName calls are made, the last call overrides
// the previous calls' values.
func WithCloudFormationTypeName(v string) ResourceTypeOptionsFunc {
	return func(o *resourceTypeOptions) error {
		o.cfTypeName = v

		return nil
	}
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema.
// If multiple WithTerraformSchema calls are made, the last call overrides
// the previous calls' values.
func WithTerraformSchema(v schema.Schema) ResourceTypeOptionsFunc {
	return func(o *resourceTypeOptions) error {
		o.tfSchema = v

		return nil
	}
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name.
// If multiple WithTerraformTypeName calls are made, the last call overrides
// the previous calls' values.
func WithTerraformTypeName(v string) ResourceTypeOptionsFunc {
	return func(o *resourceTypeOptions) error {
		o.tfTypeName = v

		return nil
	}
}

// IsImmutableType is a helper function to construct functional options
// that set a resource type's immutability flag.
// If multiple IsImmutableType calls are made, the last call overrides
// the previous calls' values.
func IsImmutableType(v bool) ResourceTypeOptionsFunc {
	return func(o *resourceTypeOptions) error {
		o.isImmutableType = v

		return nil
	}
}

// WithPrimaryIdentifierPath is a helper function to construct functional options
// that set a resource type's primary identifier path (JSON Pointer).
// If multiple WithPrimaryIdentifierPath calls are made, the last call overrides
// the previous calls' values.
func WithPrimaryIdentifierPath(v string) ResourceTypeOptionsFunc {
	return func(o *resourceTypeOptions) error {
		identifierAttributePath, err := propertyPathToAttributePath(v)

		if err != nil {
			return fmt.Errorf("error creating identifier attribute path (%s): %w", v, err)
		}

		o.identifierAttributePath = identifierAttributePath

		return nil
	}
}

// WithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths (JSON Pointer).
// If multiple WithWriteOnlyPropertyPaths calls are made, the last call overrides
// the previous calls' values.
func WithWriteOnlyPropertyPaths(v []string) ResourceTypeOptionsFunc {
	return func(o *resourceTypeOptions) error {
		writeOnlyAttributePaths := make([]*tftypes.AttributePath, 0)

		for _, writeOnlyPropertyPath := range v {
			writeOnlyAttributePath, err := propertyPathToAttributePath(writeOnlyPropertyPath)

			if err != nil {
				return fmt.Errorf("error creating write-only attribute path (%s): %w", writeOnlyPropertyPath, err)
			}

			writeOnlyAttributePaths = append(writeOnlyAttributePaths, writeOnlyAttributePath)
		}

		o.writeOnlyAttributePaths = writeOnlyAttributePaths

		return nil
	}
}

// ResourceTypeOptions is a type alias for a slice of resource type functional options.
type ResourceTypeOptions []ResourceTypeOptionsFunc

// WithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithCloudFormationTypeName(v string) ResourceTypeOptions {
	return append(opts, WithCloudFormationTypeName(v))
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithTerraformSchema(v schema.Schema) ResourceTypeOptions {
	return append(opts, WithTerraformSchema(v))
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithTerraformTypeName(v string) ResourceTypeOptions {
	return append(opts, WithTerraformTypeName(v))
}

// IsImmutableType is a helper function to construct functional options
// that set a resource type's Terraform immutability flag, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) IsImmutableType(v bool) ResourceTypeOptions {
	return append(opts, IsImmutableType(v))
}

// WithPrimaryIdentifierPath is a helper function to construct functional options
// that set a resource type's primary identifier path, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithPrimaryIdentifierPath(v string) ResourceTypeOptions {
	return append(opts, WithPrimaryIdentifierPath(v))
}

// WithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceTypeOptions) WithWriteOnlyPropertyPaths(v []string) ResourceTypeOptions {
	return append(opts, WithWriteOnlyPropertyPaths(v))
}

// resourceType implements tfsdk.ResourceType.
type resourceType struct {
	cfTypeName              string                   // CloudFormation type name for the resource type
	tfSchema                schema.Schema            // Terraform schema for the resource type
	tfTypeName              string                   // Terraform type name for resource type
	isImmutableType         bool                     // Resources cannot be updated and must be recreated
	identifierAttributePath *tftypes.AttributePath   // Path to the resource's primary identifier attribute
	writeOnlyAttributePaths []*tftypes.AttributePath // Paths to any write-only attributes
}

// NewResourceType returns a new ResourceType from the specified varidaic list of functional options.
// It's public as it's called from generated code.
func NewResourceType(ctx context.Context, optFns ...ResourceTypeOptionsFunc) (tfsdk.ResourceType, error) {
	var options resourceTypeOptions

	for _, optFn := range optFns {
		err := optFn(&options)

		if err != nil {
			return nil, err
		}
	}

	if options.cfTypeName == "" {
		return nil, fmt.Errorf("no CloudFormation type name specified")
	}
	if options.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}
	if options.identifierAttributePath == nil {
		return nil, fmt.Errorf("no primary identifier path specified")
	}

	return &resourceType{
		cfTypeName:              options.cfTypeName,
		tfSchema:                options.tfSchema,
		tfTypeName:              options.tfTypeName,
		isImmutableType:         options.isImmutableType,
		identifierAttributePath: options.identifierAttributePath,
		writeOnlyAttributePaths: options.writeOnlyAttributePaths,
	}, nil
}

func (rt *resourceType) GetSchema(ctx context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return rt.tfSchema, nil
}

func (rt *resourceType) NewResource(ctx context.Context, provider tfsdk.Provider) (tfsdk.Resource, []*tfprotov6.Diagnostic) {
	return newGenericResource(provider, rt), nil
}

// Implements tfsdk.Resource.
type resource struct {
	provider     tfcloudformation.Provider
	resourceType *resourceType
}

func newGenericResource(provider tfsdk.Provider, resourceType *resourceType) tfsdk.Resource {
	return &resource{
		provider:     provider.(tfcloudformation.Provider),
		resourceType: resourceType,
	}
}

var (
	// Path to the "id" attribute required for acceptance testing.
	idAttributePath = tftypes.NewAttributePath().WithAttributeName("id")
)

func (r *resource) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := r.resourceType.cfTypeName
	tfTypeName := r.resourceType.tfTypeName

	tflog.Debug(ctx, "Resource.Create enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := r.provider.CloudFormationClient(ctx)

	tflog.Debug(ctx, "Request.Plan.Raw", "value", hclog.Fmt("%v", request.Plan.Raw))

	desiredState, err := GetCloudFormationDesiredState(ctx, request.Plan.Raw)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, DesiredStateErrorDiag("Plan", err))

		return
	}

	tflog.Debug(ctx, "CloudFormation DesiredState", "value", desiredState)

	input := &cloudformation.CreateResourceInput{
		ClientToken:  aws.String(tfresource.UniqueId()),
		DesiredState: aws.String(desiredState),
		TypeName:     aws.String(cfTypeName),
	}

	if roleARN := r.provider.RoleARN(ctx); roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.CreateResource(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "CreateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("CloudFormation", "CreateResource"))

		return
	}

	// TODO
	// TODO How long to wait for?
	// TODO
	maxWaitTime := 5 * time.Minute

	identifier, err := tfcloudformation.WaitForResourceRequestSuccess(ctx, conn, aws.ToString(output.ProgressEvent.RequestToken), maxWaitTime)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationWaiterErrorDiag("CloudFormation", "CreateResource", err))

		return
	}

	description, err := r.describe(ctx, conn, identifier)

	if tfresource.NotFound(err) {
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

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	// Set the "id" attribute required for acceptance testing.
	response.State.SetAttribute(ctx, idAttributePath, identifier)

	err = SetUnknownValuesFromCloudFormationResourceModel(ctx, &response.State, aws.ToString(description.ResourceModel))

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Creation Of Terraform State Unsuccessful",
			Detail:   fmt.Sprintf("Unable to set Terraform State Unknown values from a CloudFormation Resource Model. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		})

		return
	}

	tflog.Debug(ctx, "Response.State.Raw", "value", hclog.Fmt("%v", response.State.Raw))

	tflog.Debug(ctx, "Resource.Create exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

func (r *resource) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := r.resourceType.cfTypeName
	tfTypeName := r.resourceType.tfTypeName

	tflog.Debug(ctx, "Resource.Read enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := r.provider.CloudFormationClient(ctx)

	currentState := &request.State
	identifier, err := r.getIdentifier(ctx, currentState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	description, err := r.describe(ctx, conn, identifier)

	if tfresource.NotFound(err) {
		response.Diagnostics = append(response.Diagnostics, ResourceNotFoundWarningDiag(err))
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "GetResource", err))

		return
	}

	schema := &currentState.Schema
	val, err := GetCloudFormationResourceModelValue(ctx, schema, aws.ToString(description.ResourceModel))

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

	// Set the "id" attribute required for acceptance testing.
	response.State.SetAttribute(ctx, idAttributePath, identifier)

	tflog.Debug(ctx, "Resource.Read exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

func (r *resource) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := r.resourceType.cfTypeName
	tfTypeName := r.resourceType.tfTypeName

	tflog.Debug(ctx, "Resource.Update enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := r.provider.CloudFormationClient(ctx)

	currentState := &request.State
	identifier, err := r.getIdentifier(ctx, currentState)

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

	patchDocument, err := patchDocument(currentDesiredState, plannedDesiredState)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Creation Of JSON Patch Unsuccessful",
			Detail:   fmt.Sprintf("Unable to create a JSON Patch for resource update. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		})

		return
	}

	tflog.Debug(ctx, "CloudFormation PatchDocument", "value", patchDocument)

	input := &cloudformation.UpdateResourceInput{
		ClientToken:   aws.String(tfresource.UniqueId()),
		Identifier:    aws.String(identifier),
		PatchDocument: aws.String(patchDocument),
		TypeName:      aws.String(cfTypeName),
	}

	if roleARN := r.provider.RoleARN(ctx); roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.UpdateResource(ctx, input)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "UpdateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationEmptyResultDiag("CloudFormation", "UpdateResource"))

		return
	}

	// TODO
	// TODO How long to wait for?
	// TODO
	maxWaitTime := 5 * time.Minute

	_, err = tfcloudformation.WaitForResourceRequestSuccess(ctx, conn, aws.ToString(output.ProgressEvent.RequestToken), maxWaitTime)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationWaiterErrorDiag("CloudFormation", "UpdateResource", err))

		return
	}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	// On Update there should be nothing unknown in the planned state...
	response.State.Raw = request.Plan.Raw

	tflog.Debug(ctx, "Resource.Update exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

func (r *resource) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := r.resourceType.cfTypeName
	tfTypeName := r.resourceType.tfTypeName

	tflog.Debug(ctx, "Resource.Delete enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := r.provider.CloudFormationClient(ctx)

	identifier, err := r.getIdentifier(ctx, &request.State)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	err = tfcloudformation.DeleteResource(ctx, conn, r.provider.RoleARN(ctx), cfTypeName, identifier)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "DeleteResource", err))

		return
	}

	response.State.RemoveResource(ctx)

	tflog.Debug(ctx, "Resource.Delete exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

// describe returns the live state of the specified resource.
func (r *resource) describe(ctx context.Context, conn *cloudformation.Client, identifier string) (*cftypes.ResourceDescription, error) {
	return tfcloudformation.FindResourceByTypeNameAndID(ctx, conn, r.provider.RoleARN(ctx), r.resourceType.cfTypeName, identifier)
}

// getIdentifier returns the resource's primary identifier value from State.
func (r *resource) getIdentifier(ctx context.Context, state *tfsdk.State) (string, error) {
	val, err := state.GetAttribute(ctx, r.resourceType.identifierAttributePath)

	if err != nil {
		return "", err
	}

	if val, ok := val.(types.String); ok {
		return val.Value, nil
	}

	return "", fmt.Errorf("invalid identifier type %T", val)
}

// patchDocument returns a JSON Patch document describing the difference between `old` and `new`.
func patchDocument(old, new string) (string, error) {
	patch, err := jsonpatch.CreatePatch([]byte(old), []byte(new))

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(patch)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// propertyPathToAttributePath returns the AttributePath for the specified JSON Pointer property path.
func propertyPathToAttributePath(propertyPath string) (*tftypes.AttributePath, error) {
	segments := strings.Split(propertyPath, "/")

	if got, expected := len(segments), 3; got < expected {
		return nil, fmt.Errorf("expected at least %d property path segments, got: %d", expected, got)
	}

	if got, expected := segments[0], ""; got != expected {
		return nil, fmt.Errorf("expected %q for the initial property path segment, got: %q", expected, got)
	}

	if got, expected := segments[1], "properties"; got != expected {
		return nil, fmt.Errorf("expected %q for the second property path segment, got: %q", expected, got)
	}

	attributePath := tftypes.NewAttributePath()

	for _, segment := range segments[2:] {
		switch segment {
		case "", "*":
			return nil, fmt.Errorf("invalid property path segment: %q", segment)

		default:
			attributePath = attributePath.WithAttributeName(naming.CloudFormationPropertyToTerraformAttribute(segment))
		}
	}

	return attributePath, nil
}
