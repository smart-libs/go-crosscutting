package types

import "fmt"

type (
	DataType[T any] interface {
		EqualTo(T) bool
		IsSet() bool
		IsValid() bool
		fmt.Stringer
	}
)
