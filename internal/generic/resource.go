// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// ResourceOptionsFunc is a type alias for a resource type functional option.
type ResourceOptionsFunc func(*genericResource) error

// resourceWithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name maps.
// If multiple resourceWithAttributeNameMap calls are made, the last call overrides
// the previous calls' values.
func resourceWithAttributeNameMap(v map[string]string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if _, ok := v["id"]; !ok {
			// Synthesize a mapping for the reserved top-level "id" attribute.
			v["id"] = "ID"
		}

		cfToTfNameMap := make(map[string]string, len(v))

		for tfName, cfName := range v {
			_, ok := cfToTfNameMap[cfName]
			if ok {
				return fmt.Errorf("duplicate attribute name mapping for CloudFormation property %s", cfName)
			}
			cfToTfNameMap[cfName] = tfName
		}

		o.tfToCfNameMap = v
		o.cfToTfNameMap = cfToTfNameMap

		return nil
	}
}

// resourceWithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name.
// If multiple resourceWithCloudFormationTypeName calls are made, the last call overrides
// the previous calls' values.
func resourceWithCloudFormationTypeName(v string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.cfTypeName = v

		return nil
	}
}

// resourceWithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema.
// If multiple resourceWithTerraformSchema calls are made, the last call overrides
// the previous calls' values.
func resourceWithTerraformSchema(v schema.Schema) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.tfSchema = v

		return nil
	}
}

// resourceWithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name.
// If multiple resourceWithTerraformTypeName calls are made, the last call overrides
// the previous calls' values.
func resourceWithTerraformTypeName(v string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.tfTypeName = v

		return nil
	}
}

// resourceIsImmutableType is a helper function to construct functional options
// that set a resource type's immutability flag.
// If multiple resourceIsImmutableType calls are made, the last call overrides
// the previous calls' values.
func resourceIsImmutableType(v bool) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.isImmutableType = v

		return nil
	}
}

// resourceWithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths (JSON Pointer).
// If multiple resourceWithWriteOnlyPropertyPaths calls are made, the last call overrides
// the previous calls' values.
func resourceWithWriteOnlyPropertyPaths(v []string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		writeOnlyAttributePaths := make([]*path.Path, 0)

		for _, writeOnlyPropertyPath := range v {
			writeOnlyAttributePath, err := o.propertyPathToAttributePath(writeOnlyPropertyPath)

			if err != nil {
				// return fmt.Errorf("creating write-only attribute path (%s): %w", writeOnlyPropertyPath, err)
				continue
			}

			writeOnlyAttributePaths = append(writeOnlyAttributePaths, writeOnlyAttributePath)
		}

		o.writeOnlyAttributePaths = writeOnlyAttributePaths

		return nil
	}
}

const (
	resourceMaxWaitTimeCreate = 120 * time.Minute
	resourceMaxWaitTimeUpdate = 120 * time.Minute
	resourceMaxWaitTimeDelete = 120 * time.Minute
)

// resourceWithCreateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's create timeout (in minutes).
// If multiple resourceWithCreateTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithCreateTimeoutInMinutes(v int) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if v > 0 {
			o.createTimeout = time.Duration(v) * time.Minute
		} else {
			o.createTimeout = resourceMaxWaitTimeCreate
		}

		return nil
	}
}

// resourceWithUpdateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's update timeout (in minutes).
// If multiple resourceWithUpdateTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithUpdateTimeoutInMinutes(v int) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if v > 0 {
			o.updateTimeout = time.Duration(v) * time.Minute
		} else {
			o.updateTimeout = resourceMaxWaitTimeUpdate
		}

		return nil
	}
}

// resourceWithDeleteTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's delete timeout (in minutes).
// If multiple resourceWithDeleteTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithDeleteTimeoutInMinutes(v int) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if v > 0 {
			o.deleteTimeout = time.Duration(v) * time.Minute
		} else {
			o.deleteTimeout = resourceMaxWaitTimeDelete
		}

		return nil
	}
}

// resourceWithConfigValidators is a helper function to construct functional options
// that set a resource type's config validators.
// If multiple resourceWithConfigValidators calls are made, the last call overrides
// the previous calls' values.
func resourceWithConfigValidators(vs ...resource.ConfigValidator) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.configValidators = vs

		return nil
	}
}

// ResourceOptions is a type alias for a slice of resource type functional options.
type ResourceOptions []ResourceOptionsFunc

// WithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name map, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithAttributeNameMap(v map[string]string) ResourceOptions {
	return append(opts, resourceWithAttributeNameMap(v))
}

