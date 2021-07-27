package cloudformation

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/tfresource"
)

func FindResourceByTypeNameAndID(ctx context.Context, conn *cloudformation.Client, roleARN, typeName, id string) (*types.ResourceDescription, error) {
	input := &cloudformation.GetResourceInput{
		Identifier: aws.String(id),
		TypeName:   aws.String(typeName),
	}

	if roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.GetResource(ctx, input)

	if rnfe := (*types.ResourceNotFoundException)(nil); errors.As(err, &rnfe) {
		return nil, &tfresource.NotFoundError{LastError: err}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.ResourceDescription == nil {
		return nil, &tfresource.NotFoundError{Message: "Empty result"}
	}

	return output.ResourceDescription, nil
}
