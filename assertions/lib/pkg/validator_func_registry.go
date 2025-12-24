package assertions

import (
	"fmt"
	"github.com/smart-libs/go-crosscutting/assertions/lib/pkg/check"
	"github.com/smart-libs/go-crosscutting/serror/lib/pkg"
	"reflect"
)

type (
	RegisterPerTypeOption[T any] func(newFunc ValidationFunc[T]) ValidationFunc[T]

	genericValidationFunc = func(any) error
	ValidationFunc[T any] func(T) error

	registryEntry struct {
		baseTFunc   any
		baseAnyFunc ValidationFunc[any]
	}
	ValidatorFuncPerType map[reflect.Type]*registryEntry
)

var (
	// DefaultValidationFunc is used when IsValid[T] does not find a function to validate the given value
	DefaultValidationFunc ValidationFunc[any]

	// validationFuncPerType is used as first option, if no validator is found, try per Kind
	validationFuncPerType = ValidatorFuncPerType{}

	// validationFuncPerKind is used when there is no specific per
	validationFuncPerKind = map[reflect.Kind]genericValidationFunc{}
)

func findTypeFrom[T any]() reflect.Type {
	var target T
	typeOf := reflect.TypeOf(target)
	if typeOf == nil {
		var targetInterface *T
		return reflect.TypeOf(targetInterface)
	}
	return typeOf
}

func findTypeFromVal(val any) reflect.Type {
	typeOf := reflect.TypeOf(val)
	if typeOf == nil {
		panic(Failed(serror.IllegalArgumentValue("val", "nil")))
	}
	return typeOf
}

func findValidationFuncByType(typeOf reflect.Type) (ValidationFunc[any], bool) {
	entryFound, found := validationFuncPerType[typeOf]
	return entryFound.baseAnyFunc, found
}

func findValidationFuncByKind(kind reflect.Kind) (genericValidationFunc, bool) {
	functionFound, found := validationFuncPerKind[kind]
	return functionFound, found
}

// findValidationFunc tries to find a ValidationFunc:
// 1 - check per type T
// 2 - check whether is a pointer and tries pointed type
// 3 - check kind of T
// 4 - returns false
func findValidationFunc[T any]() (ValidationFunc[any], bool) {
	typeOf := findTypeFrom[T]()
	validationFunc, found := findValidationFuncByType(typeOf)
	if !found {
		if typeOf.Kind() == reflect.Ptr {
			// When we wrap the ValidationFunc to any, we convert pointer to value so we can use the element type
			validationFunc, found = findValidationFuncByType(typeOf.Elem())
		}
		if !found {
			validationFunc, found = findValidationFuncByKind(typeOf.Kind())
		}
	}
	return validationFunc, found
}

func assertValueIsOfTheTypeGivenOrPointerToIt(targetType reflect.Type, val any) (any, error) {
	valType := findTypeFromVal(val)
	givenValueTypeIsNotTheTargetType := valType != targetType
	givenValueTypeIsNotPointerToTheTargetType := !(valType.Kind() == reflect.Ptr && valType.Elem() == targetType)

	if givenValueTypeIsNotTheTargetType {
		if givenValueTypeIsNotPointerToTheTargetType {
			return nil, fmt.Errorf("value to be validated is of type=[%s] that is not the type=[%s]",
				valType.String(), targetType.String())
		}

		valOf := reflect.ValueOf(val)
		if err := ValueOfIsNotNil(valOf); err != nil {
			return nil, err // Type(val) is a pointer to T but val is nil
		}
		val = valOf.Elem().Interface() // get the pointed value
	}
	return val, nil
}

func wrapAsValidationFuncOfAny[T any](newFunc ValidationFunc[T], internalErrorDecorator func(error) error) ValidationFunc[any] {
	typeOf := findTypeFrom[T]()
	return func(val any) error {
		assertedValue, err := assertValueIsOfTheTypeGivenOrPointerToIt(typeOf, val)
		if err != nil {
			return internalErrorDecorator(err)
		}
		return newFunc(assertedValue.(T))
	}
}

func (v ValidatorFuncPerType) GetForType(t reflect.Type) ValidationFunc[any] {
	if entryFound, found := v[t]; found {
		return entryFound.baseAnyFunc
	}
	return nil
}

func RegisterPerTypeWith[T any](registry ValidatorFuncPerType, newFunc ValidationFunc[T]) ValidationFunc[T] {
	typeOf := findTypeFrom[T]()
	newEntry := registryEntry{
		baseTFunc: newFunc,
		baseAnyFunc: wrapAsValidationFuncOfAny[T](newFunc, func(err error) error {
			return WrapAsInternalError(err)
		}),
	}
	var oldEntry *registryEntry
	oldEntry, registry[typeOf] = registry[typeOf], &newEntry
	if check.IsNil(oldEntry) {
		return nil
	}
	return oldEntry.baseTFunc.(ValidationFunc[T])
}

func RegisterPerType[T any](newFunc ValidationFunc[T]) ValidationFunc[T] {
	return RegisterPerTypeWith(validationFuncPerType, newFunc)
}

func RegisterPerKind(kind reflect.Kind, newFunc ValidationFunc[any]) (old ValidationFunc[any]) {
	old, validationFuncPerKind[kind] = validationFuncPerKind[reflect.Struct], newFunc
	return old
}