// WithCloudFormationTypeName is a helper function to construct functional options
// that set a resource type's CloudFormation type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithCloudFormationTypeName(v string) ResourceOptions {
	return append(opts, resourceWithCloudFormationTypeName(v))
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithTerraformSchema(v schema.Schema) ResourceOptions {
	return append(opts, resourceWithTerraformSchema(v))
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithTerraformTypeName(v string) ResourceOptions {
	return append(opts, resourceWithTerraformTypeName(v))
}

// IsImmutableType is a helper function to construct functional options
// that set a resource type's Terraform immutability flag, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) IsImmutableType(v bool) ResourceOptions {
	return append(opts, resourceIsImmutableType(v))
}

// WithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithWriteOnlyPropertyPaths(v []string) ResourceOptions {
	return append(opts, resourceWithWriteOnlyPropertyPaths(v))
}

// WithCreateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's create timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithCreateTimeoutInMinutes(v int) ResourceOptions {
	return append(opts, resourceWithCreateTimeoutInMinutes(v))
}

// WithUpdateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's update timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithUpdateTimeoutInMinutes(v int) ResourceOptions {
	return append(opts, resourceWithUpdateTimeoutInMinutes(v))
}

// WithDeleteTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's delete timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithDeleteTimeoutInMinutes(v int) ResourceOptions {
	return append(opts, resourceWithDeleteTimeoutInMinutes(v))
}

// WithConfigValidators is a helper function to construct functional options
// that set a resource type's config validators, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithConfigValidators(vs ...resource.ConfigValidator) ResourceOptions {
	return append(opts, resourceWithConfigValidators(vs...))
}

// NewResource returns a new Resource from the specified varidaic list of functional options.
// It's public as it's called from generated code.
func NewResource(_ context.Context, optFns ...ResourceOptionsFunc) (resource.Resource, error) {
	v := &genericResource{}

	for _, optFn := range optFns {
		err := optFn(v)

		if err != nil {
			return nil, err
		}
	}

	if v.cfTypeName == "" {
		return nil, fmt.Errorf("no CloudFormation type name specified")
	}
	if v.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}

	return v, nil
}

// Implements resource.Resource.
type genericResource struct {
	cfTypeName              string                     // CloudFormation type name for the resource type
	tfSchema                schema.Schema              // Terraform schema for the resource type
	tfTypeName              string                     // Terraform type name for resource type
	tfToCfNameMap           map[string]string          // Map of Terraform attribute name to CloudFormation property name
	cfToTfNameMap           map[string]string          // Map of CloudFormation property name to Terraform attribute name
	isImmutableType         bool                       // Resources cannot be updated and must be recreated
	writeOnlyAttributePaths []*path.Path               // Paths to any write-only attributes
	createTimeout           time.Duration              // Maximum wait time for resource creation
	updateTimeout           time.Duration              // Maximum wait time for resource update
	deleteTimeout           time.Duration              // Maximum wait time for resource deletion
	configValidators        []resource.ConfigValidator // Required attributes validators
	provider                tfcloudcontrol.Provider
}

var (
	// Path to the "id" attribute which uniquely (for a specific resource type) identifies the resource.
	// This attribute is required for acceptance testing.
	idAttributePath = path.Root("id")
)

func (r *genericResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = r.tfTypeName
}

func (r *genericResource) Schema(_ context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = r.tfSchema
}

func (r *genericResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) { //nolint:unparam
	if v := request.ProviderData; v != nil {
		r.provider = v.(tfcloudcontrol.Provider)
	}
}

func (r *genericResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Create")

	conn := r.provider.CloudControlAPIClient(ctx)

	tflog.Debug(ctx, "Request.Plan.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", request.Plan.Raw),
	})

	translator := toCloudControl{tfToCfNameMap: r.tfToCfNameMap}
	desiredState, err := translator.AsString(ctx, request.Plan.Schema, request.Plan.Raw)

	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Plan", err))

		return
	}

	tflog.Debug(ctx, "CloudControl DesiredState", map[string]interface{}{
		"value": desiredState,
	})

	input := &cloudcontrol.CreateResourceInput{
		ClientToken:  aws.String(tfresource.UniqueId()),
		DesiredState: aws.String(desiredState),
		TypeName:     aws.String(r.cfTypeName),
	}

	if roleARN := r.provider.RoleARN(ctx); roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.CreateResource(ctx, input)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "CreateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics.Append(ServiceOperationEmptyResultDiag("Cloud Control API", "CreateResource"))

		return
	}

	var progressEvent *cctypes.ProgressEvent
	waiter := cloudcontrol.NewResourceRequestSuccessWaiter(conn, func(o *cloudcontrol.ResourceRequestSuccessWaiterOptions) {
		o.Retryable = tfcloudcontrol.RetryGetResourceRequestStatus(&progressEvent)
	})

	err = waiter.Wait(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: output.ProgressEvent.RequestToken}, r.createTimeout)

	var id string
	if progressEvent != nil {
		id = aws.ToString(progressEvent.Identifier)
	}

	if err != nil {
		response.Diagnostics.Append(ServiceOperationWaiterErrorDiag("Cloud Control API", "CreateResource", err))

		// Save any ID to state so that the resource will be marked as tainted.
		if id != "" {
			if err := r.setId(ctx, id, &response.State); err != nil {
				response.Diagnostics.Append(ResourceIdentifierNotSetDiag(err))
			}
		}

		return
	}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	// Set the "id" attribute.
	if err = r.setId(ctx, id, &response.State); err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotSetDiag(err))

		return
	}

	response.Diagnostics.Append(r.populateUnknownValues(ctx, id, &response.State)...)
	if response.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "Resource.Create")
}

