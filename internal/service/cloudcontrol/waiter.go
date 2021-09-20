package cloudformation

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithywaiter "github.com/aws/smithy-go/waiter"
)

const (
	waitForResourceRequestMinDelay = 5 * time.Second
)

// WaitForResourceRequestSuccess waits up to maxWaitTime for the specified reource request to succeed.
func WaitForResourceRequestSuccess(ctx context.Context, client *cloudcontrol.Client, requestToken string, maxWaitTime time.Duration) (string, error) {
	var progressEvent *types.ProgressEvent

	waiter := NewResourceRequestStatusSuccessWaiter(client, func(o *ResourceRequestStatusSuccessWaiterOptions) {
		o.MinDelay = waitForResourceRequestMinDelay

		o.Retryable = func(ctx context.Context, input *cloudcontrol.GetResourceRequestStatusInput, output *cloudcontrol.GetResourceRequestStatusOutput, err error) (bool, error) {
			if err == nil {
				progressEvent = output.ProgressEvent
				switch value := progressEvent.OperationStatus; value {
				case types.OperationStatusSuccess:
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
	})

	err := waiter.Wait(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: aws.String(requestToken)}, maxWaitTime)

	if err != nil {
		return "", err
	}

	return aws.ToString(progressEvent.Identifier), nil
}

//
// Based on cloudformation.TypeRegistrationCompleteWaiter.
// Should be deleted when an SDK generated waiter is available.
//

// GetResourceRequestStatusAPIClient is a client that implements the
// GetResourceRequestStatus operation.
type GetResourceRequestStatusAPIClient interface {
	GetResourceRequestStatus(context.Context, *cloudcontrol.GetResourceRequestStatusInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error)
}

var _ GetResourceRequestStatusAPIClient = (*cloudcontrol.Client)(nil)

// ResourceRequestStatusSuccessWaiterOptions are waiter options for
// ResourceRequestStatusSuccessWaiter
type ResourceRequestStatusSuccessWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ResourceRequestStatusSuccessWaiter will use default minimum delay of 30 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ResourceRequestStatusSuccessWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *cloudcontrol.GetResourceRequestStatusInput, *cloudcontrol.GetResourceRequestStatusOutput, error) (bool, error)
}

const (
	resourceRequestStatusSuccessWaiterDefaultMinDelay = 30 * time.Second
	resourceRequestStatusSuccessWaiterDefaultMaxDelay = 120 * time.Second
)

// ResourceRequestStatusSuccessWaiter defines the waiters for ResourceRequestStatusSuccess
type ResourceRequestStatusSuccessWaiter struct {
	client GetResourceRequestStatusAPIClient

	options ResourceRequestStatusSuccessWaiterOptions
}

// NewResourceRequestStatusSuccessWaiter constructs a ResourceRequestStatusSuccessWaiter.
func NewResourceRequestStatusSuccessWaiter(client GetResourceRequestStatusAPIClient, optFns ...func(*ResourceRequestStatusSuccessWaiterOptions)) *ResourceRequestStatusSuccessWaiter {
	options := ResourceRequestStatusSuccessWaiterOptions{
		MinDelay: resourceRequestStatusSuccessWaiterDefaultMinDelay,
		MaxDelay: resourceRequestStatusSuccessWaiterDefaultMaxDelay,
		Retryable: func(context.Context, *cloudcontrol.GetResourceRequestStatusInput, *cloudcontrol.GetResourceRequestStatusOutput, error) (bool, error) {
			return false, fmt.Errorf("retryable not implemented")
		},
	}

	for _, fn := range optFns {
		fn(&options)
	}
	return &ResourceRequestStatusSuccessWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ResourceRequestStatusSuccess waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ResourceRequestStatusSuccessWaiter) Wait(ctx context.Context, params *cloudcontrol.GetResourceRequestStatusInput, maxWaitDur time.Duration, optFns ...func(*ResourceRequestStatusSuccessWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = resourceRequestStatusSuccessWaiterDefaultMaxDelay
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetResourceRequestStatus(ctx, params, func(o *cloudcontrol.Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ResourceRequestStatusSuccess waiter")
}
