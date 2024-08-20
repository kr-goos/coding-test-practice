package caesarcipher

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	s = [TESTCOUNT]string{
		"AB",
		"z",
		"a B z",
	}

	n      = [TESTCOUNT]int{1, 1, 4}
	result = [TESTCOUNT]string{"BC", "a", "e F d"}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("s : ", s[i])
		fmt.Println("n : ", n[i])
		fmt.Println("result : ", result[i])
		fmt.Println("solution result : ", solution(s[i], n[i]))
	}
}