func (r *genericResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Read")

	tflog.Debug(ctx, "Request.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", request.State.Raw),
	})

	conn := r.provider.CloudControlAPIClient(ctx)

	currentState := &request.State
	id, err := r.getId(ctx, currentState)

	if err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotFoundDiag(err))

		return
	}

	description, err := r.describe(ctx, conn, id)

	if tfresource.NotFound(err) {
		response.Diagnostics.Append(ResourceNotFoundWarningDiag(err))
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))

		return
	}

	translator := toTerraform{cfToTfNameMap: r.cfToTfNameMap}
	schema := currentState.Schema
	val, err := translator.FromString(ctx, schema, aws.ToString(description.Properties))

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to create a Terraform State value from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	response.State = tfsdk.State{
		Schema: schema,
		Raw:    val,
	}

	// Copy over any write-only values.
	// They can only be in the current state.
	for _, path := range r.writeOnlyAttributePaths {
		response.Diagnostics.Append(copyStateValueAtPath(ctx, &response.State, &request.State, *path)...)
		if response.Diagnostics.HasError() {
			return
		}
	}

	// Set the "id" attribute.
	if err := r.setId(ctx, id, &response.State); err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotSetDiag(err))

		return
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "Resource.Read")
}

func (r *genericResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Update")

	conn := r.provider.CloudControlAPIClient(ctx)

	currentState := &request.State
	id, err := r.getId(ctx, currentState)

	if err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotFoundDiag(err))

		return
	}

	// Clear any write-only values.
	// This forces patch document generation to always add values.
	currentStateRaw := currentState.Raw
	if len(r.writeOnlyAttributePaths) > 0 {
		currentStateRaw, err = tftypes.Transform(currentStateRaw, func(tfPath *tftypes.AttributePath, val tftypes.Value) (tftypes.Value, error) {
			if len(tfPath.Steps()) < 1 {
				return val, nil
			}

			path, diags := attributePath(ctx, tfPath, currentState.Schema)
			if diags.HasError() {
				return val, ccdiag.DiagnosticsError(diags)
			}

			for _, woPath := range r.writeOnlyAttributePaths {
				if woPath.Equal(path) {
					return tftypes.NewValue(val.Type(), nil), nil
				}
			}

			return val, nil
		})
		if err != nil {
			response.Diagnostics.Append(DesiredStateErrorDiag("Prior State", err))

			return
		}
	}

	translator := toCloudControl{tfToCfNameMap: r.tfToCfNameMap}
	currentDesiredState, err := translator.AsString(ctx, currentState.Schema, currentStateRaw)

	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Prior State", err))

		return
	}

	plannedDesiredState, err := translator.AsString(ctx, request.Plan.Schema, request.Plan.Raw)

	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Plan", err))

		return
	}

	patchDocument, err := patchDocument(currentDesiredState, plannedDesiredState)

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of JSON Patch Unsuccessful",
			fmt.Sprintf("Unable to create a JSON Patch for resource update. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	tflog.Debug(ctx, "Cloud Control API PatchDocument", map[string]interface{}{
		"value": patchDocument,
	})

	input := &cloudcontrol.UpdateResourceInput{
		ClientToken:   aws.String(tfresource.UniqueId()),
		Identifier:    aws.String(id),
		PatchDocument: aws.String(patchDocument),
		TypeName:      aws.String(r.cfTypeName),
	}

	if roleARN := r.provider.RoleARN(ctx); roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.UpdateResource(ctx, input)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "UpdateResource", err))

		return
	}

	if output == nil || output.ProgressEvent == nil {
		response.Diagnostics.Append(ServiceOperationEmptyResultDiag("Cloud Control API", "UpdateResource"))

		return
	}

	waiter := cloudcontrol.NewResourceRequestSuccessWaiter(conn, func(o *cloudcontrol.ResourceRequestSuccessWaiterOptions) {
		o.Retryable = tfcloudcontrol.RetryGetResourceRequestStatus(nil)
	})

	err = waiter.Wait(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: output.ProgressEvent.RequestToken}, r.updateTimeout)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationWaiterErrorDiag("Cloud Control API", "UpdateResource", err))

		return
	}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	response.Diagnostics.Append(r.populateUnknownValues(ctx, id, &response.State)...)
	if response.Diagnostics.HasError() {
		return
	}

	traceExit(ctx, "Resource.Update")
}

