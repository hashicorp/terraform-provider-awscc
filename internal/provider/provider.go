package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/smithy-go/logging"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func New() tfsdk.Provider {
	return &AwsCloudControlProvider{}
}

type AwsCloudControlProvider struct {
	cfClient *cloudformation.Client
	region   string
	roleARN  string
}

func (p *AwsCloudControlProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Version: 1,
		Attributes: map[string]tfsdk.Attribute{
			"access_key": {
				Type:        types.StringType,
				Description: "This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment variable, or via a shared credentials file if `profile` is specified.",
				Optional:    true,
			},

			"insecure": {
				Type:        types.BoolType,
				Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted, default value is `false`.",
				Optional:    true,
			},

			"profile": {
				Type:        types.StringType,
				Description: "This is the AWS profile name as set in the shared credentials file.",
				Optional:    true,
			},

			"region": {
				Type:        types.StringType,
				Description: "This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment variables, or via a shared credentials file if `profile` is specified.",
				Optional:    true,
			},

			"role_arn": {
				Type:        types.StringType,
				Description: "Amazon Resource Name of the AWS CloudFormation service role that is used on your behalf to perform operations.",
				Optional:    true,
			},

			"secret_key": {
				Type:        types.StringType,
				Description: "This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment variable, or via a shared credentials file if `profile` is specified.",
				Optional:    true,
			},

			"shared_credentials_file": {
				Type:        types.StringType,
				Description: "The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.",
				Optional:    true,
			},

			"skip_medatadata_api_check": {
				Type:        types.BoolType,
				Description: "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint.  Setting to `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
				Optional:    true,
			},

			"token": {
				Type:        types.StringType,
				Description: "Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.  It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.",
				Optional:    true,
			},

			"assume_role": {
				Attributes: tfsdk.SingleNestedAttributes(
					map[string]tfsdk.Attribute{
						"role_arn": {
							Type:        types.StringType,
							Description: "Amazon Resource Name (ARN) of the IAM Role to assume.",
							Required:    true,
						},
						"duration_seconds": {
							Type:        types.NumberType,
							Description: "Number of seconds to restrict the assume role session duration. You can provide a value from 900 seconds (15 minutes) up to the maximum session duration setting for the role.",
							Optional:    true,
						},
						"external_id": {
							Type:        types.StringType,
							Description: "External identifier to use when assuming the role.",
							Optional:    true,
						},
						"session_name": {
							Type:        types.StringType,
							Description: "Session name to use when assuming the role.",
							Optional:    true,
						},
						// "tags": {
						// 	Description: "Map of assume role session tags.",
						// 	Attributes: schema.SetNestedAttributes(
						// 		map[string]schema.Attribute{
						// 			"key": {
						// 				Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						// 				Type:        types.StringType,
						// 				Required:    true,
						// 			},
						// 			"value": {
						// 				Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						// 				Type:        types.StringType,
						// 				Required:    true,
						// 			},
						// 		},
						// 		schema.SetNestedAttributesOptions{},
						// 	),
						// 	Optional: true,
						// },
						// "transitive_tag_keys": {
						// 	Description: "Set of assume role session tag keys to pass to any subsequent sessions.",
						// 	Type:        providertypes.SetType{ElemType: types.StringType},
						// 	Optional:    true,
						// },
					},
				),
				Optional:    true,
				Description: "An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration.",
			},
		},
	}, nil
}

type providerData struct {
	AccessKey            types.String    `tfsdk:"access_key"`
	CredsFilename        types.String    `tfsdk:"shared_credentials_file"`
	Insecure             types.Bool      `tfsdk:"insecure"`
	Profile              types.String    `tfsdk:"profile"`
	Region               types.String    `tfsdk:"region"`
	RoleARN              types.String    `tfsdk:"role_arn"`
	SecretKey            types.String    `tfsdk:"secret_key"`
	SkipMetadataApiCheck types.Bool      `tfsdk:"skip_medatadata_api_check"`
	Token                types.String    `tfsdk:"token"`
	AssumeRole           *assumeRoleData `tfsdk:"assume_role"`
}

type assumeRoleData struct {
	RoleARN         types.String `tfsdk:"role_arn"`
	DurationSeconds types.Number `tfsdk:"duration_seconds"`
	ExternalID      types.String `tfsdk:"external_id"`
	SessionName     types.String `tfsdk:"session_name"`
}

