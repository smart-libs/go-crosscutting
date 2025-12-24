package currency

import (
	"cmp"
	"fmt"
)

const UnknownID = ""

type (
	// Serializable has the minimum data to allow a currency to be inflated with its own GO data type
	// This struct should be also comparable which means all its fields must be comparable.
	Serializable[T fmt.Stringer] struct {
		ID   string `json:"id"`
		Type T      `json:"type"`
	}
)

func (s Serializable[T]) GoString() string { return s.ID }
func (s Serializable[T]) String() string   { return s.ID }

var (
	_ fmt.Stringer = unknownSerializableType{}

	// To force Serializable to be a Currency
	//_ Currency = Serializable[unknownSerializableType]{}
	// To force Serializable to be comparable
	_ = cmp.Or(Serializable[unknownSerializableType]{})
)
