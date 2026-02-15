// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	cctypes_sdk "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
)

func TestThrottlingRetry(t *testing.T) {
	tests := []struct {
		name          string
		err           error
		expectRetry   aws.Ternary
	}{
		{
			name:        "ThrottlingException should retry",
			err:         &cctypes_sdk.ThrottlingException{Message: aws.String("Rate exceeded")},
			expectRetry: aws.TrueTernary,
		},
		{
			name:        "Non-throttling error should not retry",
			err:         &cctypes_sdk.InvalidRequestException{Message: aws.String("Invalid request")},
			expectRetry: aws.UnknownTernary,
		},
		{
			name:        "Generic error should not retry",
			err:         errors.New("some other error"),
			expectRetry: aws.UnknownTernary,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var throttlingErr *cctypes_sdk.ThrottlingException
			result := aws.UnknownTernary
			if errors.As(tt.err, &throttlingErr) {
				result = aws.TrueTernary
			}

			if result != tt.expectRetry {
				t.Errorf("Expected retry=%v, got %v", tt.expectRetry, result)
			}
		})
	}
}
