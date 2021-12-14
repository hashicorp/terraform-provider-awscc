package cloudcontrol

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func FindResourceByTypeNameAndID(ctx context.Context, conn *cloudcontrol.Client, roleARN, typeName, id string) (*types.ResourceDescription, error) {
	log.Printf("[DEBUG] FindResourceByTypeNameAndID. cfTypeName: %s, id: %s", typeName, id)

	input := &cloudcontrol.GetResourceInput{
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
		// "api error ResourceNotFound: AWS::Logs::LogGroup Handler returned status FAILED: Resource of type 'AWS::Logs::LogGroup' with identifier '{"/properties/LogGroupName":"JPZD4AtrMPJQ4DJLl0EUpXRYS-7otPmOvuX3oo"}' was not found."
		if strings.Contains(err.Error(), "api error ResourceNotFound") {
			return nil, &tfresource.NotFoundError{LastError: err}
		}

		return nil, err
	}

	if output == nil || output.ResourceDescription == nil {
		return nil, &tfresource.NotFoundError{Message: "Empty result"}
	}

	log.Printf("[DEBUG] ResourceDescription.ResourceModel. value: %s", aws.ToString(output.ResourceDescription.Properties))

	return output.ResourceDescription, nil
}
