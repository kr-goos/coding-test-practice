package findremainderone

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

func TestSolution(t *testing.T) {

	n := [TESTCOUNT]int{10, 12}
	result := [TESTCOUNT]int{3, 11}
	for i := 0; i < 2; i++ {
		if r := solution(n[i]); result[i] == r {
			fmt.Printf("testcase %d sucecss\n", i+1)
		} else {
			t.Fatalf("testcase %d result : %d, solution result : %d", i+1, result[i], r)
		}
	}

}
