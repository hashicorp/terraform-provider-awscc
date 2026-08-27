// SPDX-License-Identifier: MPL-2.0

package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
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
