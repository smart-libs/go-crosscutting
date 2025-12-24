package types

import (
	"encoding/json"
	"fmt"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
	"strings"
	"time"
)

type (
	DateTemplate[T metadata.Metadata[time.Time]] struct {
		DataTypeTemplate[time.Time, T]
	}
)

func (d DateTemplate[T]) Time() time.Time { return d.value }

func (d DateTemplate[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func dateFromString(s string) (t time.Time, err error) {
	if strings.Contains(s, "-") {
		if len(s) > 10 {
			// truncate date time values
			s = s[:10]
		}
		t, err = time.Parse("2006-01-02", s)
	} else {
		if len(s) == 6 {
			t, err = time.Parse("060102", s)
		} else {
			t, err = time.Parse("20060102", s)
		}
	}
	if err == nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	}
	return
}

func (d *DateTemplate[T]) UnmarshalJSON(b []byte) error {
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
	d.value, err = dateFromString(s)
	return err
}

func NewDateTemplate[T metadata.Metadata[time.Time]](y int, m time.Month, d int) DateTemplate[T] {
	t := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	return DateTemplate[T]{
		DataTypeTemplate: DataTypeTemplate[time.Time, T]{value: t},
	}
}

func NewDateTemplateFromTime[T metadata.Metadata[time.Time]](t time.Time) DateTemplate[T] {
	return NewDateTemplate[T](t.Year(), t.Month(), t.Day())
}

func NewDateTemplateFromString[T metadata.Metadata[time.Time]](s string) (t DateTemplate[T], err error) {
	t.value, err = dateFromString(s)
	return
}
