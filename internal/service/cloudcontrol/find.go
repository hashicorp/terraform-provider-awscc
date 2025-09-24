// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-awscc/internal/errs"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func FindResourceByTypeNameAndID(ctx context.Context, conn *cloudcontrol.Client, roleARN, typeName, id string) (*types.ResourceDescription, error) {
	tflog.Debug(ctx, "FindResourceByTypeNameAndID", map[string]interface{}{
		"cfTypeName": typeName,
		"id":         id,
	})

	input := &cloudcontrol.GetResourceInput{
		Identifier: aws.String(id),
		TypeName:   aws.String(typeName),
	}

	if roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	output, err := conn.GetResource(ctx, input)

	if errs.IsA[*types.ResourceNotFoundException](err) {
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

	tflog.Debug(ctx, "ResourceDescription.ResourceModel", map[string]interface{}{
		"value": aws.ToString(output.ResourceDescription.Properties),
	})

	return output.ResourceDescription, nil
}
