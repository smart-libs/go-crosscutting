package assertions

import (
	"fmt"
	"reflect"
)

// IsValid is a function that actively uses the ValidationFunc registry to identify the function that
// must be used to validate the given value.
func IsValid[T any](val T, objectName string) error {
	valOf := reflect.ValueOf(val)
	if err := ValueOfIsNotNil(valOf); err != nil {
		return WrapAsIllegalArgumentValueWithCause(objectName, "nil", err)
	}

	validationFunc, found := findValidationFunc[T]()
	if !found {
		if DefaultValidationFunc != nil {
			return WrapError(objectName, val, DefaultValidationFunc(val))
		}
		return WrapAsInternalError(fmt.Errorf("no ValidationFunc found for type=[%T] or kind=[%d]", val, valOf.Kind()))
	}

	return WrapError(objectName, val, validationFunc(val))
}
