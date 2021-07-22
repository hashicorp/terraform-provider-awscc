package waiter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func ResourceRequestStatusProgressEventOperationStatusSuccess(ctx context.Context, conn *cloudformation.Client, requestToken string, timeout time.Duration) (*cftypes.ProgressEvent, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(cftypes.OperationStatusInProgress),
			string(cftypes.OperationStatusPending),
		},
		Target:  []string{string(cftypes.OperationStatusSuccess)},
		Refresh: ResourceRequestStatusProgressEventOperationStatus(ctx, conn, requestToken),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*cftypes.ProgressEvent); ok {
		if err != nil && output != nil {
			newErr := fmt.Errorf("%v", output)

			var te *resource.TimeoutError
			var use *resource.UnexpectedStateError
			if ok := errors.As(err, &te); ok && te.LastError == nil {
				te.LastError = newErr
			} else if ok := errors.As(err, &use); ok && use.LastError == nil {
				use.LastError = newErr
			}
		}

		return output, err
	}

	return nil, err
}
