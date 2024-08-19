package middlechar

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	s = [TESTCOUNT]string{
		"abcde",
		"qwer",
	}

	result = [TESTCOUNT]string{
		"c",
		"we",
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("s : ", s[i])
		fmt.Println("result : ", result[i])
		fmt.Println("solution result : ", solution(s[i]))
	}
}
