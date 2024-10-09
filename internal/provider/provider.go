// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	awsbase "github.com/hashicorp/aws-sdk-go-base/v2"
	basediag "github.com/hashicorp/aws-sdk-go-base/v2/diag"
	baselogging "github.com/hashicorp/aws-sdk-go-base/v2/logging"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/flex"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

const (
	defaultMaxRetries         = 25
	defaultAssumeRoleDuration = 1 * time.Hour
)

// providerData is returned from the provider's Configure method and
// is passed to each resource and data source in their Configure methods.
type providerData struct {
	ccAPIClient *cloudcontrol.Client
	logger      baselogging.Logger
	region      string
	roleARN     string
}

func (p *providerData) CloudControlAPIClient(_ context.Context) *cloudcontrol.Client {
	return p.ccAPIClient
}

func (p *providerData) Region(_ context.Context) string {
	return p.region
}

func (p *providerData) RegisterLogger(ctx context.Context) context.Context {
	return baselogging.RegisterLogger(ctx, p.logger)
}

func (p *providerData) RoleARN(_ context.Context) string {
	return p.roleARN
}

type ccProvider struct {
	providerData *providerData // Used in acceptance tests.
}

func New() provider.Provider {
	return &ccProvider{}
}

// ProviderData is used in acceptance testing to get access to configured API client etc.
func (p *ccProvider) ProviderData() any {
	return p.providerData
}

func (p *ccProvider) Metadata(ctx context.Context, request provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "awscc"
	response.Version = Version
}

