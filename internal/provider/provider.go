package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/smithy-go/logging"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

const (
	defaultMaxRetries         = 25
	defaultAssumeRoleDuration = 1 * time.Hour
)

func New() tfsdk.Provider {
	return &AwsCloudControlApiProvider{}
}

type AwsCloudControlApiProvider struct {
	ccClient *cloudcontrol.Client
	region   string
	roleARN  string
}

func (p *AwsCloudControlApiProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Version: 1,
		Attributes: map[string]tfsdk.Attribute{
			"access_key": {
				Type:        types.StringType,
				Description: "This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment variable, or via a shared credentials file if `profile` is specified.",
				Optional:    true,
			},

			"http_proxy": {
				Type:        types.StringType,
				Description: "The address of an HTTP proxy to use when accessing the AWS API. Can also be configured using the `HTTP_PROXY` or `HTTPS_PROXY` environment variables.",
				Optional:    true,
			},

			"insecure": {
				Type:        types.BoolType,
				Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If not set, defaults to `false`.",
				Optional:    true,
			},

			"max_retries": {
				Type:        types.Int64Type,
				Description: fmt.Sprintf("The maximum number of times an AWS API request is retried on failure. If not set, defaults to %d.", defaultMaxRetries),
				Optional:    true,
			},

			"profile": {
				Type:        types.StringType,
				Description: "This is the AWS profile name as set in the shared credentials file.",
				Optional:    true,
			},

			"region": {
				Type:        types.StringType,
				Description: "This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment variables, or via a shared config file.",
				Optional:    true,
			},

			"role_arn": {
				Type:        types.StringType,
				Description: "Amazon Resource Name of the AWS CloudFormation service role that is used on your behalf to perform operations.",
				Optional:    true,
				Validators: []tfsdk.AttributeValidator{
					validate.ARN(),
				},
			},

			"secret_key": {
				Type:        types.StringType,
				Description: "This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment variable, or via a shared credentials file if `profile` is specified.",
				Optional:    true,
			},

			"shared_config_files": {
				Type:        types.ListType{ElemType: types.StringType},
				Description: "List of paths to shared config files. If not set, defaults to `~/.aws/config`.",
				Optional:    true,
			},

			"shared_credentials_files": {
				Type:        types.ListType{ElemType: types.StringType},
				Description: "List of paths to shared credentials files. If not set, defaults to `~/.aws/credentials`.",
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

			"user_agent": {
				Attributes: tfsdk.ListNestedAttributes(
					map[string]tfsdk.Attribute{
						"product_name": {
							Type:        types.StringType,
							Description: "Product name. At least one of `product_name` or `comment` must be set.",
							Required:    true,
						},
						"product_version": {
							Type:        types.StringType,
							Description: "Product version. Optional, and should only be set when `product_name` is set.",
							Optional:    true,
						},
						"comment": {
							Type:        types.StringType,
							Description: "User-Agent comment. At least one of `comment` or `product_name` must be set.",
							Optional:    true,
						},
					},
					tfsdk.ListNestedAttributesOptions{},
				),
				Description: "Product details to append to User-Agent string in all AWS API calls.",
				Optional:    true,
			},

			"assume_role": {
				Attributes: tfsdk.SingleNestedAttributes(
					map[string]tfsdk.Attribute{
						"role_arn": {
							Type:        types.StringType,
							Description: "Amazon Resource Name (ARN) of the IAM Role to assume.",
							Required:    true,
							Validators: []tfsdk.AttributeValidator{
								validate.ARN(),
							},
						},

						"duration": {
							Type: cctypes.DurationType,
							Description: "Duration of the assume role session. You can provide a value from 15 minutes up to the maximum session duration setting for the role. " +
								cctypes.DurationType.Description() +
								fmt.Sprintf(" Default value is %s", defaultAssumeRoleDuration),
							Optional: true,
						},

						"external_id": {
							Type:        types.StringType,
							Description: "External identifier to use when assuming the role.",
							Optional:    true,
						},

						"policy": {
							Type:        types.StringType,
							Description: "IAM policy in JSON format to use as a session policy. The effective permissions for the session will be the intersection between this polcy and the role's policies.",
							Optional:    true,
							Validators: []tfsdk.AttributeValidator{
								validate.StringLenAtMost(2048),
								validate.StringIsJsonObject(),
							},
						},

						"policy_arns": {
							Type:        types.ListType{ElemType: types.StringType},
							Description: "Amazon Resource Names (ARNs) of IAM Policies to use as managed session policies. The effective permissions for the session will be the intersection between these polcy and the role's policies.",
							Optional:    true,
							Validators: []tfsdk.AttributeValidator{
								validate.ArrayForEach(
									validate.IAMPolicyARN(),
								),
							},
						},

						"session_name": {
							Type:        types.StringType,
							Description: "Session name to use when assuming the role.",
							Optional:    true,
						},

						"tags": {
							Description: "Map of assume role session tags.",
							Type:        types.MapType{ElemType: types.StringType},
							Optional:    true,
						},

						"transitive_tag_keys": {
							Description: "Set of assume role session tag keys to pass to any subsequent sessions.",
							Type:        types.SetType{ElemType: types.StringType},
							Optional:    true,
						},
					},
				),
				Optional:    true,
				Description: "An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration.",
			},
		},
	}, nil
}

