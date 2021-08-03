package cloudformation

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/tfresource"
)

func DeleteResource(ctx context.Context, conn *cloudformation.Client, roleARN, typeName, id string) error {
	tflog.Debug(ctx, "DeleteResource", "cfTypeName", typeName, "id", id)

	input := &cloudformation.DeleteResourceInput{
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

	// TODO
	// TODO How long to wait for?
	// TODO
	maxWaitTime := 5 * time.Minute
	waiter := NewResourceRequestStatusSuccessWaiter(conn)

	err = waiter.Wait(ctx, &cloudformation.GetResourceRequestStatusInput{RequestToken: output.ProgressEvent.RequestToken}, maxWaitTime)

	// TODO How is HandlerErrorCodeNotFound propagated?
	// if progressEvent != nil && progressEvent.ErrorCode == types.HandlerErrorCodeNotFound {
	// 	return nil
	// }

	if err != nil {
		return err
	}

	return nil
}
