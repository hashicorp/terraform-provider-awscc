package awsbase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

func getCredentialsProvider(ctx context.Context, c *Config) (aws.CredentialsProvider, error) {
	loadOptions := append(
		commonLoadOptions(c),
		config.WithSharedConfigProfile(c.Profile),
		// Bypass retries when validating authentication
		config.WithRetryer(func() aws.Retryer {
			return aws.NopRetryer{}
		}),
	)
	if c.AccessKey != "" || c.SecretKey != "" || c.Token != "" {
		loadOptions = append(
			loadOptions,
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					c.AccessKey,
					c.SecretKey,
					c.Token,
				),
			),
		)
	}
	if len(c.SharedCredentialsFiles) > 0 {
		loadOptions = append(
			loadOptions,
			config.WithSharedCredentialsFiles(c.SharedCredentialsFiles),
		)
	}

	cfg, err := config.LoadDefaultConfig(ctx, loadOptions...)
	if err != nil {
		return nil, fmt.Errorf("loading configuration: %w", err)
	}

	_, err = cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, c.NewNoValidCredentialSourcesError(err)
	}

	if c.AssumeRoleARN == "" {
		return cfg.Credentials, nil
	}

	return assumeRoleCredentialsProvider(ctx, cfg, c)
}

func assumeRoleCredentialsProvider(ctx context.Context, awsConfig aws.Config, c *Config) (aws.CredentialsProvider, error) {
	// When assuming a role, we need to first authenticate the base credentials above, then assume the desired role
	log.Printf("[INFO] Attempting to AssumeRole %s (SessionName: %q, ExternalId: %q)",
		c.AssumeRoleARN, c.AssumeRoleSessionName, c.AssumeRoleExternalID)

	client := sts.NewFromConfig(awsConfig)

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
	_, err := appCreds.Retrieve(ctx)
	if err != nil {
		return nil, c.NewCannotAssumeRoleError(err)
	}
	return aws.NewCredentialsCache(appCreds), nil
}