type providerData struct {
	AccessKey              types.String       `tfsdk:"access_key"`
	HTTPProxy              types.String       `tfsdk:"http_proxy"`
	Insecure               types.Bool         `tfsdk:"insecure"`
	MaxRetries             types.Int64        `tfsdk:"max_retries"`
	Profile                types.String       `tfsdk:"profile"`
	Region                 types.String       `tfsdk:"region"`
	RoleARN                types.String       `tfsdk:"role_arn"`
	SecretKey              types.String       `tfsdk:"secret_key"`
	SharedConfigFiles      types.List         `tfsdk:"shared_config_files"`
	SharedCredentialsFiles types.List         `tfsdk:"shared_credentials_files"`
	SkipMetadataApiCheck   types.Bool         `tfsdk:"skip_medatadata_api_check"`
	Token                  types.String       `tfsdk:"token"`
	AssumeRole             *assumeRoleData    `tfsdk:"assume_role"`
	UserAgent              []userAgentProduct `tfsdk:"user_agent"`
	terraformVersion       string
}

type userAgentProduct struct {
	ProductName    types.String `tfsdk:"product_name"`
	ProductVersion types.String `tfsdk:"product_version"`
	Comment        types.String `tfsdk:"comment"`
}

type assumeRoleData struct {
	RoleARN           types.String     `tfsdk:"role_arn"`
	Duration          cctypes.Duration `tfsdk:"duration"`
	ExternalID        types.String     `tfsdk:"external_id"`
	Policy            types.String     `tfsdk:"policy"`
	PolicyARNs        types.List       `tfsdk:"policy_arns"`
	SessionName       types.String     `tfsdk:"session_name"`
	Tags              types.Map        `tfsdk:"tags"`
	TransitiveTagKeys types.Set        `tfsdk:"transitive_tag_keys"`
}

func (a assumeRoleData) Config() *awsbase.AssumeRole {
	assumeRole := &awsbase.AssumeRole{
		RoleARN:         a.RoleARN.Value,
		DurationSeconds: int(a.Duration.Value.Seconds()),
		ExternalID:      a.ExternalID.Value,
		Policy:          a.Policy.Value,
		SessionName:     a.SessionName.Value,
	}
	if !a.PolicyARNs.Null {
		arns := make([]string, len(a.PolicyARNs.Elems))
		for i, v := range a.PolicyARNs.Elems {
			arns[i] = v.(types.String).Value
		}
		assumeRole.PolicyARNs = arns
	}
	if !a.Tags.Null {
		tags := make(map[string]string)
		for key, value := range a.Tags.Elems {
			tags[key] = value.(types.String).Value
		}
		assumeRole.Tags = tags
	}
	if !a.TransitiveTagKeys.Null {
		tagKeys := make([]string, len(a.TransitiveTagKeys.Elems))
		for i, v := range a.TransitiveTagKeys.Elems {
			tagKeys[i] = v.(types.String).Value
		}
		assumeRole.TransitiveTagKeys = tagKeys
	}

	return assumeRole
}

