package slices

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
