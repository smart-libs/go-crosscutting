package metadata

import "time"

type (
	DateDefaultStringer struct{}

	Date struct {
		IsSetDefaultValue[time.Time]
		DefaultValueFor[time.Time]
		ComparableEqualTo[time.Time]
		IsValidIfIsSet[time.Time]
		DateDefaultStringer
	}

	OrdinaryDate struct {
		Date
	}
)

func (o OrdinaryDate) DataTypeName() string { return "date" }

func (d DateDefaultStringer) String(t time.Time) string { return t.Format("2006-01-02") }

var (
	_ Metadata[time.Time] = OrdinaryDate{}
)
