// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"

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
				waiterErr := newWaiterErr(string(value), aws.ToString(progressEvent.StatusMessage))
				waiterErr.withErrorCode(string(progressEvent.ErrorCode))

				// Add hook information if available
				if len(output.HooksProgressEvent) > 0 {
					for _, hookEvent := range output.HooksProgressEvent {
						hookStatus := aws.ToString(hookEvent.HookStatus)
						// HOOK_COMPLETE_FAILED: The Hook invocation is complete with a failed result.
						// HOOK_FAILED: The Hook invocation didn't complete successfully.
						if hookStatus == "HOOK_COMPLETE_FAILED" || hookStatus == "HOOK_FAILED" {
							waiterErr.withHookEvent(hookEvent)
						}
					}
				}

				return false, waiterErr
			}
		}

		return true, err
	}
}
