package basicstringhandling

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	s = [TESTCOUNT]string{
		"a234",
		"1234",
	}
	result = [TESTCOUNT]bool{
		false,
		true,
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("s : ", s[i])
		fmt.Println("result : ", result[i])
		answer := solution(s[i])
		fmt.Println("solution result : ", answer)
		if answer != result[i] {
			t.Errorf("[testcase %d] failure", i)
		}
	}
}
