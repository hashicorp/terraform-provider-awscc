package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func TestTranslateToCloudFormation(t *testing.T) {
	testCases := []struct {
		TestName      string
		Plan          tfsdk.Plan
		TfToCfNameMap map[string]string
		ExpectedError bool
		ExpectedState map[string]interface{}
	}{
		{
			TestName:      "simple Plan",
			Plan:          makeSimpleTestPlan(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedState: map[string]interface{}{
				"Name": "testing",
			},
		},
		{
			TestName:      "simple Plan with Optional",
			Plan:          makeSimpleTestPlanWithOptionalPopulated(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedState: map[string]interface{}{
				"Name":   "testing",
				"Number": float64(42),
			},
		},
		{
			TestName:      "complex Plan",
			Plan:          makeComplexTestPlan(),
			TfToCfNameMap: complexTfToCfNameMap,
			ExpectedState: map[string]interface{}{
				"Name":        "hello, world",
				"MachineType": "e2-medium",
				"Ports":       []interface{}{float64(80), float64(443)},
				"Tags":        []interface{}{"red", "blue", "green"},
				"Disks": []interface{}{
					map[string]interface{}{
						"Id":                 "disk0",
						"DeleteWithInstance": true,
					},
					map[string]interface{}{
						"Id":                 "disk1",
						"DeleteWithInstance": false,
					},
				},
				"BootDisk": map[string]interface{}{
					"Id":                 "bootdisk",
					"DeleteWithInstance": true,
				},
				"ScratchDisk": map[string]interface{}{
					"Interface": "SCSI",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			translator := toCloudFormation{tfToCfNameMap: testCase.TfToCfNameMap}
			got, err := translator.AsRaw(context.TODO(), testCase.Plan.Raw)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil {
				if diff := cmp.Diff(got, testCase.ExpectedState); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}
