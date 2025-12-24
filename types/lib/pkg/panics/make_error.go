package panics

import (
	"errors"
	"fmt"
)

func MakeErrorAddingContextInfo(recoveryArg any, makeMessage func() string) error {
	return MakeError(recoveryArg, func(panicErr error) error {
		return errors.Join(panicErr, errors.New(makeMessage()))
	})
}

func MakeError(recoveryArg any, makeContextError func(panicErr error) error) error {
	if recoveryArg == nil {
		return nil
	}
	err, ok := recoveryArg.(error)
	if !ok {
		err = fmt.Errorf("panic(%v)", recoveryArg)
	}
	return makeContextError(err)
}
