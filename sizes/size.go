package sizes

import (
	"strconv"
)

var units = map[byte]int64{
	'B': 1,
	'K': 1024,
	'M': 1024 * 1024,
	'G': 1024 * 1024 * 1024,
	'T': 1024 * 1024 * 1024 * 1024,
}

// ToByte (T/G/M/K to B)
func ToByte(s string) (int64, error) {
	for u, m := range units {
		if s[len(s)-1] != u {
			continue
		}
		v, err := strconv.ParseInt(s[:len(s)-1], 10, 64)
		if err != nil {
			return 0, err
		}
		return v * m, nil
	}

	return strconv.ParseInt(s, 10, 64)
}
