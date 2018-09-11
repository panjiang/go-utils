package times

import (
	"testing"
	"time"
)

func TestMondayZeroTime(t *testing.T) {
	var tt time.Time

	tt = time.Now()
	t.Log(tt.Weekday(), int(tt.Weekday()))
	t.Log(MondayZeroTime(tt))

	tt = time.Now().Add(-1 * 24 * time.Hour * 2)
	t.Log(tt.Weekday(), int(tt.Weekday()))
	t.Log(MondayZeroTime(tt))
}
