// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"reflect"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

// JSON Schema minLength/maxLength count characters (code points), so the
// emitted validators must count runes rather than bytes.
func TestStringValidatorsCountCharacters(t *testing.T) {
	t.Parallel()

	stringType := cfschema.Type(cfschema.PropertyTypeString)
	min, max := 1, 200

	testCases := map[string]struct {
		minLength *int
		maxLength *int
		want      []string
	}{
		"at least": {minLength: &min, want: []string{"stringvalidator.UTF8LengthAtLeast(1)"}},
		"at most":  {maxLength: &max, want: []string{"stringvalidator.UTF8LengthAtMost(200)"}},
		"between":  {minLength: &min, maxLength: &max, want: []string{"stringvalidator.UTF8LengthBetween(1,200)"}},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			property := &cfschema.Property{Type: &stringType, MinLength: testCase.minLength, MaxLength: testCase.maxLength}

			_, got, err := stringValidators([]string{"Description"}, property)
			if err != nil {
				t.Fatalf("stringValidators() error = %v", err)
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("stringValidators() = %v, want %v", got, testCase.want)
			}
		})
	}
}
