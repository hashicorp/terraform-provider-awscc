package generic

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestPropertyPathToAttributePath(t *testing.T) {
	testCases := []struct {
		TestName      string
		PropertyPath  string
		ExpectedError bool
		ExpectedValue *tftypes.AttributePath
	}{
		{
			TestName:      "empty property path",
			ExpectedError: true,
		},
		{
			TestName:      "no seperators",
			PropertyPath:  "test",
			ExpectedError: true,
		},
		{
			TestName:      "just root",
			PropertyPath:  "/properties",
			ExpectedError: true,
		},
		{
			TestName:      "invalid root",
			PropertyPath:  "/definitions/BasicAuthParameters/Password",
			ExpectedError: true,
		},
		{
			TestName:      "one path segment",
			PropertyPath:  "/properties/Tags",
			ExpectedValue: tftypes.NewAttributePath().WithAttributeName("tags"),
		},
		{
			TestName:      "two path segments",
			PropertyPath:  "/properties/BasicAuthParameters/Password",
			ExpectedValue: tftypes.NewAttributePath().WithAttributeName("basic_auth_parameters").WithAttributeName("password"),
		},
		{
			TestName:      "empty segment",
			PropertyPath:  "/properties//Password",
			ExpectedError: true,
		},
		{
			TestName:      "segment with *",
			PropertyPath:  "/properties/Actions/*/AuthenticateOidcConfig/ClientSecret",
			ExpectedError: true,
		},
	}

	rt := resourceType{
		cfToTfNameMap: map[string]string{
			"BasicAuthParameters": "basic_auth_parameters",
			"Password":            "password",
			"Tags":                "tags",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			attributePath, err := rt.propertyPathToAttributePath(testCase.PropertyPath)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from propertyPathToAttributePath")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from propertyPathToAttributePath: %s", err)
			}

			if !attributePath.Equal(testCase.ExpectedValue) {
				t.Errorf("got: %s, expected: %s", attributePath, testCase.ExpectedValue)
			}
		})
	}
}
