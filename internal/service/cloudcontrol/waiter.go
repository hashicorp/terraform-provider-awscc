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

				// Attempt to retrieve hook results
				hookResultsMsg := ""
				if output.HooksProgressEvent != nil {
					var details []string
					for _, hook := range output.HooksProgressEvent {
						details = append(details, fmt.Sprintf("StatusReason: %s", aws.ToString(hook.HookStatusMessage)))
					}
					hookResults := strings.Join(details, "\n")
					if hookResults != "" {
						hookResultsMsg = fmt.Sprintf(" Hook results: \n%s", hookResults)
					}
					return false, fmt.Errorf("%s", hookResultsMsg)
				}

				return false, fmt.Errorf("waiter state transitioned to %s. StatusMessage: %s. ErrorCode: %s", value, aws.ToString(progressEvent.StatusMessage), progressEvent.ErrorCode)
			}
		}
		return true, err
	}
}
