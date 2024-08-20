package sortstringdescending

import "sort"

func solution(s string) string {
	b := []byte(s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if b[i] < b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return string(b)
}

func simpleSolution(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	return string(b)
}
