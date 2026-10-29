// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

// A two-level slice of a de-recursed (WAFv2-like) schema: the level-2 statement
// keeps "match" but its "and_statement" is a placeholder, as the code generator
// emits for a property pruned at the depth boundary.
func makeDerecursedTestSchema() schema.Schema {
	placeholder := schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{},
		Optional:   true,
		Validators: []validator.Object{fwvalidators.NestingDepthExceededObject()},
	}

	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"statement": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"match": schema.StringAttribute{Optional: true},
					"and_statement": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"statements": schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match":         schema.StringAttribute{Optional: true},
										"and_statement": placeholder,
									},
								},
								Optional: true,
							},
						},
						Optional: true,
					},
				},
				Optional: true,
			},
		},
	}
}

var derecursedCfToTfNameMap = map[string]string{
	"AndStatement": "and_statement",
	"Match":        "match",
	"Statement":    "statement",
	"Statements":   "statements",
}

// A remote resource model nested one level deeper than the schema.
const overDepthResourceModel = `{
  "Statement": {
    "AndStatement": {
      "Statements": [
        {"Match": "a"},
        {"AndStatement": {"Statements": [{"Match": "b"}, {"Match": "c"}]}}
      ]
    }
  }
}`

func newDerecursedTestResource() *genericResource {
	return &genericResource{
		tfSchema:                  makeDerecursedTestSchema(),
		cfToTfNameMap:             derecursedCfToTfNameMap,
		unrepresentableProperties: []string{"AndStatement"},
	}
}

func TestTranslateToTerraformDerecursedPlaceholder(t *testing.T) {
	ctx := context.Background()
	r := newDerecursedTestResource()

	var skipped []skippedProperty
	translator := toTerraform{cfToTfNameMap: r.cfToTfNameMap, skipped: &skipped, unrepresentable: r.unrepresentablePropertySet()}
	got, err := translator.FromString(ctx, r.tfSchema, overDepthResourceModel, nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	placeholderType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{}}
	level2Type := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"and_statement": placeholderType,
		"match":         tftypes.String,
	}}
	andType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"statements": tftypes.List{ElementType: level2Type},
	}}
	statementType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"and_statement": andType,
		"match":         tftypes.String,
	}}
	rootType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"statement": statementType,
	}}

	// The over-depth AndStatement is skipped: the placeholder stays null instead of
	// being read as an (empty, non-null) object that would then plan for removal.
	want := tftypes.NewValue(rootType, map[string]tftypes.Value{
		"statement": tftypes.NewValue(statementType, map[string]tftypes.Value{
			"and_statement": tftypes.NewValue(andType, map[string]tftypes.Value{
				"statements": tftypes.NewValue(tftypes.List{ElementType: level2Type}, []tftypes.Value{
					tftypes.NewValue(level2Type, map[string]tftypes.Value{
						"and_statement": tftypes.NewValue(placeholderType, nil),
						"match":         tftypes.NewValue(tftypes.String, "a"),
					}),
					tftypes.NewValue(level2Type, map[string]tftypes.Value{
						"and_statement": tftypes.NewValue(placeholderType, nil),
						"match":         tftypes.NewValue(tftypes.String, nil),
					}),
				}),
			}),
			"match": tftypes.NewValue(tftypes.String, nil),
		}),
	})

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	if len(skipped) != 1 || skipped[0].PropertyName != "AndStatement" {
		t.Fatalf("expected exactly the over-depth AndStatement to be skipped, got %v", skipped)
	}

	if diag := r.unrepresentablePropertiesWarning(skipped); diag == nil {
		t.Error("expected a warning diagnostic for the skipped AndStatement")
	}
}

func TestFindUnrepresentablePropertyPathsDerecursedPlaceholder(t *testing.T) {
	ctx := context.Background()
	r := newDerecursedTestResource()

	got := r.findUnrepresentablePropertyPaths(ctx, overDepthResourceModel)
	want := []string{"/Statement/AndStatement/Statements/1/AndStatement"}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}

	testCases := []struct {
		name          string
		patch         string
		expectBlocked bool
	}{
		{
			name:          "replace ancestor",
			patch:         `[{"op":"replace","path":"/Statement","value":{}}]`,
			expectBlocked: true,
		},
		{
			name:          "replace sibling",
			patch:         `[{"op":"replace","path":"/Statement/AndStatement/Statements/0/Match","value":"z"}]`,
			expectBlocked: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			blocked, _ := r.patchTouchesUnrepresentableProperties(ctx, testCase.patch, overDepthResourceModel)
			if blocked != testCase.expectBlocked {
				t.Errorf("blocked = %t, want %t", blocked, testCase.expectBlocked)
			}
		})
	}
}
