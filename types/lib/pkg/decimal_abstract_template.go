package types

import (
	"encoding/json"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/decimal"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
)

type (
	// DecimalAbstractTemplate doesn't even know the decimal implementation
	DecimalAbstractTemplate[T any, M metadata.Decimal[T]] struct {
		DataTypeTemplate[T, M]
	}
)

func (d DecimalAbstractTemplate[T, M]) IsZero() bool {
	return d.metadata.IsZero(d.value)
}

func (d DecimalAbstractTemplate[T, M]) IsPositive() bool {
	return d.metadata.IsPositive(d.value)
}

func (d DecimalAbstractTemplate[T, M]) IsNegative() bool {
	return d.metadata.IsNegative(d.value)
}

func (d *DecimalAbstractTemplate[T, M]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &d.value)
}

func (d DecimalAbstractTemplate[T, M]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

func FromAbstractDecimalToDecimal[T any, M metadata.Decimal[T], O metadata.Decimal[T]](i DecimalAbstractTemplate[T, M]) (result DecimalAbstractTemplate[T, O]) {
	result.value = i.value
	return result
}

func AbstractDecimalFromString[T any, M metadata.Decimal[T]](s string, options ...decimal.ParserOptions) (result DecimalAbstractTemplate[T, M], err error) {
	var m M
	result.value, err = m.FromString(s, options...)
	return
}

func AbstractDecimalFromFloat64[T any, M metadata.Decimal[T]](f float64, options ...decimal.ParserOptions) (result DecimalAbstractTemplate[T, M]) {
	var m M
	result.value = m.FromFloat64(f, options...)
	return
}
