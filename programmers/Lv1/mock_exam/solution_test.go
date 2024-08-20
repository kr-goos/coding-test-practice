package mockexam

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	answers = [TESTCOUNT][]int{
		{1, 2, 3, 4, 5},
		{1, 3, 2, 4, 2},
	}

	result = [TESTCOUNT][]int{
		{1},
		{1, 2, 3},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("answers : %v\n", answers[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(answers[i]))
	}
}
