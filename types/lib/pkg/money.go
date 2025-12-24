package types

import (
	"fmt"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/currency"
)

type (
	// Money TODO
	Money interface {
		Amount() Decimal
		Currency() currency.Unit
		fmt.Stringer
	}

	defaultMoney struct {
		amount   Decimal
		currency currency.Unit
	}
)

func (d defaultMoney) Amount() Decimal         { return d.amount }
func (d defaultMoney) Currency() currency.Unit { return d.currency }
func (d defaultMoney) String() string          { return fmt.Sprintf("%s%s", d.currency.String(), d.currency) }

func NewMoney(amount Decimal, currency currency.Unit) Money {
	return defaultMoney{amount: amount, currency: currency}
}
