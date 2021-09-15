package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Provider is the interface implemented by AWS Cloud Control client providers.
type Provider interface {
	// CloudControlClient returns an AWS Cloud Control client.
	CloudControlClient(context.Context) *cloudcontrol.Client

	// Region returns and AWS Cloud Control client's region
	Region(ctx context.Context) string

	// RoleARN returns an AWS Cloud Control service role ARN.
	RoleARN(context.Context) string
}
