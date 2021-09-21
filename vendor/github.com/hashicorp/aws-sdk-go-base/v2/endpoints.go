package awsbase

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func endpointResolver(c *Config) aws.EndpointResolver {
	resolver := func(service, region string) (aws.Endpoint, error) {
		log.Printf("[DEBUG] Resolving endpoint for %q in %q", service, region)
		switch service {
		case iam.ServiceID:
			if endpoint := c.IamEndpoint; endpoint != "" {
				log.Printf("[INFO] Setting custom IAM endpoint: %s", endpoint)
				return aws.Endpoint{
					URL:    endpoint,
					Source: aws.EndpointSourceCustom,
				}, nil
			}
		case sts.ServiceID:
			if endpoint := c.StsEndpoint; endpoint != "" {
				log.Printf("[INFO] Setting custom STS endpoint: %s", endpoint)
				return aws.Endpoint{
					URL:    endpoint,
					Source: aws.EndpointSourceCustom,
				}, nil
			}
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	}

	return aws.EndpointResolverFunc(resolver)
}
