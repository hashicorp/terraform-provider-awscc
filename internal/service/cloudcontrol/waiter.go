package cloudformation

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
)

// RetryGetResourceRequestStatus is a custom retryable function for the GetResourceRequestStatus operation.
func RetryGetResourceRequestStatus(ctx context.Context, input *cloudcontrol.GetResourceRequestStatusInput, output *cloudcontrol.GetResourceRequestStatusOutput, err error) (bool, error) {
	if err == nil {
		progressEvent := output.ProgressEvent
		switch value := progressEvent.OperationStatus; value {
		case types.OperationStatusSuccess, types.OperationStatusCancelComplete:
			return false, nil

		case types.OperationStatusFailed:
			if progressEvent.ErrorCode == types.HandlerErrorCodeNotFound && progressEvent.Operation == types.OperationDelete {
				// Resource not found error on delete is OK.
				return false, nil
			}

			return false, fmt.Errorf("waiter state transitioned to %s. StatusMessage: %s. ErrorCode: %s", value, aws.ToString(progressEvent.StatusMessage), progressEvent.ErrorCode)
		}
	}

	return true, nil
}
