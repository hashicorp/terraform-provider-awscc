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
	loadOptions, err := commonLoadOptions(c)
	if err != nil {
		return nil, err
	}
	loadOptions = append(
		loadOptions,
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

	if c.AssumeRole == nil || c.AssumeRole.RoleARN == "" {
		return cfg.Credentials, nil
	}

	return assumeRoleCredentialsProvider(ctx, cfg, c)
}

func assumeRoleCredentialsProvider(ctx context.Context, awsConfig aws.Config, c *Config) (aws.CredentialsProvider, error) {
	ar := c.AssumeRole
	// When assuming a role, we need to first authenticate the base credentials above, then assume the desired role
	log.Printf("[INFO] Attempting to AssumeRole %s (SessionName: %q, ExternalId: %q)",
		ar.RoleARN, ar.SessionName, ar.ExternalID)

	client := sts.NewFromConfig(awsConfig)

	appCreds := stscreds.NewAssumeRoleProvider(client, ar.RoleARN, func(opts *stscreds.AssumeRoleOptions) {
		opts.RoleSessionName = ar.SessionName
		opts.Duration = time.Duration(ar.DurationSeconds) * time.Second

		if ar.ExternalID != "" {
			opts.ExternalID = aws.String(ar.ExternalID)
		}

		if ar.Policy != "" {
			opts.Policy = aws.String(ar.Policy)
		}

		if len(ar.PolicyARNs) > 0 {
			var policyDescriptorTypes []types.PolicyDescriptorType

			for _, policyARN := range ar.PolicyARNs {
				policyDescriptorType := types.PolicyDescriptorType{
					Arn: aws.String(policyARN),
				}
				policyDescriptorTypes = append(policyDescriptorTypes, policyDescriptorType)
			}

			opts.PolicyARNs = policyDescriptorTypes
		}

		if len(ar.Tags) > 0 {
			var tags []types.Tag
			for k, v := range ar.Tags {
				tag := types.Tag{
					Key:   aws.String(k),
					Value: aws.String(v),
				}
				tags = append(tags, tag)
			}

			opts.Tags = tags
		}

		if len(ar.TransitiveTagKeys) > 0 {
			opts.TransitiveTagKeys = ar.TransitiveTagKeys
		}
	})
	_, err := appCreds.Retrieve(ctx)
	if err != nil {
		return nil, c.NewCannotAssumeRoleError(err)
	}
	return aws.NewCredentialsCache(appCreds), nil
}
