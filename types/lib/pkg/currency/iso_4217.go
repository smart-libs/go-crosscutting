package currency

import (
	"golang.org/x/text/currency"
)

type (
	ISO4217 = currency.Unit
)

var (
	_ Unit = ISO4217{}
)
