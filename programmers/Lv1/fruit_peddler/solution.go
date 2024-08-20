package fruitpeddler

import (
	"sort"
)

func solution(k int, m int, score []int) int {

	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	var profit int

	for i := 0; i+m <= len(score); i += m {
		box := score[i : i+m]
		profit += box[m-1] * m
	}

	return profit
}
