package assertions

import (
	"fmt"
	"strings"
)

type (
	MultiError    []error
	InternalError struct {
		cause error
	}

	IllegalArgument struct {
		causes  MultiError
		message string
	}
)

var (
	WrapAsInternalError = func(err error) error {
		return InternalError{cause: err}
	}

	WrapAsIllegalArgumentValue = func(name string, value any) error {
		return WrapAsIllegalArgumentValueWithCause(name, value, nil)
	}

	WrapAsIllegalArgumentValueWithCause = func(name string, value any, causes ...error) error {
		return IllegalArgument{
			causes:  causes,
			message: fmt.Sprintf("invalid value=[%v] for argument=[%s]", value, name),
		}
	}
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func (m MultiError) Error() string {
	var buffer strings.Builder
	numOfCauses := len(m)
	if numOfCauses > 0 {
		if numOfCauses == 1 {
			buffer.WriteString(m[0].Error())
		} else {
			for i, e := range m {
				if i > 0 {
					buffer.WriteRune(',')
				}
				_, _ = fmt.Fprintf(&buffer, "[%d]={%s}", i, e.Error())
			}
		}
	}
	return buffer.String()
}

func (i IllegalArgument) Error() string {
	return fmt.Sprintf("illegal_argument: %s, caused by %s", i.message, i.causes.Error())
}

func is[T error](err error) bool {
	if _, ok := err.(T); ok {
		return true
	}
	return false

}

func (i IllegalArgument) Is(err error) bool {
	return is[IllegalArgument](err) || is[*IllegalArgument](err)
}

func (i IllegalArgument) Unwrap() error {
	return i.causes
}

func (i InternalError) Error() string {
	return fmt.Sprintf("internal_error: caused by %s", i.cause.Error())
}

func (i InternalError) Unwrap() error {
	return i.cause
}

func (i InternalError) Is(err error) bool {
	return is[InternalError](err) || is[*InternalError](err)
}

func WrapError(objectNameToBeValidated string, val any, err error) error {
	switch castedErr := err.(type) {
	case nil:
		return nil
	case InternalError, *InternalError:
		return castedErr
	case IllegalArgument:
		return WrapAsIllegalArgumentValueWithCause(objectNameToBeValidated, val, castedErr.causes...)
	case *IllegalArgument:
		return WrapAsIllegalArgumentValueWithCause(objectNameToBeValidated, val, castedErr.causes...)
	default:
		return WrapAsIllegalArgumentValueWithCause(objectNameToBeValidated, val, err)
	}
}

var (
	_ error = InternalError{}
	_ error = IllegalArgument{}
)
