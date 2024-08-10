package numericstringsandenglish

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(s string) int {

	for i, v := range initEnglishArray() {
		idx := strings.Index(s, v)
		if idx == -1 {
			continue
		}
		s = strings.ReplaceAll(s, v, fmt.Sprintf("%d", i))
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return v
}

func initEnglishArray() []string {
	return []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
}
