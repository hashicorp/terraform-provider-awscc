// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
)

func TestIdentifierStatePath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Identifier
		expected path.Path
	}{
		{
			name: "top-level identifier",
			input: Identifier{
				Name: "qev_2_idc_application_arn",
			},
			expected: path.Root("qev_2_idc_application_arn"),
		},
		{
			name: "nested identifier",
			input: Identifier{
				Name:          "storage_lens_configuration_id",
				AttributePath: []string{"storage_lens_configuration", "id"},
			},
			expected: path.Root("storage_lens_configuration").AtName("id"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.input.StatePath(); !got.Equal(tt.expected) {
				t.Errorf("StatePath() = %s, want %s", got, tt.expected)
			}
		})
	}
}
