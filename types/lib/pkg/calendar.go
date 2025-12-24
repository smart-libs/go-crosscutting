package types

type (
	Calendar struct{}
)

var (
	DefaultCalendar Calendar
)

func (c Calendar) GetFirstDateOfYear(year int) Date {
	return NewDateFromDay(year, 1, 1)
}

func (c Calendar) GetFirstDateOfMonth(d Date) Date {
	myTime := d.Time()
	return NewDateFromDay(myTime.Year(), int(myTime.Month()), 1)
}

func (c Calendar) GetNumOfMonthsBetween(start, end Date) int {
	// Extract the year and month for both dates
	startYear, startMonth, _ := start.Time().Date()
	endYear, endMonth, _ := end.Time().Date()

	// Calculate the differences in years and months
	yearDiff := endYear - startYear
	monthDiff := int(endMonth - startMonth)

	// Total months difference
	totalMonths := yearDiff*12 + monthDiff

	return totalMonths
}

func (c Calendar) GetLastDateOfYear(year int) Date {
	return NewDateFromDay(year, 12, 31)
}

func (c Calendar) GetLastDateOfMonth(d Date) Date {
	myTime := d.Time()
	firstDayOfMonth := NewDateFromDay(myTime.Year(), int(myTime.Month()), 1).Time()
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	return NewDateFromTime(lastDayOfMonth)
}
