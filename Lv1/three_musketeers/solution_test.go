package threemusketeers

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	number = [TESTCOUNT][]int{
		{-2, 3, 0, 2, -5},
		{-3, -2, -1, 0, 1, 2, 3},
		{-1, 1, -1, 1},
	}
	result = [TESTCOUNT]int{2, 5, 0}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("number[i] : ", number[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("solution result : ", solution(number[i]))
	}
}
