package metadata

import "reflect"

type (
	DataTypeName interface {
		DataTypeName() string
	}

	IsValid[T any] interface {
		IsValid(T) bool
	}

	EqualTo[T any] interface {
		EqualTo(T, T) bool
	}

	DefaultValue[T any] interface {
		DefaultValue() T
	}

	IsSetValidation[T any] interface {
		IsSet(T) bool
	}

	Stringer[T any] interface {
		String(T) string
	}

	Metadata[T any] interface {
		DataTypeName
		EqualTo[T]
		IsSetValidation[T]
		IsValid[T]
		Stringer[T]
	}

	AlwaysIsValid[T any]            struct{}
	IsValidIfIsSet[T comparable]    struct{}
	ComparableEqualTo[T comparable] struct{}
	DeepEqualTo[T any]              struct{}
	DefaultValueFor[T any]          struct{}
	IsSetDefaultValue[T comparable] struct{}
)

func (e AlwaysIsValid[T]) IsValid(_ T) bool {
	return true
}

func (e ComparableEqualTo[T]) EqualTo(a, b T) bool {
	return a == b
}

func (e DeepEqualTo[T]) EqualTo(a, b T) bool {
	return reflect.DeepEqual(a, b)
}

func (d DefaultValueFor[T]) DefaultValue() T {
	var t T
	return t
}

func (d IsSetDefaultValue[T]) IsSet(v T) bool {
	defaultValue := DefaultValueFor[T]{}.DefaultValue()
	return v != defaultValue
}

func (d IsValidIfIsSet[T]) IsValid(v T) bool {
	defaultValue := DefaultValueFor[T]{}.DefaultValue()
	return v != defaultValue
}
