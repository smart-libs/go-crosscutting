package serror

import (
	"github.com/joomcode/errorx"
)

type (
	CallbackCondition struct {
		Condition func(err error) bool
		Callback  func(err error)
	}
)

func IsIllegalArgumentError(err error) bool {
	return errorx.IsOfType(err, errorx.IllegalArgument)
}

func IsIllegalConfigError(err error) bool {
	return errorx.IsOfType(err, IllegalConfig)
}

func IsTimeoutError(err error) bool {
	return errorx.IsTimeout(err)
}

func IsNotFoundError(err error) bool {
	return errorx.IsNotFound(err)
}

func IsDuplicateError(err error) bool {
	return errorx.IsDuplicate(err)
}

// IdentifyRootCause invokes the first CallbackCondition that matches. It returns true if a callback function was invoked,
// otherwise it returns false.
func IdentifyRootCause(err error, fallback func(error), onErrorCallback ...CallbackCondition) bool {
	if err == nil {
		return false
	}
	for _, callback := range onErrorCallback {
		if callback.Condition != nil && callback.Condition(err) {
			if callback.Callback != nil {
				callback.Callback(err)
				return true
			}
		}
	}
	if fallback != nil {
		fallback(err)
		return true
	}
	return false
}
