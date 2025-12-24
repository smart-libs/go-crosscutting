package shopspring

import (
	"github.com/leekchan/accounting"
	impl "github.com/shopspring/decimal"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/decimal"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
)

type (
	Metadata struct{}
)

func (m Metadata) WithOptions(d DecimalImpl, options decimal.Options) DecimalImpl {
	var result DecimalImpl
	if options.Precision == decimal.ParserMaxSizeOptions.Precision {
		// use the precision identified
		result = d
		result.options = options
		result.options.Precision = -int(d.wrapped.Exponent())
		return result
	}
	if options.TruncEnabled {
		result.wrapped = d.wrapped.Truncate(int32(options.Precision))
		result.options = options
	} else {
		result.wrapped = d.wrapped.Round(int32(options.Precision))
		result.options = options
	}
	return result
}

func fromOp[T any](m Metadata, i T, from func(T) impl.Decimal, options ...decimal.ParserOptions) (result DecimalImpl) {
	implDec := from(i)
	result.wrapped = implDec
	if len(options) > 0 {
		return m.WithOptions(result, options[0].Options)
	}
	optionWithParsedPrecision := decimal.DefaultOptions
	if implDec.IsInteger() {
		optionWithParsedPrecision.Precision = 0
	} else {
		optionWithParsedPrecision.Precision = -int(implDec.Exponent())
	}
	return m.WithOptions(result, optionWithParsedPrecision)
}

func (m Metadata) FromString(s string, options ...decimal.ParserOptions) (result DecimalImpl, err error) {
	return fromOp(m, s, func(s string) impl.Decimal {
		result.wrapped, err = impl.NewFromString(s)
		return result.wrapped
	}, options...), err
}

func (m Metadata) FromFloat64(f float64, options ...decimal.ParserOptions) (result DecimalImpl) {
	return fromOp(m, f, impl.NewFromFloat, options...)
}

func (m Metadata) FromInt(i int, options ...decimal.ParserOptions) (result DecimalImpl) {
	return fromOp(m, int64(i), impl.NewFromInt, options...)
}

func (m Metadata) IsZero(t DecimalImpl) bool     { return t.wrapped.IsZero() }
func (m Metadata) IsPositive(t DecimalImpl) bool { return t.wrapped.IsPositive() }
func (m Metadata) IsNegative(t DecimalImpl) bool { return t.wrapped.IsNegative() }

func (m Metadata) Abs(t DecimalImpl) DecimalImpl {
	return DecimalImpl{wrapped: t.wrapped.Abs(), options: t.options}
}

func (m Metadata) EqualTo(t DecimalImpl, t2 DecimalImpl) bool {
	return t.wrapped.Equal(t2.wrapped)
}

func (m Metadata) DataTypeName() string { return "DecimalImpl" }

func (m Metadata) IsSet(_ DecimalImpl) bool { return true }

func (m Metadata) IsValid(DecimalImpl) bool { return true }

func (m Metadata) String(t DecimalImpl) string {
	return t.wrapped.StringFixed(int32(t.options.Precision))
}

func (m Metadata) doOpe(d2 DecimalImpl, op func(impl.Decimal) impl.Decimal, options decimal.OperationOptions) DecimalImpl {
	result := DecimalImpl{
		wrapped: op(d2.wrapped).Round(int32(options.ResultOptions.Precision)),
		options: options.ResultOptions,
	}
	return result
}

func resolveOptions(d1Options decimal.Options, opOptions ...decimal.OperationOptions) decimal.OperationOptions {
	if len(opOptions) > 0 {
		return opOptions[0]
	}
	return decimal.OperationOptions{
		ResultOptions: d1Options,
		TruncEnabled:  d1Options.TruncEnabled,
	}
}

func (m Metadata) Sub(d1, d2 DecimalImpl, options ...decimal.OperationOptions) DecimalImpl {
	return m.doOpe(d2, d1.wrapped.Sub, resolveOptions(d1.options, options...))
}

func (m Metadata) Add(d1, d2 DecimalImpl, options ...decimal.OperationOptions) DecimalImpl {
	return m.doOpe(d2, d1.wrapped.Add, resolveOptions(d1.options, options...))
}

func (m Metadata) Divide(d1, d2 DecimalImpl, options ...decimal.OperationOptions) DecimalImpl {
	return m.doOpe(d2, d1.wrapped.Div, resolveOptions(d1.options, options...))
}

func (m Metadata) Multiply(d1, d2 DecimalImpl, options ...decimal.OperationOptions) DecimalImpl {
	return m.doOpe(d2, d1.wrapped.Mul, resolveOptions(d1.options, options...))
}

func (m Metadata) Pow(d1, d2 DecimalImpl, options ...decimal.OperationOptions) DecimalImpl {
	return m.doOpe(d2, d1.wrapped.Pow, resolveOptions(d1.options, options...))
}

func (m Metadata) Format(d DecimalImpl, options decimal.FormatOptions) string {
	ac := accounting.Accounting{
		Symbol:         options.Symbol,
		Precision:      options.Precision,
		Thousand:       options.Thousand,
		Decimal:        options.Decimal,
		Format:         options.Format,
		FormatNegative: options.FormatNegative,
		FormatZero:     options.FormatZero,
	}
	return ac.FormatMoneyDecimal(d.wrapped)
}

var (
	_ metadata.Decimal[DecimalImpl] = Metadata{}
)
