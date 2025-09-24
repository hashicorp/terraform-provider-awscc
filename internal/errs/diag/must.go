// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package diag

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-provider-awscc/internal/errs"
)

// Must is a generic implementation of the Go Must idiom [1, 2]. It panics if
// the provided Diagnostics has errors and returns x otherwise.
//
// [1]: https://pkg.go.dev/text/template#Must
// [2]: https://pkg.go.dev/regexp#MustCompile
func Must[T any](x T, diags diag.Diagnostics) T {
	return errs.Must(x, DiagnosticsError(diags))
}
