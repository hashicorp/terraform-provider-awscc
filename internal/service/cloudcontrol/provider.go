// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Provider is the interface implemented by AWS Cloud Control API client providers.
// It's role is similar to terraform-aws-provider's 'conns.AWSClient'.
type Provider interface {
	// CloudControlApiClient returns an AWS Cloud Control API client.
	CloudControlAPIClient(context.Context) *cloudcontrol.Client

	// Region returns an AWS Cloud Control API client's region
	Region(ctx context.Context) string

	// RegisterLogger places the configured logger into Context so it can be used via `tflog`.
	RegisterLogger(ctx context.Context) context.Context

	// RoleARN returns an AWS Cloud Control API service role ARN.
	RoleARN(context.Context) string
}
