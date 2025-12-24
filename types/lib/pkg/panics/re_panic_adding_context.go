package panics

import (
	"errors"
)

// RePanicAddingContextInfo recovers from a panic and re-panics with context-specific information added to the error.
func RePanicAddingContextInfo(recoveryArg any, makeMessage func() string) {
	RePanic(recoveryArg, func(panicErr error) error {
		return errors.Join(panicErr, errors.New(makeMessage()))
	})
}

// RePanic recovers from a panic and re-panics with context-specific information added to the error.
func RePanic(recoveryArg any, makeContextError func(panicErr error) error) {
	err := MakeError(recoveryArg, makeContextError)
	if err == nil {
		return
	}
	panic(err)
}
