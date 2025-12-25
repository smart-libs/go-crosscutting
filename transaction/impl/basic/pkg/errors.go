package transaction

import (
	"errors"
	"fmt"
	"github.com/joomcode/errorx"
)

type (
	IllegalStateError struct {
		error
	}

	PanicError struct {
		error
	}
)

func (i IllegalStateError) Unwrap() error { return i.error }
func (i IllegalStateError) Error() string {
	return fmt.Sprintf("IllegalStateError: %s", i.error.Error())
}

func (i PanicError) Unwrap() error { return i.error }
func (i PanicError) Error() string {
	return fmt.Sprintf("PanicError: %s", i.error.Error())
}

var (
	// ErrorNamespace is the namespace that identifies the go-transaction errors.
	// A go-db error is an error caused by go-db due to a malfunctioning.
	// Notice that
	ErrorNamespace = errorx.NewNamespace("go-transaction")

	IllegalState  = ErrorNamespace.NewType("illegal_state", errorx.Temporary())
	InternalError = ErrorNamespace.NewType("internal_error")

	WrapAsIllegalStateError = func(err error) error { return IllegalStateError{err} }
	WrapAsPanicError        = func(err error) error { return PanicError{err} }
)

func WrapWithNSifNeeded(err error) error {
	if err == nil {
		return nil
	}
	var cause error
	if e := errorx.Cast(err); e != nil {
		if e.Type() != nil && e.Type().Namespace().String() != "" {
			return e
		}
		cause = WrapWithNSifNeeded(e.Cause())
	} else {
		cause = WrapWithNSifNeeded(errors.Unwrap(err))
	}
	if cause == nil {
		return InternalError.WrapWithNoMessage(err)
	}
	return err
}