func (p *ccProvider) Schema(ctx context.Context, request provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_key": schema.StringAttribute{
				Description: "This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment variable, or via a shared credentials file if `profile` is specified.",
				Optional:    true,
			},
			"assume_role": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"duration": schema.StringAttribute{
						CustomType:  cctypes.DurationType,
						Description: "The duration, between 15 minutes and 12 hours, of the role session. Valid time units are ns, us (or µs), ms, s, h, or m.",
						Optional:    true,
					},
					"external_id": schema.StringAttribute{
						Description: "External identifier to use when assuming the role.",
						Optional:    true,
					},
					"policy": schema.StringAttribute{
						CustomType:  jsontypes.ExactType{},
						Description: "IAM policy in JSON format to use as a session policy. The effective permissions for the session will be the intersection between this polcy and the role's policies.",
						Optional:    true,
					},
					"policy_arns": schema.ListAttribute{
						ElementType: cctypes.ARNType,
						Description: "Amazon Resource Names (ARNs) of IAM Policies to use as managed session policies. The effective permissions for the session will be the intersection between these polcy and the role's policies.",
						Optional:    true,
					},
					"role_arn": schema.StringAttribute{
						CustomType:  cctypes.ARNType,
						Description: "Amazon Resource Name (ARN) of the IAM Role to assume.",
						Required:    true,
					},
					"session_name": schema.StringAttribute{
						Description: "Session name to use when assuming the role.",
						Optional:    true,
					},
					"tags": schema.MapAttribute{
						ElementType: types.StringType,
						Description: "Map of assume role session tags.",
						Optional:    true,
					},
					"transitive_tag_keys": schema.SetAttribute{
						ElementType: types.StringType,
						Description: "Set of assume role session tag keys to pass to any subsequent sessions.",
						Optional:    true,
					},
				},
				Optional:    true,
				Description: "An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration.",
			},
			"assume_role_with_web_identity": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"duration": schema.StringAttribute{
						CustomType:  cctypes.DurationType,
						Description: "The duration, between 15 minutes and 12 hours, of the role session. Valid time units are ns, us (or µs), ms, s, h, or m.",
						Optional:    true,
					},
					"policy": schema.StringAttribute{
						CustomType:  jsontypes.ExactType{},
						Description: "IAM policy in JSON format to use as a session policy. The effective permissions for the session will be the intersection between this polcy and the role's policies.",
						Optional:    true,
					},
					"policy_arns": schema.ListAttribute{
						ElementType: cctypes.ARNType,
						Description: "Amazon Resource Names (ARNs) of IAM Policies to use as managed session policies. The effective permissions for the session will be the intersection between these polcy and the role's policies.",
						Optional:    true,
					},
					"role_arn": schema.StringAttribute{
						CustomType:  cctypes.ARNType,
						Description: "Amazon Resource Name (ARN) of the IAM Role to assume. Can also be set with the environment variable `AWS_ROLE_ARN`.",
						Required:    true,
					},
					"session_name": schema.StringAttribute{
						Description: "Session name to use when assuming the role. Can also be set with the environment variable `AWS_ROLE_SESSION_NAME`.",
						Optional:    true,
					},
					"web_identity_token": schema.StringAttribute{
						Description: "The value of a web identity token from an OpenID Connect (OIDC) or OAuth provider. One of `web_identity_token` or `web_identity_token_file` is required.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.LengthBetween(4, 20000),
						},
					},
					"web_identity_token_file": schema.StringAttribute{
						Description: "File containing a web identity token from an OpenID Connect (OIDC) or OAuth provider. Can also be set with the  environment variable`AWS_WEB_IDENTITY_TOKEN_FILE`. One of `web_identity_token_file` or `web_identity_token` is required.",
						Optional:    true,
					},
				},
				Optional:    true,
				Description: "An `assume_role_with_web_identity` block (documented below). Only one `assume_role_with_web_identity` block may be in the configuration.",
			},
			"endpoints": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"cloudcontrolapi": schema.StringAttribute{
						Optional:    true,
						Description: "Use this to override the default Cloud Control API service endpoint URL",
					},
					"iam": schema.StringAttribute{
						Optional:    true,
						Description: "Use this to override the default IAM service endpoint URL",
					},
					"sso": schema.StringAttribute{
						Optional:    true,
						Description: "Use this to override the default SSO service endpoint URL",
					},
					"sts": schema.StringAttribute{
						Optional:    true,
						Description: "Use this to override the default STS service endpoint URL",
					},
				},
				Optional:    true,
				Description: "An `endpoints` block (documented below). Only one `endpoints` block may be in the configuration.",
			},
			"http_proxy": schema.StringAttribute{
				Description: "URL of a proxy to use for HTTP requests when accessing the AWS API. Can also be set using the `HTTP_PROXY` or `http_proxy` environment variables.",
				Optional:    true,
			},
			"https_proxy": schema.StringAttribute{
				Description: "URL of a proxy to use for HTTPS requests when accessing the AWS API. Can also be set using the `HTTPS_PROXY` or `https_proxy` environment variables.",
				Optional:    true,
			},
			"insecure": schema.BoolAttribute{
				Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If not set, defaults to `false`.",
				Optional:    true,
			},
			"max_retries": schema.Int64Attribute{
				Description: fmt.Sprintf("The maximum number of times an AWS API request is retried on failure. If not set, defaults to %d.", defaultMaxRetries),
				Optional:    true,
			},
			"no_proxy": schema.StringAttribute{
				Description: "Comma-separated list of hosts that should not use HTTP or HTTPS proxies. Can also be set using the `NO_PROXY` or `no_proxy` environment variables.",
				Optional:    true,
			},
			"profile": schema.StringAttribute{
				Description: "This is the AWS profile name as set in the shared credentials file.",
				Optional:    true,
			},
			"region": schema.StringAttribute{
				Description: "This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment variables, via a shared config file, or from the EC2 Instance Metadata Service if used.",
				Optional:    true,
			},
			"role_arn": schema.StringAttribute{
				CustomType:  cctypes.ARNType,
				Description: "Amazon Resource Name of the AWS CloudFormation service role that is used on your behalf to perform operations.",
				Optional:    true,
			},
			"secret_key": schema.StringAttribute{
				Description: "This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment variable, or via a shared credentials file if `profile` is specified.",
				Optional:    true,
			},
			"shared_config_files": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of paths to shared config files. If not set, defaults to `~/.aws/config`.",
				Optional:    true,
			},
			"shared_credentials_files": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of paths to shared credentials files. If not set, defaults to `~/.aws/credentials`.",
				Optional:    true,
			},
			"skip_medatadata_api_check": schema.BoolAttribute{
				Description:        "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint.  Setting to `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
				Optional:           true,
				DeprecationMessage: `Use "skip_metadata_api_check" instead`,
			},
			"skip_metadata_api_check": schema.BoolAttribute{
				Description: "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint.  Setting to `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
				Optional:    true,
			},
			"token": schema.StringAttribute{
				Description: "Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.  It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.",
				Optional:    true,
			},
			"user_agent": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							Description: "User-Agent comment. At least one of `comment` or `product_name` must be set.",
							Optional:    true,
						},
						"product_name": schema.StringAttribute{
							Description: "Product name. At least one of `product_name` or `comment` must be set.",
							Required:    true,
						},
						"product_version": schema.StringAttribute{
							Description: "Product version. Optional, and should only be set when `product_name` is set.",
							Optional:    true,
						},
					},
				},
				Description: "Product details to append to User-Agent string in all AWS API calls.",
				Optional:    true,
			},
		},
	}
}

