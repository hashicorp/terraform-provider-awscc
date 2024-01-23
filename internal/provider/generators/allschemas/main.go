// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		ui.Error(fmt.Sprintf("error loading AWS SDK config: %s", err))
		os.Exit(1)
	}

	ccClient := cloudcontrol.NewFromConfig(cfg)
	cfClient := cloudformation.NewFromConfig(cfg)

	input := &cloudformation.ListTypesInput{
		ProvisioningType: types.ProvisioningTypeFullyMutable,
		Visibility:       types.VisibilityPublic,
	}
	var typeSummaries []types.TypeSummary
	for {
		output, err := cfClient.ListTypes(ctx, input)

		if err != nil {
			ui.Error(fmt.Sprintf("error listing fully-mutable, public CloudFormation types: %s", err))
			os.Exit(1)
		}

		typeSummaries = append(typeSummaries, output.TypeSummaries...)

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	input = &cloudformation.ListTypesInput{
		ProvisioningType: types.ProvisioningTypeImmutable,
		Visibility:       types.VisibilityPublic,
	}
	for {
		output, err := cfClient.ListTypes(ctx, input)

		if err != nil {
			ui.Error(fmt.Sprintf("error listing immutable, public CloudFormation types: %s", err))
			os.Exit(1)
		}

		typeSummaries = append(typeSummaries, output.TypeSummaries...)

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	var cfTypeNames []string
	for _, typeSummary := range typeSummaries {
		typeName := aws.ToString(typeSummary.TypeName)
		org, _, _, err := naming.ParseCloudFormationTypeName(typeName)

		if err == nil && org != naming.OrganizationNameAWS {
			continue
		}

		cfTypeNames = append(cfTypeNames, typeName)
	}
	sort.Strings(cfTypeNames)

	ui.Output(fmt.Sprintf("# %d CloudFormation resource types schemas are available for use with the Cloud Control API.", len(cfTypeNames)))

	for _, cfTypeName := range cfTypeNames {
		org, svc, res, err := naming.ParseCloudFormationTypeName(cfTypeName)

		if err != nil {
			ui.Error(fmt.Sprintf("error parsing CloudFormation type name (%s): %s", cfTypeName, err))
			os.Exit(1)
		}

		tfTypeName := strings.Join([]string{strings.ToLower(org), strings.ToLower(svc), naming.CloudFormationPropertyToTerraformAttribute(res)}, "_")

		// Determine Plural Data Source (if supported)
		input := &cloudcontrol.ListResourcesInput{
			TypeName: aws.String(cfTypeName),
		}

		var suppressPluralDataSource bool
		if _, err = ccClient.ListResources(ctx, input); err != nil {
			suppressPluralDataSource = true
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
		ui.Output(block)
	}
}
