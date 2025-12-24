package decimal

import "math"

type (
	// Options to configure a decimal value
	Options struct {
		// Precision number of decimals
		Precision int `json:"precision,omitempty"`
		// TruncEnabled enable or disable the truncation of the decimal
		TruncEnabled   bool `json:"trunc_enabled,omitempty"`
		RoundingMethod `json:"rounding_method,omitempty"`
	}

	// ParserOptions provides options for the decimal parser
	ParserOptions struct {
		Options
	}

	// OperationOptions specifies the options of the operation result and also the truncation behavior
	OperationOptions struct {
		// ResultOptions are the options of the decimal resulted from the operation
		ResultOptions Options
		// TruncEnabled enable or disable the truncation of the resulted decimal for the operation only
		TruncEnabled bool
	}
)

var (
	// DefaultOptions for new decimals
	DefaultOptions = Options{
		Precision:    2,
		TruncEnabled: false,
	}

	// ParserMaxSizeOptions for new decimals when you want to use the number os decimals given in the string value
	ParserMaxSizeOptions = ParserOptions{Options{
		Precision:    math.MaxInt64,
		TruncEnabled: false,
	}}
)