func (r *genericResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Delete")

	conn := r.provider.CloudControlAPIClient(ctx)

	id, err := r.getId(ctx, &request.State)

	if err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotFoundDiag(err))

		return
	}

	err = tfcloudcontrol.DeleteResource(ctx, conn, r.provider.RoleARN(ctx), r.cfTypeName, id, r.deleteTimeout)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "DeleteResource", err))

		return
	}

	response.State.RemoveResource(ctx)

	traceExit(ctx, "Resource.Delete")
}

func (r *genericResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.ImportState")

	tflog.Debug(ctx, "Request.ID", map[string]interface{}{
		"value": hclog.Fmt("%v", request.ID),
	})

	resource.ImportStatePassthroughID(ctx, idAttributePath, request, response)

	traceExit(ctx, "Resource.ImportState")
}

// ConfigValidators returns a list of functions which will all be performed during validation.
func (r *genericResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	validators := make([]resource.ConfigValidator, 0)

	if len(r.configValidators) > 0 {
		validators = append(validators, r.configValidators...)
	}

	return validators
}

// describe returns the live state of the specified resource.
func (r *genericResource) describe(ctx context.Context, conn *cloudcontrol.Client, id string) (*cctypes.ResourceDescription, error) {
	return tfcloudcontrol.FindResourceByTypeNameAndID(ctx, conn, r.provider.RoleARN(ctx), r.cfTypeName, id)
}

// getId returns the resource's primary identifier value from State.
func (r *genericResource) getId(ctx context.Context, state *tfsdk.State) (string, error) {
	var val string
	diags := state.GetAttribute(ctx, idAttributePath, &val)

	if diags.HasError() {
		return "", ccdiag.DiagnosticsError(diags)
	}

	return val, nil
}

// setId sets the resource's primary identifier value in State.
func (r *genericResource) setId(ctx context.Context, val string, state *tfsdk.State) error {
	diags := state.SetAttribute(ctx, idAttributePath, val)

	if diags.HasError() {
		return ccdiag.DiagnosticsError(diags)
	}

	return nil
}

// populateUnknownValues populates and unknown values in State with values from the current resource description.
func (r *genericResource) populateUnknownValues(ctx context.Context, id string, state *tfsdk.State) diag.Diagnostics {
	var diags diag.Diagnostics

	unknowns, err := UnknownValuePaths(ctx, state.Raw)

	if err != nil {
		diags.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to set Terraform State Unknown values from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return diags
	}

	if len(unknowns) == 0 {
		return nil
	}

	description, err := r.describe(ctx, r.provider.CloudControlAPIClient(ctx), id)

	if tfresource.NotFound(err) {
		diags.Append(ResourceNotFoundAfterWriteDiag(err))

		return diags
	}

	if err != nil {
		diags.Append(ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))

		return diags
	}

	if description == nil {
		diags.Append(ServiceOperationEmptyResultDiag("Cloud Control API", "GetResource"))

		return diags
	}

	err = SetUnknownValuesFromResourceModel(ctx, state, unknowns, aws.ToString(description.Properties), r.cfToTfNameMap)

	if err != nil {
		diags.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to set Terraform State Unknown values from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return diags
	}

	return nil
}

// bootstrapContext injects the CloudFormation type name into logger contexts.
func (r *genericResource) bootstrapContext(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, LoggingKeyCFNType, r.cfTypeName)
	ctx = r.provider.RegisterLogger(ctx)

	return ctx
}

// propertyPathToAttributePath returns the AttributePath for the specified JSON Pointer property path.
func (r *genericResource) propertyPathToAttributePath(propertyPath string) (*path.Path, error) {
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

	attributePath := path.Empty()

	for _, segment := range segments[2:] {
		switch segment {
		case "", "*":
			return nil, fmt.Errorf("invalid property path segment: %q", segment)

		default:
			attributeName, ok := r.cfToTfNameMap[segment]
			if !ok {
				return nil, fmt.Errorf("attribute name mapping not found: %s", segment)
			}
			attributePath = attributePath.AtName(attributeName)
		}
	}

	return &attributePath, nil
}