type configModel struct {
	AccessKey                 types.String                   `tfsdk:"access_key"`
	AssumeRole                *assumeRoleModel               `tfsdk:"assume_role"`
	AssumeRoleWithWebIdentity *assumeRoleWithWebIdentityData `tfsdk:"assume_role_with_web_identity"`
	Endpoints                 *endpointData                  `tfsdk:"endpoints"`
	HTTPProxy                 types.String                   `tfsdk:"http_proxy"`
	HTTPSProxy                types.String                   `tfsdk:"https_proxy"`
	Insecure                  types.Bool                     `tfsdk:"insecure"`
	MaxRetries                types.Int64                    `tfsdk:"max_retries"`
	NoProxy                   types.String                   `tfsdk:"no_proxy"`
	Profile                   types.String                   `tfsdk:"profile"`
	Region                    types.String                   `tfsdk:"region"`
	RoleARN                   cctypes.ARN                    `tfsdk:"role_arn"`
	SecretKey                 types.String                   `tfsdk:"secret_key"`
	SharedConfigFiles         types.List                     `tfsdk:"shared_config_files"`
	SharedCredentialsFiles    types.List                     `tfsdk:"shared_credentials_files"`
	SkipMedatadataApiCheck    types.Bool                     `tfsdk:"skip_medatadata_api_check"`
	SkipMetadataApiCheck      types.Bool                     `tfsdk:"skip_metadata_api_check"`
	Token                     types.String                   `tfsdk:"token"`
	UserAgent                 []userAgentProduct             `tfsdk:"user_agent"`

	terraformVersion string
}

type userAgentProduct struct {
	Comment        types.String `tfsdk:"comment"`
	ProductName    types.String `tfsdk:"product_name"`
	ProductVersion types.String `tfsdk:"product_version"`
}

type assumeRoleModel struct {
	Duration          cctypes.Duration `tfsdk:"duration"`
	ExternalID        types.String     `tfsdk:"external_id"`
	Policy            jsontypes.Exact  `tfsdk:"policy"`
	PolicyARNs        types.List       `tfsdk:"policy_arns"`
	RoleARN           cctypes.ARN      `tfsdk:"role_arn"`
	SessionName       types.String     `tfsdk:"session_name"`
	Tags              types.Map        `tfsdk:"tags"`
	TransitiveTagKeys types.Set        `tfsdk:"transitive_tag_keys"`
}

func (a assumeRoleModel) Config() awsbase.AssumeRole {
	assumeRole := awsbase.AssumeRole{
		Duration:    a.Duration.ValueDuration(),
		ExternalID:  a.ExternalID.ValueString(),
		Policy:      a.Policy.ValueString(),
		RoleARN:     a.RoleARN.ValueString(),
		SessionName: a.SessionName.ValueString(),
	}
	if !a.PolicyARNs.IsNull() {
		arns := make([]string, len(a.PolicyARNs.Elements()))
		for i, v := range a.PolicyARNs.Elements() {
			arns[i] = v.(types.String).ValueString()
		}
		assumeRole.PolicyARNs = arns
	}
	if !a.Tags.IsNull() {
		tags := make(map[string]string)
		for key, value := range a.Tags.Elements() {
			tags[key] = value.(types.String).ValueString()
		}
		assumeRole.Tags = tags
	}
	if !a.TransitiveTagKeys.IsNull() {
		tagKeys := make([]string, len(a.TransitiveTagKeys.Elements()))
		for i, v := range a.TransitiveTagKeys.Elements() {
			tagKeys[i] = v.(types.String).ValueString()
		}
		assumeRole.TransitiveTagKeys = tagKeys
	}

	return assumeRole
}

type endpointData struct {
	CloudControlAPI types.String `tfsdk:"cloudcontrolapi"`
	IAM             types.String `tfsdk:"iam"`
	SSO             types.String `tfsdk:"sso"`
	STS             types.String `tfsdk:"sts"`
}

