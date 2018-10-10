package slices

import (
	"testing"
)

func TestMapToStr(t *testing.T) {
	checkError := func(err error) {
		if err != nil {
			t.Error("error:", err)
		}
	}

	a := []int{1, 2, 3}
	b := []int64{1, 2, 3}
	c := []float64{1.1, 2.2, 3.3}

	var err error
	var z []string

	z, err = MapToStr(a)
	checkError(err)
	t.Logf("%q", z)

	z, err = MapToStr(b)
	checkError(err)
	t.Logf("%q", z)

	z, err = MapToStr(c)
	checkError(err)
	t.Logf("%q", z)

	z, err = MapToStr(1)
	checkError(err)
	t.Logf("%q", z)
}
