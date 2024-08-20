package nearestduplicatecharacter

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	const TESTCOUNT = 2
	sTestcases := [TESTCOUNT]string{
		"banana",
		"foobar",
	}

	successResults := [TESTCOUNT][]int{
		{-1, -1, -1, 2, 2, 2},
		{-1, -1, 1, -1, -1, -1},
	}

	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("testcase s : ", sTestcases[i])
		fmt.Println("success result : ", successResults[i])
		fmt.Println("solution result : ", solution(sTestcases[i]))
	}

}
