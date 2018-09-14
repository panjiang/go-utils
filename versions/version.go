package versions

import (
	"regexp"
	"strconv"
	"strings"
)

var reg = regexp.MustCompile("^\\d+[\\.\\d]*$")

// Valid checks whether matchs version pattern
func Valid(version string) bool {
	return reg.FindString(version) != ""
}

// Cmp compares two version
// 0: v1 == v2
// 1: v1 > v2
// -1: v1 < v2
// Attention: 1 > 0.1, 1.1 > 1.0.1, 1.1.0 == 1.1
func Cmp(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}

	varr1 := strings.Split(v1, ".")
	varr2 := strings.Split(v2, ".")
	maxLen := len(varr1)
	if len(varr2) > len(varr1) {
		maxLen = len(varr2)
	}
	for i := 0; i < maxLen; i++ {
		var part1 int64
		var part2 int64
		if len(varr1) >= i+1 {
			part1, _ = strconv.ParseInt(varr1[i], 10, 64)
		}
		if len(varr2) >= i+1 {
			part2, _ = strconv.ParseInt(varr2[i], 10, 64)
		}
		if part1 > part2 {
			return 1
		} else if part1 < part2 {
			return -1
		}
	}

	return 0
}
