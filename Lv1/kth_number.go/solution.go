package kthnumber

import "sort"

func solution(array []int, commands [][]int) []int {

	var answer []int
	for _, cmd := range commands {
		var temp []int
		temp = append(temp, array[cmd[0]-1:cmd[1]]...)
		sort.Ints(temp)
		answer = append(answer, temp[cmd[2]-1])
	}
	return answer
}
