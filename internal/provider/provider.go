package provider

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	awsbase "github.com/hashicorp/aws-sdk-go-base"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func New() tfsdk.Provider {
	return &awsCloudControlProvider{}
}

type awsCloudControlProvider struct {
	cfClient *cloudformation.Client
	roleARN  string
}

func (p *awsCloudControlProvider) GetSchema(ctx context.Context) (tfsdk.Schema, []*tfprotov6.Diagnostic) {
	return tfsdk.Schema{
		Version: 1,
		Attributes: map[string]tfsdk.Attribute{
			"access_key": {
				Type:        types.StringType,
				Description: "The access key for API operations.",
				Optional:    true,
			},

			"insecure": {
				Type:        types.BoolType,
				Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted, default value is `false`.",
				Optional:    true,
			},

			"profile": {
				Type:        types.StringType,
				Description: "The profile for API operations. If not set, the default profile created with `aws configure` will be used.",
				Optional:    true,
			},

			"region": {
				Type:        types.StringType,
				Description: "The region where AWS operations will take place.",
				Optional:    true,
			},

			"role_arn": {
				Type:        types.StringType,
				Description: "Amazon Resource Name of the AWS CloudFormation service role that is used on your behalf to perform operations.",
				Optional:    true,
			},

			"secret_key": {
				Type:        types.StringType,
				Description: "The secret key for API operations.",
				Optional:    true,
			},

			"shared_credentials_file": {
				Type:        types.StringType,
				Description: "The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.",
				Optional:    true,
			},

			"skip_medatadata_api_check": {
				Type:        types.BoolType,
				Description: "Skip the AWS Metadata API check. Used for AWS API implementations that do not have a Metadata API endpoint.",
				Optional:    true,
			},

			"token": {
				Type:        types.StringType,
				Description: "Session token. A session token is only required if you are using temporary security credentials.",
				Optional:    true,
			},
		},
	}, nil
}

type providerData struct {
	AccessKey            types.String `tfsdk:"access_key"`
	CredsFilename        types.String `tfsdk:"shared_credentials_file"`
	Insecure             types.Bool   `tfsdk:"insecure"`
	Profile              types.String `tfsdk:"profile"`
	Region               types.String `tfsdk:"region"`
	RoleARN              types.String `tfsdk:"role_arn"`
	SecretKey            types.String `tfsdk:"secret_key"`
	SkipMetadataApiCheck types.Bool   `tfsdk:"skip_medatadata_api_check"`
	Token                types.String `tfsdk:"token"`
}

func (p *awsCloudControlProvider) Configure(ctx context.Context, request tfsdk.ConfigureProviderRequest, response *tfsdk.ConfigureProviderResponse) {
	var config providerData

	diags := request.Config.Get(ctx, &config)

	if tfresource.DiagsHasError(diags) {
		response.Diagnostics = append(response.Diagnostics, diags...)

		return
	}

	// TODO
	// TODO Is this the correct thing to do for any Unknown values?
	// TODO
	anyUnknownConfigValues := false

	if config.AccessKey.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("access_key"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.CredsFilename.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("shared_credentials_file"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.Insecure.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("insecure"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.Profile.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("profile"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.Region.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("region"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.RoleARN.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("role_arn"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.SecretKey.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("secret_key"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.SkipMetadataApiCheck.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("skip_medatadata_api_check"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if config.Token.Unknown {
		response.AddAttributeError(tftypes.NewAttributePath().WithAttributeName("token"), "Unknown Value", "Attribute value is not yet known")
		anyUnknownConfigValues = true
	}

	if anyUnknownConfigValues {
		return
	}

	cfClient, err := newCloudFormationClient(ctx, &config)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error configuring AWS CloudFormation client",
			Detail:   fmt.Sprintf("Error configuring the AWS CloudFormation client, this is an error in the provider.\n%s\n", err),
		})

		return
	}

	p.cfClient = cfClient
	p.roleARN = config.RoleARN.Value
}

func (p *awsCloudControlProvider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic
	resources := make(map[string]tfsdk.ResourceType)

	for name, factory := range registry.ResourceFactories() {
		resourceType, err := factory(ctx)

		if err != nil {
			diags = append(diags, &tfprotov6.Diagnostic{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error getting Resource",
				Detail:   fmt.Sprintf("Error getting the %s Resource, this is an error in the provider.\n%s\n", name, err),
			})

			continue
		}

		resources[name] = resourceType
	}

	return resources, diags
}

func (p *awsCloudControlProvider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, []*tfprotov6.Diagnostic) {
	return nil, nil
}

func (p *awsCloudControlProvider) CloudFormationClient(_ context.Context) *cloudformation.Client {
	return p.cfClient
}

func (p *awsCloudControlProvider) RoleARN(_ context.Context) string {
	return p.roleARN
}

// newCloudFormationClient configures and returns a fully initialized AWS CloudFormation client.
func newCloudFormationClient(ctx context.Context, pd *providerData) (*cloudformation.Client, error) {
	logLevel := os.Getenv("TF_LOG")
	config := awsbase.Config{
		AccessKey:            pd.AccessKey.Value,
		CredsFilename:        pd.CredsFilename.Value,
		DebugLogging:         strings.EqualFold(logLevel, "DEBUG") || strings.EqualFold(logLevel, "TRACE"),
		Insecure:             pd.Insecure.Value,
		Profile:              pd.Profile.Value,
		Region:               pd.Region.Value,
		SecretKey:            pd.SecretKey.Value,
		SkipMetadataApiCheck: pd.SkipMetadataApiCheck.Value,
		Token:                pd.Token.Value,
	}

	cfg, err := awsbase.GetAwsConfig(ctx, &config)

	if err != nil {
		return nil, err
	}

	return cloudformation.NewFromConfig(cfg), nil
}