type assumeRoleWithWebIdentityData struct {
	Duration             cctypes.Duration `tfsdk:"duration"`
	Policy               jsontypes.Exact  `tfsdk:"policy"`
	PolicyARNs           types.List       `tfsdk:"policy_arns"`
	RoleARN              cctypes.ARN      `tfsdk:"role_arn"`
	SessionName          types.String     `tfsdk:"session_name"`
	WebIdentityToken     types.String     `tfsdk:"web_identity_token"`
	WebIdentityTokenFile types.String     `tfsdk:"web_identity_token_file"`
}

func (a assumeRoleWithWebIdentityData) Config() *awsbase.AssumeRoleWithWebIdentity {
	assumeRole := &awsbase.AssumeRoleWithWebIdentity{
		Duration:             a.Duration.ValueDuration(),
		Policy:               a.Policy.ValueString(),
		RoleARN:              a.RoleARN.ValueString(),
		SessionName:          a.SessionName.ValueString(),
		WebIdentityToken:     a.WebIdentityToken.ValueString(),
		WebIdentityTokenFile: a.WebIdentityTokenFile.ValueString(),
	}
	if !a.PolicyARNs.IsNull() {
		arns := make([]string, len(a.PolicyARNs.Elements()))
		for i, v := range a.PolicyARNs.Elements() {
			arns[i] = v.(types.String).ValueString()
		}
		assumeRole.PolicyARNs = arns
	}

	return assumeRole
}

func (p *ccProvider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	var config configModel

	response.Diagnostics.Append(request.Config.Get(ctx, &config)...)
	if response.Diagnostics.HasError() {
		return
	}

	if !request.Config.Raw.IsFullyKnown() {
		response.Diagnostics.AddError("Unknown Value", "An attribute value is not yet known")
	}

	config.terraformVersion = request.TerraformVersion

	providerData, diags := newProviderData(ctx, &config)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	p.providerData = providerData
	response.DataSourceData = providerData
	response.ResourceData = providerData
}

func (p *ccProvider) Resources(ctx context.Context) []func() resource.Resource {
	var diags diag.Diagnostics
	var resources = make([]func() resource.Resource, 0)

	for name, factory := range registry.ResourceFactories() {
		v, err := factory(ctx)

		if err != nil {
			diags.AddError(
				"Error getting Resource",
				fmt.Sprintf("Error getting the %s Resource, this is an error in the provider.\n%s\n", name, err),
			)

			continue
		}

		resources = append(resources, func() resource.Resource {
			return v
		})
	}

	return resources
}

func (p *ccProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	var diags diag.Diagnostics
	dataSources := make([]func() datasource.DataSource, 0)

	for name, factory := range registry.DataSourceFactories() {
		v, err := factory(ctx)

		if err != nil {
			diags.AddError(
				"Error getting Data Source",
				fmt.Sprintf("Error getting the %s Data Source, this is an error in the provider.\n%s\n", name, err),
			)

			continue
		}

		dataSources = append(dataSources, func() datasource.DataSource {
			return v
		})
	}

	return dataSources
}

