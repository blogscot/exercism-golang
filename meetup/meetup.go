package meetup

import "time"

// WeekSchedule specifies the week in a month
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func lengthOfMonth(year int, month time.Month) int {
	date := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return date.Day()
}

// Day is
func Day(schedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	var day int

	switch schedule {
	case First:
		day = 1
	case Second:
		day = 8
	case Third:
		day = 15
	case Fourth:
		day = 22
	case Teenth:
		day = 13
	case Last:
		day = lengthOfMonth(year, month) - 6
	}

	d := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	for d.Weekday() != weekday {
		d = d.AddDate(0, 0, 1)
	}
	return d.Day()
}
