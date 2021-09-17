package tfresource

import (
	"errors"

	multierror "github.com/hashicorp/go-multierror"
	tfdiag "github.com/hashicorp/terraform-plugin-framework/diag"
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

func DiagsError(diags tfdiag.Diagnostics) error {
	var errs *multierror.Error

	for _, diag := range diags {
		if diag == nil {
			continue
		}

		if diag.Severity() == tfdiag.SeverityError {
			errs = multierror.Append(errs, errors.New(diag.Detail()))
		}
	}

	return errs.ErrorOrNil()
}
