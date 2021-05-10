package generators

import (
	"path/filepath"
	"testing"
)

func TestNewConfigPath(t *testing.T) {
	testCases := []struct {
		TestDescription string
		ConfigFilePath  string
		ExpectError     bool
	}{
		{
			TestDescription: "valid configuration",
			ConfigFilePath:  "config.good.hcl",
		},
		{
			TestDescription: "invalid configuration",
			ConfigFilePath:  "config.bad.hcl",
			ExpectError:     true,
		},
		{
			TestDescription: "missing configuration file",
			ConfigFilePath:  "missing file",
			ExpectError:     true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			_, err := NewConfigPath(filepath.Join("testdata", testCase.ConfigFilePath))

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}
