package metadata

import "time"

type (
	TimestampDefaultStringer struct{}

	Timestamp struct {
		IsSetDefaultValue[time.Time]
		TimestampDefaultStringer
		AlwaysIsValid[time.Time]
		ComparableEqualTo[time.Time]
	}

	OrdinaryTimestamp struct {
		Timestamp
	}
)

func (s TimestampDefaultStringer) String(t time.Time) string { return t.Format(time.RFC3339Nano) }

func (u OrdinaryTimestamp) DataTypeName() string { return "time.Time" }

var (
	_ Metadata[time.Time] = OrdinaryTimestamp{}
)
