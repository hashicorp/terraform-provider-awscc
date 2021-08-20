package tfresource

import (
	"errors"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
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

func DiagsHasError(diags []*tfprotov6.Diagnostic) bool {
	for _, diag := range diags {
		if diag == nil {
			continue
		}

		if diag.Severity == tfprotov6.DiagnosticSeverityError {
			return true
		}
	}

	return false
}

func DiagsError(diags []*tfprotov6.Diagnostic) error {
	var errs *multierror.Error

	for _, diag := range diags {
		if diag == nil {
			continue
		}

		if diag.Severity == tfprotov6.DiagnosticSeverityError {
			errs = multierror.Append(errs, errors.New(diag.Detail))
		}
	}

	return errs.ErrorOrNil()
}
