package types

import (
	"github.com/smart-libs/go-crosscutting/types/impl/decimal/shopspring/pkg"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/decimal"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
)

type (
	// DecimalTemplate fixes a decimal implementation but allows metadata extensions
	DecimalTemplate[M metadata.Decimal[shopspring.DecimalImpl]] struct {
		DecimalAbstractTemplate[shopspring.DecimalImpl, M]
	}
)

func (d DecimalTemplate[M]) WithOptions(options decimal.Options) (result DecimalTemplate[M]) {
	result.value = d.metadata.WithOptions(d.value, options)
	return result
}

func (d DecimalTemplate[M]) Ptr() *DecimalTemplate[M] {
	return &d
}

func (d DecimalTemplate[M]) EqualTo(o DecimalTemplate[M]) bool {
	return d.metadata.EqualTo(d.value, o.value)
}

func (d DecimalTemplate[M]) Add(o DecimalTemplate[M], options ...decimal.OperationOptions) (result DecimalTemplate[M]) {
	result.value = d.metadata.Add(d.value, o.value, options...)
	return result
}

func (d DecimalTemplate[M]) Sub(o DecimalTemplate[M], options ...decimal.OperationOptions) (result DecimalTemplate[M]) {
	result.value = d.metadata.Sub(d.value, o.value, options...)
	return result
}

func (d DecimalTemplate[M]) Divide(o DecimalTemplate[M], options ...decimal.OperationOptions) (result DecimalTemplate[M]) {
	result.value = d.metadata.Divide(d.value, o.value, options...)
	return result
}

func (d DecimalTemplate[M]) Multiply(o DecimalTemplate[M], options ...decimal.OperationOptions) (result DecimalTemplate[M]) {
	result.value = d.metadata.Multiply(d.value, o.value, options...)
	return result
}

func (d DecimalTemplate[M]) Pow(o DecimalTemplate[M], options ...decimal.OperationOptions) (result DecimalTemplate[M]) {
	result.value = d.metadata.Pow(d.value, o.value, options...)
	return result
}

func (d DecimalTemplate[M]) String() string {
	return d.metadata.String(d.value)
}

func (d DecimalTemplate[M]) Abs() (result DecimalTemplate[M]) {
	result.value = d.metadata.Abs(d.value)
	return result
}

func (d DecimalTemplate[M]) Format(options decimal.FormatOptions) string {
	return d.metadata.Format(d.value, options)
}
