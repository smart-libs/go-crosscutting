package assertions

import (
	"fmt"
	"strings"
)

func makePermittedValueList[T any](list ...T) string {
	firstValue := true
	result := strings.Builder{}
	result.WriteRune('[')
	for _, element := range list {
		if !firstValue {
			result.WriteRune(',')
		} else {
			firstValue = false
		}
		result.WriteString(fmt.Sprintf("%v", element))
	}
	result.WriteRune(']')
	return result.String()
}

// IsIn returns error if given value v is not in the given list of expected values
func IsIn[T comparable](v T, list ...T) error {
	for _, element := range list {
		if v == element {
			return nil
		}
	}
	return fmt.Errorf("value=[%v] not in the permitted value list=%s", v, makePermittedValueList(list...))
}

// IsInFunc returns a validation function to be used in NewValidatorWithFunc
func IsInFunc[T comparable](list ...T) func(T) error {
	return func(v T) error {
		return IsIn(v, list...)
	}
}
