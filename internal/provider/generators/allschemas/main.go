package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/mitchellh/cli"
)

const (
	PluralDataSource   = "plural"
	SingularDataSource = "singular"
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

	client := cloudformation.NewFromConfig(cfg)

	input := &cloudformation.ListTypesInput{
		Filters: &types.TypeFilters{
			Category: types.CategoryAwsTypes,
		},
		ProvisioningType: types.ProvisioningTypeFullyMutable,
		Visibility:       types.VisibilityPublic,
	}
	var typeSummaries []types.TypeSummary
	for {
		output, err := client.ListTypes(ctx, input)

		if err != nil {
			ui.Error(fmt.Sprintf("error listing fully-mutable CloudFormation types: %s", err))
			os.Exit(1)
		}

		typeSummaries = append(typeSummaries, output.TypeSummaries...)

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
	}

	input = &cloudformation.ListTypesInput{
		Filters: &types.TypeFilters{
			Category: types.CategoryAwsTypes,
		},
		ProvisioningType: types.ProvisioningTypeImmutable,
		Visibility:       types.VisibilityPublic,
	}
	for {
		output, err := client.ListTypes(ctx, input)

		if err != nil {
			ui.Error(fmt.Sprintf("error listing immutable CloudFormation types: %s", err))
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
		cfTypeNames = append(cfTypeNames, aws.ToString(typeSummary.TypeName))
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

		block := fmt.Sprintf(`
resource_schema %[1]q {
  cloudformation_type_name = %[2]q
}`, tfTypeName, cfTypeName)
		ui.Output(block)

	}

	// Data Sources
	for _, cfTypeName := range cfTypeNames {
		org, svc, res, err := naming.ParseCloudFormationTypeName(cfTypeName)

		if err != nil {
			ui.Error(fmt.Sprintf("error parsing CloudFormation type name (%s): %s", cfTypeName, err))
			os.Exit(1)
		}

		tfTypeName := strings.Join([]string{strings.ToLower(org), strings.ToLower(svc), naming.CloudFormationPropertyToTerraformAttribute(res)}, "_")

		types := []string{SingularDataSource}

		// Plural Data Source (if supported)
		input := &cloudformation.ListResourcesInput{
			TypeName: aws.String(cfTypeName),
		}

		if _, err = client.ListResources(ctx, input); err == nil {
			types = append(types, PluralDataSource)
		}

		dsBlock := fmt.Sprintf(`
data_schema %[1]q {
  cloudformation_type_name = %[2]q
  data_source_types        = %[3]s
}`, tfTypeName, cfTypeName, strings.Replace(fmt.Sprintf("%q", types), " ", ",", -1))
		ui.Output(dsBlock)
	}
}
