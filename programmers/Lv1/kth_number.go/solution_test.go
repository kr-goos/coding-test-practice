package kthnumber

import (
	"fmt"
	"testing"
)

var (
	array    = []int{1, 5, 2, 6, 3, 7, 4}
	commands = [][]int{
		{2, 5, 3}, {4, 4, 1}, {1, 7, 3},
	}
	result = []int{5, 6, 3}
)

func TestSolution(t *testing.T) {
	fmt.Printf("result : %v\n", result)
	fmt.Printf("solution result : %v\n", solution(array, commands))
}
