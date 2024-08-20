package removesmallestnumber

import (
	"sort"
)

func solution(arr []int) []int {

	if len(arr) == 0 || len(arr) == 1 {
		return []int{-1}
	}

	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	for i, v := range arr {
		if v == min {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	return arr
}

func solution2(arr []int) []int {

	if len(arr) == 0 || len(arr) == 1 {
		return []int{-1}
	}

	temp := make([]int, len(arr))
	copy(temp, arr)

	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j]
	})

	for i, v := range arr {
		if v == temp[len(temp)-1] {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}

	}

	return arr
}
