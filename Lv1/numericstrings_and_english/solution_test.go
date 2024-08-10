package numericstringsandenglish

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 4

var (
	s = [TESTCOUNT]string{
		"one4seveneight",
		"23four5six7",
		"2three45sixseven",
		"123",
	}

	result = []int{
		1478,
		234567,
		234567,
		123,
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("===== TEST %d =====\n", i)
		fmt.Printf("s : %v\n", s[i])
		fmt.Printf("answer : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(s[i]))
	}
}
