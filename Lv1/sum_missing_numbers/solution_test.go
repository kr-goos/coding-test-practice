package summissingnumbers

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

func TestSolution(t *testing.T) {
	numbers := [TESTCOUNT][]int{
		{1, 2, 3, 4, 6, 7, 8, 0},
		{5, 8, 4, 0, 6, 7, 9},
	}
	result := [TESTCOUNT]int{14, 6}

	for i := 0; i < TESTCOUNT; i++ {
		v := solution(numbers[i])

		if v == result[i] {
			fmt.Printf("testcase %d sucecss\n", i+1)
		} else {
			t.Fatalf("testcase %d result : %d, solution result : %d", i+1, result[i], v)
		}
	}
}
