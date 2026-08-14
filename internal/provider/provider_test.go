// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
)

func TestThrottlingErrRetryableFunc(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		expectRetry aws.Ternary
	}{
		{
			name:        "ThrottlingException should retry",
			err:         &cctypes.ThrottlingException{Message: aws.String("Rate exceeded")},
			expectRetry: aws.TrueTernary,
		},
		{
			name:        "Non-throttling error should not retry",
			err:         &cctypes.InvalidRequestException{Message: aws.String("Invalid request")},
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
			result := ThrottlingErrRetryableFunc(tt.err)
			if result != tt.expectRetry {
				t.Errorf("Expected retry=%v, got %v", tt.expectRetry, result)
			}
		})
	}
}
