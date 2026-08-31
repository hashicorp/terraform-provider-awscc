// Copyright IBM Corp. 2021, 2026
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
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
	"golang.org/x/sync/errgroup"
)

// discoverRegion is forced regardless of the ambient AWS_REGION: the CloudFormation
// registry is queried from a single reference region for reproducibility.
const discoverRegion = "us-east-1"

// discoverConcurrency bounds the number of in-flight DescribeType calls.
// DescribeType is throttled by AWS, so the SDK client retries with backoff;
// concurrency trades a higher (bounded) request rate for wall-clock time. The
// speedup over the serial crawl is real but limited by the account's DescribeType
// throttle ceiling, so a much larger value mostly adds retries, not throughput.
const discoverConcurrency = 10

// discovered pairs a resource row with the sanitized CloudFormation schema bytes
// captured in the same DescribeType call (§10 "one crawl"). err records a
// per-type failure without aborting the whole crawl; schema is nil on failure.
type discovered struct {
	row    resourceRow
	schema []byte
	err    error
}

// discover queries AWS for the provisionable public resource types
// (FULLY_MUTABLE + IMMUTABLE, LIVE) and, in one bounded-concurrency sweep,
// derives per-type metadata (Terraform name, plural-DS support) and captures the
// sanitized schema bytes. This is the "what's new" half of the mandate plus the
// cache input, fetched once.
func discover(ctx context.Context) ([]discovered, error) {
	conn, err := newCFNClient(ctx)
	if err != nil {
		return nil, err
	}

	stepf("Listing available CloudFormation resource types (%s)…", discoverRegion)
	names, err := listTypeNames(ctx, conn)
	if err != nil {
		return nil, err
	}

	stepf("Describing %d types (concurrency %d, DescribeType is throttled)…", len(names), discoverConcurrency)
	bar := newBar(len(names), "describe")

	// Fan out DescribeType across a bounded worker pool. Each goroutine writes
	// its own slot, so no mutex is needed; per-type errors are captured in-band.
	results := make([]discovered, len(names))
	var g errgroup.Group
	g.SetLimit(discoverConcurrency)
	for i, cfn := range names {
		g.Go(func() error {
			results[i] = describeOne(ctx, conn, cfn)
			_ = bar.Add(1)
			return nil
		})
	}
	_ = g.Wait() // goroutines never return an error; failures are in-band.
	_ = bar.Finish()

	var failures int
	for _, d := range results {
		if d.err != nil {
			failures++
		}
	}
	if failures > 0 {
		fmt.Fprintf(os.Stderr, "bigdiffer: discovery: %d of %d types failed to describe (see per-type warnings)\n", failures, len(results))
	}
	return results, nil
}

// newCFNClient builds a CloudFormation client pinned to the reference region with
// an aggressive retryer so concurrent DescribeType calls ride through throttling.
func newCFNClient(ctx context.Context) (*cloudformation.Client, error) {
	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(discoverRegion))
	if err != nil {
		return nil, fmt.Errorf("loading AWS config: %w", err)
	}
	return cloudformation.NewFromConfig(cfg, func(o *cloudformation.Options) {
		o.Retryer = retry.NewStandard(func(so *retry.StandardOptions) {
			so.MaxAttempts = 25
			so.RateLimiter = ratelimit.None
		})
	}), nil
}

// listTypeNames pages ListTypes for both provisionable visibilities and returns
// the AWS-owned CloudFormation type names, deduplicated and sorted.
func listTypeNames(ctx context.Context, conn *cloudformation.Client) ([]string, error) {
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
	return awsTypeNames(summaries), nil
}

// awsTypeNames filters type summaries to AWS-owned types, deduplicates, and
// sorts. Pure: unit-tested without AWS.
func awsTypeNames(summaries []types.TypeSummary) []string {
	seen := make(map[string]bool)
	var names []string
	for _, s := range summaries {
		name := aws.ToString(s.TypeName)
		if name == "" {
			continue
		}
		if org, _, _, err := naming.ParseCloudFormationTypeName(name); err != nil || org != naming.OrganizationNameAWS {
			continue
		}
		if !seen[name] {
			seen[name] = true
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return names
}

// describeOne fetches and sanitizes one type's schema and derives its row. A
// failure is returned in-band (discovered.err) so the crawl continues.
func describeOne(ctx context.Context, conn *cloudformation.Client, cfn string) discovered {
	org, svc, res, err := naming.ParseCloudFormationTypeName(cfn)
	if err != nil {
		return discovered{row: resourceRow{CloudFormationTypeName: cfn}, err: fmt.Errorf("parsing %q: %w", cfn, err)}
	}
	label := strings.ToLower(org) + "_" + strings.ToLower(svc) + "_" + naming.CloudFormationPropertyToTerraformAttribute(res)

	schema, pluralSupported, err := describeSchema(ctx, conn, cfn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bigdiffer: %v\n", err)
		return discovered{
			row: resourceRow{ResourceTypeName: label, CloudFormationTypeName: cfn},
			err: err,
		}
	}
	return discovered{
		row: resourceRow{
			ResourceTypeName:                   label,
			CloudFormationTypeName:             cfn,
			SuppressPluralDataSourceGeneration: !pluralSupported,
		},
		schema: schema,
	}
}

// describeSchema returns the sanitized schema bytes and whether a plural data
// source is supported (a list handler with no required arguments).
func describeSchema(ctx context.Context, conn *cloudformation.Client, cfn string) ([]byte, bool, error) {
	out, err := conn.DescribeType(ctx, &cloudformation.DescribeTypeInput{
		Type:     types.RegistryTypeResource,
		TypeName: aws.String(cfn),
	})
	if err != nil {
		return nil, false, fmt.Errorf("describing %s: %w", cfn, err)
	}
	sanitized, err := cfschema.Sanitize(aws.ToString(out.Schema))
	if err != nil {
		return nil, false, fmt.Errorf("sanitizing %s: %w", cfn, err)
	}
	return []byte(sanitized), pluralSupported(sanitized), nil
}

// pluralSupported reports whether the sanitized schema has a list handler with no
// required arguments. Any parse error is treated as "no plural data source."
func pluralSupported(sanitized string) bool {
	doc, err := cfschema.NewResourceJsonSchemaDocument(sanitized)
	if err != nil {
		return false
	}
	resource, err := doc.Resource()
	if err != nil {
		return false
	}
	handler, ok := resource.Handlers[cfschema.HandlerTypeList]
	if !ok {
		return false
	}
	hs := handler.HandlerSchema
	return hs == nil || (len(hs.AllOf) == 0 && len(hs.AnyOf) == 0 && len(hs.OneOf) == 0 && len(hs.Required) == 0)
}
