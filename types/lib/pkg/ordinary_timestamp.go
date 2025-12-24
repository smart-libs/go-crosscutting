package types

import (
	"fmt"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
	"time"
)

type (
	TimestampDataTypeTemplate[T metadata.Metadata[time.Time]] struct {
		DataTypeTemplate[time.Time, T]
	}

	OrdinaryTimestamp struct {
		TimestampDataTypeTemplate[metadata.OrdinaryTimestamp]
	}
)

var (
	_ DataType[time.Time] = OrdinaryTimestamp{}
)

func (d TimestampDataTypeTemplate[T]) Time() time.Time { return d.value }

func FromTimestamp[T metadata.Metadata[time.Time]](t TimestampDataTypeTemplate[T]) TimestampDataTypeTemplate[T] {
	return TimestampDataTypeTemplate[T]{
		DataTypeTemplate: DataTypeTemplate[time.Time, T]{value: t.value},
	}
}

func NewTimestamp[T metadata.Metadata[time.Time]](t time.Time) TimestampDataTypeTemplate[T] {
	return TimestampDataTypeTemplate[T]{
		DataTypeTemplate: DataTypeTemplate[time.Time, T]{value: t},
	}
}

func NewOrdinaryTimestamp(t time.Time) OrdinaryTimestamp {
	return OrdinaryTimestamp{
		TimestampDataTypeTemplate: NewTimestamp[metadata.OrdinaryTimestamp](t),
	}
}

func NewOrdinaryTimestampFromString(s string) (t time.Time, err error) {
	const RFC3339Milli = "2006-01-02T15:04:05.999"
	if len(s) > len(time.RFC3339) {
		t, err = time.Parse(time.RFC3339Nano, s)
	} else {
		if len(s) > len(RFC3339Milli) {
			t, err = time.Parse(time.RFC3339, s)
		} else {
			t, err = time.ParseInLocation(RFC3339Milli, s, time.Local)
		}
	}
	return
}

func (d *TimestampDataTypeTemplate[T]) UnmarshalJSON(b []byte) error {
	const quote = byte('"')
	size := len(b)
	if size == 0 || string(b) == "null" {
		return nil
	}
	if size < 8 || b[0] != quote || b[size-1] != quote {
		return fmt.Errorf("invalid JSON data=[%s], it must be sorrounded by quotes", string(b))
	}
	var (
		s   = string(b[1 : size-1])
		err error
	)
	d.value, err = NewOrdinaryTimestampFromString(s)
	return err
}