func (p *AwsCloudControlProvider) Configure(ctx context.Context, request tfsdk.ConfigureProviderRequest, response *tfsdk.ConfigureProviderResponse) {
	var config providerData

	diags := request.Config.Get(ctx, &config)

	if diags.HasError() {
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

	cfClient, region, err := newCloudFormationClient(ctx, &config)

	if err != nil {
		response.Diagnostics.AddError(
			"Error configuring AWS CloudFormation client",
			fmt.Sprintf("Error configuring the AWS CloudFormation client, this is an error in the provider.\n%s\n", err),
		)

		return
	}

	p.cfClient = cfClient
	p.region = region
	p.roleARN = config.RoleARN.Value
}

func (p *AwsCloudControlProvider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	var diags diag.Diagnostics
	resources := make(map[string]tfsdk.ResourceType)

	for name, factory := range registry.ResourceFactories() {
		resourceType, err := factory(ctx)

		if err != nil {
			diags.AddError(
				"Error getting Resource",
				fmt.Sprintf("Error getting the %s Resource, this is an error in the provider.\n%s\n", name, err),
			)

			continue
		}

		resources[name] = resourceType
	}

	return resources, diags
}

func (p *AwsCloudControlProvider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	var diags diag.Diagnostics
	dataSources := make(map[string]tfsdk.DataSourceType)

	for name, factory := range registry.DataSourceFactories() {
		dataSourceType, err := factory(ctx)

		if err != nil {
			diags.AddError(
				"Error getting Data Source",
				fmt.Sprintf("Error getting the %s Data Source, this is an error in the provider.\n%s\n", name, err),
			)

			continue
		}

		dataSources[name] = dataSourceType
	}

	return dataSources, diags
}

func (p *AwsCloudControlProvider) CloudFormationClient(_ context.Context) *cloudformation.Client {
	return p.cfClient
}

func (p *AwsCloudControlProvider) Region(_ context.Context) string {
	return p.region
}

func (p *AwsCloudControlProvider) RoleARN(_ context.Context) string {
	return p.roleARN
}

// newCloudFormationClient configures and returns a fully initialized AWS CloudFormation client with the configured region.
func newCloudFormationClient(ctx context.Context, pd *providerData) (*cloudformation.Client, string, error) {
	logLevel := os.Getenv("TF_LOG")
	config := awsbase.Config{
		AccessKey:              pd.AccessKey.Value,
		DebugLogging:           strings.EqualFold(logLevel, "DEBUG") || strings.EqualFold(logLevel, "TRACE"),
		Insecure:               pd.Insecure.Value,
		Profile:                pd.Profile.Value,
		Region:                 pd.Region.Value,
		SecretKey:              pd.SecretKey.Value,
		SharedCredentialsFiles: []string{pd.CredsFilename.Value},
		SkipMetadataApiCheck:   pd.SkipMetadataApiCheck.Value,
		Token:                  pd.Token.Value,
	}
	if pd.AssumeRole != nil && !pd.AssumeRole.RoleARN.Null {
		config.AssumeRoleARN = pd.AssumeRole.RoleARN.Value

		if !pd.AssumeRole.DurationSeconds.Null {
			v, _ := pd.AssumeRole.DurationSeconds.Value.Int64()
			config.AssumeRoleDurationSeconds = int(v)
		}

		if !pd.AssumeRole.ExternalID.Null {
			config.AssumeRoleExternalID = pd.AssumeRole.ExternalID.Value
		}

		if !pd.AssumeRole.SessionName.Null {
			config.AssumeRoleSessionName = pd.AssumeRole.SessionName.Value
		}
	}

	cfg, err := awsbase.GetAwsConfig(ctx, &config)

	if err != nil {
		return nil, "", err
	}

	return cloudformation.NewFromConfig(cfg, func(o *cloudformation.Options) { o.Logger = awsSdkLogger{} }), cfg.Region, nil
}

type awsSdkLogger struct{}
type awsSdkContextLogger struct {
	ctx context.Context
}

func (l awsSdkLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	log.Printf("[%s] [aws-sdk-go-v2] %s", classification, fmt.Sprintf(format, v...))
}

func (l awsSdkLogger) WithContext(ctx context.Context) logging.Logger {
	return awsSdkContextLogger{ctx: ctx}
}

func (l awsSdkContextLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	switch classification {
	case logging.Warn:
		tflog.Warn(l.ctx, "[aws-sdk-go-v2]", "message", hclog.Fmt(format, v...))
	default:
		tflog.Debug(l.ctx, "[aws-sdk-go-v2]", "message", hclog.Fmt(format, v...))
	}
}
