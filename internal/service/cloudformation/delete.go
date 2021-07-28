package cloudformation

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation/waiter"
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
	progressEvent, err := waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.ToString(output.ProgressEvent.RequestToken), 5*time.Minute)

	if progressEvent != nil && progressEvent.ErrorCode == types.HandlerErrorCodeNotFound {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
