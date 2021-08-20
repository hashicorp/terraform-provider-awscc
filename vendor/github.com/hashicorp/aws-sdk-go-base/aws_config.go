package awsbase

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/aws/smithy-go/logging"
	"github.com/hashicorp/go-cleanhttp"
)

func GetAwsConfig(ctx context.Context, c *Config) (aws.Config, error) {
	credentialsProvider, err := credentialsProvider(c)
	if err != nil {
		return aws.Config{}, err
	}

	var logMode aws.ClientLogMode
	var logger logging.Logger
	if c.DebugLogging {
		logMode = aws.LogRequestWithBody | aws.LogResponseWithBody | aws.LogRetries
		logger = debugLogger{}
	}

	imdsEnableState := imds.ClientDefaultEnableState
	if c.SkipMetadataApiCheck {
		imdsEnableState = imds.ClientDisabled
	}

	httpClient := cleanhttp.DefaultClient()
	if c.Insecure {
		transport := httpClient.Transport.(*http.Transport)
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentialsProvider),
		config.WithRegion(c.Region),
		config.WithSharedCredentialsFiles([]string{c.CredsFilename}),
		config.WithSharedConfigProfile(c.Profile),
		config.WithEndpointResolver(endpointResolver(c)),
		config.WithClientLogMode(logMode),
		config.WithLogger(logger),
		config.WithEC2IMDSClientEnableState(imdsEnableState),
		config.WithHTTPClient(httpClient),
	)

	if c.AssumeRoleARN == "" {
		return cfg, err
	}

	// When assuming a role, we need to first authenticate the base credentials above, then assume the desired role
	log.Printf("[INFO] Attempting to AssumeRole %s (SessionName: %q, ExternalId: %q)",
		c.AssumeRoleARN, c.AssumeRoleSessionName, c.AssumeRoleExternalID)

	client := sts.NewFromConfig(cfg)

	appCreds := stscreds.NewAssumeRoleProvider(client, c.AssumeRoleARN, func(opts *stscreds.AssumeRoleOptions) {
		opts.RoleSessionName = c.AssumeRoleSessionName
		opts.Duration = time.Duration(c.AssumeRoleDurationSeconds) * time.Second

		if c.AssumeRoleExternalID != "" {
			opts.ExternalID = aws.String(c.AssumeRoleExternalID)
		}

		if c.AssumeRolePolicy != "" {
			opts.Policy = aws.String(c.AssumeRolePolicy)
		}

		if len(c.AssumeRolePolicyARNs) > 0 {
			var policyDescriptorTypes []types.PolicyDescriptorType

			for _, policyARN := range c.AssumeRolePolicyARNs {
				policyDescriptorType := types.PolicyDescriptorType{
					Arn: aws.String(policyARN),
				}
				policyDescriptorTypes = append(policyDescriptorTypes, policyDescriptorType)
			}

			opts.PolicyARNs = policyDescriptorTypes
		}

		if len(c.AssumeRoleTags) > 0 {
			var tags []types.Tag

			for k, v := range c.AssumeRoleTags {
				tag := types.Tag{
					Key:   aws.String(k),
					Value: aws.String(v),
				}
				tags = append(tags, tag)
			}

			opts.Tags = tags
		}

		if len(c.AssumeRoleTransitiveTagKeys) > 0 {
			opts.TransitiveTagKeys = c.AssumeRoleTransitiveTagKeys
		}
	})
	_, err = appCreds.Retrieve(ctx)
	if err != nil {
		return aws.Config{}, fmt.Errorf("error assuming role: %w", err)
	}

	cfg.Credentials = aws.NewCredentialsCache(appCreds)

	return cfg, err
}
