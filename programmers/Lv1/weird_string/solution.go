package weirdstring

import (
	"unicode"
)

func solution(s string) string {
	runes := []rune(s)

	var idx int
	for i, r := range runes {
		if r == ' ' {
			idx = 0
			continue
		}

		if idx%2 == 0 {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
		idx++
	}

	return string(runes)
}
