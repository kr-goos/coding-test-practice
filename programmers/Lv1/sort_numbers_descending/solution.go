package sortnumbersdescending

import (
	"sort"
	"strconv"
)

func solution(n int64) int64 {
	s := []byte(strconv.Itoa(int(n)))
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	v, _ := strconv.ParseInt(string(s), 10, 64)
	return v
}
