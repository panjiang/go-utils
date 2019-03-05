package slices

import (
	"errors"
	"fmt"
	"reflect"
)

// ContainsStr whether cotains certain string
func ContainsStr(list []string, s string) bool {
	for _, ss := range list {
		if ss == s {
			return true
		}
	}
	return false
}

// ContainsInt64 whether cotains certain int64
func ContainsInt64(list []int64, s int64) bool {
	for _, ss := range list {
		if ss == s {
			return true
		}
	}
	return false
}

// ContainsInt whether cotains certain int
func ContainsInt(list []int, s int) bool {
	for _, ss := range list {
		if ss == s {
			return true
		}
	}
	return false
}

// MapToStr converts any kind slice to string slice
func MapToStr(slice interface{}) (sliceStr []string, err error) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		err = errors.New("must be slice")
		return
	}

	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		sliceStr = append(sliceStr, fmt.Sprint(s.Index(i)))
	}
	return
}
