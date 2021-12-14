package cloudcontrol

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func DeleteResource(ctx context.Context, conn *cloudcontrol.Client, roleARN, typeName, id string, maxWaitTime time.Duration) error {
	log.Printf("[DEBUG] DeleteResource. cfTypeName: %s, id: %s", typeName, id)

	input := &cloudcontrol.DeleteResourceInput{
		ClientToken: aws.String(tfresource.UniqueId()),
		Identifier:  aws.String(id),
		TypeName:    aws.String(typeName),
	}

	if roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.DeleteResource(ctx, input)

	if err != nil {
		return err
	}

	if output == nil || output.ProgressEvent == nil {
		return fmt.Errorf("empty result")
	}

	waiter := cloudcontrol.NewResourceRequestSuccessWaiter(conn, func(o *cloudcontrol.ResourceRequestSuccessWaiterOptions) {
		o.Retryable = RetryGetResourceRequestStatus(nil)
	})

	err = waiter.Wait(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: output.ProgressEvent.RequestToken}, maxWaitTime)

	if err != nil {
		return err
	}

	return nil
}
