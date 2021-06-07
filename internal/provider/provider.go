package provider

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsbase "github.com/hashicorp/aws-sdk-go-base"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"region": {
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"AWS_REGION",
						"AWS_DEFAULT_REGION",
					}, nil),
					Description:  "The region where AWS operations will take place.",
					InputDefault: "us-east-1",
				},

				"role_arn": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Amazon Resource Name of an IAM Role that is used to do the actual provisioning.",
				},
			},

			ResourcesMap: resources,
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		config := Config{
			Region: d.Get("region").(string),
		}

		return config.Client()
	}
}

var resources map[string]*schema.Resource
var resourcesMu sync.Mutex

func registerResource(name string, r *schema.Resource) {
	resourcesMu.Lock()
	defer resourcesMu.Unlock()

	if resources == nil {
		resources = map[string]*schema.Resource{}
	}
	resources[name] = r
}

// Minimal AWS client.
type AWSClient struct {
	cfconn *cloudformation.CloudFormation
}

type Config struct {
	Region string
}

// Client configures and returns a fully initialized AWSClient.
func (c *Config) Client() (interface{}, diag.Diagnostics) {
	awsbaseConfig := &awsbase.Config{
		DebugLogging: logging.IsDebugOrHigher(),
		Region:       c.Region,
	}

	sess, _, _, err := awsbase.GetSessionWithAccountIDAndPartition(awsbaseConfig)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("error configuring Terraform AWS Provider: %w", err))
	}

	client := &AWSClient{
		cfconn: cloudformation.New(sess.Copy(&aws.Config{})),
	}

	return client, nil
}
