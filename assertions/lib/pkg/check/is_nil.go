package check

import (
	"reflect"
)

func IsValueOfNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func IsNil(val any) bool {
	switch val.(type) {
	case nil:
		return true
	default:
		return IsValueOfNil(reflect.ValueOf(val))
	}
}
