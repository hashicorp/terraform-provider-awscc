// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func ListResourcesByTypeName(ctx context.Context, conn *cloudcontrol.Client, roleARN, typeName string) ([]types.ResourceDescription, error) {
	tflog.Debug(ctx, "ListResourcesByTypeName", map[string]interface{}{
		"cfTypeName": typeName,
	})

	var resourceDescriptions []types.ResourceDescription
	input := &cloudcontrol.ListResourcesInput{
		TypeName: aws.String(typeName),
	}

	if roleARN != "" {
		input.RoleArn = aws.String(roleARN)
	}

	paginator := cloudcontrol.NewListResourcesPaginator(conn, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		resourceDescriptions = append(resourceDescriptions, page.ResourceDescriptions...)
	}

	if len(resourceDescriptions) == 0 {
		return nil, &tfresource.NotFoundError{Message: "Empty result"}
	}

	return resourceDescriptions, nil
}
