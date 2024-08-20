package gymclothes

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	n    = [TESTCOUNT]int{5, 5, 3}
	lost = [TESTCOUNT][]int{
		{2, 4},
		{2, 4},
		{3},
	}
	reserve = [TESTCOUNT][]int{
		{1, 3, 5},
		{3},
		{1},
	}
	result = [TESTCOUNT]int{5, 4, 2}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("n : %v\n", n[i])
		fmt.Printf("lost : %v\n", lost[i])
		fmt.Printf("reserve : %v\n", reserve[i])
		fmt.Printf("result : %d\n", result[i])
		fmt.Printf("solution result : %d\n", solution(n[i], lost[i], reserve[i]))
	}
}
