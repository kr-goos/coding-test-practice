package repeatedstring

import "strings"

func solution(n int) string {
	pattern := []rune("수박")
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteRune(pattern[i%len(pattern)])
	}
	return sb.String()
}
