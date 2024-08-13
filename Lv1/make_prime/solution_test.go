package makeprime

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	nums = [TESTCOUNT][]int{
		{1, 2, 3, 4},
		{1, 2, 7, 6, 4},
	}
	result = [TESTCOUNT]int{1, 4}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("nums : %v\n", nums[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(nums[i]))
	}
}
