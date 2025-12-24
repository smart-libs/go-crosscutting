package types

import (
	"fmt"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
	"time"
)

type (
	Date struct {
		DateTemplate[metadata.OrdinaryDate]
	}
)

func (d Date) GoString() string { return d.String() }

func DateNotSet() Date { return Date{} }
func Today() Date      { return NewDateFromTime(time.Now()) }

func NewDateFromTime(t time.Time) Date {
	return Date{
		DateTemplate: NewDateTemplate[metadata.OrdinaryDate](t.Year(), t.Month(), t.Day()),
	}
}

func LastDateOfMonth(y, m int) Date {
	return NewDateFromTime(
		time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC).
			AddDate(0, 1, -1),
	)
}

func NewDateFromDay(y, m, d int) Date {
	return Date{
		DateTemplate: NewDateTemplate[metadata.OrdinaryDate](y, time.Month(m), d),
	}
}

func NewDateFromString(s string) (t Date, err error) {
	t.value, err = dateFromString(s)
	return
}

var _ fmt.GoStringer = Date{}
