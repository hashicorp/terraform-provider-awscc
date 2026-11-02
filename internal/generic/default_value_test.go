// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestStringDefaultValue(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		plan  types.String
		state types.String
		want  types.String
	}{
		"omitted on create remains omitted": {
			plan:  types.StringNull(),
			state: types.StringNull(),
			want:  types.StringNull(),
		},
		"matching service default is preserved": {
			plan:  types.StringNull(),
			state: types.StringValue("MCP"),
			want:  types.StringValue("MCP"),
		},
		"non-default state is not treated as the default": {
			plan:  types.StringNull(),
			state: types.StringValue("NONE"),
			want:  types.StringNull(),
		},
		"configured value is unchanged": {
			plan:  types.StringValue("MCP"),
			state: types.StringNull(),
			want:  types.StringValue("MCP"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			response := planmodifier.StringResponse{PlanValue: test.plan}
			StringDefaultValue("MCP").PlanModifyString(context.Background(), planmodifier.StringRequest{
				PlanValue:  test.plan,
				StateValue: test.state,
			}, &response)

			if diff := cmp.Diff(test.want, response.PlanValue); diff != "" {
				t.Errorf("unexpected plan value (-want, +got):\n%s", diff)
			}
		})
	}
}
