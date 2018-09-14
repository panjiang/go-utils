package sizes

import (
	"fmt"
	"strconv"

	"github.com/panjiang/go-utils/maths"
)

var toFixed = maths.ToFixedString

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

// ToString 1.2K 1.2M 1.2G 1.2T
func ToString(v int64) string {
	f := float64(v)
	if f/1024 < 1 {
		return fmt.Sprintf("%dB", v)
	}

	f = float64(f) / 1024
	if f/1024 < 1 {
		return fmt.Sprintf("%sK", toFixed(f, 1))
	}

	f = float64(f) / 1024
	if f/1024 < 1 {
		return fmt.Sprintf("%sM", toFixed(f, 1))
	}

	f = float64(f) / 1024
	if f/1024 < 1 {
		return fmt.Sprintf("%sG", toFixed(f, 1))
	}

	f = float64(f) / 1024
	return fmt.Sprintf("%sT", toFixed(f, 1))
}
