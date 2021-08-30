package validate

import (
	"testing"

	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		required    []string
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required": {
			names:    []string{"alpha", "beta", "gamma"},
			required: []string{"alpha", "gamma"},
		},
		"all required": {
			names:    []string{"alpha", "beta", "gamma"},
			required: []string{"alpha", "beta", "gamma"},
		},
		"missing one": {
			names:       []string{"alpha", "beta", "gamma"},
			required:    []string{"beta", "delta"},
			expectError: true,
		},
		"missing all": {
			names:       []string{"alpha", "beta", "gamma"},
			required:    []string{"sigma", "tau"},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := Required(test.required...)(test.names)

			if !tfresource.DiagsHasError(diags) && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if tfresource.DiagsHasError(diags) && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestAllOfRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), Required("gamma")},
		},
		"all required": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), AllOfRequired(Required("beta"), Required("gamma"))},
		},
		"missing one": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("beta"), Required("delta")},
			expectError: true,
		},
		"missing all": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("sigma"), AllOfRequired(Required("tau"))},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := AllOfRequired(test.fs...)(test.names)

			if !tfresource.DiagsHasError(diags) && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if tfresource.DiagsHasError(diags) && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestAnyOfRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), Required("delta")},
		},
		"nested allOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested allOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("epsilon"))},
			expectError: true,
		},
		"nested anyOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested anyOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("sigma"), Required("tau"))},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := AnyOfRequired(test.fs...)(test.names)

			if !tfresource.DiagsHasError(diags) && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if tfresource.DiagsHasError(diags) && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}

func TestOneOfRequired(t *testing.T) {
	t.Parallel()

	type testCase struct {
		names       []string
		fs          []RequiredAttributesFunc
		expectError bool
	}
	tests := map[string]testCase{
		"both empty": {},
		"none required": {
			names: []string{"alpha", "beta", "gamma"},
		},
		"some required OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("alpha"), Required("delta")},
		},
		"some required error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("alpha"), Required("beta")},
			expectError: true,
		},
		"nested allOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested allOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AllOfRequired(Required("beta"), Required("epsilon"))},
			expectError: true,
		},
		"nested anyOf OK": {
			names: []string{"alpha", "beta", "gamma"},
			fs:    []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("beta"), Required("gamma"))},
		},
		"nested anyOf error": {
			names:       []string{"alpha", "beta", "gamma"},
			fs:          []RequiredAttributesFunc{Required("delta"), AnyOfRequired(Required("sigma"), Required("tau"))},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			diags := OneOfRequired(test.fs...)(test.names)

			if !tfresource.DiagsHasError(diags) && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if tfresource.DiagsHasError(diags) && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(diags))
			}
		})
	}
}
