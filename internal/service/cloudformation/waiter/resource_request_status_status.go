package waiter

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func ResourceRequestStatusProgressEventOperationStatus(ctx context.Context, conn *cloudformation.Client, requestToken string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		input := &cloudformation.GetResourceRequestStatusInput{
			RequestToken: aws.String(requestToken),
		}

		output, err := conn.GetResourceRequestStatus(ctx, input)

		if err != nil {
			return nil, "", err
		}

		if output == nil || output.ProgressEvent == nil {
			return nil, "", nil
		}

		return output.ProgressEvent, string(output.ProgressEvent.OperationStatus), nil
	}
}
