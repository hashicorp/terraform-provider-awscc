// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"strings"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/cli"
)

// A property pruned at the depth boundary of a de-recursed schema is kept as a
// marked placeholder. A resource emits it as an Optional attribute that is an error
// to set (so over-depth configuration is not silently discarded); a data source
// omits it.
func TestEmitterDerecursedPlaceholder(t *testing.T) {
	t.Parallel()

	stringType := cfschema.Type(cfschema.PropertyTypeString)
	objectType := cfschema.Type(cfschema.PropertyTypeObject)
	comment := derecursedPlaceholderComment
	description := "Not supported."
	resource := &cfschema.Resource{
		Properties: map[string]*cfschema.Property{
			"Match": {
				Type: &stringType,
			},
			"AndStatement": {
				Comment:     &comment,
				Description: &description,
				Type:        &objectType,
			},
		},
	}

	testCases := map[string]struct {
		isDataSource bool
		want         []string
		notWant      []string
	}{
		"resource": {
			want: []string{
				`"and_statement":schema.SingleNestedAttribute{`,
				`Attributes:map[string]schema.Attribute{},`,
				`Description:"Not supported.",`,
				`Optional:true,`,
				`fwvalidators.NestingDepthExceededObject(),`,
			},
			notWant: []string{
				`jsontypes.NormalizedType`,
			},
		},
		"data source": {
			isDataSource: true,
			want: []string{
				`"match":`,
			},
			notWant: []string{
				`and_statement`,
				`AndStatement`,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var rootSchema strings.Builder
			emitter := Emitter{
				CfResource:   resource,
				IsDataSource: testCase.isDataSource,
				Ui:           cli.NewMockUi(),
				Writer:       &rootSchema,
			}

			features, _, err := emitter.EmitRootPropertiesSchema("awscc_test_thing", make(map[string]string))
			if err != nil {
				t.Fatal(err)
			}

			got := rootSchema.String()
			for _, want := range testCase.want {
				if !strings.Contains(got, want) {
					t.Errorf("emitted schema does not contain %q:\n%s", want, got)
				}
			}
			for _, notWant := range testCase.notWant {
				if strings.Contains(got, notWant) {
					t.Errorf("emitted schema unexpectedly contains %q:\n%s", notWant, got)
				}
			}

			if !testCase.isDataSource && (!features.HasValidator || !features.UsesInternalValidatorsPackage) {
				t.Errorf("placeholder features = %+v, want HasValidator and UsesInternalValidatorsPackage", features)
			}
		})
	}
}
