package awsbase

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func endpointResolver(c *Config) aws.EndpointResolver {
	resolver := func(service, region string) (aws.Endpoint, error) {
		switch service {
		// case ec2metadata.ServiceName:
		// 	if endpoint := os.Getenv("AWS_METADATA_URL"); endpoint != "" {
		// 		log.Printf("[INFO] Setting custom EC2 metadata endpoint: %s", endpoint)
		// 		resolvedEndpoint.URL = endpoint
		// 	}
		// case iam.ServiceName:
		// 	if endpoint := c.IamEndpoint; endpoint != "" {
		// 		log.Printf("[INFO] Setting custom IAM endpoint: %s", endpoint)
		// 		resolvedEndpoint.URL = endpoint
		// 	}
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
