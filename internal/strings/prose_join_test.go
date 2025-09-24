// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package strings

import "testing"

func TestProseJoin(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		values   []string
		expected string
	}{
		"empty array": {
			values:   []string{},
			expected: "",
		},
		"single element": {
			values:   []string{"single"},
			expected: "single",
		},
		"two elements": {
			values:   []string{"one", "two"},
			expected: "one and two",
		},
		"three elements": {
			values:   []string{"one", "two", "three"},
			expected: "one, two, and three",
		},
	}

	for name, testcase := range testcases {
		name, testcase := name, testcase

		t.Run(name, func(t *testing.T) {
			actual := ProseJoin(testcase.values)

			if actual != testcase.expected {
				t.Errorf("expected %q, got %q", testcase.expected, actual)
			}
		})
	}
}
