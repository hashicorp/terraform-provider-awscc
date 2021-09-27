package cloudcontrol

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Provider is the interface implemented by AWS Cloud Control API client providers.
type Provider interface {
	// CloudControlApiClient returns an AWS Cloud Control API client.
	CloudControlApiClient(context.Context) *cloudcontrol.Client

	// Region returns and AWS Cloud Control API client's region
	Region(ctx context.Context) string

	// RoleARN returns an AWS Cloud Control API service role ARN.
	RoleARN(context.Context) string
}
