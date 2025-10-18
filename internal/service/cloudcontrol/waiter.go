// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
)

// RetryGetResourceRequestStatus returns a custom retryable function for the GetResourceRequestStatus operation.
func RetryGetResourceRequestStatus(pProgressEvent **types.ProgressEvent) func(context.Context, *cloudcontrol.GetResourceRequestStatusInput, *cloudcontrol.GetResourceRequestStatusOutput, error) (bool, error) {
	return func(ctx context.Context, input *cloudcontrol.GetResourceRequestStatusInput, output *cloudcontrol.GetResourceRequestStatusOutput, err error) (bool, error) {
		if err == nil {
			progressEvent := output.ProgressEvent
			if pProgressEvent != nil {
				*pProgressEvent = progressEvent
			}

			switch value := progressEvent.OperationStatus; value {
			case types.OperationStatusSuccess, types.OperationStatusCancelComplete:
				return false, nil

			case types.OperationStatusFailed:
				if progressEvent.ErrorCode == types.HandlerErrorCodeNotFound && progressEvent.Operation == types.OperationDelete {
					// Resource not found error on delete is OK.
					return false, nil
				}

				// Build enhanced error message with hook information
				errorMsg := fmt.Sprintf("waiter state transitioned to %s. StatusMessage: %s",
					value, aws.ToString(progressEvent.StatusMessage))
				
				if progressEvent.ErrorCode != "" {
					errorMsg += fmt.Sprintf(". ErrorCode: %s", string(progressEvent.ErrorCode))
				}

				// Add hook information if available
				if len(output.HooksProgressEvent) > 0 {
					var hookErrors []string
					for _, hookEvent := range output.HooksProgressEvent {
						hookStatus := aws.ToString(hookEvent.HookStatus)
						// HOOK_COMPLETE_FAILED: The Hook invocation is complete with a failed result.
						// HOOK_FAILED: The Hook invocation didn't complete successfully.
						if hookStatus == "HOOK_COMPLETE_FAILED" || hookStatus == "HOOK_FAILED" {
							hookErrors = append(hookErrors, fmt.Sprintf("HookName: %s, HookArn: %s, HookVersion: %s, Time: %s, HookMessage: %s",
								aws.ToString(hookEvent.HookTypeName),
								aws.ToString(hookEvent.HookTypeArn),
								aws.ToString(hookEvent.HookTypeVersionId),
								hookEvent.HookEventTime.Format("2006-01-02T15:04:05Z"),
								aws.ToString(hookEvent.HookStatusMessage)))
						}
					}
					if len(hookErrors) > 0 {
						errorMsg += ". Hook failures: " + strings.Join(hookErrors, "; ")
					}
				}

				return false, fmt.Errorf("%s", errorMsg)
			}
		}

		return true, err
	}
}
