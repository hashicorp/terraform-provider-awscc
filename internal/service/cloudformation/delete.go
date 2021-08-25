package cloudformation

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func DeleteResource(ctx context.Context, conn *cloudformation.Client, roleARN, typeName, id string, maxWaitTime time.Duration) error {
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

	_, err = WaitForResourceRequestSuccess(ctx, conn, aws.ToString(output.ProgressEvent.RequestToken), maxWaitTime)

	if err != nil {
		return err
	}

	return nil
}
