package utils

import "time"

//Function GetStartOfDay returns the first millisecond of the given date
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
