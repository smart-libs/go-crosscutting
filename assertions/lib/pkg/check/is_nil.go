package check

import (
	"reflect"
)

func IsValueOfNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func IsNil(val any) bool {
	switch v := val.(type) {
	case nil:
		return true
	case reflect.Value:
		return IsValueOfNil(v)
	default:
		return IsValueOfNil(reflect.ValueOf(val))
	}
}
