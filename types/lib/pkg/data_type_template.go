package types

import (
	"encoding/json"
	"fmt"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
)

type (
	DataTypeTemplate[T any, M metadata.Metadata[T]] struct {
		value    T
		metadata M
	}
)

func (d DataTypeTemplate[T, M]) PanicIfNotSet() {
	if err := d.ErrorIfNotSet(); err != nil {
		panic(err)
	}
}

func (d DataTypeTemplate[T, M]) ErrorIfNotSet() error {
	if !d.metadata.IsSet(d.value) {
		return fmt.Errorf("%s is not set", d.metadata.DataTypeName())
	}
	return nil
}

func (d DataTypeTemplate[T, M]) EqualTo(t T) bool { return d.metadata.EqualTo(t, d.value) }

func (d DataTypeTemplate[T, M]) IsValid() bool { return d.metadata.IsValid(d.value) }

func (d DataTypeTemplate[T, M]) IsSet() bool { return d.metadata.IsSet(d.value) }

func (d DataTypeTemplate[T, M]) String() string { return d.metadata.String(d.value) }

func (d DataTypeTemplate[T, M]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}
