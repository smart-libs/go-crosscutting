package assertions

import (
	"context"
	"errors"
	"fmt"
	"githup.com/smart-libs/go-crosscutting/assertions/lib/pkg/check"
	"reflect"
	"strings"
	"time"
)

func ValueOfIsNotNil(v reflect.Value) error {
	if check.IsValueOfNil(v) {
		return errors.New("ValueOf.AnyIsNotNil() == true")
	}
	return nil
}

// AnyIsNotNil returns error if given any value is nil
func AnyIsNotNil(value any) error {
	if check.IsNil(value) {
		return fmt.Errorf("invalid nil value")
	}
	return nil
}

// IsContextValid returns nil if the context is valid, otherwise returns:
// - IllegalArgumentValue if it is nil
// -
// an errors.RouteError instance whose Name is errors.InternalError.
func IsContextValid(ctx context.Context, serviceName string) error {
	if ctx == nil {
		return errors.New("context is nil")
	}
	return ctx.Err()
}

// IsNotNilArray returns success if the given array is not nil
func IsNotNilArray[T any](a []T) error {
	if a == nil {
		return errors.New("array is nil")
	}
	return nil
}

// IsNotEmptyArray returns success if the given array is not empty
func IsNotEmptyArray[T any](a []T) error {
	err := IsNotNilArray(a)
	if err == nil && len(a) == 0 {
		return errors.New("array is empty")
	}
	return err
}

// IsNotEmptyOrBlank check whether the given string is empty or blank and if it is not, then it returns true, otherwise it returns false and set a bad request error.
func IsNotEmptyOrBlank(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return fmt.Errorf("[%s] cannot be empty or blank", s)
	}
	return nil
}

// IsPositiveInt check whether the given integer value is positive
func IsPositiveInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v T) error {
	if v <= 0 {
		return fmt.Errorf("%d <= 0", v)
	}
	return nil
}

// IsPositiveFloat check whether the given float value is positive
func IsPositiveFloat[T ~float32 | ~float64](v T, objName string) error {
	if v <= 0 {
		return fmt.Errorf("%f <= 0", v)
	}
	return nil
}

var defaultTime = time.Time{}

// IsValidDate check whether the given integer value is positive
func IsValidDate(v time.Time, objName string) error {
	if v == defaultTime {
		return fmt.Errorf("%s is the default value", v)
	}
	return nil
}
