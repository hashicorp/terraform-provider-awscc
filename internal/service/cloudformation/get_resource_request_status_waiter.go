package cloudformation

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithywaiter "github.com/aws/smithy-go/waiter"
)

//
// Based on cloudformation.TypeRegistrationCompleteWaiter.
//

// GetResourceRequestStatusAPIClient is a client that implements the
// GetResourceRequestStatus operation.
type GetResourceRequestStatusAPIClient interface {
	GetResourceRequestStatus(context.Context, *cloudformation.GetResourceRequestStatusInput, ...func(*cloudformation.Options)) (*cloudformation.GetResourceRequestStatusOutput, error)
}

var _ GetResourceRequestStatusAPIClient = (*cloudformation.Client)(nil)

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
	Retryable func(context.Context, *cloudformation.GetResourceRequestStatusInput, *cloudformation.GetResourceRequestStatusOutput, error) (bool, error)
}

// ResourceRequestStatusSuccessWaiter defines the waiters for ResourceRequestStatusSuccess
type ResourceRequestStatusSuccessWaiter struct {
	client GetResourceRequestStatusAPIClient

	options ResourceRequestStatusSuccessWaiterOptions
}

// NewResourceRequestStatusSuccessWaiter constructs a ResourceRequestStatusSuccessWaiter.
func NewResourceRequestStatusSuccessWaiter(client GetResourceRequestStatusAPIClient, optFns ...func(*ResourceRequestStatusSuccessWaiterOptions)) *ResourceRequestStatusSuccessWaiter {
	options := ResourceRequestStatusSuccessWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = resourceRequestStatusSuccessStateRetryable

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
func (w *ResourceRequestStatusSuccessWaiter) Wait(ctx context.Context, params *cloudformation.GetResourceRequestStatusInput, maxWaitDur time.Duration, optFns ...func(*ResourceRequestStatusSuccessWaiterOptions)) (*cloudformation.GetResourceRequestStatusOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v", options.MinDelay, options.MaxDelay)
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

		out, err := w.client.GetResourceRequestStatus(ctx, params, func(o *cloudformation.Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
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
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ResourceRequestStatusSuccess waiter")
}

func resourceRequestStatusSuccessStateRetryable(ctx context.Context, input *cloudformation.GetResourceRequestStatusInput, output *cloudformation.GetResourceRequestStatusOutput, err error) (bool, error) {
	if err == nil {
		switch value := output.ProgressEvent.OperationStatus; value {
		case types.OperationStatusSuccess:
			return false, nil
		case types.OperationStatusFailed:
			if output.ProgressEvent.ErrorCode == types.HandlerErrorCodeNotFound && output.ProgressEvent.Operation == types.OperationDelete {
				return false, nil
			}
			return false, fmt.Errorf("waiter state transitioned to %s", value)
		}
	}

	return true, nil
}
