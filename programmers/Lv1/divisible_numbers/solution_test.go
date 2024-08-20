package divisiblenumbers

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	arr = [TESTCOUNT][]int{
		{5, 9, 7, 10},
		{2, 36, 1, 3},
		{3, 2, 6},
	}
	divisor = [TESTCOUNT]int{5, 1, 10}
	result  = [TESTCOUNT][]int{
		{5, 10},
		{1, 2, 3, 36},
		{-1},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("arr : ", arr[i])
		fmt.Println("divisor : ", divisor[i])
		fmt.Println("result : ", result[i])
		fmt.Println("solution result : ", solution(arr[i], divisor[i]))
	}
}
