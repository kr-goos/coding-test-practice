package slpitstrings

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	const TESTCOUNT = 3
	sTestcases := [TESTCOUNT]string{
		"banana",
		"abracadabra",
		"aaabbaccccabba",
	}

	successResults := [TESTCOUNT]int{
		3,
		6,
		3,
	}
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("s testcase : ", sTestcases[i])
		fmt.Println("success result : ", successResults[i])
		fmt.Println("result : ", solution(sTestcases[i]))
	}

}
