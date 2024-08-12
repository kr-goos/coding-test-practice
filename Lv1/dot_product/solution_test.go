package dotproduct

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	a = [TESTCOUNT][]int{
		{1, 2, 3, 4},
		{-1, 0, 1},
	}

	b = [TESTCOUNT][]int{
		{-3, -1, 0, 2},
		{1, 0, -1},
	}
	result = [TESTCOUNT]int{3, -2}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("a : %v\n", a[i])
		fmt.Printf("b : %v\n", b[i])
		fmt.Printf("result : %d\n", result[i])
		fmt.Printf("solution result : %d\n", solution(a[i], b[i]))
	}
}
