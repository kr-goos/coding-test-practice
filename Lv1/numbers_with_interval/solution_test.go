package numberswithinterval

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	x      = [TESTCOUNT]int{2, 4, -4}
	n      = [TESTCOUNT]int{5, 3, 2}
	answer = [TESTCOUNT][]int{
		{2, 4, 6, 8, 10},
		{4, 8, 12},
		{-4, -8},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("x : %d\n", x[i])
		fmt.Printf("n : %d\n", n[i])
		fmt.Printf("answer : %v\n", answer[i])
		fmt.Printf("solution answer : %v\n", solution(x[i], n[i]))
	}
}