func newProviderData(ctx context.Context, c *configModel) (*providerData, diag.Diagnostics) {
	var diags diag.Diagnostics

	ctx, logger := baselogging.NewTfLogger(ctx)

	awsbaseConfig := awsbase.Config{
		AccessKey:              c.AccessKey.ValueString(),
		CallerDocumentationURL: "https://registry.terraform.io/providers/hashicorp/awscc",
		CallerName:             "Terraform AWS Cloud Control Provider",
		HTTPProxy:              flex.StringFromFramework(ctx, c.HTTPProxy),
		HTTPProxyMode:          awsbase.HTTPProxyModeLegacy,
		HTTPSProxy:             flex.StringFromFramework(ctx, c.HTTPSProxy),
		Insecure:               c.Insecure.ValueBool(),
		Logger:                 logger,
		NoProxy:                c.NoProxy.ValueString(),
		Profile:                c.Profile.ValueString(),
		Region:                 c.Region.ValueString(),
		SecretKey:              c.SecretKey.ValueString(),
		Token:                  c.Token.ValueString(),
		APNInfo: &awsbase.APNInfo{
			PartnerName: "HashiCorp",
			Products: []awsbase.UserAgentProduct{
				{Name: "Terraform", Version: c.terraformVersion, Comment: "+https://www.terraform.io"},
				{Name: "terraform-provider-awscc", Version: Version, Comment: "+https://registry.terraform.io/providers/hashicorp/awscc"},
			},
		},
	}
	awsbaseConfig.UserAgent = userAgentProducts(c.UserAgent)
	if c.MaxRetries.IsNull() {
		awsbaseConfig.MaxRetries = defaultMaxRetries
	} else {
		awsbaseConfig.MaxRetries = int(c.MaxRetries.ValueInt64())
	}
	if !c.SharedConfigFiles.IsNull() {
		cf := make([]string, len(c.SharedConfigFiles.Elements()))
		for i, v := range c.SharedConfigFiles.Elements() {
			cf[i] = v.(types.String).ValueString()
		}
		awsbaseConfig.SharedConfigFiles = cf
	}
	if !c.SharedCredentialsFiles.IsNull() {
		cf := make([]string, len(c.SharedCredentialsFiles.Elements()))
		for i, v := range c.SharedCredentialsFiles.Elements() {
			cf[i] = v.(types.String).ValueString()
		}
		awsbaseConfig.SharedCredentialsFiles = cf
	}
	if c.AssumeRole != nil {
		awsbaseConfig.AssumeRole = []awsbase.AssumeRole{c.AssumeRole.Config()}
	}

	if c.AssumeRoleWithWebIdentity != nil {
		awsbaseConfig.AssumeRoleWithWebIdentity = c.AssumeRoleWithWebIdentity.Config()
	}

	if c.SkipMetadataApiCheck.IsNull() {
		if c.SkipMedatadataApiCheck.IsNull() {
			awsbaseConfig.EC2MetadataServiceEnableState = imds.ClientDefaultEnableState
		} else if !c.SkipMedatadataApiCheck.ValueBool() {
			awsbaseConfig.EC2MetadataServiceEnableState = imds.ClientDisabled
		} else {
			awsbaseConfig.EC2MetadataServiceEnableState = imds.ClientEnabled
		}
	} else if !c.SkipMetadataApiCheck.ValueBool() {
		awsbaseConfig.EC2MetadataServiceEnableState = imds.ClientDisabled
	} else {
		awsbaseConfig.EC2MetadataServiceEnableState = imds.ClientEnabled
	}

	if c.Endpoints != nil && !c.Endpoints.IAM.IsNull() {
		awsbaseConfig.IamEndpoint = c.Endpoints.IAM.ValueString()
	}
	if c.Endpoints != nil && !c.Endpoints.SSO.IsNull() {
		awsbaseConfig.SsoEndpoint = c.Endpoints.SSO.ValueString()
	}
	if c.Endpoints != nil && !c.Endpoints.STS.IsNull() {
		awsbaseConfig.StsEndpoint = c.Endpoints.STS.ValueString()
	}

	_, cfg, awsDiags := awsbase.GetAwsConfig(ctx, &awsbaseConfig)

	for _, d := range awsDiags {
		switch d.Severity() {
		case basediag.SeverityWarning:
			diags = append(diags, diag.NewWarningDiagnostic(d.Summary(), d.Detail()))
		case basediag.SeverityError:
			diags = append(diags, diag.NewErrorDiagnostic(d.Summary(), d.Detail()))
		}
	}

	if diags.HasError() {
		return nil, diags
	}

	ccAPIClient := cloudcontrol.NewFromConfig(cfg, func(o *cloudcontrol.Options) {
		if c.Endpoints != nil {
			o.BaseEndpoint = flex.StringFromFramework(ctx, c.Endpoints.CloudControlAPI)
		}
	})

	providerData := &providerData{
		ccAPIClient: ccAPIClient,
		logger:      logger,
		region:      cfg.Region,
		roleARN:     c.RoleARN.ValueString(),
	}

	return providerData, diags
}

func userAgentProducts(products []userAgentProduct) []awsbase.UserAgentProduct {
	results := make([]awsbase.UserAgentProduct, len(products))
	for i, p := range products {
		results[i] = awsbase.UserAgentProduct{
			Name:    p.ProductName.ValueString(),
			Version: p.ProductVersion.ValueString(),
			Comment: p.Comment.ValueString(),
		}
	}
	return results
}
