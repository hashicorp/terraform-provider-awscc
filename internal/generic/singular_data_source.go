package generic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tflog "github.com/hashicorp/terraform-plugin-log"
	tfcloudformation "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

type SingularDataSourceType DataSourceType

func DataSourceWithAttributeNameMap(v map[string]string) DataSourceTypeOptionsFunc {
	return func(o *DataSourceType) error {
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

		o.cfToTfNameMap = cfToTfNameMap

		return nil
	}
}

func (opts DataSourceTypeOptions) WithAttributeNameMap(v map[string]string) DataSourceTypeOptions {
	return append(opts, DataSourceWithAttributeNameMap(v))
}

// NewSingularDataSourceType returns a new SingularDataSourceType from the specified variadic list of functional options.
// It's public as it's called from generated code.
func NewSingularDataSourceType(_ context.Context, optFns ...DataSourceTypeOptionsFunc) (tfsdk.DataSourceType, error) {
	dataSourceType := &DataSourceType{}

	for _, optFn := range optFns {
		err := optFn(dataSourceType)

		if err != nil {
			return nil, err
		}
	}

	if dataSourceType.cfTypeName == "" {
		return nil, fmt.Errorf("no CloudFormation type name specified")
	}

	if dataSourceType.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}

	var sdt *SingularDataSourceType

	return sdt.New(dataSourceType), nil
}

func (sdt *SingularDataSourceType) New(dst *DataSourceType) *SingularDataSourceType {
	return &SingularDataSourceType{
		cfToTfNameMap: dst.cfToTfNameMap,
		cfTypeName:    dst.cfTypeName,
		tfSchema:      dst.tfSchema,
		tfTypeName:    dst.tfTypeName,
	}
}

func (sdt *SingularDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, []*tfprotov6.Diagnostic) {
	return sdt.tfSchema, nil
}

func (sdt *SingularDataSourceType) NewDataSource(ctx context.Context, provider tfsdk.Provider) (tfsdk.DataSource, []*tfprotov6.Diagnostic) {
	return newGenericSingularDataSource(provider, sdt), nil
}

// Implements tfsdk.DataSource
type singularDataSource struct {
	provider       tfcloudformation.Provider
	dataSourceType *SingularDataSourceType
}

func newGenericSingularDataSource(provider tfsdk.Provider, singularDataSourceType *SingularDataSourceType) tfsdk.DataSource {
	return &singularDataSource{
		provider:       provider.(tfcloudformation.Provider),
		dataSourceType: singularDataSourceType,
	}
}

func (sd *singularDataSource) Read(ctx context.Context, request tfsdk.ReadDataSourceRequest, response *tfsdk.ReadDataSourceResponse) {
	ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())

	cfTypeName := sd.dataSourceType.cfTypeName
	tfTypeName := sd.dataSourceType.tfTypeName

	tflog.Debug(ctx, "DataSource.Read enter", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)

	conn := sd.provider.CloudFormationClient(ctx)

	currentConfig := &request.Config

	id, err := sd.getId(ctx, currentConfig)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotFoundDiag(err))

		return
	}

	description, err := sd.describe(ctx, conn, id)

	if tfresource.NotFound(err) {
		response.Diagnostics = append(response.Diagnostics, ResourceNotFoundWarningDiag(err))

		return
	}

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ServiceOperationErrorDiag("CloudFormation", "GetResource", err))

		return
	}

	translator := toTerraform{cfToTfNameMap: sd.dataSourceType.cfToTfNameMap}
	schema := &currentConfig.Schema
	val, err := translator.FromString(ctx, schema, aws.ToString(description.ResourceModel))

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Creation Of Terraform State Unsuccessful",
			Detail:   fmt.Sprintf("Unable to create a Terraform State value from a CloudFormation Resource Model. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		})

		return
	}

	response.State = tfsdk.State{
		Schema: *schema,
		Raw:    val,
	}

	err = sd.setId(ctx, aws.ToString(description.Identifier), &response.State)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, ResourceIdentifierNotSetDiag(err))

		return
	}

	tflog.Debug(ctx, "Response.State.Raw", "value", hclog.Fmt("%v", response.State.Raw))

	tflog.Debug(ctx, "DataSource.Read exit", "cfTypeName", cfTypeName, "tfTypeName", tfTypeName)
}

func (sd *singularDataSource) describe(ctx context.Context, conn *cloudformation.Client, id string) (*cftypes.ResourceDescription, error) {
	return tfcloudformation.FindResourceByTypeNameAndID(ctx, conn, sd.provider.RoleARN(ctx), sd.dataSourceType.cfTypeName, id)
}

func (sd *singularDataSource) getId(ctx context.Context, config *tfsdk.Config) (string, error) {
	val, diags := config.GetAttribute(ctx, idAttributePath)

	if tfresource.DiagsHasError(diags) {
		return "", tfresource.DiagsError(diags)
	}

	if val, ok := val.(types.String); ok {
		return val.Value, nil
	}

	return "", fmt.Errorf("invalid identifier type %T", val)
}

// setId sets the resource's primary identifier value in State.
func (sd *singularDataSource) setId(ctx context.Context, val string, state *tfsdk.State) error {
	diags := state.SetAttribute(ctx, idAttributePath, val)

	if tfresource.DiagsHasError(diags) {
		return tfresource.DiagsError(diags)
	}

	return nil
}
