package times

import (
	"strconv"
	"time"
)

// Layout
const (
	LayoutDateValue = "20060102"
)

// DateTimeZero returns zero o'clock of the datetime
func DateTimeZero(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// DateNum returns the date number of specify time
func DateNum(t time.Time) int64 {
	s := t.Format(LayoutDateValue)
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}
