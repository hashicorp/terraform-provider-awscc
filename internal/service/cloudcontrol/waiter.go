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

				statusMessage := aws.ToString(progressEvent.StatusMessage)

				// Attempt to retrieve hook results
				if output.HooksProgressEvent != nil {
					var details []string
					for _, hook := range output.HooksProgressEvent {
						if *hook.HookStatus == "HOOK_COMPLETE_FAILED" {
							details = append(details, fmt.Sprintf("Reason: %s", aws.ToString(hook.HookStatusMessage)))
						}
					}
					hookResults := strings.Join(details, "\n")
					if hookResults != "" {
						statusMessage = fmt.Sprintf("Hook invocation failed: \n%s", hookResults)
					}
				}
				return false, fmt.Errorf("waiter state transitioned to %s. StatusMessage: %s. ErrorCode: %s", value, statusMessage, progressEvent.ErrorCode)
			}
		}
		return true, err
	}
}
