package pointers

import (
	"fmt"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg"
	"reflect"
)

var (
	nillableKinds = map[reflect.Kind]bool{
		reflect.Chan:          true,
		reflect.Func:          true,
		reflect.Map:           true,
		reflect.Ptr:           true,
		reflect.UnsafePointer: true,
		reflect.Interface:     true,
		reflect.Slice:         true,
	}
)

func ToString(s string) *string { return &s }

func ToDecimal(d types.Decimal) *types.Decimal { return &d }

func To[T any](d T) *T { return &d }

func IsNil[T any](value T) bool {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Invalid {
		return true // nil
	}
	if nillableKinds[v.Kind()] {
		return v.IsNil()
	}
	return false
}

func AssertNotNil[T any](value T) {
	if IsNil(value) {
		panic(fmt.Errorf("AssertNotNil: %T value cannot be nil", value))
	}
}
