package types

import "fmt"

type (
	DateInterval struct {
		StartDate Date
		EndDate   Date
	}

	DateIntervalGetter interface {
		GetStartDate() Date
		GetEndDate() Date
		GetDateInterval() DateInterval
	}
)

func (i DateInterval) GetStartDate() Date            { return i.StartDate }
func (i DateInterval) GetEndDate() Date              { return i.EndDate }
func (i DateInterval) GetDateInterval() DateInterval { return i }

func (i DateInterval) IsSet() bool {
	return i.StartDate.IsSet() && i.EndDate.IsSet()
}

func (i DateInterval) String() string {
	return fmt.Sprintf("{StartDate: %s, EndDate: %s}", i.StartDate, i.EndDate)
}

func (i DateInterval) IsValid() bool {
	return i.StartDate.IsValid() && i.EndDate.IsValid() && i.StartDate.Time().Compare(i.EndDate.Time()) <= 0
}

func (i DateInterval) DateBelongsToInterval(d Date) bool {
	return !d.Time().Before(i.StartDate.Time()) && !d.Time().After(i.EndDate.Time())
}

var (
	_ DateIntervalGetter = DateInterval{}
)
