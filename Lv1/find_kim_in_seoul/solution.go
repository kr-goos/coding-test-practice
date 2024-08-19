package findkiminseoul

import "fmt"

func solution(seoul []string) string {
	var idx int
	for i := range seoul {
		if seoul[i] == "Kim" {
			idx = i
			break
		}
	}
	return fmt.Sprintf("김서방은 %d에 있다", idx)
}
