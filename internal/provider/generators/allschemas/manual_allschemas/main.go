// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/ratelimit"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
	tfslices "github.com/hashicorp/terraform-provider-awscc/internal/slices"
)

func main() {
	g := common.NewGenerator()
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		g.Fatalf("loading AWS SDK config: %s", err)
	}

	conn := cloudformation.NewFromConfig(cfg, func(o *cloudformation.Options) {
		o.Retryer = retry.NewStandard(func(so *retry.StandardOptions) {
			so.MaxAttempts = 25
			so.RateLimiter = ratelimit.None
		})
	})

	var typeSummaries []types.TypeSummary

	for _, input := range []*cloudformation.ListTypesInput{
		{
			DeprecatedStatus: types.DeprecatedStatusLive,
			ProvisioningType: types.ProvisioningTypeFullyMutable,
			Visibility:       types.VisibilityPublic,
		},
		{
			DeprecatedStatus: types.DeprecatedStatusLive,
			ProvisioningType: types.ProvisioningTypeImmutable,
			Visibility:       types.VisibilityPublic,
		},
	} {
		pages := cloudformation.NewListTypesPaginator(conn, input)
		for pages.HasMorePages() {
			page, err := pages.NextPage(ctx)

			if err != nil {
				g.Fatalf("listing CloudFormation types: %s", err)
			}

			typeSummaries = append(typeSummaries, page.TypeSummaries...)
		}
	}

	var cfTypeNames []string
	for _, typeSummary := range typeSummaries {
		typeName := aws.ToString(typeSummary.TypeName)
		org, _, _, err := naming.ParseCloudFormationTypeName(typeName)

		if err == nil && org != naming.OrganizationNameAWS {
			continue
		}

		cfTypeNames = tfslices.AppendUnique(cfTypeNames, typeName)
	}
	sort.Strings(cfTypeNames)

	g.Infof("# %d CloudFormation resource types schemas are available for use with the Cloud Control API.", len(cfTypeNames))

	for _, cfTypeName := range cfTypeNames {
		org, svc, res, err := naming.ParseCloudFormationTypeName(cfTypeName)

		if err != nil {
			g.Fatalf("parsing CloudFormation type name (%s): %s", cfTypeName, err)
		}

		tfTypeName := strings.Join([]string{strings.ToLower(org), strings.ToLower(svc), naming.CloudFormationPropertyToTerraformAttribute(res)}, "_")

		// Determine Plural Data Source (if supported)
		suppressPluralDataSource := true
		input := &cloudformation.DescribeTypeInput{
			Type:     types.RegistryTypeResource,
			TypeName: aws.String(cfTypeName),
		}

		output, err := conn.DescribeType(ctx, input)

		if err != nil {
			g.Errorf("describing CloudFormation type (%s): %s", cfTypeName, err)
		} else {
			schema, err := cfschema.Sanitize(aws.ToString(output.Schema))

			if err != nil {
				g.Errorf("sanitizing CloudFormation type (%s) schema: %s", cfTypeName, err)
			} else {
				resourceSchema, err := cfschema.NewResourceJsonSchemaDocument(schema)

				if err != nil {
					g.Errorf("parsing CloudFormation type (%s) schema: %s", cfTypeName, err)
				} else {
					resource, err := resourceSchema.Resource()

					if err != nil {
						g.Errorf("parsing CloudFormation type (%s) resource schema: %s", cfTypeName, err)
					} else {
						if handler, ok := resource.Handlers[cfschema.HandlerTypeList]; ok {
							// Ensure no required arguments.
							if handlerSchema := handler.HandlerSchema; handlerSchema == nil ||
								(len(handlerSchema.AllOf) == 0 && len(handlerSchema.AnyOf) == 0 && len(handlerSchema.OneOf) == 0 && len(handlerSchema.Required) == 0) {
								suppressPluralDataSource = false
							}
						}
					}
				}
			}
		}

		var block string
		if suppressPluralDataSource {
			block = fmt.Sprintf(`
resource_schema %[1]q {
  cloudformation_type_name               = %[2]q
  suppress_plural_data_source_generation = %[3]t
}`, tfTypeName, cfTypeName, suppressPluralDataSource)
		} else {
			block = fmt.Sprintf(`
resource_schema %[1]q {
  cloudformation_type_name = %[2]q
}`, tfTypeName, cfTypeName)
		}
		g.Infof("%s", block)
	}
}
