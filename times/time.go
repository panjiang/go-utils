package times

import (
	"strconv"
	"time"
)

// Layout
const (
	LayoutDateValue = "20060102"
)

// MondayZeroTime returns week's begin time
func MondayZeroTime(t time.Time) time.Time {
	year, month, day := t.Date()
	dayNum := int(t.Weekday())
	if dayNum == 0 {
		dayNum = 7
	}
	interval := time.Hour * 24 * time.Duration(dayNum-1)
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Add(interval * -1)
}

// DateZeroTime returns day's begin time
func DateZeroTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// DateNum returns the date number of specify time
func DateNum(t time.Time) int64 {
	s := t.Format(LayoutDateValue)
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

// TimestampMilli returns integer timestamp rounded to millisecond
func TimestampMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
