package set

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestSetTypeTerraformType(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    SetType
		expected tftypes.Type
	}
	tests := map[string]testCase{
		"set-of-strings": {
			input: SetType{
				ElemType: types.StringType,
			},
			expected: tftypes.Set{
				ElementType: tftypes.String,
			},
		},
		"set-of-set-of-bools": {
			input: SetType{
				ElemType: SetType{
					ElemType: types.BoolType,
				},
			},
			expected: tftypes.Set{
				ElementType: tftypes.Set{
					ElementType: tftypes.Bool,
				},
			},
		},
		"set-of-set-of-list-of-number": {
			input: SetType{
				ElemType: SetType{
					ElemType: types.ListType{
						ElemType: types.NumberType,
					},
				},
			},
			expected: tftypes.Set{
				ElementType: tftypes.Set{
					ElementType: tftypes.List{
						ElementType: tftypes.Number,
					},
				},
			},
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			got := test.input.TerraformType(context.Background())
			if !got.Is(test.expected) {
				t.Errorf("Expected %s, got %s", test.expected, got)
			}
		})
	}
}

func TestSetTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver    SetType
		input       tftypes.Value
		expected    attr.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"set-of-strings": {
			receiver: SetType{
				ElemType: types.StringType,
			},
			input: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
				tftypes.NewValue(tftypes.String, "world"),
			}),
			expected: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
		},
		"unknown-set": {
			receiver: SetType{
				ElemType: types.StringType,
			},
			input: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, tftypes.UnknownValue),
			expected: Set{
				ElemType: types.StringType,
				Unknown:  true,
			},
		},
		"partially-unknown-set": {
			receiver: SetType{
				ElemType: types.StringType,
			},
			input: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
				tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
			expected: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
		},
		"null-set": {
			receiver: SetType{
				ElemType: types.StringType,
			},
			input: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, nil),
			expected: Set{
				ElemType: types.StringType,
				Null:     true,
			},
		},
		"partially-null-set": {
			receiver: SetType{
				ElemType: types.StringType,
			},
			input: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
				tftypes.NewValue(tftypes.String, nil),
			}),
			expected: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			got, gotErr := test.receiver.ValueFromTerraform(context.Background(), test.input)
			if gotErr != nil {
				if test.expectedErr != "" {
					if gotErr.Error() != test.expectedErr {
						t.Errorf("Expected error to be %q, got %q", test.expectedErr, gotErr.Error())
						return
					}
				}
				t.Errorf("Unexpected error: %s", gotErr.Error())
				return
			}
			if gotErr == nil && test.expectedErr != "" {
				t.Errorf("Expected error to be %q, got nil", test.expectedErr)
				return
			}
			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("Unexpected diff (-expected, +got): %s", diff)
			}
		})
	}
}

func TestSetTypeEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver SetType
		input    attr.Type
		expected bool
	}
	tests := map[string]testCase{
		"equal": {
			receiver: SetType{ElemType: types.StringType},
			input:    SetType{ElemType: types.StringType},
			expected: true,
		},
		"diff": {
			receiver: SetType{ElemType: types.StringType},
			input:    SetType{ElemType: types.NumberType},
			expected: false,
		},
		"wrongType": {
			receiver: SetType{ElemType: types.StringType},
			input:    types.NumberType,
			expected: false,
		},
		"nil": {
			receiver: SetType{ElemType: types.StringType},
			input:    nil,
			expected: false,
		},
		"nil-elem": {
			receiver: SetType{},
			input:    SetType{},
			// SetTypes with nil ElemTypes are invalid, and
			// aren't equal to anything
			expected: false,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := test.receiver.Equal(test.input)
			if test.expected != got {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestSetToTerraformValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Set
		expectation interface{}
	}
	tests := map[string]testCase{
		"value": {
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			expectation: []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
				tftypes.NewValue(tftypes.String, "world"),
			},
		},
		"unknown": {
			input:       Set{Unknown: true},
			expectation: tftypes.UnknownValue,
		},
		"null": {
			input:       Set{Null: true},
			expectation: nil,
		},
		"partial-unknown": {
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Unknown: true},
					types.String{Value: "hello, world"},
				},
			},
			expectation: []tftypes.Value{
				tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				tftypes.NewValue(tftypes.String, "hello, world"),
			},
		},
		"partial-null": {
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Null: true},
					types.String{Value: "hello, world"},
				},
			},
			expectation: []tftypes.Value{
				tftypes.NewValue(tftypes.String, nil),
				tftypes.NewValue(tftypes.String, "hello, world"),
			},
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := test.input.ToTerraformValue(context.Background())
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}
			if diff := cmp.Diff(got, test.expectation); diff != "" {
				t.Errorf("Unexpected result (+got, -expected): %s", diff)
			}
		})
	}
}

func TestSetEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver Set
		input    attr.Value
		expected bool
	}
	tests := map[string]testCase{
		"set-value-set-value": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			expected: true,
		},
		"set-value-diff": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "goodnight"},
					types.String{Value: "moon"},
				},
			},
			expected: false,
		},
		"set-value-diff-order": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "world"},
					types.String{Value: "hello"},
				},
			},
			expected: true,
		},
		"set-value-count-diff": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
					types.String{Value: "test"},
				},
			},
			expected: false,
		},
		"set-value-type-diff": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input: Set{
				ElemType: types.BoolType,
				Elems: []attr.Value{
					types.Bool{Value: false},
					types.Bool{Value: true},
				},
			},
			expected: false,
		},
		"set-value-unknown": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input:    Set{Unknown: true},
			expected: false,
		},
		"set-value-null": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input:    Set{Null: true},
			expected: false,
		},
		"set-value-wrongType": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input:    types.String{Value: "hello, world"},
			expected: false,
		},
		"set-value-nil": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			input:    nil,
			expected: false,
		},
		"partially-known-set-value-set-value": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
			expected: true,
		},
		"partially-known-set-value-diff": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			expected: false,
		},
		"partially-known-set-value-unknown": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
			input:    Set{Unknown: true},
			expected: false,
		},
		"partially-known-set-value-null": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
			input:    Set{Null: true},
			expected: false,
		},
		"partially-known-set-value-wrongType": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
			input:    types.String{Value: "hello, world"},
			expected: false,
		},
		"partially-known-set-value-nil": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Unknown: true},
				},
			},
			input:    nil,
			expected: false,
		},
		"partially-null-set-value-set-value": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
			expected: true,
		},
		"partially-null-set-value-diff": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
			input: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Value: "world"},
				},
			},
			expected: false,
		},
		"partially-null-set-value-unknown": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
			input: Set{
				Unknown: true,
			},
			expected: false,
		},
		"partially-null-set-value-null": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
			input: Set{
				Null: true,
			},
			expected: false,
		},
		"partially-null-set-value-wrongType": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
			input:    types.String{Value: "hello, world"},
			expected: false,
		},
		"partially-null-set-value-nil": {
			receiver: Set{
				ElemType: types.StringType,
				Elems: []attr.Value{
					types.String{Value: "hello"},
					types.String{Null: true},
				},
			},
			input:    nil,
			expected: false,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := test.receiver.Equal(test.input)
			if got != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}
