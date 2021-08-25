package awsbase

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	multierror "github.com/hashicorp/go-multierror"
)

func credentialsProvider(c *Config) (aws.CredentialsProvider, error) {
	var providers []aws.CredentialsProvider
	if c.AccessKey != "" {
		providers = append(providers,
			credentials.NewStaticCredentialsProvider(
				c.AccessKey,
				c.SecretKey,
				c.Token,
			))
	}
	if len(providers) == 0 {
		return nil, nil
	}

	return aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		var errs *multierror.Error
		for _, p := range providers {
			creds, err := p.Retrieve(ctx)
			if err == nil {
				return creds, nil
			}
			errs = multierror.Append(errs, err)
		}

		return aws.Credentials{}, fmt.Errorf("No valid providers found: %w", errs)
	}), nil
}
