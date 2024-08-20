package divisiblenumbers

import "sort"

func solution(arr []int, divisor int) []int {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var answer []int
	for _, element := range arr {
		if element%divisor == 0 {
			answer = append(answer, element)
		}
	}

	if answer == nil {
		return []int{-1}
	}

	return answer
}
