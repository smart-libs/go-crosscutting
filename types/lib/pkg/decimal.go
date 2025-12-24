package types

import (
	"github.com/smart-libs/go-crosscutting/types/impl/decimal/shopspring/pkg"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/decimal"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
)

type (
	// Decimal is the default decimal implementation
	Decimal = DecimalTemplate[shopspring.Metadata]
)

var (
	// DecimalZero to be used as constant to return zero as decimal
	DecimalZero   Decimal
	Decimal100    Decimal = DecimalFromInt(100)
	DecimalMinus1 Decimal = DecimalFromInt(-1)
	DecimalOne    Decimal = DecimalFromInt(1)
)

func DecimalFromDecimal[O metadata.Decimal[shopspring.DecimalImpl]](i Decimal) (result DecimalTemplate[O]) {
	result.value = i.value
	return result
}

func DecimalFromString(s string, options ...decimal.ParserOptions) (result Decimal, err error) {
	result.value, err = result.metadata.FromString(s, options...)
	return
}

func DecimalFromStringOrPanic(s string, options ...decimal.ParserOptions) Decimal {
	value, err := DecimalFromString(s, options...)
	if err != nil {
		panic(err)
	}
	return value
}

func DecimalFromFloat64(f float64, options ...decimal.ParserOptions) (result Decimal) {
	result.value = result.metadata.FromFloat64(f, options...)
	return
}

func DecimalFromInt(i int, options ...decimal.ParserOptions) (result Decimal) {
	result.value = result.metadata.FromInt(i, options...)
	return
}
