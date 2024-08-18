package sumdigits

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	N = [TESTCOUNT]int{
		123,
		987,
	}
	answer = [TESTCOUNT]int{
		6,
		24,
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("N : ", N[i])
		fmt.Println("answer : ", answer[i])
		r := solution(N[i])
		fmt.Println("solution result : ", r)

		if r != answer[i] {
			t.Errorf("testcase [%d] failure", i)
		}
	}
}
