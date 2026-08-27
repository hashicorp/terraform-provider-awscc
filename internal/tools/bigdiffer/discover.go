// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"os"
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
)

// discoverRegion is forced regardless of the ambient AWS_REGION: the CloudFormation
// registry is queried from a single reference region for reproducibility.
const discoverRegion = "us-east-1"

// discover queries AWS for the set of provisionable public resource types
// (FULLY_MUTABLE + IMMUTABLE, LIVE) and returns them as resource rows. This is
// the "what's new" half of the mandate: bigdiffer gets the list itself rather
// than diffing checked-in available_schemas.<date>.hcl snapshots.
//
// It intentionally mirrors the old manual_allschemas generator's logic but lives
// here so bigdiffer stays self-contained.
func discover(ctx context.Context) ([]resourceRow, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(discoverRegion))
	if err != nil {
		return nil, fmt.Errorf("loading AWS config: %w", err)
	}

	conn := cloudformation.NewFromConfig(cfg, func(o *cloudformation.Options) {
		o.Retryer = retry.NewStandard(func(so *retry.StandardOptions) {
			so.MaxAttempts = 25
			so.RateLimiter = ratelimit.None
		})
	})

	var summaries []types.TypeSummary
	for _, in := range []*cloudformation.ListTypesInput{
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
		pages := cloudformation.NewListTypesPaginator(conn, in)
		for pages.HasMorePages() {
			page, err := pages.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("listing CloudFormation types: %w", err)
			}
			summaries = append(summaries, page.TypeSummaries...)
		}
	}

	seen := make(map[string]bool)
	var cfnNames []string
	for _, s := range summaries {
		name := aws.ToString(s.TypeName)
		if org, _, _, err := naming.ParseCloudFormationTypeName(name); err == nil && org != naming.OrganizationNameAWS {
			continue
		}
		if !seen[name] {
			seen[name] = true
			cfnNames = append(cfnNames, name)
		}
	}
	sort.Strings(cfnNames)

	rows := make([]resourceRow, 0, len(cfnNames))
	for _, cfn := range cfnNames {
		org, svc, res, err := naming.ParseCloudFormationTypeName(cfn)
		if err != nil {
			return nil, fmt.Errorf("parsing CloudFormation type name %q: %w", cfn, err)
		}
		label := strings.ToLower(org) + "_" + strings.ToLower(svc) + "_" + naming.CloudFormationPropertyToTerraformAttribute(res)

		rows = append(rows, resourceRow{
			ResourceTypeName:                   label,
			CloudFormationTypeName:             cfn,
			SuppressPluralDataSourceGeneration: !supportsPluralDataSource(ctx, conn, cfn),
		})
	}
	return rows, nil
}

// supportsPluralDataSource reports whether the type has a List handler with no
// required arguments (the condition under which a plural data source is
// generated). Any error is treated as "no plural data source" and logged.
func supportsPluralDataSource(ctx context.Context, conn *cloudformation.Client, cfn string) bool {
	out, err := conn.DescribeType(ctx, &cloudformation.DescribeTypeInput{
		Type:     types.RegistryTypeResource,
		TypeName: aws.String(cfn),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "bigdiffer: describing %s: %v\n", cfn, err)
		return false
	}
	schema, err := cfschema.Sanitize(aws.ToString(out.Schema))
	if err != nil {
		fmt.Fprintf(os.Stderr, "bigdiffer: sanitizing %s: %v\n", cfn, err)
		return false
	}
	doc, err := cfschema.NewResourceJsonSchemaDocument(schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bigdiffer: parsing %s schema: %v\n", cfn, err)
		return false
	}
	resource, err := doc.Resource()
	if err != nil {
		fmt.Fprintf(os.Stderr, "bigdiffer: parsing %s resource: %v\n", cfn, err)
		return false
	}
	handler, ok := resource.Handlers[cfschema.HandlerTypeList]
	if !ok {
		return false
	}
	hs := handler.HandlerSchema
	return hs == nil || (len(hs.AllOf) == 0 && len(hs.AnyOf) == 0 && len(hs.OneOf) == 0 && len(hs.Required) == 0)
}
