package generator

import (
	"testing"

	"github.com/iancoleman/strcase"
)

func TestToSnake(t *testing.T) {
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
			Input:    "abc",
			Expected: "abc",
		},
		{
			TestName: "first char capitalized",
			Input:    "Abc",
			Expected: "abc",
		},
		{
			TestName: "last char capitalized",
			Input:    "abcZ",
			Expected: "abc_z",
		},
		{
			TestName: "all caps",
			Input:    "ABCZ",
			Expected: "abcz",
		},
		{
			TestName: "mixture",
			Input:    "abcDeFGh",
			Expected: "abc_de_f_gh",
		},
		{
			TestName: "leading digits",
			Input:    "012abCd",
			Expected: "012_ab_cd",
		},
		{
			TestName: "trailing digits",
			Input:    "abCd012",
			Expected: "ab_cd",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got := strcase.ToSnake(testCase.Input)

			if testCase.Expected != got {
				t.Errorf("got %q, expected %q", got, testCase.Expected)
			}
		})
	}
}
