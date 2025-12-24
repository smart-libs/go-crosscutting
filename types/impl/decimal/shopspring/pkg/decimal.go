package shopspring

import (
	"encoding/json"
	impl "github.com/shopspring/decimal"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/decimal"
)

type (
	DecimalImpl struct {
		wrapped impl.Decimal
		options decimal.Options
	}
)

func (c *DecimalImpl) UnmarshalJSON(buf []byte) (err error) {
	if err = json.Unmarshal(buf, &c.wrapped); err == nil {
		c.options.Precision = -int(c.wrapped.Exponent())
		c.options.TruncEnabled = decimal.DefaultOptions.TruncEnabled
	}
	return
}

func (c DecimalImpl) MarshalJSON() (buf []byte, err error) {
	return json.Marshal(c.wrapped)
}
