package cloudcontrol

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func ListResourcesByTypeName(ctx context.Context, conn *cloudcontrol.Client, roleARN, typeName string) ([]types.ResourceDescription, error) {
	log.Printf("[DEBUG] ListResourcesByTypeName. cfTypeName: %s", typeName)

	input := &cloudcontrol.ListResourcesInput{
		TypeName: aws.String(typeName),
	}

	if roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.ListResources(ctx, input)

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, &tfresource.NotFoundError{Message: "Empty result"}
	}

	return output.ResourceDescriptions, nil
}
