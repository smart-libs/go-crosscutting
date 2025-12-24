package decimal

type (
	FormatOptions struct {
		Symbol         string // currency symbol (required)
		Precision      int    // currency precision (decimal places) (optional / default: 0)
		Thousand       string // thousand separator (optional / default: ,)
		Decimal        string // decimal separator (optional / default: .)
		Format         string // simple format string allows control of symbol position (%v = value, %s = symbol) (default: %s%v)
		FormatNegative string // format string for negative values (optional / default: strings.Replace(strings.Replace(accounting.Format, "-", "", -1), "%v", "-%v", -1))
		FormatZero     string // format string for zero values (optional / default: Format)
	}
)
