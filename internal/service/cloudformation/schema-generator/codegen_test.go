package generator

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		TestName string
		Input    string
		Expected string
	}{
		{
			TestName: "empty string",
		},
		{
			TestName: "no caps",
			Input:    "abc0123",
			Expected: "abc0123",
		},
		{
			TestName: "first char capitalized",
			Input:    "Abc0123",
			Expected: "abc0123",
		},
		{
			TestName: "last char capitalized",
			Input:    "abc0123Z",
			Expected: "abc0123z",
		},
		{
			TestName: "all caps",
			Input:    "ABCZ",
			Expected: "a_b_c_z",
		},
		{
			TestName: "mixture",
			Input:    "abcDeFGh",
			Expected: "abc_de_f_gh",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got := ToSnakeCase(testCase.Input)

			if testCase.Expected != got {
				t.Errorf("got %q, expected %q", got, testCase.Expected)
			}
		})
	}
}
