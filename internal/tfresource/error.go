// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfresource

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type NotFoundError struct {
	LastError error
	Message   string
}

func (e *NotFoundError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "couldn't find resource"
}

func (e *NotFoundError) Unwrap() error {
	return e.LastError
}

// NotFound returns true if the error represents a "resource not found" condition.
// Specifically, NotFound returns true if the error or a wrapped error is of type
// resource.NotFoundError.
func NotFound(err error) bool {
	var e *NotFoundError
	return errors.As(err, &e)
}

func DiagnosticsError(diags diag.Diagnostics) error {
	var errs []error

	for _, d := range diags.Errors() {
		errs = append(errs, errors.New(DiagnosticString(d)))
	}

	return errors.Join(errs...)
}

func DiagnosticString(d diag.Diagnostic) string {
	var buf strings.Builder

	fmt.Fprint(&buf, d.Summary())
	if d.Detail() != "" {
		fmt.Fprintf(&buf, "\n\n%s", d.Detail())
	}
	if withPath, ok := d.(diag.DiagnosticWithPath); ok {
		fmt.Fprintf(&buf, "\n%s", withPath.Path().String())
	}

	return buf.String()
}