func (p *AwsCloudControlApiProvider) Configure(ctx context.Context, request tfsdk.ConfigureProviderRequest, response *tfsdk.ConfigureProviderResponse) {
	var config providerData

	diags := request.Config.Get(ctx, &config)

	if diags.HasError() {
		response.Diagnostics.Append(diags...)

		return
	}

	if !request.Config.Raw.IsFullyKnown() {
		response.AddError("Unknown Value", "An attribute value is not yet known")
	}

	config.terraformVersion = request.TerraformVersion

	ccClient, region, err := newCloudControlClient(ctx, &config)

	if err != nil {
		response.Diagnostics.AddError(
			"Error configuring AWS CloudControl client",
			fmt.Sprintf("Error configuring the AWS Cloud Control API client, this is an error in the provider.\n%s\n", err),
		)

		return
	}

	p.ccClient = ccClient
	p.region = region
	p.roleARN = config.RoleARN.Value
}

func (p *AwsCloudControlApiProvider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
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

func (p *AwsCloudControlApiProvider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
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

func (p *AwsCloudControlApiProvider) CloudControlApiClient(_ context.Context) *cloudcontrol.Client {
	return p.ccClient
}

func (p *AwsCloudControlApiProvider) Region(_ context.Context) string {
	return p.region
}

func (p *AwsCloudControlApiProvider) RoleARN(_ context.Context) string {
	return p.roleARN
}

// newCloudControlClient configures and returns a fully initialized AWS Cloud Control API client with the configured region.
func newCloudControlClient(ctx context.Context, pd *providerData) (*cloudcontrol.Client, string, error) {
	logLevel := os.Getenv("TF_LOG")
	config := awsbase.Config{
		AccessKey:              pd.AccessKey.Value,
		CallerDocumentationURL: "https://registry.terraform.io/providers/hashicorp/awscc",
		CallerName:             "Terraform AWS Cloud Control Provider",
		DebugLogging:           strings.EqualFold(logLevel, "DEBUG") || strings.EqualFold(logLevel, "TRACE"),
		HTTPProxy:              pd.HTTPProxy.Value,
		Insecure:               pd.Insecure.Value,
		Profile:                pd.Profile.Value,
		Region:                 pd.Region.Value,
		SecretKey:              pd.SecretKey.Value,
		SkipMetadataApiCheck:   pd.SkipMetadataApiCheck.Value,
		Token:                  pd.Token.Value,
		APNInfo: &awsbase.APNInfo{
			PartnerName: "HashiCorp",
			Products: []awsbase.UserAgentProduct{
				{Name: "Terraform", Version: pd.terraformVersion, Comment: "+https://www.terraform.io"},
				{Name: "terraform-provider-awscc", Version: Version, Comment: "+https://registry.terraform.io/providers/hashicorp/awscc"},
			},
		},
	}
	config.UserAgent = userAgentProducts(pd.UserAgent)
	if pd.MaxRetries.Null {
		config.MaxRetries = defaultMaxRetries
	} else {
		config.MaxRetries = int(pd.MaxRetries.Value)
	}
	if !pd.SharedConfigFiles.Null {
		cf := make([]string, len(pd.SharedConfigFiles.Elems))
		for i, v := range pd.SharedConfigFiles.Elems {
			cf[i] = v.(types.String).Value
		}
		config.SharedConfigFiles = cf
	}
	if !pd.SharedCredentialsFiles.Null {
		cf := make([]string, len(pd.SharedCredentialsFiles.Elems))
		for i, v := range pd.SharedCredentialsFiles.Elems {
			cf[i] = v.(types.String).Value
		}
		config.SharedCredentialsFiles = cf
	}
	if pd.AssumeRole != nil {
		config.AssumeRole = pd.AssumeRole.Config()
	}

	cfg, err := awsbase.GetAwsConfig(ctx, &config)

	if err != nil {
		return nil, "", err
	}

	return cloudcontrol.NewFromConfig(cfg, func(o *cloudcontrol.Options) { o.Logger = awsSdkLogger{} }), cfg.Region, nil
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

func userAgentProducts(products []userAgentProduct) []awsbase.UserAgentProduct {
	results := make([]awsbase.UserAgentProduct, len(products))
	for i, p := range products {
		results[i] = awsbase.UserAgentProduct{
			Name:    p.ProductName.Value,
			Version: p.ProductVersion.Value,
			Comment: p.Comment.Value,
		}
	}
	return results
}
