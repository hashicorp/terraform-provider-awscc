// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Test_unorderedArrayPaths(t *testing.T) {
	tests := []struct {
		name          string
		attrs         map[string]schema.Attribute
		tfToCfNameMap map[string]string
		want          map[string]bool
	}{
		{
			name: "top-level set nested attribute",
			attrs: map[string]schema.Attribute{
				"config": schema.SetNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"day": schema.StringAttribute{},
						},
					},
				},
				"name": schema.StringAttribute{},
			},
			tfToCfNameMap: map[string]string{"config": "Config", "name": "Name", "day": "Day"},
			want:          map[string]bool{"/Config": true},
		},
		{
			name: "top-level primitive set attribute",
			attrs: map[string]schema.Attribute{
				"subnet_ids": schema.SetAttribute{ElementType: types.StringType},
			},
			tfToCfNameMap: map[string]string{"subnet_ids": "SubnetIds"},
			want:          map[string]bool{"/SubnetIds": true},
		},
		{
			name: "list nested attribute is ordered, not flagged, but still recursed into",
			attrs: map[string]schema.Attribute{
				"queue_configs": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"tags": schema.SetNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{},
									},
								},
							},
						},
					},
				},
			},
			tfToCfNameMap: map[string]string{"queue_configs": "QueueConfigs", "tags": "Tags"},
			want:          map[string]bool{"/QueueConfigs/Tags": true},
		},
		{
			name: "set nested attribute inside single nested attribute",
			attrs: map[string]schema.Attribute{
				"vpc_config": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"security_group_ids": schema.SetAttribute{ElementType: types.StringType},
					},
				},
			},
			tfToCfNameMap: map[string]string{"vpc_config": "VpcConfig", "security_group_ids": "SecurityGroupIds"},
			want:          map[string]bool{"/VpcConfig/SecurityGroupIds": true},
		},
		{
			name: "no unordered attributes",
			attrs: map[string]schema.Attribute{
				"name":  schema.StringAttribute{},
				"ports": schema.ListAttribute{ElementType: types.NumberType},
			},
			tfToCfNameMap: map[string]string{"name": "Name", "ports": "Ports"},
			want:          map[string]bool{},
		},
		{
			// insertionOrder:false + non-unique items renders as a List (not a Set)
			// carrying the Multiset plan modifier; it is equally order-unstable.
			name: "top-level multiset list nested attribute",
			attrs: map[string]schema.Attribute{
				"config": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"day": schema.StringAttribute{},
						},
					},
					PlanModifiers: []planmodifier.List{Multiset()},
				},
			},
			tfToCfNameMap: map[string]string{"config": "Config", "day": "Day"},
			want:          map[string]bool{"/Config": true},
		},
		{
			name: "top-level multiset primitive list attribute",
			attrs: map[string]schema.Attribute{
				"values": schema.ListAttribute{
					ElementType:   types.StringType,
					PlanModifiers: []planmodifier.List{Multiset()},
				},
			},
			tfToCfNameMap: map[string]string{"values": "Values"},
			want:          map[string]bool{"/Values": true},
		},
		{
			// A plain ordered list (no Multiset modifier) must not be flagged, even
			// when it carries other plan modifiers.
			name: "ordered list with non-multiset plan modifier is not flagged",
			attrs: map[string]schema.Attribute{
				"steps": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{},
						},
					},
					PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				},
			},
			tfToCfNameMap: map[string]string{"steps": "Steps", "name": "Name"},
			want:          map[string]bool{},
		},
		{
			// Multiset nested inside an ordinary ordered list: only the inner,
			// order-unstable array is flagged, at its index-free skeleton path.
			name: "multiset nested inside an ordered list",
			attrs: map[string]schema.Attribute{
				"containers": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"command": schema.ListAttribute{
								ElementType:   types.StringType,
								PlanModifiers: []planmodifier.List{Multiset()},
							},
						},
					},
				},
			},
			tfToCfNameMap: map[string]string{"containers": "Containers", "command": "Command"},
			want:          map[string]bool{"/Containers/Command": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unorderedArrayPaths(tt.attrs, tt.tfToCfNameMap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unorderedArrayPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
