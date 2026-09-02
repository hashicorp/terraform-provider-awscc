// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

func TestAWSTypeNames(t *testing.T) {
	t.Parallel()

	summary := func(name string) types.TypeSummary {
		return types.TypeSummary{TypeName: aws.String(name)}
	}

	in := []types.TypeSummary{
		summary("AWS::S3::Bucket"),
		summary("AWS::Logs::LogGroup"),
		summary("AWS::S3::Bucket"),   // duplicate (appears in both provisioning lists)
		summary("Alexa::ASK::Skill"), // non-AWS org: dropped
		summary("AWS::EC2::Instance"),
		summary(""), // empty: dropped
	}

	got := awsTypeNames(in)
	want := []string{"AWS::EC2::Instance", "AWS::Logs::LogGroup", "AWS::S3::Bucket"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("awsTypeNames() = %v, want %v", got, want)
	}
}

// TestClassifyAbsentProbe covers the design doc's classification table
// (contributing/docs/absent-row-probe-design.md) without a live AWS call: the
// four probe outcomes and the class (or lack of one) each must produce.
func TestClassifyAbsentProbe(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		out  *cloudformation.DescribeTypeOutput
		err  error
		want changeClass
	}{
		{
			name: "deprecated on success is withdrawn regardless of provisioning type",
			out: &cloudformation.DescribeTypeOutput{
				DeprecatedStatus: types.DeprecatedStatusDeprecated,
				ProvisioningType: types.ProvisioningTypeFullyMutable,
			},
			want: classWithdrawn,
		},
		{
			name: "non-provisionable, not deprecated, is non-provisionable-live",
			out: &cloudformation.DescribeTypeOutput{
				DeprecatedStatus: types.DeprecatedStatusLive,
				ProvisioningType: types.ProvisioningTypeNonProvisionable,
			},
			want: classNonProvisionable,
		},
		{
			name: "type not found is withdrawn",
			err:  &types.TypeNotFoundException{Message: aws.String("The type does not exist.")},
			want: classWithdrawn,
		},
		{
			name: "wrapped type not found is still withdrawn (errors.As through %w)",
			err:  fmt.Errorf("describing AWS::Old::Gone: %w", &types.TypeNotFoundException{Message: aws.String("gone")}),
			want: classWithdrawn,
		},
		{
			name: "transient error is no decision",
			err:  errors.New("RequestError: throttled"),
			want: "",
		},
		{
			name: "live and provisionable on success is no decision (listing blip)",
			out: &cloudformation.DescribeTypeOutput{
				DeprecatedStatus: types.DeprecatedStatusLive,
				ProvisioningType: types.ProvisioningTypeFullyMutable,
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := classifyAbsentProbe(tt.out, tt.err); got != tt.want {
				t.Errorf("classifyAbsentProbe() = %q, want %q", got, tt.want)
			}
		})
	}
}
