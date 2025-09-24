// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package slices

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAppendUnique(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    []string
		append   []string
		expected []string
	}
	tests := map[string]testCase{
		"all nil": {},
		"all empty": {
			input:    []string{},
			append:   []string{},
			expected: []string{},
		},
		"append to nil": {
			append:   []string{"alpha", "beta", "alpha"},
			expected: []string{"alpha", "beta"},
		},
		"append nothing": {
			input:    []string{"alpha", "beta"},
			append:   []string{},
			expected: []string{"alpha", "beta"},
		},
		"append no unique": {
			input:    []string{"alpha", "beta", "gamma"},
			append:   []string{"beta", "gamma", "alpha"},
			expected: []string{"alpha", "beta", "gamma"},
		},
		"append one unique": {
			input:    []string{"alpha", "beta", "gamma"},
			append:   []string{"beta", "delta"},
			expected: []string{"alpha", "beta", "gamma", "delta"},
		},
		"append three unique": {
			input:    []string{"alpha", "beta", "gamma"},
			append:   []string{"delta", "gamma", "epsilon", "alpha", "epsilon", "zeta"},
			expected: []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta"},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := AppendUnique(test.input, test.append...)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
