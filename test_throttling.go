package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
)

func main() {
	// Test waiter retry logic
	retryFunc := tfcloudcontrol.RetryGetResourceRequestStatus(nil)
	
	// Simulate throttling error
	output := &cloudcontrol.GetResourceRequestStatusOutput{
		ProgressEvent: &types.ProgressEvent{
			OperationStatus: types.OperationStatusFailed,
			ErrorCode:       "Throttling",
			StatusMessage:   aws.String("Rate exceeded for operation 'CREATE'"),
		},
	}
	
	shouldRetry, err := retryFunc(context.Background(), nil, output, nil)
	
	fmt.Printf("Throttling error - Should retry: %v, Error: %v\n", shouldRetry, err)
	
	// Test non-throttling error
	output2 := &cloudcontrol.GetResourceRequestStatusOutput{
		ProgressEvent: &types.ProgressEvent{
			OperationStatus: types.OperationStatusFailed,
			ErrorCode:       "ValidationError",
			StatusMessage:   aws.String("Invalid parameter"),
		},
	}
	
	shouldRetry2, err2 := retryFunc(context.Background(), nil, output2, nil)
	
	fmt.Printf("Validation error - Should retry: %v, Error: %v\n", shouldRetry2, err2)
}
