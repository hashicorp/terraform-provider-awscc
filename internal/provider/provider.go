package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsbase "github.com/hashicorp/aws-sdk-go-base"

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
	cfClient *cloudformation.CloudFormation
}

func (p *awsCloudAPIProvider) GetSchema(ctx context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	tflog.Debug(ctx, "Provider.GetSchema() enter")

	return schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"region": {
				Type:        types.StringType,
				Description: "The region where AWS operations will take place.",
				Required:    true,
			},

			"role_arn": {
				Type:        types.StringType,
				Description: "Amazon Resource Name of an IAM Role that is used to do the actual provisioning.",
				Optional:    true,
			},
		},
	}, nil
}

type providerData struct {
	Region  types.String `tfsdk:"region"`
	RoleARN types.String `tfsdk:"role_arn"`
}

func (p *awsCloudAPIProvider) Configure(ctx context.Context, input tfsdk.ConfigureProviderRequest, output *tfsdk.ConfigureProviderResponse) {
	tflog.Debug(ctx, "Provider.Configure() enter")

	var config providerData

	err := input.Config.Get(ctx, &config)

	if err != nil {
		output.Diagnostics = append(output.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error parsing provider configuration",
			Detail:   fmt.Sprintf("Error parsing the provider configuration, this is an error in the provider.\n%s\n", err),
		})

		return
	}

	if config.Region.Null || config.Region.Unknown || config.RoleARN.Null || config.RoleARN.Unknown {
		tflog.Info(ctx, "One or more configuration values is Null or Unknown")

		return
	}

	cfClient, err := newCloudFormationClient(&config)

	if err != nil {
		output.Diagnostics = append(output.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error configuring AWS CloudFormation client",
			Detail:   fmt.Sprintf("Error configuring the AWS CloudFormation client, this is an error in the provider.\n%s\n", err),
		})

		return
	}

	p.cfClient = cfClient
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

func (p *awsCloudAPIProvider) CloudFormationClient(_ context.Context) (*cloudformation.CloudFormation, error) {
	return p.cfClient, nil
}

// newCloudFormationClient configures and returns a fully initialized AWS CloudFormation client.
func newCloudFormationClient(config *providerData) (*cloudformation.CloudFormation, error) {
	awsbaseConfig := &awsbase.Config{
		//DebugLogging: logging.IsDebugOrHigher(),
		Region: config.Region.Value,
	}

	sess, _, _, err := awsbase.GetSessionWithAccountIDAndPartition(awsbaseConfig)
	if err != nil {
		return nil, fmt.Errorf("error getting AWS SDK session: %w", err)
	}

	return cloudformation.New(sess.Copy(&aws.Config{})), nil
}
