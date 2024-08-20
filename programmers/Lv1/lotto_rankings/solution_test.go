package lottorankings

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	lottos = [TESTCOUNT][]int{
		{44, 1, 0, 0, 31, 25},
		{0, 0, 0, 0, 0, 0},
		{45, 4, 35, 20, 3, 9},
	}
	win_nums = [TESTCOUNT][]int{
		{31, 10, 45, 1, 6, 19},
		{38, 19, 20, 40, 15, 25},
		{20, 9, 3, 45, 4, 35},
	}
	result = [TESTCOUNT][]int{
		{3, 5},
		{1, 6},
		{1, 1},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("lottos : %v\n", lottos[i])
		fmt.Printf("win_nums : %v\n", win_nums[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(lottos[i], win_nums[i]))
	}
}
