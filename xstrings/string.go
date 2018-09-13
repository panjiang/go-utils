package xstrings

import (
	"unicode"
)

// LowerFirst converts the first char to lower
func LowerFirst(s string) string {
	ss := []rune(s)
	ss[0] = unicode.ToLower(ss[0])
	return string(ss)
}
