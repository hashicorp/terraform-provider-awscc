package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func New() tfsdk.Provider {
	return &awsCloudAPIProvider{}
}

type awsCloudAPIProvider struct {
	cfClient *cloudformation.Client
	roleARN  string
}

func (p *awsCloudAPIProvider) GetSchema(ctx context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	tflog.Debug(ctx, "Provider.GetSchema() enter")

	return schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
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
		},
	}, nil
}

type providerData struct {
	Region  types.String `tfsdk:"region"`
	RoleARN types.String `tfsdk:"role_arn"`
}

func (p *awsCloudAPIProvider) Configure(ctx context.Context, request tfsdk.ConfigureProviderRequest, response *tfsdk.ConfigureProviderResponse) {
	tflog.Debug(ctx, "Provider.Configure() enter")

	var config providerData

	err := request.Config.Get(ctx, &config)

	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error parsing provider configuration",
			Detail:   fmt.Sprintf("Error parsing the provider configuration, this is an error in the provider.\n%s\n", err),
		})

		return
	}

	if config.Region.Unknown {
		tflog.Info(ctx, "AWS Region is Unknown")

		return
	}

	if config.RoleARN.Unknown {
		tflog.Info(ctx, "Role ARN is Unknown")

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

func (p *awsCloudAPIProvider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, []*tfprotov6.Diagnostic) {
	tflog.Debug(ctx, "Provider.GetResources() enter")

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

func (p *awsCloudAPIProvider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, []*tfprotov6.Diagnostic) {
	tflog.Debug(ctx, "Provider.GetDataSources() enter")

	return nil, nil
}

func (p *awsCloudAPIProvider) CloudFormationClient(_ context.Context) *cloudformation.Client {
	return p.cfClient
}

func (p *awsCloudAPIProvider) RoleARN(_ context.Context) string {
	return p.roleARN
}

// newCloudFormationClient configures and returns a fully initialized AWS CloudFormation client.
func newCloudFormationClient(ctx context.Context, pd *providerData) (*cloudformation.Client, error) {
	optFns := make([]func(*config.LoadOptions) error, 0)

	if region := pd.Region.Value; region != "" {
		optFns = append(optFns, config.WithRegion(region))
	}

	cfg, err := config.LoadDefaultConfig(ctx, optFns...)

	if err != nil {
		return nil, err
	}

	return cloudformation.NewFromConfig(cfg), nil
}
