// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// decodeOneRow decodes a single-block overlay chunk into a resourceRow.
func decodeOneRow(t *testing.T, src string) resourceRow {
	t.Helper()
	var af availableFile
	if err := hclsimple.Decode("block.hcl", []byte(src), nil, &af); err != nil {
		t.Fatalf("decoding mutated block: %v\n%s", err, src)
	}
	if len(af.Resources) != 1 {
		t.Fatalf("expected 1 block, got %d\n%s", len(af.Resources), src)
	}
	return af.Resources[0]
}

func TestSetBlockAttributes(t *testing.T) {
	t.Parallel()

	t.Run("adds non_provisionable, preserves comment and existing attr", func(t *testing.T) {
		t.Parallel()
		in := `resource_schema "aws_appstream_stack_fleet_association" {
  # Suppression Reason: keep this note
  cloudformation_type_name               = "AWS::AppStream::StackFleetAssociation"
  suppress_plural_data_source_generation = true
}
`
		out, err := setBlockAttributes(in, map[string]string{attrNonProvisionable: "true"})
		if err != nil {
			t.Fatalf("setBlockAttributes: %v", err)
		}
		if !strings.Contains(out, "# Suppression Reason: keep this note") {
			t.Errorf("comment not preserved:\n%s", out)
		}
		row := decodeOneRow(t, out)
		if !row.NonProvisionable {
			t.Errorf("non_provisionable not set:\n%s", out)
		}
		if !row.SuppressPluralDataSourceGeneration {
			t.Errorf("existing suppress_plural lost:\n%s", out)
		}
		if row.CloudFormationTypeName != "AWS::AppStream::StackFleetAssociation" {
			t.Errorf("cfn type name altered: %q", row.CloudFormationTypeName)
		}
	})

	t.Run("adds frozen_since as a quoted string", func(t *testing.T) {
		t.Parallel()
		in := `resource_schema "aws_logs_log_group" {
  cloudformation_type_name = "AWS::Logs::LogGroup"
}
`
		out, err := setBlockAttributes(in, map[string]string{attrFrozenSince: "2026-08-27"})
		if err != nil {
			t.Fatalf("setBlockAttributes: %v", err)
		}
		if !strings.Contains(out, "frozen_since") || !strings.Contains(out, `"2026-08-27"`) {
			t.Errorf("frozen_since not written as a quoted string:\n%s", out)
		}
		if decodeOneRow(t, out).FrozenSince != "2026-08-27" {
			t.Errorf("frozen_since not decodable:\n%s", out)
		}
	})

	t.Run("updates an existing attribute in place", func(t *testing.T) {
		t.Parallel()
		in := `resource_schema "aws_svc_thing" {
  cloudformation_type_name = "AWS::Svc::Thing"
  frozen_since             = "2020-01-01"
}
`
		out, err := setBlockAttributes(in, map[string]string{attrFrozenSince: "2026-08-27"})
		if err != nil {
			t.Fatalf("setBlockAttributes: %v", err)
		}
		row := decodeOneRow(t, out)
		if row.FrozenSince != "2026-08-27" {
			t.Errorf("frozen_since not updated, got %q\n%s", row.FrozenSince, out)
		}
		if strings.Contains(out, "2020-01-01") {
			t.Errorf("stale value remained:\n%s", out)
		}
	})

	t.Run("no attrs is a no-op", func(t *testing.T) {
		t.Parallel()
		in := "resource_schema \"aws_svc_thing\" {\n  cloudformation_type_name = \"AWS::Svc::Thing\"\n}\n"
		out, err := setBlockAttributes(in, nil)
		if err != nil || out != in {
			t.Errorf("expected unchanged text, got err=%v\n%s", err, out)
		}
	})
}
