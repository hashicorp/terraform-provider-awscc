// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package naming_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
)

// TestPluralize guards this copy against drifting from internal/naming, whose
// TestPluralize is the canonical source; the cases here mirror the behaviors the
// generator relies on, including the already-plural rule that keeps the plural
// data source name distinct from the singular.
func TestPluralize(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name, in, want string
	}{
		{"empty", "", ""},
		{"inflection pluralizes", "aws_wafv2_web_acl", "aws_wafv2_web_acls"},
		{"ends in s", "aws_cloudwatch_event_bus", "aws_cloudwatch_event_buses"},
		{"ends in a digit", "aws_datasync_location_s3", "aws_datasync_location_s3s"},
		{"custom list: nfs", "aws_example_nfs", "aws_example_nfs_plural"},
		{"custom list: windows", "aws_datasync_windows", "aws_datasync_windows_plural"},
		// Already-plural names inflection leaves unchanged that are not in the
		// custom list: without the general fallback these collide with the
		// singular identifier (the SSMGuiConnect::Preferences failure mode).
		{"already plural: preferences", "awscc_ssmguiconnect_preferences", "awscc_ssmguiconnect_preferences_plural"},
		{"already plural: series", "aws_example_series", "aws_example_series_plural"},
		{"already plural: xfs", "aws_example_xfs", "aws_example_xfs_plural"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := naming.Pluralize(tc.in); got != tc.want {
				t.Errorf("Pluralize(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
