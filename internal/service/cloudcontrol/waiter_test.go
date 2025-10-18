// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
)

func TestRetryGetResourceRequestStatus_WithHookFailures(t *testing.T) {
	retryFunc := RetryGetResourceRequestStatus(nil)
	testTime := time.Date(2025, 10, 18, 11, 10, 45, 0, time.UTC)

	output := &cloudcontrol.GetResourceRequestStatusOutput{
		ProgressEvent: &types.ProgressEvent{
			OperationStatus: types.OperationStatusFailed,
			StatusMessage:   aws.String("test-token"),
		},
		HooksProgressEvent: []types.HookProgressEvent{
			{
				HookTypeName:        aws.String("Private::Security::S3PublicAccessBlock"),
				HookTypeArn:         aws.String("arn:aws:cloudformation:us-east-1:123456789012:type/hook/test"),
				HookTypeVersionId:   aws.String("00000001"),
				HookStatus:          aws.String("HOOK_FAILED"),
				HookStatusMessage:   aws.String("Unable to assume role"),
				HookEventTime:       &testTime,
			},
		},
	}

	_, err := retryFunc(context.Background(), nil, output, nil)

	if err == nil {
		t.Fatal("Expected error but got nil")
	}

	errorMsg := err.Error()
	expectedParts := []string{
		"waiter state transitioned to FAILED",
		"Hook failures:",
		"Private::Security::S3PublicAccessBlock",
		"arn:aws:cloudformation:us-east-1:123456789012:type/hook/test",
		"v00000001",
		"at 2025-10-18T11:10:45Z",
		"Unable to assume role",
	}

	for _, part := range expectedParts {
		if !strings.Contains(errorMsg, part) {
			t.Errorf("Error message missing expected part %q. Got: %s", part, errorMsg)
		}
	}
}

func TestRetryGetResourceRequestStatus_WithHookCompleteFailed(t *testing.T) {
	retryFunc := RetryGetResourceRequestStatus(nil)
	testTime := time.Date(2025, 10, 18, 11, 15, 30, 0, time.UTC)

	output := &cloudcontrol.GetResourceRequestStatusOutput{
		ProgressEvent: &types.ProgressEvent{
			OperationStatus: types.OperationStatusFailed,
			StatusMessage:   aws.String("test-token"),
		},
		HooksProgressEvent: []types.HookProgressEvent{
			{
				HookTypeName:        aws.String("Private::Security::S3PublicAccessBlock"),
				HookTypeArn:         aws.String("arn:aws:cloudformation:us-east-1:123456789012:type/hook/test"),
				HookTypeVersionId:   aws.String("00000002"),
				HookStatus:          aws.String("HOOK_COMPLETE_FAILED"),
				HookStatusMessage:   aws.String("Template failed validation, the following rule(s) failed: s3-security-rules.guard/default."),
				HookEventTime:       &testTime,
			},
		},
	}

	_, err := retryFunc(context.Background(), nil, output, nil)

	if err == nil {
		t.Fatal("Expected error but got nil")
	}

	errorMsg := err.Error()
	expectedParts := []string{
		"waiter state transitioned to FAILED",
		"Hook failures:",
		"Private::Security::S3PublicAccessBlock",
		"arn:aws:cloudformation:us-east-1:123456789012:type/hook/test",
		"v00000002",
		"at 2025-10-18T11:15:30Z",
		"Template failed validation",
	}

	for _, part := range expectedParts {
		if !strings.Contains(errorMsg, part) {
			t.Errorf("Error message missing expected part %q. Got: %s", part, errorMsg)
		}
	}
}
