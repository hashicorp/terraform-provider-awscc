package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnknowns(t *testing.T) {
	testCases := []struct {
		TestName      string
		Value         tftypes.Value
		TfToCfNameMap map[string]string
		ExpectedError bool
		ExpectedPaths []*tftypes.AttributePath
	}{
		{
			TestName:      "simple State",
			Value:         makeSimpleValueWithUnknowns(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedPaths: []*tftypes.AttributePath{
				tftypes.NewAttributePath().WithAttributeName("arn"),
				tftypes.NewAttributePath().WithAttributeName("identifier"),
			},
		},
		{
			TestName:      "complex State",
			Value:         makeComplexValueWithUnknowns(),
			TfToCfNameMap: complexTfToCfNameMap,
			ExpectedPaths: []*tftypes.AttributePath{
				tftypes.NewAttributePath().WithAttributeName("identifier"),
			},
		},
	}

	opts := cmp.Options{
		cmpopts.SortSlices(func(i, j *tftypes.AttributePath) bool {
			return i.String() < j.String()
		}),
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got, err := UnknownValuePaths(context.TODO(), testCase.Value)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil {
				if diff := cmp.Diff(got, testCase.ExpectedPaths, opts); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}

func TestUnknowsSetValue(t *testing.T) {
	testCases := []struct {
		TestName      string
		State         tfsdk.State
		ResourceModel string
		CfToTfNameMap map[string]string
		ExpectedError bool
		ExpectedState tfsdk.State
	}{
		{
			TestName: "simple State",
			State: tfsdk.State{
				Raw:    makeSimpleValueWithUnknowns(),
				Schema: testSimpleSchema,
			},
			ResourceModel: `{"Arn": "arn:aws:test:::test"}`,
			CfToTfNameMap: simpleCfToTfNameMap,
			ExpectedState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"identifier": tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arn:aws:test:::test"),
					"identifier": tftypes.NewValue(tftypes.String, nil),
					"name":       tftypes.NewValue(tftypes.String, "testing"),
					"number":     tftypes.NewValue(tftypes.Number, 42),
				}),
				Schema: testSimpleSchema,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			unknowns, err := UnknownValuePaths(context.TODO(), testCase.State.Raw)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			err = SetUnknownValuesFromResourceModel(context.TODO(), &testCase.State, unknowns, testCase.ResourceModel, testCase.CfToTfNameMap)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil {
				if diff := cmp.Diff(testCase.State, testCase.ExpectedState); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}
