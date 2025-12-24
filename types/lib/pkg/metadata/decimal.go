package metadata

import (
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/decimal"
)

type (
	Decimal[T any] interface {
		Metadata[T]
		FromString(s string, options ...decimal.ParserOptions) (T, error)
		FromFloat64(f float64, options ...decimal.ParserOptions) T
		FromInt(i int, options ...decimal.ParserOptions) T
		Sub(d1, d2 T, options ...decimal.OperationOptions) T
		Add(d1, d2 T, options ...decimal.OperationOptions) T
		Divide(d1, d2 T, options ...decimal.OperationOptions) T
		Multiply(d1, d2 T, options ...decimal.OperationOptions) T
		Pow(d1, d2 T, options ...decimal.OperationOptions) T
		IsZero(d T) bool
		IsPositive(d T) bool
		IsNegative(d T) bool
		WithOptions(T, decimal.Options) T
		Abs(d T) T
		Format(T, decimal.FormatOptions) string
	}
)
